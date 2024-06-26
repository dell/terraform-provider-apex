/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://mozilla.org/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"fmt"
	"net/http"

	jmsClient "dell/apex-job-client"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/constants"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &mobilityGroupsCopyResource{}
	_ resource.ResourceWithConfigure = &mobilityGroupsCopyResource{}
)

// NewMobilityGroupsCopyResource returns a storage system resource object
func NewMobilityGroupsCopyResource() resource.Resource {
	return &mobilityGroupsCopyResource{}
}

// mobilityGroupsCopyResource is the resource implementation.
type mobilityGroupsCopyResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *mobilityGroupsCopyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_mobility_groups_copy"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *mobilityGroupsCopyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used in the process of copying data from one storage group (the source) to another storage group (the target) within the APEX Navigator environment.",
		MarkdownDescription: "This Terraform resource is used in the process of copying data from one storage group (the source) to another storage group (the target) within the APEX Navigator environment.",
		Attributes: map[string]schema.Attribute{
			"mobility_source_id": schema.StringAttribute{
				MarkdownDescription: "Source ID of the Mobility Group",
				Description:         "Source ID of the Mobility Group",
				Required:            true,
			},
			"mobility_target_id": schema.StringAttribute{
				MarkdownDescription: "Target ID of the Mobility Group",
				Description:         "Target ID of the Mobility Group",
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the Mobility Group Copy Job",
				Description:         "Status of the Mobility Group Copy Job",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the Mobility Group Copy Job",
				Description:         "ID of the Mobility Group Copy Job",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"powerflex_source": schema.SingleNestedBlock{
				Attributes: PowerflexAsyncInfo.Attributes,
			},
			"powerflex_target": schema.SingleNestedBlock{
				Attributes: PowerflexAsyncInfo.Attributes,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *mobilityGroupsCopyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clients, ok := req.ProviderData.(Clients)
	if !ok {
		resp.Diagnostics.AddError(
			constants.ResourceConfigureErrorMsg,
			fmt.Sprintf(constants.GeneralConfigureErrorMsg, req.ProviderData),
		)
		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *mobilityGroupsCopyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.MobilityGroupCopyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get/Validate the mobility group id set in the plan
	mobilityGroup, status, err := helper.GetMobilityGroup(r.client, plan.MobilitySourceID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityGroupReadErrorMsg,
			constants.BlockMobilityGroupReadDetailMsg+newErr,
		)
		return
	}

	// Get/Validate the mobility target value from Apex Navigator
	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, plan.MobilityTargetID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityGroupReadErrorMsg,
			constants.BlockMobilityGroupReadDetailMsg+newErr,
		)
		return
	}

	// Activate Powerflex Source and check every so often while the copy is in progress
	stopSource := make(chan bool)
	actErrSource := make(chan error)
	go helper.AsyncCheckActivation(ctx, stopSource, actErrSource, r.client, mobilityGroup.SystemId, *plan.PowerFlexClientSource, client.STORAGEPRODUCTENUM_POWERFLEX)

	// Activate Powerflex Target and check every so often while the copy is in progress
	stopTarget := make(chan bool)
	actErrTarget := make(chan error)
	go helper.AsyncCheckActivation(ctx, stopTarget, actErrTarget, r.client, mobilityTarget.SystemId, *plan.PowerFlexClientTarget, client.STORAGEPRODUCTENUM_POWERFLEX)

	// Check for Activation Error in a non blocking way
	select {
	case chanErrSource, ok := <-actErrSource:
		if ok {
			resp.Diagnostics.AddError(
				constants.ErrorActivatingPowerFlexSystem,
				fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityGroup.SystemId, chanErrSource.Error()),
			)
			// Closing channel for activation routine
			close(stopSource)
			close(stopTarget)
			return
		}
	case chanErrTarget, ok := <-actErrTarget:
		if ok {
			resp.Diagnostics.AddError(
				constants.ErrorActivatingPowerFlexSystem,
				fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityTarget.SystemId, chanErrTarget.Error()),
			)
			// Closing channel for activation routine
			close(stopSource)
			close(stopTarget)
			return
		}
	default:
		// do nothing since the proccess is still running without error
	}

	// Create Mobility Groups POST request
	createReq := r.client.MobilityGroupsAPI.MobilityGroupsCopy(ctx, plan.MobilitySourceID.ValueString())
	startCopyInput := *client.NewStartCopyInput()
	mobilityTargetIDArray := make([]string, 0)
	mobilityTargetIDArray = append(mobilityTargetIDArray, plan.MobilityTargetID.ValueString())

	startCopyInput.SetMobilityTargetIds(mobilityTargetIDArray)

	// Executing copy request request
	job, status, err := helper.CopyMobilityGroups(createReq, startCopyInput)

	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityGroupCopyErrorMsg,
			constants.BlockMobilityGroupCopyDetailMsg+newErr,
		)
		close(stopSource)
		close(stopTarget)
		return
	}
	// Waiting for job to complete
	_, err = helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		// Check for Activation Error in a non blocking way
		select {
		case chanErrSource, ok := <-actErrSource:
			if ok {
				resp.Diagnostics.AddError(
					constants.ErrorActivatingPowerFlexSystem,
					fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityGroup.SystemId, chanErrSource.Error()),
				)
				// Closing channel for activation routine
				close(stopSource)
				close(stopTarget)
				return
			}
		case chanErrTarget, ok := <-actErrTarget:
			if ok {
				resp.Diagnostics.AddError(
					constants.ErrorActivatingPowerFlexSystem,
					fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityTarget.SystemId, chanErrTarget.Error()),
				)
				// Closing channel for activation routine
				close(stopSource)
				close(stopTarget)
				return
			}
		default:
			// do nothing since the proccess is still running without error
		}

		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.GeneralJobError+err.Error(),
		)
		close(stopSource)
		close(stopTarget)
		return
	}

	// Fetching Job Status
	jobStatus, err := helper.GetJobStatus(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			constants.ErrorGettingJob,
			constants.JobRetrieveError+err.Error(),
		)
		close(stopSource)
		close(stopTarget)
		return
	}

	// Check for Activation Error in a non blocking way
	select {
	case chanErrSource, ok := <-actErrSource:
		if ok {
			resp.Diagnostics.AddError(
				constants.ErrorActivatingPowerFlexSystem,
				fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityGroup.SystemId, chanErrSource.Error()),
			)
			// Closing channel for activation routine
			close(stopSource)
			close(stopTarget)
			return
		}
	case chanErrTarget, ok := <-actErrTarget:
		if ok {
			resp.Diagnostics.AddError(
				constants.ErrorActivatingPowerFlexSystem,
				fmt.Sprintf(constants.ErrorActivatingPowerFlexSpecificSystemDetail, mobilityTarget.SystemId, chanErrTarget.Error()),
			)
			// Closing channel for activation routine
			close(stopSource)
			close(stopTarget)
			return
		}
	default:
		// do nothing since the proccess is still running without error
	}

	tflog.Debug(ctx, "Closing Activation Channels")
	// Closing channel for activation routine
	close(stopSource)
	close(stopTarget)

	// Updating TFState with Mobility group copy info
	result := models.MobilityGroupCopyModel{
		ID:               types.StringValue(job.Id),
		MobilitySourceID: plan.MobilitySourceID,
		MobilityTargetID: plan.MobilityTargetID,
		Status:           types.StringValue(string(*jobStatus.State)),
	}

	result.PowerFlexClientSource = helper.SetPowerflexAysncClientState(*plan.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexAysncClientState(*plan.PowerFlexClientTarget)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (r *mobilityGroupsCopyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { //nolint:dupl
	// Get current state
	var state models.MobilityGroupCopyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	result := models.MobilityGroupCopyModel{
		ID:               state.ID,
		MobilitySourceID: state.MobilitySourceID,
		MobilityTargetID: state.MobilityTargetID,
		Status:           state.Status,
	}

	result.PowerFlexClientSource = helper.SetPowerflexAysncClientState(*state.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexAysncClientState(*state.PowerFlexClientTarget)

	// Set refresded state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *mobilityGroupsCopyResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning(
		constants.UpdatesNotSupportedErrorMsg,
		constants.UpdatesNotSupportedErrorMsg,
	)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *mobilityGroupsCopyResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	resp.Diagnostics.AddWarning(
		constants.DeleteIsNotSupportedErrorMsg,
		constants.DeleteIsNotSupportedErrorMsg,
	)
}

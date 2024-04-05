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

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
		Description:         "This Terraform resource is used to copy source Data Mobility Group to target on Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to copy source Data Mobility Group to target on Apex Navigator.",
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
				Attributes: PowerflexInfo.Attributes,
			},
			"powerflex_target": schema.SingleNestedBlock{
				Attributes: PowerflexInfo.Attributes,
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
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *provider.Clients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
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
			"Error Reading Apex Navigator mobility group",
			"Could not read Apex Navigator mobility group, unexpected error: "+newErr,
		)
		return
	}

	// Get/Validate the mobility target value from Apex Navigator
	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, plan.MobilityTargetID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator mobility target",
			"Could not read Mobility target, unexpected error: "+newErr,
		)
		return
	}

	// Activate Powerflex
	actErr := helper.ActivateSystemPowerflexSystem(ctx, r.client, mobilityGroup.SystemId, *plan.PowerFlexClientSource, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			"Error activating Powerflex System",
			"Could not activate powerflex system, please check username/password and system id are correct: "+actErr.Error(),
		)
		return
	}

	// Activate Powerflex
	act2Err := helper.ActivateSystemPowerflexSystem(ctx, r.client, mobilityTarget.SystemId, *plan.PowerFlexClientTarget, client.STORAGEPRODUCTENUM_POWERFLEX)
	if act2Err != nil {
		resp.Diagnostics.AddError(
			"Error activating Powerflex System",
			"Could not activate powerflex system, please check username/password and system id are correct: "+act2Err.Error(),
		)
		return
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
			"Error creating Mobility Group Copy",
			"Could not create Mobility Group Copy, unexpected error: "+newErr,
		)
		return
	}

	// Waiting for job to complete
	_, err = helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			helper.ResourceRetrieveError+err.Error(),
		)
		return
	}

	// Fetching Job Status
	jobStatus, err := helper.GetJobStatus(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting job",
			helper.JobRetrieveError+err.Error(),
		)
		return
	}

	// Updating TFState with Mobility group copy info
	result := models.MobilityGroupCopyModel{
		ID:               types.StringValue(job.Id),
		MobilitySourceID: plan.MobilitySourceID,
		MobilityTargetID: plan.MobilityTargetID,
		Status:           types.StringValue(string(*jobStatus.State)),
	}

	result.PowerFlexClientSource = helper.SetPowerflexClientState(*plan.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexClientState(*plan.PowerFlexClientTarget)

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

	result.PowerFlexClientSource = helper.SetPowerflexClientState(*state.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexClientState(*state.PowerFlexClientTarget)

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
		"Updates are not supported for this resource",
		"Updates are not supported for this resource",
	)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *mobilityGroupsCopyResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	resp.Diagnostics.AddWarning(
		"Deletes are not supported for this resource",
		"Deletes are not supported for this resource",
	)
}

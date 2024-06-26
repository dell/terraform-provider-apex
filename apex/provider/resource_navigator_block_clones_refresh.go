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
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &clonesRefreshResource{}
	_ resource.ResourceWithConfigure = &clonesRefreshResource{}
)

// NewClonesRefreshResource returns a storage system resource object
func NewClonesRefreshResource() resource.Resource {
	return &clonesRefreshResource{}
}

// clonesRefreshResource is the resource implementation.
type clonesRefreshResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *clonesRefreshResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_clones_refresh"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *clonesRefreshResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Description: "This Terraform resource is used to manage clones refreshes on Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"clone_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the clone to be refreshed.",
				Description:         "Unique identifier for the clone to be refreshed.",
				Required:            true,
			},
			"status": schema.StringAttribute{
				Description:         "Status of the clone refreh Job",
				MarkdownDescription: "Status of the clone refresh Job",
				Computed:            true,
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Description:         "ID for the clone refresh Job",
				MarkdownDescription: "ID for the clone refresh Job",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"powerflex": schema.SingleNestedBlock{
				Attributes: PowerflexInfo.Attributes,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *clonesRefreshResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *clonesRefreshResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.ClonesRefreshModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.ActivationClientModel, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actErr.Error(),
		)
		return
	}

	// Create Clones POST request
	refreshReq := r.client.ClonesAPI.ClonesRefresh(ctx, plan.CloneID.ValueString())

	// Executing refresh request request
	job, status, err := helper.RefreshClone(refreshReq)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneRefreshErrorMsg,
			constants.BlockCloneRefreshDetailMsg+newErr,
		)
		return
	}

	// Waiting for job to complete
	_, err = helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.GeneralJobError+err.Error(),
		)
		return
	}

	// Fetching Job Status
	jobStatus, err := helper.GetJobStatus(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			constants.ErrorGettingJob,
			constants.JobRetrieveError+err.Error(),
		)
		return
	}

	// Updating TFState with Clones refresh info
	result := models.ClonesRefreshModel{
		ID:       types.StringValue(job.Id),
		CloneID:  plan.CloneID,
		Status:   types.StringValue(string(*jobStatus.State)),
		SystemID: plan.SystemID,
	}

	result.ActivationClientModel = helper.SetPowerflexClientState(*plan.ActivationClientModel)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (r *clonesRefreshResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state models.ClonesRefreshModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state = models.ClonesRefreshModel{
		ID:                    state.ID,
		CloneID:               state.CloneID,
		Status:                state.Status,
		SystemID:              state.SystemID,
		ActivationClientModel: state.ActivationClientModel,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *clonesRefreshResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning(
		constants.UpdatesNotSupportedErrorMsg,
		constants.UpdatesNotSupportedErrorMsg,
	)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *clonesRefreshResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	resp.Diagnostics.AddWarning(
		constants.DeleteIsNotSupportedErrorMsg,
		constants.DeleteIsNotSupportedErrorMsg,
	)
}

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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &mobilityGroupsCopyResource{}
	_ resource.ResourceWithConfigure   = &mobilityGroupsCopyResource{}
	_ resource.ResourceWithImportState = &mobilityGroupsCopyResource{}
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
	resp.TypeName = req.ProviderTypeName + "_mobility_groups_copy"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *mobilityGroupsCopyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"mobility_source_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"mobility_target_id": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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

	// Create Mobility Groups POST request
	createReq := r.client.MobilityGroupsAPI.MobilityGroupsCopy(ctx, plan.MobilitySourceID.ValueString())
	startCopyInput := *client.NewStartCopyInput()
	mobilityTargetIDArray := make([]string, 0, len(plan.MobilityTargetID))
	for _, targetID := range plan.MobilityTargetID {
		mobilityTargetIDArray = append(mobilityTargetIDArray, targetID.ValueString())
	}
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

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
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
	state = models.MobilityGroupCopyModel{
		ID:               state.ID,
		MobilitySourceID: state.MobilitySourceID,
		MobilityTargetID: state.MobilityTargetID,
		Status:           state.Status,
	}

	// Set refresded state
	diags = resp.State.Set(ctx, &state)
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

func (r *mobilityGroupsCopyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

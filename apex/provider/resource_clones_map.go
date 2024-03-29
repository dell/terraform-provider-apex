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

package provider //nolint:dupl

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
	_ resource.Resource              = &clonesMapResource{}
	_ resource.ResourceWithConfigure = &clonesMapResource{}
)

// NewClonesMapResource returns a storage system resource object
func NewClonesMapResource() resource.Resource {
	return &clonesMapResource{}
}

// clonesMapResource is the resource implementation.
type clonesMapResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *clonesMapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clones_map"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *clonesMapResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"clone_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"host_mappings": schema.ListAttribute{
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
func (r *clonesMapResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *clonesMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.ClonesMapModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create Clones POST request
	mapReq := r.client.ClonesAPI.ClonesMap(ctx, plan.CloneID.ValueString())

	// Executing map request request
	job, status, err := helper.MapClones(ctx, mapReq, plan.CloneID.ValueString(), plan.HostMappings)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error creating Clones Map Request",
			"Could not create Clones Map request, unexpected error: "+newErr,
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

	// Updating TFState with Clones map info
	result := models.ClonesMapModel{
		ID:           types.StringValue(job.Id),
		CloneID:      plan.CloneID,
		HostMappings: plan.HostMappings,
		Status:       types.StringValue(string(*jobStatus.State)),
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (r *clonesMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { //nolint:dupl
	// Get current state
	var state models.ClonesMapModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state = models.ClonesMapModel{
		ID:           state.ID,
		CloneID:      state.CloneID,
		HostMappings: state.HostMappings,
		Status:       state.Status,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *clonesMapResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning(
		"Updates are not supported for this resource",
		"Updates are not supported for this resource",
	)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *clonesMapResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Deletes are not supported for this resource",
		"Deletes are not supported for this resource",
	)
}

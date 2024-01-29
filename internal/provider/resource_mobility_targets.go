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

	poller "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/internal/jobs"
	client "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/pkg/gen/apex/client"
	jmsClient "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/pkg/gen/jobapi/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &mobilityTargetsResource{}
	_ resource.ResourceWithConfigure   = &mobilityTargetsResource{}
	_ resource.ResourceWithImportState = &mobilityTargetsResource{}
)

// NewMobilityTargetsResource returns a storage system resource object
func NewMobilityTargetsResource() resource.Resource {
	return &mobilityTargetsResource{}
}

// mobilityTargetsResource is the resource implementation.
type mobilityTargetsResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *mobilityTargetsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mobility_targets"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *mobilityTargetsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: " ",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"system_type": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"source_mobility_group_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"creation_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            false,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"image_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            false,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_members": schema.ListNestedAttribute{
				Optional: false,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional: false,
							Computed: true,
						},
						"parent_id": schema.StringAttribute{
							Optional: false,
							Computed: true,
						},
						"name": schema.StringAttribute{
							Optional: false,
							Computed: true,
						},
						"size": schema.StringAttribute{
							Optional: false,
							Computed: true,
						},
					},
				},
			},
			"last_copy_job_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            false,
				Computed:            true,
			},
			"bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"target_system_options": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *mobilityTargetsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clients, ok := req.ProviderData.(Clients)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *provider.CLients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *mobilityTargetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan MobilityTargetModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create Mobility Target request
	createReq := r.client.MobilityTargetsAPI.MobilityTargetsCreate(ctx)

	targetSystemOptions := client.TargetSystemOptions{
		StoragePool: plan.TargetSystemOptions.ValueStringPointer(),
	}

	mobilityTargetsInput := *client.NewCreateTargetInput(plan.SourceMobilityGroupID.ValueString(), plan.Name.ValueString(), plan.SystemID.ValueString(), *plan.SystemType.Ptr(), targetSystemOptions)
	createReq = createReq.CreateTargetInput(mobilityTargetsInput)

	// Executing job request
	job, status, err := createReq.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error creating Mobility Target",
			"Could not create Mobility Target, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Fetching Job Status
	poller := poller.NewPoller(r.jobsClient)
	resourceID, err := poller.WaitForResource(ctx, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			ResourceRetrieveError+err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Fetching Mobility group after Job Completes
	mobilityTarget, status, err := r.client.MobilityTargetsAPI.MobilityTargetsInstance(ctx, resourceID).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving Mobility group",
			"Could not retrieve Mobility group, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Updating TFState with Mobility target info
	result := GetMobilityTargetModel(*mobilityTarget)
	// Updating TFState to include target system options
	result.TargetSystemOptions = plan.TargetSystemOptions

	// Set state to fully populated data
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (r *mobilityTargetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state MobilityTargetModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed mobility group value from Apex Navigator
	mobilityTarget, status, err := r.client.MobilityTargetsAPI.MobilityTargetsInstance(context.Background(), state.ID.ValueString()).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator mobility target",
			"Could not read Mobility target, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	if mobilityTarget.SystemId != state.SystemID.ValueString() || *mobilityTarget.SystemType != *state.SystemType || mobilityTarget.SourceMobilityGroupId != state.SourceMobilityGroupID.ValueString() {
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator mobility target",
			"Attempted to modify unchangeable attribute(s) [SystemID, SystemType, sourceMobilityGroupID]"+fmt.Sprintf(("%s, %s"), *mobilityTarget.SystemType, *state.SystemType),
		)
		return
	}

	// Overwrite items with refreshed state
	state = GetMobilityTargetModel(*mobilityTarget)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *mobilityTargetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan MobilityTargetModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the update request
	req2 := r.client.MobilityTargetsAPI.MobilityTargetsModify(ctx, plan.ID.ValueString())
	limit := int32(plan.BandwidthLimit.ValueInt64())

	updateInput := client.UpdateMobilityTargetInput{
		Name:           plan.Name.ValueStringPointer(),
		Description:    plan.Description.ValueStringPointer(),
		BandwidthLimit: &limit,
	}
	if limit <= 0 {
		updateInput.BandwidthLimit = nil
	}

	req2 = req2.UpdateMobilityTargetInput(updateInput)

	// Execute Update Job
	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error executing Update Mobility Target Job",
			"Could not execute update mobility target, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Fetching Job Status
	poller := poller.NewPoller(r.jobsClient)
	resourceID, err := poller.WaitForResource(ctx, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			ResourceRetrieveError+err.Error(),
		)
		return
	}

	// Fetching Clone after Job Completes
	mobilityTarget, status, err := r.client.MobilityTargetsAPI.MobilityTargetsInstance(ctx, resourceID).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving Mobility Target",
			"Could not retrieve Mobility Target, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Updating TFState with Mobility Target info
	result := GetMobilityTargetModel(*mobilityTarget)
	result.TargetSystemOptions = plan.TargetSystemOptions

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
// nolint:dupl
func (r *mobilityTargetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan MobilityTargetModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing mobility target
	req2 := r.client.MobilityTargetsAPI.MobilityTargetsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Deleting Apex Navigator Mobility Target",
			"Could not delete mobility target, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}
	poller := poller.NewPoller(r.jobsClient)
	resourceID, err := poller.WaitForResource(ctx, job.Id)
	if (err != nil) || (resourceID == "") {
		resp.Diagnostics.AddError(
			"Error getting Delete Job ID",
			JobRetrieveError+err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}
}

func (r *mobilityTargetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

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
	"encoding/json"
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
	_ resource.Resource                = &mobilityGroupsResource{}
	_ resource.ResourceWithConfigure   = &mobilityGroupsResource{}
	_ resource.ResourceWithImportState = &mobilityGroupsResource{}
)

// NewMobilityGroupsResource returns a storage system resource object
func NewMobilityGroupsResource() resource.Resource {
	return &mobilityGroupsResource{}
}

// mobilityGroupsResource is the resource implementation.
type mobilityGroupsResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *mobilityGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_mobility_groups"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *mobilityGroupsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to manage SOurce Data Mobility Groups on Apex Navigator.We can create, read, update, delete Data Mobility Groups on Apex Navigator.We can also import existing Data Mobility Groups from Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to manage Data Mobility Groups on Apex Navigator.We can create, read, update, delete Data Mobility Groups on Apex Navigator.We can also import existing Data Mobility Groups from Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Mobility group identifier",
				Description:         "Mobility group identifier",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Mobility group name",
				Description:         "Mobility group name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Mobility group description",
				Description:         "Mobility group description",
				Optional:            true,
				Computed:            true,
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the system for the mobility group members",
				Description:         "Identifier of the system for the mobility group members",
				Required:            true,
			},
			"system_type": schema.StringAttribute{
				MarkdownDescription: "The source system type (e.g.: POWERFLEX)",
				Description:         "The source system type (e.g.: POWERFLEX)",
				Required:            true,
			},
			"creation_timestamp": schema.StringAttribute{
				MarkdownDescription: "When the mobility group was created",
				Description:         "When the mobility group was created",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"volume_id": schema.ListAttribute{
				Description:         "List of volume ids you want to add to the group",
				MarkdownDescription: "List of volume ids you want to add to the group",
				ElementType:         types.StringType,
				Required:            true,
			},
			"members": schema.ListNestedAttribute{
				Description:         "A mobility member is an object (e.g. volume) that is part of a mobility group that will be the source of mobility copy operations.",
				MarkdownDescription: "A mobility member is an object (e.g. volume) that is part of a mobility group that will be the source of mobility copy operations.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Identifier of the member (e.g. PowerFlex volume identifier)",
							MarkdownDescription: "Identifier of the member (e.g. PowerFlex volume identifier)",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the member (e.g. name of the volume)",
							MarkdownDescription: "Name of the member (e.g. name of the volume)",
							Computed:            true,
						},
						"size": schema.StringAttribute{
							Computed:            true,
							Description:         "Size of the member (e.g. volume size in bytes)",
							MarkdownDescription: "Size of the member (e.g. volume size in bytes)",
						},
					},
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
func (r *mobilityGroupsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *mobilityGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.MobilityGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.ActivationClientModel, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			"Error activating Powerflex System",
			"Could not activate powerflex system, please check username/password and system id are correct: "+actErr.Error(),
		)
		return
	}
	// Executing job request
	job, status, err := helper.CreateMobilityGroup(ctx, r.client, plan)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error creating Mobility Group",
			"Could not create Mobility Group, unexpected error: "+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, err := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			helper.ResourceRetrieveError+err.Error(),
		)
		return
	}

	// Fetching Mobility group after Job Completes
	mobilityGroup, status, err := helper.GetMobilityGroup(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving created Mobility group",
			"Could not retrieve created Mobility group, unexpected error: "+newErr,
		)
		return
	}

	// Updating TFState with Mobility group info
	result := helper.GetMobilityGroupModel(*mobilityGroup)
	if plan.VolumeID != nil {
		result.VolumeID = append(result.VolumeID, plan.VolumeID...)
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
func (r *mobilityGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state models.MobilityGroupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed mobility group value from Apex Navigator
	mobilityGroup, status, err := helper.GetMobilityGroup(r.client, state.ID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator mobility group",
			"Could not read Apex Navigator mobility group, unexpected error: "+newErr,
		)
		return
	}

	// Overwrite items with refreshed state
	result := helper.GetMobilityGroupModel(*mobilityGroup)

	if state.VolumeID != nil {
		result.VolumeID = append(result.VolumeID, state.VolumeID...)
	}

	result.ActivationClientModel = helper.SetPowerflexClientState(*state.ActivationClientModel)

	// Set refresded state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *mobilityGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan models.MobilityGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.ActivationClientModel, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			"Error activating Powerflex System",
			"Could not activate powerflex system, please check username/password and system id are correct: "+actErr.Error(),
		)
		return
	}

	// Get current state
	var state models.MobilityGroupModel
	diagsState := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.SystemID.ValueString() != plan.SystemID.ValueString() || *state.SystemType != *plan.SystemType {
		resp.Diagnostics.AddError(
			"Error updating Mobility Group",
			"Cannot update Mobility Group, attempted to update unchangeable attribute [SystemID, SystemType]",
		)
		return
	}

	// Execute Update Job
	job, status, err := helper.UpdateMobilityGroup(ctx, r.client, plan)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error executing Update Mobility Group Job",
			"Could not execute update Mobility group, unexpected error: "+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, err := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			helper.ResourceRetrieveError+err.Error(),
		)
		return
	}

	// Fetching Clone after Job Completes
	mobilityGroup, status, err := helper.GetMobilityGroup(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving updated Mobility Group",
			"Could not retrieve updated Mobility Group, unexpected error: "+newErr,
		)
		return
	}

	// Updating TFState with Mobility Group info
	result := helper.GetMobilityGroupModel(*mobilityGroup)
	if plan.VolumeID != nil {
		result.VolumeID = append(result.VolumeID, plan.VolumeID...)
	}
	result.ActivationClientModel = helper.SetPowerflexClientState(*plan.ActivationClientModel)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *mobilityGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan models.MobilityGroupModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.ActivationClientModel, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			"Error activating Powerflex System",
			"Could not activate powerflex system, please check username/password and system id are correct: "+actErr.Error(),
		)
		return
	}

	// Delete existing mobility group
	req2 := r.client.MobilityGroupsAPI.MobilityGroupsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Deleting Apex Navigator Mobility Group",
			"Could not delete mobility group, unexpected error: "+newErr,
		)
		return
	}

	resourceID, err := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if (err != nil) || (resourceID == "") {
		resp.Diagnostics.AddError(
			"Error getting Delete Job ID",
			helper.JobRetrieveError+err.Error(),
		)
		return
	}

}

func (r *mobilityGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Scheme   string `json:"scheme"`
		Insecure bool   `json:"insecure"`
		ID       string `json:"id"`
	}

	var p params
	err := json.Unmarshal([]byte(req.ID), &p)
	if err != nil {
		resp.Diagnostics.AddError(", make sure you include username, password, host, scheme, insecure and id", err.Error())
	}

	mobilityGroup, status, err := helper.GetMobilityGroup(r.client, p.ID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator mobility group",
			"Could not retrieve Mobility Groups during import: "+newErr,
		)
		return
	}

	// Overwrite items with refreshed state
	result := helper.GetMobilityGroupModel(*mobilityGroup)

	result.ActivationClientModel = &models.ActivationClientModel{
		Username: types.StringValue(p.Username),
		Password: types.StringValue(p.Password),
		Host:     types.StringValue(p.Host),
		Scheme:   types.StringValue(p.Scheme),
		Insecure: types.BoolValue(p.Insecure),
	}

	diags := resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

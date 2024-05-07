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
	resp.TypeName = req.ProviderTypeName + "_navigator_block_mobility_targets"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *mobilityTargetsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to manage Mobility Targets on Apex Navigator. We can create, read, update, delete Data Mobility Targets on Apex Navigator.We can also import existing Data Mobility Targets from Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to manage Mobility Targets on Apex Navigator. We can create, read, update, delete Data Mobility Targets on Apex Navigator.We can also import existing Data Mobility Targets from Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Idenifier of this target mobility group",
				Description:         "Idenifier of this target mobility group",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the mobility target",
				Description:         "Name of the mobility target",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the mobility target",
				Description:         "Description of the mobility target",
				Optional:            true,
				Computed:            true,
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: "ID of the target system",
				Description:         "ID of the target system",
				Required:            true,
			},
			"system_type": schema.StringAttribute{
				MarkdownDescription: "The source system type (e.g.: POWERFLEX)",
				Description:         "The source system type (e.g.: POWERFLEX)",
				Required:            true,
			},
			"source_mobility_group_id": schema.StringAttribute{
				MarkdownDescription: "ID of the source mobility group",
				Description:         "ID of the source mobility group",
				Required:            true,
			},
			"creation_timestamp": schema.StringAttribute{
				MarkdownDescription: "Timestamp from when the group was created",
				Description:         "Timestamp from when the group was created",
				Optional:            false,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"image_timestamp": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the last source image copied to this target",
				Description:         "Timestamp of the last source image copied to this target",
				Optional:            false,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_members": schema.ListNestedAttribute{
				Optional:            false,
				Computed:            true,
				Description:         "A mobility member map is a mapping of a mobility member and it's related member. For example a target volume with a reference to the source volume. Or a clone volume and its related target volume.",
				MarkdownDescription: "A mobility member map is a mapping of a mobility member and it's related member. For example a target volume with a reference to the source volume. Or a clone volume and its related target volume.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional:            false,
							Computed:            true,
							Description:         "ID of the member",
							MarkdownDescription: "ID of the member",
						},
						"parent_id": schema.StringAttribute{
							Optional:            false,
							Computed:            true,
							Description:         "Identifier of the related mobility member",
							MarkdownDescription: "Identifier of the related mobility member",
						},
						"name": schema.StringAttribute{
							Optional:            false,
							Computed:            true,
							Description:         "Name of the member",
							MarkdownDescription: "Name of the member",
						},
						"size": schema.StringAttribute{
							Optional:            false,
							Computed:            true,
							Description:         "Size of the member",
							MarkdownDescription: "Size of the member",
						},
					},
				},
			},
			"last_copy_job_id": schema.StringAttribute{
				MarkdownDescription: "Last copy job ID",
				Description:         "Last copy job ID",
				Optional:            false,
				Computed:            true,
			},
			"bandwidth_limit": schema.Int64Attribute{
				MarkdownDescription: "Bandwidth limit in Mbps (Mega bits per second)",
				Description:         "Bandwidth limit in Mbps (Mega bits per second)",
				Optional:            true,
				Computed:            true,
			},
			"target_system_options": schema.StringAttribute{
				Description:         "Storage pool id to use for allocating target volumes",
				MarkdownDescription: "Storage pool id to use for allocating target volumes",
				Required:            true,
			},
			"type": schema.StringAttribute{
				Computed: true,
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
func (r *mobilityTargetsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *mobilityTargetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.MobilityTargetModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex Source
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientSource, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actErr.Error(),
		)
		return
	}

	// Activate Powerflex Target
	actTargetErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientTarget, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actTargetErr.Error(),
		)
		return
	}

	// Create Mobility Target request
	createReq := r.client.MobilityTargetsAPI.MobilityTargetsCreate(ctx)

	targetSystemOptions := client.TargetSystemOptions{
		StoragePool: plan.TargetSystemOptions.ValueStringPointer(),
	}

	mobilityTargetsInput := *client.NewCreateTargetInput(plan.SourceMobilityGroupID.ValueString(), plan.Name.ValueString(), plan.SystemID.ValueString(), *plan.SystemType.Ptr(), targetSystemOptions)

	mobilityTargetsInput.Description = plan.Description.ValueStringPointer()
	if plan.BandwidthLimit.ValueInt64() != 0 {
		limit := int32(plan.BandwidthLimit.ValueInt64())
		mobilityTargetsInput.BandwidthLimit = &limit
	}
	// Executing job request
	job, status, err := helper.CreateMobilityTarget(createReq, mobilityTargetsInput)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetCreateErrorMsg,
			constants.BlockMobilityTargetCreateDetailMsg+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, err := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.GeneralJobError+err.Error(),
		)
		return
	}

	// Fetching Mobility group after Job Completes
	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetReadErrorMsg,
			constants.BlockMobilityTargetReadDetailMsg+newErr,
		)
		return
	}

	// Updating TFState with Mobility target info
	result := helper.GetMobilityTargetModel(*mobilityTarget)
	// Updating TFState to include target system options
	result.TargetSystemOptions = plan.TargetSystemOptions

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
func (r *mobilityTargetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state models.MobilityTargetModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed mobility group value from Apex Navigator
	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, state.ID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetReadErrorMsg,
			constants.BlockMobilityTargetReadDetailMsg+newErr,
		)
		return
	}

	// Overwrite items with refreshed state
	result := helper.GetMobilityTargetModel(*mobilityTarget)
	result.TargetSystemOptions = state.TargetSystemOptions
	result.PowerFlexClientSource = helper.SetPowerflexClientState(*state.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexClientState(*state.PowerFlexClientTarget)

	// Set refreshed state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *mobilityTargetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan models.MobilityTargetModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve values from state
	var state models.MobilityTargetModel
	diagsState := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.SystemID.ValueString() != state.SystemID.ValueString() || *plan.SystemType != *state.SystemType || plan.SourceMobilityGroupID.ValueString() != state.SourceMobilityGroupID.ValueString() {
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetUpdateErrorMsg,
			constants.BlockMobilityTargetUpdateInvalidFieldUpdateErrorMsg,
		)
		return
	}

	// Activate Powerflex Source
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientSource, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actErr.Error(),
		)
		return
	}

	// Activate Powerflex Target
	actTargetErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientTarget, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actTargetErr.Error(),
		)
		return
	}

	// Create the update request
	updateReq := r.client.MobilityTargetsAPI.MobilityTargetsModify(ctx, plan.ID.ValueString())
	limit := int32(plan.BandwidthLimit.ValueInt64())

	updateBlockInput := client.UpdateMobilityTargetBlockInput{
		Name:           plan.Name.ValueStringPointer(),
		Description:    plan.Description.ValueStringPointer(),
		BandwidthLimit: &limit,
	}
	if limit <= 0 {
		updateBlockInput.BandwidthLimit = nil
	}

	updateInput := client.UpdateMobilityTargetInput{
		UpdateMobilityTargetBlockInput: &updateBlockInput,
	}

	// Execute Update Job
	job, status, err := helper.UpdateMobilityTarget(updateReq, updateInput)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetUpdateErrorMsg,
			constants.BlockMobilityTargetUpdateDetailMsg+newErr,
		)
		return
	}
	// Fetching Job Status
	resourceID, err := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.GeneralJobError+err.Error(),
		)
		return
	}

	// Fetching Clone after Job Completes
	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetReadErrorMsg,
			constants.BlockMobilityTargetReadDetailMsg+newErr,
		)
		return
	}

	// Updating TFState with Mobility Target info
	result := helper.GetMobilityTargetModel(*mobilityTarget)
	result.TargetSystemOptions = plan.TargetSystemOptions
	result.PowerFlexClientSource = helper.SetPowerflexClientState(*plan.PowerFlexClientSource)
	result.PowerFlexClientTarget = helper.SetPowerflexClientState(*plan.PowerFlexClientTarget)
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
	var plan models.MobilityTargetModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Activate Powerflex Source
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientSource, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actErr.Error(),
		)
		return
	}

	// Activate Powerflex Target
	actTargetErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.PowerFlexClientTarget, client.STORAGEPRODUCTENUM_POWERFLEX)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerFlexSystem,
			constants.ErrorActivatingPowerFlexSystemDetail+actTargetErr.Error(),
		)
		return
	}

	// Delete existing mobility target
	req2 := r.client.MobilityTargetsAPI.MobilityTargetsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetDeleteErrorMsg,
			constants.BlockMobilityTargetDeleteDetailMsg+newErr,
		)
		return
	}

	poller := helper.NewPoller(r.jobsClient)
	resourceID, err := poller.WaitForResource(ctx, job.Id)
	if (err != nil) || (resourceID == "") {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.JobRetrieveError+err.Error(),
		)
		return
	}

}

// ImportState imports the Terraform resource
func (r *mobilityTargetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	type params struct {
		TargetUsername string `json:"target_username"`
		TargetPassword string `json:"target_password"`
		TargetHost     string `json:"target_host"`
		SourceUsername string `json:"source_username"`
		SourcePassword string `json:"source_password"`
		SourceHost     string `json:"source_host"`
		Scheme         string `json:"scheme"`
		Insecure       bool   `json:"insecure"`
		ID             string `json:"id"`
	}

	var p params
	err := json.Unmarshal([]byte(req.ID), &p)
	if err != nil {
		resp.Diagnostics.AddError("Error while unmarshalling import, make sure you include username, password, host, scheme, insecure and id", err.Error())
	}

	mobilityTarget, status, err := helper.GetMobilityTarget(r.client, p.ID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockMobilityTargetReadErrorMsg,
			constants.BlockMobilityTargetImportReadErrorMsg+newErr,
		)
		return
	}
	result := helper.GetMobilityTargetModel(*mobilityTarget)
	result.TargetSystemOptions = types.StringValue("")
	result.PowerFlexClientTarget = &models.ActivationClientModel{
		Username: types.StringValue(p.TargetUsername),
		Password: types.StringValue(p.TargetPassword),
		Host:     types.StringValue(p.TargetHost),
		Scheme:   types.StringValue(p.Scheme),
		Insecure: types.BoolValue(p.Insecure),
	}

	result.PowerFlexClientSource = &models.ActivationClientModel{
		Username: types.StringValue(p.SourceUsername),
		Password: types.StringValue(p.SourcePassword),
		Host:     types.StringValue(p.SourceHost),
		Scheme:   types.StringValue(p.Scheme),
		Insecure: types.BoolValue(p.Insecure),
	}

	diags := resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

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
	"regexp"

	jmsClient "dell/apex-job-client"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/constants"
	helper "github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &clonesResource{}
	_ resource.ResourceWithConfigure   = &clonesResource{}
	_ resource.ResourceWithImportState = &clonesResource{}
)

// NewClonesResource is a helper function to simplify the provider implementation.
func NewClonesResource() resource.Resource {
	return &clonesResource{}
}

// clonesResource is the resource implementation.
type clonesResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *clonesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_clones"
}

// Schema defines the schema for the resource.
func (r *clonesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to manage clones on Apex Navigator. We can create, read, update and delete the Clones using this resource.We can also import an existing clone from Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to manage clones on Apex Navigator. We can create, read, update and delete the Clones using this resource.We can also import an existing clone from Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the clone",
				Description:         "The ID of the clone",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the clone",
				Description:         "Name of the clone",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				Description:         "Description of the clone",
				MarkdownDescription: "Description of the clone",
				Optional:            true,
				Computed:            true,
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"mobility_target_id": schema.StringAttribute{
				MarkdownDescription: "Target ID",
				Description:         "Target ID",
				Required:            true,
			},
			"creation_timestamp": schema.StringAttribute{
				MarkdownDescription: "When the clone was created",
				Description:         "When the clone was created",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"refresh_timestamp": schema.StringAttribute{
				MarkdownDescription: "When the clone was last updated",
				Description:         "When the clone was last updated",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"image_timestamp": schema.StringAttribute{
				MarkdownDescription: "Image timestamp",
				Description:         "Image timestamp",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"clone_volumes": schema.ListNestedAttribute{
				Description:         "A clone mobility member provides details of clone volume",
				MarkdownDescription: "A clone mobility member provides details of clone volume",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "ID of the member",
							MarkdownDescription: "ID of the member",
							Computed:            true,
						},
						"parent_id": schema.StringAttribute{
							Description:         "Identifier of the related mobility member",
							MarkdownDescription: "Identifier of the related mobility member",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the member",
							MarkdownDescription: "Name of the member",
							Computed:            true,
						},
						"size": schema.StringAttribute{
							Description:         "Size of the member",
							MarkdownDescription: "Size of the member",
							Computed:            true,
						},
					},
				},
			},
			"host_mappings": schema.ListNestedAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "This represents the mapping of a repurposed (clone) storage object to a host (presumably using the clone for some analytical workload)",
				MarkdownDescription: "This represents the mapping of a repurposed (clone) storage object to a host (presumably using the clone for some analytical workload)",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							Description:         "Name of host/SDC to be mapped to the clone",
							MarkdownDescription: "Name of host/SDC to be mapped to the clone",
							Computed:            true,
						},
						"host_ip": schema.StringAttribute{
							Description:         "IP address of host",
							MarkdownDescription: "IP address of host",
							Computed:            true,
						},
						"host_id": schema.StringAttribute{
							Description:         "Identifier of the host",
							MarkdownDescription: "Identifier of the host",
							Required:            true,
						},
						"id": schema.StringAttribute{
							Description:         "This is a host mappings id generated by APEX Navigator for Multicloud Storage",
							MarkdownDescription: "This is a host mappings id generated by APEX Navigator for Multicloud Storage",
							Computed:            true,
						},
						"nqn": schema.StringAttribute{
							Description:         "NVMe qualified name. Only applicable if host_initiator_protocol is NVMe",
							MarkdownDescription: "NVMe qualified name. Only applicable if host_initiator_protocol is NVMe",
							Computed:            true,
						},
						"host_initiator_protocol": schema.StringAttribute{
							Description:         "Type of the host",
							MarkdownDescription: "Type of the host",
							Computed:            true,
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

// Configure adds the provider configured client to the resource.
func (r *clonesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *clonesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan models.ClonesModel
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

	cloneCreateInput := *client.NewCloneCreateInput(plan.Name.ValueString(), plan.MobilityTargetID.ValueString())

	if plan.Description.ValueString() != "" {
		cloneCreateInput.Description = plan.Description.ValueStringPointer()
	}
	// Fetches the hostId using regex and appends it to the cloneCreateInput
	if len(plan.HostMappings.Elements()) != 0 {
		for _, mapping := range plan.HostMappings.Elements() {
			mapValue := mapping.String()
			re := regexp.MustCompile(`\"host_id\":\"(.*?)\"`)
			match := re.FindStringSubmatch(mapValue)
			if len(match) > 1 {
				hostIDValue := match[1]
				cloneCreateInput.HostMappings = append(cloneCreateInput.HostMappings, hostIDValue)
			}
		}
	}

	// Create new Clone
	job, status, err := helper.CreateClone(r.client.ClonesAPI.ClonesCreate(ctx), cloneCreateInput)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneCreateErrorMsg,
			constants.BlockCloneCreateDetailMsg+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, jobErr := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if jobErr != nil {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			constants.GeneralJobError+jobErr.Error(),
		)
		return
	}

	// Fetching Clone after Job Completes
	clone, status, err := helper.GetCloneInstance(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneReadErrorMsg,
			constants.BlockCloneReadDetailMsg+newErr,
		)
		return
	}
	// Updating TFState with Clone info
	result := helper.GetClonesModel(*clone, plan.SystemID.ValueString())
	result.ActivationClientModel = helper.SetPowerflexClientState(*plan.ActivationClientModel)
	// Set state to fully populated data
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *clonesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state models.ClonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	getID := state.ID.ValueString()

	// Get refreshed clones value from Apex Navigator
	clone, status, err := helper.GetCloneInstance(r.client, getID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneReadErrorMsg,
			constants.BlockCloneReadDetailMsg+newErr,
		)
		return
	}
	// set the updated state.
	result := helper.GetClonesModel(*clone, state.SystemID.ValueString())
	result.ActivationClientModel = helper.SetPowerflexClientState(*state.ActivationClientModel)
	// Set refreshed state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// isInvalidUpdate checks if the clone is an invalid update based on the plan.
func isInvalidUpdate(clone *client.Clone, plan models.ClonesModel) bool {
	if *clone.MobilityTargetId != plan.MobilityTargetID.ValueString() {
		return true
	}
	if len(clone.HostMappings) == len(plan.HostMappings.Elements()) {
		if len(clone.HostMappings) == 0 || len(plan.HostMappings.Elements()) == 0 {
			return false
		}
		for _, mapping := range clone.HostMappings {
			mapValue := plan.HostMappings.String()
			re := regexp.MustCompile(`\"host_id\":\"(.*?)\"`)
			match := re.FindStringSubmatch(mapValue)
			if len(match) > 1 {
				hostIDValue := match[1]
				if mapping.HostId != hostIDValue {
					return true
				}
			}
		}
	} else {
		return true
	}
	return false
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *clonesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan models.ClonesModel
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

	clone, _, _ := helper.GetCloneInstance(r.client, plan.ID.ValueString())
	if isInvalidUpdate(clone, plan) {
		resp.Diagnostics.AddError(
			constants.BlockCloneUpdateErrorMsg,
			constants.BlockCloneInvalidFieldUpdateErrorMsg,
		)
		return
	}

	// Update existing clones
	updateInput := client.UpdateCloneInput{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
	}

	// Execute Update Job
	job, status, err := helper.UpdateClone(r.client.ClonesAPI.ClonesModify(ctx, plan.ID.ValueString()), updateInput)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneUpdateErrorMsg,
			constants.BlockCloneUpdateDetailMsg+newErr,
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
	clone, status, err = helper.GetCloneInstance(r.client, resourceID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneReadErrorMsg,
			constants.BlockCloneReadDetailMsg+newErr,
		)
		return
	}
	// Updating TFState with Clone info
	result := helper.GetClonesModel(*clone, plan.SystemID.ValueString())
	result.ActivationClientModel = helper.SetPowerflexClientState(*plan.ActivationClientModel)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *clonesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var plan models.ClonesModel
	diags := req.State.Get(ctx, &plan)
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

	// Delete existing clones
	req2 := r.client.ClonesAPI.ClonesDelete(ctx, plan.ID.ValueString())
	d := client.DeleteCloneInput{
		RemoveFromSystem: true,
	}
	req2 = req2.DeleteCloneInput(d)

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneDeleteErrorMsg,
			constants.BlockCloneDeleteDetailMsg+newErr,
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

func (r *clonesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Scheme   string `json:"scheme"`
		Insecure bool   `json:"insecure"`
		ID       string `json:"id"`
		SystemID string `json:"system_id"`
	}

	var p params
	err := json.Unmarshal([]byte(req.ID), &p)
	if err != nil {
		resp.Diagnostics.AddError(", make sure you include system_id, username, password, host, scheme, insecure and id", err.Error())
	}

	// Get refreshed clones value from Apex Navigator
	clone, status, err := helper.GetCloneInstance(r.client, p.ID)
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockCloneReadErrorMsg,
			constants.BlockCloneImportReadErrorMsg+newErr,
		)
		return
	}
	// set the updated state.
	result := helper.GetClonesModel(*clone, p.SystemID)
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

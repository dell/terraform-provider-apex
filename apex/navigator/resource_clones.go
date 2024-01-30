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

package navigator

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	poller "github.com/dell/terraform-provider-apex/apex/jobs"
	client "github.com/dell/terraform-provider-apex/client/apex"
	jmsClient "github.com/dell/terraform-provider-apex/client/jobs"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	resp.TypeName = req.ProviderTypeName + "_clones"
}

// Schema defines the schema for the resource.
func (r *clonesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the clone",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"mobility_target_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"creation_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"refresh_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"image_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"clone_volumes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"parent_id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"size": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"host_mappings": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							Computed: true,
						},
						"host_ip": schema.StringAttribute{
							Computed: true,
						},
						"host_id": schema.StringAttribute{
							Required: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"nqn": schema.StringAttribute{
							Computed: true,
						},
						"host_initiator_protocol": schema.StringAttribute{
							Computed: true,
						},
					},
				},
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
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *provider.CLients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *clonesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ClonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create new Clone
	req2 := r.client.ClonesAPI.ClonesCreate(ctx)

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

	req2 = req2.CloneCreateInput(cloneCreateInput)

	// Executing Job Request
	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error creating Clones",
			"Could not create Clones, unexpected error: "+newErr,
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

	// Fetching Clone after Job Completes
	clone, status, err := r.client.ClonesAPI.ClonesInstance(ctx, resourceID).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving created Clone",
			"Could not retrieve created Clone, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Updating TFState with Clone info
	result := GetClonesModel(*clone)

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
	var state ClonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	getID := state.ID.ValueString()

	// Get refreshed clones value from Apex Navigator
	clone, status, err := r.client.ClonesAPI.ClonesInstance(context.Background(), getID).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator Clones",
			"Could not read Clones, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Overwrite items with refreshed state
	state = GetClonesModel(*clone)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// isInvalidUpdate checks if the clone is an invalid update based on the plan.
func isInvalidUpdate(clone *client.Clone, plan ClonesModel) bool {
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
	var plan ClonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clone, status, _ := r.client.ClonesAPI.ClonesInstance(ctx, plan.ID.ValueString()).Execute()
	if isInvalidUpdate(clone, plan) {
		resp.Diagnostics.AddError(
			"Error updating Clones",
			"Could not update Clones, invalid field modified [Modifiable fields: Name, Description]",
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Update existing clones
	req2 := r.client.ClonesAPI.ClonesModify(context.Background(), plan.ID.ValueString())
	u := client.UpdateCloneInput{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
	}
	req2 = req2.UpdateCloneInput(u)

	// Execute Update Job
	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error executing Update Clone Job",
			"Could not execute update clones, unexpected error: "+newErr,
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
	clone, status, err = r.client.ClonesAPI.ClonesInstance(ctx, resourceID).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error retrieving updated Clone",
			"Could not retrieve updated Clone, unexpected error: "+newErr,
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Updating TFState with Clone info
	result := GetClonesModel(*clone)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *clonesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var plan ClonesModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
		newErr := GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Deleting Apex Navigator Clones",
			"Could not delete clones, unexpected error: "+newErr,
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

func (r *clonesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

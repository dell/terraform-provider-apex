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
	"strings"

	jmsClient "dell/apex-job-client"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/constants"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &blockStorageResource{}
	_ resource.ResourceWithConfigure   = &blockStorageResource{}
	_ resource.ResourceWithImportState = &blockStorageResource{}
)

// NewBlockStorageResource returns a storage system resource object
func NewBlockStorageResource() resource.Resource {
	return &blockStorageResource{}
}

// storageProductsResource is the resource implementation.
type blockStorageResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *blockStorageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_storage"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *blockStorageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to manage Block Storage on Apex Navigator. We can create, read, update, delete Block Storage on Apex Navigator.We can also import existing Block Storage from Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to manage Block Storage on Apex Navigator. We can create, read, update, delete Block Storage on Apex Navigator.We can also import existing Block Storage from Apex Navigator.",
		Attributes:          GetStorageSystemSchema("block"),
		Blocks: map[string]schema.Block{
			"powerflex": schema.SingleNestedBlock{
				Attributes: PowerflexInfo.Attributes,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *blockStorageResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *blockStorageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.StorageModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	validationErr := helper.ValidateCreateStorageParamsBlock(plan)
	if validationErr != nil {
		resp.Diagnostics.AddError(
			constants.BlockStorageCreateErrorMsg,
			validationErr.Error(),
		)
		return
	}

	var systemCreateInput client.StorageSystemDeploymentRequest
	switch {
	case plan.DeploymentDetails == nil:
		break
	case plan.DeploymentDetails.SystemPublicCloud != nil:
		tier := client.PowerFlexTierEnum(plan.DeploymentDetails.SystemPublicCloud.TierType.ValueString())
		minIops := int32(plan.DeploymentDetails.SystemPublicCloud.MinimumIops.ValueInt64())
		minUC := int32(plan.DeploymentDetails.SystemPublicCloud.MinimumCapacity.ValueInt64())

		storageParams := client.StorageOptions{
			PowerFlexStorageOptions: &client.PowerFlexStorageOptions{
				SystemType:            plan.StorageSystemType.ValueString(),
				Version:               plan.ProductVersion.ValueStringPointer(),
				Tier:                  &tier,
				MinimumIops:           &minIops,
				MinimumUsableCapacity: &minUC,
			},
		}

		var subnetOptions []client.SubnetOptions
		for _, subnetOption := range plan.DeploymentDetails.SystemPublicCloud.SubnetOptions {
			var subnetType *client.SubnetTypeEnum
			if subnetOption.SubnetType != nil {
				subnetType = subnetOption.SubnetType.Ptr()
			} else {
				subnetType = client.SUBNETTYPEENUM_UNDEFINED.Ptr()
			}
			newOption := client.SubnetOptions{
				Id:        subnetOption.SubnetID.ValueStringPointer(),
				CidrBlock: subnetOption.CidrBlock.ValueStringPointer(),
				Type:      subnetType,
			}
			subnetOptions = append(subnetOptions, newOption)
		}

		cloudParams := client.CloudOptions{
			AWSCloudOptions: &client.AWSCloudOptions{
				CloudType:                string(*plan.DeploymentDetails.SystemPublicCloud.CloudType),
				AccountId:                plan.DeploymentDetails.SystemPublicCloud.CloudAccount.ValueString(),
				Region:                   plan.DeploymentDetails.SystemPublicCloud.CloudRegion.ValueString(),
				AvailabilityZoneTopology: plan.DeploymentDetails.SystemPublicCloud.AvailabilityZoneTopology.Ptr(),
				SshKeyName:               plan.DeploymentDetails.SystemPublicCloud.SSHKeyName.ValueString(),
				Vpc: client.Vpc{
					IsNewVpc: plan.DeploymentDetails.SystemPublicCloud.Vpc.IsNewVpc.ValueBoolPointer(),
					Id:       plan.DeploymentDetails.SystemPublicCloud.Vpc.VpcID.ValueStringPointer(),
					Name:     plan.DeploymentDetails.SystemPublicCloud.Vpc.VpcName.ValueStringPointer(),
				},
				SubnetOptions: subnetOptions,
			},
		}
		systemCreateInput = *client.NewStorageSystemDeploymentRequest(plan.Name.ValueString(), cloudParams, storageParams, true)
	case plan.DeploymentDetails.SystemOnPrem != nil:
		defaultValue := "default"
		minValue := int32(1)
		storageOptions := client.StorageOptions{
			PowerFlexStorageOptions: &client.PowerFlexStorageOptions{
				SystemType:            defaultValue,
				Version:               &defaultValue,
				Tier:                  client.POWERFLEXTIERENUM_BALANCED.Ptr(),
				MinimumIops:           &minValue,
				MinimumUsableCapacity: &minValue,
			},
		}
		var subnetOptions []client.SubnetOptions

		newOption := client.SubnetOptions{
			Id:        &defaultValue,
			CidrBlock: &defaultValue,
			Type:      client.SUBNETTYPEENUM_EXTERNAL.Ptr(),
		}
		subnetOptions = append(subnetOptions, newOption)

		falseBool := false
		cloudOptions := client.CloudOptions{
			AWSCloudOptions: &client.AWSCloudOptions{
				CloudType:                "default",
				AccountId:                "default",
				Region:                   "default",
				AvailabilityZoneTopology: client.AVAILABILITYZONETOPOLOGYENUM_MULTIPLE_AVAILABILITY_ZONE.Ptr(),
				SshKeyName:               "default",
				Vpc: client.Vpc{
					IsNewVpc: &falseBool,
					Id:       &defaultValue,
					Name:     &defaultValue,
				},
				SubnetOptions: subnetOptions,
			},
		}
		systemCreateInput = *client.NewStorageSystemDeploymentRequest(plan.Name.ValueString(), cloudOptions, storageOptions, true)
	default:
		resp.Diagnostics.AddError(
			constants.BlockStorageCreateErrorMsg,
			constants.StorageSystemInvalidDeploymentType,
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Creating Block Storage: %v", systemCreateInput))
	job, status, err := helper.CreateStorageSystem(r.client.StorageSystemsAPI.StorageSystemsCreate(ctx), systemCreateInput)

	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockStorageCreateErrorMsg,
			constants.BlockStorageCreateDetailMsg+newErr,
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

	// After the long job make sure to update the client token
	_ = helper.UpdateToken(ctx, r.client, r.jobsClient)

	// Fetching Block storage after Job Completes
	storageSystem, status, err := helper.GetStorageInstance(r.client, resourceID)
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		// Try again with POWERFLEX- prefix
		if status.StatusCode == http.StatusNotFound {
			resourceIDWithPrefix := "POWERFLEX-" + resourceID
			storageSystem, status, err = helper.GetStorageInstance(r.client, resourceIDWithPrefix)
			if (err != nil) || (status.StatusCode != http.StatusOK) {
				newErr := helper.GetErrorString(err, status)
				resp.Diagnostics.AddError(
					constants.StorageReadErrorMsg,
					constants.StorageReadDetailMsg+newErr,
				)
				return
			}
		} else {
			newErr := helper.GetErrorString(err, status)
			resp.Diagnostics.AddError(
				constants.StorageReadErrorMsg,
				constants.StorageReadDetailMsg+newErr,
			)
			return
		}
	}

	// Updating TFState with Block Storage info
	result := helper.GetStorageSystemBlock(*storageSystem)

	if result.SystemType.ValueString() == plan.StorageSystemType.ValueString() {
		result.StorageSystemType = plan.StorageSystemType
	}

	if strings.Contains(result.Version.ValueString(), plan.ProductVersion.ValueString()) {
		result.ProductVersion = plan.ProductVersion
	}
	// Need to check for cloud deployment details
	if plan.DeploymentDetails != nil &&
		plan.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpcBlock(plan, result)
	}

	if result.DeploymentDetails == nil {
		resp.Diagnostics.AddError(
			constants.UnexpectedSystemType,
			constants.UnexpectedSystemType,
		)
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
func (r *blockStorageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state models.StorageModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed storage systems value from Apex Navigator
	storageSystem, status, err := helper.GetStorageInstance(r.client, state.ID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.StorageReadErrorMsg,
			constants.StorageReadDetailMsg+newErr,
		)
		return
	}

	// Overwrite items with refreshed state
	result := helper.GetStorageSystemBlock(*storageSystem)
	if result.SystemType.ValueString() == state.StorageSystemType.ValueString() {
		result.StorageSystemType = state.StorageSystemType
	}

	if strings.Contains(result.Version.ValueString(), state.ProductVersion.ValueString()) {
		result.ProductVersion = state.ProductVersion
	}

	// Need to check for on prem deployment details
	if state.DeploymentDetails != nil &&
		state.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpcBlock(state, result)
	}

	if state.DeploymentDetails == nil {
		resp.Diagnostics.AddError(
			constants.UnexpectedSystemType,
			constants.UnexpectedSystemType,
		)
	}

	result.ActivationClientModel = helper.SetPowerflexClientState(*state.ActivationClientModel)

	// Set refresdhed state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *blockStorageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan models.StorageModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Retrieve values from state
	var state models.StorageModel
	diagsState := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Allow the user to update the activation user/password
	if (plan.ActivationClientModel.Username.ValueString() != state.ActivationClientModel.Username.ValueString()) ||
		(plan.ActivationClientModel.Password.ValueString() != state.ActivationClientModel.Password.ValueString()) ||
		(plan.ActivationClientModel.Scheme.ValueString() != state.ActivationClientModel.Scheme.ValueString()) ||
		(plan.ActivationClientModel.Host.ValueString() != state.ActivationClientModel.Host.ValueString()) {
		state.ActivationClientModel = helper.SetPowerflexClientState(*plan.ActivationClientModel)
		state.StorageSystemType = plan.StorageSystemType
		if strings.Contains(state.Version.ValueString(), plan.ProductVersion.ValueString()) {
			state.ProductVersion = plan.ProductVersion
		}

		// Need to check for on prem deployment details
		if state.DeploymentDetails != nil &&
			state.DeploymentDetails.SystemPublicCloud != nil {
			ioOps := state.DeploymentDetails.SystemPublicCloud.MinimumIops
			minCap := state.DeploymentDetails.SystemPublicCloud.MinimumCapacity
			if len(plan.DeploymentDetails.SystemPublicCloud.SubnetOptions) == 0 {
				helper.SetCloudConfigSubnetAndVpcBlock(plan, state)

			}
			state.DeploymentDetails.SystemPublicCloud.MinimumIops = ioOps
			state.DeploymentDetails.SystemPublicCloud.MinimumCapacity = minCap
		}

		// Set refresh the state with the updated poweflex client
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		return
	}

	// If it is anything besides the activation username/password show an error
	if state != plan {
		resp.Diagnostics.AddError(
			constants.BlockStorageUpdateErrorMsg,
			constants.BlockStorageUpdateDetailMsg,
		)
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *blockStorageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan models.StorageModel
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

	// Delete existing block storage
	req2 := r.client.StorageSystemsAPI.StorageSystemsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockStorageDecomissionErrorMsg,
			constants.BlockStorageDecomissionDetailMsg+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, jobErr := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if jobErr != nil || resourceID == "" {
		resp.Diagnostics.AddError(
			constants.GeneralJobError,
			jobErr.Error(),
		)
		return
	}

}

func (r *blockStorageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Scheme   string `json:"scheme"`
		Insecure bool   `json:"insecure"`
		SystemID string `json:"system_id"`
	}
	var p params
	err := json.Unmarshal([]byte(req.ID), &p)
	if err != nil {
		resp.Diagnostics.AddError("make sure you include system_id, username, password, host, scheme, insecure", err.Error())
	}

	getReq := r.client.StorageSystemsAPI.StorageSystemsInstance(ctx, p.SystemID)

	// Get refreshed storage systems value from Apex Navigator
	storageSystem, status, err := getReq.Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.StorageReadErrorMsg,
			constants.StorageReadDetailMsg+newErr,
		)
		return
	}

	result := helper.GetStorageSystemBlock(*storageSystem)
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

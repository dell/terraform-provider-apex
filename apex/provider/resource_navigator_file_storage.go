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
	_ resource.Resource                = &fileStorageResource{}
	_ resource.ResourceWithConfigure   = &fileStorageResource{}
	_ resource.ResourceWithImportState = &fileStorageResource{}
)

// NewFileStorageResource returns a storage system resource object
func NewFileStorageResource() resource.Resource {
	return &fileStorageResource{}
}

// fileStorageResource is the resource implementation.
type fileStorageResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *fileStorageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_file_storage"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *fileStorageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen

	resp.Schema = schema.Schema{
		Attributes: GetStorageSystemSchema("file"),
		Blocks: map[string]schema.Block{
			"powerflex": schema.SingleNestedBlock{
				Attributes: PowerflexInfo.Attributes,
			},
		},
	}
}

// Configures the provider configured client to the data source.
func (r *fileStorageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// r.fileStorage.Configure(ctx, req, resp)
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
func (r *fileStorageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.StorageModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var systemCreateInput client.StorageSystemDeploymentRequest
	switch {
	case plan.DeploymentDetails == nil:
		break
	case plan.DeploymentDetails.SystemPublicCloud != nil:
		tier := client.PowerScaleTierEnum(plan.DeploymentDetails.SystemPublicCloud.TierType.ValueString())
		rawCapacity := plan.DeploymentDetails.SystemPublicCloud.RawCapacity.ValueStringPointer()

		storageParams := client.StorageSystemDeploymentRequestStorageOptions{
			PowerScaleStorageOptions: &client.PowerScaleStorageOptions{
				SystemType:  plan.StorageSystemType.ValueString(),
				Version:     plan.ProductVersion.ValueStringPointer(),
				Tier:        &tier,
				RawCapacity: rawCapacity,
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

		cloudParams := client.StorageSystemDeploymentRequestCloudOptions{
			AWSCloudOptions: &client.AWSCloudOptions{
				CloudType:                string(*plan.DeploymentDetails.SystemPublicCloud.CloudType),
				AccountId:                plan.DeploymentDetails.SystemPublicCloud.CloudAccount.ValueString(),
				Region:                   plan.DeploymentDetails.SystemPublicCloud.CloudRegion.ValueString(),
				AvailabilityZoneTopology: plan.DeploymentDetails.SystemPublicCloud.AvailabilityZoneTopology.Ptr(),
				SshKeyName:               plan.DeploymentDetails.SystemPublicCloud.SSHKeyName.ValueString(),
				IamInstanceProfile:       plan.DeploymentDetails.SystemPublicCloud.IAMInstanceProfile.ValueStringPointer(),
				Vpc: client.Vpc{
					Id:   plan.DeploymentDetails.SystemPublicCloud.Vpc.VpcID.ValueStringPointer(),
					Name: plan.DeploymentDetails.SystemPublicCloud.Vpc.VpcName.ValueStringPointer(),
				},
				SubnetOptions: subnetOptions,
			},
		}
		systemCreateInput = *client.NewStorageSystemDeploymentRequest(plan.Name.ValueString(), cloudParams, storageParams, true)
	case plan.DeploymentDetails.SystemOnPrem != nil:
		defaultValue := "default"
		rawCapacity := "1"
		storageOptions := client.StorageSystemDeploymentRequestStorageOptions{
			PowerScaleStorageOptions: &client.PowerScaleStorageOptions{
				SystemType:  defaultValue,
				Version:     &defaultValue,
				Tier:        client.POWERSCALETIERENUM_BALANCED.Ptr(),
				RawCapacity: &rawCapacity,
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
		defaultProfile := "profile"
		cloudOptions := client.StorageSystemDeploymentRequestCloudOptions{
			AWSCloudOptions: &client.AWSCloudOptions{
				CloudType:                "default",
				AccountId:                "default",
				Region:                   "default",
				AvailabilityZoneTopology: client.AVAILABILITYZONETOPOLOGYENUM_MULTIPLE_AVAILABILITY_ZONE.Ptr(),
				SshKeyName:               "default",
				IamInstanceProfile:       &defaultProfile,
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
			constants.FileStorageCreateErrorMsg,
			constants.StorageSystemInvalidDeploymentType,
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Creating File Storage: %v", systemCreateInput))
	job, status, err := helper.CreateStorageSystem(r.client.StorageSystemsAPI.StorageSystemsCreate(ctx), systemCreateInput)

	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.FileStorageCreateErrorMsg,
			constants.FileStorageCreateDetailMsg+newErr,
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

	// Fetching File storage after Job Completes
	storageSystem, status, err := helper.GetStorageInstance(r.client, resourceID)
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		// Try again with POWERSCALE- prefix
		if status.StatusCode == http.StatusNotFound {
			resourceIDWithPrefix := "ONEFS-" + resourceID
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

	// Updating TFState with File Storage info
	result := helper.GetStorageSystem(*storageSystem)
	if result.SystemType.ValueString() == "ISILON" {
		result.StorageSystemType = plan.StorageSystemType
	}

	if strings.Contains(result.Version.ValueString(), plan.ProductVersion.ValueString()) {
		result.ProductVersion = plan.ProductVersion
	}

	// Need to check for cloud deployment details
	if plan.DeploymentDetails != nil &&
		plan.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpc(plan, result)
		result.DeploymentDetails.SystemPublicCloud.RawCapacity = plan.DeploymentDetails.SystemPublicCloud.RawCapacity
		result.DeploymentDetails.SystemPublicCloud.IAMInstanceProfile = plan.DeploymentDetails.SystemPublicCloud.IAMInstanceProfile

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
func (r *fileStorageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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
	result := helper.GetStorageSystem(*storageSystem)
	if result.SystemType.ValueString() == "ISILON" {
		result.StorageSystemType = state.StorageSystemType
	}

	if strings.Contains(result.Version.ValueString(), state.ProductVersion.ValueString()) {
		result.ProductVersion = state.ProductVersion
	}

	// Need to check for on prem deployment details
	if state.DeploymentDetails != nil &&
		state.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpc(state, result)
		result.DeploymentDetails.SystemPublicCloud.RawCapacity = state.DeploymentDetails.SystemPublicCloud.RawCapacity
		result.DeploymentDetails.SystemPublicCloud.IAMInstanceProfile = state.DeploymentDetails.SystemPublicCloud.IAMInstanceProfile

	}

	if state.DeploymentDetails == nil {
		resp.Diagnostics.AddError(
			constants.UnexpectedSystemType,
			constants.UnexpectedSystemType,
		)
	}

	result.ActivationClientModel = helper.SetPowerflexClientState(*state.ActivationClientModel)

	// Set refreshed state
	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *fileStorageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
		if strings.Contains(state.Version.ValueString(), plan.ProductVersion.ValueString()) {
			state.ProductVersion = plan.ProductVersion
		}
		state.StorageSystemType = plan.StorageSystemType

		// Need to check for on prem deployment details
		if state.DeploymentDetails != nil &&
			state.DeploymentDetails.SystemPublicCloud != nil {
			ioOps := state.DeploymentDetails.SystemPublicCloud.MinimumIops
			minCap := state.DeploymentDetails.SystemPublicCloud.MinimumCapacity
			if len(plan.DeploymentDetails.SystemPublicCloud.SubnetOptions) == 0 {
				helper.SetCloudConfigSubnetAndVpc(plan, state)
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
			constants.FileStorageUpdateErrorMsg,
			constants.FileStorageUpdateDetailMsg,
		)
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *fileStorageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan models.StorageModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update the DI token for subsquent post calls to activate the system
	_ = helper.UpdateToken(ctx, r.client, r.jobsClient)

	// Activate PowerScale
	actErr := helper.ActivateSystemClientSystem(ctx, r.client, plan.SystemID.ValueString(), *plan.ActivationClientModel, client.STORAGEPRODUCTENUM_POWERSCALE)
	if actErr != nil {
		resp.Diagnostics.AddError(
			constants.ErrorActivatingPowerScaleSystem,
			constants.ErrorActivatingPowerScaleSystemDetail+actErr.Error(),
		)
		return
	}

	// Delete existing storage system
	req2 := r.client.StorageSystemsAPI.StorageSystemsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.FileStorageDecomissionErrorMsg,
			constants.FileStorageDecomissionDetailMsg+newErr,
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

func (r *fileStorageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

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

	//  get a new token for GET Storage system
	_ = helper.UpdateToken(ctx, r.client, r.jobsClient)

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

	result := helper.GetStorageSystem(*storageSystem)

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

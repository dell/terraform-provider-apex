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
	"strings"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	jmsClient "github.com/dell/terraform-provider-apex/client/jobsclient/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	resp.TypeName = req.ProviderTypeName + "_block_storage"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *blockStorageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"system_id": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"system_type": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"capacity_impact": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"capacity_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"compression_savings": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"configuration_impact": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"configuration_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"configured_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"connectivity_status": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"license_type": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"license_expiration_date_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"contract_coverage_type": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"cirrus_deployed": schema.BoolAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"contract_expiration_date_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"data_protection_impact": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"data_protection_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"display_identifier": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"free_percent": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"free_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"health_connectivity_status": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"health_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"health_score": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"health_state": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"iops": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"last_contact_timestamp": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"latency": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"logical_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"model": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"overall_efficiency": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"performance_impact": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"performance_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"site_name": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"snaps_savings": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"system_health_impact": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"system_health_issue_count": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"thin_savings": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"total_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"unconfigured_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"used_percent": schema.Float64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"used_size": schema.Int64Attribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"vendor": schema.StringAttribute{
				MarkdownDescription: " ",
				Optional:            true,
				Computed:            true,
			},
			"product_version": schema.StringAttribute{
				MarkdownDescription: " ",
				Required:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: " ",
				Computed:            true,
			},
			"deployment_details": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"system_on_prem": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"deployment_type": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"site_name": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"location": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"country": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"state": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"city": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"street_address_1": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"street_address_2": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"zip_code": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
						},
					},
					"system_public_cloud": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"deployment_type": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"cloud_type": schema.StringAttribute{
								Required: true,
							},
							"cloud_account": schema.StringAttribute{
								Required: true,
							},
							"cloud_region": schema.StringAttribute{
								Required: true,
							},
							"availability_zone_topology": schema.StringAttribute{
								Optional: true,
							},
							"virtual_private_cloud": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"cloud_management_address": schema.StringAttribute{
								Optional: true,
								Computed: true,
							},
							"minimum_iops": schema.Int64Attribute{
								Required: true,
							},
							"minimum_capacity": schema.Int64Attribute{
								Required: true,
							},
							"tier_type": schema.StringAttribute{
								Required: true,
							},
							"ssh_key_name": schema.StringAttribute{
								Required: true,
							},
							"vpc": schema.SingleNestedAttribute{
								Required: true,
								Attributes: map[string]schema.Attribute{
									"is_new_vpc": schema.BoolAttribute{
										Optional: true,
									},
									"vpc_id": schema.StringAttribute{
										Optional: true,
									},
									"vpc_name": schema.StringAttribute{
										Optional: true,
									},
								},
							},
							"subnet_options": schema.ListNestedAttribute{
								Required: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"subnet_id": schema.StringAttribute{
											Optional: true,
										},
										"cidr_block": schema.StringAttribute{
											Optional: true,
										},
										"subnet_type": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
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
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *provider.CLients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *blockStorageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit
	// Retrieve values from plan
	var plan models.BlockStorageModel
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
		tier := client.PowerFlexTierEnum(plan.DeploymentDetails.SystemPublicCloud.TierType.ValueString())
		minIops := int32(plan.DeploymentDetails.SystemPublicCloud.MinimumIops.ValueInt64())
		minUC := int32(plan.DeploymentDetails.SystemPublicCloud.MinimumCapacity.ValueInt64())

		storageParams := client.StorageSystemDeploymentRequestStorageOptions{
			PowerFlexStorageOptions: &client.PowerFlexStorageOptions{
				SystemType:            plan.SystemType.ValueString(),
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

		cloudParams := client.StorageSystemDeploymentRequestCloudOptions{
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
		storageOptions := client.StorageSystemDeploymentRequestStorageOptions{
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
		cloudOptions := client.StorageSystemDeploymentRequestCloudOptions{
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
			"Error creating block storage post request",
			"Could not create block storage post request",
		)
		return
	}
	job, status, err := helper.CreateBlockStorage(r.client.StorageSystemsAPI.StorageSystemsCreate(ctx), systemCreateInput)

	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error creating Block Storage",
			"Could not create Block Storage, unexpected error: "+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, jobErr := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if jobErr != nil {
		resp.Diagnostics.AddError(
			"Error getting resourceID",
			helper.ResourceRetrieveError+jobErr.Error(),
		)
		return
	}

	// Fetching Block storage after Job Completes
	storageSystem, status, err := helper.GetBlockStorageInstance(r.client, resourceID)
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		// Try again with POWERFLEX- prefix
		if status.StatusCode == http.StatusNotFound {
			resourceIDWithPrefix := "POWERFLEX-" + resourceID
			storageSystem, status, err = helper.GetBlockStorageInstance(r.client, resourceIDWithPrefix)
			if (err != nil) || (status.StatusCode != http.StatusOK) {
				newErr := helper.GetErrorString(err, status)
				resp.Diagnostics.AddError(
					"Error retrieving created Block storage",
					"Could not retrieve created Block storage, unexpected error: "+newErr,
				)
				return
			}
		} else {
			newErr := helper.GetErrorString(err, status)
			resp.Diagnostics.AddError(
				"Error retrieving created Block storage",
				"Could not retrieve created Block storage, unexpected error: "+newErr,
			)
			return
		}
	}

	// Updating TFState with Block Storage info
	result := helper.GetBlockStorageSystem(*storageSystem)

	if strings.Contains(result.Version.ValueString(), plan.ProductVersion.ValueString()) {
		result.ProductVersion = plan.ProductVersion
	}
	// Need to check for cloud deployment details
	if plan.DeploymentDetails != nil &&
		plan.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpc(plan, result)
	}

	if result.DeploymentDetails == nil {
		resp.Diagnostics.AddError(
			"Unexpected system type",
			"Unexpected system type",
		)
	}

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
	var state models.BlockStorageModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed storage systems value from Apex Navigator
	storageSystem, status, err := helper.GetBlockStorageInstance(r.client, state.ID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator block storage",
			"Could not read Apex Navigator block storage, unexpected error: "+newErr,
		)
		return
	}

	// Overwrite items with refreshed state
	result := helper.GetBlockStorageSystem(*storageSystem)

	if strings.Contains(result.Version.ValueString(), state.ProductVersion.ValueString()) {
		result.ProductVersion = state.ProductVersion
	}

	// Need to check for on prem deployment details
	if state.DeploymentDetails != nil &&
		state.DeploymentDetails.SystemPublicCloud != nil {
		helper.SetCloudConfigSubnetAndVpc(state, result)
	}

	if state.DeploymentDetails == nil {
		resp.Diagnostics.AddError(
			"Unexpected system type",
			"Unexpected system type",
		)
	}

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
	var plan models.BlockStorageModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Retrieve values from state
	var state models.BlockStorageModel
	diagsState := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state != plan {
		resp.Diagnostics.AddError(
			"Unable to update Apex Navigator block storage",
			"Update of Block Storage is not supported, please create a new resource",
		)
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *blockStorageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan models.BlockStorageModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing block storage
	req2 := r.client.StorageSystemsAPI.StorageSystemsDelete(ctx, plan.ID.ValueString())

	job, status, err := req2.Async(true).Execute()
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Deleting Apex Navigator Storage System",
			"Could not delete block storage, unexpected error: "+newErr,
		)
		return
	}

	// Fetching Job Status
	resourceID, jobErr := helper.WaitForJobToComplete(ctx, r.jobsClient, job.Id)
	if jobErr != nil || resourceID == "" {
		resp.Diagnostics.AddError(
			"Error getting Delete Job ID",
			"Could not retrieve delete job id, unexpected error: "+jobErr.Error(),
		)
		return
	}

}

func (r *blockStorageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

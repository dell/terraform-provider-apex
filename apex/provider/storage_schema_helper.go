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
	"net/http"

	client "dell/apex-client"
	jmsClient "dell/apex-job-client"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// StorageSystemResource is a resource that represents a storage system
type StorageSystemResource struct {
	StorageSystem resource.Resource
	client        *client.APIClient
	jobsClient    *jmsClient.APIClient
}

// GetStorageSystemSchema returns a map of attribute names and types
func GetStorageSystemSchema(systemType string) map[string]schema.Attribute {

	var onPremDetails map[string]schema.Attribute = map[string]schema.Attribute{
		"deployment_type": schema.StringAttribute{
			Description:         "System deployment types (e.g. PUBLIC_CLOUD)",
			MarkdownDescription: "System deployment types (e.g. PUBLIC_CLOUD)",
			Optional:            true,
			Computed:            true,
		},
		"site_name": schema.StringAttribute{
			MarkdownDescription: "Name of the site where the system is located",
			Description:         "Name of the site where the system is located",
			Optional:            true,
			Computed:            true,
		},
		"location": schema.StringAttribute{
			Description:         "User provided description of where the system is located.",
			MarkdownDescription: "User provided description of where the system is located.",
			Optional:            true,
			Computed:            true,
		},
		"country": schema.StringAttribute{
			Description:         "Name of the country where the system is located.",
			MarkdownDescription: "Name of the country where the system is located.",
			Optional:            true,
			Computed:            true,
		},
		"state": schema.StringAttribute{
			Description:         "Name of the state where the system is located.",
			MarkdownDescription: "Name of the state where the system is located.",
			Optional:            true,
			Computed:            true,
		},
		"city": schema.StringAttribute{
			Description:         "Name of the city where the system is located.",
			MarkdownDescription: "Name of the city where the system is located.",
			Optional:            true,
			Computed:            true,
		},
		"street_address_1": schema.StringAttribute{
			Description:         "Street address 1 of where the system is located",
			MarkdownDescription: "Street address 1 of where the system is located",
			Optional:            true,
			Computed:            true,
		},
		"street_address_2": schema.StringAttribute{
			Description:         "Street address 2 of where the system is located",
			MarkdownDescription: "Street address 2 of where the system is located",
			Optional:            true,
			Computed:            true,
		},
		"zip_code": schema.StringAttribute{
			Description:         "Zip code of where the system is located",
			MarkdownDescription: "Zip code of where the system is located",
			Optional:            true,
			Computed:            true,
		},
	}

	var publicCloudDetails map[string]schema.Attribute = map[string]schema.Attribute{
		"deployment_type": schema.StringAttribute{
			Description:         "System deployment types (e.g. PUBLIC_CLOUD)",
			MarkdownDescription: "System deployment types (e.g. PUBLIC_CLOUD)",
			Optional:            true,
			Computed:            true,
		},
		"cloud_type": schema.StringAttribute{
			Description:         "Enum for all the supported cloud providers * AWS - Amazon Web Services",
			MarkdownDescription: "Enum for all the supported cloud providers * AWS - Amazon Web Services",
			Required:            true,
		},
		"cloud_account": schema.StringAttribute{
			Description:         "Cloud provider account where the storage system resides",
			MarkdownDescription: "Cloud provider account where the storage system resides",
			Required:            true,
		},
		"cloud_region": schema.StringAttribute{
			Description:         "Cloud provider region where the storage system resides",
			MarkdownDescription: "Cloud provider region where the storage system resides",
			Required:            true,
		},
		"availability_zone_topology": schema.StringAttribute{
			Description:         "This enum represents all the availability zone topology * SINGLE_AVAILABILITY_ZONE * MULTIPLE_AVAILABILITY_ZONE",
			MarkdownDescription: "This enum represents all the availability zone topology * SINGLE_AVAILABILITY_ZONE * MULTIPLE_AVAILABILITY_ZONE",
			Optional:            true,
		},
		"virtual_private_cloud": schema.StringAttribute{
			Description:         "Cloud virtual private environment identifier",
			MarkdownDescription: "Cloud virtual private environment identifier",
			Optional:            true,
			Computed:            true,
		},
		"cloud_management_address": schema.StringAttribute{
			Description:         "Management IPv4 or IPv6 address or DNS name of the storage system in cloud",
			MarkdownDescription: "Management IPv4 or IPv6 address or DNS name of the storage system in cloud",
			Optional:            true,
			Computed:            true,
		},

		"tier_type": schema.StringAttribute{
			Description:         "Tier type requested during the deployment time",
			MarkdownDescription: "Tier type requested during the deployment time",
			Required:            true,
		},
		"ssh_key_name": schema.StringAttribute{
			Required: true,
		},
		"vpc": schema.SingleNestedAttribute{
			Description:         "VPC",
			MarkdownDescription: "VPC",
			Required:            true,
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
			Description:         "Subnet options",
			MarkdownDescription: "Subnet options",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"subnet_id": schema.StringAttribute{
						Optional: true,
					},
					"cidr_block": schema.StringAttribute{
						Description:         "CIDR block for subnet. It is not supported when existing vpc is used to create block storage",
						MarkdownDescription: "CIDR block for subnet. It is not supported when existing vpc is used to create block storage",
						Optional:            true,
					},
					"subnet_type": schema.StringAttribute{
						Description:         "subnet type is one of EXTERNAL, INTERNAL, SCG",
						MarkdownDescription: "subnet type is one of EXTERNAL, INTERNAL, SCG",
						Optional:            true,
						Validators: []validator.String{stringvalidator.OneOf(
							"EXTERNAL",
							"INTERNAL",
							"SCG",
						)},
					},
				},
			},
		},
	}

	switch systemType {
	case "block":
		var minIops schema.Int64Attribute = schema.Int64Attribute{
			Description:         "Minimum IOPS requested during the deployment time - Unit: IO/s",
			MarkdownDescription: "Minimum IOPS requested during the deployment time - Unit: IO/s",
			Required:            true,
		}
		var minCapacity schema.Int64Attribute = schema.Int64Attribute{
			Description:         "Minimum capacity requested during the deployment time - Unit: bytes",
			MarkdownDescription: "Minimum capacity requested during the deployment time - Unit: bytes",
			Required:            true,
		}
		var rawCapacity schema.StringAttribute = schema.StringAttribute{
			Description:         "Raw capacity requested during the deployment time - Unit: bytes",
			MarkdownDescription: "Raw capacity requested during the deployment time - Unit: bytes",
			Optional:            true,
		}
		var iamProfile schema.StringAttribute = schema.StringAttribute{
			Description:         "IAM instance profile requested during the deployment time - Unit: string",
			MarkdownDescription: "IAM instance profile requested during the deployment time - Unit: string",
			Optional:            true,
		}
		publicCloudDetails["minimum_iops"] = minIops
		publicCloudDetails["minimum_capacity"] = minCapacity
		publicCloudDetails["raw_capacity"] = rawCapacity
		publicCloudDetails["iam_instance_profile"] = iamProfile

	case "file":
		var minIops schema.Int64Attribute = schema.Int64Attribute{
			Description:         "Minimum IOPS requested during the deployment time - Unit: IO/s",
			MarkdownDescription: "Minimum IOPS requested during the deployment time - Unit: IO/s",
			Optional:            true,
			Computed:            true,
		}
		var minCapacity schema.Int64Attribute = schema.Int64Attribute{
			Description:         "Minimum capacity requested during the deployment time - Unit: bytes",
			MarkdownDescription: "Minimum capacity requested during the deployment time - Unit: bytes",
			Optional:            true,
			Computed:            true,
		}
		var rawCapacity schema.StringAttribute = schema.StringAttribute{
			Description:         "Raw capacity requested during the deployment time - Unit: bytes",
			MarkdownDescription: "Raw capacity requested during the deployment time - Unit: bytes",
			Required:            true,
		}
		var iamProfile schema.StringAttribute = schema.StringAttribute{
			Description:         "IAM instance profile requested during the deployment time - Unit: string",
			MarkdownDescription: "IAM instance profile requested during the deployment time - Unit: string",
			Required:            true,
		}
		publicCloudDetails["minimum_iops"] = minIops
		publicCloudDetails["minimum_capacity"] = minCapacity
		publicCloudDetails["raw_capacity"] = rawCapacity
		publicCloudDetails["iam_instance_profile"] = iamProfile

	default:
		// error to provide valid system type
	}

	var storageSystemSchemaAttr map[string]schema.Attribute = map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Identifier for the storage system",
			Description:         "Identifier for the storage system",
			Computed:            true,
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"system_id": schema.StringAttribute{
			MarkdownDescription: "Unique identifier for the device or appliance",
			Description:         "Unique identifier for the device or appliance",
			Optional:            true,
			Computed:            true,
		},
		"storage_system_type": schema.StringAttribute{
			MarkdownDescription: "Type of the system",
			Description:         "Type of the system",
			Required:            true,
		},
		"system_type": schema.StringAttribute{
			MarkdownDescription: "Type of the system",
			Description:         "Type of the system",
			Computed:            true,
		},
		"bandwidth": schema.Int64Attribute{
			MarkdownDescription: "The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
			Description:         "The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
			Optional:            true,
			Computed:            true,
		},
		"capacity_impact": schema.Int64Attribute{
			MarkdownDescription: "Impact point of highest impacting issue in the capacity health category",
			Description:         "Impact point of highest impacting issue in the capacity health category",
			Optional:            true,
			Computed:            true,
		},
		"capacity_issue_count": schema.Int64Attribute{
			MarkdownDescription: "Total number of issues in the capacity health category",
			Description:         "Total number of issues in the capacity health category",
			Optional:            true,
			Computed:            true,
		},
		"compression_savings": schema.Float64Attribute{
			Description:         "Storage efficiency ratio of data which has compression applied to it on the system.",
			MarkdownDescription: "Storage efficiency ratio of data which has compression applied to it on the system.",
			Optional:            true,
			Computed:            true,
		},
		"configuration_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the configuration health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the configuration health category.",
			Optional:            true,
			Computed:            true,
		},
		"configuration_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the configuration health category.",
			MarkdownDescription: "Total number of issues in the configuration health category.",
			Optional:            true,
			Computed:            true,
		},
		"configured_size": schema.Int64Attribute{
			Description:         "The configured size for this system - Unit: bytes.",
			MarkdownDescription: "The configured size for this system - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"connectivity_status": schema.StringAttribute{
			Description:         "Connectivity status.",
			MarkdownDescription: "Connectivity status.",
			Optional:            true,
			Computed:            true,
		},
		"license_type": schema.StringAttribute{
			Description:         "Type of the license on the system.",
			MarkdownDescription: "Type of the license on the system.",
			Optional:            true,
			Computed:            true,
		},
		"license_expiration_date_timestamp": schema.StringAttribute{
			Description:         "Expiration date for the license on the system.",
			MarkdownDescription: "Expiration date for the license on the system.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"contract_coverage_type": schema.StringAttribute{
			Description:         "Type of the service contract of the system.",
			MarkdownDescription: "Type of the service contract of the system.",
			Optional:            true,
			Computed:            true,
		},
		"contract_expiration_date_timestamp": schema.StringAttribute{
			Description:         "Expiration date for the service contract of the system.",
			MarkdownDescription: "Expiration date for the service contract of the system.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"data_protection_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the data protection health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the data protection health category.",
			Optional:            true,
			Computed:            true,
		},
		"data_protection_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the data protection health category.",
			MarkdownDescription: "Total number of issues in the data protection health category.",
			Optional:            true,
			Computed:            true,
		},
		"display_identifier": schema.StringAttribute{
			Description:         "Unique identifier for the system.",
			MarkdownDescription: "Unique identifier for the system.",
			Optional:            true,
			Computed:            true,
		},
		"free_percent": schema.Float64Attribute{
			Description:         "The %free capacity.",
			MarkdownDescription: "The %free capacity.",
			Optional:            true,
			Computed:            true,
		},
		"free_size": schema.Int64Attribute{
			Description:         "The free size value - Unit: bytes.",
			MarkdownDescription: "The free size value - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"health_connectivity_status": schema.StringAttribute{
			Description:         "Health connectivity status.",
			MarkdownDescription: "Health connectivity status.",
			Optional:            true,
			Computed:            true,
		},
		"health_issue_count": schema.Int64Attribute{
			Description:         "Total amount of health issues.",
			MarkdownDescription: "Total amount of health issues.",
			Optional:            true,
			Computed:            true,
		},
		"health_score": schema.Int64Attribute{
			Description:         "The overall health score of the system.",
			MarkdownDescription: "The overall health score of the system.",
			Optional:            true,
			Computed:            true,
		},
		"health_state": schema.StringAttribute{
			Description:         "Health state of the system.",
			MarkdownDescription: "Health state of the system.",
			Optional:            true,
			Computed:            true,
		},
		"iops": schema.Int64Attribute{
			Description:         "The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s.",
			MarkdownDescription: "The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s.",
			Optional:            true,
			Computed:            true,
		},
		"ipv4_address": schema.StringAttribute{
			Description:         "IPv4 address of the system.",
			MarkdownDescription: "IPv4 address of the system.",
			Optional:            true,
			Computed:            true,
		},
		"ipv6_address": schema.StringAttribute{
			Description:         "IPv6 address of the system.",
			MarkdownDescription: "IPv6 address of the system.",
			Optional:            true,
			Computed:            true,
		},
		"last_contact_timestamp": schema.StringAttribute{
			Description:         "Last time that CloudIQ received data from the system.",
			MarkdownDescription: "Last time that CloudIQ received data from the system.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"latency": schema.Int64Attribute{
			Description:         "The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds.",
			MarkdownDescription: "The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds.",
			Optional:            true,
			Computed:            true,
		},
		"logical_size": schema.Int64Attribute{
			Description:         "The logical size written - Unit: bytes.",
			MarkdownDescription: "The logical size written - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"model": schema.StringAttribute{
			Description:         "The model of the system.",
			MarkdownDescription: "The model of the system.",
			Optional:            true,
			Computed:            true,
		},
		"name": schema.StringAttribute{
			Description:         "The user-defined name of the system.",
			MarkdownDescription: "The user-defined name of the system.",
			Required:            true,
		},
		"overall_efficiency": schema.Float64Attribute{
			Description:         "The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction.",
			MarkdownDescription: "The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction.",
			Optional:            true,
			Computed:            true,
		},
		"performance_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the performance health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the performance health category.",
			Optional:            true,
			Computed:            true,
		},
		"performance_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the performance health category.",
			MarkdownDescription: "Total number of issues in the performance health category.",
			Optional:            true,
			Computed:            true,
		},
		"serial_number": schema.StringAttribute{
			Description:         "The serial number for this system.",
			MarkdownDescription: "The serial number for this system.",
			Optional:            true,
			Computed:            true,
		},
		"site_name": schema.StringAttribute{
			Description:         "Name of the site where the system is registered to.",
			MarkdownDescription: "Name of the site where the system is registered to.",
			Optional:            true,
			Computed:            true,
		},
		"snaps_savings": schema.Float64Attribute{
			Description:         "The snaps savings for this system.",
			MarkdownDescription: "The snaps savings for this system.",
			Optional:            true,
			Computed:            true,
		},
		"system_health_impact": schema.Int64Attribute{
			Description:         "Health impact for the system.",
			MarkdownDescription: "Health impact for the system.",
			Optional:            true,
			Computed:            true,
		},
		"system_health_issue_count": schema.Int64Attribute{
			Description:         "Total amount of health issues for the system.",
			MarkdownDescription: "Total amount of health issues for the system.",
			Optional:            true,
			Computed:            true,
		},
		"thin_savings": schema.Float64Attribute{
			Description:         "The savings due to thin provisioning.",
			MarkdownDescription: "The savings due to thin provisioning.",
			Optional:            true,
			Computed:            true,
		},
		"total_size": schema.Int64Attribute{
			Description:         "The total size of the system - Unit: bytes.",
			MarkdownDescription: "The total size of the system - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"unconfigured_size": schema.Int64Attribute{
			Description:         "The unconfigured capacity for this system - Unit: bytes.",
			MarkdownDescription: "The unconfigured capacity for this system - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"used_percent": schema.Float64Attribute{
			Description:         "Percentage of capacity used for this system.",
			MarkdownDescription: "Percentage of capacity used for this system.",
			Optional:            true,
			Computed:            true,
		},
		"used_size": schema.Int64Attribute{
			Description:         "The value of used capacity for this system - Unit: bytes.",
			MarkdownDescription: "The value of used capacity for this system - Unit: bytes.",
			Optional:            true,
			Computed:            true,
		},
		"vendor": schema.StringAttribute{
			Description:         "Name of the vendor who makes the system.",
			MarkdownDescription: "Name of the vendor who makes the system.",
			Optional:            true,
			Computed:            true,
		},
		"product_version": schema.StringAttribute{
			MarkdownDescription: "Product version.",
			Description:         "Product version.",
			Required:            true,
		},
		"version": schema.StringAttribute{
			Description:         "Version identifier.",
			MarkdownDescription: "Version identifier.",
			Computed:            true,
		},
		"deployment_details": schema.SingleNestedAttribute{
			Optional:            true,
			MarkdownDescription: "Details of deployment",
			Description:         "Details of deployment",
			Attributes: map[string]schema.Attribute{
				"system_on_prem": schema.SingleNestedAttribute{
					Optional:   true,
					Attributes: onPremDetails,
				},
				"system_public_cloud": schema.SingleNestedAttribute{
					Optional:   true,
					Attributes: publicCloudDetails,
				},
			},
		},
	}

	return storageSystemSchemaAttr
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *StorageSystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:dupl
	// Retrieve values from state
	var plan models.StorageModel
	diags := req.State.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing storage system
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

// ImportState imports the Terraform resource
func (r *StorageSystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

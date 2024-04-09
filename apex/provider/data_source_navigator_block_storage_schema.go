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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// BlockStorageDataSourceSchema defines the schema for a storage system instance data source and
// its attributes are used in storage systems collection data source
var BlockStorageDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Identifier for the storage system",
			Description:         "Identifier for the storage system",
			Optional:            true,
		},
		"system_id": schema.StringAttribute{
			MarkdownDescription: "Unique identifier for the device or appliance",
			Description:         "Unique identifier for the device or appliance",
			Computed:            true,
		},
		"system_type": schema.StringAttribute{
			MarkdownDescription: "Type of the system",
			Description:         "Type of the system",
			Computed:            true,
		},
		"bandwidth": schema.Int64Attribute{
			MarkdownDescription: "The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
			Description:         "The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
			Computed:            true,
		},
		"capacity_impact": schema.Int64Attribute{
			MarkdownDescription: "Impact point of highest impacting issue in the capacity health category",
			Description:         "Impact point of highest impacting issue in the capacity health category",
			Computed:            true,
		},
		"capacity_issue_count": schema.Int64Attribute{
			MarkdownDescription: "Total number of issues in the capacity health category",
			Description:         "Total number of issues in the capacity health category",
			Computed:            true,
		},
		"compression_savings": schema.Float64Attribute{
			Description:         "Storage efficiency ratio of data which has compression applied to it on the system.",
			MarkdownDescription: "Storage efficiency ratio of data which has compression applied to it on the system.",
			Computed:            true,
		},
		"configuration_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the configuration health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the configuration health category.",
			Computed:            true,
		},
		"configuration_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the configuration health category.",
			MarkdownDescription: "Total number of issues in the configuration health category.",
			Computed:            true,
		},
		"configured_size": schema.Int64Attribute{
			Description:         "The configured size for this system - Unit: bytes.",
			MarkdownDescription: "The configured size for this system - Unit: bytes.",
			Computed:            true,
		},
		"connectivity_status": schema.StringAttribute{
			Description:         "Connectivity status.",
			MarkdownDescription: "Connectivity status.",
			Computed:            true,
		},
		"license_type": schema.StringAttribute{
			Description:         "Type of the license on the system.",
			MarkdownDescription: "Type of the license on the system.",
			Computed:            true,
		},
		"license_expiration_date_timestamp": schema.StringAttribute{
			Description:         "Expiration date for the license on the system.",
			MarkdownDescription: "Expiration date for the license on the system.",
			Computed:            true,
		},
		"contract_coverage_type": schema.StringAttribute{
			Description:         "Type of the service contract of the system.",
			MarkdownDescription: "Type of the service contract of the system.",
			Computed:            true,
		},
		"contract_expiration_date_timestamp": schema.StringAttribute{
			Description:         "Expiration date for the service contract of the system.",
			MarkdownDescription: "Expiration date for the service contract of the system.",
			Computed:            true,
		},
		"data_protection_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the data protection health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the data protection health category.",
			Computed:            true,
		},
		"data_protection_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the data protection health category.",
			MarkdownDescription: "Total number of issues in the data protection health category.",
			Computed:            true,
		},
		"display_identifier": schema.StringAttribute{
			Description:         "Unique identifier for the system.",
			MarkdownDescription: "Unique identifier for the system.",
			Computed:            true,
		},
		"free_percent": schema.Float64Attribute{
			Description:         "The %free capacity.",
			MarkdownDescription: "The %free capacity.",
			Computed:            true,
		},
		"free_size": schema.Int64Attribute{
			Description:         "The free size value - Unit: bytes.",
			MarkdownDescription: "The free size value - Unit: bytes.",
			Computed:            true,
		},
		"health_connectivity_status": schema.StringAttribute{
			Description:         "Health connectivity status.",
			MarkdownDescription: "Health connectivity status.",
			Computed:            true,
		},
		"health_issue_count": schema.Int64Attribute{
			Description:         "Total amount of health issues.",
			MarkdownDescription: "Total amount of health issues.",
			Computed:            true,
		},
		"health_score": schema.Int64Attribute{
			Description:         "The overall health score of the system.",
			MarkdownDescription: "The overall health score of the system.",
			Computed:            true,
		},
		"health_state": schema.StringAttribute{
			Description:         "Health state of the system.",
			MarkdownDescription: "Health state of the system.",
			Computed:            true,
		},
		"iops": schema.Int64Attribute{
			Description:         "The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s.",
			MarkdownDescription: "The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s.",
			Computed:            true,
		},
		"ipv4_address": schema.StringAttribute{
			Description:         "IPv4 address of the system.",
			MarkdownDescription: "IPv4 address of the system.",
			Computed:            true,
		},
		"ipv6_address": schema.StringAttribute{
			Description:         "IPv6 address of the system.",
			MarkdownDescription: "IPv6 address of the system.",
			Computed:            true,
		},
		"last_contact_timestamp": schema.StringAttribute{
			Description:         "Last time that CloudIQ received data from the system.",
			MarkdownDescription: "Last time that CloudIQ received data from the system.",
			Computed:            true,
		},
		"latency": schema.Int64Attribute{
			Description:         "The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds.",
			MarkdownDescription: "The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds.",
			Computed:            true,
		},
		"logical_size": schema.Int64Attribute{
			Description:         "The logical size written - Unit: bytes.",
			MarkdownDescription: "The logical size written - Unit: bytes.",
			Computed:            true,
		},
		"model": schema.StringAttribute{
			Description:         "The model of the system.",
			MarkdownDescription: "The model of the system.",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			Description:         "The user-defined name of the system.",
			MarkdownDescription: "The user-defined name of the system.",
			Computed:            true,
		},
		"overall_efficiency": schema.Float64Attribute{
			Description:         "The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction.",
			MarkdownDescription: "The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction.",
			Computed:            true,
		},
		"performance_impact": schema.Int64Attribute{
			Description:         "Impact point of highest impacting issue in the performance health category.",
			MarkdownDescription: "Impact point of highest impacting issue in the performance health category.",
			Computed:            true,
		},
		"performance_issue_count": schema.Int64Attribute{
			Description:         "Total number of issues in the performance health category.",
			MarkdownDescription: "Total number of issues in the performance health category.",
			Computed:            true,
		},
		"serial_number": schema.StringAttribute{
			Description:         "The serial number for this system.",
			MarkdownDescription: "The serial number for this system.",
			Computed:            true,
		},
		"site_name": schema.StringAttribute{
			Description:         "Name of the site where the system is registered to.",
			MarkdownDescription: "Name of the site where the system is registered to.",
			Computed:            true,
		},
		"snaps_savings": schema.Float64Attribute{
			Description:         "The snaps savings for this system.",
			MarkdownDescription: "The snaps savings for this system.",
			Computed:            true,
		},
		"system_health_impact": schema.Int64Attribute{
			Description:         "Health impact for the system.",
			MarkdownDescription: "Health impact for the system.",
			Computed:            true,
		},
		"system_health_issue_count": schema.Int64Attribute{
			Description:         "Total amount of health issues for the system.",
			MarkdownDescription: "Total amount of health issues for the system.",
			Computed:            true,
		},
		"thin_savings": schema.Float64Attribute{
			Description:         "The savings due to thin provisioning.",
			MarkdownDescription: "The savings due to thin provisioning.",
			Computed:            true,
		},
		"total_size": schema.Int64Attribute{
			Description:         "The total size of the system - Unit: bytes.",
			MarkdownDescription: "The total size of the system - Unit: bytes.",
			Computed:            true,
		},
		"unconfigured_size": schema.Int64Attribute{
			Description:         "The unconfigured capacity for this system - Unit: bytes.",
			MarkdownDescription: "The unconfigured capacity for this system - Unit: bytes.",
			Computed:            true,
		},
		"used_percent": schema.Float64Attribute{
			Description:         "Percentage of capacity used for this system.",
			MarkdownDescription: "Percentage of capacity used for this system.",
			Computed:            true,
		},
		"used_size": schema.Int64Attribute{
			Description:         "The value of used capacity for this system - Unit: bytes.",
			MarkdownDescription: "The value of used capacity for this system - Unit: bytes.",
			Computed:            true,
		},
		"vendor": schema.StringAttribute{
			Description:         "Name of the vendor who makes the system.",
			MarkdownDescription: "Name of the vendor who makes the system.",
			Computed:            true,
		},
		"product_version": schema.StringAttribute{
			MarkdownDescription: "Product version.",
			Description:         "Product version.",
			Computed:            true,
		},
		"version": schema.StringAttribute{
			Description:         "Version identifier.",
			MarkdownDescription: "Version identifier.",
			Computed:            true,
		},
		"cirrus_deployed": schema.BoolAttribute{
			MarkdownDescription: "If the system deployed in Cirrus",
			Description:         "If the system deployed in Cirrus",
			Optional:            true,
		},
		"deployment_details": schema.SingleNestedAttribute{
			MarkdownDescription: "Details of the deployment",
			Description:         "Details of the deployment",
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"system_on_prem": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"deployment_type": schema.StringAttribute{
							Description:         "System deployment types (e.g. PUBLIC_CLOUD)",
							MarkdownDescription: "System deployment types (e.g. PUBLIC_CLOUD)",
							Computed:            true,
						},
						"site_name": schema.StringAttribute{
							MarkdownDescription: "Name of the site where the system is located",
							Description:         "Name of the site where the system is located",
							Computed:            true,
						},
						"location": schema.StringAttribute{
							Description:         "User provided description of where the system is located.",
							MarkdownDescription: "User provided description of where the system is located.",
							Computed:            true,
						},
						"country": schema.StringAttribute{
							Description:         "Name of the country where the system is located.",
							MarkdownDescription: "Name of the country where the system is located.",
							Computed:            true,
						},
						"state": schema.StringAttribute{
							Description:         "Name of the state where the system is located.",
							MarkdownDescription: "Name of the state where the system is located.",
							Computed:            true,
						},
						"city": schema.StringAttribute{
							Description:         "Name of the city where the system is located.",
							MarkdownDescription: "Name of the city where the system is located.",
							Computed:            true,
						},
						"street_address_1": schema.StringAttribute{
							Description:         "Street address 1 of where the system is located",
							MarkdownDescription: "Street address 1 of where the system is located",
							Computed:            true,
						},
						"street_address_2": schema.StringAttribute{
							Description:         "Street address 2 of where the system is located",
							MarkdownDescription: "Street address 2 of where the system is located",
							Computed:            true,
						},
						"zip_code": schema.StringAttribute{
							Description:         "Zip code of where the system is located",
							MarkdownDescription: "Zip code of where the system is located",
							Computed:            true,
						},
					},
				},
				"system_public_cloud": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"deployment_type": schema.StringAttribute{
							Description:         "System deployment types (e.g. PUBLIC_CLOUD)",
							MarkdownDescription: "System deployment types (e.g. PUBLIC_CLOUD)",
							Computed:            true,
						},
						"cloud_type": schema.StringAttribute{
							Description:         "Enum for all the supported cloud providers * AWS - Amazon Web Services",
							MarkdownDescription: "Enum for all the supported cloud providers * AWS - Amazon Web Services",
							Computed:            true,
						},
						"cloud_account": schema.StringAttribute{
							Description:         "Cloud provider account where the storage system resides",
							MarkdownDescription: "Cloud provider account where the storage system resides",
							Computed:            true,
						},
						"cloud_region": schema.StringAttribute{
							Description:         "Cloud provider region where the storage system resides",
							MarkdownDescription: "Cloud provider region where the storage system resides",
							Computed:            true,
						},
						"availability_zone_topology": schema.StringAttribute{
							Description:         "This enum represents all the availability zone topology * SINGLE_AVAILABILITY_ZONE * MULTIPLE_AVAILABILITY_ZONE",
							MarkdownDescription: "This enum represents all the availability zone topology * SINGLE_AVAILABILITY_ZONE * MULTIPLE_AVAILABILITY_ZONE",
							Computed:            true,
						},
						"virtual_private_cloud": schema.StringAttribute{
							Description:         "Cloud virtual private environment identifier",
							MarkdownDescription: "Cloud virtual private environment identifier",
							Computed:            true,
						},
						"cloud_management_address": schema.StringAttribute{
							Description:         "Management IPv4 or IPv6 address or DNS name of the storage system in cloud",
							MarkdownDescription: "Management IPv4 or IPv6 address or DNS name of the storage system in cloud",
							Computed:            true,
						},
						"minimum_iops": schema.Int64Attribute{
							Description:         "Minimum IOPS requested during the deployment time - Unit: IO/s",
							MarkdownDescription: "Minimum IOPS requested during the deployment time - Unit: IO/s",
							Computed:            true,
						},
						"minimum_capacity": schema.Int64Attribute{
							Description:         "Minimum capacity requested during the deployment time - Unit: bytes",
							MarkdownDescription: "Minimum capacity requested during the deployment time - Unit: bytes",
							Computed:            true,
						},
						"tier_type": schema.StringAttribute{
							Description:         "Tier type requested during the deployment time",
							MarkdownDescription: "Tier type requested during the deployment time",
							Computed:            true,
						},
						"ssh_key_name": schema.StringAttribute{
							Computed: true,
						},
						"vpc": schema.SingleNestedAttribute{
							Description:         "VPC",
							MarkdownDescription: "VPC",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"is_new_vpc": schema.BoolAttribute{
									Computed: true,
								},
								"vpc_id": schema.StringAttribute{
									Computed: true,
								},
								"vpc_name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"subnet_options": schema.ListNestedAttribute{
							Description:         "Subnet options",
							MarkdownDescription: "Subnet options",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"subnet_option": schema.ListNestedAttribute{
										Computed: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"subnet_id": schema.StringAttribute{
													Computed: true,
												},
												"cidr_block": schema.StringAttribute{
													Computed: true,
												},
												"subnet_type": schema.StringAttribute{
													Computed: true,
												},
											},
										},
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

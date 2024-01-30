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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// BlockStorageDataSourceSchema defines the schema for a storage system instance data source and
// its attributes are used in storage systems collection data source
var BlockStorageDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Unique Host Claim ID",
			Optional:            true,
		},
		"system_id": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"system_type": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"bandwidth": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"capacity_impact": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"capacity_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"compression_savings": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"configuration_impact": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"configuration_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"configured_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"connectivity_status": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"license_type": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"license_expiration_date_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"contract_coverage_type": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"contract_expiration_date_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"data_protection_impact": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"data_protection_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"display_identifier": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"free_percent": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"free_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"health_connectivity_status": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"health_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"health_score": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"health_state": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"iops": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"ipv4_address": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"ipv6_address": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"last_contact_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"latency": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"logical_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"model": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"overall_efficiency": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"performance_impact": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"performance_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"serial_number": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"site_name": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"snaps_savings": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"system_health_impact": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"system_health_issue_count": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"thin_savings": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"total_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"unconfigured_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"used_percent": schema.Float64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"used_size": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"vendor": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"product_version": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"version": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"cirrus_deployed": schema.BoolAttribute{
			MarkdownDescription: " ",
			Optional:            true,
		},
		"deployment_details": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"system_on_prem": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"deployment_type": schema.StringAttribute{
							Computed: true,
						},
						"site_name": schema.StringAttribute{
							Computed: true,
						},
						"location": schema.StringAttribute{
							Computed: true,
						},
						"country": schema.StringAttribute{
							Computed: true,
						},
						"state": schema.StringAttribute{
							Computed: true,
						},
						"city": schema.StringAttribute{
							Computed: true,
						},
						"street_address_1": schema.StringAttribute{
							Computed: true,
						},
						"street_address_2": schema.StringAttribute{
							Computed: true,
						},
						"zip_code": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				"system_public_cloud": schema.SingleNestedAttribute{
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"deployment_type": schema.StringAttribute{
							Computed: true,
						},
						"cloud_type": schema.StringAttribute{
							Computed: true,
						},
						"cloud_account": schema.StringAttribute{
							Computed: true,
						},
						"cloud_region": schema.StringAttribute{
							Computed: true,
						},
						"availability_zone_topology": schema.StringAttribute{
							Computed: true,
						},
						"virtual_private_cloud": schema.StringAttribute{
							Computed: true,
						},
						"cloud_management_address": schema.StringAttribute{
							Computed: true,
						},
						"minimum_iops": schema.Int64Attribute{
							Computed: true,
						},
						"minimum_capacity": schema.Int64Attribute{
							Computed: true,
						},
						"tier_type": schema.StringAttribute{
							Computed: true,
						},
						"ssh_key_name": schema.StringAttribute{
							Computed: true,
						},
						"vpc": schema.SingleNestedAttribute{
							Computed: true,
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
							Computed: true,
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

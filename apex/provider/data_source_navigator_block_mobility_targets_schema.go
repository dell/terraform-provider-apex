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

// MobilityTargetsDataSourceSchema defines tha acceptable confirguation and state attribute names and types
var MobilityTargetsDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Unique Host Claim ID",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"system_id": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"system_type": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"source_mobility_group_id": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"image_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"target_members": schema.ListNestedAttribute{
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
		"last_copy_job_id": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"bandwidth_limit": schema.Int64Attribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"target_system_options": schema.StringAttribute{
			Computed: true,
		},
		"type": schema.StringAttribute{
			Optional: true,
		},
	},
}

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
			MarkdownDescription: "Idenifier of this target mobility group",
			Description:         "Idenifier of this target mobility group",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Description:         "Name of the mobility target",
			MarkdownDescription: "Name of the mobility target",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the mobility target",
			Description:         "Description of the mobility target",
			Computed:            true,
		},
		"system_id": schema.StringAttribute{
			Description:         "ID of the target system",
			MarkdownDescription: "ID of the target system",
			Computed:            true,
		},
		"system_type": schema.StringAttribute{
			Description:         "The source system type (e.g.: POWERFLEX)",
			MarkdownDescription: "The source system type (e.g.: POWERFLEX)",
			Computed:            true,
		},
		"source_mobility_group_id": schema.StringAttribute{
			Description:         "ID of the source mobility group",
			MarkdownDescription: "ID of the source mobility group",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			Description:         "Timestamp from when the group was created",
			MarkdownDescription: "Timestamp from when the group was created",
			Computed:            true,
		},
		"image_timestamp": schema.StringAttribute{
			Description:         "Timestamp of the last source image copied to this target",
			MarkdownDescription: "Timestamp of the last source image copied to this target",
			Computed:            true,
		},
		"target_members": schema.ListNestedAttribute{
			Description:         "A mobility member map is a mapping of a mobility member and it's related member. Ex: a target volume with a reference to the source volume or a clone volume and its related target volume.",
			MarkdownDescription: "A mobility member map is a mapping of a mobility member and it's related member. Ex: a target volume with a reference to the source volume or a clone volume and its related target volume.",
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
		"last_copy_job_id": schema.StringAttribute{
			MarkdownDescription: "Last copy job ID",
			Description:         "Last copy job ID",
			Computed:            true,
		},
		"bandwidth_limit": schema.Int64Attribute{
			Description:         "Bandwidth limit in Mbps (Mega bits per second).",
			MarkdownDescription: "Bandwidth limit in Mbps (Mega bits per second).",
			Computed:            true,
		},
		"target_system_options": schema.StringAttribute{
			Description:         "Storage pool id to use for allocating target volumes",
			MarkdownDescription: "Storage pool id to use for allocating target volumes",
			Computed:            true,
		},
		"type": schema.StringAttribute{
			Optional: true,
		},
	},
}

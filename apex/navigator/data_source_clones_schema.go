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

// ClonesDataSourceSchema defines tha acceptable confirguation and state attribute names and types
var ClonesDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name of the clone",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"mobility_target_id": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"refresh_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
		},
		"image_timestamp": schema.StringAttribute{
			MarkdownDescription: " ",
			Computed:            true,
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
						Computed: true,
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

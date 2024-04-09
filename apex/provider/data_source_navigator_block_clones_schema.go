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

// ClonesDataSourceSchema defines tha acceptable confirguation and state attribute names and types
var ClonesDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description:         "Identifier of the cloned object",
			MarkdownDescription: "Identifier of the cloned object",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Description:         "Name of the clone",
			MarkdownDescription: "Name of the clone",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Description:         "Description of the clone",
			MarkdownDescription: "Description of the clone",
			Computed:            true,
		},
		"mobility_target_id": schema.StringAttribute{
			Description:         "Mobility target ID",
			MarkdownDescription: "Mobility target ID",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			Description:         "When the clone was created",
			MarkdownDescription: "When the clone was created",
			Computed:            true,
		},
		"refresh_timestamp": schema.StringAttribute{
			Description:         "When the clone was last updated",
			MarkdownDescription: "When the clone was last updated",
			Computed:            true,
		},
		"image_timestamp": schema.StringAttribute{
			Description:         "Image timestamp",
			MarkdownDescription: "Image timestamp",
			Computed:            true,
		},
		"clone_volumes": schema.ListNestedAttribute{
			Description:         "A clone mobility member provides details of clone volume",
			MarkdownDescription: "A clone mobility member provides details of clone volume",
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
		"host_mappings": schema.ListNestedAttribute{
			Description:         "This represents the mapping of a repurposed (clone) storage object to a host (presumably using the clone for some analytical workload)",
			MarkdownDescription: "This represents the mapping of a repurposed (clone) storage object to a host (presumably using the clone for some analytical workload)",
			Computed:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"host_name": schema.StringAttribute{
						Description:         "Name of host/SDC to be mapped to the clone",
						MarkdownDescription: "Name of host/SDC to be mapped to the clone",
						Computed:            true,
					},
					"host_ip": schema.StringAttribute{
						Description:         "IP address of host",
						MarkdownDescription: "IP address of host",
						Computed:            true,
					},
					"host_id": schema.StringAttribute{
						Description:         "Identifier of the host",
						MarkdownDescription: "Identifier of the host",
						Computed:            true,
					},
					"id": schema.StringAttribute{
						Description:         "This is a host mappings id generated by APEX Navigator for Multicloud Storage",
						MarkdownDescription: "This is a host mappings id generated by APEX Navigator for Multicloud Storage",
						Computed:            true,
					},
					"nqn": schema.StringAttribute{
						Description:         "NVMe qualified name. Only applicable if host_initiator_protocol is NVMe",
						MarkdownDescription: "NVMe qualified name. Only applicable if host_initiator_protocol is NVMe",
						Computed:            true,
					},
					"host_initiator_protocol": schema.StringAttribute{
						Description:         "Type of the host",
						MarkdownDescription: "Type of the host",
						Computed:            true,
					},
				},
			},
		},
	},
}
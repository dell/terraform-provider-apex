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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MobilityGroupsDataSourceSchema defines the acceptable configuration and state attributes names and types
var MobilityGroupsDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description:         "Mobility group identifier",
			MarkdownDescription: "Mobility group identifier",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Description:         "Name of the mobility group",
			MarkdownDescription: "Name of the mobility group",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Description:         "Description of the mobility group",
			MarkdownDescription: "Description of the mobility group",
			Computed:            true,
		},
		"system_id": schema.StringAttribute{
			Description:         "Identifier of the system for the mobility group members",
			MarkdownDescription: "Identifier of the system for the mobility group members",
			Computed:            true,
		},
		"system_type": schema.StringAttribute{
			Description:         "The source system type (e.g.: POWERFLEX)",
			MarkdownDescription: "The source system type (e.g.: POWERFLEX)",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			Description:         "When the mobility group was created",
			MarkdownDescription: "When the mobility group was created",
			Computed:            true,
		},
		"volume_id": schema.ListAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"members": schema.ListNestedAttribute{
			Computed:    true,
			Description: "A mobility member is an object (e.g. volume) that is part of a mobility group that will be the source of mobility copy operations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:         "Identifier of the member (e.g. PowerFlex volume identifier)",
						MarkdownDescription: "Identifier of the member (e.g. PowerFlex volume identifier)",
						Computed:            true,
					},
					"name": schema.StringAttribute{
						Description:         "Name of the member (e.g. name of the volume)",
						MarkdownDescription: "Name of the member (e.g. name of the volume)",
						Computed:            true,
					},
					"size": schema.StringAttribute{
						Description:         "Size of the member (e.g. volume size in bytes)",
						MarkdownDescription: "Size of the member (e.g. volume size in bytes)",
						Computed:            true,
					},
				},
			},
		},
	},
}

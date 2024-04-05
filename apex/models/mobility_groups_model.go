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

package models

import (
	client "dell/apex-client"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// MobilityGroupModel defines the attribute names and types for a Mobility Target TF model
type MobilityGroupModel struct {
	ID                    types.String                  `tfsdk:"id"`
	Name                  types.String                  `tfsdk:"name"`
	Description           types.String                  `tfsdk:"description"`
	SystemID              types.String                  `tfsdk:"system_id"`
	SystemType            *client.StorageSystemTypeEnum `tfsdk:"system_type"`
	CreationTimeStamp     types.String                  `tfsdk:"creation_timestamp"`
	Members               basetypes.ListValue           `tfsdk:"members"`
	VolumeID              []types.String                `tfsdk:"volume_id"`
	ActivationClientModel *ActivationClientModel        `tfsdk:"powerflex"`
}

// MobilityGroupModelDs defines the attribute names and types for a Mobility Target TF model
type MobilityGroupModelDs struct {
	ID                types.String                  `tfsdk:"id"`
	Name              types.String                  `tfsdk:"name"`
	Description       types.String                  `tfsdk:"description"`
	SystemID          types.String                  `tfsdk:"system_id"`
	SystemType        *client.StorageSystemTypeEnum `tfsdk:"system_type"`
	CreationTimeStamp types.String                  `tfsdk:"creation_timestamp"`
	Members           basetypes.ListValue           `tfsdk:"members"`
	VolumeID          []types.String                `tfsdk:"volume_id"`
}

// MobilityGroupsDataSourceModel defines the attribute names and types for a Mobility Target TF model
type MobilityGroupsDataSourceModel struct {
	MobilityGroups []MobilityGroupModelDs   `tfsdk:"mobility_groups"`
	ID             types.String             `tfsdk:"id"`
	Filter         *MobilityGroupFilterType `tfsdk:"filter"`
}

// MobilityGroupFilterType describes the filter data model.
type MobilityGroupFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

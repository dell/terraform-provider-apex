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

// MobilityTargetModel defines the attribute names and types for a Mobility Target TF model
type MobilityTargetModel struct {
	ID                    types.String                  `tfsdk:"id"`
	Name                  types.String                  `tfsdk:"name"`
	Description           types.String                  `tfsdk:"description"`
	SystemID              types.String                  `tfsdk:"system_id"`
	SystemType            *client.StorageSystemTypeEnum `tfsdk:"system_type"`
	SourceMobilityGroupID types.String                  `tfsdk:"source_mobility_group_id"`
	CreationTimestamp     types.String                  `tfsdk:"creation_timestamp"`
	ImageTimestamp        types.String                  `tfsdk:"image_timestamp"`
	LastCopyJobID         types.String                  `tfsdk:"last_copy_job_id"`
	TargetMembers         basetypes.ListValue           `tfsdk:"target_members"`
	TargetSystemOptions   types.String                  `tfsdk:"target_system_options"`
	BandwidthLimit        types.Int64                   `tfsdk:"bandwidth_limit"`
	Type                  types.String                  `tfsdk:"type"`
	ActivationClientModel *ActivationClientModel        `tfsdk:"powerflex"`
}

// MobilityTargetModelDs defines the attribute names and types for a Mobility Target TF model
type MobilityTargetModelDs struct {
	ID                    types.String                  `tfsdk:"id"`
	Name                  types.String                  `tfsdk:"name"`
	Description           types.String                  `tfsdk:"description"`
	SystemID              types.String                  `tfsdk:"system_id"`
	SystemType            *client.StorageSystemTypeEnum `tfsdk:"system_type"`
	SourceMobilityGroupID types.String                  `tfsdk:"source_mobility_group_id"`
	CreationTimestamp     types.String                  `tfsdk:"creation_timestamp"`
	ImageTimestamp        types.String                  `tfsdk:"image_timestamp"`
	LastCopyJobID         types.String                  `tfsdk:"last_copy_job_id"`
	TargetMembers         basetypes.ListValue           `tfsdk:"target_members"`
	TargetSystemOptions   types.String                  `tfsdk:"target_system_options"`
	BandwidthLimit        types.Int64                   `tfsdk:"bandwidth_limit"`
	Type                  types.String                  `tfsdk:"type"`
}

// MobilityTargetsDataSourceModel defines the attribute names and types for a Mobility Targets Collection
type MobilityTargetsDataSourceModel struct {
	MobilityTargets []MobilityTargetModelDs   `tfsdk:"mobility_targets"`
	ID              types.String              `tfsdk:"id"`
	Filter          *MobilityTargetFilterType `tfsdk:"filter"`
}

// MobilityTargetFilterType describes the filter data model.
type MobilityTargetFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

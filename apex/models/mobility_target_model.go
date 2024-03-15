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
	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
}

// MobilityTargetsDataSourceModel defines the attribute names and types for a Mobility Targets Collection
type MobilityTargetsDataSourceModel struct {
	MobilityTargets []MobilityTargetModel `tfsdk:"mobility_targets"`
	ID              types.String          `tfsdk:"id"`
}

// GetMobilityTargetModel creates a new MobilityTargetModel based on the provided mobilityTarget
func GetMobilityTargetModel(mobilityTarget client.MobilityTarget) (model MobilityTargetModel) {
	model = MobilityTargetModel{
		ID:                    types.StringValue(mobilityTarget.Id),
		Name:                  types.StringValue(mobilityTarget.Name),
		Description:           types.StringValue(mobilityTarget.Description),
		SystemID:              types.StringValue(mobilityTarget.SystemId),
		SystemType:            (mobilityTarget.SystemType),
		SourceMobilityGroupID: types.StringValue(mobilityTarget.SourceMobilityGroupId),
		CreationTimestamp:     types.StringValue(mobilityTarget.CreationTimestamp),
		Type:                  types.StringValue(mobilityTarget.Type),
	}
	if mobilityTarget.ImageTimestamp != nil {
		model.ImageTimestamp = types.StringValue(*mobilityTarget.ImageTimestamp)
	}
	if mobilityTarget.LastCopyJobId != nil {
		model.LastCopyJobID = types.StringValue(*mobilityTarget.LastCopyJobId)
	}
	if mobilityTarget.BandwidthLimit != nil {
		model.BandwidthLimit = types.Int64Value(int64(*mobilityTarget.BandwidthLimit))
	}

	attrTypes := map[string]attr.Type{
		"id":        types.StringType,
		"parent_id": types.StringType,
		"name":      types.StringType,
		"size":      types.StringType,
	}
	var targetValues []attr.Value

	if mobilityTarget.TargetMembers != nil { //nolint:dupl
		for _, target := range mobilityTarget.TargetMembers {
			attrValues := map[string]attr.Value{
				"id":        types.StringValue(target.Id),
				"parent_id": types.StringValue(target.ParentId),
				"name":      types.StringValue(target.Name),
				"size":      types.StringValue(target.Size),
			}
			// TF Object representing a target member
			object, _ := types.ObjectValue(attrTypes, attrValues)
			targetValues = append(targetValues, object)
		}
		newObjectType := types.ObjectType{}
		newObjectType.AttrTypes = map[string]attr.Type{
			"id":        types.StringType,
			"parent_id": types.StringType,
			"name":      types.StringType,
			"size":      types.StringType,
		}

		model.TargetMembers = basetypes.NewListValueMust(newObjectType, targetValues)
	}
	return model
}
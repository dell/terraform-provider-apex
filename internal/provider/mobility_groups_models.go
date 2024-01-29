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
	client "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/pkg/gen/cirrusv1/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// MobilityGroupModel defines the attribute names and types for a Mobility Target TF model
type MobilityGroupModel struct {
	ID                types.String                  `tfsdk:"id"`
	Name              types.String                  `tfsdk:"name"`
	Description       types.String                  `tfsdk:"description"`
	SystemID          types.String                  `tfsdk:"system_id"`
	SystemType        *client.StorageSystemTypeEnum `tfsdk:"system_type"`
	CreationTimeStamp types.String                  `tfsdk:"creation_timestamp"`
	Members           basetypes.ListValue           `tfsdk:"members"`
	VolumeID          []types.String                `tfsdk:"volume_id"`
}

// GetMobilityGroupModel returns a new mobilityGroupModel based on data from the given mobilityGroup
func GetMobilityGroupModel(mobilityGroup client.MobilityGroup) (model MobilityGroupModel) {
	model = MobilityGroupModel{
		ID:                types.StringValue(mobilityGroup.Id),
		Name:              types.StringValue(mobilityGroup.Name),
		Description:       types.StringValue(*mobilityGroup.Description),
		SystemID:          types.StringValue(mobilityGroup.SystemId),
		SystemType:        (&mobilityGroup.SystemType),
		CreationTimeStamp: types.StringValue(mobilityGroup.CreationTimestamp),
	}
	attrTypes := map[string]attr.Type{
		"id":   types.StringType,
		"name": types.StringType,
		"size": types.StringType,
	}
	var groupValues []attr.Value

	if mobilityGroup.Members != nil { //nolint:dupl
		for _, member := range mobilityGroup.Members {
			attrValues := map[string]attr.Value{
				"id":   types.StringValue(member.Id),
				"name": types.StringValue(member.Name),
				"size": types.StringValue(member.Size),
			}
			// TF Object representing a group member
			object, _ := types.ObjectValue(attrTypes, attrValues)
			groupValues = append(groupValues, object)
		}
		newObjectType := types.ObjectType{}
		newObjectType.AttrTypes = map[string]attr.Type{
			"id":   types.StringType,
			"name": types.StringType,
			"size": types.StringType,
		}

		model.Members = basetypes.NewListValueMust(newObjectType, groupValues)
	}
	return model
}

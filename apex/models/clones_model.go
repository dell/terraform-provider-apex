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

// ClonesModel defines the attribute names and types for a Clones TF model
type ClonesModel struct {
	ID                types.String        `tfsdk:"id"`
	Name              types.String        `tfsdk:"name"`
	Description       types.String        `tfsdk:"description"`
	MobilityTargetID  types.String        `tfsdk:"mobility_target_id"`
	CreationTimestamp types.String        `tfsdk:"creation_timestamp"`
	RefreshTimestamp  types.String        `tfsdk:"refresh_timestamp"`
	ImageTimestamp    types.String        `tfsdk:"image_timestamp"`
	CloneVolumes      basetypes.ListValue `tfsdk:"clone_volumes"`
	HostMappings      basetypes.ListValue `tfsdk:"host_mappings"`
	// HostMappings      []cloneToHostMappingModel `tfsdk:"host_mappings"`
}

// ClonesDataSourceModel maps the data source schema data.
type ClonesDataSourceModel struct {
	Clones []ClonesModel `tfsdk:"clones"`
	ID     types.String  `tfsdk:"id"`
}

// ClonesRefreshModel defines the attribute names and types for a Clones Refresh TF model
type ClonesRefreshModel struct {
	ID      types.String `tfsdk:"id"`
	CloneID types.String `tfsdk:"clone_id"`
	Status  types.String `tfsdk:"status"`
}

// GetClonesModel returns a new clonesModel based on data from the given clones
func GetClonesModel(clone client.Clone) (model ClonesModel) {
	model = ClonesModel{
		ID:                types.StringValue(clone.Id),
		Name:              types.StringValue(clone.Name),
		Description:       types.StringValue(clone.Description),
		MobilityTargetID:  types.StringValue(*clone.MobilityTargetId),
		CreationTimestamp: types.StringValue(clone.CreationTimestamp),
		RefreshTimestamp:  types.StringValue(clone.RefreshTimestamp),
		ImageTimestamp:    types.StringValue(clone.ImageTimestamp),
	}
	attrTypes := map[string]attr.Type{
		"id":        types.StringType,
		"parent_id": types.StringType,
		"name":      types.StringType,
		"size":      types.StringType,
	}
	var volumeValues []attr.Value

	if clone.CloneVolumes != nil { //nolint:dupl
		for _, cloneVolume := range clone.CloneVolumes {
			values := map[string]attr.Value{
				"id":        types.StringValue(cloneVolume.Id),
				"parent_id": types.StringValue(cloneVolume.ParentId),
				"name":      types.StringValue(cloneVolume.Name),
				"size":      types.StringValue(cloneVolume.Size),
			}
			object, _ := types.ObjectValue(attrTypes, values)
			volumeValues = append(volumeValues, object)
		}
		newObjectType := types.ObjectType{}
		newObjectType.AttrTypes = map[string]attr.Type{
			"id":        types.StringType,
			"parent_id": types.StringType,
			"name":      types.StringType,
			"size":      types.StringType,
		}

		model.CloneVolumes = basetypes.NewListValueMust(newObjectType, volumeValues)
	}
	attrHostTypes := map[string]attr.Type{
		"host_name":               types.StringType,
		"host_ip":                 types.StringType,
		"host_id":                 types.StringType,
		"id":                      types.StringType,
		"nqn":                     types.StringType,
		"host_initiator_protocol": types.StringType,
	}
	var hostValues []attr.Value

	if clone.HostMappings != nil {
		for _, hostMapping := range clone.HostMappings {
			mapValues := map[string]attr.Value{
				"host_name":               types.StringValue(hostMapping.HostName),
				"host_ip":                 types.StringValue(hostMapping.HostIp),
				"host_id":                 types.StringValue(hostMapping.HostId),
				"id":                      types.StringValue(*hostMapping.Id),
				"nqn":                     types.StringValue(*hostMapping.Nqn),
				"host_initiator_protocol": types.StringValue(string(*hostMapping.HostInitiatorProtocol)),
			}
			mapObject, _ := types.ObjectValue(attrHostTypes, mapValues)
			hostValues = append(hostValues, mapObject)
		}
		newHostObjectType := types.ObjectType{}
		newHostObjectType.AttrTypes = map[string]attr.Type{
			"host_name":               types.StringType,
			"host_ip":                 types.StringType,
			"host_id":                 types.StringType,
			"id":                      types.StringType,
			"nqn":                     types.StringType,
			"host_initiator_protocol": types.StringType,
		}
		model.HostMappings = basetypes.NewListValueMust(newHostObjectType, hostValues)
	}

	return model
}

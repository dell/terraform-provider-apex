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

package helper

import (
	"context"
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GetCloneCollection is a collection of clones
func GetCloneCollection(client *client.APIClient, filter string) (*client.ClonesCollection200Response, *http.Response, error) {
	// Check for empty filter
	if filter == "" {
		return client.ClonesAPI.ClonesCollection(context.Background()).Limit(500).Execute()
	}
	return client.ClonesAPI.ClonesCollection(context.Background()).Limit(500).Filter(filter).Execute()
}

// GetCloneInstance is a clone
func GetCloneInstance(client *client.APIClient, id string) (*client.Clone, *http.Response, error) {
	return client.ClonesAPI.ClonesInstance(context.Background(), id).Execute()
}

// CreateClone sends the create clone request
func CreateClone(request client.ApiClonesCreateRequest, input client.CloneCreateInput) (*client.Job, *http.Response, error) {
	request = request.CloneCreateInput(input)
	return request.Async(true).Execute()
}

// UpdateClone sends the update clone request
func UpdateClone(request client.ApiClonesModifyRequest, input client.UpdateCloneInput) (*client.Job, *http.Response, error) {
	request = request.UpdateCloneInput(input)
	return request.Async(true).Execute()
}

// MapClones maps clones to a host
func MapClones(ctx context.Context, mapReq client.ApiClonesMapRequest, cloneID string, hosts []basetypes.StringValue) (*client.Job, *http.Response, error) {
	// Create Clones POST request
	hostIds := make([]string, 0, len(hosts))
	for _, mapping := range hosts {
		hostIds = append(hostIds, mapping.ValueString())
	}
	mapInput := *client.NewMapInput(hostIds)
	mapReq = mapReq.MapInput(mapInput)
	return mapReq.Async(true).Execute()
}

// RefreshClone refreshes a clone
func RefreshClone(refreshReq client.ApiClonesRefreshRequest) (*client.Job, *http.Response, error) {
	// Executing refresh request request
	return refreshReq.Async(true).Execute()
}

// UnmapClones unmaps clones from a host
func UnmapClones(unmapReq client.ApiClonesUnmapRequest, hosts []basetypes.StringValue) (*client.Job, *http.Response, error) {
	hostIds := make([]string, 0, len(hosts))
	for _, mapping := range hosts {
		hostIds = append(hostIds, mapping.ValueString())
	}
	unmapInput := *client.NewUnmapInput(hostIds)
	unmapReq = unmapReq.UnmapInput(unmapInput)
	return unmapReq.Async(true).Execute()
}

// GetClonesModel returns a new clonesModel based on data from the given clones
func GetClonesModel(clone client.Clone, systemID string) (model models.ClonesModel) {
	model = models.ClonesModel{
		ID:                types.StringValue(clone.Id),
		Name:              types.StringValue(clone.Name),
		Description:       types.StringValue(clone.Description),
		MobilityTargetID:  types.StringPointerValue(clone.MobilityTargetId),
		CreationTimestamp: types.StringValue(clone.CreationTimestamp),
		RefreshTimestamp:  types.StringValue(clone.RefreshTimestamp),
		ImageTimestamp:    types.StringValue(clone.ImageTimestamp),
		SystemID:          types.StringValue(systemID),
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
			hInitator := "unknown"
			if hostMapping.HostInitiatorProtocol != nil {
				hInitator = string(*hostMapping.HostInitiatorProtocol)
			}
			mapValues := map[string]attr.Value{
				"host_name":               types.StringValue(hostMapping.HostName),
				"host_ip":                 types.StringValue(hostMapping.HostIp),
				"host_id":                 types.StringValue(hostMapping.HostId),
				"id":                      types.StringPointerValue(hostMapping.Id),
				"nqn":                     types.StringPointerValue(hostMapping.Nqn),
				"host_initiator_protocol": types.StringValue(hInitator),
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

// GetClonesModelDs returns a new clonesModelDs based on data from the given clones
func GetClonesModelDs(clone client.Clone) (model models.ClonesModelDs) {
	model = models.ClonesModelDs{
		ID:                types.StringValue(clone.Id),
		Name:              types.StringValue(clone.Name),
		Description:       types.StringValue(clone.Description),
		MobilityTargetID:  types.StringPointerValue(clone.MobilityTargetId),
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
			hInitator := "unknown"
			if hostMapping.HostInitiatorProtocol != nil {
				hInitator = string(*hostMapping.HostInitiatorProtocol)
			}
			mapValues := map[string]attr.Value{
				"host_name":               types.StringValue(hostMapping.HostName),
				"host_ip":                 types.StringValue(hostMapping.HostIp),
				"host_id":                 types.StringValue(hostMapping.HostId),
				"id":                      types.StringPointerValue(hostMapping.Id),
				"nqn":                     types.StringPointerValue(hostMapping.Nqn),
				"host_initiator_protocol": types.StringValue(hInitator),
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

// UpdateHostMappings updates the host mappings after map/unmap operations
func UpdateHostMappings(clone *client.Clone) basetypes.ListValue {
	attrHostTypes := map[string]attr.Type{
		"host_name":               types.StringType,
		"host_ip":                 types.StringType,
		"host_id":                 types.StringType,
		"id":                      types.StringType,
		"nqn":                     types.StringType,
		"host_initiator_protocol": types.StringType,
	}
	var hostValues []attr.Value
	newHostObjectType := types.ObjectType{}
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
		newHostObjectType.AttrTypes = map[string]attr.Type{
			"host_name":               types.StringType,
			"host_ip":                 types.StringType,
			"host_id":                 types.StringType,
			"id":                      types.StringType,
			"nqn":                     types.StringType,
			"host_initiator_protocol": types.StringType,
		}

	}
	return basetypes.NewListValueMust(newHostObjectType, hostValues)
}

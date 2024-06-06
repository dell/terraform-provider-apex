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

// GetMobilityTargetCollection returns a list of all Mobility Targets
func GetMobilityTargetCollection(client *client.APIClient, filter string) (*client.MobilityTargetsCollection200Response, *http.Response, error) {
	// Check for empty filter
	if filter == "" {
		return client.MobilityTargetsAPI.MobilityTargetsCollection(context.Background()).Limit(500).Execute()
	}
	return client.MobilityTargetsAPI.MobilityTargetsCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// GetMobilityTarget returns a single Mobility Target
func GetMobilityTarget(client *client.APIClient, id string) (*client.MobilityTarget, *http.Response, error) {
	return client.MobilityTargetsAPI.MobilityTargetsInstance(context.Background(), id).Execute()
}

// CreateMobilityTarget sends the create mobility target request
func CreateMobilityTarget(request client.ApiMobilityTargetsCreateRequest, input client.CreateTargetInput) (*client.Job, *http.Response, error) {
	request = request.CreateTargetInput(input)
	return request.Async(true).Execute()
}

// UpdateMobilityTarget sends the update mobility target request
func UpdateMobilityTarget(request client.ApiMobilityTargetsModifyRequest, input client.UpdateMobilityTargetInput) (*client.Job, *http.Response, error) {
	request = request.UpdateMobilityTargetInput(input)
	return request.Async(true).Execute()
}

// GetMobilityTargetModel creates a new MobilityTargetModel based on the provided mobilityTarget
func GetMobilityTargetModel(mobilityTarget client.MobilityTarget) (model models.MobilityTargetModel) {
	model = models.MobilityTargetModel{
		ID:                    types.StringValue(mobilityTarget.Id),
		Name:                  types.StringValue(mobilityTarget.Name),
		Description:           types.StringValue(mobilityTarget.Description),
		SystemID:              types.StringValue(mobilityTarget.SystemId),
		SystemType:            (mobilityTarget.SystemType),
		SourceMobilityGroupID: types.StringValue(mobilityTarget.SourceMobilityGroupId),
		CreationTimestamp:     types.StringValue(mobilityTarget.CreationTimestamp),
		ImageTimestamp:        types.StringPointerValue(mobilityTarget.ImageTimestamp),
		LastCopyJobID:         types.StringPointerValue(mobilityTarget.LastCopyJobId),
	}

	// Do the null check specifically instead of using `types.Int64PointerValue` since we have to do the int64 conversion
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
				"name":      types.StringPointerValue(target.Name),
				"size":      types.StringPointerValue(target.Size),
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

// GetMobilityTargetModelDs creates a new MobilityTargetModel based on the provided mobilityTarget
func GetMobilityTargetModelDs(mobilityTarget client.MobilityTarget) (model models.MobilityTargetModelDs) {
	model = models.MobilityTargetModelDs{
		ID:                    types.StringValue(mobilityTarget.Id),
		Name:                  types.StringValue(mobilityTarget.Name),
		Description:           types.StringValue(mobilityTarget.Description),
		SystemID:              types.StringValue(mobilityTarget.SystemId),
		SystemType:            (mobilityTarget.SystemType),
		SourceMobilityGroupID: types.StringValue(mobilityTarget.SourceMobilityGroupId),
		CreationTimestamp:     types.StringValue(mobilityTarget.CreationTimestamp),
		ImageTimestamp:        types.StringPointerValue(mobilityTarget.ImageTimestamp),
		LastCopyJobID:         types.StringPointerValue(mobilityTarget.LastCopyJobId),
	}

	// Do the null check specifically instead of using `types.Int64PointerValue` since we have to do the int64 conversion
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
				"name":      types.StringPointerValue(target.Name),
				"size":      types.StringPointerValue(target.Size),
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

	} else { // Send an empty list

		attrValues := map[string]attr.Value{
			"id":        types.StringValue(""),
			"parent_id": types.StringValue(""),
			"name":      types.StringValue(""),
			"size":      types.StringValue(""),
		}
		object, _ := types.ObjectValue(attrTypes, attrValues)
		targetValues = append(targetValues, object)
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

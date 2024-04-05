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

// GetMobilityGroupCollection returns a list of all Mobility Groups
func GetMobilityGroupCollection(client *client.APIClient, filter string) (*client.MobilityGroupsCollection200Response, *http.Response, error) {
	return client.MobilityGroupsAPI.MobilityGroupsCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// GetMobilityGroup returns a single Mobility Group
func GetMobilityGroup(client *client.APIClient, id string) (*client.MobilityGroup, *http.Response, error) {
	return client.MobilityGroupsAPI.MobilityGroupsInstance(context.Background(), id).Execute()
}

// CreateMobilityGroup creates a new Mobility Group
func CreateMobilityGroup(ctx context.Context, clientAPI *client.APIClient, plan models.MobilityGroupModel) (*client.Job, *http.Response, error) {
	// Create Mobility Groups POST request
	var createReq client.ApiMobilityGroupsCreateRequest

	createReq = clientAPI.MobilityGroupsAPI.MobilityGroupsCreate(ctx)

	var members []string
	if plan.VolumeID != nil {
		for _, member := range plan.VolumeID {
			members = append(members, member.ValueString())
		}
	}

	mobilityGroupsInput := *client.NewSourceMobilityGroupInput(plan.Name.ValueString(), plan.SystemID.ValueString(), *plan.SystemType, members)
	mobilityGroupsInput.Description = plan.Description.ValueStringPointer()
	createReq = createReq.SourceMobilityGroupInput(mobilityGroupsInput)
	return createReq.Async(true).Execute()
}

// UpdateMobilityGroup updates an existing Mobility Group
func UpdateMobilityGroup(ctx context.Context, clientAPI *client.APIClient, plan models.MobilityGroupModel) (*client.Job, *http.Response, error) {
	// Create the update request
	req2 := clientAPI.MobilityGroupsAPI.MobilityGroupsModify(ctx, plan.ID.ValueString())

	updateInput := client.UpdateMobilityGroupInput{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
	}

	if plan.VolumeID != nil {
		for _, member := range plan.VolumeID {
			updateInput.Members = append(updateInput.Members, member.ValueString())
		}
	}

	req2 = req2.UpdateMobilityGroupInput(updateInput)
	return req2.Async(true).Execute()

}

// GetMobilityGroupModel returns a new mobilityGroupModel based on data from the given mobilityGroup
func GetMobilityGroupModel(mobilityGroup client.MobilityGroup) (model models.MobilityGroupModel) {
	model = models.MobilityGroupModel{
		ID:                types.StringValue(mobilityGroup.Id),
		Name:              types.StringValue(mobilityGroup.Name),
		Description:       types.StringPointerValue(mobilityGroup.Description),
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

// GetMobilityGroupModelDs returns a new mobilityGroupModel based on data from the given mobilityGroup
func GetMobilityGroupModelDs(mobilityGroup client.MobilityGroup) (model models.MobilityGroupModelDs) {
	model = models.MobilityGroupModelDs{
		ID:                types.StringValue(mobilityGroup.Id),
		Name:              types.StringValue(mobilityGroup.Name),
		Description:       types.StringPointerValue(mobilityGroup.Description),
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

// CopyMobilityGroups copies a Mobility Group
func CopyMobilityGroups(request client.ApiMobilityGroupsCopyRequest, input client.StartCopyInput) (*client.Job, *http.Response, error) {
	request = request.StartCopyInput(input)
	return request.Async(true).Execute()
}

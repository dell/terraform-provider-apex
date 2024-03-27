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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetVolumesCollection returns a list of all Volumes
func GetVolumesCollection(client *client.APIClient, filter string) (*client.VolumesCollection200Response, *http.Response, error) {
	return client.VolumesAPI.VolumesCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// MapVolumesToState maps the given array of client.VolumesInstance to an array of models.VolumesModel.
func MapVolumesToState(volumes []client.VolumesInstance) []models.VolumesModel {
	volumesArray := make([]models.VolumesModel, 0)
	for _, volume := range volumes {
		volumeState := models.VolumesModel{
			ID:                      types.StringValue(volume.Id),
			SystemID:                types.StringPointerValue(volume.SystemId),
			SystemType:              types.StringPointerValue(volume.SystemType),
			AllocatedSize:           types.Int64PointerValue(volume.AllocatedSize),
			Bandwidth:               types.Int64PointerValue(volume.Bandwidth),
			ConsistencyGroupName:    types.StringPointerValue(volume.ConsistencyGroupName),
			DataReductionPercent:    types.Float64PointerValue(volume.DataReductionPercent),
			DataReductionRatio:      types.Float64PointerValue(volume.DataReductionRatio),
			DataReductionSavedSize:  types.Int64PointerValue(volume.DataReductionSavedSize),
			IoLimitPolicyName:       types.StringPointerValue(volume.IoLimitPolicyName),
			Iops:                    types.Int64PointerValue(volume.Iops),
			IsCompressedOrDeduped:   types.StringPointerValue(volume.IsCompressedOrDeduped),
			IsThinEnabled:           types.BoolPointerValue(volume.IsThinEnabled),
			IssueCount:              types.Int64PointerValue(volume.IssueCount),
			Latency:                 types.Int64PointerValue(volume.Latency),
			LogicalSize:             types.Int64PointerValue(volume.LogicalSize),
			Name:                    types.StringPointerValue(volume.Name),
			NativeID:                types.StringPointerValue(volume.NativeId),
			Type:                    types.StringPointerValue(volume.Type),
			PoolID:                  types.StringPointerValue(volume.PoolId),
			PoolName:                types.StringPointerValue(volume.PoolName),
			PoolType:                types.StringPointerValue(volume.PoolType),
			SnapshotCount:           types.Int64PointerValue(volume.SnapshotCount),
			SnapshotPolicy:          types.StringPointerValue(volume.SnapshotPolicy),
			SnapshotSize:            types.Int64PointerValue(volume.SnapshotSize),
			StorageResourceID:       types.StringPointerValue(volume.StorageResourceId),
			StorageResourceNativeID: types.StringPointerValue(volume.StorageResourceNativeId),
			SystemModel:             types.StringPointerValue(volume.SystemModel),
			SystemName:              types.StringPointerValue(volume.SystemName),
			TotalSize:               types.Int64PointerValue(volume.TotalSize),
			UsedSize:                types.Int64PointerValue(volume.UsedSize),
			UsedSizeUnique:          types.Int64PointerValue(volume.UsedSizeUnique),
		}
		volumesArray = append(volumesArray, volumeState)
	}
	return volumesArray
}

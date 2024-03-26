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

// GetSourcePoolsCollection returns a list of all Source Pools
func GetSourcePoolsCollection(client *client.APIClient, filter string) (*client.PoolsCollection200Response, *http.Response, error) {
	return client.PoolsAPI.PoolsCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// MapPoolsToState maps the pools to the terraform state
func MapPoolsToState(pools []client.PoolsInstance) []models.PoolsModel {
	poolsArray := make([]models.PoolsModel, 0)
	for _, pool := range pools {
		poolState := models.PoolsModel{
			ID:                   types.StringValue(pool.Id),
			SystemID:             types.StringPointerValue(pool.SystemId),
			SystemType:           types.StringPointerValue(pool.SystemType),
			FreeSize:             types.Int64PointerValue(pool.FreeSize),
			IssueCount:           types.Int64PointerValue(pool.IssueCount),
			Name:                 types.StringPointerValue(pool.Name),
			NativeID:             types.StringPointerValue(pool.NativeId),
			SubscribedPercent:    types.Float64PointerValue(pool.SubscribedPercent),
			SubscribedSize:       types.Int64PointerValue(pool.SubscribedSize),
			SystemModel:          types.StringPointerValue(pool.SystemModel),
			SystemName:           types.StringPointerValue(pool.SystemName),
			TimeToFullPrediction: types.StringPointerValue(pool.TimeToFullPrediction),
			TotalSize:            types.Int64PointerValue(pool.TotalSize),
			Type:                 types.StringPointerValue(pool.Type),
			UsedPercent:          types.Float64PointerValue(pool.UsedPercent),
			UsedSize:             types.Int64PointerValue(pool.UsedSize),
		}
		poolsArray = append(poolsArray, poolState)
	}
	return poolsArray
}

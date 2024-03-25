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

	"github.com/dell/terraform-provider-apex/apex/models"
	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetStorageProductsCollection returns a list of all Stroage Products
func GetStorageProductsCollection(client *client.APIClient, filter string) (*client.StorageProductsCollection200Response, *http.Response, error) {
	return client.StorageProductsAPI.StorageProductsCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// MapStorageProduct maps storage product to terraformsdk model
func MapStorageProduct(storageProducts *client.StorageProductsCollection200Response) models.StorageProductsDataSourceModel {
	var state models.StorageProductsDataSourceModel
	// Map response body to model
	for _, storageProduct := range storageProducts.Results {
		storageProductState := models.StorageProductsModel{
			ID:            types.StringValue(storageProduct.Id),
			Name:          types.StringValue(storageProduct.Name),
			Description:   types.StringValue(storageProduct.Description),
			LatestVersion: types.StringValue(storageProduct.LatestVersion),
			SystemType:    client.StorageSystemTypeEnum(storageProduct.SystemType),
			StorageType:   storageProduct.StorageType,
			CloudType:     storageProduct.CloudType,
		}

		for _, supportmap := range storageProduct.SupportMap {
			storageProductState.SupportMaps = append(storageProductState.SupportMaps, models.SupportMapModel{
				ID:                        types.StringValue(supportmap.Id),
				SupportedEvaluationPeriod: types.Int64Value(int64(supportmap.SupportedEvaluationPeriod)),
				Version:                   types.StringValue(supportmap.Version),
				SupportedActions:          []client.StorageProductActionEnum{},
			})
		}

		state.StorageProducts = append(state.StorageProducts, storageProductState)
	}
	return state
}

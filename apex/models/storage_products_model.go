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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StorageProductsDataSourceModel maps the data source schema data.
type StorageProductsDataSourceModel struct {
	StorageProducts []StorageProductsModel     `tfsdk:"storage_products"`
	ID              types.String               `tfsdk:"id"`
	Filter          *StorageProductsFilterType `tfsdk:"filter"`
}

// StorageProductsModel maps hosts schema data.
type StorageProductsModel struct {
	ID            types.String                 `tfsdk:"id"`
	Name          types.String                 `tfsdk:"name"`
	SystemType    client.StorageSystemTypeEnum `tfsdk:"system_type"`
	StorageType   client.StorageTypeEnum       `tfsdk:"storage_type"`
	Description   types.String                 `tfsdk:"description"`
	CloudType     client.CloudProviderEnum     `tfsdk:"cloud_type"`
	LatestVersion types.String                 `tfsdk:"latest_version"`
	SupportMaps   []SupportMapModel            `tfsdk:"support_map"`
}

// SupportMapModel maps supportMap schema data
type SupportMapModel struct {
	ID                        types.String                      `tfsdk:"id"`
	SupportedEvaluationPeriod types.Int64                       `tfsdk:"supported_evaluation_period"`
	Version                   types.String                      `tfsdk:"version"`
	SupportedActions          []client.StorageProductActionEnum `tfsdk:"supported_actions"`
}

// StorageProductsFilterType describes the filter data model.
type StorageProductsFilterType struct {
	SystemType types.String `tfsdk:"system_type"`
}

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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PoolsDataSourceModel maps the data source schema data.
type PoolsDataSourceModel struct {
	Pools  []PoolsModel     `tfsdk:"pools"`
	ID     types.String     `tfsdk:"id"`
	Filter *PoolsFilterType `tfsdk:"filter"`
}

// PoolsFilterType describes the filter data model.
type PoolsFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

// PoolsModel maps pools schema data.
type PoolsModel struct {
	ID                   types.String  `tfsdk:"id"`
	SystemID             types.String  `tfsdk:"system_id"`
	SystemType           types.String  `tfsdk:"system_type"`
	FreeSize             types.Int64   `tfsdk:"free_size"`
	IssueCount           types.Int64   `tfsdk:"issue_count"`
	Name                 types.String  `tfsdk:"name"`
	NativeID             types.String  `tfsdk:"native_id"`
	SubscribedPercent    types.Float64 `tfsdk:"subscribed_percent"`
	SubscribedSize       types.Int64   `tfsdk:"subscribed_size"`
	SystemModel          types.String  `tfsdk:"system_model"`
	SystemName           types.String  `tfsdk:"system_name"`
	TimeToFullPrediction types.String  `tfsdk:"time_to_full_prediction"`
	TotalSize            types.Int64   `tfsdk:"total_size"`
	Type                 types.String  `tfsdk:"type"`
	UsedPercent          types.Float64 `tfsdk:"used_percent"`
	UsedSize             types.Int64   `tfsdk:"used_size"`
}

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

import "github.com/hashicorp/terraform-plugin-framework/types"

// VolumesModel maps volumes schema data.
type VolumesModel struct {
	ID                      types.String  `tfsdk:"id"`
	SystemID                types.String  `tfsdk:"system_id"`
	SystemType              types.String  `tfsdk:"system_type"`
	AllocatedSize           types.Int64   `tfsdk:"allocated_size"`
	Bandwidth               types.Int64   `tfsdk:"bandwidth"`
	ConsistencyGroupName    types.String  `tfsdk:"consistency_group_name"`
	DataReductionPercent    types.Float64 `tfsdk:"data_reduction_percent"`
	DataReductionRatio      types.Float64 `tfsdk:"data_reduction_ratio"`
	DataReductionSavedSize  types.Int64   `tfsdk:"data_reduction_saved_size"`
	IoLimitPolicyName       types.String  `tfsdk:"io_limit_policy_name"`
	Iops                    types.Int64   `tfsdk:"iops"`
	IsCompressedOrDeduped   types.String  `tfsdk:"is_compressed_or_deduped"`
	IsThinEnabled           types.Bool    `tfsdk:"is_thin_enabled"`
	IssueCount              types.Int64   `tfsdk:"issue_count"`
	Latency                 types.Int64   `tfsdk:"latency"`
	LogicalSize             types.Int64   `tfsdk:"logical_size"`
	Name                    types.String  `tfsdk:"name"`
	NativeID                types.String  `tfsdk:"native_id"`
	Type                    types.String  `tfsdk:"type"`
	PoolID                  types.String  `tfsdk:"pool_id"`
	PoolName                types.String  `tfsdk:"pool_name"`
	PoolType                types.String  `tfsdk:"pool_type"`
	SnapshotCount           types.Int64   `tfsdk:"snap_shot_count"`
	SnapshotPolicy          types.String  `tfsdk:"snap_shot_policy"`
	SnapshotSize            types.Int64   `tfsdk:"snap_shot_size"`
	StorageResourceID       types.String  `tfsdk:"storage_resource_id"`
	StorageResourceNativeID types.String  `tfsdk:"storage_resource_native_id"`
	SystemModel             types.String  `tfsdk:"system_model"`
	SystemName              types.String  `tfsdk:"system_name"`
	TotalSize               types.Int64   `tfsdk:"total_size"`
	UsedSize                types.Int64   `tfsdk:"used_size"`
	UsedSizeUnique          types.Int64   `tfsdk:"used_size_unique"`
}

// VolumesDataSourceModel maps the data source schema data.
type VolumesDataSourceModel struct {
	Volumes []VolumesModel `tfsdk:"volumes"`
	ID      types.String   `tfsdk:"id"`
}

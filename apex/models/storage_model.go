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
	client "dell/apex-client"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DeploymentDetailsModel maps storage system schema data.
type DeploymentDetailsModel struct {
	SystemOnPrem      *SystemOnPremDeploymentDetailsModel      `tfsdk:"system_on_prem"`
	SystemPublicCloud *SystemPublicCloudDeploymentDetailsModel `tfsdk:"system_public_cloud"`
}

// StoragesDataSourceModel maps storage system schema data.
type StoragesDataSourceModel struct {
	Storages []StorageModelDs   `tfsdk:"storages"`
	ID       types.String       `tfsdk:"id"`
	Filter   *StorageFilterType `tfsdk:"filter"`
}

// SystemOnPremDeploymentDetailsModel maps storage system schema data.
type SystemOnPremDeploymentDetailsModel struct {
	DeploymentType *client.SystemDeploymentTypeEnum `tfsdk:"deployment_type"`
	SiteName       types.String                     `tfsdk:"site_name"`
	Location       types.String                     `tfsdk:"location"`
	Country        types.String                     `tfsdk:"country"`
	State          types.String                     `tfsdk:"state"`
	City           types.String                     `tfsdk:"city"`
	StreetAddress1 types.String                     `tfsdk:"street_address_1"`
	StreetAddress2 types.String                     `tfsdk:"street_address_2"`
	ZipCode        types.String                     `tfsdk:"zip_code"`
}

// StorageFilterType describes the filter data model.
type StorageFilterType struct {
	IDs        []types.String `tfsdk:"ids"`
	SystemType types.String   `tfsdk:"system_type"`
}

// SystemPublicCloudDeploymentDetailsModel maps storage system schema data.
type SystemPublicCloudDeploymentDetailsModel struct {
	DeploymentType           *client.SystemDeploymentTypeEnum     `tfsdk:"deployment_type"`
	CloudType                *client.CloudProviderEnum            `tfsdk:"cloud_type"`
	CloudAccount             types.String                         `tfsdk:"cloud_account"`
	CloudRegion              types.String                         `tfsdk:"cloud_region"`
	AvailabilityZoneTopology *client.AvailabilityZoneTopologyEnum `tfsdk:"availability_zone_topology"`
	VirtualPrivateCloud      types.String                         `tfsdk:"virtual_private_cloud"`
	CloudManagementAddress   types.String                         `tfsdk:"cloud_management_address"`
	MinimumIops              types.Int64                          `tfsdk:"minimum_iops"`
	MinimumCapacity          types.Int64                          `tfsdk:"minimum_capacity"`
	RawCapacity              types.String                         `tfsdk:"raw_capacity"`
	TierType                 types.String                         `tfsdk:"tier_type"`
	SSHKeyName               types.String                         `tfsdk:"ssh_key_name"`
	IAMInstanceProfile       types.String                         `tfsdk:"iam_instance_profile"`
	Vpc                      *vpcModel                            `tfsdk:"vpc"`
	SubnetOptions            []SubnetOptionModel                  `tfsdk:"subnet_options"`
}

// SubnetOptionModel maps subnet schema data.
type SubnetOptionModel struct {
	SubnetID   types.String           `tfsdk:"subnet_id"`
	CidrBlock  types.String           `tfsdk:"cidr_block"`
	SubnetType *client.SubnetTypeEnum `tfsdk:"subnet_type"`
}

type vpcModel struct {
	IsNewVpc types.Bool   `tfsdk:"is_new_vpc"`
	VpcID    types.String `tfsdk:"vpc_id"`
	VpcName  types.String `tfsdk:"vpc_name"`
}

// StorageModel maps storage system schema data.
type StorageModel struct {
	ID                              types.String            `tfsdk:"id"`
	SystemID                        types.String            `tfsdk:"system_id"`
	StorageSystemType               types.String            `tfsdk:"storage_system_type"`
	SystemType                      types.String            `tfsdk:"system_type"`
	Bandwidth                       types.Int64             `tfsdk:"bandwidth"`
	CapacityImpact                  types.Int64             `tfsdk:"capacity_impact"`
	CapacityIssueCount              types.Int64             `tfsdk:"capacity_issue_count"`
	CompressionSavings              types.Float64           `tfsdk:"compression_savings"`
	ConfigurationImpact             types.Int64             `tfsdk:"configuration_impact"`
	ConfigurationIssueCount         types.Int64             `tfsdk:"configuration_issue_count"`
	ConfiguredSize                  types.Int64             `tfsdk:"configured_size"`
	ConnectivityStatus              types.String            `tfsdk:"connectivity_status"`
	LicenseType                     types.String            `tfsdk:"license_type"`
	LicenseExpirationDateTimestamp  types.String            `tfsdk:"license_expiration_date_timestamp"`
	ContractCoverageType            types.String            `tfsdk:"contract_coverage_type"`
	ContractExpirationDateTimestamp types.String            `tfsdk:"contract_expiration_date_timestamp"`
	DataProtectionImpact            types.Int64             `tfsdk:"data_protection_impact"`
	DataProtectionIssueCount        types.Int64             `tfsdk:"data_protection_issue_count"`
	DisplayIdentifier               types.String            `tfsdk:"display_identifier"`
	FreePercent                     types.Float64           `tfsdk:"free_percent"`
	FreeSize                        types.Int64             `tfsdk:"free_size"`
	HealthConnectivityStatus        types.String            `tfsdk:"health_connectivity_status"`
	HealthIssueCount                types.Int64             `tfsdk:"health_issue_count"`
	HealthScore                     types.Int64             `tfsdk:"health_score"`
	HealthState                     types.String            `tfsdk:"health_state"`
	Iops                            types.Int64             `tfsdk:"iops"`
	Ipv4Address                     types.String            `tfsdk:"ipv4_address"`
	Ipv6Address                     types.String            `tfsdk:"ipv6_address"`
	LastContactTimestamp            types.String            `tfsdk:"last_contact_timestamp"`
	Latency                         types.Int64             `tfsdk:"latency"`
	LogicalSize                     types.Int64             `tfsdk:"logical_size"`
	Model                           types.String            `tfsdk:"model"`
	Name                            types.String            `tfsdk:"name"`
	OverallEfficiency               types.Float64           `tfsdk:"overall_efficiency"`
	PerformanceImpact               types.Int64             `tfsdk:"performance_impact"`
	PerformanceIssueCount           types.Int64             `tfsdk:"performance_issue_count"`
	SerialNumber                    types.String            `tfsdk:"serial_number"`
	SiteName                        types.String            `tfsdk:"site_name"`
	SnapsSavings                    types.Float64           `tfsdk:"snaps_savings"`
	SystemHealthImpact              types.Int64             `tfsdk:"system_health_impact"`
	SystemHealthIssueCount          types.Int64             `tfsdk:"system_health_issue_count"`
	ThinSavings                     types.Float64           `tfsdk:"thin_savings"`
	TotalSize                       types.Int64             `tfsdk:"total_size"`
	UnconfiguredSize                types.Int64             `tfsdk:"unconfigured_size"`
	UsedPercent                     types.Float64           `tfsdk:"used_percent"`
	UsedSize                        types.Int64             `tfsdk:"used_size"`
	Vendor                          types.String            `tfsdk:"vendor"`
	ProductVersion                  types.String            `tfsdk:"product_version"`
	Version                         types.String            `tfsdk:"version"`
	DeploymentDetails               *DeploymentDetailsModel `tfsdk:"deployment_details"`
	ActivationClientModel           *ActivationClientModel  `tfsdk:"powerflex"`
}

// StorageModelDs maps storage system schema data.
type StorageModelDs struct {
	ID                              types.String            `tfsdk:"id"`
	SystemID                        types.String            `tfsdk:"system_id"`
	StorageSystemType               types.String            `tfsdk:"storage_system_type"`
	SystemType                      types.String            `tfsdk:"system_type"`
	Bandwidth                       types.Int64             `tfsdk:"bandwidth"`
	CapacityImpact                  types.Int64             `tfsdk:"capacity_impact"`
	CapacityIssueCount              types.Int64             `tfsdk:"capacity_issue_count"`
	CompressionSavings              types.Float64           `tfsdk:"compression_savings"`
	ConfigurationImpact             types.Int64             `tfsdk:"configuration_impact"`
	ConfigurationIssueCount         types.Int64             `tfsdk:"configuration_issue_count"`
	ConfiguredSize                  types.Int64             `tfsdk:"configured_size"`
	ConnectivityStatus              types.String            `tfsdk:"connectivity_status"`
	LicenseType                     types.String            `tfsdk:"license_type"`
	LicenseExpirationDateTimestamp  types.String            `tfsdk:"license_expiration_date_timestamp"`
	ContractCoverageType            types.String            `tfsdk:"contract_coverage_type"`
	ContractExpirationDateTimestamp types.String            `tfsdk:"contract_expiration_date_timestamp"`
	DataProtectionImpact            types.Int64             `tfsdk:"data_protection_impact"`
	DataProtectionIssueCount        types.Int64             `tfsdk:"data_protection_issue_count"`
	DisplayIdentifier               types.String            `tfsdk:"display_identifier"`
	FreePercent                     types.Float64           `tfsdk:"free_percent"`
	FreeSize                        types.Int64             `tfsdk:"free_size"`
	HealthConnectivityStatus        types.String            `tfsdk:"health_connectivity_status"`
	HealthIssueCount                types.Int64             `tfsdk:"health_issue_count"`
	HealthScore                     types.Int64             `tfsdk:"health_score"`
	HealthState                     types.String            `tfsdk:"health_state"`
	Iops                            types.Int64             `tfsdk:"iops"`
	Ipv4Address                     types.String            `tfsdk:"ipv4_address"`
	Ipv6Address                     types.String            `tfsdk:"ipv6_address"`
	LastContactTimestamp            types.String            `tfsdk:"last_contact_timestamp"`
	Latency                         types.Int64             `tfsdk:"latency"`
	LogicalSize                     types.Int64             `tfsdk:"logical_size"`
	Model                           types.String            `tfsdk:"model"`
	Name                            types.String            `tfsdk:"name"`
	OverallEfficiency               types.Float64           `tfsdk:"overall_efficiency"`
	PerformanceImpact               types.Int64             `tfsdk:"performance_impact"`
	PerformanceIssueCount           types.Int64             `tfsdk:"performance_issue_count"`
	SerialNumber                    types.String            `tfsdk:"serial_number"`
	SiteName                        types.String            `tfsdk:"site_name"`
	SnapsSavings                    types.Float64           `tfsdk:"snaps_savings"`
	SystemHealthImpact              types.Int64             `tfsdk:"system_health_impact"`
	SystemHealthIssueCount          types.Int64             `tfsdk:"system_health_issue_count"`
	ThinSavings                     types.Float64           `tfsdk:"thin_savings"`
	TotalSize                       types.Int64             `tfsdk:"total_size"`
	UnconfiguredSize                types.Int64             `tfsdk:"unconfigured_size"`
	UsedPercent                     types.Float64           `tfsdk:"used_percent"`
	UsedSize                        types.Int64             `tfsdk:"used_size"`
	Vendor                          types.String            `tfsdk:"vendor"`
	ProductVersion                  types.String            `tfsdk:"product_version"`
	Version                         types.String            `tfsdk:"version"`
	DeploymentDetails               *DeploymentDetailsModel `tfsdk:"deployment_details"`
}

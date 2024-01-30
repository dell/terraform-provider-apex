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

package navigator

import (
	"fmt"

	client "github.com/dell/terraform-provider-apex/client/apex"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type deploymentDetailsModel struct {
	SystemOnPrem      *systemOnPremDeploymentDetailsModel      `tfsdk:"system_on_prem"`
	SystemPublicCloud *systemPublicCloudDeploymentDetailsModel `tfsdk:"system_public_cloud"`
}
type systemOnPremDeploymentDetailsModel struct {
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
type systemPublicCloudDeploymentDetailsModel struct {
	DeploymentType           *client.SystemDeploymentTypeEnum     `tfsdk:"deployment_type"`
	CloudType                *client.CloudProviderEnum            `tfsdk:"cloud_type"`
	CloudAccount             types.String                         `tfsdk:"cloud_account"`
	CloudRegion              types.String                         `tfsdk:"cloud_region"`
	AvailabilityZoneTopology *client.AvailabilityZoneTopologyEnum `tfsdk:"availability_zone_topology"`
	VirtualPrivateCloud      types.String                         `tfsdk:"virtual_private_cloud"`
	CloudManagementAddress   types.String                         `tfsdk:"cloud_management_address"`
	MinimumIops              types.Int64                          `tfsdk:"minimum_iops"`
	MinimumCapacity          types.Int64                          `tfsdk:"minimum_capacity"`
	TierType                 types.String                         `tfsdk:"tier_type"`
	SSHKeyName               types.String                         `tfsdk:"ssh_key_name"`
	Vpc                      *vpcModel                            `tfsdk:"vpc"`
	SubnetOptions            []subnetOptionModel                  `tfsdk:"subnet_options"`
}

type subnetOptionModel struct {
	SubnetID   types.String           `tfsdk:"subnet_id"`
	CidrBlock  types.String           `tfsdk:"cidr_block"`
	SubnetType *client.SubnetTypeEnum `tfsdk:"subnet_type"`
}

type vpcModel struct {
	IsNewVpc types.Bool   `tfsdk:"is_new_vpc"`
	VpcID    types.String `tfsdk:"vpc_id"`
	VpcName  types.String `tfsdk:"vpc_name"`
}

// BlockStorageModel maps storage system schema data.
type BlockStorageModel struct {
	ID                              types.String            `tfsdk:"id"`
	SystemID                        types.String            `tfsdk:"system_id"`
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
	DeploymentDetails               *deploymentDetailsModel `tfsdk:"deployment_details"`
	CirrusDeployed                  types.Bool              `tfsdk:"cirrus_deployed"`
}

// GetBlockStorageSystem returns a storage systems model based on the given storage system instance
func GetBlockStorageSystem(storageSystem client.StorageSystemsInstance) (blockStorageState BlockStorageModel) {
	blockStorageState = BlockStorageModel{
		ID:                              types.StringValue(storageSystem.Id),
		SystemID:                        types.StringPointerValue(storageSystem.SystemId),
		SystemType:                      types.StringPointerValue(storageSystem.SystemType),
		Bandwidth:                       types.Int64PointerValue(storageSystem.Bandwidth),
		CapacityImpact:                  types.Int64PointerValue(storageSystem.CapacityImpact),
		CapacityIssueCount:              types.Int64PointerValue(storageSystem.CapacityIssueCount),
		CompressionSavings:              types.Float64PointerValue(storageSystem.CompressionSavings),
		ConfigurationImpact:             types.Int64PointerValue(storageSystem.ConfigurationImpact),
		ConfigurationIssueCount:         types.Int64PointerValue(storageSystem.ConfigurationIssueCount),
		ConfiguredSize:                  types.Int64PointerValue(storageSystem.ConfiguredSize),
		ConnectivityStatus:              types.StringPointerValue(storageSystem.ConnectivityStatus),
		LicenseType:                     types.StringPointerValue(storageSystem.LicenseType),
		LicenseExpirationDateTimestamp:  types.StringValue(convertTimeToString(storageSystem.LicenseExpirationDateTimestamp)),
		ContractCoverageType:            types.StringPointerValue(storageSystem.ContractCoverageType),
		ContractExpirationDateTimestamp: types.StringValue(convertTimeToString(storageSystem.ContractExpirationDateTimestamp)),
		DataProtectionImpact:            types.Int64PointerValue(storageSystem.DataProtectionImpact),
		DataProtectionIssueCount:        types.Int64PointerValue(storageSystem.DataProtectionIssueCount),
		DisplayIdentifier:               types.StringPointerValue(storageSystem.DisplayIdentifier),
		FreePercent:                     types.Float64PointerValue(storageSystem.FreePercent),
		FreeSize:                        types.Int64PointerValue(storageSystem.FreeSize),
		HealthConnectivityStatus:        types.StringPointerValue(storageSystem.HealthConnectivityStatus),
		HealthIssueCount:                types.Int64PointerValue(storageSystem.HealthIssueCount),
		HealthScore:                     types.Int64PointerValue(storageSystem.HealthScore),
		HealthState:                     types.StringPointerValue(storageSystem.HealthState),
		Iops:                            types.Int64PointerValue(storageSystem.Iops),
		Ipv4Address:                     types.StringPointerValue(storageSystem.Ipv4Address),
		Ipv6Address:                     types.StringPointerValue(storageSystem.Ipv6Address),
		LastContactTimestamp:            types.StringValue(convertTimeToString(storageSystem.LastContactTimestamp)),
		Latency:                         types.Int64PointerValue(storageSystem.Latency),
		LogicalSize:                     types.Int64PointerValue(storageSystem.LogicalSize),
		Model:                           types.StringPointerValue(storageSystem.Model),
		Name:                            types.StringPointerValue(storageSystem.Name),
		OverallEfficiency:               types.Float64PointerValue(storageSystem.OverallEfficiency),
		PerformanceImpact:               types.Int64PointerValue(storageSystem.PerformanceImpact),
		PerformanceIssueCount:           types.Int64PointerValue(storageSystem.PerformanceIssueCount),
		SerialNumber:                    types.StringPointerValue(storageSystem.SerialNumber),
		SiteName:                        types.StringPointerValue(storageSystem.SiteName),
		SnapsSavings:                    types.Float64PointerValue(storageSystem.SnapsSavings),
		SystemHealthImpact:              types.Int64PointerValue(storageSystem.SystemHealthImpact),
		SystemHealthIssueCount:          types.Int64PointerValue(storageSystem.SystemHealthIssueCount),
		ThinSavings:                     types.Float64PointerValue(storageSystem.ThinSavings),
		TotalSize:                       types.Int64PointerValue(storageSystem.TotalSize),
		UnconfiguredSize:                types.Int64PointerValue(storageSystem.UnconfiguredSize),
		UsedPercent:                     types.Float64PointerValue(storageSystem.UsedPercent),
		UsedSize:                        types.Int64PointerValue(storageSystem.UsedSize),
		Vendor:                          types.StringPointerValue(storageSystem.Vendor),
		Version:                         types.StringPointerValue(storageSystem.Version),
		CirrusDeployed:                  types.BoolPointerValue(storageSystem.CirrusDeployed),
	}

	switch {
	case storageSystem.DeploymentDetails == nil:
	case storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails != nil:
		blockStorageState.DeploymentDetails = &deploymentDetailsModel{
			SystemOnPrem: &systemOnPremDeploymentDetailsModel{
				DeploymentType: (storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.DeploymentType),
				SiteName:       types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.SiteName),
				Location:       types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.Location),
				Country:        types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.Country),
				State:          types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.State),
				City:           types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.City),
				StreetAddress1: types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.StreetAddress1),
				StreetAddress2: types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.StreetAddress2),
				ZipCode:        types.StringPointerValue(storageSystem.DeploymentDetails.SystemOnPremDeploymentDetails.ZipCode),
			},
		}
	case storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails != nil:
		blockStorageState.DeploymentDetails = &deploymentDetailsModel{
			SystemPublicCloud: &systemPublicCloudDeploymentDetailsModel{
				DeploymentType:         (storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.DeploymentType),
				CloudAccount:           types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.CloudAccount),
				CloudRegion:            types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.CloudRegion),
				VirtualPrivateCloud:    types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.VirtualPrivateCloud),
				CloudManagementAddress: types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.CloudManagementAddress),
				MinimumIops:            types.Int64PointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.MinimumIops),
				MinimumCapacity:        types.Int64PointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.MinimumCapacity),
				TierType:               types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.TierType),
			},
		}

		if storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.CloudType != nil {
			blockStorageState.DeploymentDetails.SystemPublicCloud.CloudType = storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.CloudType
		}
		if storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.AvailabilityZoneTopology != nil {
			blockStorageState.DeploymentDetails.SystemPublicCloud.AvailabilityZoneTopology = storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.AvailabilityZoneTopology
		}

	default:
		fmt.Printf("Unexpected system type")
	}

	return blockStorageState
}

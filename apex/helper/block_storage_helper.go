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
	"fmt"
	"math"
	"net/http"

	"github.com/dell/terraform-provider-apex/apex/models"
	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GetBlockStorageCollection gets list of block storage instances
func GetBlockStorageCollection(client *client.APIClient, filter string) (*client.StorageSystemsCollection200Response, *http.Response, error) {
	return client.StorageSystemsAPI.StorageSystemsCollection(context.Background()).Limit(500).Filter(filter).Execute()
}

// GetBlockStorageInstance gets block storage instance
func GetBlockStorageInstance(client *client.APIClient, id string) (*client.StorageSystemsInstance, *http.Response, error) {
	return client.StorageSystemsAPI.StorageSystemsInstance(context.Background(), id).Execute()
}

// CreateBlockStorage creates block storage
func CreateBlockStorage(client client.ApiStorageSystemsCreateRequest, systemCreateInput client.StorageSystemDeploymentRequest) (*client.Job, *http.Response, error) {
	return client.StorageSystemDeploymentRequest(systemCreateInput).Async(true).Execute()
}

// SetCloudConfigSubnetAndVpc sets cloud config from plan
func SetCloudConfigSubnetAndVpc(plan models.BlockStorageModel, result models.BlockStorageModel) {
	result.DeploymentDetails.SystemPublicCloud.SSHKeyName = plan.DeploymentDetails.SystemPublicCloud.SSHKeyName
	result.DeploymentDetails.SystemPublicCloud.Vpc = plan.DeploymentDetails.SystemPublicCloud.Vpc
	result.DeploymentDetails.SystemPublicCloud.MinimumCapacity = basetypes.NewInt64Value(result.DeploymentDetails.SystemPublicCloud.MinimumCapacity.ValueInt64() / (int64)(math.Pow(1024, 4)))
	result.DeploymentDetails.SystemPublicCloud.MinimumIops = basetypes.NewInt64Value(result.DeploymentDetails.SystemPublicCloud.MinimumIops.ValueInt64() / 1000)

	for _, subnetOption := range plan.DeploymentDetails.SystemPublicCloud.SubnetOptions {
		result.DeploymentDetails.SystemPublicCloud.SubnetOptions = append(result.DeploymentDetails.SystemPublicCloud.SubnetOptions, models.SubnetOptionModel{
			SubnetID:   subnetOption.SubnetID,
			CidrBlock:  subnetOption.CidrBlock,
			SubnetType: subnetOption.SubnetType,
		})
	}
}

// GetBlockStorageSystem returns a storage systems model based on the given storage system instance
func GetBlockStorageSystem(storageSystem client.StorageSystemsInstance) (blockStorageState models.BlockStorageModel) {
	blockStorageState = models.BlockStorageModel{
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
		LicenseExpirationDateTimestamp:  types.StringValue(ConvertTimeToString(storageSystem.LicenseExpirationDateTimestamp)),
		ContractCoverageType:            types.StringPointerValue(storageSystem.ContractCoverageType),
		ContractExpirationDateTimestamp: types.StringValue(ConvertTimeToString(storageSystem.ContractExpirationDateTimestamp)),
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
		LastContactTimestamp:            types.StringValue(ConvertTimeToString(storageSystem.LastContactTimestamp)),
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
		blockStorageState.DeploymentDetails = &models.DeploymentDetailsModel{
			SystemOnPrem: &models.SystemOnPremDeploymentDetailsModel{
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
		blockStorageState.DeploymentDetails = &models.DeploymentDetailsModel{
			SystemPublicCloud: &models.SystemPublicCloudDeploymentDetailsModel{
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

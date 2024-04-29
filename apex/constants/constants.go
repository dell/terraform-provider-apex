/*
Copyright (c) 2022-2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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

package constants

const (

	// General Messages ####################################################//

	// FilterGeneralErrorMsg is the Error message for general filter failure
	FilterGeneralErrorMsg = "one more more of the ids set in the filter is invalid."

	// FilterErrorSystemTypeMsg is the Error message for system type filter failure
	FilterErrorSystemTypeMsg = "Failed to filter storage product."

	// GeneralConfigureErrorMsg is the Error message for general configuration failure
	GeneralConfigureErrorMsg = "Expected *openapi.Client, got: %T. Please report this issue to the provider developers."

	// DatasourceConfigureErrorMsg is the Error message for datasource configuration failure
	DatasourceConfigureErrorMsg = "Unexpected Data Source Configure Type"

	// ResourceConfigureErrorMsg is the Error message for resource configuration failure
	ResourceConfigureErrorMsg = "Unexpected Resource Configure Type"

	// UnexpectedSysteType is the error message for an unexpected system type
	UnexpectedSysteType = "Unexpected system type"

	// UpdateJobErrorMsg is the Error message for update job failure
	UpdateJobErrorMsg = "Error occurred during update job: "

	// GeneralJobError is the Error message for general job failure
	GeneralJobError = "Error occurred during job"

	// ErrorGettingJob is the Error message for getting job failure
	ErrorGettingJob = "Error getting job"

	// UpdatesNotSupportedErrorMsg is the Error message for updates not supported
	UpdatesNotSupportedErrorMsg = "Updates are not supported for this resource"

	// DeleteIsNotSupportedErrorMsg is the Error message for delete not supported
	DeleteIsNotSupportedErrorMsg = "Delete is not supported for this resource"

	// JobRetrieveError is the Error message for retrieveing job ID failure
	JobRetrieveError = "Could not retrieve job, unexpected error: "

	// PowerFlex Activation Messages ####################################################//

	// ErrorActivatingPowerFlexSystem is the Error message for powerflex activation failure
	ErrorActivatingPowerFlexSystem = "Error activating Powerflex System"

	// ErrorActivatingPowerFlexSystemDetail is the Error message for powerflex activation failure
	ErrorActivatingPowerFlexSystemDetail = "Could not activate powerflex system, please check username/password and system id are correct: "

	// ErrorActivatingPowerFlexSpecificSystemDetail is the Error message for powerflex activation failure
	ErrorActivatingPowerFlexSpecificSystemDetail = "Could not activate powerflex system %s, please check username/password are correct: %s"

	// AWS Account Messages ####################################################//

	// AwsAccountReadErrorMsg is the Error message for read failure
	AwsAccountReadErrorMsg = "Unable to Read Apex Navigator Aws Accounts"

	// AwsAccountReadDetailMsg is the Error message for read failure
	AwsAccountReadDetailMsg = "Could not read aws account, unexpected error: "

	// AwsAccountFilterErrorMsg is the Error message for filter failure
	AwsAccountFilterErrorMsg = "Failed to filter aws accounts."

	// AwsConnectionErrorMsg is the Error message for connection failure
	AwsConnectionErrorMsg = "Error connecting AWS Account"

	// AwsConnectionDetailMsg is the Error message for connection failure
	AwsConnectionDetailMsg = "Could not connect AWS Account, unexpected error: "

	// AwsDisconnectionErrorMsg is the Error message for disconnection failure
	AwsDisconnectionErrorMsg = "Error disconnecting AWS Account"

	// AwsDisconnectionDetailMsg is the Error message for disconnection failure
	AwsDisconnectionDetailMsg = "Could not disconnect AWS Account, unexpected error: "

	// AWS Permission Messages ####################################################//

	// AwsPermissionReadErrorMsg is the Error message for read failure
	AwsPermissionReadErrorMsg = "Unable to Read Apex Navigator AWS Permission Policies"

	// AwsPermissionFilterErrorMsg is the Error message for filter failure
	AwsPermissionFilterErrorMsg = "Failed to filter AWS Permission Policies."

	// AwsAccountUpdateErrorMsg is the Error message for update failure
	AwsAccountUpdateErrorMsg = "Error Updating Apex Navigator Aws Accounts"

	// AwsAccountUpdateDetailMsg is the Error message for update failure
	AwsAccountUpdateDetailMsg = "Could not update aws account, unexpected error: "

	// AwsAccountUpdateIDErrorMsg is the Error message for update failure
	AwsAccountUpdateIDErrorMsg = "Account ID cannot be updated"

	// AwsGeneratePolicyErrorMsg is the Error message for update failure
	AwsGeneratePolicyErrorMsg = "Error generating Trust Policy"

	// AwsGeneratePolicyDetailMsg is the Error message for update failure
	AwsGeneratePolicyDetailMsg = "Could not generate Trust Policy, unexpected error: "

	// AwsUpdateInvalidFieldUpdateErrorMsg is the Error message for update failure
	AwsUpdateInvalidFieldUpdateErrorMsg = "[account_id] cannot be updated, please use the lifecycle block to make sure this is run as a new create each terraform apply"

	// Block Clone Messages ####################################################//

	// BlockCloneReadErrorMsg is the Error message for read failure
	BlockCloneReadErrorMsg = "Unable to Read Apex Navigator Clones"

	// BlockCloneReadDetailMsg is the Error message for read failure
	BlockCloneReadDetailMsg = "Could not read clone, unexpected error: "

	// BlockCloneFilterErrorMsg is the Error message for filter failure
	BlockCloneFilterErrorMsg = "Failed to filter clones."

	// BlockCloneMapError fails to map clones
	BlockCloneMapError = "Failed to map clones."

	// BlockCloneMapDetailErrMsg is the Error message for map failure
	BlockCloneMapDetailErrMsg = "Could not map clones, unexpected error: "

	// BlockCloneRefreshErrorMsg is the Error message for refresh failure
	BlockCloneRefreshErrorMsg = "Error refreshing Apex Navigator Clones"

	// BlockCloneRefreshDetailMsg is the Error message for refresh failure
	BlockCloneRefreshDetailMsg = "Could not refresh clone, unexpected error: "

	// BlockCloneUnmapErrorMsg is the Error message for unmap failure
	BlockCloneUnmapErrorMsg = "Error unmapping Apex Navigator Clones"

	// BlockCloneUnmapDetailMsg is the Error message for unmap failure
	BlockCloneUnmapDetailMsg = "Could not unmap clone, unexpected error: "

	// BlockCloneCreateErrorMsg is the Error message for create failure
	BlockCloneCreateErrorMsg = "Error creating Apex Navigator Clones"

	// BlockCloneCreateDetailMsg is the Error message for create failure
	BlockCloneCreateDetailMsg = "Could not create clone, unexpected error: "

	// BlockCloneUpdateErrorMsg is the Error message for update failure
	BlockCloneUpdateErrorMsg = "Error updating Apex Navigator Clones"

	// BlockCloneUpdateDetailMsg is the Error message for update failure
	BlockCloneUpdateDetailMsg = "Could not update clone, unexpected error: "

	// BlockCloneInvalidFieldUpdateErrorMsg is the Error message for update failure
	BlockCloneInvalidFieldUpdateErrorMsg = "Could not update Clones, invalid field modified [Modifiable fields: Name, Description]"

	// BlockCloneDeleteErrorMsg is the Error message for delete failure
	BlockCloneDeleteErrorMsg = "Error deleting Apex Navigator Clones"

	// BlockCloneDeleteDetailMsg is the Error message for delete failure
	BlockCloneDeleteDetailMsg = "Could not delete clone, unexpected error: "

	// BlockCloneImportReadErrorMsg is the Error message for read failure
	BlockCloneImportReadErrorMsg = "Could not retrieve Clones during import: "

	// Block Hosts Messages ####################################################//

	// BlockHostsReadErrorMsg is the Error message for read failure
	BlockHostsReadErrorMsg = "Unable to Read Apex Navigator Hosts"

	// BlockHostsFilterErrorMsg is the Error message for filter failure
	BlockHostsFilterErrorMsg = "Failed to filter hosts."

	// Block Mobility Group Messages ####################################################//

	// BlockMobilityGroupReadErrorMsg is the Error message for read failure
	BlockMobilityGroupReadErrorMsg = "Unable to Read Apex Navigator Mobility Groups"

	// BlockMobilityGroupReadDetailMsg is the Error message for read failure
	BlockMobilityGroupReadDetailMsg = "Could not read mobility group, unexpected error: "

	// BlockMobilityGroupFilterErrorMsg is the Error message for filter failure
	BlockMobilityGroupFilterErrorMsg = "Failed to filter mobility groups."

	// BlockMobilityGroupCopyErrorMsg is the Error message for copy failure
	BlockMobilityGroupCopyErrorMsg = "Error copying Apex Navigator Mobility Groups"

	// BlockMobilityGroupCopyDetailMsg is the Error message for copy failure
	BlockMobilityGroupCopyDetailMsg = "Could not copy mobility group, unexpected error: "

	// MobilityGroupCreateErrorMsg is the Error message for create failure
	MobilityGroupCreateErrorMsg = "Error creating Apex Navigator Mobility Groups"

	// MobilityGroupCreateDetailMsg is the Error message for create failure
	MobilityGroupCreateDetailMsg = "Could not create mobility group, unexpected error: "

	// MobilityGroupUpdateErrorMsg is the Error message for update failure
	MobilityGroupUpdateErrorMsg = "Error updating Apex Navigator Mobility Groups"

	// MobilityGroupUpdateInvalidFieldUpdateErrorMsg is the Error message for update failure
	MobilityGroupUpdateInvalidFieldUpdateErrorMsg = "Cannot update Mobility Group, attempted to update unchangeable attribute [SystemID, SystemType]"

	// MobilityGroupUpdateDetailMsg is the Error message for update failure
	MobilityGroupUpdateDetailMsg = "Could not update mobility group, unexpected error: "

	// MobilityGroupDeleteErrorMsg is the Error message for delete failure
	MobilityGroupDeleteErrorMsg = "Error deleting Apex Navigator Mobility Groups"

	// MobilityGroupDeleteDetailMsg is the Error message for delete failure
	MobilityGroupDeleteDetailMsg = "Could not delete mobility group, unexpected error: "

	// MobilityGroupImportReadErrorMsg is the Error message for read failure
	MobilityGroupImportReadErrorMsg = "Could not retrieve Mobility Groups during import: "

	// Block Mobility Target Messages ####################################################//

	// BlockMobilityTargetReadErrorMsg is the Error message for read failure
	BlockMobilityTargetReadErrorMsg = "Unable to Read Apex Navigator Mobility Targets"

	// BlockMobilityTargetReadDetailMsg is the Error message for read failure
	BlockMobilityTargetReadDetailMsg = "Could not read mobility target, unexpected error: "

	// BlockMobilityTargetFilterErrorMsg is the Error message for filter failure
	BlockMobilityTargetFilterErrorMsg = "Failed to filter mobility targets."

	// BlockMobilityTargetCreateErrorMsg is the Error message for create failure
	BlockMobilityTargetCreateErrorMsg = "Error creating Apex Navigator Mobility Targets"

	// BlockMobilityTargetCreateDetailMsg is the Error message for create failure
	BlockMobilityTargetCreateDetailMsg = "Could not create mobility target, unexpected error: "

	// BlockMobilityTargetUpdateErrorMsg is the Error message for update failure
	BlockMobilityTargetUpdateErrorMsg = "Error updating Apex Navigator Mobility Targets"

	// BlockMobilityTargetUpdateInvalidFieldUpdateErrorMsg is the Error message for update failure
	BlockMobilityTargetUpdateInvalidFieldUpdateErrorMsg = "Attempted to modify unchangeable attribute(s) [SystemID, SystemType, sourceMobilityGroupID]"

	// BlockMobilityTargetUpdateDetailMsg is the Error message for update failure
	BlockMobilityTargetUpdateDetailMsg = "Could not update mobility target, unexpected error: "

	// BlockMobilityTargetDeleteErrorMsg is the Error message for delete failure
	BlockMobilityTargetDeleteErrorMsg = "Error deleting Apex Navigator Mobility Targets"

	// BlockMobilityTargetDeleteDetailMsg is the Error message for delete failure
	BlockMobilityTargetDeleteDetailMsg = "Could not delete mobility target, unexpected error: "

	// BlockMobilityTargetImportReadErrorMsg is the Error message for read failure
	BlockMobilityTargetImportReadErrorMsg = "Could not retrieve Mobility Targets during import: "

	// Block Pools Messages ####################################################//

	// BlockPoolsReadErrorMsg is the Error message for read failure
	BlockPoolsReadErrorMsg = "Unable to Read Apex Navigator Pools"

	// BlockPoolsFilterErrorMsg is the Error message for filter failure
	BlockPoolsFilterErrorMsg = "Failed to filter Apex Navigator Pools."

	// Block Volumes Messages ####################################################//

	// BlockVolumesReadErrorMsg is the Error message for read failure
	BlockVolumesReadErrorMsg = "Unable to Read Apex Navigator Volumes"

	// BlockVolumesFilterErrorMsg is the Error message for filter failure
	BlockVolumesFilterErrorMsg = "Failed to filter volumes."

	// Storage Products Messages ####################################################//

	// StorageProductReadErrorMsg is the Error message for read failure
	StorageProductReadErrorMsg = "Unable to Read Apex Navigator Storage Products"

	// StorageProductFilterErrorMsg is the Error message for filter failure
	StorageProductFilterErrorMsg = "the system_type in the filter is invalid."

	// Storage Messages ####################################################//

	// StorageReadErrorMsg is the Error message for read failure
	StorageReadErrorMsg = "Unable to Read Apex Navigator Storage"

	// StorageReadDetailMsg is the Error message for read failure
	StorageReadDetailMsg = "Could not read storage, unexpected error: "

	// StorageFilterErrorMsg is the Error message for filter failure
	StorageFilterErrorMsg = "Failed to filter storage."

	// BlockStorageInvalidDeploymentType invalid deployment type
	BlockStorageInvalidDeploymentType = "Invalid deployment type."

	// BlockStorageCreateErrorMsg is the Error message for create failure
	BlockStorageCreateErrorMsg = "Error creating Apex Navigator Block Storage"

	// BlockStorageCreateDetailMsg is the Error message for create failure
	BlockStorageCreateDetailMsg = "Could not create storage, unexpected error: "

	// BlockStorageUpdateErrorMsg is the Error message for update failure
	BlockStorageUpdateErrorMsg = "Error updating Apex Navigator Block Storage"

	// BlockStorageUpdateDetailMsg is the Error message for update failure
	BlockStorageUpdateDetailMsg = "Update of Block Storage is not supported, please create a new resource"

	// BlockStorageDecomissionErrorMsg is the Error message for decomission failure
	BlockStorageDecomissionErrorMsg = "Error decomissioning Apex Navigator Block Storage"

	// BlockStorageDecomissionDetailMsg is the Error message for decomission failure
	BlockStorageDecomissionDetailMsg = "Could not decomission block storage, unexpected error: "
)

# SystemPublicCloudDeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | Pointer to [**SystemDeploymentTypeEnum**](SystemDeploymentTypeEnum.md) |  | [optional] [default to SYSTEMDEPLOYMENTTYPEENUM_ONPREM]
**CloudType** | Pointer to [**CloudProviderEnum**](CloudProviderEnum.md) |  | [optional] [default to CLOUDPROVIDERENUM_AWS]
**CloudAccount** | Pointer to **string** | Cloud provider account where the storage system resides. | [optional] 
**CloudRegion** | Pointer to **string** | Cloud provider region where the storage system resides. | [optional] 
**AvailabilityZoneTopology** | Pointer to [**AvailabilityZoneTopologyEnum**](AvailabilityZoneTopologyEnum.md) |  | [optional] 
**VirtualPrivateCloud** | Pointer to **string** | Cloud virtual private environment identifier. | [optional] 
**CloudManagementAddress** | Pointer to **string** | Management IPv4 or IPv6 address or DNS name of the storage system in cloud. | [optional] 
**MinimumIops** | Pointer to **int64** | Minimum IOPS requested during the deployment time - Unit: IO/s | [optional] 
**MinimumCapacity** | Pointer to **int64** | Minimum capacity requested during the deployment time - Unit: bytes | [optional] 
**TierType** | Pointer to **string** | Tier type requested during the deployment time. | [optional] 

## Methods

### NewSystemPublicCloudDeploymentDetails

`func NewSystemPublicCloudDeploymentDetails() *SystemPublicCloudDeploymentDetails`

NewSystemPublicCloudDeploymentDetails instantiates a new SystemPublicCloudDeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPublicCloudDeploymentDetailsWithDefaults

`func NewSystemPublicCloudDeploymentDetailsWithDefaults() *SystemPublicCloudDeploymentDetails`

NewSystemPublicCloudDeploymentDetailsWithDefaults instantiates a new SystemPublicCloudDeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *SystemPublicCloudDeploymentDetails) GetDeploymentType() SystemDeploymentTypeEnum`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *SystemPublicCloudDeploymentDetails) GetDeploymentTypeOk() (*SystemDeploymentTypeEnum, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *SystemPublicCloudDeploymentDetails) SetDeploymentType(v SystemDeploymentTypeEnum)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *SystemPublicCloudDeploymentDetails) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetCloudType

`func (o *SystemPublicCloudDeploymentDetails) GetCloudType() CloudProviderEnum`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *SystemPublicCloudDeploymentDetails) GetCloudTypeOk() (*CloudProviderEnum, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *SystemPublicCloudDeploymentDetails) SetCloudType(v CloudProviderEnum)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *SystemPublicCloudDeploymentDetails) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudAccount

`func (o *SystemPublicCloudDeploymentDetails) GetCloudAccount() string`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *SystemPublicCloudDeploymentDetails) GetCloudAccountOk() (*string, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *SystemPublicCloudDeploymentDetails) SetCloudAccount(v string)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *SystemPublicCloudDeploymentDetails) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### GetCloudRegion

`func (o *SystemPublicCloudDeploymentDetails) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *SystemPublicCloudDeploymentDetails) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *SystemPublicCloudDeploymentDetails) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *SystemPublicCloudDeploymentDetails) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetAvailabilityZoneTopology

`func (o *SystemPublicCloudDeploymentDetails) GetAvailabilityZoneTopology() AvailabilityZoneTopologyEnum`

GetAvailabilityZoneTopology returns the AvailabilityZoneTopology field if non-nil, zero value otherwise.

### GetAvailabilityZoneTopologyOk

`func (o *SystemPublicCloudDeploymentDetails) GetAvailabilityZoneTopologyOk() (*AvailabilityZoneTopologyEnum, bool)`

GetAvailabilityZoneTopologyOk returns a tuple with the AvailabilityZoneTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneTopology

`func (o *SystemPublicCloudDeploymentDetails) SetAvailabilityZoneTopology(v AvailabilityZoneTopologyEnum)`

SetAvailabilityZoneTopology sets AvailabilityZoneTopology field to given value.

### HasAvailabilityZoneTopology

`func (o *SystemPublicCloudDeploymentDetails) HasAvailabilityZoneTopology() bool`

HasAvailabilityZoneTopology returns a boolean if a field has been set.

### GetVirtualPrivateCloud

`func (o *SystemPublicCloudDeploymentDetails) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *SystemPublicCloudDeploymentDetails) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *SystemPublicCloudDeploymentDetails) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.

### HasVirtualPrivateCloud

`func (o *SystemPublicCloudDeploymentDetails) HasVirtualPrivateCloud() bool`

HasVirtualPrivateCloud returns a boolean if a field has been set.

### GetCloudManagementAddress

`func (o *SystemPublicCloudDeploymentDetails) GetCloudManagementAddress() string`

GetCloudManagementAddress returns the CloudManagementAddress field if non-nil, zero value otherwise.

### GetCloudManagementAddressOk

`func (o *SystemPublicCloudDeploymentDetails) GetCloudManagementAddressOk() (*string, bool)`

GetCloudManagementAddressOk returns a tuple with the CloudManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManagementAddress

`func (o *SystemPublicCloudDeploymentDetails) SetCloudManagementAddress(v string)`

SetCloudManagementAddress sets CloudManagementAddress field to given value.

### HasCloudManagementAddress

`func (o *SystemPublicCloudDeploymentDetails) HasCloudManagementAddress() bool`

HasCloudManagementAddress returns a boolean if a field has been set.

### GetMinimumIops

`func (o *SystemPublicCloudDeploymentDetails) GetMinimumIops() int64`

GetMinimumIops returns the MinimumIops field if non-nil, zero value otherwise.

### GetMinimumIopsOk

`func (o *SystemPublicCloudDeploymentDetails) GetMinimumIopsOk() (*int64, bool)`

GetMinimumIopsOk returns a tuple with the MinimumIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIops

`func (o *SystemPublicCloudDeploymentDetails) SetMinimumIops(v int64)`

SetMinimumIops sets MinimumIops field to given value.

### HasMinimumIops

`func (o *SystemPublicCloudDeploymentDetails) HasMinimumIops() bool`

HasMinimumIops returns a boolean if a field has been set.

### GetMinimumCapacity

`func (o *SystemPublicCloudDeploymentDetails) GetMinimumCapacity() int64`

GetMinimumCapacity returns the MinimumCapacity field if non-nil, zero value otherwise.

### GetMinimumCapacityOk

`func (o *SystemPublicCloudDeploymentDetails) GetMinimumCapacityOk() (*int64, bool)`

GetMinimumCapacityOk returns a tuple with the MinimumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCapacity

`func (o *SystemPublicCloudDeploymentDetails) SetMinimumCapacity(v int64)`

SetMinimumCapacity sets MinimumCapacity field to given value.

### HasMinimumCapacity

`func (o *SystemPublicCloudDeploymentDetails) HasMinimumCapacity() bool`

HasMinimumCapacity returns a boolean if a field has been set.

### GetTierType

`func (o *SystemPublicCloudDeploymentDetails) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *SystemPublicCloudDeploymentDetails) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *SystemPublicCloudDeploymentDetails) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *SystemPublicCloudDeploymentDetails) HasTierType() bool`

HasTierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



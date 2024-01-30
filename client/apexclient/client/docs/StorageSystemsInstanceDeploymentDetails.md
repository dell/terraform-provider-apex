# StorageSystemsInstanceDeploymentDetails

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
**SiteName** | Pointer to **string** | Name of the site where the system is located. | [optional] 
**Location** | Pointer to **string** | User provided description of where the system is located. | [optional] 
**Country** | Pointer to **string** | Name of the country where the system is located. | [optional] 
**State** | Pointer to **string** | Name of the state where the system is located. | [optional] 
**City** | Pointer to **string** | Name of the city where the system is located. | [optional] 
**StreetAddress1** | Pointer to **string** | Street address 1 of where the system is located. | [optional] 
**StreetAddress2** | Pointer to **string** | Street address 2 of where the system is located. | [optional] 
**ZipCode** | Pointer to **string** | State ZIP code of where the system is located. | [optional] 

## Methods

### NewStorageSystemsInstanceDeploymentDetails

`func NewStorageSystemsInstanceDeploymentDetails() *StorageSystemsInstanceDeploymentDetails`

NewStorageSystemsInstanceDeploymentDetails instantiates a new StorageSystemsInstanceDeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemsInstanceDeploymentDetailsWithDefaults

`func NewStorageSystemsInstanceDeploymentDetailsWithDefaults() *StorageSystemsInstanceDeploymentDetails`

NewStorageSystemsInstanceDeploymentDetailsWithDefaults instantiates a new StorageSystemsInstanceDeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *StorageSystemsInstanceDeploymentDetails) GetDeploymentType() SystemDeploymentTypeEnum`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetDeploymentTypeOk() (*SystemDeploymentTypeEnum, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *StorageSystemsInstanceDeploymentDetails) SetDeploymentType(v SystemDeploymentTypeEnum)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *StorageSystemsInstanceDeploymentDetails) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetCloudType

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudType() CloudProviderEnum`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudTypeOk() (*CloudProviderEnum, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *StorageSystemsInstanceDeploymentDetails) SetCloudType(v CloudProviderEnum)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *StorageSystemsInstanceDeploymentDetails) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudAccount

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudAccount() string`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudAccountOk() (*string, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *StorageSystemsInstanceDeploymentDetails) SetCloudAccount(v string)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *StorageSystemsInstanceDeploymentDetails) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### GetCloudRegion

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *StorageSystemsInstanceDeploymentDetails) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *StorageSystemsInstanceDeploymentDetails) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetAvailabilityZoneTopology

`func (o *StorageSystemsInstanceDeploymentDetails) GetAvailabilityZoneTopology() AvailabilityZoneTopologyEnum`

GetAvailabilityZoneTopology returns the AvailabilityZoneTopology field if non-nil, zero value otherwise.

### GetAvailabilityZoneTopologyOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetAvailabilityZoneTopologyOk() (*AvailabilityZoneTopologyEnum, bool)`

GetAvailabilityZoneTopologyOk returns a tuple with the AvailabilityZoneTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneTopology

`func (o *StorageSystemsInstanceDeploymentDetails) SetAvailabilityZoneTopology(v AvailabilityZoneTopologyEnum)`

SetAvailabilityZoneTopology sets AvailabilityZoneTopology field to given value.

### HasAvailabilityZoneTopology

`func (o *StorageSystemsInstanceDeploymentDetails) HasAvailabilityZoneTopology() bool`

HasAvailabilityZoneTopology returns a boolean if a field has been set.

### GetVirtualPrivateCloud

`func (o *StorageSystemsInstanceDeploymentDetails) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *StorageSystemsInstanceDeploymentDetails) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.

### HasVirtualPrivateCloud

`func (o *StorageSystemsInstanceDeploymentDetails) HasVirtualPrivateCloud() bool`

HasVirtualPrivateCloud returns a boolean if a field has been set.

### GetCloudManagementAddress

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudManagementAddress() string`

GetCloudManagementAddress returns the CloudManagementAddress field if non-nil, zero value otherwise.

### GetCloudManagementAddressOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCloudManagementAddressOk() (*string, bool)`

GetCloudManagementAddressOk returns a tuple with the CloudManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManagementAddress

`func (o *StorageSystemsInstanceDeploymentDetails) SetCloudManagementAddress(v string)`

SetCloudManagementAddress sets CloudManagementAddress field to given value.

### HasCloudManagementAddress

`func (o *StorageSystemsInstanceDeploymentDetails) HasCloudManagementAddress() bool`

HasCloudManagementAddress returns a boolean if a field has been set.

### GetMinimumIops

`func (o *StorageSystemsInstanceDeploymentDetails) GetMinimumIops() int64`

GetMinimumIops returns the MinimumIops field if non-nil, zero value otherwise.

### GetMinimumIopsOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetMinimumIopsOk() (*int64, bool)`

GetMinimumIopsOk returns a tuple with the MinimumIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIops

`func (o *StorageSystemsInstanceDeploymentDetails) SetMinimumIops(v int64)`

SetMinimumIops sets MinimumIops field to given value.

### HasMinimumIops

`func (o *StorageSystemsInstanceDeploymentDetails) HasMinimumIops() bool`

HasMinimumIops returns a boolean if a field has been set.

### GetMinimumCapacity

`func (o *StorageSystemsInstanceDeploymentDetails) GetMinimumCapacity() int64`

GetMinimumCapacity returns the MinimumCapacity field if non-nil, zero value otherwise.

### GetMinimumCapacityOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetMinimumCapacityOk() (*int64, bool)`

GetMinimumCapacityOk returns a tuple with the MinimumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCapacity

`func (o *StorageSystemsInstanceDeploymentDetails) SetMinimumCapacity(v int64)`

SetMinimumCapacity sets MinimumCapacity field to given value.

### HasMinimumCapacity

`func (o *StorageSystemsInstanceDeploymentDetails) HasMinimumCapacity() bool`

HasMinimumCapacity returns a boolean if a field has been set.

### GetTierType

`func (o *StorageSystemsInstanceDeploymentDetails) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *StorageSystemsInstanceDeploymentDetails) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *StorageSystemsInstanceDeploymentDetails) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### GetSiteName

`func (o *StorageSystemsInstanceDeploymentDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *StorageSystemsInstanceDeploymentDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *StorageSystemsInstanceDeploymentDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetLocation

`func (o *StorageSystemsInstanceDeploymentDetails) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageSystemsInstanceDeploymentDetails) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StorageSystemsInstanceDeploymentDetails) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountry

`func (o *StorageSystemsInstanceDeploymentDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StorageSystemsInstanceDeploymentDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StorageSystemsInstanceDeploymentDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *StorageSystemsInstanceDeploymentDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageSystemsInstanceDeploymentDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageSystemsInstanceDeploymentDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *StorageSystemsInstanceDeploymentDetails) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StorageSystemsInstanceDeploymentDetails) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *StorageSystemsInstanceDeploymentDetails) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *StorageSystemsInstanceDeploymentDetails) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *StorageSystemsInstanceDeploymentDetails) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *StorageSystemsInstanceDeploymentDetails) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.

### HasStreetAddress1

`func (o *StorageSystemsInstanceDeploymentDetails) HasStreetAddress1() bool`

HasStreetAddress1 returns a boolean if a field has been set.

### GetStreetAddress2

`func (o *StorageSystemsInstanceDeploymentDetails) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *StorageSystemsInstanceDeploymentDetails) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *StorageSystemsInstanceDeploymentDetails) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *StorageSystemsInstanceDeploymentDetails) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetZipCode

`func (o *StorageSystemsInstanceDeploymentDetails) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *StorageSystemsInstanceDeploymentDetails) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *StorageSystemsInstanceDeploymentDetails) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *StorageSystemsInstanceDeploymentDetails) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



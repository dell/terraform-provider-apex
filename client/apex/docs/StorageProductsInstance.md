# StorageProductsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | Name of the storage product | 
**SystemType** | [**StorageProductEnum**](StorageProductEnum.md) |  | 
**StorageType** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Description** | **string** | Description of the storage product and its capabilities | 
**CloudType** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | [default to CLOUDPROVIDERENUM_AWS]
**LatestVersion** | **string** | Latest supported version of the storage product on the cloud | 
**SupportMap** | [**[]SupportMap**](SupportMap.md) | Supported cloud and version details | 

## Methods

### NewStorageProductsInstance

`func NewStorageProductsInstance(id string, name string, systemType StorageProductEnum, storageType StorageTypeEnum, description string, cloudType CloudProviderEnum, latestVersion string, supportMap []SupportMap, ) *StorageProductsInstance`

NewStorageProductsInstance instantiates a new StorageProductsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProductsInstanceWithDefaults

`func NewStorageProductsInstanceWithDefaults() *StorageProductsInstance`

NewStorageProductsInstanceWithDefaults instantiates a new StorageProductsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageProductsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageProductsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageProductsInstance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StorageProductsInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProductsInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProductsInstance) SetName(v string)`

SetName sets Name field to given value.


### GetSystemType

`func (o *StorageProductsInstance) GetSystemType() StorageProductEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *StorageProductsInstance) GetSystemTypeOk() (*StorageProductEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *StorageProductsInstance) SetSystemType(v StorageProductEnum)`

SetSystemType sets SystemType field to given value.


### GetStorageType

`func (o *StorageProductsInstance) GetStorageType() StorageTypeEnum`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *StorageProductsInstance) GetStorageTypeOk() (*StorageTypeEnum, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *StorageProductsInstance) SetStorageType(v StorageTypeEnum)`

SetStorageType sets StorageType field to given value.


### GetDescription

`func (o *StorageProductsInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageProductsInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageProductsInstance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCloudType

`func (o *StorageProductsInstance) GetCloudType() CloudProviderEnum`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *StorageProductsInstance) GetCloudTypeOk() (*CloudProviderEnum, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *StorageProductsInstance) SetCloudType(v CloudProviderEnum)`

SetCloudType sets CloudType field to given value.


### GetLatestVersion

`func (o *StorageProductsInstance) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *StorageProductsInstance) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *StorageProductsInstance) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### GetSupportMap

`func (o *StorageProductsInstance) GetSupportMap() []SupportMap`

GetSupportMap returns the SupportMap field if non-nil, zero value otherwise.

### GetSupportMapOk

`func (o *StorageProductsInstance) GetSupportMapOk() (*[]SupportMap, bool)`

GetSupportMapOk returns a tuple with the SupportMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMap

`func (o *StorageProductsInstance) SetSupportMap(v []SupportMap)`

SetSupportMap sets SupportMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



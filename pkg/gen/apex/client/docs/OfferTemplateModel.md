# OfferTemplateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**SystemType** | [**StorageProductEnum**](StorageProductEnum.md) |  | 
**StorageProductVersion** | **string** | Version of the storage product | 
**CloudType** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | [default to CLOUDPROVIDERENUM_AWS]
**SupportedTierInfo** | [**[]TierInfo**](TierInfo.md) | Supported tier information of storage product  | 

## Methods

### NewOfferTemplateModel

`func NewOfferTemplateModel(systemType StorageProductEnum, storageProductVersion string, cloudType CloudProviderEnum, supportedTierInfo []TierInfo, ) *OfferTemplateModel`

NewOfferTemplateModel instantiates a new OfferTemplateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferTemplateModelWithDefaults

`func NewOfferTemplateModelWithDefaults() *OfferTemplateModel`

NewOfferTemplateModelWithDefaults instantiates a new OfferTemplateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OfferTemplateModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfferTemplateModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfferTemplateModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OfferTemplateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSystemType

`func (o *OfferTemplateModel) GetSystemType() StorageProductEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *OfferTemplateModel) GetSystemTypeOk() (*StorageProductEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *OfferTemplateModel) SetSystemType(v StorageProductEnum)`

SetSystemType sets SystemType field to given value.


### GetStorageProductVersion

`func (o *OfferTemplateModel) GetStorageProductVersion() string`

GetStorageProductVersion returns the StorageProductVersion field if non-nil, zero value otherwise.

### GetStorageProductVersionOk

`func (o *OfferTemplateModel) GetStorageProductVersionOk() (*string, bool)`

GetStorageProductVersionOk returns a tuple with the StorageProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProductVersion

`func (o *OfferTemplateModel) SetStorageProductVersion(v string)`

SetStorageProductVersion sets StorageProductVersion field to given value.


### GetCloudType

`func (o *OfferTemplateModel) GetCloudType() CloudProviderEnum`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *OfferTemplateModel) GetCloudTypeOk() (*CloudProviderEnum, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *OfferTemplateModel) SetCloudType(v CloudProviderEnum)`

SetCloudType sets CloudType field to given value.


### GetSupportedTierInfo

`func (o *OfferTemplateModel) GetSupportedTierInfo() []TierInfo`

GetSupportedTierInfo returns the SupportedTierInfo field if non-nil, zero value otherwise.

### GetSupportedTierInfoOk

`func (o *OfferTemplateModel) GetSupportedTierInfoOk() (*[]TierInfo, bool)`

GetSupportedTierInfoOk returns a tuple with the SupportedTierInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTierInfo

`func (o *OfferTemplateModel) SetSupportedTierInfo(v []TierInfo)`

SetSupportedTierInfo sets SupportedTierInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TierInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | Tier name | 
**TierType** | [**TierEnum**](TierEnum.md) |  | 
**Description** | **string** | Description of the tier | 
**StorageOptions** | [**[]TierOption**](TierOption.md) | Supported storage options for a tier | 
**CloudOptions** | [**[]TierOption**](TierOption.md) | Supported cloud options for a tier | 

## Methods

### NewTierInfo

`func NewTierInfo(id string, name string, tierType TierEnum, description string, storageOptions []TierOption, cloudOptions []TierOption, ) *TierInfo`

NewTierInfo instantiates a new TierInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierInfoWithDefaults

`func NewTierInfoWithDefaults() *TierInfo`

NewTierInfoWithDefaults instantiates a new TierInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TierInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TierInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TierInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TierInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierInfo) SetName(v string)`

SetName sets Name field to given value.


### GetTierType

`func (o *TierInfo) GetTierType() TierEnum`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *TierInfo) GetTierTypeOk() (*TierEnum, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *TierInfo) SetTierType(v TierEnum)`

SetTierType sets TierType field to given value.


### GetDescription

`func (o *TierInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TierInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TierInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStorageOptions

`func (o *TierInfo) GetStorageOptions() []TierOption`

GetStorageOptions returns the StorageOptions field if non-nil, zero value otherwise.

### GetStorageOptionsOk

`func (o *TierInfo) GetStorageOptionsOk() (*[]TierOption, bool)`

GetStorageOptionsOk returns a tuple with the StorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptions

`func (o *TierInfo) SetStorageOptions(v []TierOption)`

SetStorageOptions sets StorageOptions field to given value.


### GetCloudOptions

`func (o *TierInfo) GetCloudOptions() []TierOption`

GetCloudOptions returns the CloudOptions field if non-nil, zero value otherwise.

### GetCloudOptionsOk

`func (o *TierInfo) GetCloudOptionsOk() (*[]TierOption, bool)`

GetCloudOptionsOk returns a tuple with the CloudOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudOptions

`func (o *TierInfo) SetCloudOptions(v []TierOption)`

SetCloudOptions sets CloudOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



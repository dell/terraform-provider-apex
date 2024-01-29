# PowerFlexStorageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemType** | **string** | Type of the Dell Apex storage system to be deployed. Only \&quot;PowerFlex\&quot; is supported in this version. | 
**Version** | Pointer to **string** | Storage product version to deploy | [optional] 
**Tier** | Pointer to [**PowerFlexTierEnum**](PowerFlexTierEnum.md) |  | [optional] 
**MinimumIops** | Pointer to **int32** |  | [optional] 
**MinimumUsableCapacity** | Pointer to **int32** |  | [optional] 

## Methods

### NewPowerFlexStorageOptions

`func NewPowerFlexStorageOptions(systemType string, ) *PowerFlexStorageOptions`

NewPowerFlexStorageOptions instantiates a new PowerFlexStorageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerFlexStorageOptionsWithDefaults

`func NewPowerFlexStorageOptionsWithDefaults() *PowerFlexStorageOptions`

NewPowerFlexStorageOptionsWithDefaults instantiates a new PowerFlexStorageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemType

`func (o *PowerFlexStorageOptions) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *PowerFlexStorageOptions) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *PowerFlexStorageOptions) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.


### GetVersion

`func (o *PowerFlexStorageOptions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PowerFlexStorageOptions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PowerFlexStorageOptions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PowerFlexStorageOptions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTier

`func (o *PowerFlexStorageOptions) GetTier() PowerFlexTierEnum`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PowerFlexStorageOptions) GetTierOk() (*PowerFlexTierEnum, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PowerFlexStorageOptions) SetTier(v PowerFlexTierEnum)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PowerFlexStorageOptions) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetMinimumIops

`func (o *PowerFlexStorageOptions) GetMinimumIops() int32`

GetMinimumIops returns the MinimumIops field if non-nil, zero value otherwise.

### GetMinimumIopsOk

`func (o *PowerFlexStorageOptions) GetMinimumIopsOk() (*int32, bool)`

GetMinimumIopsOk returns a tuple with the MinimumIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIops

`func (o *PowerFlexStorageOptions) SetMinimumIops(v int32)`

SetMinimumIops sets MinimumIops field to given value.

### HasMinimumIops

`func (o *PowerFlexStorageOptions) HasMinimumIops() bool`

HasMinimumIops returns a boolean if a field has been set.

### GetMinimumUsableCapacity

`func (o *PowerFlexStorageOptions) GetMinimumUsableCapacity() int32`

GetMinimumUsableCapacity returns the MinimumUsableCapacity field if non-nil, zero value otherwise.

### GetMinimumUsableCapacityOk

`func (o *PowerFlexStorageOptions) GetMinimumUsableCapacityOk() (*int32, bool)`

GetMinimumUsableCapacityOk returns a tuple with the MinimumUsableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUsableCapacity

`func (o *PowerFlexStorageOptions) SetMinimumUsableCapacity(v int32)`

SetMinimumUsableCapacity sets MinimumUsableCapacity field to given value.

### HasMinimumUsableCapacity

`func (o *PowerFlexStorageOptions) HasMinimumUsableCapacity() bool`

HasMinimumUsableCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



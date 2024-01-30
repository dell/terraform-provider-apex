# TargetSystemOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoragePool** | Pointer to **string** | Storage pool to use for allocsting target volumes | [optional] 

## Methods

### NewTargetSystemOptions

`func NewTargetSystemOptions() *TargetSystemOptions`

NewTargetSystemOptions instantiates a new TargetSystemOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetSystemOptionsWithDefaults

`func NewTargetSystemOptionsWithDefaults() *TargetSystemOptions`

NewTargetSystemOptionsWithDefaults instantiates a new TargetSystemOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoragePool

`func (o *TargetSystemOptions) GetStoragePool() string`

GetStoragePool returns the StoragePool field if non-nil, zero value otherwise.

### GetStoragePoolOk

`func (o *TargetSystemOptions) GetStoragePoolOk() (*string, bool)`

GetStoragePoolOk returns a tuple with the StoragePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePool

`func (o *TargetSystemOptions) SetStoragePool(v string)`

SetStoragePool sets StoragePool field to given value.

### HasStoragePool

`func (o *TargetSystemOptions) HasStoragePool() bool`

HasStoragePool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



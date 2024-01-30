# StorageProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemType** | **string** | Type of the Dell Apex storage system to be deployed. Only \&quot;PowerFlex\&quot; is supported in this version. | 
**Version** | Pointer to **string** | Storage product version to deploy | [optional] 

## Methods

### NewStorageProduct

`func NewStorageProduct(systemType string, ) *StorageProduct`

NewStorageProduct instantiates a new StorageProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProductWithDefaults

`func NewStorageProductWithDefaults() *StorageProduct`

NewStorageProductWithDefaults instantiates a new StorageProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemType

`func (o *StorageProduct) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *StorageProduct) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *StorageProduct) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.


### GetVersion

`func (o *StorageProduct) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageProduct) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageProduct) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageProduct) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



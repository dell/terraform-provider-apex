# StorageSystemTokensInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemType** | [**StorageSystemTypeEnum**](StorageSystemTypeEnum.md) |  | 
**SystemId** | **string** | Unique identifier for the storage system. | 
**ExpirationTimestamp** | **time.Time** | Expiration date for the token of the storage system. | 
**IsTokenValid** | **bool** |  | 

## Methods

### NewStorageSystemTokensInstance

`func NewStorageSystemTokensInstance(systemType StorageSystemTypeEnum, systemId string, expirationTimestamp time.Time, isTokenValid bool, ) *StorageSystemTokensInstance`

NewStorageSystemTokensInstance instantiates a new StorageSystemTokensInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemTokensInstanceWithDefaults

`func NewStorageSystemTokensInstanceWithDefaults() *StorageSystemTokensInstance`

NewStorageSystemTokensInstanceWithDefaults instantiates a new StorageSystemTokensInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemType

`func (o *StorageSystemTokensInstance) GetSystemType() StorageSystemTypeEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *StorageSystemTokensInstance) GetSystemTypeOk() (*StorageSystemTypeEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *StorageSystemTokensInstance) SetSystemType(v StorageSystemTypeEnum)`

SetSystemType sets SystemType field to given value.


### GetSystemId

`func (o *StorageSystemTokensInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *StorageSystemTokensInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *StorageSystemTokensInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetExpirationTimestamp

`func (o *StorageSystemTokensInstance) GetExpirationTimestamp() time.Time`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *StorageSystemTokensInstance) GetExpirationTimestampOk() (*time.Time, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *StorageSystemTokensInstance) SetExpirationTimestamp(v time.Time)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetIsTokenValid

`func (o *StorageSystemTokensInstance) GetIsTokenValid() bool`

GetIsTokenValid returns the IsTokenValid field if non-nil, zero value otherwise.

### GetIsTokenValidOk

`func (o *StorageSystemTokensInstance) GetIsTokenValidOk() (*bool, bool)`

GetIsTokenValidOk returns a tuple with the IsTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenValid

`func (o *StorageSystemTokensInstance) SetIsTokenValid(v bool)`

SetIsTokenValid sets IsTokenValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



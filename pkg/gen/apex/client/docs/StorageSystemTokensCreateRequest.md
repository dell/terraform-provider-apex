# StorageSystemTokensCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemType** | [**StorageProductEnum**](StorageProductEnum.md) |  | 
**SystemId** | **string** | Unique identifier for the storage system. | 
**AccessToken** | **string** | Access token of the system. | 

## Methods

### NewStorageSystemTokensCreateRequest

`func NewStorageSystemTokensCreateRequest(systemType StorageProductEnum, systemId string, accessToken string, ) *StorageSystemTokensCreateRequest`

NewStorageSystemTokensCreateRequest instantiates a new StorageSystemTokensCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemTokensCreateRequestWithDefaults

`func NewStorageSystemTokensCreateRequestWithDefaults() *StorageSystemTokensCreateRequest`

NewStorageSystemTokensCreateRequestWithDefaults instantiates a new StorageSystemTokensCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemType

`func (o *StorageSystemTokensCreateRequest) GetSystemType() StorageProductEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *StorageSystemTokensCreateRequest) GetSystemTypeOk() (*StorageProductEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *StorageSystemTokensCreateRequest) SetSystemType(v StorageProductEnum)`

SetSystemType sets SystemType field to given value.


### GetSystemId

`func (o *StorageSystemTokensCreateRequest) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *StorageSystemTokensCreateRequest) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *StorageSystemTokensCreateRequest) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetAccessToken

`func (o *StorageSystemTokensCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *StorageSystemTokensCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *StorageSystemTokensCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



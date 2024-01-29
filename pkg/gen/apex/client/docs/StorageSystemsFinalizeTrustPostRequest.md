# StorageSystemsFinalizeTrustPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Client name for establishing the trust between on-prem storage system and APEX Navigator for Multicloud Storage. | [optional] 
**ClientId** | **string** | Client id for establishing the trust between on-prem storage system and APEX Navigator for Multicloud Storage. | 
**ClientSecret** | **string** | Client secret for establishing the trust between on-prem storage system and APEX Navigator for Multicloud Storage. | 

## Methods

### NewStorageSystemsFinalizeTrustPostRequest

`func NewStorageSystemsFinalizeTrustPostRequest(clientId string, clientSecret string, ) *StorageSystemsFinalizeTrustPostRequest`

NewStorageSystemsFinalizeTrustPostRequest instantiates a new StorageSystemsFinalizeTrustPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemsFinalizeTrustPostRequestWithDefaults

`func NewStorageSystemsFinalizeTrustPostRequestWithDefaults() *StorageSystemsFinalizeTrustPostRequest`

NewStorageSystemsFinalizeTrustPostRequestWithDefaults instantiates a new StorageSystemsFinalizeTrustPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageSystemsFinalizeTrustPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSystemsFinalizeTrustPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSystemsFinalizeTrustPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSystemsFinalizeTrustPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientId

`func (o *StorageSystemsFinalizeTrustPostRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StorageSystemsFinalizeTrustPostRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StorageSystemsFinalizeTrustPostRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *StorageSystemsFinalizeTrustPostRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *StorageSystemsFinalizeTrustPostRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *StorageSystemsFinalizeTrustPostRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageSystemsInitializeTrustResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Client name for establish the trust between on-prem storage system and APEX Navigator for Multicloud Storage. | 
**EncryptionKey** | **string** | Public key to encrypt the client secret. | 
**RedirectUris** | **[]string** | The location where the Storage System Authorization server sends the user once the APEX Navigator for Multicloud Storage  has been successfully authorized and granted an authorization code or access token. | 

## Methods

### NewStorageSystemsInitializeTrustResponse

`func NewStorageSystemsInitializeTrustResponse(name string, encryptionKey string, redirectUris []string, ) *StorageSystemsInitializeTrustResponse`

NewStorageSystemsInitializeTrustResponse instantiates a new StorageSystemsInitializeTrustResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemsInitializeTrustResponseWithDefaults

`func NewStorageSystemsInitializeTrustResponseWithDefaults() *StorageSystemsInitializeTrustResponse`

NewStorageSystemsInitializeTrustResponseWithDefaults instantiates a new StorageSystemsInitializeTrustResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageSystemsInitializeTrustResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSystemsInitializeTrustResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSystemsInitializeTrustResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEncryptionKey

`func (o *StorageSystemsInitializeTrustResponse) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *StorageSystemsInitializeTrustResponse) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *StorageSystemsInitializeTrustResponse) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.


### GetRedirectUris

`func (o *StorageSystemsInitializeTrustResponse) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *StorageSystemsInitializeTrustResponse) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *StorageSystemsInitializeTrustResponse) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



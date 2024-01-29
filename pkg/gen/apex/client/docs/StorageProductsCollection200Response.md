# StorageProductsCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]StorageProductsInstance**](StorageProductsInstance.md) |  | [optional] 
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewStorageProductsCollection200Response

`func NewStorageProductsCollection200Response() *StorageProductsCollection200Response`

NewStorageProductsCollection200Response instantiates a new StorageProductsCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProductsCollection200ResponseWithDefaults

`func NewStorageProductsCollection200ResponseWithDefaults() *StorageProductsCollection200Response`

NewStorageProductsCollection200ResponseWithDefaults instantiates a new StorageProductsCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *StorageProductsCollection200Response) GetResults() []StorageProductsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StorageProductsCollection200Response) GetResultsOk() (*[]StorageProductsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StorageProductsCollection200Response) SetResults(v []StorageProductsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *StorageProductsCollection200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPaging

`func (o *StorageProductsCollection200Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *StorageProductsCollection200Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *StorageProductsCollection200Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *StorageProductsCollection200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



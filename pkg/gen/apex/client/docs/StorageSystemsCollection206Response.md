# StorageSystemsCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]StorageSystemsInstance**](StorageSystemsInstance.md) |  | [optional] 

## Methods

### NewStorageSystemsCollection206Response

`func NewStorageSystemsCollection206Response() *StorageSystemsCollection206Response`

NewStorageSystemsCollection206Response instantiates a new StorageSystemsCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemsCollection206ResponseWithDefaults

`func NewStorageSystemsCollection206ResponseWithDefaults() *StorageSystemsCollection206Response`

NewStorageSystemsCollection206ResponseWithDefaults instantiates a new StorageSystemsCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *StorageSystemsCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *StorageSystemsCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *StorageSystemsCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *StorageSystemsCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *StorageSystemsCollection206Response) GetResults() []StorageSystemsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StorageSystemsCollection206Response) GetResultsOk() (*[]StorageSystemsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StorageSystemsCollection206Response) SetResults(v []StorageSystemsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *StorageSystemsCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



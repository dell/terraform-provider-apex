# PoolsCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]PoolsInstance**](PoolsInstance.md) |  | [optional] 

## Methods

### NewPoolsCollection206Response

`func NewPoolsCollection206Response() *PoolsCollection206Response`

NewPoolsCollection206Response instantiates a new PoolsCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsCollection206ResponseWithDefaults

`func NewPoolsCollection206ResponseWithDefaults() *PoolsCollection206Response`

NewPoolsCollection206ResponseWithDefaults instantiates a new PoolsCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *PoolsCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PoolsCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PoolsCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PoolsCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *PoolsCollection206Response) GetResults() []PoolsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PoolsCollection206Response) GetResultsOk() (*[]PoolsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PoolsCollection206Response) SetResults(v []PoolsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *PoolsCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



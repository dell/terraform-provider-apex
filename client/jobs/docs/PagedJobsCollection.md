# PagedJobsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]JobsInstance**](JobsInstance.md) |  | [optional] 

## Methods

### NewPagedJobsCollection

`func NewPagedJobsCollection() *PagedJobsCollection`

NewPagedJobsCollection instantiates a new PagedJobsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedJobsCollectionWithDefaults

`func NewPagedJobsCollectionWithDefaults() *PagedJobsCollection`

NewPagedJobsCollectionWithDefaults instantiates a new PagedJobsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *PagedJobsCollection) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PagedJobsCollection) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PagedJobsCollection) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PagedJobsCollection) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *PagedJobsCollection) GetResults() []JobsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PagedJobsCollection) GetResultsOk() (*[]JobsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PagedJobsCollection) SetResults(v []JobsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *PagedJobsCollection) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



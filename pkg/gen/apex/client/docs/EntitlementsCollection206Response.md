# EntitlementsCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]EntitlementsInstance**](EntitlementsInstance.md) |  | [optional] 

## Methods

### NewEntitlementsCollection206Response

`func NewEntitlementsCollection206Response() *EntitlementsCollection206Response`

NewEntitlementsCollection206Response instantiates a new EntitlementsCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsCollection206ResponseWithDefaults

`func NewEntitlementsCollection206ResponseWithDefaults() *EntitlementsCollection206Response`

NewEntitlementsCollection206ResponseWithDefaults instantiates a new EntitlementsCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *EntitlementsCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *EntitlementsCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *EntitlementsCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *EntitlementsCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *EntitlementsCollection206Response) GetResults() []EntitlementsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EntitlementsCollection206Response) GetResultsOk() (*[]EntitlementsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EntitlementsCollection206Response) SetResults(v []EntitlementsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *EntitlementsCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ClonesCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]Clone**](Clone.md) |  | [optional] 

## Methods

### NewClonesCollection206Response

`func NewClonesCollection206Response() *ClonesCollection206Response`

NewClonesCollection206Response instantiates a new ClonesCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonesCollection206ResponseWithDefaults

`func NewClonesCollection206ResponseWithDefaults() *ClonesCollection206Response`

NewClonesCollection206ResponseWithDefaults instantiates a new ClonesCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *ClonesCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ClonesCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ClonesCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ClonesCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *ClonesCollection206Response) GetResults() []Clone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ClonesCollection206Response) GetResultsOk() (*[]Clone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ClonesCollection206Response) SetResults(v []Clone)`

SetResults sets Results field to given value.

### HasResults

`func (o *ClonesCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



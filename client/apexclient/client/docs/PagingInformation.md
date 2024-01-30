# PagingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInstances** | Pointer to **int32** | The number of instances in the complete queried collection (not just the size of the current response). If the request was filtered, then this is the size of the complete filtered collection. | [optional] 
**First** | Pointer to **string** | The URL to fetch the first page of this collection (including filtering and sorting parameters). | [optional] 
**Last** | Pointer to **string** | The URL to fetch the last page of this collection (including filtering and sorting parameters). | [optional] 
**Next** | Pointer to **string** | The URL to fetch the next page of this collection (including filtering and sorting parameters). This will not be returned with the last page of a collection. | [optional] 
**Prev** | Pointer to **string** | The URL to fetch the previous page of this collection (including filtering and sorting parameters). This will not be returned with the first page of a collection. | [optional] 

## Methods

### NewPagingInformation

`func NewPagingInformation() *PagingInformation`

NewPagingInformation instantiates a new PagingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingInformationWithDefaults

`func NewPagingInformationWithDefaults() *PagingInformation`

NewPagingInformationWithDefaults instantiates a new PagingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInstances

`func (o *PagingInformation) GetTotalInstances() int32`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *PagingInformation) GetTotalInstancesOk() (*int32, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *PagingInformation) SetTotalInstances(v int32)`

SetTotalInstances sets TotalInstances field to given value.

### HasTotalInstances

`func (o *PagingInformation) HasTotalInstances() bool`

HasTotalInstances returns a boolean if a field has been set.

### GetFirst

`func (o *PagingInformation) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PagingInformation) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PagingInformation) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PagingInformation) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PagingInformation) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PagingInformation) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PagingInformation) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *PagingInformation) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *PagingInformation) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PagingInformation) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PagingInformation) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PagingInformation) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *PagingInformation) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PagingInformation) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PagingInformation) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PagingInformation) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



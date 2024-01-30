# MobilityTargetsCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]MobilityTarget**](MobilityTarget.md) |  | [optional] 

## Methods

### NewMobilityTargetsCollection206Response

`func NewMobilityTargetsCollection206Response() *MobilityTargetsCollection206Response`

NewMobilityTargetsCollection206Response instantiates a new MobilityTargetsCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobilityTargetsCollection206ResponseWithDefaults

`func NewMobilityTargetsCollection206ResponseWithDefaults() *MobilityTargetsCollection206Response`

NewMobilityTargetsCollection206ResponseWithDefaults instantiates a new MobilityTargetsCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *MobilityTargetsCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *MobilityTargetsCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *MobilityTargetsCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *MobilityTargetsCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *MobilityTargetsCollection206Response) GetResults() []MobilityTarget`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MobilityTargetsCollection206Response) GetResultsOk() (*[]MobilityTarget, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MobilityTargetsCollection206Response) SetResults(v []MobilityTarget)`

SetResults sets Results field to given value.

### HasResults

`func (o *MobilityTargetsCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



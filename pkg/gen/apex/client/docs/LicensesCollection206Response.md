# LicensesCollection206Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Results** | Pointer to [**[]LicensesInstance**](LicensesInstance.md) |  | [optional] 

## Methods

### NewLicensesCollection206Response

`func NewLicensesCollection206Response() *LicensesCollection206Response`

NewLicensesCollection206Response instantiates a new LicensesCollection206Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensesCollection206ResponseWithDefaults

`func NewLicensesCollection206ResponseWithDefaults() *LicensesCollection206Response`

NewLicensesCollection206ResponseWithDefaults instantiates a new LicensesCollection206Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *LicensesCollection206Response) GetPaging() PagingInformation`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *LicensesCollection206Response) GetPagingOk() (*PagingInformation, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *LicensesCollection206Response) SetPaging(v PagingInformation)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *LicensesCollection206Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *LicensesCollection206Response) GetResults() []LicensesInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LicensesCollection206Response) GetResultsOk() (*[]LicensesInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LicensesCollection206Response) SetResults(v []LicensesInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *LicensesCollection206Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



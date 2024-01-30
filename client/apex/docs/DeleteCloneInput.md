# DeleteCloneInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveFromSystem** | **bool** | If true remove the clone from the system if false the clone will remain on the system but will be removed from APEX management. | 

## Methods

### NewDeleteCloneInput

`func NewDeleteCloneInput(removeFromSystem bool, ) *DeleteCloneInput`

NewDeleteCloneInput instantiates a new DeleteCloneInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCloneInputWithDefaults

`func NewDeleteCloneInputWithDefaults() *DeleteCloneInput`

NewDeleteCloneInputWithDefaults instantiates a new DeleteCloneInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveFromSystem

`func (o *DeleteCloneInput) GetRemoveFromSystem() bool`

GetRemoveFromSystem returns the RemoveFromSystem field if non-nil, zero value otherwise.

### GetRemoveFromSystemOk

`func (o *DeleteCloneInput) GetRemoveFromSystemOk() (*bool, bool)`

GetRemoveFromSystemOk returns a tuple with the RemoveFromSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFromSystem

`func (o *DeleteCloneInput) SetRemoveFromSystem(v bool)`

SetRemoveFromSystem sets RemoveFromSystem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



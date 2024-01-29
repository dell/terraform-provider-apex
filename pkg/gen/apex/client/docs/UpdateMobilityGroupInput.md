# UpdateMobilityGroupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Updated name of the mobility group | [optional] 
**Description** | Pointer to **string** | Updated description of the mobility group | [optional] 
**Members** | Pointer to **[]string** | Update the list of members (e.g. volumes).  This is the list of identifiers of members for the mobility group.  It will replace the current list of members. | [optional] 

## Methods

### NewUpdateMobilityGroupInput

`func NewUpdateMobilityGroupInput() *UpdateMobilityGroupInput`

NewUpdateMobilityGroupInput instantiates a new UpdateMobilityGroupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMobilityGroupInputWithDefaults

`func NewUpdateMobilityGroupInputWithDefaults() *UpdateMobilityGroupInput`

NewUpdateMobilityGroupInputWithDefaults instantiates a new UpdateMobilityGroupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMobilityGroupInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMobilityGroupInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMobilityGroupInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMobilityGroupInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMobilityGroupInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMobilityGroupInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMobilityGroupInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMobilityGroupInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembers

`func (o *UpdateMobilityGroupInput) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateMobilityGroupInput) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateMobilityGroupInput) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *UpdateMobilityGroupInput) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



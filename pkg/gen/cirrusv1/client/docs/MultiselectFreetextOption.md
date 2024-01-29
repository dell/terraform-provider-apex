# MultiselectFreetextOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | **string** | Description of the option | 
**AvailableOptions** | **[]string** | All the available values for the option | 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** | Type of the option | [default to "MULTISELECT_FREETEXT"]

## Methods

### NewMultiselectFreetextOption

`func NewMultiselectFreetextOption(name string, description string, availableOptions []string, isConfigurable bool, type_ string, ) *MultiselectFreetextOption`

NewMultiselectFreetextOption instantiates a new MultiselectFreetextOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiselectFreetextOptionWithDefaults

`func NewMultiselectFreetextOptionWithDefaults() *MultiselectFreetextOption`

NewMultiselectFreetextOptionWithDefaults instantiates a new MultiselectFreetextOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MultiselectFreetextOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultiselectFreetextOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultiselectFreetextOption) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MultiselectFreetextOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiselectFreetextOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiselectFreetextOption) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAvailableOptions

`func (o *MultiselectFreetextOption) GetAvailableOptions() []string`

GetAvailableOptions returns the AvailableOptions field if non-nil, zero value otherwise.

### GetAvailableOptionsOk

`func (o *MultiselectFreetextOption) GetAvailableOptionsOk() (*[]string, bool)`

GetAvailableOptionsOk returns a tuple with the AvailableOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOptions

`func (o *MultiselectFreetextOption) SetAvailableOptions(v []string)`

SetAvailableOptions sets AvailableOptions field to given value.


### GetIsConfigurable

`func (o *MultiselectFreetextOption) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *MultiselectFreetextOption) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *MultiselectFreetextOption) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *MultiselectFreetextOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiselectFreetextOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiselectFreetextOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



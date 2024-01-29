# EnumOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | **string** | Description of the option | 
**AvailableOptions** | **[]string** | All the available values for the option | 
**DefaultValue** | Pointer to **string** | Default value for the option | [optional] 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** | Type of the option | [default to "ENUMERATION"]

## Methods

### NewEnumOption

`func NewEnumOption(name string, description string, availableOptions []string, isConfigurable bool, type_ string, ) *EnumOption`

NewEnumOption instantiates a new EnumOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumOptionWithDefaults

`func NewEnumOptionWithDefaults() *EnumOption`

NewEnumOptionWithDefaults instantiates a new EnumOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnumOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumOption) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnumOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumOption) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAvailableOptions

`func (o *EnumOption) GetAvailableOptions() []string`

GetAvailableOptions returns the AvailableOptions field if non-nil, zero value otherwise.

### GetAvailableOptionsOk

`func (o *EnumOption) GetAvailableOptionsOk() (*[]string, bool)`

GetAvailableOptionsOk returns a tuple with the AvailableOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOptions

`func (o *EnumOption) SetAvailableOptions(v []string)`

SetAvailableOptions sets AvailableOptions field to given value.


### GetDefaultValue

`func (o *EnumOption) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EnumOption) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EnumOption) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EnumOption) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *EnumOption) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *EnumOption) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *EnumOption) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *EnumOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnumOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnumOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BooleanOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | Pointer to **string** | Description for the option | [optional] 
**DefaultValue** | Pointer to **bool** | User input value for the option | [optional] 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** |  | [default to "BOOLEAN"]

## Methods

### NewBooleanOption

`func NewBooleanOption(name string, isConfigurable bool, type_ string, ) *BooleanOption`

NewBooleanOption instantiates a new BooleanOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanOptionWithDefaults

`func NewBooleanOptionWithDefaults() *BooleanOption`

NewBooleanOptionWithDefaults instantiates a new BooleanOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BooleanOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BooleanOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BooleanOption) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BooleanOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooleanOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooleanOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooleanOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *BooleanOption) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BooleanOption) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BooleanOption) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BooleanOption) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *BooleanOption) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *BooleanOption) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *BooleanOption) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *BooleanOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooleanOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooleanOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



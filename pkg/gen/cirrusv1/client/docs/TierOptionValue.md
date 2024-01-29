# TierOptionValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | **string** | Description for the option | 
**MinValue** | **int64** | Minimum value for the range selection | 
**MaxValue** | **int64** | Maximum value for the range selection | 
**LinearInterval** | **int64** | Linear interval for the range to increase | 
**Unit** | **string** | Unit for the range | 
**DefaultValue** | Pointer to **bool** | User input value for the option | [optional] 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** |  | [default to "BOOLEAN"]
**AvailableOptions** | **[]string** | All the available values for the option | 

## Methods

### NewTierOptionValue

`func NewTierOptionValue(name string, description string, minValue int64, maxValue int64, linearInterval int64, unit string, isConfigurable bool, type_ string, availableOptions []string, ) *TierOptionValue`

NewTierOptionValue instantiates a new TierOptionValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierOptionValueWithDefaults

`func NewTierOptionValueWithDefaults() *TierOptionValue`

NewTierOptionValueWithDefaults instantiates a new TierOptionValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TierOptionValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierOptionValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierOptionValue) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TierOptionValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TierOptionValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TierOptionValue) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMinValue

`func (o *TierOptionValue) GetMinValue() int64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *TierOptionValue) GetMinValueOk() (*int64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *TierOptionValue) SetMinValue(v int64)`

SetMinValue sets MinValue field to given value.


### GetMaxValue

`func (o *TierOptionValue) GetMaxValue() int64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *TierOptionValue) GetMaxValueOk() (*int64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *TierOptionValue) SetMaxValue(v int64)`

SetMaxValue sets MaxValue field to given value.


### GetLinearInterval

`func (o *TierOptionValue) GetLinearInterval() int64`

GetLinearInterval returns the LinearInterval field if non-nil, zero value otherwise.

### GetLinearIntervalOk

`func (o *TierOptionValue) GetLinearIntervalOk() (*int64, bool)`

GetLinearIntervalOk returns a tuple with the LinearInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearInterval

`func (o *TierOptionValue) SetLinearInterval(v int64)`

SetLinearInterval sets LinearInterval field to given value.


### GetUnit

`func (o *TierOptionValue) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TierOptionValue) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TierOptionValue) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDefaultValue

`func (o *TierOptionValue) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *TierOptionValue) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *TierOptionValue) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *TierOptionValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *TierOptionValue) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *TierOptionValue) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *TierOptionValue) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *TierOptionValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TierOptionValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TierOptionValue) SetType(v string)`

SetType sets Type field to given value.


### GetAvailableOptions

`func (o *TierOptionValue) GetAvailableOptions() []string`

GetAvailableOptions returns the AvailableOptions field if non-nil, zero value otherwise.

### GetAvailableOptionsOk

`func (o *TierOptionValue) GetAvailableOptionsOk() (*[]string, bool)`

GetAvailableOptionsOk returns a tuple with the AvailableOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOptions

`func (o *TierOptionValue) SetAvailableOptions(v []string)`

SetAvailableOptions sets AvailableOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



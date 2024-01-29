# RangeOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | Pointer to **string** | Description of the option | [optional] 
**MinValue** | **int64** | Minimum value for the range selection | 
**MaxValue** | **int64** | Maximum value for the range selection | 
**LinearInterval** | **int64** | Linear interval for the range to increase | 
**Unit** | **string** | Unit for the range | 
**DefaultValue** | Pointer to **int64** | Default value for the range | [optional] 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** | Type of the option | [default to "RANGE"]

## Methods

### NewRangeOption

`func NewRangeOption(name string, minValue int64, maxValue int64, linearInterval int64, unit string, isConfigurable bool, type_ string, ) *RangeOption`

NewRangeOption instantiates a new RangeOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeOptionWithDefaults

`func NewRangeOptionWithDefaults() *RangeOption`

NewRangeOptionWithDefaults instantiates a new RangeOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RangeOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RangeOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RangeOption) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RangeOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RangeOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RangeOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RangeOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMinValue

`func (o *RangeOption) GetMinValue() int64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *RangeOption) GetMinValueOk() (*int64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *RangeOption) SetMinValue(v int64)`

SetMinValue sets MinValue field to given value.


### GetMaxValue

`func (o *RangeOption) GetMaxValue() int64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *RangeOption) GetMaxValueOk() (*int64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *RangeOption) SetMaxValue(v int64)`

SetMaxValue sets MaxValue field to given value.


### GetLinearInterval

`func (o *RangeOption) GetLinearInterval() int64`

GetLinearInterval returns the LinearInterval field if non-nil, zero value otherwise.

### GetLinearIntervalOk

`func (o *RangeOption) GetLinearIntervalOk() (*int64, bool)`

GetLinearIntervalOk returns a tuple with the LinearInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearInterval

`func (o *RangeOption) SetLinearInterval(v int64)`

SetLinearInterval sets LinearInterval field to given value.


### GetUnit

`func (o *RangeOption) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RangeOption) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RangeOption) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDefaultValue

`func (o *RangeOption) GetDefaultValue() int64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RangeOption) GetDefaultValueOk() (*int64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RangeOption) SetDefaultValue(v int64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RangeOption) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *RangeOption) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *RangeOption) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *RangeOption) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *RangeOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RangeOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RangeOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



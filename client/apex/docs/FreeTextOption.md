# FreeTextOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the option | 
**Description** | Pointer to **string** | Description for the option | [optional] 
**IsConfigurable** | **bool** | Is it mandatory | 
**Type** | **string** |  | [default to "FREETEXT"]

## Methods

### NewFreeTextOption

`func NewFreeTextOption(name string, isConfigurable bool, type_ string, ) *FreeTextOption`

NewFreeTextOption instantiates a new FreeTextOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeTextOptionWithDefaults

`func NewFreeTextOptionWithDefaults() *FreeTextOption`

NewFreeTextOptionWithDefaults instantiates a new FreeTextOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FreeTextOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FreeTextOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FreeTextOption) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FreeTextOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FreeTextOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FreeTextOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FreeTextOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *FreeTextOption) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *FreeTextOption) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *FreeTextOption) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.


### GetType

`func (o *FreeTextOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeTextOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeTextOption) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



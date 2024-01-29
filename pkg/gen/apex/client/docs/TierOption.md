# TierOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Key** | **string** | Unique key for tier option | 
**Order** | **int32** | Order of the option selection | 
**Value** | [**TierOptionValue**](TierOptionValue.md) |  | 

## Methods

### NewTierOption

`func NewTierOption(id string, key string, order int32, value TierOptionValue, ) *TierOption`

NewTierOption instantiates a new TierOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierOptionWithDefaults

`func NewTierOptionWithDefaults() *TierOption`

NewTierOptionWithDefaults instantiates a new TierOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TierOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TierOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TierOption) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *TierOption) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TierOption) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TierOption) SetKey(v string)`

SetKey sets Key field to given value.


### GetOrder

`func (o *TierOption) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TierOption) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TierOption) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetValue

`func (o *TierOption) GetValue() TierOptionValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TierOption) GetValueOk() (*TierOptionValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TierOption) SetValue(v TierOptionValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AwsTrustPolicyStatementInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to [**AwsTrustPolicyStatementInnerPrincipal**](AwsTrustPolicyStatementInnerPrincipal.md) |  | [optional] 
**Condition** | Pointer to [**AwsTrustPolicyStatementInnerCondition**](AwsTrustPolicyStatementInnerCondition.md) |  | [optional] 

## Methods

### NewAwsTrustPolicyStatementInner

`func NewAwsTrustPolicyStatementInner() *AwsTrustPolicyStatementInner`

NewAwsTrustPolicyStatementInner instantiates a new AwsTrustPolicyStatementInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTrustPolicyStatementInnerWithDefaults

`func NewAwsTrustPolicyStatementInnerWithDefaults() *AwsTrustPolicyStatementInner`

NewAwsTrustPolicyStatementInnerWithDefaults instantiates a new AwsTrustPolicyStatementInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *AwsTrustPolicyStatementInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AwsTrustPolicyStatementInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AwsTrustPolicyStatementInner) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *AwsTrustPolicyStatementInner) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetAction

`func (o *AwsTrustPolicyStatementInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AwsTrustPolicyStatementInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AwsTrustPolicyStatementInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AwsTrustPolicyStatementInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrincipal

`func (o *AwsTrustPolicyStatementInner) GetPrincipal() AwsTrustPolicyStatementInnerPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AwsTrustPolicyStatementInner) GetPrincipalOk() (*AwsTrustPolicyStatementInnerPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AwsTrustPolicyStatementInner) SetPrincipal(v AwsTrustPolicyStatementInnerPrincipal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AwsTrustPolicyStatementInner) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetCondition

`func (o *AwsTrustPolicyStatementInner) GetCondition() AwsTrustPolicyStatementInnerCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AwsTrustPolicyStatementInner) GetConditionOk() (*AwsTrustPolicyStatementInnerCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AwsTrustPolicyStatementInner) SetCondition(v AwsTrustPolicyStatementInnerCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AwsTrustPolicyStatementInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



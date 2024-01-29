# AwsPermissionPoliciesInstancePermissionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Statement** | Pointer to [**[]AwsPermissionPoliciesInstancePermissionPolicyStatementInner**](AwsPermissionPoliciesInstancePermissionPolicyStatementInner.md) |  | [optional] 

## Methods

### NewAwsPermissionPoliciesInstancePermissionPolicy

`func NewAwsPermissionPoliciesInstancePermissionPolicy() *AwsPermissionPoliciesInstancePermissionPolicy`

NewAwsPermissionPoliciesInstancePermissionPolicy instantiates a new AwsPermissionPoliciesInstancePermissionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPermissionPoliciesInstancePermissionPolicyWithDefaults

`func NewAwsPermissionPoliciesInstancePermissionPolicyWithDefaults() *AwsPermissionPoliciesInstancePermissionPolicy`

NewAwsPermissionPoliciesInstancePermissionPolicyWithDefaults instantiates a new AwsPermissionPoliciesInstancePermissionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatement

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) GetStatement() []AwsPermissionPoliciesInstancePermissionPolicyStatementInner`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) GetStatementOk() (*[]AwsPermissionPoliciesInstancePermissionPolicyStatementInner, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) SetStatement(v []AwsPermissionPoliciesInstancePermissionPolicyStatementInner)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *AwsPermissionPoliciesInstancePermissionPolicy) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



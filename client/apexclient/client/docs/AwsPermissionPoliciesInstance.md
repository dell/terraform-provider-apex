# AwsPermissionPoliciesInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**PermissionPolicy** | Pointer to [**AwsPermissionPoliciesInstancePermissionPolicy**](AwsPermissionPoliciesInstancePermissionPolicy.md) |  | [optional] 

## Methods

### NewAwsPermissionPoliciesInstance

`func NewAwsPermissionPoliciesInstance() *AwsPermissionPoliciesInstance`

NewAwsPermissionPoliciesInstance instantiates a new AwsPermissionPoliciesInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPermissionPoliciesInstanceWithDefaults

`func NewAwsPermissionPoliciesInstanceWithDefaults() *AwsPermissionPoliciesInstance`

NewAwsPermissionPoliciesInstanceWithDefaults instantiates a new AwsPermissionPoliciesInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsPermissionPoliciesInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsPermissionPoliciesInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsPermissionPoliciesInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsPermissionPoliciesInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AwsPermissionPoliciesInstance) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AwsPermissionPoliciesInstance) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AwsPermissionPoliciesInstance) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AwsPermissionPoliciesInstance) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPermissionPolicy

`func (o *AwsPermissionPoliciesInstance) GetPermissionPolicy() AwsPermissionPoliciesInstancePermissionPolicy`

GetPermissionPolicy returns the PermissionPolicy field if non-nil, zero value otherwise.

### GetPermissionPolicyOk

`func (o *AwsPermissionPoliciesInstance) GetPermissionPolicyOk() (*AwsPermissionPoliciesInstancePermissionPolicy, bool)`

GetPermissionPolicyOk returns a tuple with the PermissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionPolicy

`func (o *AwsPermissionPoliciesInstance) SetPermissionPolicy(v AwsPermissionPoliciesInstancePermissionPolicy)`

SetPermissionPolicy sets PermissionPolicy field to given value.

### HasPermissionPolicy

`func (o *AwsPermissionPoliciesInstance) HasPermissionPolicy() bool`

HasPermissionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



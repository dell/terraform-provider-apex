# AwsAccountsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Customer&#39;s AWS account id | [optional] 
**AwsAccountAlias** | Pointer to **string** | Customer&#39;s AWS account alias | [optional] 
**Status** | Pointer to **string** | Describes the state of the AWS account that is being added for Dell to manage | [optional] 
**RoleArn** | Pointer to **string** | Amazon Resource Name (ARN) for the role created by customer as part of establishing trust. | [optional] 

## Methods

### NewAwsAccountsInstance

`func NewAwsAccountsInstance() *AwsAccountsInstance`

NewAwsAccountsInstance instantiates a new AwsAccountsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAccountsInstanceWithDefaults

`func NewAwsAccountsInstanceWithDefaults() *AwsAccountsInstance`

NewAwsAccountsInstanceWithDefaults instantiates a new AwsAccountsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsAccountsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsAccountsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsAccountsInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsAccountsInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAwsAccountAlias

`func (o *AwsAccountsInstance) GetAwsAccountAlias() string`

GetAwsAccountAlias returns the AwsAccountAlias field if non-nil, zero value otherwise.

### GetAwsAccountAliasOk

`func (o *AwsAccountsInstance) GetAwsAccountAliasOk() (*string, bool)`

GetAwsAccountAliasOk returns a tuple with the AwsAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountAlias

`func (o *AwsAccountsInstance) SetAwsAccountAlias(v string)`

SetAwsAccountAlias sets AwsAccountAlias field to given value.

### HasAwsAccountAlias

`func (o *AwsAccountsInstance) HasAwsAccountAlias() bool`

HasAwsAccountAlias returns a boolean if a field has been set.

### GetStatus

`func (o *AwsAccountsInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsAccountsInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsAccountsInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsAccountsInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRoleArn

`func (o *AwsAccountsInstance) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsAccountsInstance) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsAccountsInstance) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AwsAccountsInstance) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



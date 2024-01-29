# Vpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNewVpc** | Pointer to **bool** | Create a new VPC. Default is false. | [optional] [default to false]
**Id** | Pointer to **string** | AWS VPC resource identifier.  Provide this value if you want Apex Navigator to deploy the storage system in your existing VPC.  | [optional] 
**Name** | Pointer to **string** | AWS VPC name. Provide only name (not id) if you want Apex Navigator to deploy the storage system in a new VPC.  | [optional] 

## Methods

### NewVpc

`func NewVpc() *Vpc`

NewVpc instantiates a new Vpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcWithDefaults

`func NewVpcWithDefaults() *Vpc`

NewVpcWithDefaults instantiates a new Vpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNewVpc

`func (o *Vpc) GetIsNewVpc() bool`

GetIsNewVpc returns the IsNewVpc field if non-nil, zero value otherwise.

### GetIsNewVpcOk

`func (o *Vpc) GetIsNewVpcOk() (*bool, bool)`

GetIsNewVpcOk returns a tuple with the IsNewVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewVpc

`func (o *Vpc) SetIsNewVpc(v bool)`

SetIsNewVpc sets IsNewVpc field to given value.

### HasIsNewVpc

`func (o *Vpc) HasIsNewVpc() bool`

HasIsNewVpc returns a boolean if a field has been set.

### GetId

`func (o *Vpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vpc) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vpc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Vpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vpc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vpc) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



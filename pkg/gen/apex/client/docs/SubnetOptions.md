# SubnetOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | AWS subnet resource identifier | [optional] 
**CidrBlock** | Pointer to **string** | IP range that Apex Navigator will use to create new subnets and use it based on the  subnet option type.  If this is a new VPC, Apex Navigator also associates this CIDR block to the new VPC as primary or secondary  CIDR block.  | [optional] 
**Type** | Pointer to [**SubnetTypeEnum**](SubnetTypeEnum.md) |  | [optional] [default to SUBNETTYPEENUM_UNDEFINED]

## Methods

### NewSubnetOptions

`func NewSubnetOptions() *SubnetOptions`

NewSubnetOptions instantiates a new SubnetOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetOptionsWithDefaults

`func NewSubnetOptionsWithDefaults() *SubnetOptions`

NewSubnetOptionsWithDefaults instantiates a new SubnetOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubnetOptions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubnetOptions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubnetOptions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubnetOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCidrBlock

`func (o *SubnetOptions) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *SubnetOptions) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *SubnetOptions) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *SubnetOptions) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetType

`func (o *SubnetOptions) GetType() SubnetTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubnetOptions) GetTypeOk() (*SubnetTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubnetOptions) SetType(v SubnetTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *SubnetOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AWSCloudOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** |  | 
**AccountId** | **string** | AWS account ID that uniquely identifies an account. The destination of the AWS resources created for the deployment | 
**Region** | **string** | AWS cloud region for deployment | 
**AvailabilityZoneTopology** | Pointer to [**AvailabilityZoneTopologyEnum**](AvailabilityZoneTopologyEnum.md) |  | [optional] 
**SshKeyName** | **string** | Provide a ssh key name for Apex Navigator to create and use to launch Apex storage system EC2 instance with AWS. You can retrieve the private key from your AWS Secrets Manager using the same name | 
**Vpc** | [**Vpc**](Vpc.md) |  | 
**SubnetOptions** | [**[]SubnetOptions**](SubnetOptions.md) |  | 

## Methods

### NewAWSCloudOptions

`func NewAWSCloudOptions(cloudType string, accountId string, region string, sshKeyName string, vpc Vpc, subnetOptions []SubnetOptions, ) *AWSCloudOptions`

NewAWSCloudOptions instantiates a new AWSCloudOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCloudOptionsWithDefaults

`func NewAWSCloudOptionsWithDefaults() *AWSCloudOptions`

NewAWSCloudOptionsWithDefaults instantiates a new AWSCloudOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *AWSCloudOptions) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AWSCloudOptions) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AWSCloudOptions) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### GetAccountId

`func (o *AWSCloudOptions) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSCloudOptions) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSCloudOptions) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetRegion

`func (o *AWSCloudOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSCloudOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSCloudOptions) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAvailabilityZoneTopology

`func (o *AWSCloudOptions) GetAvailabilityZoneTopology() AvailabilityZoneTopologyEnum`

GetAvailabilityZoneTopology returns the AvailabilityZoneTopology field if non-nil, zero value otherwise.

### GetAvailabilityZoneTopologyOk

`func (o *AWSCloudOptions) GetAvailabilityZoneTopologyOk() (*AvailabilityZoneTopologyEnum, bool)`

GetAvailabilityZoneTopologyOk returns a tuple with the AvailabilityZoneTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneTopology

`func (o *AWSCloudOptions) SetAvailabilityZoneTopology(v AvailabilityZoneTopologyEnum)`

SetAvailabilityZoneTopology sets AvailabilityZoneTopology field to given value.

### HasAvailabilityZoneTopology

`func (o *AWSCloudOptions) HasAvailabilityZoneTopology() bool`

HasAvailabilityZoneTopology returns a boolean if a field has been set.

### GetSshKeyName

`func (o *AWSCloudOptions) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *AWSCloudOptions) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *AWSCloudOptions) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetVpc

`func (o *AWSCloudOptions) GetVpc() Vpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AWSCloudOptions) GetVpcOk() (*Vpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AWSCloudOptions) SetVpc(v Vpc)`

SetVpc sets Vpc field to given value.


### GetSubnetOptions

`func (o *AWSCloudOptions) GetSubnetOptions() []SubnetOptions`

GetSubnetOptions returns the SubnetOptions field if non-nil, zero value otherwise.

### GetSubnetOptionsOk

`func (o *AWSCloudOptions) GetSubnetOptionsOk() (*[]SubnetOptions, bool)`

GetSubnetOptionsOk returns a tuple with the SubnetOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOptions

`func (o *AWSCloudOptions) SetSubnetOptions(v []SubnetOptions)`

SetSubnetOptions sets SubnetOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



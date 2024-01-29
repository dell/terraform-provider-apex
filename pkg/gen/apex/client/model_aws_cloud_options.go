/*
APEX Navigator for Multicloud Storage REST APIs

The public API definitions for APEX Navigator for Multicloud Storage

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AWSCloudOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSCloudOptions{}

// AWSCloudOptions AWS Cloud
type AWSCloudOptions struct {
	CloudType string `json:"cloud_type"`
	// AWS account ID that uniquely identifies an account. The destination of the AWS resources created for the deployment
	AccountId string `json:"account_id"`
	// AWS cloud region for deployment
	Region string `json:"region"`
	AvailabilityZoneTopology *AvailabilityZoneTopologyEnum `json:"availability_zone_topology,omitempty"`
	// Provide a ssh key name for Apex Navigator to create and use to launch Apex storage system EC2 instance with AWS. You can retrieve the private key from your AWS Secrets Manager using the same name
	SshKeyName string `json:"ssh_key_name"`
	Vpc Vpc `json:"vpc"`
	SubnetOptions []SubnetOptions `json:"subnet_options"`
}

type _AWSCloudOptions AWSCloudOptions

// NewAWSCloudOptions instantiates a new AWSCloudOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCloudOptions(cloudType string, accountId string, region string, sshKeyName string, vpc Vpc, subnetOptions []SubnetOptions) *AWSCloudOptions {
	this := AWSCloudOptions{}
	this.CloudType = cloudType
	this.AccountId = accountId
	this.Region = region
	this.SshKeyName = sshKeyName
	this.Vpc = vpc
	this.SubnetOptions = subnetOptions
	return &this
}

// NewAWSCloudOptionsWithDefaults instantiates a new AWSCloudOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCloudOptionsWithDefaults() *AWSCloudOptions {
	this := AWSCloudOptions{}
	return &this
}

// GetCloudType returns the CloudType field value
func (o *AWSCloudOptions) GetCloudType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetCloudTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudType, true
}

// SetCloudType sets field value
func (o *AWSCloudOptions) SetCloudType(v string) {
	o.CloudType = v
}

// GetAccountId returns the AccountId field value
func (o *AWSCloudOptions) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AWSCloudOptions) SetAccountId(v string) {
	o.AccountId = v
}

// GetRegion returns the Region field value
func (o *AWSCloudOptions) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AWSCloudOptions) SetRegion(v string) {
	o.Region = v
}

// GetAvailabilityZoneTopology returns the AvailabilityZoneTopology field value if set, zero value otherwise.
func (o *AWSCloudOptions) GetAvailabilityZoneTopology() AvailabilityZoneTopologyEnum {
	if o == nil || IsNil(o.AvailabilityZoneTopology) {
		var ret AvailabilityZoneTopologyEnum
		return ret
	}
	return *o.AvailabilityZoneTopology
}

// GetAvailabilityZoneTopologyOk returns a tuple with the AvailabilityZoneTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetAvailabilityZoneTopologyOk() (*AvailabilityZoneTopologyEnum, bool) {
	if o == nil || IsNil(o.AvailabilityZoneTopology) {
		return nil, false
	}
	return o.AvailabilityZoneTopology, true
}

// HasAvailabilityZoneTopology returns a boolean if a field has been set.
func (o *AWSCloudOptions) HasAvailabilityZoneTopology() bool {
	if o != nil && !IsNil(o.AvailabilityZoneTopology) {
		return true
	}

	return false
}

// SetAvailabilityZoneTopology gets a reference to the given AvailabilityZoneTopologyEnum and assigns it to the AvailabilityZoneTopology field.
func (o *AWSCloudOptions) SetAvailabilityZoneTopology(v AvailabilityZoneTopologyEnum) {
	o.AvailabilityZoneTopology = &v
}

// GetSshKeyName returns the SshKeyName field value
func (o *AWSCloudOptions) GetSshKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshKeyName
}

// GetSshKeyNameOk returns a tuple with the SshKeyName field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetSshKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshKeyName, true
}

// SetSshKeyName sets field value
func (o *AWSCloudOptions) SetSshKeyName(v string) {
	o.SshKeyName = v
}

// GetVpc returns the Vpc field value
func (o *AWSCloudOptions) GetVpc() Vpc {
	if o == nil {
		var ret Vpc
		return ret
	}

	return o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetVpcOk() (*Vpc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vpc, true
}

// SetVpc sets field value
func (o *AWSCloudOptions) SetVpc(v Vpc) {
	o.Vpc = v
}

// GetSubnetOptions returns the SubnetOptions field value
func (o *AWSCloudOptions) GetSubnetOptions() []SubnetOptions {
	if o == nil {
		var ret []SubnetOptions
		return ret
	}

	return o.SubnetOptions
}

// GetSubnetOptionsOk returns a tuple with the SubnetOptions field value
// and a boolean to check if the value has been set.
func (o *AWSCloudOptions) GetSubnetOptionsOk() ([]SubnetOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetOptions, true
}

// SetSubnetOptions sets field value
func (o *AWSCloudOptions) SetSubnetOptions(v []SubnetOptions) {
	o.SubnetOptions = v
}

func (o AWSCloudOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSCloudOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_type"] = o.CloudType
	toSerialize["account_id"] = o.AccountId
	toSerialize["region"] = o.Region
	if !IsNil(o.AvailabilityZoneTopology) {
		toSerialize["availability_zone_topology"] = o.AvailabilityZoneTopology
	}
	toSerialize["ssh_key_name"] = o.SshKeyName
	toSerialize["vpc"] = o.Vpc
	toSerialize["subnet_options"] = o.SubnetOptions
	return toSerialize, nil
}

func (o *AWSCloudOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_type",
		"account_id",
		"region",
		"ssh_key_name",
		"vpc",
		"subnet_options",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAWSCloudOptions := _AWSCloudOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAWSCloudOptions)

	if err != nil {
		return err
	}

	*o = AWSCloudOptions(varAWSCloudOptions)

	return err
}

type NullableAWSCloudOptions struct {
	value *AWSCloudOptions
	isSet bool
}

func (v NullableAWSCloudOptions) Get() *AWSCloudOptions {
	return v.value
}

func (v *NullableAWSCloudOptions) Set(val *AWSCloudOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSCloudOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSCloudOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSCloudOptions(val *AWSCloudOptions) *NullableAWSCloudOptions {
	return &NullableAWSCloudOptions{value: val, isSet: true}
}

func (v NullableAWSCloudOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSCloudOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



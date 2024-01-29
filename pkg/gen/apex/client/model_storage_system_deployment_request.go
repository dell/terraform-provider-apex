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

// checks if the StorageSystemDeploymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageSystemDeploymentRequest{}

// StorageSystemDeploymentRequest Storage system deployment request input object
type StorageSystemDeploymentRequest struct {
	// Unique name to identify the deployed system
	Name string `json:"name"`
	CloudOptions StorageSystemDeploymentRequestCloudOptions `json:"cloud_options"`
	StorageOptions StorageSystemDeploymentRequestStorageOptions `json:"storage_options"`
	// By setting this option to true, you, on your behalf of your company, agree the following will apply:  Your evaluation of Dell APEX Navigator for Multicloud is subject to and governed by Dell’s Cloud Service Offering Agreement (https://www.dell.com/learn/us/en/uscorp1/legal_terms-conditions_dellwebpage/csoa-agreement) and  the Dell APEX Navigator for Multicloud Storage Service Offering Description (https://www.dell.com/learn/us/en/uscorp1/apex-services).  Your evaluation of Dell APEX Block Storage for AWS is subject to and governed by Dell’s Software Evaluation Agreement  (https://vault.pactsafe.io/s/3cb1f636-b99f-47ab-ad13-3ce168930a55/legal.html?g=38113#contract-s12zpzeii).  Your evaluation of Dell APEX Navigator for Multicloud will be available for 90 days free of charge. At the end of the evaluation, the Dell APEX Navigator for Multicloud Storage evaluation service and associated data will remain accessible, subject to the terms identified in this paragraph, however, certain features will be limited until you order a new subscription. Payment for any required public cloud service provider infrastructure requirements is your company’s responsibility.
	IsTermsAndConditionsAgreed bool `json:"is_terms_and_conditions_agreed"`
}

type _StorageSystemDeploymentRequest StorageSystemDeploymentRequest

// NewStorageSystemDeploymentRequest instantiates a new StorageSystemDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSystemDeploymentRequest(name string, cloudOptions StorageSystemDeploymentRequestCloudOptions, storageOptions StorageSystemDeploymentRequestStorageOptions, isTermsAndConditionsAgreed bool) *StorageSystemDeploymentRequest {
	this := StorageSystemDeploymentRequest{}
	this.Name = name
	this.CloudOptions = cloudOptions
	this.StorageOptions = storageOptions
	this.IsTermsAndConditionsAgreed = isTermsAndConditionsAgreed
	return &this
}

// NewStorageSystemDeploymentRequestWithDefaults instantiates a new StorageSystemDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSystemDeploymentRequestWithDefaults() *StorageSystemDeploymentRequest {
	this := StorageSystemDeploymentRequest{}
	return &this
}

// GetName returns the Name field value
func (o *StorageSystemDeploymentRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageSystemDeploymentRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageSystemDeploymentRequest) SetName(v string) {
	o.Name = v
}

// GetCloudOptions returns the CloudOptions field value
func (o *StorageSystemDeploymentRequest) GetCloudOptions() StorageSystemDeploymentRequestCloudOptions {
	if o == nil {
		var ret StorageSystemDeploymentRequestCloudOptions
		return ret
	}

	return o.CloudOptions
}

// GetCloudOptionsOk returns a tuple with the CloudOptions field value
// and a boolean to check if the value has been set.
func (o *StorageSystemDeploymentRequest) GetCloudOptionsOk() (*StorageSystemDeploymentRequestCloudOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudOptions, true
}

// SetCloudOptions sets field value
func (o *StorageSystemDeploymentRequest) SetCloudOptions(v StorageSystemDeploymentRequestCloudOptions) {
	o.CloudOptions = v
}

// GetStorageOptions returns the StorageOptions field value
func (o *StorageSystemDeploymentRequest) GetStorageOptions() StorageSystemDeploymentRequestStorageOptions {
	if o == nil {
		var ret StorageSystemDeploymentRequestStorageOptions
		return ret
	}

	return o.StorageOptions
}

// GetStorageOptionsOk returns a tuple with the StorageOptions field value
// and a boolean to check if the value has been set.
func (o *StorageSystemDeploymentRequest) GetStorageOptionsOk() (*StorageSystemDeploymentRequestStorageOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageOptions, true
}

// SetStorageOptions sets field value
func (o *StorageSystemDeploymentRequest) SetStorageOptions(v StorageSystemDeploymentRequestStorageOptions) {
	o.StorageOptions = v
}

// GetIsTermsAndConditionsAgreed returns the IsTermsAndConditionsAgreed field value
func (o *StorageSystemDeploymentRequest) GetIsTermsAndConditionsAgreed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTermsAndConditionsAgreed
}

// GetIsTermsAndConditionsAgreedOk returns a tuple with the IsTermsAndConditionsAgreed field value
// and a boolean to check if the value has been set.
func (o *StorageSystemDeploymentRequest) GetIsTermsAndConditionsAgreedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTermsAndConditionsAgreed, true
}

// SetIsTermsAndConditionsAgreed sets field value
func (o *StorageSystemDeploymentRequest) SetIsTermsAndConditionsAgreed(v bool) {
	o.IsTermsAndConditionsAgreed = v
}

func (o StorageSystemDeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSystemDeploymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["cloud_options"] = o.CloudOptions
	toSerialize["storage_options"] = o.StorageOptions
	toSerialize["is_terms_and_conditions_agreed"] = o.IsTermsAndConditionsAgreed
	return toSerialize, nil
}

func (o *StorageSystemDeploymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"cloud_options",
		"storage_options",
		"is_terms_and_conditions_agreed",
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

	varStorageSystemDeploymentRequest := _StorageSystemDeploymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageSystemDeploymentRequest)

	if err != nil {
		return err
	}

	*o = StorageSystemDeploymentRequest(varStorageSystemDeploymentRequest)

	return err
}

type NullableStorageSystemDeploymentRequest struct {
	value *StorageSystemDeploymentRequest
	isSet bool
}

func (v NullableStorageSystemDeploymentRequest) Get() *StorageSystemDeploymentRequest {
	return v.value
}

func (v *NullableStorageSystemDeploymentRequest) Set(val *StorageSystemDeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSystemDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSystemDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSystemDeploymentRequest(val *StorageSystemDeploymentRequest) *NullableStorageSystemDeploymentRequest {
	return &NullableStorageSystemDeploymentRequest{value: val, isSet: true}
}

func (v NullableStorageSystemDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSystemDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



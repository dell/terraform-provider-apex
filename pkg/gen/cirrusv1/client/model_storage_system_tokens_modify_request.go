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

// checks if the StorageSystemTokensModifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageSystemTokensModifyRequest{}

// StorageSystemTokensModifyRequest Storage system tokens modify request
type StorageSystemTokensModifyRequest struct {
	// Update access token of the system.
	AccessToken string `json:"access_token"`
}

type _StorageSystemTokensModifyRequest StorageSystemTokensModifyRequest

// NewStorageSystemTokensModifyRequest instantiates a new StorageSystemTokensModifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSystemTokensModifyRequest(accessToken string) *StorageSystemTokensModifyRequest {
	this := StorageSystemTokensModifyRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewStorageSystemTokensModifyRequestWithDefaults instantiates a new StorageSystemTokensModifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSystemTokensModifyRequestWithDefaults() *StorageSystemTokensModifyRequest {
	this := StorageSystemTokensModifyRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *StorageSystemTokensModifyRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *StorageSystemTokensModifyRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *StorageSystemTokensModifyRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o StorageSystemTokensModifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSystemTokensModifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	return toSerialize, nil
}

func (o *StorageSystemTokensModifyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
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

	varStorageSystemTokensModifyRequest := _StorageSystemTokensModifyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageSystemTokensModifyRequest)

	if err != nil {
		return err
	}

	*o = StorageSystemTokensModifyRequest(varStorageSystemTokensModifyRequest)

	return err
}

type NullableStorageSystemTokensModifyRequest struct {
	value *StorageSystemTokensModifyRequest
	isSet bool
}

func (v NullableStorageSystemTokensModifyRequest) Get() *StorageSystemTokensModifyRequest {
	return v.value
}

func (v *NullableStorageSystemTokensModifyRequest) Set(val *StorageSystemTokensModifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSystemTokensModifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSystemTokensModifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSystemTokensModifyRequest(val *StorageSystemTokensModifyRequest) *NullableStorageSystemTokensModifyRequest {
	return &NullableStorageSystemTokensModifyRequest{value: val, isSet: true}
}

func (v NullableStorageSystemTokensModifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSystemTokensModifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



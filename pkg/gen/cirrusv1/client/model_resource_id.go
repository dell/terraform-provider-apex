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

// checks if the ResourceId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceId{}

// ResourceId Resource Identifier
type ResourceId struct {
	// ID of the resource
	Id string `json:"id"`
}

type _ResourceId ResourceId

// NewResourceId instantiates a new ResourceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceId(id string) *ResourceId {
	this := ResourceId{}
	this.Id = id
	return &this
}

// NewResourceIdWithDefaults instantiates a new ResourceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceIdWithDefaults() *ResourceId {
	this := ResourceId{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceId) SetId(v string) {
	o.Id = v
}

func (o ResourceId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ResourceId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varResourceId := _ResourceId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceId)

	if err != nil {
		return err
	}

	*o = ResourceId(varResourceId)

	return err
}

type NullableResourceId struct {
	value *ResourceId
	isSet bool
}

func (v NullableResourceId) Get() *ResourceId {
	return v.value
}

func (v *NullableResourceId) Set(val *ResourceId) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceId) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceId(val *ResourceId) *NullableResourceId {
	return &NullableResourceId{value: val, isSet: true}
}

func (v NullableResourceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



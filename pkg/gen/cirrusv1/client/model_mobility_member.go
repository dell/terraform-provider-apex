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

// checks if the MobilityMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobilityMember{}

// MobilityMember A mobility member is an object (e.g. volume) that is part of a mobility group that will be the source of mobility copy operations.
type MobilityMember struct {
	// Identifier of the member (e.g. PowerFlex volume identifier)
	Id string `json:"id"`
	// Name of the member (e.g. name of the volume)
	Name string `json:"name"`
	// Size of the member (e.g. volume size in bytes)
	Size string `json:"size"`
}

type _MobilityMember MobilityMember

// NewMobilityMember instantiates a new MobilityMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobilityMember(id string, name string, size string) *MobilityMember {
	this := MobilityMember{}
	this.Id = id
	this.Name = name
	this.Size = size
	return &this
}

// NewMobilityMemberWithDefaults instantiates a new MobilityMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobilityMemberWithDefaults() *MobilityMember {
	this := MobilityMember{}
	return &this
}

// GetId returns the Id field value
func (o *MobilityMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MobilityMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MobilityMember) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MobilityMember) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MobilityMember) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MobilityMember) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value
func (o *MobilityMember) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *MobilityMember) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *MobilityMember) SetSize(v string) {
	o.Size = v
}

func (o MobilityMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobilityMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *MobilityMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"size",
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

	varMobilityMember := _MobilityMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMobilityMember)

	if err != nil {
		return err
	}

	*o = MobilityMember(varMobilityMember)

	return err
}

type NullableMobilityMember struct {
	value *MobilityMember
	isSet bool
}

func (v NullableMobilityMember) Get() *MobilityMember {
	return v.value
}

func (v *NullableMobilityMember) Set(val *MobilityMember) {
	v.value = val
	v.isSet = true
}

func (v NullableMobilityMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMobilityMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobilityMember(val *MobilityMember) *NullableMobilityMember {
	return &NullableMobilityMember{value: val, isSet: true}
}

func (v NullableMobilityMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobilityMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



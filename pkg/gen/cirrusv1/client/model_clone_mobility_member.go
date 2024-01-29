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

// checks if the CloneMobilityMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneMobilityMember{}

// CloneMobilityMember A clone mobility member provides details of clone volume.
type CloneMobilityMember struct {
	// ID of the member
	Id string `json:"id"`
	// Identifier of the related mobility member
	ParentId string `json:"parent_id"`
	// Name of the member
	Name string `json:"name"`
	// Size of the member
	Size string `json:"size"`
}

type _CloneMobilityMember CloneMobilityMember

// NewCloneMobilityMember instantiates a new CloneMobilityMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneMobilityMember(id string, parentId string, name string, size string) *CloneMobilityMember {
	this := CloneMobilityMember{}
	this.Id = id
	this.ParentId = parentId
	this.Name = name
	this.Size = size
	return &this
}

// NewCloneMobilityMemberWithDefaults instantiates a new CloneMobilityMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneMobilityMemberWithDefaults() *CloneMobilityMember {
	this := CloneMobilityMember{}
	return &this
}

// GetId returns the Id field value
func (o *CloneMobilityMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloneMobilityMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloneMobilityMember) SetId(v string) {
	o.Id = v
}

// GetParentId returns the ParentId field value
func (o *CloneMobilityMember) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *CloneMobilityMember) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *CloneMobilityMember) SetParentId(v string) {
	o.ParentId = v
}

// GetName returns the Name field value
func (o *CloneMobilityMember) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneMobilityMember) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneMobilityMember) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value
func (o *CloneMobilityMember) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloneMobilityMember) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloneMobilityMember) SetSize(v string) {
	o.Size = v
}

func (o CloneMobilityMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneMobilityMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parent_id"] = o.ParentId
	toSerialize["name"] = o.Name
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *CloneMobilityMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"parent_id",
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

	varCloneMobilityMember := _CloneMobilityMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloneMobilityMember)

	if err != nil {
		return err
	}

	*o = CloneMobilityMember(varCloneMobilityMember)

	return err
}

type NullableCloneMobilityMember struct {
	value *CloneMobilityMember
	isSet bool
}

func (v NullableCloneMobilityMember) Get() *CloneMobilityMember {
	return v.value
}

func (v *NullableCloneMobilityMember) Set(val *CloneMobilityMember) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneMobilityMember) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneMobilityMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneMobilityMember(val *CloneMobilityMember) *NullableCloneMobilityMember {
	return &NullableCloneMobilityMember{value: val, isSet: true}
}

func (v NullableCloneMobilityMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneMobilityMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



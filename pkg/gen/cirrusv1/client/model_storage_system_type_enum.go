/*
APEX Navigator for Multicloud Storage REST APIs

The public API definitions for APEX Navigator for Multicloud Storage

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// StorageSystemTypeEnum The source system type (e.g.: POWERFLEX)
type StorageSystemTypeEnum string

// List of StorageSystemTypeEnum
const (
	STORAGESYSTEMTYPEENUM_POWERFLEX StorageSystemTypeEnum = "POWERFLEX"
)

// All allowed values of StorageSystemTypeEnum enum
var AllowedStorageSystemTypeEnumEnumValues = []StorageSystemTypeEnum{
	"POWERFLEX",
}

func (v *StorageSystemTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageSystemTypeEnum(value)
	for _, existing := range AllowedStorageSystemTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageSystemTypeEnum", value)
}

// NewStorageSystemTypeEnumFromValue returns a pointer to a valid StorageSystemTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageSystemTypeEnumFromValue(v string) (*StorageSystemTypeEnum, error) {
	ev := StorageSystemTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageSystemTypeEnum: valid values are %v", v, AllowedStorageSystemTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageSystemTypeEnum) IsValid() bool {
	for _, existing := range AllowedStorageSystemTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageSystemTypeEnum value
func (v StorageSystemTypeEnum) Ptr() *StorageSystemTypeEnum {
	return &v
}

type NullableStorageSystemTypeEnum struct {
	value *StorageSystemTypeEnum
	isSet bool
}

func (v NullableStorageSystemTypeEnum) Get() *StorageSystemTypeEnum {
	return v.value
}

func (v *NullableStorageSystemTypeEnum) Set(val *StorageSystemTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSystemTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSystemTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSystemTypeEnum(val *StorageSystemTypeEnum) *NullableStorageSystemTypeEnum {
	return &NullableStorageSystemTypeEnum{value: val, isSet: true}
}

func (v NullableStorageSystemTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSystemTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server


import (
	"fmt"
)


// StorageTypeEnum : This enum represents all the supported storage types * BLOCK 
type StorageTypeEnum string

// List of StorageTypeEnum
const (
	STORAGETYPEENUM_BLOCK StorageTypeEnum = "BLOCK"
)

// AllowedStorageTypeEnumEnumValues is all the allowed values of StorageTypeEnum enum
var AllowedStorageTypeEnumEnumValues = []StorageTypeEnum{
	"BLOCK",
}

// validStorageTypeEnumEnumValue provides a map of StorageTypeEnums for fast verification of use input
var validStorageTypeEnumEnumValues = map[StorageTypeEnum]struct{}{
	"BLOCK": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageTypeEnum) IsValid() bool {
	_, ok := validStorageTypeEnumEnumValues[v]
	return ok
}

// NewStorageTypeEnumFromValue returns a pointer to a valid StorageTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageTypeEnumFromValue(v string) (StorageTypeEnum, error) {
	ev := StorageTypeEnum(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for StorageTypeEnum: valid values are %v", v, AllowedStorageTypeEnumEnumValues)
	}
}



// AssertStorageTypeEnumRequired checks if the required fields are not zero-ed
func AssertStorageTypeEnumRequired(obj StorageTypeEnum) error {
	return nil
}

// AssertStorageTypeEnumConstraints checks if the values respects the defined constraints
func AssertStorageTypeEnumConstraints(obj StorageTypeEnum) error {
	return nil
}

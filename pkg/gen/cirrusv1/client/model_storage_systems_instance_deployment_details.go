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

// StorageSystemsInstanceDeploymentDetails - struct for StorageSystemsInstanceDeploymentDetails
type StorageSystemsInstanceDeploymentDetails struct {
	SystemOnPremDeploymentDetails *SystemOnPremDeploymentDetails
	SystemPublicCloudDeploymentDetails *SystemPublicCloudDeploymentDetails
}

// SystemOnPremDeploymentDetailsAsStorageSystemsInstanceDeploymentDetails is a convenience function that returns SystemOnPremDeploymentDetails wrapped in StorageSystemsInstanceDeploymentDetails
func SystemOnPremDeploymentDetailsAsStorageSystemsInstanceDeploymentDetails(v *SystemOnPremDeploymentDetails) StorageSystemsInstanceDeploymentDetails {
	return StorageSystemsInstanceDeploymentDetails{
		SystemOnPremDeploymentDetails: v,
	}
}

// SystemPublicCloudDeploymentDetailsAsStorageSystemsInstanceDeploymentDetails is a convenience function that returns SystemPublicCloudDeploymentDetails wrapped in StorageSystemsInstanceDeploymentDetails
func SystemPublicCloudDeploymentDetailsAsStorageSystemsInstanceDeploymentDetails(v *SystemPublicCloudDeploymentDetails) StorageSystemsInstanceDeploymentDetails {
	return StorageSystemsInstanceDeploymentDetails{
		SystemPublicCloudDeploymentDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageSystemsInstanceDeploymentDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SystemOnPremDeploymentDetails
	err = newStrictDecoder(data).Decode(&dst.SystemOnPremDeploymentDetails)
	if err == nil {
		jsonSystemOnPremDeploymentDetails, _ := json.Marshal(dst.SystemOnPremDeploymentDetails)
		if string(jsonSystemOnPremDeploymentDetails) == "{}" { // empty struct
			dst.SystemOnPremDeploymentDetails = nil
		} else {
			match++
		}
	} else {
		dst.SystemOnPremDeploymentDetails = nil
	}

	// try to unmarshal data into SystemPublicCloudDeploymentDetails
	err = newStrictDecoder(data).Decode(&dst.SystemPublicCloudDeploymentDetails)
	if err == nil {
		jsonSystemPublicCloudDeploymentDetails, _ := json.Marshal(dst.SystemPublicCloudDeploymentDetails)
		if string(jsonSystemPublicCloudDeploymentDetails) == "{}" { // empty struct
			dst.SystemPublicCloudDeploymentDetails = nil
		} else {
			match++
		}
	} else {
		dst.SystemPublicCloudDeploymentDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SystemOnPremDeploymentDetails = nil
		dst.SystemPublicCloudDeploymentDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StorageSystemsInstanceDeploymentDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StorageSystemsInstanceDeploymentDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageSystemsInstanceDeploymentDetails) MarshalJSON() ([]byte, error) {
	if src.SystemOnPremDeploymentDetails != nil {
		return json.Marshal(&src.SystemOnPremDeploymentDetails)
	}

	if src.SystemPublicCloudDeploymentDetails != nil {
		return json.Marshal(&src.SystemPublicCloudDeploymentDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageSystemsInstanceDeploymentDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SystemOnPremDeploymentDetails != nil {
		return obj.SystemOnPremDeploymentDetails
	}

	if obj.SystemPublicCloudDeploymentDetails != nil {
		return obj.SystemPublicCloudDeploymentDetails
	}

	// all schemas are nil
	return nil
}

type NullableStorageSystemsInstanceDeploymentDetails struct {
	value *StorageSystemsInstanceDeploymentDetails
	isSet bool
}

func (v NullableStorageSystemsInstanceDeploymentDetails) Get() *StorageSystemsInstanceDeploymentDetails {
	return v.value
}

func (v *NullableStorageSystemsInstanceDeploymentDetails) Set(val *StorageSystemsInstanceDeploymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSystemsInstanceDeploymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSystemsInstanceDeploymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSystemsInstanceDeploymentDetails(val *StorageSystemsInstanceDeploymentDetails) *NullableStorageSystemsInstanceDeploymentDetails {
	return &NullableStorageSystemsInstanceDeploymentDetails{value: val, isSet: true}
}

func (v NullableStorageSystemsInstanceDeploymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSystemsInstanceDeploymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



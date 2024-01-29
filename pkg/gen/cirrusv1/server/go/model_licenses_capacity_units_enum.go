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


// LicensesCapacityUnitsEnum : Valid values of license capacity_units attribute. * COUNT - A specific quantity limit. * PERCENT - A percentage value in the 0-100 range. * REQUEST_P_MIN - Requests per minute. * BYTE - Bytes of data. * BYTE_P_SEC- Bytes per second of throughput. * IO_P_SEC - IOs per second of throughput. 
type LicensesCapacityUnitsEnum string

// List of LicensesCapacityUnitsEnum
const (
	LICENSESCAPACITYUNITSENUM_COUNT LicensesCapacityUnitsEnum = "COUNT"
	LICENSESCAPACITYUNITSENUM_PERCENT LicensesCapacityUnitsEnum = "PERCENT"
	LICENSESCAPACITYUNITSENUM_REQUEST_P_MIN LicensesCapacityUnitsEnum = "REQUEST_P_MIN"
	LICENSESCAPACITYUNITSENUM_BYTE LicensesCapacityUnitsEnum = "BYTE"
	LICENSESCAPACITYUNITSENUM_BYTE_P_SEC LicensesCapacityUnitsEnum = "BYTE_P_SEC"
	LICENSESCAPACITYUNITSENUM_IO_P_SEC LicensesCapacityUnitsEnum = "IO_P_SEC"
)

// AllowedLicensesCapacityUnitsEnumEnumValues is all the allowed values of LicensesCapacityUnitsEnum enum
var AllowedLicensesCapacityUnitsEnumEnumValues = []LicensesCapacityUnitsEnum{
	"COUNT",
	"PERCENT",
	"REQUEST_P_MIN",
	"BYTE",
	"BYTE_P_SEC",
	"IO_P_SEC",
}

// validLicensesCapacityUnitsEnumEnumValue provides a map of LicensesCapacityUnitsEnums for fast verification of use input
var validLicensesCapacityUnitsEnumEnumValues = map[LicensesCapacityUnitsEnum]struct{}{
	"COUNT": {},
	"PERCENT": {},
	"REQUEST_P_MIN": {},
	"BYTE": {},
	"BYTE_P_SEC": {},
	"IO_P_SEC": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicensesCapacityUnitsEnum) IsValid() bool {
	_, ok := validLicensesCapacityUnitsEnumEnumValues[v]
	return ok
}

// NewLicensesCapacityUnitsEnumFromValue returns a pointer to a valid LicensesCapacityUnitsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicensesCapacityUnitsEnumFromValue(v string) (LicensesCapacityUnitsEnum, error) {
	ev := LicensesCapacityUnitsEnum(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for LicensesCapacityUnitsEnum: valid values are %v", v, AllowedLicensesCapacityUnitsEnumEnumValues)
	}
}



// AssertLicensesCapacityUnitsEnumRequired checks if the required fields are not zero-ed
func AssertLicensesCapacityUnitsEnumRequired(obj LicensesCapacityUnitsEnum) error {
	return nil
}

// AssertLicensesCapacityUnitsEnumConstraints checks if the values respects the defined constraints
func AssertLicensesCapacityUnitsEnumConstraints(obj LicensesCapacityUnitsEnum) error {
	return nil
}

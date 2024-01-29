/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// MobilityMemberMap - A mobility member map is a mapping of a mobility member and it's related member.  For example a target volume with a reference to the source volume.  Or a clone volume and its related target volume.
type MobilityMemberMap struct {

	// ID of the member
	Id string `json:"id"`

	// Identifier of the related mobility member
	ParentId string `json:"parent_id"`

	// Name of the member
	Name string `json:"name"`

	// Size of the member
	Size string `json:"size"`
}

// AssertMobilityMemberMapRequired checks if the required fields are not zero-ed
func AssertMobilityMemberMapRequired(obj MobilityMemberMap) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"parent_id": obj.ParentId,
		"name": obj.Name,
		"size": obj.Size,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMobilityMemberMapConstraints checks if the values respects the defined constraints
func AssertMobilityMemberMapConstraints(obj MobilityMemberMap) error {
	return nil
}

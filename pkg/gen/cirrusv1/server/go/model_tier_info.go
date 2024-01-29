/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// TierInfo - This model represents the tier info for a storage template
type TierInfo struct {

	Id string `json:"id"`

	// Tier name
	Name string `json:"name"`

	TierType TierEnum `json:"tier_type"`

	// Description of the tier
	Description string `json:"description"`

	// Supported storage options for a tier
	StorageOptions []TierOption `json:"storage_options"`

	// Supported cloud options for a tier
	CloudOptions []TierOption `json:"cloud_options"`
}

// AssertTierInfoRequired checks if the required fields are not zero-ed
func AssertTierInfoRequired(obj TierInfo) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"name": obj.Name,
		"tier_type": obj.TierType,
		"description": obj.Description,
		"storage_options": obj.StorageOptions,
		"cloud_options": obj.CloudOptions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.StorageOptions {
		if err := AssertTierOptionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CloudOptions {
		if err := AssertTierOptionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTierInfoConstraints checks if the values respects the defined constraints
func AssertTierInfoConstraints(obj TierInfo) error {
	return nil
}

/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// SupportMap - This model shows the cloud support mapping, all supported storage product versions, supported actions, and the latest supported version.
type SupportMap struct {

	Id string `json:"id"`

	// Evaluation period in days. After the evaluation period is expired, you need to purchase a license from Dell, to continue the use the product.
	SupportedEvaluationPeriod int32 `json:"supported_evaluation_period"`

	// Version of the storage product on the cloud
	Version string `json:"version"`

	// Part group of PowerFlex
	PartGroup string `json:"part_group,omitempty"`

	// All the supported actions for the storage products version
	SupportedActions []StorageProductActionEnum `json:"supported_actions"`
}

// AssertSupportMapRequired checks if the required fields are not zero-ed
func AssertSupportMapRequired(obj SupportMap) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"supported_evaluation_period": obj.SupportedEvaluationPeriod,
		"version": obj.Version,
		"supported_actions": obj.SupportedActions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSupportMapConstraints checks if the values respects the defined constraints
func AssertSupportMapConstraints(obj SupportMap) error {
	return nil
}

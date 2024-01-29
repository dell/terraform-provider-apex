/*
 * APEX Orchestration Platform - Job Management System (JMS) API
 *
 * Provides management and visibility for APEX Orchestration Platform Jobs
 *
 * API version: IGNORED - see resource tag's x-api-version for the specific version of this resource.
 * Contact: apex.mars@dell.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server


import (
	"time"
)



// ErrorMessage - A message describing the failure, a contributing factor to the failure, or possibly the aftermath of the failure.
type ErrorMessage struct {

	Severity SeverityEnum `json:"severity,omitempty"`

	// Identifier for this kind of message. This is a string that can be used to look up  additional information on the support website. (Note - specific format can be determined  by platform - hex value codes are common in midrange storage.) 
	Code string `json:"code,omitempty"`

	// Message string in English.
	Message string `json:"message,omitempty"`

	// The time at which the error occurred.
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Localized message string.
	MessageL10n string `json:"message_l10n,omitempty"`

	// Ordered list of substitution args for the error message. Must match up with  the {0}, {1}, etc... actually in the message referenced by the message code  above, if any. 
	Arguments []string `json:"arguments,omitempty"`
}

// AssertErrorMessageRequired checks if the required fields are not zero-ed
func AssertErrorMessageRequired(obj ErrorMessage) error {
	return nil
}

// AssertErrorMessageConstraints checks if the values respects the defined constraints
func AssertErrorMessageConstraints(obj ErrorMessage) error {
	return nil
}

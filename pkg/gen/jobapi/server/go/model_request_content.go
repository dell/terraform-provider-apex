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




// RequestContent - Generic request object. This is used as a generic type for request bodies that job instances preserve for clients to query in the future. 
type RequestContent struct {

	// This is the name of the OpenAPI definition of the specific request body in the request_body property. For example \"LDAPCreate\", \"LocalAccountModify\", etc... 
	OpenapiType string `json:"openapi_type,omitempty"`

	// The request body that invoked this job.
	Body map[string]interface{} `json:"body,omitempty"`
}

// AssertRequestContentRequired checks if the required fields are not zero-ed
func AssertRequestContentRequired(obj RequestContent) error {
	return nil
}

// AssertRequestContentConstraints checks if the values respects the defined constraints
func AssertRequestContentConstraints(obj RequestContent) error {
	return nil
}

/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// Vpc - VPC
type Vpc struct {

	// Create a new VPC. Default is false.
	IsNewVpc bool `json:"is_new_vpc,omitempty"`

	// AWS VPC resource identifier.  Provide this value if you want Apex Navigator to deploy the storage system in your existing VPC. 
	Id string `json:"id,omitempty"`

	// AWS VPC name. Provide only name (not id) if you want Apex Navigator to deploy the storage system in a new VPC. 
	Name string `json:"name,omitempty"`
}

// AssertVpcRequired checks if the required fields are not zero-ed
func AssertVpcRequired(obj Vpc) error {
	return nil
}

// AssertVpcConstraints checks if the values respects the defined constraints
func AssertVpcConstraints(obj Vpc) error {
	return nil
}

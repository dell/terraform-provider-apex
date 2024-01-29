/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// SubnetOptions - Subnet Options
type SubnetOptions struct {

	// AWS subnet resource identifier
	Id string `json:"id,omitempty"`

	// IP range that Apex Navigator will use to create new subnets and use it based on the  subnet option type.  If this is a new VPC, Apex Navigator also associates this CIDR block to the new VPC as primary or secondary  CIDR block. 
	CidrBlock string `json:"cidr_block,omitempty"`

	Type SubnetTypeEnum `json:"type,omitempty"`
}

// AssertSubnetOptionsRequired checks if the required fields are not zero-ed
func AssertSubnetOptionsRequired(obj SubnetOptions) error {
	return nil
}

// AssertSubnetOptionsConstraints checks if the values respects the defined constraints
func AssertSubnetOptionsConstraints(obj SubnetOptions) error {
	return nil
}

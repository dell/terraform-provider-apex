/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




type AwsTrustPolicyStatementInner struct {

	Effect string `json:"Effect,omitempty"`

	Action string `json:"Action,omitempty"`

	Principal AwsTrustPolicyStatementInnerPrincipal `json:"Principal,omitempty"`

	Condition AwsTrustPolicyStatementInnerCondition `json:"Condition,omitempty"`
}

// AssertAwsTrustPolicyStatementInnerRequired checks if the required fields are not zero-ed
func AssertAwsTrustPolicyStatementInnerRequired(obj AwsTrustPolicyStatementInner) error {
	if err := AssertAwsTrustPolicyStatementInnerPrincipalRequired(obj.Principal); err != nil {
		return err
	}
	if err := AssertAwsTrustPolicyStatementInnerConditionRequired(obj.Condition); err != nil {
		return err
	}
	return nil
}

// AssertAwsTrustPolicyStatementInnerConstraints checks if the values respects the defined constraints
func AssertAwsTrustPolicyStatementInnerConstraints(obj AwsTrustPolicyStatementInner) error {
	return nil
}

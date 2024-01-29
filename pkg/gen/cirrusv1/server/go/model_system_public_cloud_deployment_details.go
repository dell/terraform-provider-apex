/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// SystemPublicCloudDeploymentDetails - Public cloud deployment details.
type SystemPublicCloudDeploymentDetails struct {

	DeploymentType SystemDeploymentTypeEnum `json:"deployment_type,omitempty"`

	CloudType CloudProviderEnum `json:"cloud_type,omitempty"`

	// Cloud provider account where the storage system resides.
	CloudAccount string `json:"cloud_account,omitempty"`

	// Cloud provider region where the storage system resides.
	CloudRegion string `json:"cloud_region,omitempty"`

	AvailabilityZoneTopology AvailabilityZoneTopologyEnum `json:"availability_zone_topology,omitempty"`

	// Cloud virtual private environment identifier.
	VirtualPrivateCloud string `json:"virtual_private_cloud,omitempty"`

	// Management IPv4 or IPv6 address or DNS name of the storage system in cloud.
	CloudManagementAddress string `json:"cloud_management_address,omitempty"`

	// Minimum IOPS requested during the deployment time - Unit: IO/s
	MinimumIops int64 `json:"minimum_iops,omitempty"`

	// Minimum capacity requested during the deployment time - Unit: bytes
	MinimumCapacity int64 `json:"minimum_capacity,omitempty"`

	// Tier type requested during the deployment time.
	TierType string `json:"tier_type,omitempty"`
}

// AssertSystemPublicCloudDeploymentDetailsRequired checks if the required fields are not zero-ed
func AssertSystemPublicCloudDeploymentDetailsRequired(obj SystemPublicCloudDeploymentDetails) error {
	return nil
}

// AssertSystemPublicCloudDeploymentDetailsConstraints checks if the values respects the defined constraints
func AssertSystemPublicCloudDeploymentDetailsConstraints(obj SystemPublicCloudDeploymentDetails) error {
	return nil
}

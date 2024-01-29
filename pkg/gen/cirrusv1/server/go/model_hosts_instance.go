/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// HostsInstance - The host object.
type HostsInstance struct {

	// Host identifier.
	Id string `json:"id"`

	// Unique identifier for the system that the host is connected to.
	SystemId string `json:"system_id,omitempty"`

	// Product type of the system.
	SystemType string `json:"system_type,omitempty"`

	// Description of the host.
	Description string `json:"description,omitempty"`

	// Number of initiators that are connected between the host or server and the monitored system.
	InitiatorCount int64 `json:"initiator_count,omitempty"`

	// Type of initiator (FC or iSCSI) that the host or server uses to connect to a monitored system.
	InitiatorProtocol string `json:"initiator_protocol,omitempty"`

	// Number of health issues that are present on the host or server.
	IssueCount int64 `json:"issue_count,omitempty"`

	// Name of the host or server.
	Name string `json:"name,omitempty"`

	// Identifier of the host, defined by the system.
	NativeId string `json:"native_id,omitempty"`

	// IPv4 or IPv6 IP addresses, domain names, or netgroup name associated with the host or server.
	NetworkAddresses string `json:"network_addresses,omitempty"`

	// Type of the host.
	Type string `json:"type,omitempty"`

	// Operating system of the host or server.
	OperatingSystem string `json:"operating_system,omitempty"`

	// Model of the system.
	SystemModel string `json:"system_model,omitempty"`

	// Name of the system.
	SystemName string `json:"system_name,omitempty"`

	// Total size of all LUNs or Volumes that are provisioned to the host or server from the system - Unit: bytes
	TotalSize int64 `json:"total_size,omitempty"`
}

// AssertHostsInstanceRequired checks if the required fields are not zero-ed
func AssertHostsInstanceRequired(obj HostsInstance) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertHostsInstanceConstraints checks if the values respects the defined constraints
func AssertHostsInstanceConstraints(obj HostsInstance) error {
	return nil
}

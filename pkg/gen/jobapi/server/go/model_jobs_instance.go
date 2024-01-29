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



// JobsInstance - Information about a job.
type JobsInstance struct {

	// Unique identifier of the job.
	Id string `json:"id,omitempty"`

	// Name of the job, if not provided it will default to <resource_type>.<resource_action>.
	Name string `json:"name,omitempty"`

	Resource ResourceRef `json:"resource,omitempty"`

	ResourceAction ResourceActionEnum `json:"resource_action,omitempty"`

	// When resource_action==OTHER_ACTION, this value should be a single token that names a product  specific action. This should _only_ be used for resource types of OTHER_RESOURCE,  representing native APIs outside of the /dtapi REST namespace. 
	ResourceActionOtherName string `json:"resource_action_other_name,omitempty"`

	// Description of the job.
	DescriptionL10n string `json:"description_l10n,omitempty"`

	State JobStateEnum `json:"state,omitempty"`

	// Whether or not the pause operation is possible on this job.
	IsPausable bool `json:"is_pausable,omitempty"`

	// Whether or not the cancel operation is possible on this job.
	IsCancellable bool `json:"is_cancellable,omitempty"`

	// Date and time when the job was instantiated, whether due to an external request or other reason.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Date and time when the job execution started.
	StartTime time.Time `json:"start_time,omitempty"`

	// Date and time when the job execution completed, if succeeded, completed, or failed.
	EndTime time.Time `json:"end_time,omitempty"`

	// Date and time when the job should start.
	RequestedStartTime time.Time `json:"requested_start_time,omitempty"`

	// Date and time past which job execution should not be allowed to continue. If the job is  cancellable then it will be cancelled if it has not completed by this tme. Since not all  jobs can be cancelled, a job could run past the requested_end_time. 
	RequestedEndTime time.Time `json:"requested_end_time,omitempty"`

	// Estimated completion date and time, if not completed or failed.
	EstimatedCompletionTime time.Time `json:"estimated_completion_time,omitempty"`

	// Approximate percent complete of the job. This may be percentage relative to number of required operations and not reflect time 
	Progress int32 `json:"progress,omitempty"`

	// Unique identifier of the parent job, if applicable.
	ParentId string `json:"parent_id,omitempty"`

	// List of child jobs. This list can be _either_ a list of other job resource instances whose parent_id is this job, _or_ a list of jobs without id values, that are  queriable only as children of this job, but not a mix of both types. 
	Children []JobsInstance `json:"children,omitempty"`

	// Unique identifier of the root job, if applicable.  The root job is the job at the top of the parent hierarchy. 
	RootId string `json:"root_id,omitempty"`

	// List of jobs in this family. A job family is all the jobs sharing the same root_id. Only the root job (the one with the root_id) has this property. 
	FamilyIds []string `json:"family_ids,omitempty"`

	// Unique id of the administrative user responsible for the job.
	UserId string `json:"user_id,omitempty"`

	Request RequestContent `json:"request,omitempty"`

	// Name of the OpenAPI object definition in the request property above. Normally will be called  <ResourceType><ResourceAction>, such as JobDelete or X509CertificateGenerateCSR. 
	RequestDefinitionName string `json:"request_definition_name,omitempty"`

	Response ResponseContent `json:"response,omitempty"`

	// Name of the OpenAPI object definition in the response property above. Normally will be called  <ResourceType><ResourceAction>Response, such as X509CertificateGenerateCSRResponse for 200 status responses. Null for 204 status. ErrorResponse for 4xx/5xx responses. 
	ResponseDefinitionName string `json:"response_definition_name,omitempty"`

	ResponseStatus HttpStatusEnum `json:"response_status,omitempty"`

	// IP address from which the job request originated, if external to the system.
	ClientAddress string `json:"client_address,omitempty"`

	// Start order of a given job step with respect to its siblings within the job hierarchy. Even jobs run concurrently will start in some specific order. 
	StepOrder int32 `json:"step_order,omitempty"`

	// Indicates whether a command is internal or visible to the end user.
	IsInternal bool `json:"is_internal,omitempty"`

	PauseMessage ErrorMessage `json:"pause_message,omitempty"`

	// Opaque value that may be set by request originator. It is persisted with the job and may be used for:  * Filtering a job query to find an async job (whether intentional or an unintentional disconnect).  * To retry a request. If a request is made with a duplicate correlationId value, then the system must    return the result of the initial invocation with that correlationId, waiting for that initial job to    finish if necessary. This does not apply when a completed or failed jobs has already expired and    been deleted. * The platform _should_ log this value anywhere log entries occur during the execution of the request. 
	CorrelationId string `json:"correlation_id,omitempty"`
}

// AssertJobsInstanceRequired checks if the required fields are not zero-ed
func AssertJobsInstanceRequired(obj JobsInstance) error {
	if err := AssertResourceRefRequired(obj.Resource); err != nil {
		return err
	}
	for _, el := range obj.Children {
		if err := AssertJobsInstanceRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRequestContentRequired(obj.Request); err != nil {
		return err
	}
	if err := AssertResponseContentRequired(obj.Response); err != nil {
		return err
	}
	if err := AssertErrorMessageRequired(obj.PauseMessage); err != nil {
		return err
	}
	return nil
}

// AssertJobsInstanceConstraints checks if the values respects the defined constraints
func AssertJobsInstanceConstraints(obj JobsInstance) error {
	return nil
}

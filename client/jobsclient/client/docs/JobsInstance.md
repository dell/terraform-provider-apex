# JobsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the job. | [optional] 
**Name** | Pointer to **string** | Name of the job, if not provided it will default to &lt;resource_type&gt;.&lt;resource_action&gt;. | [optional] 
**Resource** | Pointer to [**ResourceRef**](ResourceRef.md) |  | [optional] 
**ResourceAction** | Pointer to [**ResourceActionEnum**](ResourceActionEnum.md) |  | [optional] 
**ResourceActionOtherName** | Pointer to **string** | When resource_action&#x3D;&#x3D;OTHER_ACTION, this value should be a single token that names a product  specific action. This should _only_ be used for resource types of OTHER_RESOURCE,  representing native APIs outside of the /dtapi REST namespace.  | [optional] 
**DescriptionL10n** | Pointer to **string** | Description of the job. | [optional] 
**State** | Pointer to [**JobStateEnum**](JobStateEnum.md) |  | [optional] 
**IsPausable** | Pointer to **bool** | Whether or not the pause operation is possible on this job. | [optional] 
**IsCancellable** | Pointer to **bool** | Whether or not the cancel operation is possible on this job. | [optional] 
**CreateTime** | Pointer to **time.Time** | Date and time when the job was instantiated, whether due to an external request or other reason. | [optional] 
**StartTime** | Pointer to **time.Time** | Date and time when the job execution started. | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the job execution completed, if succeeded, completed, or failed. | [optional] 
**RequestedStartTime** | Pointer to **time.Time** | Date and time when the job should start. | [optional] 
**RequestedEndTime** | Pointer to **time.Time** | Date and time past which job execution should not be allowed to continue. If the job is  cancellable then it will be cancelled if it has not completed by this tme. Since not all  jobs can be cancelled, a job could run past the requested_end_time.  | [optional] 
**EstimatedCompletionTime** | Pointer to **time.Time** | Estimated completion date and time, if not completed or failed. | [optional] 
**Progress** | Pointer to **int32** | Approximate percent complete of the job. This may be percentage relative to number of required operations and not reflect time  | [optional] 
**ParentId** | Pointer to **string** | Unique identifier of the parent job, if applicable. | [optional] 
**Children** | Pointer to [**[]JobsInstance**](JobsInstance.md) | List of child jobs. This list can be _either_ a list of other job resource instances whose parent_id is this job, _or_ a list of jobs without id values, that are  queriable only as children of this job, but not a mix of both types.  | [optional] 
**RootId** | Pointer to **string** | Unique identifier of the root job, if applicable.  The root job is the job at the top of the parent hierarchy.  | [optional] 
**FamilyIds** | Pointer to **[]string** | List of jobs in this family. A job family is all the jobs sharing the same root_id. Only the root job (the one with the root_id) has this property.  | [optional] 
**UserId** | Pointer to **string** | Unique id of the administrative user responsible for the job. | [optional] 
**Request** | Pointer to [**RequestContent**](RequestContent.md) |  | [optional] 
**RequestDefinitionName** | Pointer to **string** | Name of the OpenAPI object definition in the request property above. Normally will be called  &lt;ResourceType&gt;&lt;ResourceAction&gt;, such as JobDelete or X509CertificateGenerateCSR.  | [optional] 
**Response** | Pointer to [**ResponseContent**](ResponseContent.md) |  | [optional] 
**ResponseDefinitionName** | Pointer to **string** | Name of the OpenAPI object definition in the response property above. Normally will be called  &lt;ResourceType&gt;&lt;ResourceAction&gt;Response, such as X509CertificateGenerateCSRResponse for 200 status responses. Null for 204 status. ErrorResponse for 4xx/5xx responses.  | [optional] 
**ResponseStatus** | Pointer to [**HTTPStatusEnum**](HTTPStatusEnum.md) |  | [optional] 
**ClientAddress** | Pointer to **string** | IP address from which the job request originated, if external to the system. | [optional] 
**StepOrder** | Pointer to **int32** | Start order of a given job step with respect to its siblings within the job hierarchy. Even jobs run concurrently will start in some specific order.  | [optional] 
**IsInternal** | Pointer to **bool** | Indicates whether a command is internal or visible to the end user. | [optional] 
**PauseMessage** | Pointer to [**ErrorMessage**](ErrorMessage.md) |  | [optional] 
**CorrelationId** | Pointer to **string** | Opaque value that may be set by request originator. It is persisted with the job and may be used for:  * Filtering a job query to find an async job (whether intentional or an unintentional disconnect).  * To retry a request. If a request is made with a duplicate correlationId value, then the system must    return the result of the initial invocation with that correlationId, waiting for that initial job to    finish if necessary. This does not apply when a completed or failed jobs has already expired and    been deleted. * The platform _should_ log this value anywhere log entries occur during the execution of the request.  | [optional] 

## Methods

### NewJobsInstance

`func NewJobsInstance() *JobsInstance`

NewJobsInstance instantiates a new JobsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsInstanceWithDefaults

`func NewJobsInstanceWithDefaults() *JobsInstance`

NewJobsInstanceWithDefaults instantiates a new JobsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobsInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobsInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JobsInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobsInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobsInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobsInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResource

`func (o *JobsInstance) GetResource() ResourceRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *JobsInstance) GetResourceOk() (*ResourceRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *JobsInstance) SetResource(v ResourceRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *JobsInstance) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceAction

`func (o *JobsInstance) GetResourceAction() ResourceActionEnum`

GetResourceAction returns the ResourceAction field if non-nil, zero value otherwise.

### GetResourceActionOk

`func (o *JobsInstance) GetResourceActionOk() (*ResourceActionEnum, bool)`

GetResourceActionOk returns a tuple with the ResourceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAction

`func (o *JobsInstance) SetResourceAction(v ResourceActionEnum)`

SetResourceAction sets ResourceAction field to given value.

### HasResourceAction

`func (o *JobsInstance) HasResourceAction() bool`

HasResourceAction returns a boolean if a field has been set.

### GetResourceActionOtherName

`func (o *JobsInstance) GetResourceActionOtherName() string`

GetResourceActionOtherName returns the ResourceActionOtherName field if non-nil, zero value otherwise.

### GetResourceActionOtherNameOk

`func (o *JobsInstance) GetResourceActionOtherNameOk() (*string, bool)`

GetResourceActionOtherNameOk returns a tuple with the ResourceActionOtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceActionOtherName

`func (o *JobsInstance) SetResourceActionOtherName(v string)`

SetResourceActionOtherName sets ResourceActionOtherName field to given value.

### HasResourceActionOtherName

`func (o *JobsInstance) HasResourceActionOtherName() bool`

HasResourceActionOtherName returns a boolean if a field has been set.

### GetDescriptionL10n

`func (o *JobsInstance) GetDescriptionL10n() string`

GetDescriptionL10n returns the DescriptionL10n field if non-nil, zero value otherwise.

### GetDescriptionL10nOk

`func (o *JobsInstance) GetDescriptionL10nOk() (*string, bool)`

GetDescriptionL10nOk returns a tuple with the DescriptionL10n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionL10n

`func (o *JobsInstance) SetDescriptionL10n(v string)`

SetDescriptionL10n sets DescriptionL10n field to given value.

### HasDescriptionL10n

`func (o *JobsInstance) HasDescriptionL10n() bool`

HasDescriptionL10n returns a boolean if a field has been set.

### GetState

`func (o *JobsInstance) GetState() JobStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobsInstance) GetStateOk() (*JobStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobsInstance) SetState(v JobStateEnum)`

SetState sets State field to given value.

### HasState

`func (o *JobsInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsPausable

`func (o *JobsInstance) GetIsPausable() bool`

GetIsPausable returns the IsPausable field if non-nil, zero value otherwise.

### GetIsPausableOk

`func (o *JobsInstance) GetIsPausableOk() (*bool, bool)`

GetIsPausableOk returns a tuple with the IsPausable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausable

`func (o *JobsInstance) SetIsPausable(v bool)`

SetIsPausable sets IsPausable field to given value.

### HasIsPausable

`func (o *JobsInstance) HasIsPausable() bool`

HasIsPausable returns a boolean if a field has been set.

### GetIsCancellable

`func (o *JobsInstance) GetIsCancellable() bool`

GetIsCancellable returns the IsCancellable field if non-nil, zero value otherwise.

### GetIsCancellableOk

`func (o *JobsInstance) GetIsCancellableOk() (*bool, bool)`

GetIsCancellableOk returns a tuple with the IsCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellable

`func (o *JobsInstance) SetIsCancellable(v bool)`

SetIsCancellable sets IsCancellable field to given value.

### HasIsCancellable

`func (o *JobsInstance) HasIsCancellable() bool`

HasIsCancellable returns a boolean if a field has been set.

### GetCreateTime

`func (o *JobsInstance) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *JobsInstance) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *JobsInstance) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *JobsInstance) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetStartTime

`func (o *JobsInstance) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobsInstance) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobsInstance) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobsInstance) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *JobsInstance) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobsInstance) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobsInstance) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobsInstance) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRequestedStartTime

`func (o *JobsInstance) GetRequestedStartTime() time.Time`

GetRequestedStartTime returns the RequestedStartTime field if non-nil, zero value otherwise.

### GetRequestedStartTimeOk

`func (o *JobsInstance) GetRequestedStartTimeOk() (*time.Time, bool)`

GetRequestedStartTimeOk returns a tuple with the RequestedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStartTime

`func (o *JobsInstance) SetRequestedStartTime(v time.Time)`

SetRequestedStartTime sets RequestedStartTime field to given value.

### HasRequestedStartTime

`func (o *JobsInstance) HasRequestedStartTime() bool`

HasRequestedStartTime returns a boolean if a field has been set.

### GetRequestedEndTime

`func (o *JobsInstance) GetRequestedEndTime() time.Time`

GetRequestedEndTime returns the RequestedEndTime field if non-nil, zero value otherwise.

### GetRequestedEndTimeOk

`func (o *JobsInstance) GetRequestedEndTimeOk() (*time.Time, bool)`

GetRequestedEndTimeOk returns a tuple with the RequestedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEndTime

`func (o *JobsInstance) SetRequestedEndTime(v time.Time)`

SetRequestedEndTime sets RequestedEndTime field to given value.

### HasRequestedEndTime

`func (o *JobsInstance) HasRequestedEndTime() bool`

HasRequestedEndTime returns a boolean if a field has been set.

### GetEstimatedCompletionTime

`func (o *JobsInstance) GetEstimatedCompletionTime() time.Time`

GetEstimatedCompletionTime returns the EstimatedCompletionTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeOk

`func (o *JobsInstance) GetEstimatedCompletionTimeOk() (*time.Time, bool)`

GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTime

`func (o *JobsInstance) SetEstimatedCompletionTime(v time.Time)`

SetEstimatedCompletionTime sets EstimatedCompletionTime field to given value.

### HasEstimatedCompletionTime

`func (o *JobsInstance) HasEstimatedCompletionTime() bool`

HasEstimatedCompletionTime returns a boolean if a field has been set.

### GetProgress

`func (o *JobsInstance) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *JobsInstance) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *JobsInstance) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *JobsInstance) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetParentId

`func (o *JobsInstance) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *JobsInstance) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *JobsInstance) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *JobsInstance) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildren

`func (o *JobsInstance) GetChildren() []JobsInstance`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *JobsInstance) GetChildrenOk() (*[]JobsInstance, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *JobsInstance) SetChildren(v []JobsInstance)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *JobsInstance) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRootId

`func (o *JobsInstance) GetRootId() string`

GetRootId returns the RootId field if non-nil, zero value otherwise.

### GetRootIdOk

`func (o *JobsInstance) GetRootIdOk() (*string, bool)`

GetRootIdOk returns a tuple with the RootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootId

`func (o *JobsInstance) SetRootId(v string)`

SetRootId sets RootId field to given value.

### HasRootId

`func (o *JobsInstance) HasRootId() bool`

HasRootId returns a boolean if a field has been set.

### GetFamilyIds

`func (o *JobsInstance) GetFamilyIds() []string`

GetFamilyIds returns the FamilyIds field if non-nil, zero value otherwise.

### GetFamilyIdsOk

`func (o *JobsInstance) GetFamilyIdsOk() (*[]string, bool)`

GetFamilyIdsOk returns a tuple with the FamilyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyIds

`func (o *JobsInstance) SetFamilyIds(v []string)`

SetFamilyIds sets FamilyIds field to given value.

### HasFamilyIds

`func (o *JobsInstance) HasFamilyIds() bool`

HasFamilyIds returns a boolean if a field has been set.

### GetUserId

`func (o *JobsInstance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JobsInstance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JobsInstance) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JobsInstance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRequest

`func (o *JobsInstance) GetRequest() RequestContent`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *JobsInstance) GetRequestOk() (*RequestContent, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *JobsInstance) SetRequest(v RequestContent)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *JobsInstance) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRequestDefinitionName

`func (o *JobsInstance) GetRequestDefinitionName() string`

GetRequestDefinitionName returns the RequestDefinitionName field if non-nil, zero value otherwise.

### GetRequestDefinitionNameOk

`func (o *JobsInstance) GetRequestDefinitionNameOk() (*string, bool)`

GetRequestDefinitionNameOk returns a tuple with the RequestDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDefinitionName

`func (o *JobsInstance) SetRequestDefinitionName(v string)`

SetRequestDefinitionName sets RequestDefinitionName field to given value.

### HasRequestDefinitionName

`func (o *JobsInstance) HasRequestDefinitionName() bool`

HasRequestDefinitionName returns a boolean if a field has been set.

### GetResponse

`func (o *JobsInstance) GetResponse() ResponseContent`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *JobsInstance) GetResponseOk() (*ResponseContent, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *JobsInstance) SetResponse(v ResponseContent)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *JobsInstance) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseDefinitionName

`func (o *JobsInstance) GetResponseDefinitionName() string`

GetResponseDefinitionName returns the ResponseDefinitionName field if non-nil, zero value otherwise.

### GetResponseDefinitionNameOk

`func (o *JobsInstance) GetResponseDefinitionNameOk() (*string, bool)`

GetResponseDefinitionNameOk returns a tuple with the ResponseDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDefinitionName

`func (o *JobsInstance) SetResponseDefinitionName(v string)`

SetResponseDefinitionName sets ResponseDefinitionName field to given value.

### HasResponseDefinitionName

`func (o *JobsInstance) HasResponseDefinitionName() bool`

HasResponseDefinitionName returns a boolean if a field has been set.

### GetResponseStatus

`func (o *JobsInstance) GetResponseStatus() HTTPStatusEnum`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *JobsInstance) GetResponseStatusOk() (*HTTPStatusEnum, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *JobsInstance) SetResponseStatus(v HTTPStatusEnum)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *JobsInstance) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetClientAddress

`func (o *JobsInstance) GetClientAddress() string`

GetClientAddress returns the ClientAddress field if non-nil, zero value otherwise.

### GetClientAddressOk

`func (o *JobsInstance) GetClientAddressOk() (*string, bool)`

GetClientAddressOk returns a tuple with the ClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAddress

`func (o *JobsInstance) SetClientAddress(v string)`

SetClientAddress sets ClientAddress field to given value.

### HasClientAddress

`func (o *JobsInstance) HasClientAddress() bool`

HasClientAddress returns a boolean if a field has been set.

### GetStepOrder

`func (o *JobsInstance) GetStepOrder() int32`

GetStepOrder returns the StepOrder field if non-nil, zero value otherwise.

### GetStepOrderOk

`func (o *JobsInstance) GetStepOrderOk() (*int32, bool)`

GetStepOrderOk returns a tuple with the StepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepOrder

`func (o *JobsInstance) SetStepOrder(v int32)`

SetStepOrder sets StepOrder field to given value.

### HasStepOrder

`func (o *JobsInstance) HasStepOrder() bool`

HasStepOrder returns a boolean if a field has been set.

### GetIsInternal

`func (o *JobsInstance) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *JobsInstance) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *JobsInstance) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *JobsInstance) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### GetPauseMessage

`func (o *JobsInstance) GetPauseMessage() ErrorMessage`

GetPauseMessage returns the PauseMessage field if non-nil, zero value otherwise.

### GetPauseMessageOk

`func (o *JobsInstance) GetPauseMessageOk() (*ErrorMessage, bool)`

GetPauseMessageOk returns a tuple with the PauseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseMessage

`func (o *JobsInstance) SetPauseMessage(v ErrorMessage)`

SetPauseMessage sets PauseMessage field to given value.

### HasPauseMessage

`func (o *JobsInstance) HasPauseMessage() bool`

HasPauseMessage returns a boolean if a field has been set.

### GetCorrelationId

`func (o *JobsInstance) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *JobsInstance) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *JobsInstance) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *JobsInstance) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



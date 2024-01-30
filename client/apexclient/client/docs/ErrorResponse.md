# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpStatusCode** | Pointer to [**HttpStatusEnum**](HttpStatusEnum.md) |  | [optional] 
**Messages** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) | A list of messages describing the failure encountered by this request. At least  one will be of Error severity because Info and Warning conditions do not cause  the request to fail.  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatusCode

`func (o *ErrorResponse) GetHttpStatusCode() HttpStatusEnum`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *ErrorResponse) GetHttpStatusCodeOk() (*HttpStatusEnum, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *ErrorResponse) SetHttpStatusCode(v HttpStatusEnum)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *ErrorResponse) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorResponse) GetMessages() []ErrorMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorResponse) GetMessagesOk() (*[]ErrorMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorResponse) SetMessages(v []ErrorMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



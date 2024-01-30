# ErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to [**SeverityEnum**](SeverityEnum.md) |  | [optional] 
**Code** | Pointer to **string** | Identifier for this kind of message. This is a string that can be used to look up  additional information on the support website. (Note - specific format can be determined  by platform - hex value codes are common in midrange storage.)  | [optional] 
**Message** | Pointer to **string** | Message string in English. | [optional] 
**Timestamp** | Pointer to **time.Time** | The time at which the error occurred. | [optional] 
**MessageL10n** | Pointer to **string** | Localized message string. | [optional] 
**Arguments** | Pointer to **[]string** | Ordered list of substitution args for the error message. Must match up with  the {0}, {1}, etc... actually in the message referenced by the message code  above, if any.  | [optional] 

## Methods

### NewErrorMessage

`func NewErrorMessage() *ErrorMessage`

NewErrorMessage instantiates a new ErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageWithDefaults

`func NewErrorMessageWithDefaults() *ErrorMessage`

NewErrorMessageWithDefaults instantiates a new ErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *ErrorMessage) GetSeverity() SeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ErrorMessage) GetSeverityOk() (*SeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ErrorMessage) SetSeverity(v SeverityEnum)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ErrorMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCode

`func (o *ErrorMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorMessage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorMessage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *ErrorMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ErrorMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ErrorMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ErrorMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessageL10n

`func (o *ErrorMessage) GetMessageL10n() string`

GetMessageL10n returns the MessageL10n field if non-nil, zero value otherwise.

### GetMessageL10nOk

`func (o *ErrorMessage) GetMessageL10nOk() (*string, bool)`

GetMessageL10nOk returns a tuple with the MessageL10n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageL10n

`func (o *ErrorMessage) SetMessageL10n(v string)`

SetMessageL10n sets MessageL10n field to given value.

### HasMessageL10n

`func (o *ErrorMessage) HasMessageL10n() bool`

HasMessageL10n returns a boolean if a field has been set.

### GetArguments

`func (o *ErrorMessage) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ErrorMessage) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ErrorMessage) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ErrorMessage) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



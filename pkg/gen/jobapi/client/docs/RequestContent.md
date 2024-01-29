# RequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenapiType** | Pointer to **string** | This is the name of the OpenAPI definition of the specific request body in the request_body property. For example \&quot;LDAPCreate\&quot;, \&quot;LocalAccountModify\&quot;, etc...  | [optional] 
**Body** | Pointer to **map[string]interface{}** | The request body that invoked this job. | [optional] 

## Methods

### NewRequestContent

`func NewRequestContent() *RequestContent`

NewRequestContent instantiates a new RequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestContentWithDefaults

`func NewRequestContentWithDefaults() *RequestContent`

NewRequestContentWithDefaults instantiates a new RequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapiType

`func (o *RequestContent) GetOpenapiType() string`

GetOpenapiType returns the OpenapiType field if non-nil, zero value otherwise.

### GetOpenapiTypeOk

`func (o *RequestContent) GetOpenapiTypeOk() (*string, bool)`

GetOpenapiTypeOk returns a tuple with the OpenapiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapiType

`func (o *RequestContent) SetOpenapiType(v string)`

SetOpenapiType sets OpenapiType field to given value.

### HasOpenapiType

`func (o *RequestContent) HasOpenapiType() bool`

HasOpenapiType returns a boolean if a field has been set.

### GetBody

`func (o *RequestContent) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RequestContent) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RequestContent) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *RequestContent) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



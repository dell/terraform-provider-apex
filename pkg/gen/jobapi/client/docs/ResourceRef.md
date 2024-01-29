# ResourceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ResourceTypeEnum**](ResourceTypeEnum.md) |  | 
**Id** | **string** | Unique identifier of the resource on which the job is operating. | 
**Name** | Pointer to **string** | The name of the referenced resource, if applicable. This is copied from the resource as a convenience for clients. | [optional] 
**ResourceUrl** | Pointer to **string** | The URL of the referenced resource instance. For specific resource types in ResourceType enum, the URL will be of the form &lt;base_uri&gt;/&lt;type&gt;/&lt;id&gt;. When the type is OTHER_RESOURCE, the value can be any valid URI that resolves to the referenced resource. | [optional] 

## Methods

### NewResourceRef

`func NewResourceRef(type_ ResourceTypeEnum, id string, ) *ResourceRef`

NewResourceRef instantiates a new ResourceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRefWithDefaults

`func NewResourceRefWithDefaults() *ResourceRef`

NewResourceRefWithDefaults instantiates a new ResourceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceRef) GetType() ResourceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceRef) GetTypeOk() (*ResourceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceRef) SetType(v ResourceTypeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ResourceRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceRef) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResourceRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceUrl

`func (o *ResourceRef) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *ResourceRef) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *ResourceRef) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *ResourceRef) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



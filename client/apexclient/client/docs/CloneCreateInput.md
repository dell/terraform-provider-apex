# CloneCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of this clone group | 
**Description** | Pointer to **string** | Description of this clone group | [optional] 
**MobilityTargetId** | **string** | ID of the mobility target that was used to create the clone | 
**HostMappings** | Pointer to **[]string** | IDs of hosts that this clone should be exposed to | [optional] 

## Methods

### NewCloneCreateInput

`func NewCloneCreateInput(name string, mobilityTargetId string, ) *CloneCreateInput`

NewCloneCreateInput instantiates a new CloneCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneCreateInputWithDefaults

`func NewCloneCreateInputWithDefaults() *CloneCreateInput`

NewCloneCreateInputWithDefaults instantiates a new CloneCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CloneCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloneCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloneCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloneCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMobilityTargetId

`func (o *CloneCreateInput) GetMobilityTargetId() string`

GetMobilityTargetId returns the MobilityTargetId field if non-nil, zero value otherwise.

### GetMobilityTargetIdOk

`func (o *CloneCreateInput) GetMobilityTargetIdOk() (*string, bool)`

GetMobilityTargetIdOk returns a tuple with the MobilityTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilityTargetId

`func (o *CloneCreateInput) SetMobilityTargetId(v string)`

SetMobilityTargetId sets MobilityTargetId field to given value.


### GetHostMappings

`func (o *CloneCreateInput) GetHostMappings() []string`

GetHostMappings returns the HostMappings field if non-nil, zero value otherwise.

### GetHostMappingsOk

`func (o *CloneCreateInput) GetHostMappingsOk() (*[]string, bool)`

GetHostMappingsOk returns a tuple with the HostMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMappings

`func (o *CloneCreateInput) SetHostMappings(v []string)`

SetHostMappings sets HostMappings field to given value.

### HasHostMappings

`func (o *CloneCreateInput) HasHostMappings() bool`

HasHostMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



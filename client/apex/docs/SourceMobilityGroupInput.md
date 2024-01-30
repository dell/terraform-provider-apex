# SourceMobilityGroupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the mobility group | 
**Description** | Pointer to **string** | Description of the mobility group | [optional] 
**SystemId** | **string** | Identifier of the source system for this mobility group | 
**SystemType** | [**StorageSystemTypeEnum**](StorageSystemTypeEnum.md) |  | 
**Members** | **[]string** | List of the identifiers of the mobility group members  (e.g. volumes identifiers) | 

## Methods

### NewSourceMobilityGroupInput

`func NewSourceMobilityGroupInput(name string, systemId string, systemType StorageSystemTypeEnum, members []string, ) *SourceMobilityGroupInput`

NewSourceMobilityGroupInput instantiates a new SourceMobilityGroupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceMobilityGroupInputWithDefaults

`func NewSourceMobilityGroupInputWithDefaults() *SourceMobilityGroupInput`

NewSourceMobilityGroupInputWithDefaults instantiates a new SourceMobilityGroupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceMobilityGroupInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceMobilityGroupInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceMobilityGroupInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SourceMobilityGroupInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceMobilityGroupInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceMobilityGroupInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceMobilityGroupInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSystemId

`func (o *SourceMobilityGroupInput) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SourceMobilityGroupInput) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SourceMobilityGroupInput) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetSystemType

`func (o *SourceMobilityGroupInput) GetSystemType() StorageSystemTypeEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *SourceMobilityGroupInput) GetSystemTypeOk() (*StorageSystemTypeEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *SourceMobilityGroupInput) SetSystemType(v StorageSystemTypeEnum)`

SetSystemType sets SystemType field to given value.


### GetMembers

`func (o *SourceMobilityGroupInput) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SourceMobilityGroupInput) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SourceMobilityGroupInput) SetMembers(v []string)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



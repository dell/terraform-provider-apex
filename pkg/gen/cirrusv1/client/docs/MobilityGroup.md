# MobilityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mobility group identifier | 
**Name** | **string** | Mobility group name | 
**Description** | Pointer to **string** | Mobility group description | [optional] 
**SystemId** | **string** | Identifier of the system for the mobility group members | 
**SystemType** | [**StorageSystemTypeEnum**](StorageSystemTypeEnum.md) |  | 
**CreationTimestamp** | **string** | When the mobility group was created | 
**Members** | [**[]MobilityMember**](MobilityMember.md) | A list of members (e.g. volumes) for source groups that the user wants to copy.   | 

## Methods

### NewMobilityGroup

`func NewMobilityGroup(id string, name string, systemId string, systemType StorageSystemTypeEnum, creationTimestamp string, members []MobilityMember, ) *MobilityGroup`

NewMobilityGroup instantiates a new MobilityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobilityGroupWithDefaults

`func NewMobilityGroupWithDefaults() *MobilityGroup`

NewMobilityGroupWithDefaults instantiates a new MobilityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobilityGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobilityGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobilityGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MobilityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobilityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobilityGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MobilityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobilityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobilityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MobilityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSystemId

`func (o *MobilityGroup) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *MobilityGroup) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *MobilityGroup) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetSystemType

`func (o *MobilityGroup) GetSystemType() StorageSystemTypeEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *MobilityGroup) GetSystemTypeOk() (*StorageSystemTypeEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *MobilityGroup) SetSystemType(v StorageSystemTypeEnum)`

SetSystemType sets SystemType field to given value.


### GetCreationTimestamp

`func (o *MobilityGroup) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *MobilityGroup) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *MobilityGroup) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetMembers

`func (o *MobilityGroup) GetMembers() []MobilityMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MobilityGroup) GetMembersOk() (*[]MobilityMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MobilityGroup) SetMembers(v []MobilityMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



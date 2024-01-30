# MobilityTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Idenifier of this target mobility group | 
**Name** | **string** | Name of the mobility target | 
**Description** | **string** | Description of the mobility target | 
**SystemType** | Pointer to [**StorageSystemTypeEnum**](StorageSystemTypeEnum.md) |  | [optional] 
**SystemId** | **string** | ID of the target system | 
**SourceMobilityGroupId** | **string** | ID of the source mobility group | 
**CreationTimestamp** | **string** | Timestamp from when the group was created | 
**ImageTimestamp** | Pointer to **string** | Timestamp of the last source image copied to this target | [optional] 
**LastCopyJobId** | Pointer to **string** |  | [optional] 
**TargetMembers** | Pointer to [**[]MobilityMemberMap**](MobilityMemberMap.md) |  | [optional] 
**BandwidthLimit** | Pointer to **int32** | Bandwidth limit in Mbps (Mega bits per second). | [optional] 

## Methods

### NewMobilityTarget

`func NewMobilityTarget(id string, name string, description string, systemId string, sourceMobilityGroupId string, creationTimestamp string, ) *MobilityTarget`

NewMobilityTarget instantiates a new MobilityTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobilityTargetWithDefaults

`func NewMobilityTargetWithDefaults() *MobilityTarget`

NewMobilityTargetWithDefaults instantiates a new MobilityTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobilityTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobilityTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobilityTarget) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MobilityTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobilityTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobilityTarget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MobilityTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobilityTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobilityTarget) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSystemType

`func (o *MobilityTarget) GetSystemType() StorageSystemTypeEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *MobilityTarget) GetSystemTypeOk() (*StorageSystemTypeEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *MobilityTarget) SetSystemType(v StorageSystemTypeEnum)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *MobilityTarget) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetSystemId

`func (o *MobilityTarget) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *MobilityTarget) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *MobilityTarget) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetSourceMobilityGroupId

`func (o *MobilityTarget) GetSourceMobilityGroupId() string`

GetSourceMobilityGroupId returns the SourceMobilityGroupId field if non-nil, zero value otherwise.

### GetSourceMobilityGroupIdOk

`func (o *MobilityTarget) GetSourceMobilityGroupIdOk() (*string, bool)`

GetSourceMobilityGroupIdOk returns a tuple with the SourceMobilityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMobilityGroupId

`func (o *MobilityTarget) SetSourceMobilityGroupId(v string)`

SetSourceMobilityGroupId sets SourceMobilityGroupId field to given value.


### GetCreationTimestamp

`func (o *MobilityTarget) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *MobilityTarget) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *MobilityTarget) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetImageTimestamp

`func (o *MobilityTarget) GetImageTimestamp() string`

GetImageTimestamp returns the ImageTimestamp field if non-nil, zero value otherwise.

### GetImageTimestampOk

`func (o *MobilityTarget) GetImageTimestampOk() (*string, bool)`

GetImageTimestampOk returns a tuple with the ImageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTimestamp

`func (o *MobilityTarget) SetImageTimestamp(v string)`

SetImageTimestamp sets ImageTimestamp field to given value.

### HasImageTimestamp

`func (o *MobilityTarget) HasImageTimestamp() bool`

HasImageTimestamp returns a boolean if a field has been set.

### GetLastCopyJobId

`func (o *MobilityTarget) GetLastCopyJobId() string`

GetLastCopyJobId returns the LastCopyJobId field if non-nil, zero value otherwise.

### GetLastCopyJobIdOk

`func (o *MobilityTarget) GetLastCopyJobIdOk() (*string, bool)`

GetLastCopyJobIdOk returns a tuple with the LastCopyJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCopyJobId

`func (o *MobilityTarget) SetLastCopyJobId(v string)`

SetLastCopyJobId sets LastCopyJobId field to given value.

### HasLastCopyJobId

`func (o *MobilityTarget) HasLastCopyJobId() bool`

HasLastCopyJobId returns a boolean if a field has been set.

### GetTargetMembers

`func (o *MobilityTarget) GetTargetMembers() []MobilityMemberMap`

GetTargetMembers returns the TargetMembers field if non-nil, zero value otherwise.

### GetTargetMembersOk

`func (o *MobilityTarget) GetTargetMembersOk() (*[]MobilityMemberMap, bool)`

GetTargetMembersOk returns a tuple with the TargetMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMembers

`func (o *MobilityTarget) SetTargetMembers(v []MobilityMemberMap)`

SetTargetMembers sets TargetMembers field to given value.

### HasTargetMembers

`func (o *MobilityTarget) HasTargetMembers() bool`

HasTargetMembers returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *MobilityTarget) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *MobilityTarget) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *MobilityTarget) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *MobilityTarget) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



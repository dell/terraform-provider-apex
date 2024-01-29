# Clone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  Identifier of the cloned object. | 
**Name** | **string** | Name of the clone. | 
**Description** | **string** | Description of the clone. | 
**MobilityTargetId** | Pointer to **string** |  | [optional] 
**CreationTimestamp** | **string** | When the clone was created | 
**RefreshTimestamp** | **string** | When the clone was last updated | 
**ImageTimestamp** | **string** |  | 
**CloneVolumes** | [**[]CloneMobilityMember**](CloneMobilityMember.md) | A list of storage object identifiers representing the copied instance on the destination/target for which a clone is created.  | 
**HostMappings** | [**[]CloneToHostMapping**](CloneToHostMapping.md) | A list of the hosts/SDCs mapped to the clone | 

## Methods

### NewClone

`func NewClone(id string, name string, description string, creationTimestamp string, refreshTimestamp string, imageTimestamp string, cloneVolumes []CloneMobilityMember, hostMappings []CloneToHostMapping, ) *Clone`

NewClone instantiates a new Clone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneWithDefaults

`func NewCloneWithDefaults() *Clone`

NewCloneWithDefaults instantiates a new Clone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Clone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Clone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Clone) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Clone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Clone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Clone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Clone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Clone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Clone) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMobilityTargetId

`func (o *Clone) GetMobilityTargetId() string`

GetMobilityTargetId returns the MobilityTargetId field if non-nil, zero value otherwise.

### GetMobilityTargetIdOk

`func (o *Clone) GetMobilityTargetIdOk() (*string, bool)`

GetMobilityTargetIdOk returns a tuple with the MobilityTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilityTargetId

`func (o *Clone) SetMobilityTargetId(v string)`

SetMobilityTargetId sets MobilityTargetId field to given value.

### HasMobilityTargetId

`func (o *Clone) HasMobilityTargetId() bool`

HasMobilityTargetId returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *Clone) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Clone) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Clone) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetRefreshTimestamp

`func (o *Clone) GetRefreshTimestamp() string`

GetRefreshTimestamp returns the RefreshTimestamp field if non-nil, zero value otherwise.

### GetRefreshTimestampOk

`func (o *Clone) GetRefreshTimestampOk() (*string, bool)`

GetRefreshTimestampOk returns a tuple with the RefreshTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTimestamp

`func (o *Clone) SetRefreshTimestamp(v string)`

SetRefreshTimestamp sets RefreshTimestamp field to given value.


### GetImageTimestamp

`func (o *Clone) GetImageTimestamp() string`

GetImageTimestamp returns the ImageTimestamp field if non-nil, zero value otherwise.

### GetImageTimestampOk

`func (o *Clone) GetImageTimestampOk() (*string, bool)`

GetImageTimestampOk returns a tuple with the ImageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTimestamp

`func (o *Clone) SetImageTimestamp(v string)`

SetImageTimestamp sets ImageTimestamp field to given value.


### GetCloneVolumes

`func (o *Clone) GetCloneVolumes() []CloneMobilityMember`

GetCloneVolumes returns the CloneVolumes field if non-nil, zero value otherwise.

### GetCloneVolumesOk

`func (o *Clone) GetCloneVolumesOk() (*[]CloneMobilityMember, bool)`

GetCloneVolumesOk returns a tuple with the CloneVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneVolumes

`func (o *Clone) SetCloneVolumes(v []CloneMobilityMember)`

SetCloneVolumes sets CloneVolumes field to given value.


### GetHostMappings

`func (o *Clone) GetHostMappings() []CloneToHostMapping`

GetHostMappings returns the HostMappings field if non-nil, zero value otherwise.

### GetHostMappingsOk

`func (o *Clone) GetHostMappingsOk() (*[]CloneToHostMapping, bool)`

GetHostMappingsOk returns a tuple with the HostMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMappings

`func (o *Clone) SetHostMappings(v []CloneToHostMapping)`

SetHostMappings sets HostMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateTargetBaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the mobility target | 
**Description** | Pointer to **string** | Description of the mobility target | [optional] 
**SystemId** | **string** | Identifier of the target system | 
**SystemType** | [**StorageSystemTypeEnum**](StorageSystemTypeEnum.md) |  | 
**TargetSystemOptions** | [**TargetSystemOptions**](TargetSystemOptions.md) |  | 
**BandwidthLimit** | Pointer to **int32** | Bandwidth limit in Mbps (Mega bits per second). | [optional] 

## Methods

### NewCreateTargetBaseInput

`func NewCreateTargetBaseInput(name string, systemId string, systemType StorageSystemTypeEnum, targetSystemOptions TargetSystemOptions, ) *CreateTargetBaseInput`

NewCreateTargetBaseInput instantiates a new CreateTargetBaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTargetBaseInputWithDefaults

`func NewCreateTargetBaseInputWithDefaults() *CreateTargetBaseInput`

NewCreateTargetBaseInputWithDefaults instantiates a new CreateTargetBaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTargetBaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTargetBaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTargetBaseInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTargetBaseInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTargetBaseInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTargetBaseInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTargetBaseInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSystemId

`func (o *CreateTargetBaseInput) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *CreateTargetBaseInput) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *CreateTargetBaseInput) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetSystemType

`func (o *CreateTargetBaseInput) GetSystemType() StorageSystemTypeEnum`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *CreateTargetBaseInput) GetSystemTypeOk() (*StorageSystemTypeEnum, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *CreateTargetBaseInput) SetSystemType(v StorageSystemTypeEnum)`

SetSystemType sets SystemType field to given value.


### GetTargetSystemOptions

`func (o *CreateTargetBaseInput) GetTargetSystemOptions() TargetSystemOptions`

GetTargetSystemOptions returns the TargetSystemOptions field if non-nil, zero value otherwise.

### GetTargetSystemOptionsOk

`func (o *CreateTargetBaseInput) GetTargetSystemOptionsOk() (*TargetSystemOptions, bool)`

GetTargetSystemOptionsOk returns a tuple with the TargetSystemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystemOptions

`func (o *CreateTargetBaseInput) SetTargetSystemOptions(v TargetSystemOptions)`

SetTargetSystemOptions sets TargetSystemOptions field to given value.


### GetBandwidthLimit

`func (o *CreateTargetBaseInput) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *CreateTargetBaseInput) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *CreateTargetBaseInput) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *CreateTargetBaseInput) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateMobilityTargetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the mobility target | [optional] 
**Description** | Pointer to **string** | New description for the mobility target | [optional] 
**BandwidthLimit** | Pointer to **int32** | Bandwidth limit in Mbps (Mega bits per second). | [optional] 

## Methods

### NewUpdateMobilityTargetInput

`func NewUpdateMobilityTargetInput() *UpdateMobilityTargetInput`

NewUpdateMobilityTargetInput instantiates a new UpdateMobilityTargetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMobilityTargetInputWithDefaults

`func NewUpdateMobilityTargetInputWithDefaults() *UpdateMobilityTargetInput`

NewUpdateMobilityTargetInputWithDefaults instantiates a new UpdateMobilityTargetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMobilityTargetInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMobilityTargetInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMobilityTargetInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMobilityTargetInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMobilityTargetInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMobilityTargetInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMobilityTargetInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMobilityTargetInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *UpdateMobilityTargetInput) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *UpdateMobilityTargetInput) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *UpdateMobilityTargetInput) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *UpdateMobilityTargetInput) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



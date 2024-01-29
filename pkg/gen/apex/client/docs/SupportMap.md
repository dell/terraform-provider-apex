# SupportMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**SupportedEvaluationPeriod** | **int32** | Evaluation period in days. After the evaluation period is expired, you need to purchase a license from Dell, to continue the use the product. | 
**Version** | **string** | Version of the storage product on the cloud | 
**PartGroup** | Pointer to **string** | Part group of PowerFlex | [optional] 
**SupportedActions** | [**[]StorageProductActionEnum**](StorageProductActionEnum.md) | All the supported actions for the storage products version | 

## Methods

### NewSupportMap

`func NewSupportMap(id string, supportedEvaluationPeriod int32, version string, supportedActions []StorageProductActionEnum, ) *SupportMap`

NewSupportMap instantiates a new SupportMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportMapWithDefaults

`func NewSupportMapWithDefaults() *SupportMap`

NewSupportMapWithDefaults instantiates a new SupportMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportMap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportMap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportMap) SetId(v string)`

SetId sets Id field to given value.


### GetSupportedEvaluationPeriod

`func (o *SupportMap) GetSupportedEvaluationPeriod() int32`

GetSupportedEvaluationPeriod returns the SupportedEvaluationPeriod field if non-nil, zero value otherwise.

### GetSupportedEvaluationPeriodOk

`func (o *SupportMap) GetSupportedEvaluationPeriodOk() (*int32, bool)`

GetSupportedEvaluationPeriodOk returns a tuple with the SupportedEvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEvaluationPeriod

`func (o *SupportMap) SetSupportedEvaluationPeriod(v int32)`

SetSupportedEvaluationPeriod sets SupportedEvaluationPeriod field to given value.


### GetVersion

`func (o *SupportMap) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SupportMap) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SupportMap) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPartGroup

`func (o *SupportMap) GetPartGroup() string`

GetPartGroup returns the PartGroup field if non-nil, zero value otherwise.

### GetPartGroupOk

`func (o *SupportMap) GetPartGroupOk() (*string, bool)`

GetPartGroupOk returns a tuple with the PartGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartGroup

`func (o *SupportMap) SetPartGroup(v string)`

SetPartGroup sets PartGroup field to given value.

### HasPartGroup

`func (o *SupportMap) HasPartGroup() bool`

HasPartGroup returns a boolean if a field has been set.

### GetSupportedActions

`func (o *SupportMap) GetSupportedActions() []StorageProductActionEnum`

GetSupportedActions returns the SupportedActions field if non-nil, zero value otherwise.

### GetSupportedActionsOk

`func (o *SupportMap) GetSupportedActionsOk() (*[]StorageProductActionEnum, bool)`

GetSupportedActionsOk returns a tuple with the SupportedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedActions

`func (o *SupportMap) SetSupportedActions(v []StorageProductActionEnum)`

SetSupportedActions sets SupportedActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



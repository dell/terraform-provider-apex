# PoolsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the pool. | 
**SystemId** | Pointer to **string** | Unique identifier for the device or appliance. | [optional] 
**SystemType** | Pointer to **string** | Type of the system. | [optional] 
**FreeSize** | Pointer to **int64** | Available capacity - Unit: bytes | [optional] 
**IssueCount** | Pointer to **int64** | Number of health issues that are present on the pool. | [optional] 
**Name** | Pointer to **string** | Name of the pool. | [optional] 
**NativeId** | Pointer to **string** | Native identifier of the pool, defined by the system. | [optional] 
**SubscribedPercent** | Pointer to **float64** | Percentage of pool capacity that is provisioned. | [optional] 
**SubscribedSize** | Pointer to **int64** | Total subscribed capacity of the pool - Unit: bytes | [optional] 
**SystemModel** | Pointer to **string** | Model of the system. | [optional] 
**SystemName** | Pointer to **string** | Name of the system. | [optional] 
**TimeToFullPrediction** | Pointer to **string** | This is an enumerated type showing a prediction of when the pool may become full. Possible values are: DAY (imminent); FULL (pool is full); WEEK (full in a week); MONTH (full in a month); QUARTER (full within a quarter); BEYOND (more than a quarter to become full); LEARNING (not enough data to perform an analysis); UNPREDICTABLE (missing or invalid data); or UNKNOWN (system error). | [optional] 
**TotalSize** | Pointer to **int64** | Total capacity of the pool - Unit: bytes | [optional] 
**Type** | Pointer to **string** | The type of pool. | [optional] 
**UsedPercent** | Pointer to **float64** | Percentage of pool capacity that is being used. | [optional] 
**UsedSize** | Pointer to **int64** | Capacity of the pool that is being used - Unit: bytes | [optional] 

## Methods

### NewPoolsInstance

`func NewPoolsInstance(id string, ) *PoolsInstance`

NewPoolsInstance instantiates a new PoolsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsInstanceWithDefaults

`func NewPoolsInstanceWithDefaults() *PoolsInstance`

NewPoolsInstanceWithDefaults instantiates a new PoolsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolsInstance) SetId(v string)`

SetId sets Id field to given value.


### GetSystemId

`func (o *PoolsInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *PoolsInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *PoolsInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *PoolsInstance) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetSystemType

`func (o *PoolsInstance) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *PoolsInstance) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *PoolsInstance) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *PoolsInstance) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetFreeSize

`func (o *PoolsInstance) GetFreeSize() int64`

GetFreeSize returns the FreeSize field if non-nil, zero value otherwise.

### GetFreeSizeOk

`func (o *PoolsInstance) GetFreeSizeOk() (*int64, bool)`

GetFreeSizeOk returns a tuple with the FreeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSize

`func (o *PoolsInstance) SetFreeSize(v int64)`

SetFreeSize sets FreeSize field to given value.

### HasFreeSize

`func (o *PoolsInstance) HasFreeSize() bool`

HasFreeSize returns a boolean if a field has been set.

### GetIssueCount

`func (o *PoolsInstance) GetIssueCount() int64`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *PoolsInstance) GetIssueCountOk() (*int64, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *PoolsInstance) SetIssueCount(v int64)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *PoolsInstance) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetName

`func (o *PoolsInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolsInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolsInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolsInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeId

`func (o *PoolsInstance) GetNativeId() string`

GetNativeId returns the NativeId field if non-nil, zero value otherwise.

### GetNativeIdOk

`func (o *PoolsInstance) GetNativeIdOk() (*string, bool)`

GetNativeIdOk returns a tuple with the NativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeId

`func (o *PoolsInstance) SetNativeId(v string)`

SetNativeId sets NativeId field to given value.

### HasNativeId

`func (o *PoolsInstance) HasNativeId() bool`

HasNativeId returns a boolean if a field has been set.

### GetSubscribedPercent

`func (o *PoolsInstance) GetSubscribedPercent() float64`

GetSubscribedPercent returns the SubscribedPercent field if non-nil, zero value otherwise.

### GetSubscribedPercentOk

`func (o *PoolsInstance) GetSubscribedPercentOk() (*float64, bool)`

GetSubscribedPercentOk returns a tuple with the SubscribedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedPercent

`func (o *PoolsInstance) SetSubscribedPercent(v float64)`

SetSubscribedPercent sets SubscribedPercent field to given value.

### HasSubscribedPercent

`func (o *PoolsInstance) HasSubscribedPercent() bool`

HasSubscribedPercent returns a boolean if a field has been set.

### GetSubscribedSize

`func (o *PoolsInstance) GetSubscribedSize() int64`

GetSubscribedSize returns the SubscribedSize field if non-nil, zero value otherwise.

### GetSubscribedSizeOk

`func (o *PoolsInstance) GetSubscribedSizeOk() (*int64, bool)`

GetSubscribedSizeOk returns a tuple with the SubscribedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedSize

`func (o *PoolsInstance) SetSubscribedSize(v int64)`

SetSubscribedSize sets SubscribedSize field to given value.

### HasSubscribedSize

`func (o *PoolsInstance) HasSubscribedSize() bool`

HasSubscribedSize returns a boolean if a field has been set.

### GetSystemModel

`func (o *PoolsInstance) GetSystemModel() string`

GetSystemModel returns the SystemModel field if non-nil, zero value otherwise.

### GetSystemModelOk

`func (o *PoolsInstance) GetSystemModelOk() (*string, bool)`

GetSystemModelOk returns a tuple with the SystemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModel

`func (o *PoolsInstance) SetSystemModel(v string)`

SetSystemModel sets SystemModel field to given value.

### HasSystemModel

`func (o *PoolsInstance) HasSystemModel() bool`

HasSystemModel returns a boolean if a field has been set.

### GetSystemName

`func (o *PoolsInstance) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *PoolsInstance) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *PoolsInstance) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *PoolsInstance) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTimeToFullPrediction

`func (o *PoolsInstance) GetTimeToFullPrediction() string`

GetTimeToFullPrediction returns the TimeToFullPrediction field if non-nil, zero value otherwise.

### GetTimeToFullPredictionOk

`func (o *PoolsInstance) GetTimeToFullPredictionOk() (*string, bool)`

GetTimeToFullPredictionOk returns a tuple with the TimeToFullPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFullPrediction

`func (o *PoolsInstance) SetTimeToFullPrediction(v string)`

SetTimeToFullPrediction sets TimeToFullPrediction field to given value.

### HasTimeToFullPrediction

`func (o *PoolsInstance) HasTimeToFullPrediction() bool`

HasTimeToFullPrediction returns a boolean if a field has been set.

### GetTotalSize

`func (o *PoolsInstance) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *PoolsInstance) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *PoolsInstance) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *PoolsInstance) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetType

`func (o *PoolsInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolsInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolsInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolsInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsedPercent

`func (o *PoolsInstance) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *PoolsInstance) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *PoolsInstance) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *PoolsInstance) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetUsedSize

`func (o *PoolsInstance) GetUsedSize() int64`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *PoolsInstance) GetUsedSizeOk() (*int64, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *PoolsInstance) SetUsedSize(v int64)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *PoolsInstance) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



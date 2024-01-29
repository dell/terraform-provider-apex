# VolumesInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the volume. | 
**SystemId** | Pointer to **string** | Unique identifier for the device or appliance. | [optional] 
**SystemType** | Pointer to **string** | Type of the system for the volume. | [optional] 
**AllocatedSize** | Pointer to **int64** | The allocated size of the volume - Unit: bytes | [optional] 
**Bandwidth** | Pointer to **int64** | The bandwidth consumed by the volume. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s | [optional] 
**ConsistencyGroupName** | Pointer to **string** | Consistency group name of the volume. | [optional] 
**DataReductionPercent** | Pointer to **float64** | The data reduction percent for the volume. | [optional] 
**DataReductionRatio** | Pointer to **float64** | The data reduction ratio for the volume. | [optional] 
**DataReductionSavedSize** | Pointer to **int64** | The data reduction capacity saved for the volume - Unit: bytes | [optional] 
**IoLimitPolicyName** | Pointer to **string** | The IO limit policy name for the volume. | [optional] 
**Iops** | Pointer to **int64** | The IOPS for the volume. Aggregated for a rolling average over the last 24 hours - Unit: IO/s | [optional] 
**IsCompressedOrDeduped** | Pointer to **string** | Identifies whether the volume is compressed or deduplicated. | [optional] 
**IsThinEnabled** | Pointer to **bool** | Identifies whether the volume has thin provisioning enabled. | [optional] 
**IssueCount** | Pointer to **int64** | Number of health issues that are present on the volume. | [optional] 
**Latency** | Pointer to **int64** | The latency for the volume. Aggregated for a rolling average over the last 24 hours - Unit: microseconds | [optional] 
**LogicalSize** | Pointer to **int64** | The logical size for the volume - Unit: bytes | [optional] 
**Name** | Pointer to **string** | The name of the volume. | [optional] 
**NativeId** | Pointer to **string** | Identifier of the volume, defined by the system. | [optional] 
**Type** | Pointer to **string** | Type of the volume, which is either LUN or VOLUME. | [optional] 
**PoolId** | Pointer to **string** | The pool identifier for the volume. | [optional] 
**PoolName** | Pointer to **string** | The pool name for the volume. | [optional] 
**PoolType** | Pointer to **string** | Type of the pool. | [optional] 
**SnapshotCount** | Pointer to **int64** | The snapshot count for the volume. | [optional] 
**SnapshotPolicy** | Pointer to **string** | The snapshot policy for the volume. | [optional] 
**SnapshotSize** | Pointer to **int64** | The snapshot size for the volume - Unit: bytes | [optional] 
**StorageResourceId** | Pointer to **string** | The storage resource identifier for the volume. | [optional] 
**StorageResourceNativeId** | Pointer to **string** | The storage resource native identifier for the volume. | [optional] 
**SystemModel** | Pointer to **string** | The model of the system. | [optional] 
**SystemName** | Pointer to **string** | Name of the system for the volume. | [optional] 
**TotalSize** | Pointer to **int64** | The total provisioned size of the volume - Unit: bytes | [optional] 
**UsedSize** | Pointer to **int64** | The used size of the volume - Unit: bytes | [optional] 
**UsedSizeUnique** | Pointer to **int64** | The unique used size of the volume - Unit: bytes | [optional] 

## Methods

### NewVolumesInstance

`func NewVolumesInstance(id string, ) *VolumesInstance`

NewVolumesInstance instantiates a new VolumesInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesInstanceWithDefaults

`func NewVolumesInstanceWithDefaults() *VolumesInstance`

NewVolumesInstanceWithDefaults instantiates a new VolumesInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumesInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumesInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumesInstance) SetId(v string)`

SetId sets Id field to given value.


### GetSystemId

`func (o *VolumesInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *VolumesInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *VolumesInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *VolumesInstance) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetSystemType

`func (o *VolumesInstance) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *VolumesInstance) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *VolumesInstance) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *VolumesInstance) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *VolumesInstance) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *VolumesInstance) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *VolumesInstance) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *VolumesInstance) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetBandwidth

`func (o *VolumesInstance) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *VolumesInstance) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *VolumesInstance) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *VolumesInstance) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetConsistencyGroupName

`func (o *VolumesInstance) GetConsistencyGroupName() string`

GetConsistencyGroupName returns the ConsistencyGroupName field if non-nil, zero value otherwise.

### GetConsistencyGroupNameOk

`func (o *VolumesInstance) GetConsistencyGroupNameOk() (*string, bool)`

GetConsistencyGroupNameOk returns a tuple with the ConsistencyGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyGroupName

`func (o *VolumesInstance) SetConsistencyGroupName(v string)`

SetConsistencyGroupName sets ConsistencyGroupName field to given value.

### HasConsistencyGroupName

`func (o *VolumesInstance) HasConsistencyGroupName() bool`

HasConsistencyGroupName returns a boolean if a field has been set.

### GetDataReductionPercent

`func (o *VolumesInstance) GetDataReductionPercent() float64`

GetDataReductionPercent returns the DataReductionPercent field if non-nil, zero value otherwise.

### GetDataReductionPercentOk

`func (o *VolumesInstance) GetDataReductionPercentOk() (*float64, bool)`

GetDataReductionPercentOk returns a tuple with the DataReductionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionPercent

`func (o *VolumesInstance) SetDataReductionPercent(v float64)`

SetDataReductionPercent sets DataReductionPercent field to given value.

### HasDataReductionPercent

`func (o *VolumesInstance) HasDataReductionPercent() bool`

HasDataReductionPercent returns a boolean if a field has been set.

### GetDataReductionRatio

`func (o *VolumesInstance) GetDataReductionRatio() float64`

GetDataReductionRatio returns the DataReductionRatio field if non-nil, zero value otherwise.

### GetDataReductionRatioOk

`func (o *VolumesInstance) GetDataReductionRatioOk() (*float64, bool)`

GetDataReductionRatioOk returns a tuple with the DataReductionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionRatio

`func (o *VolumesInstance) SetDataReductionRatio(v float64)`

SetDataReductionRatio sets DataReductionRatio field to given value.

### HasDataReductionRatio

`func (o *VolumesInstance) HasDataReductionRatio() bool`

HasDataReductionRatio returns a boolean if a field has been set.

### GetDataReductionSavedSize

`func (o *VolumesInstance) GetDataReductionSavedSize() int64`

GetDataReductionSavedSize returns the DataReductionSavedSize field if non-nil, zero value otherwise.

### GetDataReductionSavedSizeOk

`func (o *VolumesInstance) GetDataReductionSavedSizeOk() (*int64, bool)`

GetDataReductionSavedSizeOk returns a tuple with the DataReductionSavedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionSavedSize

`func (o *VolumesInstance) SetDataReductionSavedSize(v int64)`

SetDataReductionSavedSize sets DataReductionSavedSize field to given value.

### HasDataReductionSavedSize

`func (o *VolumesInstance) HasDataReductionSavedSize() bool`

HasDataReductionSavedSize returns a boolean if a field has been set.

### GetIoLimitPolicyName

`func (o *VolumesInstance) GetIoLimitPolicyName() string`

GetIoLimitPolicyName returns the IoLimitPolicyName field if non-nil, zero value otherwise.

### GetIoLimitPolicyNameOk

`func (o *VolumesInstance) GetIoLimitPolicyNameOk() (*string, bool)`

GetIoLimitPolicyNameOk returns a tuple with the IoLimitPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicyName

`func (o *VolumesInstance) SetIoLimitPolicyName(v string)`

SetIoLimitPolicyName sets IoLimitPolicyName field to given value.

### HasIoLimitPolicyName

`func (o *VolumesInstance) HasIoLimitPolicyName() bool`

HasIoLimitPolicyName returns a boolean if a field has been set.

### GetIops

`func (o *VolumesInstance) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *VolumesInstance) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *VolumesInstance) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *VolumesInstance) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetIsCompressedOrDeduped

`func (o *VolumesInstance) GetIsCompressedOrDeduped() string`

GetIsCompressedOrDeduped returns the IsCompressedOrDeduped field if non-nil, zero value otherwise.

### GetIsCompressedOrDedupedOk

`func (o *VolumesInstance) GetIsCompressedOrDedupedOk() (*string, bool)`

GetIsCompressedOrDedupedOk returns a tuple with the IsCompressedOrDeduped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompressedOrDeduped

`func (o *VolumesInstance) SetIsCompressedOrDeduped(v string)`

SetIsCompressedOrDeduped sets IsCompressedOrDeduped field to given value.

### HasIsCompressedOrDeduped

`func (o *VolumesInstance) HasIsCompressedOrDeduped() bool`

HasIsCompressedOrDeduped returns a boolean if a field has been set.

### GetIsThinEnabled

`func (o *VolumesInstance) GetIsThinEnabled() bool`

GetIsThinEnabled returns the IsThinEnabled field if non-nil, zero value otherwise.

### GetIsThinEnabledOk

`func (o *VolumesInstance) GetIsThinEnabledOk() (*bool, bool)`

GetIsThinEnabledOk returns a tuple with the IsThinEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThinEnabled

`func (o *VolumesInstance) SetIsThinEnabled(v bool)`

SetIsThinEnabled sets IsThinEnabled field to given value.

### HasIsThinEnabled

`func (o *VolumesInstance) HasIsThinEnabled() bool`

HasIsThinEnabled returns a boolean if a field has been set.

### GetIssueCount

`func (o *VolumesInstance) GetIssueCount() int64`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *VolumesInstance) GetIssueCountOk() (*int64, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *VolumesInstance) SetIssueCount(v int64)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *VolumesInstance) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetLatency

`func (o *VolumesInstance) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *VolumesInstance) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *VolumesInstance) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *VolumesInstance) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLogicalSize

`func (o *VolumesInstance) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *VolumesInstance) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *VolumesInstance) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.

### HasLogicalSize

`func (o *VolumesInstance) HasLogicalSize() bool`

HasLogicalSize returns a boolean if a field has been set.

### GetName

`func (o *VolumesInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumesInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumesInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumesInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeId

`func (o *VolumesInstance) GetNativeId() string`

GetNativeId returns the NativeId field if non-nil, zero value otherwise.

### GetNativeIdOk

`func (o *VolumesInstance) GetNativeIdOk() (*string, bool)`

GetNativeIdOk returns a tuple with the NativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeId

`func (o *VolumesInstance) SetNativeId(v string)`

SetNativeId sets NativeId field to given value.

### HasNativeId

`func (o *VolumesInstance) HasNativeId() bool`

HasNativeId returns a boolean if a field has been set.

### GetType

`func (o *VolumesInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VolumesInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VolumesInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VolumesInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPoolId

`func (o *VolumesInstance) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *VolumesInstance) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *VolumesInstance) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *VolumesInstance) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *VolumesInstance) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *VolumesInstance) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *VolumesInstance) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *VolumesInstance) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolType

`func (o *VolumesInstance) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *VolumesInstance) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *VolumesInstance) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *VolumesInstance) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetSnapshotCount

`func (o *VolumesInstance) GetSnapshotCount() int64`

GetSnapshotCount returns the SnapshotCount field if non-nil, zero value otherwise.

### GetSnapshotCountOk

`func (o *VolumesInstance) GetSnapshotCountOk() (*int64, bool)`

GetSnapshotCountOk returns a tuple with the SnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCount

`func (o *VolumesInstance) SetSnapshotCount(v int64)`

SetSnapshotCount sets SnapshotCount field to given value.

### HasSnapshotCount

`func (o *VolumesInstance) HasSnapshotCount() bool`

HasSnapshotCount returns a boolean if a field has been set.

### GetSnapshotPolicy

`func (o *VolumesInstance) GetSnapshotPolicy() string`

GetSnapshotPolicy returns the SnapshotPolicy field if non-nil, zero value otherwise.

### GetSnapshotPolicyOk

`func (o *VolumesInstance) GetSnapshotPolicyOk() (*string, bool)`

GetSnapshotPolicyOk returns a tuple with the SnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicy

`func (o *VolumesInstance) SetSnapshotPolicy(v string)`

SetSnapshotPolicy sets SnapshotPolicy field to given value.

### HasSnapshotPolicy

`func (o *VolumesInstance) HasSnapshotPolicy() bool`

HasSnapshotPolicy returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *VolumesInstance) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *VolumesInstance) GetSnapshotSizeOk() (*int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *VolumesInstance) SetSnapshotSize(v int64)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *VolumesInstance) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetStorageResourceId

`func (o *VolumesInstance) GetStorageResourceId() string`

GetStorageResourceId returns the StorageResourceId field if non-nil, zero value otherwise.

### GetStorageResourceIdOk

`func (o *VolumesInstance) GetStorageResourceIdOk() (*string, bool)`

GetStorageResourceIdOk returns a tuple with the StorageResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceId

`func (o *VolumesInstance) SetStorageResourceId(v string)`

SetStorageResourceId sets StorageResourceId field to given value.

### HasStorageResourceId

`func (o *VolumesInstance) HasStorageResourceId() bool`

HasStorageResourceId returns a boolean if a field has been set.

### GetStorageResourceNativeId

`func (o *VolumesInstance) GetStorageResourceNativeId() string`

GetStorageResourceNativeId returns the StorageResourceNativeId field if non-nil, zero value otherwise.

### GetStorageResourceNativeIdOk

`func (o *VolumesInstance) GetStorageResourceNativeIdOk() (*string, bool)`

GetStorageResourceNativeIdOk returns a tuple with the StorageResourceNativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceNativeId

`func (o *VolumesInstance) SetStorageResourceNativeId(v string)`

SetStorageResourceNativeId sets StorageResourceNativeId field to given value.

### HasStorageResourceNativeId

`func (o *VolumesInstance) HasStorageResourceNativeId() bool`

HasStorageResourceNativeId returns a boolean if a field has been set.

### GetSystemModel

`func (o *VolumesInstance) GetSystemModel() string`

GetSystemModel returns the SystemModel field if non-nil, zero value otherwise.

### GetSystemModelOk

`func (o *VolumesInstance) GetSystemModelOk() (*string, bool)`

GetSystemModelOk returns a tuple with the SystemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModel

`func (o *VolumesInstance) SetSystemModel(v string)`

SetSystemModel sets SystemModel field to given value.

### HasSystemModel

`func (o *VolumesInstance) HasSystemModel() bool`

HasSystemModel returns a boolean if a field has been set.

### GetSystemName

`func (o *VolumesInstance) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *VolumesInstance) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *VolumesInstance) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *VolumesInstance) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTotalSize

`func (o *VolumesInstance) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *VolumesInstance) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *VolumesInstance) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *VolumesInstance) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetUsedSize

`func (o *VolumesInstance) GetUsedSize() int64`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *VolumesInstance) GetUsedSizeOk() (*int64, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *VolumesInstance) SetUsedSize(v int64)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *VolumesInstance) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.

### GetUsedSizeUnique

`func (o *VolumesInstance) GetUsedSizeUnique() int64`

GetUsedSizeUnique returns the UsedSizeUnique field if non-nil, zero value otherwise.

### GetUsedSizeUniqueOk

`func (o *VolumesInstance) GetUsedSizeUniqueOk() (*int64, bool)`

GetUsedSizeUniqueOk returns a tuple with the UsedSizeUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSizeUnique

`func (o *VolumesInstance) SetUsedSizeUnique(v int64)`

SetUsedSizeUnique sets UsedSizeUnique field to given value.

### HasUsedSizeUnique

`func (o *VolumesInstance) HasUsedSizeUnique() bool`

HasUsedSizeUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



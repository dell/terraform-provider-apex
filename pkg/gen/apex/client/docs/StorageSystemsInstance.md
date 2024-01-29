# StorageSystemsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the storage system. | 
**SystemId** | Pointer to **string** | Unique identifier for the device or appliance. | [optional] 
**SystemType** | Pointer to **string** | Type of the system. | [optional] 
**Bandwidth** | Pointer to **int64** | The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s | [optional] 
**CapacityImpact** | Pointer to **int64** | Impact point of highest impacting issue in the capacity health category. | [optional] 
**CapacityIssueCount** | Pointer to **int64** | Total number of issues in the capacity health category. | [optional] 
**CompressionSavings** | Pointer to **float64** | Storage efficiency ratio of data which has compression applied to it on the system. | [optional] 
**ConfigurationImpact** | Pointer to **int64** | Impact point of highest impacting issue in the configuration health category. | [optional] 
**ConfigurationIssueCount** | Pointer to **int64** | Total number of issues in the configuration health category. | [optional] 
**ConfiguredSize** | Pointer to **int64** | The configured size for this system - Unit: bytes | [optional] 
**ConnectivityStatus** | Pointer to **string** | Connectivity status. | [optional] 
**LicenseType** | Pointer to **string** | Type of the license on the system. | [optional] 
**LicenseExpirationDateTimestamp** | Pointer to **time.Time** | Expiration date for the license on the system. | [optional] 
**ContractCoverageType** | Pointer to **string** | Type of the service contract of the system. | [optional] 
**ContractExpirationDateTimestamp** | Pointer to **time.Time** | Expiration date for the service contract of the system. | [optional] 
**DataProtectionImpact** | Pointer to **int64** | Impact point of highest impacting issue in the data protection health category. | [optional] 
**DataProtectionIssueCount** | Pointer to **int64** | Total number of issues in the data protection health category. | [optional] 
**DisplayIdentifier** | Pointer to **string** | Unique identifier for the system. | [optional] 
**FreePercent** | Pointer to **float64** | The %free capacity. | [optional] 
**FreeSize** | Pointer to **int64** | The free size value - Unit: bytes | [optional] 
**HealthConnectivityStatus** | Pointer to **string** | Health connectivity status. | [optional] 
**HealthIssueCount** | Pointer to **int64** | Total amount of health issues. | [optional] 
**HealthScore** | Pointer to **int64** | The overall health score of the system. | [optional] 
**HealthState** | Pointer to **string** | Health state of the system. | [optional] 
**Iops** | Pointer to **int64** | The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s | [optional] 
**Ipv4Address** | Pointer to **string** | IPv4 address of the system. | [optional] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the system. | [optional] 
**LastContactTimestamp** | Pointer to **time.Time** | Last time that CloudIQ received data from the system. | [optional] 
**Latency** | Pointer to **int64** | The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds | [optional] 
**LogicalSize** | Pointer to **int64** | The logical size written - Unit: bytes | [optional] 
**Model** | Pointer to **string** | The model of the system. | [optional] 
**Name** | Pointer to **string** | The user-defined name of the system. | [optional] 
**OverallEfficiency** | Pointer to **float64** | The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction. | [optional] 
**PerformanceImpact** | Pointer to **int64** | Impact point of highest impacting issue in the performance health category. | [optional] 
**PerformanceIssueCount** | Pointer to **int64** | Total number of issues in the performance health category. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number for this system. | [optional] 
**SiteName** | Pointer to **string** | Name of the site where the system is registered to. | [optional] 
**SnapsSavings** | Pointer to **float64** | The snaps savings for this system. | [optional] 
**SystemHealthImpact** | Pointer to **int64** | Health impact for the system. | [optional] 
**SystemHealthIssueCount** | Pointer to **int64** | Total amount of health issues for the system. | [optional] 
**ThinSavings** | Pointer to **float64** | The savings due to thin provisioning. | [optional] 
**TotalSize** | Pointer to **int64** | The total size of the system - Unit: bytes | [optional] 
**UnconfiguredSize** | Pointer to **int64** | The unconfigured capacity for this system - Unit: bytes | [optional] 
**UsedPercent** | Pointer to **float64** | Percentage of capacity used for this system. | [optional] 
**UsedSize** | Pointer to **int64** | The value of used capacity for this system - Unit: bytes | [optional] 
**Vendor** | Pointer to **string** | Name of the vendor who makes the system. | [optional] 
**Version** | Pointer to **string** | Version identifier. | [optional] 
**DeploymentDetails** | Pointer to [**StorageSystemsInstanceDeploymentDetails**](StorageSystemsInstanceDeploymentDetails.md) |  | [optional] 

## Methods

### NewStorageSystemsInstance

`func NewStorageSystemsInstance(id string, ) *StorageSystemsInstance`

NewStorageSystemsInstance instantiates a new StorageSystemsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemsInstanceWithDefaults

`func NewStorageSystemsInstanceWithDefaults() *StorageSystemsInstance`

NewStorageSystemsInstanceWithDefaults instantiates a new StorageSystemsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageSystemsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageSystemsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageSystemsInstance) SetId(v string)`

SetId sets Id field to given value.


### GetSystemId

`func (o *StorageSystemsInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *StorageSystemsInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *StorageSystemsInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *StorageSystemsInstance) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetSystemType

`func (o *StorageSystemsInstance) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *StorageSystemsInstance) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *StorageSystemsInstance) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *StorageSystemsInstance) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetBandwidth

`func (o *StorageSystemsInstance) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *StorageSystemsInstance) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *StorageSystemsInstance) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *StorageSystemsInstance) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetCapacityImpact

`func (o *StorageSystemsInstance) GetCapacityImpact() int64`

GetCapacityImpact returns the CapacityImpact field if non-nil, zero value otherwise.

### GetCapacityImpactOk

`func (o *StorageSystemsInstance) GetCapacityImpactOk() (*int64, bool)`

GetCapacityImpactOk returns a tuple with the CapacityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityImpact

`func (o *StorageSystemsInstance) SetCapacityImpact(v int64)`

SetCapacityImpact sets CapacityImpact field to given value.

### HasCapacityImpact

`func (o *StorageSystemsInstance) HasCapacityImpact() bool`

HasCapacityImpact returns a boolean if a field has been set.

### GetCapacityIssueCount

`func (o *StorageSystemsInstance) GetCapacityIssueCount() int64`

GetCapacityIssueCount returns the CapacityIssueCount field if non-nil, zero value otherwise.

### GetCapacityIssueCountOk

`func (o *StorageSystemsInstance) GetCapacityIssueCountOk() (*int64, bool)`

GetCapacityIssueCountOk returns a tuple with the CapacityIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityIssueCount

`func (o *StorageSystemsInstance) SetCapacityIssueCount(v int64)`

SetCapacityIssueCount sets CapacityIssueCount field to given value.

### HasCapacityIssueCount

`func (o *StorageSystemsInstance) HasCapacityIssueCount() bool`

HasCapacityIssueCount returns a boolean if a field has been set.

### GetCompressionSavings

`func (o *StorageSystemsInstance) GetCompressionSavings() float64`

GetCompressionSavings returns the CompressionSavings field if non-nil, zero value otherwise.

### GetCompressionSavingsOk

`func (o *StorageSystemsInstance) GetCompressionSavingsOk() (*float64, bool)`

GetCompressionSavingsOk returns a tuple with the CompressionSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSavings

`func (o *StorageSystemsInstance) SetCompressionSavings(v float64)`

SetCompressionSavings sets CompressionSavings field to given value.

### HasCompressionSavings

`func (o *StorageSystemsInstance) HasCompressionSavings() bool`

HasCompressionSavings returns a boolean if a field has been set.

### GetConfigurationImpact

`func (o *StorageSystemsInstance) GetConfigurationImpact() int64`

GetConfigurationImpact returns the ConfigurationImpact field if non-nil, zero value otherwise.

### GetConfigurationImpactOk

`func (o *StorageSystemsInstance) GetConfigurationImpactOk() (*int64, bool)`

GetConfigurationImpactOk returns a tuple with the ConfigurationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationImpact

`func (o *StorageSystemsInstance) SetConfigurationImpact(v int64)`

SetConfigurationImpact sets ConfigurationImpact field to given value.

### HasConfigurationImpact

`func (o *StorageSystemsInstance) HasConfigurationImpact() bool`

HasConfigurationImpact returns a boolean if a field has been set.

### GetConfigurationIssueCount

`func (o *StorageSystemsInstance) GetConfigurationIssueCount() int64`

GetConfigurationIssueCount returns the ConfigurationIssueCount field if non-nil, zero value otherwise.

### GetConfigurationIssueCountOk

`func (o *StorageSystemsInstance) GetConfigurationIssueCountOk() (*int64, bool)`

GetConfigurationIssueCountOk returns a tuple with the ConfigurationIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIssueCount

`func (o *StorageSystemsInstance) SetConfigurationIssueCount(v int64)`

SetConfigurationIssueCount sets ConfigurationIssueCount field to given value.

### HasConfigurationIssueCount

`func (o *StorageSystemsInstance) HasConfigurationIssueCount() bool`

HasConfigurationIssueCount returns a boolean if a field has been set.

### GetConfiguredSize

`func (o *StorageSystemsInstance) GetConfiguredSize() int64`

GetConfiguredSize returns the ConfiguredSize field if non-nil, zero value otherwise.

### GetConfiguredSizeOk

`func (o *StorageSystemsInstance) GetConfiguredSizeOk() (*int64, bool)`

GetConfiguredSizeOk returns a tuple with the ConfiguredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredSize

`func (o *StorageSystemsInstance) SetConfiguredSize(v int64)`

SetConfiguredSize sets ConfiguredSize field to given value.

### HasConfiguredSize

`func (o *StorageSystemsInstance) HasConfiguredSize() bool`

HasConfiguredSize returns a boolean if a field has been set.

### GetConnectivityStatus

`func (o *StorageSystemsInstance) GetConnectivityStatus() string`

GetConnectivityStatus returns the ConnectivityStatus field if non-nil, zero value otherwise.

### GetConnectivityStatusOk

`func (o *StorageSystemsInstance) GetConnectivityStatusOk() (*string, bool)`

GetConnectivityStatusOk returns a tuple with the ConnectivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityStatus

`func (o *StorageSystemsInstance) SetConnectivityStatus(v string)`

SetConnectivityStatus sets ConnectivityStatus field to given value.

### HasConnectivityStatus

`func (o *StorageSystemsInstance) HasConnectivityStatus() bool`

HasConnectivityStatus returns a boolean if a field has been set.

### GetLicenseType

`func (o *StorageSystemsInstance) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *StorageSystemsInstance) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *StorageSystemsInstance) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *StorageSystemsInstance) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseExpirationDateTimestamp

`func (o *StorageSystemsInstance) GetLicenseExpirationDateTimestamp() time.Time`

GetLicenseExpirationDateTimestamp returns the LicenseExpirationDateTimestamp field if non-nil, zero value otherwise.

### GetLicenseExpirationDateTimestampOk

`func (o *StorageSystemsInstance) GetLicenseExpirationDateTimestampOk() (*time.Time, bool)`

GetLicenseExpirationDateTimestampOk returns a tuple with the LicenseExpirationDateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDateTimestamp

`func (o *StorageSystemsInstance) SetLicenseExpirationDateTimestamp(v time.Time)`

SetLicenseExpirationDateTimestamp sets LicenseExpirationDateTimestamp field to given value.

### HasLicenseExpirationDateTimestamp

`func (o *StorageSystemsInstance) HasLicenseExpirationDateTimestamp() bool`

HasLicenseExpirationDateTimestamp returns a boolean if a field has been set.

### GetContractCoverageType

`func (o *StorageSystemsInstance) GetContractCoverageType() string`

GetContractCoverageType returns the ContractCoverageType field if non-nil, zero value otherwise.

### GetContractCoverageTypeOk

`func (o *StorageSystemsInstance) GetContractCoverageTypeOk() (*string, bool)`

GetContractCoverageTypeOk returns a tuple with the ContractCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCoverageType

`func (o *StorageSystemsInstance) SetContractCoverageType(v string)`

SetContractCoverageType sets ContractCoverageType field to given value.

### HasContractCoverageType

`func (o *StorageSystemsInstance) HasContractCoverageType() bool`

HasContractCoverageType returns a boolean if a field has been set.

### GetContractExpirationDateTimestamp

`func (o *StorageSystemsInstance) GetContractExpirationDateTimestamp() time.Time`

GetContractExpirationDateTimestamp returns the ContractExpirationDateTimestamp field if non-nil, zero value otherwise.

### GetContractExpirationDateTimestampOk

`func (o *StorageSystemsInstance) GetContractExpirationDateTimestampOk() (*time.Time, bool)`

GetContractExpirationDateTimestampOk returns a tuple with the ContractExpirationDateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractExpirationDateTimestamp

`func (o *StorageSystemsInstance) SetContractExpirationDateTimestamp(v time.Time)`

SetContractExpirationDateTimestamp sets ContractExpirationDateTimestamp field to given value.

### HasContractExpirationDateTimestamp

`func (o *StorageSystemsInstance) HasContractExpirationDateTimestamp() bool`

HasContractExpirationDateTimestamp returns a boolean if a field has been set.

### GetDataProtectionImpact

`func (o *StorageSystemsInstance) GetDataProtectionImpact() int64`

GetDataProtectionImpact returns the DataProtectionImpact field if non-nil, zero value otherwise.

### GetDataProtectionImpactOk

`func (o *StorageSystemsInstance) GetDataProtectionImpactOk() (*int64, bool)`

GetDataProtectionImpactOk returns a tuple with the DataProtectionImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionImpact

`func (o *StorageSystemsInstance) SetDataProtectionImpact(v int64)`

SetDataProtectionImpact sets DataProtectionImpact field to given value.

### HasDataProtectionImpact

`func (o *StorageSystemsInstance) HasDataProtectionImpact() bool`

HasDataProtectionImpact returns a boolean if a field has been set.

### GetDataProtectionIssueCount

`func (o *StorageSystemsInstance) GetDataProtectionIssueCount() int64`

GetDataProtectionIssueCount returns the DataProtectionIssueCount field if non-nil, zero value otherwise.

### GetDataProtectionIssueCountOk

`func (o *StorageSystemsInstance) GetDataProtectionIssueCountOk() (*int64, bool)`

GetDataProtectionIssueCountOk returns a tuple with the DataProtectionIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionIssueCount

`func (o *StorageSystemsInstance) SetDataProtectionIssueCount(v int64)`

SetDataProtectionIssueCount sets DataProtectionIssueCount field to given value.

### HasDataProtectionIssueCount

`func (o *StorageSystemsInstance) HasDataProtectionIssueCount() bool`

HasDataProtectionIssueCount returns a boolean if a field has been set.

### GetDisplayIdentifier

`func (o *StorageSystemsInstance) GetDisplayIdentifier() string`

GetDisplayIdentifier returns the DisplayIdentifier field if non-nil, zero value otherwise.

### GetDisplayIdentifierOk

`func (o *StorageSystemsInstance) GetDisplayIdentifierOk() (*string, bool)`

GetDisplayIdentifierOk returns a tuple with the DisplayIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdentifier

`func (o *StorageSystemsInstance) SetDisplayIdentifier(v string)`

SetDisplayIdentifier sets DisplayIdentifier field to given value.

### HasDisplayIdentifier

`func (o *StorageSystemsInstance) HasDisplayIdentifier() bool`

HasDisplayIdentifier returns a boolean if a field has been set.

### GetFreePercent

`func (o *StorageSystemsInstance) GetFreePercent() float64`

GetFreePercent returns the FreePercent field if non-nil, zero value otherwise.

### GetFreePercentOk

`func (o *StorageSystemsInstance) GetFreePercentOk() (*float64, bool)`

GetFreePercentOk returns a tuple with the FreePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreePercent

`func (o *StorageSystemsInstance) SetFreePercent(v float64)`

SetFreePercent sets FreePercent field to given value.

### HasFreePercent

`func (o *StorageSystemsInstance) HasFreePercent() bool`

HasFreePercent returns a boolean if a field has been set.

### GetFreeSize

`func (o *StorageSystemsInstance) GetFreeSize() int64`

GetFreeSize returns the FreeSize field if non-nil, zero value otherwise.

### GetFreeSizeOk

`func (o *StorageSystemsInstance) GetFreeSizeOk() (*int64, bool)`

GetFreeSizeOk returns a tuple with the FreeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSize

`func (o *StorageSystemsInstance) SetFreeSize(v int64)`

SetFreeSize sets FreeSize field to given value.

### HasFreeSize

`func (o *StorageSystemsInstance) HasFreeSize() bool`

HasFreeSize returns a boolean if a field has been set.

### GetHealthConnectivityStatus

`func (o *StorageSystemsInstance) GetHealthConnectivityStatus() string`

GetHealthConnectivityStatus returns the HealthConnectivityStatus field if non-nil, zero value otherwise.

### GetHealthConnectivityStatusOk

`func (o *StorageSystemsInstance) GetHealthConnectivityStatusOk() (*string, bool)`

GetHealthConnectivityStatusOk returns a tuple with the HealthConnectivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthConnectivityStatus

`func (o *StorageSystemsInstance) SetHealthConnectivityStatus(v string)`

SetHealthConnectivityStatus sets HealthConnectivityStatus field to given value.

### HasHealthConnectivityStatus

`func (o *StorageSystemsInstance) HasHealthConnectivityStatus() bool`

HasHealthConnectivityStatus returns a boolean if a field has been set.

### GetHealthIssueCount

`func (o *StorageSystemsInstance) GetHealthIssueCount() int64`

GetHealthIssueCount returns the HealthIssueCount field if non-nil, zero value otherwise.

### GetHealthIssueCountOk

`func (o *StorageSystemsInstance) GetHealthIssueCountOk() (*int64, bool)`

GetHealthIssueCountOk returns a tuple with the HealthIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthIssueCount

`func (o *StorageSystemsInstance) SetHealthIssueCount(v int64)`

SetHealthIssueCount sets HealthIssueCount field to given value.

### HasHealthIssueCount

`func (o *StorageSystemsInstance) HasHealthIssueCount() bool`

HasHealthIssueCount returns a boolean if a field has been set.

### GetHealthScore

`func (o *StorageSystemsInstance) GetHealthScore() int64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *StorageSystemsInstance) GetHealthScoreOk() (*int64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *StorageSystemsInstance) SetHealthScore(v int64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *StorageSystemsInstance) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetHealthState

`func (o *StorageSystemsInstance) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *StorageSystemsInstance) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *StorageSystemsInstance) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *StorageSystemsInstance) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetIops

`func (o *StorageSystemsInstance) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *StorageSystemsInstance) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *StorageSystemsInstance) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *StorageSystemsInstance) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetIpv4Address

`func (o *StorageSystemsInstance) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *StorageSystemsInstance) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *StorageSystemsInstance) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *StorageSystemsInstance) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *StorageSystemsInstance) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *StorageSystemsInstance) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *StorageSystemsInstance) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *StorageSystemsInstance) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetLastContactTimestamp

`func (o *StorageSystemsInstance) GetLastContactTimestamp() time.Time`

GetLastContactTimestamp returns the LastContactTimestamp field if non-nil, zero value otherwise.

### GetLastContactTimestampOk

`func (o *StorageSystemsInstance) GetLastContactTimestampOk() (*time.Time, bool)`

GetLastContactTimestampOk returns a tuple with the LastContactTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactTimestamp

`func (o *StorageSystemsInstance) SetLastContactTimestamp(v time.Time)`

SetLastContactTimestamp sets LastContactTimestamp field to given value.

### HasLastContactTimestamp

`func (o *StorageSystemsInstance) HasLastContactTimestamp() bool`

HasLastContactTimestamp returns a boolean if a field has been set.

### GetLatency

`func (o *StorageSystemsInstance) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *StorageSystemsInstance) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *StorageSystemsInstance) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *StorageSystemsInstance) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLogicalSize

`func (o *StorageSystemsInstance) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *StorageSystemsInstance) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *StorageSystemsInstance) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.

### HasLogicalSize

`func (o *StorageSystemsInstance) HasLogicalSize() bool`

HasLogicalSize returns a boolean if a field has been set.

### GetModel

`func (o *StorageSystemsInstance) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageSystemsInstance) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageSystemsInstance) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageSystemsInstance) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *StorageSystemsInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSystemsInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSystemsInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSystemsInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverallEfficiency

`func (o *StorageSystemsInstance) GetOverallEfficiency() float64`

GetOverallEfficiency returns the OverallEfficiency field if non-nil, zero value otherwise.

### GetOverallEfficiencyOk

`func (o *StorageSystemsInstance) GetOverallEfficiencyOk() (*float64, bool)`

GetOverallEfficiencyOk returns a tuple with the OverallEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallEfficiency

`func (o *StorageSystemsInstance) SetOverallEfficiency(v float64)`

SetOverallEfficiency sets OverallEfficiency field to given value.

### HasOverallEfficiency

`func (o *StorageSystemsInstance) HasOverallEfficiency() bool`

HasOverallEfficiency returns a boolean if a field has been set.

### GetPerformanceImpact

`func (o *StorageSystemsInstance) GetPerformanceImpact() int64`

GetPerformanceImpact returns the PerformanceImpact field if non-nil, zero value otherwise.

### GetPerformanceImpactOk

`func (o *StorageSystemsInstance) GetPerformanceImpactOk() (*int64, bool)`

GetPerformanceImpactOk returns a tuple with the PerformanceImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceImpact

`func (o *StorageSystemsInstance) SetPerformanceImpact(v int64)`

SetPerformanceImpact sets PerformanceImpact field to given value.

### HasPerformanceImpact

`func (o *StorageSystemsInstance) HasPerformanceImpact() bool`

HasPerformanceImpact returns a boolean if a field has been set.

### GetPerformanceIssueCount

`func (o *StorageSystemsInstance) GetPerformanceIssueCount() int64`

GetPerformanceIssueCount returns the PerformanceIssueCount field if non-nil, zero value otherwise.

### GetPerformanceIssueCountOk

`func (o *StorageSystemsInstance) GetPerformanceIssueCountOk() (*int64, bool)`

GetPerformanceIssueCountOk returns a tuple with the PerformanceIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceIssueCount

`func (o *StorageSystemsInstance) SetPerformanceIssueCount(v int64)`

SetPerformanceIssueCount sets PerformanceIssueCount field to given value.

### HasPerformanceIssueCount

`func (o *StorageSystemsInstance) HasPerformanceIssueCount() bool`

HasPerformanceIssueCount returns a boolean if a field has been set.

### GetSerialNumber

`func (o *StorageSystemsInstance) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *StorageSystemsInstance) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *StorageSystemsInstance) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *StorageSystemsInstance) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteName

`func (o *StorageSystemsInstance) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *StorageSystemsInstance) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *StorageSystemsInstance) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *StorageSystemsInstance) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnapsSavings

`func (o *StorageSystemsInstance) GetSnapsSavings() float64`

GetSnapsSavings returns the SnapsSavings field if non-nil, zero value otherwise.

### GetSnapsSavingsOk

`func (o *StorageSystemsInstance) GetSnapsSavingsOk() (*float64, bool)`

GetSnapsSavingsOk returns a tuple with the SnapsSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapsSavings

`func (o *StorageSystemsInstance) SetSnapsSavings(v float64)`

SetSnapsSavings sets SnapsSavings field to given value.

### HasSnapsSavings

`func (o *StorageSystemsInstance) HasSnapsSavings() bool`

HasSnapsSavings returns a boolean if a field has been set.

### GetSystemHealthImpact

`func (o *StorageSystemsInstance) GetSystemHealthImpact() int64`

GetSystemHealthImpact returns the SystemHealthImpact field if non-nil, zero value otherwise.

### GetSystemHealthImpactOk

`func (o *StorageSystemsInstance) GetSystemHealthImpactOk() (*int64, bool)`

GetSystemHealthImpactOk returns a tuple with the SystemHealthImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemHealthImpact

`func (o *StorageSystemsInstance) SetSystemHealthImpact(v int64)`

SetSystemHealthImpact sets SystemHealthImpact field to given value.

### HasSystemHealthImpact

`func (o *StorageSystemsInstance) HasSystemHealthImpact() bool`

HasSystemHealthImpact returns a boolean if a field has been set.

### GetSystemHealthIssueCount

`func (o *StorageSystemsInstance) GetSystemHealthIssueCount() int64`

GetSystemHealthIssueCount returns the SystemHealthIssueCount field if non-nil, zero value otherwise.

### GetSystemHealthIssueCountOk

`func (o *StorageSystemsInstance) GetSystemHealthIssueCountOk() (*int64, bool)`

GetSystemHealthIssueCountOk returns a tuple with the SystemHealthIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemHealthIssueCount

`func (o *StorageSystemsInstance) SetSystemHealthIssueCount(v int64)`

SetSystemHealthIssueCount sets SystemHealthIssueCount field to given value.

### HasSystemHealthIssueCount

`func (o *StorageSystemsInstance) HasSystemHealthIssueCount() bool`

HasSystemHealthIssueCount returns a boolean if a field has been set.

### GetThinSavings

`func (o *StorageSystemsInstance) GetThinSavings() float64`

GetThinSavings returns the ThinSavings field if non-nil, zero value otherwise.

### GetThinSavingsOk

`func (o *StorageSystemsInstance) GetThinSavingsOk() (*float64, bool)`

GetThinSavingsOk returns a tuple with the ThinSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinSavings

`func (o *StorageSystemsInstance) SetThinSavings(v float64)`

SetThinSavings sets ThinSavings field to given value.

### HasThinSavings

`func (o *StorageSystemsInstance) HasThinSavings() bool`

HasThinSavings returns a boolean if a field has been set.

### GetTotalSize

`func (o *StorageSystemsInstance) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *StorageSystemsInstance) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *StorageSystemsInstance) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *StorageSystemsInstance) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetUnconfiguredSize

`func (o *StorageSystemsInstance) GetUnconfiguredSize() int64`

GetUnconfiguredSize returns the UnconfiguredSize field if non-nil, zero value otherwise.

### GetUnconfiguredSizeOk

`func (o *StorageSystemsInstance) GetUnconfiguredSizeOk() (*int64, bool)`

GetUnconfiguredSizeOk returns a tuple with the UnconfiguredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfiguredSize

`func (o *StorageSystemsInstance) SetUnconfiguredSize(v int64)`

SetUnconfiguredSize sets UnconfiguredSize field to given value.

### HasUnconfiguredSize

`func (o *StorageSystemsInstance) HasUnconfiguredSize() bool`

HasUnconfiguredSize returns a boolean if a field has been set.

### GetUsedPercent

`func (o *StorageSystemsInstance) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *StorageSystemsInstance) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *StorageSystemsInstance) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *StorageSystemsInstance) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetUsedSize

`func (o *StorageSystemsInstance) GetUsedSize() int64`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *StorageSystemsInstance) GetUsedSizeOk() (*int64, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *StorageSystemsInstance) SetUsedSize(v int64)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *StorageSystemsInstance) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.

### GetVendor

`func (o *StorageSystemsInstance) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageSystemsInstance) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageSystemsInstance) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageSystemsInstance) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *StorageSystemsInstance) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageSystemsInstance) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageSystemsInstance) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageSystemsInstance) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentDetails

`func (o *StorageSystemsInstance) GetDeploymentDetails() StorageSystemsInstanceDeploymentDetails`

GetDeploymentDetails returns the DeploymentDetails field if non-nil, zero value otherwise.

### GetDeploymentDetailsOk

`func (o *StorageSystemsInstance) GetDeploymentDetailsOk() (*StorageSystemsInstanceDeploymentDetails, bool)`

GetDeploymentDetailsOk returns a tuple with the DeploymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDetails

`func (o *StorageSystemsInstance) SetDeploymentDetails(v StorageSystemsInstanceDeploymentDetails)`

SetDeploymentDetails sets DeploymentDetails field to given value.

### HasDeploymentDetails

`func (o *StorageSystemsInstance) HasDeploymentDetails() bool`

HasDeploymentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



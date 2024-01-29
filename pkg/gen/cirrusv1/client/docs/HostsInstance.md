# HostsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Host identifier. | 
**SystemId** | Pointer to **string** | Unique identifier for the system that the host is connected to. | [optional] 
**SystemType** | Pointer to **string** | Product type of the system. | [optional] 
**Description** | Pointer to **string** | Description of the host. | [optional] 
**InitiatorCount** | Pointer to **int64** | Number of initiators that are connected between the host or server and the monitored system. | [optional] 
**InitiatorProtocol** | Pointer to **string** | Type of initiator (FC or iSCSI) that the host or server uses to connect to a monitored system. | [optional] 
**IssueCount** | Pointer to **int64** | Number of health issues that are present on the host or server. | [optional] 
**Name** | Pointer to **string** | Name of the host or server. | [optional] 
**NativeId** | Pointer to **string** | Identifier of the host, defined by the system. | [optional] 
**NetworkAddresses** | Pointer to **string** | IPv4 or IPv6 IP addresses, domain names, or netgroup name associated with the host or server. | [optional] 
**Type** | Pointer to **string** | Type of the host. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system of the host or server. | [optional] 
**SystemModel** | Pointer to **string** | Model of the system. | [optional] 
**SystemName** | Pointer to **string** | Name of the system. | [optional] 
**TotalSize** | Pointer to **int64** | Total size of all LUNs or Volumes that are provisioned to the host or server from the system - Unit: bytes | [optional] 

## Methods

### NewHostsInstance

`func NewHostsInstance(id string, ) *HostsInstance`

NewHostsInstance instantiates a new HostsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsInstanceWithDefaults

`func NewHostsInstanceWithDefaults() *HostsInstance`

NewHostsInstanceWithDefaults instantiates a new HostsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostsInstance) SetId(v string)`

SetId sets Id field to given value.


### GetSystemId

`func (o *HostsInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *HostsInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *HostsInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *HostsInstance) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetSystemType

`func (o *HostsInstance) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *HostsInstance) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *HostsInstance) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *HostsInstance) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetDescription

`func (o *HostsInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostsInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostsInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostsInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInitiatorCount

`func (o *HostsInstance) GetInitiatorCount() int64`

GetInitiatorCount returns the InitiatorCount field if non-nil, zero value otherwise.

### GetInitiatorCountOk

`func (o *HostsInstance) GetInitiatorCountOk() (*int64, bool)`

GetInitiatorCountOk returns a tuple with the InitiatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCount

`func (o *HostsInstance) SetInitiatorCount(v int64)`

SetInitiatorCount sets InitiatorCount field to given value.

### HasInitiatorCount

`func (o *HostsInstance) HasInitiatorCount() bool`

HasInitiatorCount returns a boolean if a field has been set.

### GetInitiatorProtocol

`func (o *HostsInstance) GetInitiatorProtocol() string`

GetInitiatorProtocol returns the InitiatorProtocol field if non-nil, zero value otherwise.

### GetInitiatorProtocolOk

`func (o *HostsInstance) GetInitiatorProtocolOk() (*string, bool)`

GetInitiatorProtocolOk returns a tuple with the InitiatorProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorProtocol

`func (o *HostsInstance) SetInitiatorProtocol(v string)`

SetInitiatorProtocol sets InitiatorProtocol field to given value.

### HasInitiatorProtocol

`func (o *HostsInstance) HasInitiatorProtocol() bool`

HasInitiatorProtocol returns a boolean if a field has been set.

### GetIssueCount

`func (o *HostsInstance) GetIssueCount() int64`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *HostsInstance) GetIssueCountOk() (*int64, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *HostsInstance) SetIssueCount(v int64)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *HostsInstance) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetName

`func (o *HostsInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostsInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostsInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostsInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeId

`func (o *HostsInstance) GetNativeId() string`

GetNativeId returns the NativeId field if non-nil, zero value otherwise.

### GetNativeIdOk

`func (o *HostsInstance) GetNativeIdOk() (*string, bool)`

GetNativeIdOk returns a tuple with the NativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeId

`func (o *HostsInstance) SetNativeId(v string)`

SetNativeId sets NativeId field to given value.

### HasNativeId

`func (o *HostsInstance) HasNativeId() bool`

HasNativeId returns a boolean if a field has been set.

### GetNetworkAddresses

`func (o *HostsInstance) GetNetworkAddresses() string`

GetNetworkAddresses returns the NetworkAddresses field if non-nil, zero value otherwise.

### GetNetworkAddressesOk

`func (o *HostsInstance) GetNetworkAddressesOk() (*string, bool)`

GetNetworkAddressesOk returns a tuple with the NetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddresses

`func (o *HostsInstance) SetNetworkAddresses(v string)`

SetNetworkAddresses sets NetworkAddresses field to given value.

### HasNetworkAddresses

`func (o *HostsInstance) HasNetworkAddresses() bool`

HasNetworkAddresses returns a boolean if a field has been set.

### GetType

`func (o *HostsInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostsInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostsInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostsInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *HostsInstance) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *HostsInstance) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *HostsInstance) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *HostsInstance) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetSystemModel

`func (o *HostsInstance) GetSystemModel() string`

GetSystemModel returns the SystemModel field if non-nil, zero value otherwise.

### GetSystemModelOk

`func (o *HostsInstance) GetSystemModelOk() (*string, bool)`

GetSystemModelOk returns a tuple with the SystemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModel

`func (o *HostsInstance) SetSystemModel(v string)`

SetSystemModel sets SystemModel field to given value.

### HasSystemModel

`func (o *HostsInstance) HasSystemModel() bool`

HasSystemModel returns a boolean if a field has been set.

### GetSystemName

`func (o *HostsInstance) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *HostsInstance) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *HostsInstance) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *HostsInstance) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTotalSize

`func (o *HostsInstance) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *HostsInstance) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *HostsInstance) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *HostsInstance) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LicensesInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of a licenses resource. | [optional] 
**Name** | Pointer to **string** | The name of the license. | [optional] 
**Type** | Pointer to [**LicensesTypeEnum**](LicensesTypeEnum.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | The time at which the license began.  | [optional] 
**ExpirationType** | Pointer to [**LicensesExpirationTypeEnum**](LicensesExpirationTypeEnum.md) |  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | For a TRIAL, EXT_TRIAL, TIME_LIMITED, GRACE_PERIOD, or RESTRICTED expiration type license, the time at which the license will expire.  | [optional] 
**Capacity** | Pointer to **int64** | For a CAPACITY type license, the limit on the license.  | [optional] 
**CapacityUnits** | Pointer to [**LicensesCapacityUnitsEnum**](LicensesCapacityUnitsEnum.md) |  | [optional] [default to LICENSESCAPACITYUNITSENUM_COUNT]
**EntitlementId** | Pointer to **string** | For licenses derived from ELMS entitlements, the entitlement from which this license was activated.  | [optional] 
**SystemId** | Pointer to **string** | The system on which this license is activated.  | [optional] 

## Methods

### NewLicensesInstance

`func NewLicensesInstance() *LicensesInstance`

NewLicensesInstance instantiates a new LicensesInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensesInstanceWithDefaults

`func NewLicensesInstanceWithDefaults() *LicensesInstance`

NewLicensesInstanceWithDefaults instantiates a new LicensesInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicensesInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensesInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensesInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicensesInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LicensesInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicensesInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicensesInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicensesInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *LicensesInstance) GetType() LicensesTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicensesInstance) GetTypeOk() (*LicensesTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicensesInstance) SetType(v LicensesTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *LicensesInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartTime

`func (o *LicensesInstance) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LicensesInstance) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LicensesInstance) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LicensesInstance) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetExpirationType

`func (o *LicensesInstance) GetExpirationType() LicensesExpirationTypeEnum`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *LicensesInstance) GetExpirationTypeOk() (*LicensesExpirationTypeEnum, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *LicensesInstance) SetExpirationType(v LicensesExpirationTypeEnum)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *LicensesInstance) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetExpirationTime

`func (o *LicensesInstance) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LicensesInstance) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LicensesInstance) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LicensesInstance) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetCapacity

`func (o *LicensesInstance) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *LicensesInstance) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *LicensesInstance) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *LicensesInstance) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCapacityUnits

`func (o *LicensesInstance) GetCapacityUnits() LicensesCapacityUnitsEnum`

GetCapacityUnits returns the CapacityUnits field if non-nil, zero value otherwise.

### GetCapacityUnitsOk

`func (o *LicensesInstance) GetCapacityUnitsOk() (*LicensesCapacityUnitsEnum, bool)`

GetCapacityUnitsOk returns a tuple with the CapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUnits

`func (o *LicensesInstance) SetCapacityUnits(v LicensesCapacityUnitsEnum)`

SetCapacityUnits sets CapacityUnits field to given value.

### HasCapacityUnits

`func (o *LicensesInstance) HasCapacityUnits() bool`

HasCapacityUnits returns a boolean if a field has been set.

### GetEntitlementId

`func (o *LicensesInstance) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *LicensesInstance) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *LicensesInstance) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *LicensesInstance) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### GetSystemId

`func (o *LicensesInstance) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *LicensesInstance) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *LicensesInstance) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *LicensesInstance) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EntitlementsInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of a entitlements resource. | [optional] 
**ProductType** | Pointer to [**ProductsTypeEnum**](ProductsTypeEnum.md) |  | [optional] 
**LicenseType** | Pointer to [**LicensesTypeEnum**](LicensesTypeEnum.md) |  | [optional] 
**ExpirationType** | Pointer to [**LicensesExpirationTypeEnum**](LicensesExpirationTypeEnum.md) |  | [optional] 
**PurchasedCapacity** | Pointer to **int64** | For a CAPACITY type entitlement, the limit on the entitlement, which limits the total capacity of activated licenses from this entitlement.  | [optional] 
**ActivatedCapacity** | Pointer to **int64** | For a CAPACITY type entitlement, the current sum of the capacities of the activated licenses from this entitlement.  | [optional] 
**AvailableCapacity** | Pointer to **int64** | For a CAPACITY type entitlement, the unactivated capacity, which is purchased_capacity-activated_capacity.  | [optional] 
**CapacityUnits** | Pointer to [**LicensesCapacityUnitsEnum**](LicensesCapacityUnitsEnum.md) |  | [optional] [default to LICENSESCAPACITYUNITSENUM_COUNT]
**StartTime** | Pointer to **time.Time** | The time at which the entitlement began.  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | For a TRIAL, EXT_TRIAL, TIME_LIMITED, GRACE_PERIOD, or RESTRICTED expiration type entitlement, the time at which the entitlement expires.  | [optional] 

## Methods

### NewEntitlementsInstance

`func NewEntitlementsInstance() *EntitlementsInstance`

NewEntitlementsInstance instantiates a new EntitlementsInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsInstanceWithDefaults

`func NewEntitlementsInstanceWithDefaults() *EntitlementsInstance`

NewEntitlementsInstanceWithDefaults instantiates a new EntitlementsInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementsInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementsInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementsInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementsInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *EntitlementsInstance) GetProductType() ProductsTypeEnum`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *EntitlementsInstance) GetProductTypeOk() (*ProductsTypeEnum, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *EntitlementsInstance) SetProductType(v ProductsTypeEnum)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *EntitlementsInstance) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetLicenseType

`func (o *EntitlementsInstance) GetLicenseType() LicensesTypeEnum`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *EntitlementsInstance) GetLicenseTypeOk() (*LicensesTypeEnum, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *EntitlementsInstance) SetLicenseType(v LicensesTypeEnum)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *EntitlementsInstance) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetExpirationType

`func (o *EntitlementsInstance) GetExpirationType() LicensesExpirationTypeEnum`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *EntitlementsInstance) GetExpirationTypeOk() (*LicensesExpirationTypeEnum, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *EntitlementsInstance) SetExpirationType(v LicensesExpirationTypeEnum)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *EntitlementsInstance) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetPurchasedCapacity

`func (o *EntitlementsInstance) GetPurchasedCapacity() int64`

GetPurchasedCapacity returns the PurchasedCapacity field if non-nil, zero value otherwise.

### GetPurchasedCapacityOk

`func (o *EntitlementsInstance) GetPurchasedCapacityOk() (*int64, bool)`

GetPurchasedCapacityOk returns a tuple with the PurchasedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedCapacity

`func (o *EntitlementsInstance) SetPurchasedCapacity(v int64)`

SetPurchasedCapacity sets PurchasedCapacity field to given value.

### HasPurchasedCapacity

`func (o *EntitlementsInstance) HasPurchasedCapacity() bool`

HasPurchasedCapacity returns a boolean if a field has been set.

### GetActivatedCapacity

`func (o *EntitlementsInstance) GetActivatedCapacity() int64`

GetActivatedCapacity returns the ActivatedCapacity field if non-nil, zero value otherwise.

### GetActivatedCapacityOk

`func (o *EntitlementsInstance) GetActivatedCapacityOk() (*int64, bool)`

GetActivatedCapacityOk returns a tuple with the ActivatedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedCapacity

`func (o *EntitlementsInstance) SetActivatedCapacity(v int64)`

SetActivatedCapacity sets ActivatedCapacity field to given value.

### HasActivatedCapacity

`func (o *EntitlementsInstance) HasActivatedCapacity() bool`

HasActivatedCapacity returns a boolean if a field has been set.

### GetAvailableCapacity

`func (o *EntitlementsInstance) GetAvailableCapacity() int64`

GetAvailableCapacity returns the AvailableCapacity field if non-nil, zero value otherwise.

### GetAvailableCapacityOk

`func (o *EntitlementsInstance) GetAvailableCapacityOk() (*int64, bool)`

GetAvailableCapacityOk returns a tuple with the AvailableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCapacity

`func (o *EntitlementsInstance) SetAvailableCapacity(v int64)`

SetAvailableCapacity sets AvailableCapacity field to given value.

### HasAvailableCapacity

`func (o *EntitlementsInstance) HasAvailableCapacity() bool`

HasAvailableCapacity returns a boolean if a field has been set.

### GetCapacityUnits

`func (o *EntitlementsInstance) GetCapacityUnits() LicensesCapacityUnitsEnum`

GetCapacityUnits returns the CapacityUnits field if non-nil, zero value otherwise.

### GetCapacityUnitsOk

`func (o *EntitlementsInstance) GetCapacityUnitsOk() (*LicensesCapacityUnitsEnum, bool)`

GetCapacityUnitsOk returns a tuple with the CapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUnits

`func (o *EntitlementsInstance) SetCapacityUnits(v LicensesCapacityUnitsEnum)`

SetCapacityUnits sets CapacityUnits field to given value.

### HasCapacityUnits

`func (o *EntitlementsInstance) HasCapacityUnits() bool`

HasCapacityUnits returns a boolean if a field has been set.

### GetStartTime

`func (o *EntitlementsInstance) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EntitlementsInstance) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EntitlementsInstance) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EntitlementsInstance) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *EntitlementsInstance) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *EntitlementsInstance) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *EntitlementsInstance) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *EntitlementsInstance) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



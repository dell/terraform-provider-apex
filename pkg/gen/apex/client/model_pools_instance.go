/*
APEX Navigator for Multicloud Storage REST APIs

The public API definitions for APEX Navigator for Multicloud Storage

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PoolsInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolsInstance{}

// PoolsInstance The pool object.
type PoolsInstance struct {
	// Identifier of the pool.
	Id string `json:"id"`
	// Unique identifier for the device or appliance.
	SystemId *string `json:"system_id,omitempty"`
	// Type of the system.
	SystemType *string `json:"system_type,omitempty"`
	// Available capacity - Unit: bytes
	FreeSize *int64 `json:"free_size,omitempty"`
	// Number of health issues that are present on the pool.
	IssueCount *int64 `json:"issue_count,omitempty"`
	// Name of the pool.
	Name *string `json:"name,omitempty"`
	// Native identifier of the pool, defined by the system.
	NativeId *string `json:"native_id,omitempty"`
	// Percentage of pool capacity that is provisioned.
	SubscribedPercent *float64 `json:"subscribed_percent,omitempty"`
	// Total subscribed capacity of the pool - Unit: bytes
	SubscribedSize *int64 `json:"subscribed_size,omitempty"`
	// Model of the system.
	SystemModel *string `json:"system_model,omitempty"`
	// Name of the system.
	SystemName *string `json:"system_name,omitempty"`
	// This is an enumerated type showing a prediction of when the pool may become full. Possible values are: DAY (imminent); FULL (pool is full); WEEK (full in a week); MONTH (full in a month); QUARTER (full within a quarter); BEYOND (more than a quarter to become full); LEARNING (not enough data to perform an analysis); UNPREDICTABLE (missing or invalid data); or UNKNOWN (system error).
	TimeToFullPrediction *string `json:"time_to_full_prediction,omitempty"`
	// Total capacity of the pool - Unit: bytes
	TotalSize *int64 `json:"total_size,omitempty"`
	// The type of pool.
	Type *string `json:"type,omitempty"`
	// Percentage of pool capacity that is being used.
	UsedPercent *float64 `json:"used_percent,omitempty"`
	// Capacity of the pool that is being used - Unit: bytes
	UsedSize *int64 `json:"used_size,omitempty"`
}

type _PoolsInstance PoolsInstance

// NewPoolsInstance instantiates a new PoolsInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolsInstance(id string) *PoolsInstance {
	this := PoolsInstance{}
	this.Id = id
	return &this
}

// NewPoolsInstanceWithDefaults instantiates a new PoolsInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolsInstanceWithDefaults() *PoolsInstance {
	this := PoolsInstance{}
	return &this
}

// GetId returns the Id field value
func (o *PoolsInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PoolsInstance) SetId(v string) {
	o.Id = v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *PoolsInstance) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *PoolsInstance) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *PoolsInstance) SetSystemId(v string) {
	o.SystemId = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *PoolsInstance) GetSystemType() string {
	if o == nil || IsNil(o.SystemType) {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSystemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SystemType) {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *PoolsInstance) HasSystemType() bool {
	if o != nil && !IsNil(o.SystemType) {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *PoolsInstance) SetSystemType(v string) {
	o.SystemType = &v
}

// GetFreeSize returns the FreeSize field value if set, zero value otherwise.
func (o *PoolsInstance) GetFreeSize() int64 {
	if o == nil || IsNil(o.FreeSize) {
		var ret int64
		return ret
	}
	return *o.FreeSize
}

// GetFreeSizeOk returns a tuple with the FreeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetFreeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FreeSize) {
		return nil, false
	}
	return o.FreeSize, true
}

// HasFreeSize returns a boolean if a field has been set.
func (o *PoolsInstance) HasFreeSize() bool {
	if o != nil && !IsNil(o.FreeSize) {
		return true
	}

	return false
}

// SetFreeSize gets a reference to the given int64 and assigns it to the FreeSize field.
func (o *PoolsInstance) SetFreeSize(v int64) {
	o.FreeSize = &v
}

// GetIssueCount returns the IssueCount field value if set, zero value otherwise.
func (o *PoolsInstance) GetIssueCount() int64 {
	if o == nil || IsNil(o.IssueCount) {
		var ret int64
		return ret
	}
	return *o.IssueCount
}

// GetIssueCountOk returns a tuple with the IssueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetIssueCountOk() (*int64, bool) {
	if o == nil || IsNil(o.IssueCount) {
		return nil, false
	}
	return o.IssueCount, true
}

// HasIssueCount returns a boolean if a field has been set.
func (o *PoolsInstance) HasIssueCount() bool {
	if o != nil && !IsNil(o.IssueCount) {
		return true
	}

	return false
}

// SetIssueCount gets a reference to the given int64 and assigns it to the IssueCount field.
func (o *PoolsInstance) SetIssueCount(v int64) {
	o.IssueCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolsInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolsInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolsInstance) SetName(v string) {
	o.Name = &v
}

// GetNativeId returns the NativeId field value if set, zero value otherwise.
func (o *PoolsInstance) GetNativeId() string {
	if o == nil || IsNil(o.NativeId) {
		var ret string
		return ret
	}
	return *o.NativeId
}

// GetNativeIdOk returns a tuple with the NativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetNativeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NativeId) {
		return nil, false
	}
	return o.NativeId, true
}

// HasNativeId returns a boolean if a field has been set.
func (o *PoolsInstance) HasNativeId() bool {
	if o != nil && !IsNil(o.NativeId) {
		return true
	}

	return false
}

// SetNativeId gets a reference to the given string and assigns it to the NativeId field.
func (o *PoolsInstance) SetNativeId(v string) {
	o.NativeId = &v
}

// GetSubscribedPercent returns the SubscribedPercent field value if set, zero value otherwise.
func (o *PoolsInstance) GetSubscribedPercent() float64 {
	if o == nil || IsNil(o.SubscribedPercent) {
		var ret float64
		return ret
	}
	return *o.SubscribedPercent
}

// GetSubscribedPercentOk returns a tuple with the SubscribedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSubscribedPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.SubscribedPercent) {
		return nil, false
	}
	return o.SubscribedPercent, true
}

// HasSubscribedPercent returns a boolean if a field has been set.
func (o *PoolsInstance) HasSubscribedPercent() bool {
	if o != nil && !IsNil(o.SubscribedPercent) {
		return true
	}

	return false
}

// SetSubscribedPercent gets a reference to the given float64 and assigns it to the SubscribedPercent field.
func (o *PoolsInstance) SetSubscribedPercent(v float64) {
	o.SubscribedPercent = &v
}

// GetSubscribedSize returns the SubscribedSize field value if set, zero value otherwise.
func (o *PoolsInstance) GetSubscribedSize() int64 {
	if o == nil || IsNil(o.SubscribedSize) {
		var ret int64
		return ret
	}
	return *o.SubscribedSize
}

// GetSubscribedSizeOk returns a tuple with the SubscribedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSubscribedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.SubscribedSize) {
		return nil, false
	}
	return o.SubscribedSize, true
}

// HasSubscribedSize returns a boolean if a field has been set.
func (o *PoolsInstance) HasSubscribedSize() bool {
	if o != nil && !IsNil(o.SubscribedSize) {
		return true
	}

	return false
}

// SetSubscribedSize gets a reference to the given int64 and assigns it to the SubscribedSize field.
func (o *PoolsInstance) SetSubscribedSize(v int64) {
	o.SubscribedSize = &v
}

// GetSystemModel returns the SystemModel field value if set, zero value otherwise.
func (o *PoolsInstance) GetSystemModel() string {
	if o == nil || IsNil(o.SystemModel) {
		var ret string
		return ret
	}
	return *o.SystemModel
}

// GetSystemModelOk returns a tuple with the SystemModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSystemModelOk() (*string, bool) {
	if o == nil || IsNil(o.SystemModel) {
		return nil, false
	}
	return o.SystemModel, true
}

// HasSystemModel returns a boolean if a field has been set.
func (o *PoolsInstance) HasSystemModel() bool {
	if o != nil && !IsNil(o.SystemModel) {
		return true
	}

	return false
}

// SetSystemModel gets a reference to the given string and assigns it to the SystemModel field.
func (o *PoolsInstance) SetSystemModel(v string) {
	o.SystemModel = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *PoolsInstance) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *PoolsInstance) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *PoolsInstance) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTimeToFullPrediction returns the TimeToFullPrediction field value if set, zero value otherwise.
func (o *PoolsInstance) GetTimeToFullPrediction() string {
	if o == nil || IsNil(o.TimeToFullPrediction) {
		var ret string
		return ret
	}
	return *o.TimeToFullPrediction
}

// GetTimeToFullPredictionOk returns a tuple with the TimeToFullPrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetTimeToFullPredictionOk() (*string, bool) {
	if o == nil || IsNil(o.TimeToFullPrediction) {
		return nil, false
	}
	return o.TimeToFullPrediction, true
}

// HasTimeToFullPrediction returns a boolean if a field has been set.
func (o *PoolsInstance) HasTimeToFullPrediction() bool {
	if o != nil && !IsNil(o.TimeToFullPrediction) {
		return true
	}

	return false
}

// SetTimeToFullPrediction gets a reference to the given string and assigns it to the TimeToFullPrediction field.
func (o *PoolsInstance) SetTimeToFullPrediction(v string) {
	o.TimeToFullPrediction = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *PoolsInstance) GetTotalSize() int64 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *PoolsInstance) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *PoolsInstance) SetTotalSize(v int64) {
	o.TotalSize = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PoolsInstance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PoolsInstance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PoolsInstance) SetType(v string) {
	o.Type = &v
}

// GetUsedPercent returns the UsedPercent field value if set, zero value otherwise.
func (o *PoolsInstance) GetUsedPercent() float64 {
	if o == nil || IsNil(o.UsedPercent) {
		var ret float64
		return ret
	}
	return *o.UsedPercent
}

// GetUsedPercentOk returns a tuple with the UsedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetUsedPercentOk() (*float64, bool) {
	if o == nil || IsNil(o.UsedPercent) {
		return nil, false
	}
	return o.UsedPercent, true
}

// HasUsedPercent returns a boolean if a field has been set.
func (o *PoolsInstance) HasUsedPercent() bool {
	if o != nil && !IsNil(o.UsedPercent) {
		return true
	}

	return false
}

// SetUsedPercent gets a reference to the given float64 and assigns it to the UsedPercent field.
func (o *PoolsInstance) SetUsedPercent(v float64) {
	o.UsedPercent = &v
}

// GetUsedSize returns the UsedSize field value if set, zero value otherwise.
func (o *PoolsInstance) GetUsedSize() int64 {
	if o == nil || IsNil(o.UsedSize) {
		var ret int64
		return ret
	}
	return *o.UsedSize
}

// GetUsedSizeOk returns a tuple with the UsedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsInstance) GetUsedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedSize) {
		return nil, false
	}
	return o.UsedSize, true
}

// HasUsedSize returns a boolean if a field has been set.
func (o *PoolsInstance) HasUsedSize() bool {
	if o != nil && !IsNil(o.UsedSize) {
		return true
	}

	return false
}

// SetUsedSize gets a reference to the given int64 and assigns it to the UsedSize field.
func (o *PoolsInstance) SetUsedSize(v int64) {
	o.UsedSize = &v
}

func (o PoolsInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolsInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.SystemType) {
		toSerialize["system_type"] = o.SystemType
	}
	if !IsNil(o.FreeSize) {
		toSerialize["free_size"] = o.FreeSize
	}
	if !IsNil(o.IssueCount) {
		toSerialize["issue_count"] = o.IssueCount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NativeId) {
		toSerialize["native_id"] = o.NativeId
	}
	if !IsNil(o.SubscribedPercent) {
		toSerialize["subscribed_percent"] = o.SubscribedPercent
	}
	if !IsNil(o.SubscribedSize) {
		toSerialize["subscribed_size"] = o.SubscribedSize
	}
	if !IsNil(o.SystemModel) {
		toSerialize["system_model"] = o.SystemModel
	}
	if !IsNil(o.SystemName) {
		toSerialize["system_name"] = o.SystemName
	}
	if !IsNil(o.TimeToFullPrediction) {
		toSerialize["time_to_full_prediction"] = o.TimeToFullPrediction
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UsedPercent) {
		toSerialize["used_percent"] = o.UsedPercent
	}
	if !IsNil(o.UsedSize) {
		toSerialize["used_size"] = o.UsedSize
	}
	return toSerialize, nil
}

func (o *PoolsInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPoolsInstance := _PoolsInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolsInstance)

	if err != nil {
		return err
	}

	*o = PoolsInstance(varPoolsInstance)

	return err
}

type NullablePoolsInstance struct {
	value *PoolsInstance
	isSet bool
}

func (v NullablePoolsInstance) Get() *PoolsInstance {
	return v.value
}

func (v *NullablePoolsInstance) Set(val *PoolsInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolsInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolsInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolsInstance(val *PoolsInstance) *NullablePoolsInstance {
	return &NullablePoolsInstance{value: val, isSet: true}
}

func (v NullablePoolsInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolsInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



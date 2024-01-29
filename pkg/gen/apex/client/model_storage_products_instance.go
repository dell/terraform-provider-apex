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

// checks if the StorageProductsInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageProductsInstance{}

// StorageProductsInstance Model for storage product information and its supported cloud mapping along with supported versions
type StorageProductsInstance struct {
	Id string `json:"id"`
	// Name of the storage product
	Name string `json:"name"`
	SystemType StorageProductEnum `json:"system_type"`
	StorageType StorageTypeEnum `json:"storage_type"`
	// Description of the storage product and its capabilities
	Description string `json:"description"`
	CloudType CloudProviderEnum `json:"cloud_type"`
	// Latest supported version of the storage product on the cloud
	LatestVersion string `json:"latest_version"`
	// Supported cloud and version details
	SupportMap []SupportMap `json:"support_map"`
}

type _StorageProductsInstance StorageProductsInstance

// NewStorageProductsInstance instantiates a new StorageProductsInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProductsInstance(id string, name string, systemType StorageProductEnum, storageType StorageTypeEnum, description string, cloudType CloudProviderEnum, latestVersion string, supportMap []SupportMap) *StorageProductsInstance {
	this := StorageProductsInstance{}
	this.Id = id
	this.Name = name
	this.SystemType = systemType
	this.StorageType = storageType
	this.Description = description
	this.CloudType = cloudType
	this.LatestVersion = latestVersion
	this.SupportMap = supportMap
	return &this
}

// NewStorageProductsInstanceWithDefaults instantiates a new StorageProductsInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProductsInstanceWithDefaults() *StorageProductsInstance {
	this := StorageProductsInstance{}
	var cloudType CloudProviderEnum = CLOUDPROVIDERENUM_AWS
	this.CloudType = cloudType
	return &this
}

// GetId returns the Id field value
func (o *StorageProductsInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StorageProductsInstance) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StorageProductsInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProductsInstance) SetName(v string) {
	o.Name = v
}

// GetSystemType returns the SystemType field value
func (o *StorageProductsInstance) GetSystemType() StorageProductEnum {
	if o == nil {
		var ret StorageProductEnum
		return ret
	}

	return o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetSystemTypeOk() (*StorageProductEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemType, true
}

// SetSystemType sets field value
func (o *StorageProductsInstance) SetSystemType(v StorageProductEnum) {
	o.SystemType = v
}

// GetStorageType returns the StorageType field value
func (o *StorageProductsInstance) GetStorageType() StorageTypeEnum {
	if o == nil {
		var ret StorageTypeEnum
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetStorageTypeOk() (*StorageTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *StorageProductsInstance) SetStorageType(v StorageTypeEnum) {
	o.StorageType = v
}

// GetDescription returns the Description field value
func (o *StorageProductsInstance) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *StorageProductsInstance) SetDescription(v string) {
	o.Description = v
}

// GetCloudType returns the CloudType field value
func (o *StorageProductsInstance) GetCloudType() CloudProviderEnum {
	if o == nil {
		var ret CloudProviderEnum
		return ret
	}

	return o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetCloudTypeOk() (*CloudProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudType, true
}

// SetCloudType sets field value
func (o *StorageProductsInstance) SetCloudType(v CloudProviderEnum) {
	o.CloudType = v
}

// GetLatestVersion returns the LatestVersion field value
func (o *StorageProductsInstance) GetLatestVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetLatestVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestVersion, true
}

// SetLatestVersion sets field value
func (o *StorageProductsInstance) SetLatestVersion(v string) {
	o.LatestVersion = v
}

// GetSupportMap returns the SupportMap field value
func (o *StorageProductsInstance) GetSupportMap() []SupportMap {
	if o == nil {
		var ret []SupportMap
		return ret
	}

	return o.SupportMap
}

// GetSupportMapOk returns a tuple with the SupportMap field value
// and a boolean to check if the value has been set.
func (o *StorageProductsInstance) GetSupportMapOk() ([]SupportMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportMap, true
}

// SetSupportMap sets field value
func (o *StorageProductsInstance) SetSupportMap(v []SupportMap) {
	o.SupportMap = v
}

func (o StorageProductsInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageProductsInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["system_type"] = o.SystemType
	toSerialize["storage_type"] = o.StorageType
	toSerialize["description"] = o.Description
	toSerialize["cloud_type"] = o.CloudType
	toSerialize["latest_version"] = o.LatestVersion
	toSerialize["support_map"] = o.SupportMap
	return toSerialize, nil
}

func (o *StorageProductsInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"system_type",
		"storage_type",
		"description",
		"cloud_type",
		"latest_version",
		"support_map",
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

	varStorageProductsInstance := _StorageProductsInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageProductsInstance)

	if err != nil {
		return err
	}

	*o = StorageProductsInstance(varStorageProductsInstance)

	return err
}

type NullableStorageProductsInstance struct {
	value *StorageProductsInstance
	isSet bool
}

func (v NullableStorageProductsInstance) Get() *StorageProductsInstance {
	return v.value
}

func (v *NullableStorageProductsInstance) Set(val *StorageProductsInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProductsInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProductsInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProductsInstance(val *StorageProductsInstance) *NullableStorageProductsInstance {
	return &NullableStorageProductsInstance{value: val, isSet: true}
}

func (v NullableStorageProductsInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProductsInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



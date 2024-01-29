/*
APEX Orchestration Platform - Job Management System (JMS) API

Provides management and visibility for APEX Orchestration Platform Jobs

API version: IGNORED - see resource tag's x-api-version for the specific version of this resource.
Contact: apex.mars@dell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedJobsCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedJobsCollection{}

// PagedJobsCollection struct for PagedJobsCollection
type PagedJobsCollection struct {
	Paging *PagingInformation `json:"paging,omitempty"`
	Results []JobsInstance `json:"results,omitempty"`
}

// NewPagedJobsCollection instantiates a new PagedJobsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedJobsCollection() *PagedJobsCollection {
	this := PagedJobsCollection{}
	return &this
}

// NewPagedJobsCollectionWithDefaults instantiates a new PagedJobsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedJobsCollectionWithDefaults() *PagedJobsCollection {
	this := PagedJobsCollection{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *PagedJobsCollection) GetPaging() PagingInformation {
	if o == nil || IsNil(o.Paging) {
		var ret PagingInformation
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedJobsCollection) GetPagingOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *PagedJobsCollection) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given PagingInformation and assigns it to the Paging field.
func (o *PagedJobsCollection) SetPaging(v PagingInformation) {
	o.Paging = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PagedJobsCollection) GetResults() []JobsInstance {
	if o == nil || IsNil(o.Results) {
		var ret []JobsInstance
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedJobsCollection) GetResultsOk() ([]JobsInstance, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PagedJobsCollection) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []JobsInstance and assigns it to the Results field.
func (o *PagedJobsCollection) SetResults(v []JobsInstance) {
	o.Results = v
}

func (o PagedJobsCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedJobsCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePagedJobsCollection struct {
	value *PagedJobsCollection
	isSet bool
}

func (v NullablePagedJobsCollection) Get() *PagedJobsCollection {
	return v.value
}

func (v *NullablePagedJobsCollection) Set(val *PagedJobsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedJobsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedJobsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedJobsCollection(val *PagedJobsCollection) *NullablePagedJobsCollection {
	return &NullablePagedJobsCollection{value: val, isSet: true}
}

func (v NullablePagedJobsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedJobsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



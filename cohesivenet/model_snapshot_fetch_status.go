/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesivenet

import (
	"encoding/json"
)

// SnapshotFetchStatus struct for SnapshotFetchStatus
type SnapshotFetchStatus struct {
	// Status of import
	Success *string `json:"success,omitempty"`
}

// NewSnapshotFetchStatus instantiates a new SnapshotFetchStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotFetchStatus() *SnapshotFetchStatus {
	this := SnapshotFetchStatus{}
	return &this
}

// NewSnapshotFetchStatusWithDefaults instantiates a new SnapshotFetchStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotFetchStatusWithDefaults() *SnapshotFetchStatus {
	this := SnapshotFetchStatus{}
	return &this
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SnapshotFetchStatus) GetSuccess() string {
	if o == nil || o.Success == nil {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotFetchStatus) GetSuccessOk() (*string, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SnapshotFetchStatus) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *SnapshotFetchStatus) SetSuccess(v string) {
	o.Success = &v
}


func (o SnapshotFetchStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotFetchStatus struct {
	value *SnapshotFetchStatus
	isSet bool
}

func (v NullableSnapshotFetchStatus) Get() *SnapshotFetchStatus {
	return v.value
}

func (v *NullableSnapshotFetchStatus) Set(val *SnapshotFetchStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotFetchStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotFetchStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotFetchStatus(val *SnapshotFetchStatus) *NullableSnapshotFetchStatus {
	return &NullableSnapshotFetchStatus{value: val, isSet: true}
}

func (v NullableSnapshotFetchStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotFetchStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
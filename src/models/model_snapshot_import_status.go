/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// SnapshotImportStatus struct for SnapshotImportStatus
type SnapshotImportStatus struct {
	// Status of import
	Snapshot *string `json:"snapshot,omitempty"`
	// Log details if failed
	LogLines []string `json:"log_lines,omitempty"`
}

// NewSnapshotImportStatus instantiates a new SnapshotImportStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotImportStatus() *SnapshotImportStatus {
	this := SnapshotImportStatus{}
	return &this
}

// NewSnapshotImportStatusWithDefaults instantiates a new SnapshotImportStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotImportStatusWithDefaults() *SnapshotImportStatus {
	this := SnapshotImportStatus{}
	return &this
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SnapshotImportStatus) GetSnapshot() string {
	if o == nil || o.Snapshot == nil {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotImportStatus) GetSnapshotOk() (*string, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SnapshotImportStatus) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *SnapshotImportStatus) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetLogLines returns the LogLines field value if set, zero value otherwise.
func (o *SnapshotImportStatus) GetLogLines() []string {
	if o == nil || o.LogLines == nil {
		var ret []string
		return ret
	}
	return o.LogLines
}

// GetLogLinesOk returns a tuple with the LogLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotImportStatus) GetLogLinesOk() ([]string, bool) {
	if o == nil || o.LogLines == nil {
		return nil, false
	}
	return o.LogLines, true
}

// HasLogLines returns a boolean if a field has been set.
func (o *SnapshotImportStatus) HasLogLines() bool {
	if o != nil && o.LogLines != nil {
		return true
	}

	return false
}

// SetLogLines gets a reference to the given []string and assigns it to the LogLines field.
func (o *SnapshotImportStatus) SetLogLines(v []string) {
	o.LogLines = v
}

func (o SnapshotImportStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snapshot != nil {
		toSerialize["snapshot"] = o.Snapshot
	}
	if o.LogLines != nil {
		toSerialize["log_lines"] = o.LogLines
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotImportStatus struct {
	value *SnapshotImportStatus
	isSet bool
}

func (v NullableSnapshotImportStatus) Get() *SnapshotImportStatus {
	return v.value
}

func (v *NullableSnapshotImportStatus) Set(val *SnapshotImportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotImportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotImportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotImportStatus(val *SnapshotImportStatus) *NullableSnapshotImportStatus {
	return &NullableSnapshotImportStatus{value: val, isSet: true}
}

func (v NullableSnapshotImportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotImportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


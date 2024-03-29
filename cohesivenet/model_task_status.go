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

// TaskStatus struct for TaskStatus
type TaskStatus struct {
	TaskStatus *string `json:"task_status,omitempty"`
}

// NewTaskStatus instantiates a new TaskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStatus() *TaskStatus {
	this := TaskStatus{}
	return &this
}

// NewTaskStatusWithDefaults instantiates a new TaskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStatusWithDefaults() *TaskStatus {
	this := TaskStatus{}
	return &this
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *TaskStatus) GetTaskStatus() string {
	if o == nil || o.TaskStatus == nil {
		var ret string
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetTaskStatusOk() (*string, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *TaskStatus) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given string and assigns it to the TaskStatus field.
func (o *TaskStatus) SetTaskStatus(v string) {
	o.TaskStatus = &v
}

func (o TaskStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskStatus != nil {
		toSerialize["task_status"] = o.TaskStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTaskStatus struct {
	value *TaskStatus
	isSet bool
}

func (v NullableTaskStatus) Get() *TaskStatus {
	return v.value
}

func (v *NullableTaskStatus) Set(val *TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatus(val *TaskStatus) *NullableTaskStatus {
	return &NullableTaskStatus{value: val, isSet: true}
}

func (v NullableTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



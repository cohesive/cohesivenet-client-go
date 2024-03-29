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

// CreateSnapshotRequest struct for CreateSnapshotRequest
type CreateSnapshotRequest struct {
	// Name of file. Defaults to a timestamped name
	Name *string `json:"name,omitempty"`
	// If true, will return a task id rather than wait for snapshot generation
	Async *bool `json:"async,omitempty"`
}

// NewCreateSnapshotRequest instantiates a new CreateSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotRequest() *CreateSnapshotRequest {
	this := CreateSnapshotRequest{}
	return &this
}

// NewCreateSnapshotRequestWithDefaults instantiates a new CreateSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotRequestWithDefaults() *CreateSnapshotRequest {
	this := CreateSnapshotRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSnapshotRequest) SetName(v string) {
	o.Name = &v
}

// GetAsync returns the Async field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetAsync() bool {
	if o == nil || o.Async == nil {
		var ret bool
		return ret
	}
	return *o.Async
}

// GetAsyncOk returns a tuple with the Async field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetAsyncOk() (*bool, bool) {
	if o == nil || o.Async == nil {
		return nil, false
	}
	return o.Async, true
}

// HasAsync returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasAsync() bool {
	if o != nil && o.Async != nil {
		return true
	}

	return false
}

// SetAsync gets a reference to the given bool and assigns it to the Async field.
func (o *CreateSnapshotRequest) SetAsync(v bool) {
	o.Async = &v
}

func (o CreateSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Async != nil {
		toSerialize["async"] = o.Async
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSnapshotRequest struct {
	value *CreateSnapshotRequest
	isSet bool
}

func (v NullableCreateSnapshotRequest) Get() *CreateSnapshotRequest {
	return v.value
}

func (v *NullableCreateSnapshotRequest) Set(val *CreateSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotRequest(val *CreateSnapshotRequest) *NullableCreateSnapshotRequest {
	return &NullableCreateSnapshotRequest{value: val, isSet: true}
}

func (v NullableCreateSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



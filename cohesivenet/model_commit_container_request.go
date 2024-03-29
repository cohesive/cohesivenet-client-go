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

// CommitContainerRequest struct for CommitContainerRequest
type CommitContainerRequest struct {
	// Name of new image
	Name bool `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewCommitContainerRequest instantiates a new CommitContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitContainerRequest(name bool) *CommitContainerRequest {
	this := CommitContainerRequest{}
	this.Name = name
	return &this
}

// NewCommitContainerRequestWithDefaults instantiates a new CommitContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitContainerRequestWithDefaults() *CommitContainerRequest {
	this := CommitContainerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CommitContainerRequest) GetName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommitContainerRequest) GetNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CommitContainerRequest) SetName(v bool) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommitContainerRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitContainerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommitContainerRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommitContainerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CommitContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCommitContainerRequest struct {
	value *CommitContainerRequest
	isSet bool
}

func (v NullableCommitContainerRequest) Get() *CommitContainerRequest {
	return v.value
}

func (v *NullableCommitContainerRequest) Set(val *CommitContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitContainerRequest(val *CommitContainerRequest) *NullableCommitContainerRequest {
	return &NullableCommitContainerRequest{value: val, isSet: true}
}

func (v NullableCommitContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DeletePluginInstanceRequest struct for DeletePluginInstanceRequest
type DeletePluginInstanceRequest struct {
	// Stop and delete if running
	Force *bool `json:"force,omitempty"`
}

// NewDeletePluginInstanceRequest instantiates a new DeletePluginInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePluginInstanceRequest() *DeletePluginInstanceRequest {
	this := DeletePluginInstanceRequest{}
	return &this
}

// NewDeletePluginInstanceRequestWithDefaults instantiates a new DeletePluginInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePluginInstanceRequestWithDefaults() *DeletePluginInstanceRequest {
	this := DeletePluginInstanceRequest{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DeletePluginInstanceRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePluginInstanceRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DeletePluginInstanceRequest) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DeletePluginInstanceRequest) SetForce(v bool) {
	o.Force = &v
}

func (o DeletePluginInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableDeletePluginInstanceRequest struct {
	value *DeletePluginInstanceRequest
	isSet bool
}

func (v NullableDeletePluginInstanceRequest) Get() *DeletePluginInstanceRequest {
	return v.value
}

func (v *NullableDeletePluginInstanceRequest) Set(val *DeletePluginInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePluginInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePluginInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePluginInstanceRequest(val *DeletePluginInstanceRequest) *NullableDeletePluginInstanceRequest {
	return &NullableDeletePluginInstanceRequest{value: val, isSet: true}
}

func (v NullableDeletePluginInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePluginInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



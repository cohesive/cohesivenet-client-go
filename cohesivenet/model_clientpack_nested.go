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

// ClientpackNested struct for ClientpackNested
type ClientpackNested struct {
	Clientpack *Clientpack `json:"clientpack,omitempty"`
}

// NewClientpackNested instantiates a new ClientpackNested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackNested() *ClientpackNested {
	this := ClientpackNested{}
	return &this
}

// NewClientpackNestedWithDefaults instantiates a new ClientpackNested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackNestedWithDefaults() *ClientpackNested {
	this := ClientpackNested{}
	return &this
}

// GetClientpack returns the Clientpack field value if set, zero value otherwise.
func (o *ClientpackNested) GetClientpack() Clientpack {
	if o == nil || o.Clientpack == nil {
		var ret Clientpack
		return ret
	}
	return *o.Clientpack
}

// GetClientpackOk returns a tuple with the Clientpack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackNested) GetClientpackOk() (*Clientpack, bool) {
	if o == nil || o.Clientpack == nil {
		return nil, false
	}
	return o.Clientpack, true
}

// HasClientpack returns a boolean if a field has been set.
func (o *ClientpackNested) HasClientpack() bool {
	if o != nil && o.Clientpack != nil {
		return true
	}

	return false
}

// SetClientpack gets a reference to the given Clientpack and assigns it to the Clientpack field.
func (o *ClientpackNested) SetClientpack(v Clientpack) {
	o.Clientpack = &v
}

func (o ClientpackNested) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clientpack != nil {
		toSerialize["clientpack"] = o.Clientpack
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackNested struct {
	value *ClientpackNested
	isSet bool
}

func (v NullableClientpackNested) Get() *ClientpackNested {
	return v.value
}

func (v *NullableClientpackNested) Set(val *ClientpackNested) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackNested) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackNested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackNested(val *ClientpackNested) *NullableClientpackNested {
	return &NullableClientpackNested{value: val, isSet: true}
}

func (v NullableClientpackNested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackNested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



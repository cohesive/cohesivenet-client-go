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

// StartPluginInstanceRequestEnvironmentInner struct for StartPluginInstanceRequestEnvironmentInner
type StartPluginInstanceRequestEnvironmentInner struct {
	Key *string `json:"key,omitempty"`
	Value *StartPluginInstanceRequestEnvironmentInnerValue `json:"value,omitempty"`
}

// NewStartPluginInstanceRequestEnvironmentInner instantiates a new StartPluginInstanceRequestEnvironmentInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartPluginInstanceRequestEnvironmentInner() *StartPluginInstanceRequestEnvironmentInner {
	this := StartPluginInstanceRequestEnvironmentInner{}
	return &this
}

// NewStartPluginInstanceRequestEnvironmentInnerWithDefaults instantiates a new StartPluginInstanceRequestEnvironmentInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartPluginInstanceRequestEnvironmentInnerWithDefaults() *StartPluginInstanceRequestEnvironmentInner {
	this := StartPluginInstanceRequestEnvironmentInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StartPluginInstanceRequestEnvironmentInner) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartPluginInstanceRequestEnvironmentInner) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StartPluginInstanceRequestEnvironmentInner) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StartPluginInstanceRequestEnvironmentInner) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StartPluginInstanceRequestEnvironmentInner) GetValue() StartPluginInstanceRequestEnvironmentInnerValue {
	if o == nil || o.Value == nil {
		var ret StartPluginInstanceRequestEnvironmentInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartPluginInstanceRequestEnvironmentInner) GetValueOk() (*StartPluginInstanceRequestEnvironmentInnerValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StartPluginInstanceRequestEnvironmentInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given StartPluginInstanceRequestEnvironmentInnerValue and assigns it to the Value field.
func (o *StartPluginInstanceRequestEnvironmentInner) SetValue(v StartPluginInstanceRequestEnvironmentInnerValue) {
	o.Value = &v
}

func (o StartPluginInstanceRequestEnvironmentInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStartPluginInstanceRequestEnvironmentInner struct {
	value *StartPluginInstanceRequestEnvironmentInner
	isSet bool
}

func (v NullableStartPluginInstanceRequestEnvironmentInner) Get() *StartPluginInstanceRequestEnvironmentInner {
	return v.value
}

func (v *NullableStartPluginInstanceRequestEnvironmentInner) Set(val *StartPluginInstanceRequestEnvironmentInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStartPluginInstanceRequestEnvironmentInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStartPluginInstanceRequestEnvironmentInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartPluginInstanceRequestEnvironmentInner(val *StartPluginInstanceRequestEnvironmentInner) *NullableStartPluginInstanceRequestEnvironmentInner {
	return &NullableStartPluginInstanceRequestEnvironmentInner{value: val, isSet: true}
}

func (v NullableStartPluginInstanceRequestEnvironmentInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartPluginInstanceRequestEnvironmentInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


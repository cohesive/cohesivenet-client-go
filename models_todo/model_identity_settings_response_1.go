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

// IdentitySettingsResponse1 struct for IdentitySettingsResponse1
type IdentitySettingsResponse1 struct {
	Response *IdentitySettingsResponse1Response `json:"response,omitempty"`
}

// NewIdentitySettingsResponse1 instantiates a new IdentitySettingsResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySettingsResponse1() *IdentitySettingsResponse1 {
	this := IdentitySettingsResponse1{}
	return &this
}

// NewIdentitySettingsResponse1WithDefaults instantiates a new IdentitySettingsResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySettingsResponse1WithDefaults() *IdentitySettingsResponse1 {
	this := IdentitySettingsResponse1{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *IdentitySettingsResponse1) GetResponse() IdentitySettingsResponse1Response {
	if o == nil || o.Response == nil {
		var ret IdentitySettingsResponse1Response
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySettingsResponse1) GetResponseOk() (*IdentitySettingsResponse1Response, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *IdentitySettingsResponse1) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given IdentitySettingsResponse1Response and assigns it to the Response field.
func (o *IdentitySettingsResponse1) SetResponse(v IdentitySettingsResponse1Response) {
	o.Response = &v
}

func (o IdentitySettingsResponse1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySettingsResponse1 struct {
	value *IdentitySettingsResponse1
	isSet bool
}

func (v NullableIdentitySettingsResponse1) Get() *IdentitySettingsResponse1 {
	return v.value
}

func (v *NullableIdentitySettingsResponse1) Set(val *IdentitySettingsResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySettingsResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySettingsResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySettingsResponse1(val *IdentitySettingsResponse1) *NullableIdentitySettingsResponse1 {
	return &NullableIdentitySettingsResponse1{value: val, isSet: true}
}

func (v NullableIdentitySettingsResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySettingsResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


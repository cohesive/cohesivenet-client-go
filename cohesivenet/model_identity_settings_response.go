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

// IdentitySettingsResponse struct for IdentitySettingsResponse
type IdentitySettingsResponse struct {
	Response *IdentitySettings `json:"response,omitempty"`
}

// NewIdentitySettingsResponse instantiates a new IdentitySettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySettingsResponse() *IdentitySettingsResponse {
	this := IdentitySettingsResponse{}
	return &this
}

// NewIdentitySettingsResponseWithDefaults instantiates a new IdentitySettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySettingsResponseWithDefaults() *IdentitySettingsResponse {
	this := IdentitySettingsResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *IdentitySettingsResponse) GetResponse() IdentitySettings {
	if o == nil || o.Response == nil {
		var ret IdentitySettings
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySettingsResponse) GetResponseOk() (*IdentitySettings, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *IdentitySettingsResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given IdentitySettings and assigns it to the Response field.
func (o *IdentitySettingsResponse) SetResponse(v IdentitySettings) {
	o.Response = &v
}

func (o IdentitySettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySettingsResponse struct {
	value *IdentitySettingsResponse
	isSet bool
}

func (v NullableIdentitySettingsResponse) Get() *IdentitySettingsResponse {
	return v.value
}

func (v *NullableIdentitySettingsResponse) Set(val *IdentitySettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySettingsResponse(val *IdentitySettingsResponse) *NullableIdentitySettingsResponse {
	return &NullableIdentitySettingsResponse{value: val, isSet: true}
}

func (v NullableIdentitySettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



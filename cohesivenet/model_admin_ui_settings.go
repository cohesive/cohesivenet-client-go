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

// AdminUISettings struct for AdminUISettings
type AdminUISettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewAdminUISettings instantiates a new AdminUISettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUISettings() *AdminUISettings {
	this := AdminUISettings{}
	return &this
}

// NewAdminUISettingsWithDefaults instantiates a new AdminUISettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUISettingsWithDefaults() *AdminUISettings {
	this := AdminUISettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AdminUISettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUISettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AdminUISettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AdminUISettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AdminUISettings) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUISettings) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AdminUISettings) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AdminUISettings) SetUsername(v string) {
	o.Username = &v
}

func (o AdminUISettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAdminUISettings struct {
	value *AdminUISettings
	isSet bool
}

func (v NullableAdminUISettings) Get() *AdminUISettings {
	return v.value
}

func (v *NullableAdminUISettings) Set(val *AdminUISettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUISettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUISettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUISettings(val *AdminUISettings) *NullableAdminUISettings {
	return &NullableAdminUISettings{value: val, isSet: true}
}

func (v NullableAdminUISettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUISettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// UpdateAdminUISettingsRequest struct for UpdateAdminUISettingsRequest
type UpdateAdminUISettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdminUsername *string `json:"admin_username,omitempty"`
	AdminPassword *string `json:"admin_password,omitempty"`
}

// NewUpdateAdminUISettingsRequest instantiates a new UpdateAdminUISettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAdminUISettingsRequest() *UpdateAdminUISettingsRequest {
	this := UpdateAdminUISettingsRequest{}
	return &this
}

// NewUpdateAdminUISettingsRequestWithDefaults instantiates a new UpdateAdminUISettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAdminUISettingsRequestWithDefaults() *UpdateAdminUISettingsRequest {
	this := UpdateAdminUISettingsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateAdminUISettingsRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminUISettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateAdminUISettingsRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateAdminUISettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminUsername returns the AdminUsername field value if set, zero value otherwise.
func (o *UpdateAdminUISettingsRequest) GetAdminUsername() string {
	if o == nil || o.AdminUsername == nil {
		var ret string
		return ret
	}
	return *o.AdminUsername
}

// GetAdminUsernameOk returns a tuple with the AdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminUISettingsRequest) GetAdminUsernameOk() (*string, bool) {
	if o == nil || o.AdminUsername == nil {
		return nil, false
	}
	return o.AdminUsername, true
}

// HasAdminUsername returns a boolean if a field has been set.
func (o *UpdateAdminUISettingsRequest) HasAdminUsername() bool {
	if o != nil && o.AdminUsername != nil {
		return true
	}

	return false
}

// SetAdminUsername gets a reference to the given string and assigns it to the AdminUsername field.
func (o *UpdateAdminUISettingsRequest) SetAdminUsername(v string) {
	o.AdminUsername = &v
}

// GetAdminPassword returns the AdminPassword field value if set, zero value otherwise.
func (o *UpdateAdminUISettingsRequest) GetAdminPassword() string {
	if o == nil || o.AdminPassword == nil {
		var ret string
		return ret
	}
	return *o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAdminUISettingsRequest) GetAdminPasswordOk() (*string, bool) {
	if o == nil || o.AdminPassword == nil {
		return nil, false
	}
	return o.AdminPassword, true
}

// HasAdminPassword returns a boolean if a field has been set.
func (o *UpdateAdminUISettingsRequest) HasAdminPassword() bool {
	if o != nil && o.AdminPassword != nil {
		return true
	}

	return false
}

// SetAdminPassword gets a reference to the given string and assigns it to the AdminPassword field.
func (o *UpdateAdminUISettingsRequest) SetAdminPassword(v string) {
	o.AdminPassword = &v
}

func (o UpdateAdminUISettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AdminUsername != nil {
		toSerialize["admin_username"] = o.AdminUsername
	}
	if o.AdminPassword != nil {
		toSerialize["admin_password"] = o.AdminPassword
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAdminUISettingsRequest struct {
	value *UpdateAdminUISettingsRequest
	isSet bool
}

func (v NullableUpdateAdminUISettingsRequest) Get() *UpdateAdminUISettingsRequest {
	return v.value
}

func (v *NullableUpdateAdminUISettingsRequest) Set(val *UpdateAdminUISettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAdminUISettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAdminUISettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAdminUISettingsRequest(val *UpdateAdminUISettingsRequest) *NullableUpdateAdminUISettingsRequest {
	return &NullableUpdateAdminUISettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateAdminUISettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAdminUISettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// IdentityRadiusSettings struct for IdentityRadiusSettings
type IdentityRadiusSettings struct {
	// IP address or resolvable hostname
	Server *string `json:"server,omitempty"`
	// Authentication port
	AuthPort *int32 `json:"auth_port,omitempty"`
	AccountingPort *int32 `json:"accounting_port,omitempty"`
	// Shared password
	Pass *string `json:"pass,omitempty"`
}

// NewIdentityRadiusSettings instantiates a new IdentityRadiusSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRadiusSettings() *IdentityRadiusSettings {
	this := IdentityRadiusSettings{}
	var authPort int32 = 1812
	this.AuthPort = &authPort
	var accountingPort int32 = 1812
	this.AccountingPort = &accountingPort
	return &this
}

// NewIdentityRadiusSettingsWithDefaults instantiates a new IdentityRadiusSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRadiusSettingsWithDefaults() *IdentityRadiusSettings {
	this := IdentityRadiusSettings{}
	var authPort int32 = 1812
	this.AuthPort = &authPort
	var accountingPort int32 = 1812
	this.AccountingPort = &accountingPort
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *IdentityRadiusSettings) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRadiusSettings) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *IdentityRadiusSettings) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *IdentityRadiusSettings) SetServer(v string) {
	o.Server = &v
}

// GetAuthPort returns the AuthPort field value if set, zero value otherwise.
func (o *IdentityRadiusSettings) GetAuthPort() int32 {
	if o == nil || o.AuthPort == nil {
		var ret int32
		return ret
	}
	return *o.AuthPort
}

// GetAuthPortOk returns a tuple with the AuthPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRadiusSettings) GetAuthPortOk() (*int32, bool) {
	if o == nil || o.AuthPort == nil {
		return nil, false
	}
	return o.AuthPort, true
}

// HasAuthPort returns a boolean if a field has been set.
func (o *IdentityRadiusSettings) HasAuthPort() bool {
	if o != nil && o.AuthPort != nil {
		return true
	}

	return false
}

// SetAuthPort gets a reference to the given int32 and assigns it to the AuthPort field.
func (o *IdentityRadiusSettings) SetAuthPort(v int32) {
	o.AuthPort = &v
}

// GetAccountingPort returns the AccountingPort field value if set, zero value otherwise.
func (o *IdentityRadiusSettings) GetAccountingPort() int32 {
	if o == nil || o.AccountingPort == nil {
		var ret int32
		return ret
	}
	return *o.AccountingPort
}

// GetAccountingPortOk returns a tuple with the AccountingPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRadiusSettings) GetAccountingPortOk() (*int32, bool) {
	if o == nil || o.AccountingPort == nil {
		return nil, false
	}
	return o.AccountingPort, true
}

// HasAccountingPort returns a boolean if a field has been set.
func (o *IdentityRadiusSettings) HasAccountingPort() bool {
	if o != nil && o.AccountingPort != nil {
		return true
	}

	return false
}

// SetAccountingPort gets a reference to the given int32 and assigns it to the AccountingPort field.
func (o *IdentityRadiusSettings) SetAccountingPort(v int32) {
	o.AccountingPort = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *IdentityRadiusSettings) GetPass() string {
	if o == nil || o.Pass == nil {
		var ret string
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRadiusSettings) GetPassOk() (*string, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *IdentityRadiusSettings) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given string and assigns it to the Pass field.
func (o *IdentityRadiusSettings) SetPass(v string) {
	o.Pass = &v
}

func (o IdentityRadiusSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.AuthPort != nil {
		toSerialize["auth_port"] = o.AuthPort
	}
	if o.AccountingPort != nil {
		toSerialize["accounting_port"] = o.AccountingPort
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRadiusSettings struct {
	value *IdentityRadiusSettings
	isSet bool
}

func (v NullableIdentityRadiusSettings) Get() *IdentityRadiusSettings {
	return v.value
}

func (v *NullableIdentityRadiusSettings) Set(val *IdentityRadiusSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRadiusSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRadiusSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRadiusSettings(val *IdentityRadiusSettings) *NullableIdentityRadiusSettings {
	return &NullableIdentityRadiusSettings{value: val, isSet: true}
}

func (v NullableIdentityRadiusSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRadiusSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


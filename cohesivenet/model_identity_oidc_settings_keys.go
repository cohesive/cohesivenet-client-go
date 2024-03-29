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
	"time"
)

// IdentityOIDCSettingsKeys struct for IdentityOIDCSettingsKeys
type IdentityOIDCSettingsKeys struct {
	Keys []map[string]interface{} `json:"keys,omitempty"`
	KeysDate *time.Time `json:"keys_date,omitempty"`
}

// NewIdentityOIDCSettingsKeys instantiates a new IdentityOIDCSettingsKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityOIDCSettingsKeys() *IdentityOIDCSettingsKeys {
	this := IdentityOIDCSettingsKeys{}
	return &this
}

// NewIdentityOIDCSettingsKeysWithDefaults instantiates a new IdentityOIDCSettingsKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityOIDCSettingsKeysWithDefaults() *IdentityOIDCSettingsKeys {
	this := IdentityOIDCSettingsKeys{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsKeys) GetKeys() []map[string]interface{} {
	if o == nil || o.Keys == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsKeys) GetKeysOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsKeys) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []map[string]interface{} and assigns it to the Keys field.
func (o *IdentityOIDCSettingsKeys) SetKeys(v []map[string]interface{}) {
	o.Keys = v
}

// GetKeysDate returns the KeysDate field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsKeys) GetKeysDate() time.Time {
	if o == nil || o.KeysDate == nil {
		var ret time.Time
		return ret
	}
	return *o.KeysDate
}

// GetKeysDateOk returns a tuple with the KeysDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsKeys) GetKeysDateOk() (*time.Time, bool) {
	if o == nil || o.KeysDate == nil {
		return nil, false
	}
	return o.KeysDate, true
}

// HasKeysDate returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsKeys) HasKeysDate() bool {
	if o != nil && o.KeysDate != nil {
		return true
	}

	return false
}

// SetKeysDate gets a reference to the given time.Time and assigns it to the KeysDate field.
func (o *IdentityOIDCSettingsKeys) SetKeysDate(v time.Time) {
	o.KeysDate = &v
}

func (o IdentityOIDCSettingsKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.KeysDate != nil {
		toSerialize["keys_date"] = o.KeysDate
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityOIDCSettingsKeys struct {
	value *IdentityOIDCSettingsKeys
	isSet bool
}

func (v NullableIdentityOIDCSettingsKeys) Get() *IdentityOIDCSettingsKeys {
	return v.value
}

func (v *NullableIdentityOIDCSettingsKeys) Set(val *IdentityOIDCSettingsKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityOIDCSettingsKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityOIDCSettingsKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityOIDCSettingsKeys(val *IdentityOIDCSettingsKeys) *NullableIdentityOIDCSettingsKeys {
	return &NullableIdentityOIDCSettingsKeys{value: val, isSet: true}
}

func (v NullableIdentityOIDCSettingsKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityOIDCSettingsKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



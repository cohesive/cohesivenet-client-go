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

// SSLCertsList struct for SSLCertsList
type SSLCertsList struct {
	Certs []SSLCert `json:"certs,omitempty"`
}

// NewSSLCertsList instantiates a new SSLCertsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSLCertsList() *SSLCertsList {
	this := SSLCertsList{}
	return &this
}

// NewSSLCertsListWithDefaults instantiates a new SSLCertsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSLCertsListWithDefaults() *SSLCertsList {
	this := SSLCertsList{}
	return &this
}

// GetCerts returns the Certs field value if set, zero value otherwise.
func (o *SSLCertsList) GetCerts() []SSLCert {
	if o == nil || o.Certs == nil {
		var ret []SSLCert
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLCertsList) GetCertsOk() ([]SSLCert, bool) {
	if o == nil || o.Certs == nil {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *SSLCertsList) HasCerts() bool {
	if o != nil && o.Certs != nil {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []SSLCert and assigns it to the Certs field.
func (o *SSLCertsList) SetCerts(v []SSLCert) {
	o.Certs = v
}

func (o SSLCertsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certs != nil {
		toSerialize["certs"] = o.Certs
	}
	return json.Marshal(toSerialize)
}

type NullableSSLCertsList struct {
	value *SSLCertsList
	isSet bool
}

func (v NullableSSLCertsList) Get() *SSLCertsList {
	return v.value
}

func (v *NullableSSLCertsList) Set(val *SSLCertsList) {
	v.value = val
	v.isSet = true
}

func (v NullableSSLCertsList) IsSet() bool {
	return v.isSet
}

func (v *NullableSSLCertsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSLCertsList(val *SSLCertsList) *NullableSSLCertsList {
	return &NullableSSLCertsList{value: val, isSet: true}
}

func (v NullableSSLCertsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSLCertsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



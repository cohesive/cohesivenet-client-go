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

// SSLCertsResponseResponse struct for SSLCertsResponseResponse
type SSLCertsResponseResponse struct {
	Certs []SSLCert `json:"certs,omitempty"`
}

// NewSSLCertsResponseResponse instantiates a new SSLCertsResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSLCertsResponseResponse() *SSLCertsResponseResponse {
	this := SSLCertsResponseResponse{}
	return &this
}

// NewSSLCertsResponseResponseWithDefaults instantiates a new SSLCertsResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSLCertsResponseResponseWithDefaults() *SSLCertsResponseResponse {
	this := SSLCertsResponseResponse{}
	return &this
}

// GetCerts returns the Certs field value if set, zero value otherwise.
func (o *SSLCertsResponseResponse) GetCerts() []SSLCert {
	if o == nil || o.Certs == nil {
		var ret []SSLCert
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLCertsResponseResponse) GetCertsOk() ([]SSLCert, bool) {
	if o == nil || o.Certs == nil {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *SSLCertsResponseResponse) HasCerts() bool {
	if o != nil && o.Certs != nil {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []SSLCert and assigns it to the Certs field.
func (o *SSLCertsResponseResponse) SetCerts(v []SSLCert) {
	o.Certs = v
}

func (o SSLCertsResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certs != nil {
		toSerialize["certs"] = o.Certs
	}
	return json.Marshal(toSerialize)
}

type NullableSSLCertsResponseResponse struct {
	value *SSLCertsResponseResponse
	isSet bool
}

func (v NullableSSLCertsResponseResponse) Get() *SSLCertsResponseResponse {
	return v.value
}

func (v *NullableSSLCertsResponseResponse) Set(val *SSLCertsResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSSLCertsResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSSLCertsResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSLCertsResponseResponse(val *SSLCertsResponseResponse) *NullableSSLCertsResponseResponse {
	return &NullableSSLCertsResponseResponse{value: val, isSet: true}
}

func (v NullableSSLCertsResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSLCertsResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


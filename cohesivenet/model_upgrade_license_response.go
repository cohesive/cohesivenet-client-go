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

// UpgradeLicenseResponse struct for UpgradeLicenseResponse
type UpgradeLicenseResponse struct {
	Response *UpgradeLicenseData `json:"response,omitempty"`
}

// NewUpgradeLicenseResponse instantiates a new UpgradeLicenseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeLicenseResponse() *UpgradeLicenseResponse {
	this := UpgradeLicenseResponse{}
	return &this
}

// NewUpgradeLicenseResponseWithDefaults instantiates a new UpgradeLicenseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeLicenseResponseWithDefaults() *UpgradeLicenseResponse {
	this := UpgradeLicenseResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *UpgradeLicenseResponse) GetResponse() UpgradeLicenseData {
	if o == nil || o.Response == nil {
		var ret UpgradeLicenseData
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeLicenseResponse) GetResponseOk() (*UpgradeLicenseData, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *UpgradeLicenseResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given UpgradeLicenseData and assigns it to the Response field.
func (o *UpgradeLicenseResponse) SetResponse(v UpgradeLicenseData) {
	o.Response = &v
}

func (o UpgradeLicenseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradeLicenseResponse struct {
	value *UpgradeLicenseResponse
	isSet bool
}

func (v NullableUpgradeLicenseResponse) Get() *UpgradeLicenseResponse {
	return v.value
}

func (v *NullableUpgradeLicenseResponse) Set(val *UpgradeLicenseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeLicenseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeLicenseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeLicenseResponse(val *UpgradeLicenseResponse) *NullableUpgradeLicenseResponse {
	return &NullableUpgradeLicenseResponse{value: val, isSet: true}
}

func (v NullableUpgradeLicenseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeLicenseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



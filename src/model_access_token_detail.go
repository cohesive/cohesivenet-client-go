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

// AccessTokenDetail struct for AccessTokenDetail
type AccessTokenDetail struct {
	Response *AccessToken `json:"response,omitempty"`
}

// NewAccessTokenDetail instantiates a new AccessTokenDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenDetail() *AccessTokenDetail {
	this := AccessTokenDetail{}
	return &this
}

// NewAccessTokenDetailWithDefaults instantiates a new AccessTokenDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenDetailWithDefaults() *AccessTokenDetail {
	this := AccessTokenDetail{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AccessTokenDetail) GetResponse() AccessToken {
	if o == nil || o.Response == nil {
		var ret AccessToken
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenDetail) GetResponseOk() (*AccessToken, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AccessTokenDetail) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given AccessToken and assigns it to the Response field.
func (o *AccessTokenDetail) SetResponse(v AccessToken) {
	o.Response = &v
}

func (o AccessTokenDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTokenDetail struct {
	value *AccessTokenDetail
	isSet bool
}

func (v NullableAccessTokenDetail) Get() *AccessTokenDetail {
	return v.value
}

func (v *NullableAccessTokenDetail) Set(val *AccessTokenDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenDetail(val *AccessTokenDetail) *NullableAccessTokenDetail {
	return &NullableAccessTokenDetail{value: val, isSet: true}
}

func (v NullableAccessTokenDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// ExpireRequest struct for ExpireRequest
type ExpireRequest struct {
	Expired *bool `json:"expired,omitempty"`
}

// NewExpireRequest instantiates a new ExpireRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpireRequest() *ExpireRequest {
	this := ExpireRequest{}
	return &this
}

// NewExpireRequestWithDefaults instantiates a new ExpireRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpireRequestWithDefaults() *ExpireRequest {
	this := ExpireRequest{}
	return &this
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ExpireRequest) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpireRequest) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ExpireRequest) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *ExpireRequest) SetExpired(v bool) {
	o.Expired = &v
}

func (o ExpireRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	return json.Marshal(toSerialize)
}

type NullableExpireRequest struct {
	value *ExpireRequest
	isSet bool
}

func (v NullableExpireRequest) Get() *ExpireRequest {
	return v.value
}

func (v *NullableExpireRequest) Set(val *ExpireRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExpireRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExpireRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpireRequest(val *ExpireRequest) *NullableExpireRequest {
	return &NullableExpireRequest{value: val, isSet: true}
}

func (v NullableExpireRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpireRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



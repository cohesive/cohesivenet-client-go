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

// ClientpackTagKeyRequest struct for ClientpackTagKeyRequest
type ClientpackTagKeyRequest struct {
	Key string `json:"key"`
}

// NewClientpackTagKeyRequest instantiates a new ClientpackTagKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackTagKeyRequest(key string) *ClientpackTagKeyRequest {
	this := ClientpackTagKeyRequest{}
	this.Key = key
	return &this
}

// NewClientpackTagKeyRequestWithDefaults instantiates a new ClientpackTagKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackTagKeyRequestWithDefaults() *ClientpackTagKeyRequest {
	this := ClientpackTagKeyRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *ClientpackTagKeyRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ClientpackTagKeyRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ClientpackTagKeyRequest) SetKey(v string) {
	o.Key = v
}

func (o ClientpackTagKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackTagKeyRequest struct {
	value *ClientpackTagKeyRequest
	isSet bool
}

func (v NullableClientpackTagKeyRequest) Get() *ClientpackTagKeyRequest {
	return v.value
}

func (v *NullableClientpackTagKeyRequest) Set(val *ClientpackTagKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackTagKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackTagKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackTagKeyRequest(val *ClientpackTagKeyRequest) *NullableClientpackTagKeyRequest {
	return &NullableClientpackTagKeyRequest{value: val, isSet: true}
}

func (v NullableClientpackTagKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackTagKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



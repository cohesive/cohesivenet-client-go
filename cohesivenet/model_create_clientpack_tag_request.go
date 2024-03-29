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

// CreateClientpackTagRequest struct for CreateClientpackTagRequest
type CreateClientpackTagRequest struct {
	// Alphanumeric characters allowed in snake_case or kebab-case
	Key string `json:"key"`
	Value string `json:"value"`
}

// NewCreateClientpackTagRequest instantiates a new CreateClientpackTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientpackTagRequest(key string, value string) *CreateClientpackTagRequest {
	this := CreateClientpackTagRequest{}
	this.Key = key
	this.Value = value
	return &this
}

// NewCreateClientpackTagRequestWithDefaults instantiates a new CreateClientpackTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientpackTagRequestWithDefaults() *CreateClientpackTagRequest {
	this := CreateClientpackTagRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateClientpackTagRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateClientpackTagRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateClientpackTagRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *CreateClientpackTagRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateClientpackTagRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateClientpackTagRequest) SetValue(v string) {
	o.Value = v
}

func (o CreateClientpackTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClientpackTagRequest struct {
	value *CreateClientpackTagRequest
	isSet bool
}

func (v NullableCreateClientpackTagRequest) Get() *CreateClientpackTagRequest {
	return v.value
}

func (v *NullableCreateClientpackTagRequest) Set(val *CreateClientpackTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientpackTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientpackTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientpackTagRequest(val *CreateClientpackTagRequest) *NullableCreateClientpackTagRequest {
	return &NullableCreateClientpackTagRequest{value: val, isSet: true}
}

func (v NullableCreateClientpackTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientpackTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// SimpleStatusResponse struct for SimpleStatusResponse
type SimpleStatusResponse struct {
	Response *StatusObject `json:"response,omitempty"`
}

// NewSimpleStatusResponse instantiates a new SimpleStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleStatusResponse() *SimpleStatusResponse {
	this := SimpleStatusResponse{}
	return &this
}

// NewSimpleStatusResponseWithDefaults instantiates a new SimpleStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleStatusResponseWithDefaults() *SimpleStatusResponse {
	this := SimpleStatusResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SimpleStatusResponse) GetResponse() StatusObject {
	if o == nil || o.Response == nil {
		var ret StatusObject
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleStatusResponse) GetResponseOk() (*StatusObject, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SimpleStatusResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given StatusObject and assigns it to the Response field.
func (o *SimpleStatusResponse) SetResponse(v StatusObject) {
	o.Response = &v
}

func (o SimpleStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleStatusResponse struct {
	value *SimpleStatusResponse
	isSet bool
}

func (v NullableSimpleStatusResponse) Get() *SimpleStatusResponse {
	return v.value
}

func (v *NullableSimpleStatusResponse) Set(val *SimpleStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleStatusResponse(val *SimpleStatusResponse) *NullableSimpleStatusResponse {
	return &NullableSimpleStatusResponse{value: val, isSet: true}
}

func (v NullableSimpleStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



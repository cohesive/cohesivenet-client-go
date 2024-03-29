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

// ClientpackDetailResponse struct for ClientpackDetailResponse
type ClientpackDetailResponse struct {
	Response *ClientpackNested `json:"response,omitempty"`
}

// NewClientpackDetailResponse instantiates a new ClientpackDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackDetailResponse() *ClientpackDetailResponse {
	this := ClientpackDetailResponse{}
	return &this
}

// NewClientpackDetailResponseWithDefaults instantiates a new ClientpackDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackDetailResponseWithDefaults() *ClientpackDetailResponse {
	this := ClientpackDetailResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ClientpackDetailResponse) GetResponse() ClientpackNested {
	if o == nil || o.Response == nil {
		var ret ClientpackNested
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackDetailResponse) GetResponseOk() (*ClientpackNested, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ClientpackDetailResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ClientpackNested and assigns it to the Response field.
func (o *ClientpackDetailResponse) SetResponse(v ClientpackNested) {
	o.Response = &v
}

func (o ClientpackDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackDetailResponse struct {
	value *ClientpackDetailResponse
	isSet bool
}

func (v NullableClientpackDetailResponse) Get() *ClientpackDetailResponse {
	return v.value
}

func (v *NullableClientpackDetailResponse) Set(val *ClientpackDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackDetailResponse(val *ClientpackDetailResponse) *NullableClientpackDetailResponse {
	return &NullableClientpackDetailResponse{value: val, isSet: true}
}

func (v NullableClientpackDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



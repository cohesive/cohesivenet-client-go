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

// KeyValuePairsListResponse struct for KeyValuePairsListResponse
type KeyValuePairsListResponse struct {
	Response []KeyValuePair `json:"response,omitempty"`
}

// NewKeyValuePairsListResponse instantiates a new KeyValuePairsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyValuePairsListResponse() *KeyValuePairsListResponse {
	this := KeyValuePairsListResponse{}
	return &this
}

// NewKeyValuePairsListResponseWithDefaults instantiates a new KeyValuePairsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyValuePairsListResponseWithDefaults() *KeyValuePairsListResponse {
	this := KeyValuePairsListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *KeyValuePairsListResponse) GetResponse() []KeyValuePair {
	if o == nil || o.Response == nil {
		var ret []KeyValuePair
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyValuePairsListResponse) GetResponseOk() ([]KeyValuePair, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *KeyValuePairsListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []KeyValuePair and assigns it to the Response field.
func (o *KeyValuePairsListResponse) SetResponse(v []KeyValuePair) {
	o.Response = v
}

func (o KeyValuePairsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableKeyValuePairsListResponse struct {
	value *KeyValuePairsListResponse
	isSet bool
}

func (v NullableKeyValuePairsListResponse) Get() *KeyValuePairsListResponse {
	return v.value
}

func (v *NullableKeyValuePairsListResponse) Set(val *KeyValuePairsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyValuePairsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyValuePairsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyValuePairsListResponse(val *KeyValuePairsListResponse) *NullableKeyValuePairsListResponse {
	return &NullableKeyValuePairsListResponse{value: val, isSet: true}
}

func (v NullableKeyValuePairsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyValuePairsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


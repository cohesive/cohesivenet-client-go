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

// SystemInterfaceListResponse struct for SystemInterfaceListResponse
type SystemInterfaceListResponse struct {
	Response []SystemInterface `json:"response,omitempty"`
}

// NewSystemInterfaceListResponse instantiates a new SystemInterfaceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInterfaceListResponse() *SystemInterfaceListResponse {
	this := SystemInterfaceListResponse{}
	return &this
}

// NewSystemInterfaceListResponseWithDefaults instantiates a new SystemInterfaceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInterfaceListResponseWithDefaults() *SystemInterfaceListResponse {
	this := SystemInterfaceListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SystemInterfaceListResponse) GetResponse() []SystemInterface {
	if o == nil || o.Response == nil {
		var ret []SystemInterface
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInterfaceListResponse) GetResponseOk() ([]SystemInterface, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SystemInterfaceListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []SystemInterface and assigns it to the Response field.
func (o *SystemInterfaceListResponse) SetResponse(v []SystemInterface) {
	o.Response = v
}

func (o SystemInterfaceListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInterfaceListResponse struct {
	value *SystemInterfaceListResponse
	isSet bool
}

func (v NullableSystemInterfaceListResponse) Get() *SystemInterfaceListResponse {
	return v.value
}

func (v *NullableSystemInterfaceListResponse) Set(val *SystemInterfaceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInterfaceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInterfaceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInterfaceListResponse(val *SystemInterfaceListResponse) *NullableSystemInterfaceListResponse {
	return &NullableSystemInterfaceListResponse{value: val, isSet: true}
}

func (v NullableSystemInterfaceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInterfaceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


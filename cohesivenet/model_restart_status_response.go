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

// RestartStatusResponse struct for RestartStatusResponse
type RestartStatusResponse struct {
	Response *RestartStatus `json:"response,omitempty"`
}

// NewRestartStatusResponse instantiates a new RestartStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartStatusResponse() *RestartStatusResponse {
	this := RestartStatusResponse{}
	return &this
}

// NewRestartStatusResponseWithDefaults instantiates a new RestartStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartStatusResponseWithDefaults() *RestartStatusResponse {
	this := RestartStatusResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RestartStatusResponse) GetResponse() RestartStatus {
	if o == nil || o.Response == nil {
		var ret RestartStatus
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartStatusResponse) GetResponseOk() (*RestartStatus, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RestartStatusResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given RestartStatus and assigns it to the Response field.
func (o *RestartStatusResponse) SetResponse(v RestartStatus) {
	o.Response = &v
}

func (o RestartStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableRestartStatusResponse struct {
	value *RestartStatusResponse
	isSet bool
}

func (v NullableRestartStatusResponse) Get() *RestartStatusResponse {
	return v.value
}

func (v *NullableRestartStatusResponse) Set(val *RestartStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartStatusResponse(val *RestartStatusResponse) *NullableRestartStatusResponse {
	return &NullableRestartStatusResponse{value: val, isSet: true}
}

func (v NullableRestartStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


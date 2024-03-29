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

// CreateContainerImageResponse struct for CreateContainerImageResponse
type CreateContainerImageResponse struct {
	Response *CreateContainerImageDetail `json:"response,omitempty"`
}

// NewCreateContainerImageResponse instantiates a new CreateContainerImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerImageResponse() *CreateContainerImageResponse {
	this := CreateContainerImageResponse{}
	return &this
}

// NewCreateContainerImageResponseWithDefaults instantiates a new CreateContainerImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerImageResponseWithDefaults() *CreateContainerImageResponse {
	this := CreateContainerImageResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateContainerImageResponse) GetResponse() CreateContainerImageDetail {
	if o == nil || o.Response == nil {
		var ret CreateContainerImageDetail
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerImageResponse) GetResponseOk() (*CreateContainerImageDetail, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateContainerImageResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CreateContainerImageDetail and assigns it to the Response field.
func (o *CreateContainerImageResponse) SetResponse(v CreateContainerImageDetail) {
	o.Response = &v
}

func (o CreateContainerImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableCreateContainerImageResponse struct {
	value *CreateContainerImageResponse
	isSet bool
}

func (v NullableCreateContainerImageResponse) Get() *CreateContainerImageResponse {
	return v.value
}

func (v *NullableCreateContainerImageResponse) Set(val *CreateContainerImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerImageResponse(val *CreateContainerImageResponse) *NullableCreateContainerImageResponse {
	return &NullableCreateContainerImageResponse{value: val, isSet: true}
}

func (v NullableCreateContainerImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



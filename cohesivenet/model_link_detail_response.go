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

// LinkDetailResponse struct for LinkDetailResponse
type LinkDetailResponse struct {
	Response *Link `json:"response,omitempty"`
}

// NewLinkDetailResponse instantiates a new LinkDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDetailResponse() *LinkDetailResponse {
	this := LinkDetailResponse{}
	return &this
}

// NewLinkDetailResponseWithDefaults instantiates a new LinkDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDetailResponseWithDefaults() *LinkDetailResponse {
	this := LinkDetailResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *LinkDetailResponse) GetResponse() Link {
	if o == nil || o.Response == nil {
		var ret Link
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDetailResponse) GetResponseOk() (*Link, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *LinkDetailResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given Link and assigns it to the Response field.
func (o *LinkDetailResponse) SetResponse(v Link) {
	o.Response = &v
}

func (o LinkDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDetailResponse struct {
	value *LinkDetailResponse
	isSet bool
}

func (v NullableLinkDetailResponse) Get() *LinkDetailResponse {
	return v.value
}

func (v *NullableLinkDetailResponse) Set(val *LinkDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDetailResponse(val *LinkDetailResponse) *NullableLinkDetailResponse {
	return &NullableLinkDetailResponse{value: val, isSet: true}
}

func (v NullableLinkDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


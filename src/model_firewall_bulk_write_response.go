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

// FirewallBulkWriteResponse struct for FirewallBulkWriteResponse
type FirewallBulkWriteResponse struct {
	Response *FirewallBulkWriteResponseResponse `json:"response,omitempty"`
}

// NewFirewallBulkWriteResponse instantiates a new FirewallBulkWriteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallBulkWriteResponse() *FirewallBulkWriteResponse {
	this := FirewallBulkWriteResponse{}
	return &this
}

// NewFirewallBulkWriteResponseWithDefaults instantiates a new FirewallBulkWriteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallBulkWriteResponseWithDefaults() *FirewallBulkWriteResponse {
	this := FirewallBulkWriteResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *FirewallBulkWriteResponse) GetResponse() FirewallBulkWriteResponseResponse {
	if o == nil || o.Response == nil {
		var ret FirewallBulkWriteResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallBulkWriteResponse) GetResponseOk() (*FirewallBulkWriteResponseResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *FirewallBulkWriteResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given FirewallBulkWriteResponseResponse and assigns it to the Response field.
func (o *FirewallBulkWriteResponse) SetResponse(v FirewallBulkWriteResponseResponse) {
	o.Response = &v
}

func (o FirewallBulkWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallBulkWriteResponse struct {
	value *FirewallBulkWriteResponse
	isSet bool
}

func (v NullableFirewallBulkWriteResponse) Get() *FirewallBulkWriteResponse {
	return v.value
}

func (v *NullableFirewallBulkWriteResponse) Set(val *FirewallBulkWriteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallBulkWriteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallBulkWriteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallBulkWriteResponse(val *FirewallBulkWriteResponse) *NullableFirewallBulkWriteResponse {
	return &NullableFirewallBulkWriteResponse{value: val, isSet: true}
}

func (v NullableFirewallBulkWriteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallBulkWriteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


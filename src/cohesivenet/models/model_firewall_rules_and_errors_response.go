/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// FirewallRulesAndErrorsResponse struct for FirewallRulesAndErrorsResponse
type FirewallRulesAndErrorsResponse struct {
	Response *FirewallRulesAndErrors `json:"response,omitempty"`
}

// NewFirewallRulesAndErrorsResponse instantiates a new FirewallRulesAndErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRulesAndErrorsResponse() *FirewallRulesAndErrorsResponse {
	this := FirewallRulesAndErrorsResponse{}
	return &this
}

// NewFirewallRulesAndErrorsResponseWithDefaults instantiates a new FirewallRulesAndErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRulesAndErrorsResponseWithDefaults() *FirewallRulesAndErrorsResponse {
	this := FirewallRulesAndErrorsResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *FirewallRulesAndErrorsResponse) GetResponse() FirewallRulesAndErrors {
	if o == nil || o.Response == nil {
		var ret FirewallRulesAndErrors
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRulesAndErrorsResponse) GetResponseOk() (*FirewallRulesAndErrors, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *FirewallRulesAndErrorsResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given FirewallRulesAndErrors and assigns it to the Response field.
func (o *FirewallRulesAndErrorsResponse) SetResponse(v FirewallRulesAndErrors) {
	o.Response = &v
}

func (o FirewallRulesAndErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallRulesAndErrorsResponse struct {
	value *FirewallRulesAndErrorsResponse
	isSet bool
}

func (v NullableFirewallRulesAndErrorsResponse) Get() *FirewallRulesAndErrorsResponse {
	return v.value
}

func (v *NullableFirewallRulesAndErrorsResponse) Set(val *FirewallRulesAndErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRulesAndErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRulesAndErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRulesAndErrorsResponse(val *FirewallRulesAndErrorsResponse) *NullableFirewallRulesAndErrorsResponse {
	return &NullableFirewallRulesAndErrorsResponse{value: val, isSet: true}
}

func (v NullableFirewallRulesAndErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRulesAndErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


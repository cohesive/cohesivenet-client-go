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

// FirewallV1FwsetDeleteRequest struct for FirewallV1FwsetDeleteRequest
type FirewallV1FwsetDeleteRequest struct {
	// Chained firewall rules seperated by \"\\n\"
	Rules *string `json:"rules,omitempty"`
	// Name of the Fwset. Must be valid chain that begins with one of the following: NETS_, PORTS_, LIST_. 
	Name *string `json:"name,omitempty"`
}

// NewFirewallV1FwsetDeleteRequest instantiates a new FirewallV1FwsetDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallV1FwsetDeleteRequest() *FirewallV1FwsetDeleteRequest {
	this := FirewallV1FwsetDeleteRequest{}
	return &this
}

// NewFirewallV1FwsetDeleteRequestWithDefaults instantiates a new FirewallV1FwsetDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallV1FwsetDeleteRequestWithDefaults() *FirewallV1FwsetDeleteRequest {
	this := FirewallV1FwsetDeleteRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *FirewallV1FwsetDeleteRequest) GetRules() string {
	if o == nil || o.Rules == nil {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1FwsetDeleteRequest) GetRulesOk() (*string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *FirewallV1FwsetDeleteRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
func (o *FirewallV1FwsetDeleteRequest) SetRules(v string) {
	o.Rules = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallV1FwsetDeleteRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1FwsetDeleteRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallV1FwsetDeleteRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallV1FwsetDeleteRequest) SetName(v string) {
	o.Name = &v
}

func (o FirewallV1FwsetDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallV1FwsetDeleteRequest struct {
	value *FirewallV1FwsetDeleteRequest
	isSet bool
}

func (v NullableFirewallV1FwsetDeleteRequest) Get() *FirewallV1FwsetDeleteRequest {
	return v.value
}

func (v *NullableFirewallV1FwsetDeleteRequest) Set(val *FirewallV1FwsetDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallV1FwsetDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallV1FwsetDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallV1FwsetDeleteRequest(val *FirewallV1FwsetDeleteRequest) *NullableFirewallV1FwsetDeleteRequest {
	return &NullableFirewallV1FwsetDeleteRequest{value: val, isSet: true}
}

func (v NullableFirewallV1FwsetDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallV1FwsetDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


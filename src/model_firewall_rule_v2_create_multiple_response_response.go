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

// FirewallRuleV2CreateMultipleResponseResponse struct for FirewallRuleV2CreateMultipleResponseResponse
type FirewallRuleV2CreateMultipleResponseResponse struct {
	Rules []FirewallRuleV2 `json:"rules,omitempty"`
	// List of errors for any rules that failed
	Errors []string `json:"errors,omitempty"`
}

// NewFirewallRuleV2CreateMultipleResponseResponse instantiates a new FirewallRuleV2CreateMultipleResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRuleV2CreateMultipleResponseResponse() *FirewallRuleV2CreateMultipleResponseResponse {
	this := FirewallRuleV2CreateMultipleResponseResponse{}
	return &this
}

// NewFirewallRuleV2CreateMultipleResponseResponseWithDefaults instantiates a new FirewallRuleV2CreateMultipleResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleV2CreateMultipleResponseResponseWithDefaults() *FirewallRuleV2CreateMultipleResponseResponse {
	this := FirewallRuleV2CreateMultipleResponseResponse{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *FirewallRuleV2CreateMultipleResponseResponse) GetRules() []FirewallRuleV2 {
	if o == nil || o.Rules == nil {
		var ret []FirewallRuleV2
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleV2CreateMultipleResponseResponse) GetRulesOk() ([]FirewallRuleV2, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *FirewallRuleV2CreateMultipleResponseResponse) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []FirewallRuleV2 and assigns it to the Rules field.
func (o *FirewallRuleV2CreateMultipleResponseResponse) SetRules(v []FirewallRuleV2) {
	o.Rules = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *FirewallRuleV2CreateMultipleResponseResponse) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleV2CreateMultipleResponseResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *FirewallRuleV2CreateMultipleResponseResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *FirewallRuleV2CreateMultipleResponseResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o FirewallRuleV2CreateMultipleResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallRuleV2CreateMultipleResponseResponse struct {
	value *FirewallRuleV2CreateMultipleResponseResponse
	isSet bool
}

func (v NullableFirewallRuleV2CreateMultipleResponseResponse) Get() *FirewallRuleV2CreateMultipleResponseResponse {
	return v.value
}

func (v *NullableFirewallRuleV2CreateMultipleResponseResponse) Set(val *FirewallRuleV2CreateMultipleResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRuleV2CreateMultipleResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRuleV2CreateMultipleResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRuleV2CreateMultipleResponseResponse(val *FirewallRuleV2CreateMultipleResponseResponse) *NullableFirewallRuleV2CreateMultipleResponseResponse {
	return &NullableFirewallRuleV2CreateMultipleResponseResponse{value: val, isSet: true}
}

func (v NullableFirewallRuleV2CreateMultipleResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRuleV2CreateMultipleResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// FirewallSubtableCreateResponseResponseAllOf struct for FirewallSubtableCreateResponseResponseAllOf
type FirewallSubtableCreateResponseResponseAllOf struct {
	// Rule objects that failed insertion into subtable
	Errors []FirewallSubtableCreateResponseResponseAllOfErrorsInner `json:"errors,omitempty"`
}

// NewFirewallSubtableCreateResponseResponseAllOf instantiates a new FirewallSubtableCreateResponseResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSubtableCreateResponseResponseAllOf() *FirewallSubtableCreateResponseResponseAllOf {
	this := FirewallSubtableCreateResponseResponseAllOf{}
	return &this
}

// NewFirewallSubtableCreateResponseResponseAllOfWithDefaults instantiates a new FirewallSubtableCreateResponseResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSubtableCreateResponseResponseAllOfWithDefaults() *FirewallSubtableCreateResponseResponseAllOf {
	this := FirewallSubtableCreateResponseResponseAllOf{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *FirewallSubtableCreateResponseResponseAllOf) GetErrors() []FirewallSubtableCreateResponseResponseAllOfErrorsInner {
	if o == nil || o.Errors == nil {
		var ret []FirewallSubtableCreateResponseResponseAllOfErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableCreateResponseResponseAllOf) GetErrorsOk() ([]FirewallSubtableCreateResponseResponseAllOfErrorsInner, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *FirewallSubtableCreateResponseResponseAllOf) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []FirewallSubtableCreateResponseResponseAllOfErrorsInner and assigns it to the Errors field.
func (o *FirewallSubtableCreateResponseResponseAllOf) SetErrors(v []FirewallSubtableCreateResponseResponseAllOfErrorsInner) {
	o.Errors = v
}

func (o FirewallSubtableCreateResponseResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallSubtableCreateResponseResponseAllOf struct {
	value *FirewallSubtableCreateResponseResponseAllOf
	isSet bool
}

func (v NullableFirewallSubtableCreateResponseResponseAllOf) Get() *FirewallSubtableCreateResponseResponseAllOf {
	return v.value
}

func (v *NullableFirewallSubtableCreateResponseResponseAllOf) Set(val *FirewallSubtableCreateResponseResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSubtableCreateResponseResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSubtableCreateResponseResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSubtableCreateResponseResponseAllOf(val *FirewallSubtableCreateResponseResponseAllOf) *NullableFirewallSubtableCreateResponseResponseAllOf {
	return &NullableFirewallSubtableCreateResponseResponseAllOf{value: val, isSet: true}
}

func (v NullableFirewallSubtableCreateResponseResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSubtableCreateResponseResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


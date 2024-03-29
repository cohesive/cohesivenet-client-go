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

// RuleGroup struct for RuleGroup
type RuleGroup struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// List of rule IDs in group
	Rules []string `json:"rules,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewRuleGroup instantiates a new RuleGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleGroup() *RuleGroup {
	this := RuleGroup{}
	return &this
}

// NewRuleGroupWithDefaults instantiates a new RuleGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleGroupWithDefaults() *RuleGroup {
	this := RuleGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleGroup) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *RuleGroup) GetRules() []string {
	if o == nil || o.Rules == nil {
		var ret []string
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroup) GetRulesOk() ([]string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *RuleGroup) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []string and assigns it to the Rules field.
func (o *RuleGroup) SetRules(v []string) {
	o.Rules = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RuleGroup) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroup) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RuleGroup) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RuleGroup) SetSize(v int32) {
	o.Size = &v
}

func (o RuleGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableRuleGroup struct {
	value *RuleGroup
	isSet bool
}

func (v NullableRuleGroup) Get() *RuleGroup {
	return v.value
}

func (v *NullableRuleGroup) Set(val *RuleGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleGroup(val *RuleGroup) *NullableRuleGroup {
	return &NullableRuleGroup{value: val, isSet: true}
}

func (v NullableRuleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



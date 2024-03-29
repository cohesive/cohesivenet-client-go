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

// CreateFirewallSubtableRequestRule struct for CreateFirewallSubtableRequestRule
type CreateFirewallSubtableRequestRule struct {
	Rule string `json:"rule"`
	Position int32 `json:"position"`
	Comment *string `json:"comment,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
}

// NewCreateFirewallSubtableRequestRule instantiates a new CreateFirewallSubtableRequestRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirewallSubtableRequestRule(rule string, position int32) *CreateFirewallSubtableRequestRule {
	this := CreateFirewallSubtableRequestRule{}
	this.Rule = rule
	this.Position = position
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewCreateFirewallSubtableRequestRuleWithDefaults instantiates a new CreateFirewallSubtableRequestRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirewallSubtableRequestRuleWithDefaults() *CreateFirewallSubtableRequestRule {
	this := CreateFirewallSubtableRequestRule{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetRule returns the Rule field value
func (o *CreateFirewallSubtableRequestRule) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallSubtableRequestRule) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *CreateFirewallSubtableRequestRule) SetRule(v string) {
	o.Rule = v
}

// GetPosition returns the Position field value
func (o *CreateFirewallSubtableRequestRule) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallSubtableRequestRule) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *CreateFirewallSubtableRequestRule) SetPosition(v int32) {
	o.Position = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateFirewallSubtableRequestRule) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirewallSubtableRequestRule) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateFirewallSubtableRequestRule) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateFirewallSubtableRequestRule) SetComment(v string) {
	o.Comment = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CreateFirewallSubtableRequestRule) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirewallSubtableRequestRule) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CreateFirewallSubtableRequestRule) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *CreateFirewallSubtableRequestRule) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o CreateFirewallSubtableRequestRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rule"] = o.Rule
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFirewallSubtableRequestRule struct {
	value *CreateFirewallSubtableRequestRule
	isSet bool
}

func (v NullableCreateFirewallSubtableRequestRule) Get() *CreateFirewallSubtableRequestRule {
	return v.value
}

func (v *NullableCreateFirewallSubtableRequestRule) Set(val *CreateFirewallSubtableRequestRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirewallSubtableRequestRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirewallSubtableRequestRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirewallSubtableRequestRule(val *CreateFirewallSubtableRequestRule) *NullableCreateFirewallSubtableRequestRule {
	return &NullableCreateFirewallSubtableRequestRule{value: val, isSet: true}
}

func (v NullableCreateFirewallSubtableRequestRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirewallSubtableRequestRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// AddFirewallRuleToSubtableRequest struct for AddFirewallRuleToSubtableRequest
type AddFirewallRuleToSubtableRequest struct {
	Rule string `json:"rule"`
	// Position of rule in subtable. -1 indicates the end of the subtable
	Position *int32 `json:"position,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
}

// NewAddFirewallRuleToSubtableRequest instantiates a new AddFirewallRuleToSubtableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFirewallRuleToSubtableRequest(rule string) *AddFirewallRuleToSubtableRequest {
	this := AddFirewallRuleToSubtableRequest{}
	this.Rule = rule
	var position int32 = -1
	this.Position = &position
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewAddFirewallRuleToSubtableRequestWithDefaults instantiates a new AddFirewallRuleToSubtableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFirewallRuleToSubtableRequestWithDefaults() *AddFirewallRuleToSubtableRequest {
	this := AddFirewallRuleToSubtableRequest{}
	var position int32 = -1
	this.Position = &position
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetRule returns the Rule field value
func (o *AddFirewallRuleToSubtableRequest) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *AddFirewallRuleToSubtableRequest) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *AddFirewallRuleToSubtableRequest) SetRule(v string) {
	o.Rule = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *AddFirewallRuleToSubtableRequest) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFirewallRuleToSubtableRequest) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *AddFirewallRuleToSubtableRequest) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *AddFirewallRuleToSubtableRequest) SetPosition(v int32) {
	o.Position = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AddFirewallRuleToSubtableRequest) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFirewallRuleToSubtableRequest) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AddFirewallRuleToSubtableRequest) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AddFirewallRuleToSubtableRequest) SetComment(v string) {
	o.Comment = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AddFirewallRuleToSubtableRequest) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFirewallRuleToSubtableRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AddFirewallRuleToSubtableRequest) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AddFirewallRuleToSubtableRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o AddFirewallRuleToSubtableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rule"] = o.Rule
	}
	if o.Position != nil {
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

type NullableAddFirewallRuleToSubtableRequest struct {
	value *AddFirewallRuleToSubtableRequest
	isSet bool
}

func (v NullableAddFirewallRuleToSubtableRequest) Get() *AddFirewallRuleToSubtableRequest {
	return v.value
}

func (v *NullableAddFirewallRuleToSubtableRequest) Set(val *AddFirewallRuleToSubtableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFirewallRuleToSubtableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFirewallRuleToSubtableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFirewallRuleToSubtableRequest(val *AddFirewallRuleToSubtableRequest) *NullableAddFirewallRuleToSubtableRequest {
	return &NullableAddFirewallRuleToSubtableRequest{value: val, isSet: true}
}

func (v NullableAddFirewallRuleToSubtableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFirewallRuleToSubtableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



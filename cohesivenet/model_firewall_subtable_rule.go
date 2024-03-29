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

// FirewallSubtableRule struct for FirewallSubtableRule
type FirewallSubtableRule struct {
	Rule *string `json:"rule,omitempty"`
	RuleResolved *string `json:"rule_resolved,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Id *string `json:"id,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	OsRules []FirewallSubtableRuleOsRule `json:"os_rules,omitempty"`
}

// NewFirewallSubtableRule instantiates a new FirewallSubtableRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSubtableRule() *FirewallSubtableRule {
	this := FirewallSubtableRule{}
	return &this
}

// NewFirewallSubtableRuleWithDefaults instantiates a new FirewallSubtableRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSubtableRuleWithDefaults() *FirewallSubtableRule {
	this := FirewallSubtableRule{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *FirewallSubtableRule) SetRule(v string) {
	o.Rule = &v
}

// GetRuleResolved returns the RuleResolved field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetRuleResolved() string {
	if o == nil || o.RuleResolved == nil {
		var ret string
		return ret
	}
	return *o.RuleResolved
}

// GetRuleResolvedOk returns a tuple with the RuleResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetRuleResolvedOk() (*string, bool) {
	if o == nil || o.RuleResolved == nil {
		return nil, false
	}
	return o.RuleResolved, true
}

// HasRuleResolved returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasRuleResolved() bool {
	if o != nil && o.RuleResolved != nil {
		return true
	}

	return false
}

// SetRuleResolved gets a reference to the given string and assigns it to the RuleResolved field.
func (o *FirewallSubtableRule) SetRuleResolved(v string) {
	o.RuleResolved = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *FirewallSubtableRule) SetPosition(v int32) {
	o.Position = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *FirewallSubtableRule) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FirewallSubtableRule) SetId(v string) {
	o.Id = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *FirewallSubtableRule) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetOsRules returns the OsRules field value if set, zero value otherwise.
func (o *FirewallSubtableRule) GetOsRules() []FirewallSubtableRuleOsRule {
	if o == nil || o.OsRules == nil {
		var ret []FirewallSubtableRuleOsRule
		return ret
	}
	return o.OsRules
}

// GetOsRulesOk returns a tuple with the OsRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableRule) GetOsRulesOk() ([]FirewallSubtableRuleOsRule, bool) {
	if o == nil || o.OsRules == nil {
		return nil, false
	}
	return o.OsRules, true
}

// HasOsRules returns a boolean if a field has been set.
func (o *FirewallSubtableRule) HasOsRules() bool {
	if o != nil && o.OsRules != nil {
		return true
	}

	return false
}

// SetOsRules gets a reference to the given []FirewallSubtableRuleOsRule and assigns it to the OsRules field.
func (o *FirewallSubtableRule) SetOsRules(v []FirewallSubtableRuleOsRule) {
	o.OsRules = v
}

func (o FirewallSubtableRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.RuleResolved != nil {
		toSerialize["rule_resolved"] = o.RuleResolved
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.OsRules != nil {
		toSerialize["os_rules"] = o.OsRules
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallSubtableRule struct {
	value *FirewallSubtableRule
	isSet bool
}

func (v NullableFirewallSubtableRule) Get() *FirewallSubtableRule {
	return v.value
}

func (v *NullableFirewallSubtableRule) Set(val *FirewallSubtableRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSubtableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSubtableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSubtableRule(val *FirewallSubtableRule) *NullableFirewallSubtableRule {
	return &NullableFirewallSubtableRule{value: val, isSet: true}
}

func (v NullableFirewallSubtableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSubtableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



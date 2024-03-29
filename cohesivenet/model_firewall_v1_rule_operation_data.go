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

// FirewallV1RuleOperationData struct for FirewallV1RuleOperationData
type FirewallV1RuleOperationData struct {
	Status *string `json:"status,omitempty"`
	Rule *string `json:"rule,omitempty"`
	Position *int32 `json:"position,omitempty"`
	// Task token
	Token *string `json:"token,omitempty"`
}

// NewFirewallV1RuleOperationData instantiates a new FirewallV1RuleOperationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallV1RuleOperationData() *FirewallV1RuleOperationData {
	this := FirewallV1RuleOperationData{}
	return &this
}

// NewFirewallV1RuleOperationDataWithDefaults instantiates a new FirewallV1RuleOperationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallV1RuleOperationDataWithDefaults() *FirewallV1RuleOperationData {
	this := FirewallV1RuleOperationData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallV1RuleOperationData) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1RuleOperationData) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallV1RuleOperationData) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FirewallV1RuleOperationData) SetStatus(v string) {
	o.Status = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FirewallV1RuleOperationData) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1RuleOperationData) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FirewallV1RuleOperationData) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *FirewallV1RuleOperationData) SetRule(v string) {
	o.Rule = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FirewallV1RuleOperationData) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1RuleOperationData) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FirewallV1RuleOperationData) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *FirewallV1RuleOperationData) SetPosition(v int32) {
	o.Position = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FirewallV1RuleOperationData) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1RuleOperationData) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FirewallV1RuleOperationData) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FirewallV1RuleOperationData) SetToken(v string) {
	o.Token = &v
}

func (o FirewallV1RuleOperationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallV1RuleOperationData struct {
	value *FirewallV1RuleOperationData
	isSet bool
}

func (v NullableFirewallV1RuleOperationData) Get() *FirewallV1RuleOperationData {
	return v.value
}

func (v *NullableFirewallV1RuleOperationData) Set(val *FirewallV1RuleOperationData) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallV1RuleOperationData) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallV1RuleOperationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallV1RuleOperationData(val *FirewallV1RuleOperationData) *NullableFirewallV1RuleOperationData {
	return &NullableFirewallV1RuleOperationData{value: val, isSet: true}
}

func (v NullableFirewallV1RuleOperationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallV1RuleOperationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



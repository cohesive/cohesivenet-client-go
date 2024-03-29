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

// ImportFirewallRulesRequest struct for ImportFirewallRulesRequest
type ImportFirewallRulesRequest struct {
	// Rules as string. This can be valid json for a rule object or rules delimited by newlines
	Rules string `json:"rules"`
	// Starting position for imported rules. -1 indicates end of firewall.
	Position *int32 `json:"position,omitempty"`
	// Import all rules as immediately disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Groups to add these imported rules to
	Groups []string `json:"groups,omitempty"`
}

// NewImportFirewallRulesRequest instantiates a new ImportFirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportFirewallRulesRequest(rules string) *ImportFirewallRulesRequest {
	this := ImportFirewallRulesRequest{}
	this.Rules = rules
	var position int32 = -1
	this.Position = &position
	return &this
}

// NewImportFirewallRulesRequestWithDefaults instantiates a new ImportFirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportFirewallRulesRequestWithDefaults() *ImportFirewallRulesRequest {
	this := ImportFirewallRulesRequest{}
	var position int32 = -1
	this.Position = &position
	return &this
}

// GetRules returns the Rules field value
func (o *ImportFirewallRulesRequest) GetRules() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ImportFirewallRulesRequest) GetRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *ImportFirewallRulesRequest) SetRules(v string) {
	o.Rules = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ImportFirewallRulesRequest) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFirewallRulesRequest) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ImportFirewallRulesRequest) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ImportFirewallRulesRequest) SetPosition(v int32) {
	o.Position = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ImportFirewallRulesRequest) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFirewallRulesRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ImportFirewallRulesRequest) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ImportFirewallRulesRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ImportFirewallRulesRequest) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFirewallRulesRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ImportFirewallRulesRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ImportFirewallRulesRequest) SetGroups(v []string) {
	o.Groups = v
}

func (o ImportFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableImportFirewallRulesRequest struct {
	value *ImportFirewallRulesRequest
	isSet bool
}

func (v NullableImportFirewallRulesRequest) Get() *ImportFirewallRulesRequest {
	return v.value
}

func (v *NullableImportFirewallRulesRequest) Set(val *ImportFirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportFirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportFirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportFirewallRulesRequest(val *ImportFirewallRulesRequest) *NullableImportFirewallRulesRequest {
	return &NullableImportFirewallRulesRequest{value: val, isSet: true}
}

func (v NullableImportFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportFirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
	"time"
)

// FirewallSubtableWithErrors struct for FirewallSubtableWithErrors
type FirewallSubtableWithErrors struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// table this subtable can be routed to from such as prerouting
	Type *string `json:"type,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Rules []FirewallSubtableRule `json:"rules,omitempty"`
	Errors []FirewallSubtableRuleWithError `json:"errors,omitempty"`
}

// NewFirewallSubtableWithErrors instantiates a new FirewallSubtableWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSubtableWithErrors() *FirewallSubtableWithErrors {
	this := FirewallSubtableWithErrors{}
	return &this
}

// NewFirewallSubtableWithErrorsWithDefaults instantiates a new FirewallSubtableWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSubtableWithErrorsWithDefaults() *FirewallSubtableWithErrors {
	this := FirewallSubtableWithErrors{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallSubtableWithErrors) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirewallSubtableWithErrors) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FirewallSubtableWithErrors) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FirewallSubtableWithErrors) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FirewallSubtableWithErrors) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *FirewallSubtableWithErrors) SetSize(v int32) {
	o.Size = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetRules() []FirewallSubtableRule {
	if o == nil || o.Rules == nil {
		var ret []FirewallSubtableRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetRulesOk() ([]FirewallSubtableRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []FirewallSubtableRule and assigns it to the Rules field.
func (o *FirewallSubtableWithErrors) SetRules(v []FirewallSubtableRule) {
	o.Rules = v
}

///


// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *FirewallSubtableWithErrors) GetErrors() []FirewallSubtableRuleWithError {
	if o == nil || o.Errors == nil {
		var ret []FirewallSubtableRuleWithError
		return ret
	}
	return o.Errors
}

// GetErrorssOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSubtableWithErrors) GetErrorsOk() ([]FirewallSubtableRuleWithError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *FirewallSubtableWithErrors) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []FirewallSubtableRuleWithError and assigns it to the Rules field.
func (o *FirewallSubtableWithErrors) SetErrors(v []FirewallSubtableRuleWithError) {
	o.Errors = v
}

func (o FirewallSubtableWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallSubtableWithErrors struct {
	value *FirewallSubtableWithErrors
	isSet bool
}

func (v NullableFirewallSubtableWithErrors) Get() *FirewallSubtableWithErrors {
	return v.value
}

func (v *NullableFirewallSubtableWithErrors) Set(val *FirewallSubtableWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSubtableWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSubtableWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSubtableWithErrors(val *FirewallSubtableWithErrors) *NullableFirewallSubtableWithErrors {
	return &NullableFirewallSubtableWithErrors{value: val, isSet: true}
}

func (v NullableFirewallSubtableWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSubtableWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



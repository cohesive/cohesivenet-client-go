/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

// FirewallFwset struct for FirewallFwset
type FirewallFwset struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Valid ipset type. Supported are net, port, list and service
	Type *string `json:"type,omitempty"`
	// number of entries in set
	Size *int32 `json:"size,omitempty"`
	Entries []FirewallFwsetEntry `json:"entries,omitempty"`
}

// NewFirewallFwset instantiates a new FirewallFwset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallFwset() *FirewallFwset {
	this := FirewallFwset{}
	return &this
}

// NewFirewallFwsetWithDefaults instantiates a new FirewallFwset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallFwsetWithDefaults() *FirewallFwset {
	this := FirewallFwset{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallFwset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallFwset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallFwset) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirewallFwset) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirewallFwset) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirewallFwset) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FirewallFwset) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FirewallFwset) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FirewallFwset) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FirewallFwset) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FirewallFwset) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FirewallFwset) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FirewallFwset) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FirewallFwset) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FirewallFwset) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FirewallFwset) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FirewallFwset) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *FirewallFwset) SetSize(v int32) {
	o.Size = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *FirewallFwset) GetEntries() []FirewallFwsetEntry {
	if o == nil || o.Entries == nil {
		var ret []FirewallFwsetEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallFwset) GetEntriesOk() ([]FirewallFwsetEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *FirewallFwset) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []FirewallFwsetEntry and assigns it to the Entries field.
func (o *FirewallFwset) SetEntries(v []FirewallFwsetEntry) {
	o.Entries = v
}

func (o FirewallFwset) MarshalJSON() ([]byte, error) {
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
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallFwset struct {
	value *FirewallFwset
	isSet bool
}

func (v NullableFirewallFwset) Get() *FirewallFwset {
	return v.value
}

func (v *NullableFirewallFwset) Set(val *FirewallFwset) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallFwset) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallFwset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallFwset(val *FirewallFwset) *NullableFirewallFwset {
	return &NullableFirewallFwset{value: val, isSet: true}
}

func (v NullableFirewallFwset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallFwset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


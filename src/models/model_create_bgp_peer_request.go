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
)

// CreateBGPPeerRequest struct for CreateBGPPeerRequest
type CreateBGPPeerRequest struct {
	// IP address of the desired BGP peer.
	Ipaddress string `json:"ipaddress"`
	// Autonomous system number assigned to device at ipaddress
	Asn int32 `json:"asn"`
	// Autonomous system number alias
	LocalAsnAlias *int32 `json:"local_asn_alias,omitempty"`
	// List of \"in permit CIDR\" and/or \"out permit CIDR\" statements in a string delimited by \"\\n\"
	AccessList *string `json:"access_list,omitempty"`
	// String to be agreed upon by both peers as a simple password.
	BgpPassword *string `json:"bgp_password,omitempty"`
	// Enable network distance for BGP peer
	AddNetworkDistance *bool `json:"add_network_distance,omitempty"`
	// Add distance direction for BGP peer, in or out
	AddNetworkDistanceDirection *string `json:"add_network_distance_direction,omitempty"`
	// Distance metric weight indicating distance in hops for BGP peer
	AddNetworkDistanceHops *int32 `json:"add_network_distance_hops,omitempty"`
	HoldTime *int32 `json:"hold_time,omitempty"`
	// Distance metric weight indicating distance in hops for BGP peer
	KeepaliveInterval *int32 `json:"keepalive_interval,omitempty"`
}

// NewCreateBGPPeerRequest instantiates a new CreateBGPPeerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBGPPeerRequest(ipaddress string, asn int32) *CreateBGPPeerRequest {
	this := CreateBGPPeerRequest{}
	this.Ipaddress = ipaddress
	this.Asn = asn
	return &this
}

// NewCreateBGPPeerRequestWithDefaults instantiates a new CreateBGPPeerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBGPPeerRequestWithDefaults() *CreateBGPPeerRequest {
	this := CreateBGPPeerRequest{}
	return &this
}

// GetIpaddress returns the Ipaddress field value
func (o *CreateBGPPeerRequest) GetIpaddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetIpaddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipaddress, true
}

// SetIpaddress sets field value
func (o *CreateBGPPeerRequest) SetIpaddress(v string) {
	o.Ipaddress = v
}

// GetAsn returns the Asn field value
func (o *CreateBGPPeerRequest) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *CreateBGPPeerRequest) SetAsn(v int32) {
	o.Asn = v
}

// GetLocalAsnAlias returns the LocalAsnAlias field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetLocalAsnAlias() int32 {
	if o == nil || o.LocalAsnAlias == nil {
		var ret int32
		return ret
	}
	return *o.LocalAsnAlias
}

// GetLocalAsnAliasOk returns a tuple with the LocalAsnAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetLocalAsnAliasOk() (*int32, bool) {
	if o == nil || o.LocalAsnAlias == nil {
		return nil, false
	}
	return o.LocalAsnAlias, true
}

// HasLocalAsnAlias returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasLocalAsnAlias() bool {
	if o != nil && o.LocalAsnAlias != nil {
		return true
	}

	return false
}

// SetLocalAsnAlias gets a reference to the given int32 and assigns it to the LocalAsnAlias field.
func (o *CreateBGPPeerRequest) SetLocalAsnAlias(v int32) {
	o.LocalAsnAlias = &v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetAccessList() string {
	if o == nil || o.AccessList == nil {
		var ret string
		return ret
	}
	return *o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetAccessListOk() (*string, bool) {
	if o == nil || o.AccessList == nil {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasAccessList() bool {
	if o != nil && o.AccessList != nil {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given string and assigns it to the AccessList field.
func (o *CreateBGPPeerRequest) SetAccessList(v string) {
	o.AccessList = &v
}

// GetBgpPassword returns the BgpPassword field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetBgpPassword() string {
	if o == nil || o.BgpPassword == nil {
		var ret string
		return ret
	}
	return *o.BgpPassword
}

// GetBgpPasswordOk returns a tuple with the BgpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetBgpPasswordOk() (*string, bool) {
	if o == nil || o.BgpPassword == nil {
		return nil, false
	}
	return o.BgpPassword, true
}

// HasBgpPassword returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasBgpPassword() bool {
	if o != nil && o.BgpPassword != nil {
		return true
	}

	return false
}

// SetBgpPassword gets a reference to the given string and assigns it to the BgpPassword field.
func (o *CreateBGPPeerRequest) SetBgpPassword(v string) {
	o.BgpPassword = &v
}

// GetAddNetworkDistance returns the AddNetworkDistance field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetAddNetworkDistance() bool {
	if o == nil || o.AddNetworkDistance == nil {
		var ret bool
		return ret
	}
	return *o.AddNetworkDistance
}

// GetAddNetworkDistanceOk returns a tuple with the AddNetworkDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetAddNetworkDistanceOk() (*bool, bool) {
	if o == nil || o.AddNetworkDistance == nil {
		return nil, false
	}
	return o.AddNetworkDistance, true
}

// HasAddNetworkDistance returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasAddNetworkDistance() bool {
	if o != nil && o.AddNetworkDistance != nil {
		return true
	}

	return false
}

// SetAddNetworkDistance gets a reference to the given bool and assigns it to the AddNetworkDistance field.
func (o *CreateBGPPeerRequest) SetAddNetworkDistance(v bool) {
	o.AddNetworkDistance = &v
}

// GetAddNetworkDistanceDirection returns the AddNetworkDistanceDirection field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetAddNetworkDistanceDirection() string {
	if o == nil || o.AddNetworkDistanceDirection == nil {
		var ret string
		return ret
	}
	return *o.AddNetworkDistanceDirection
}

// GetAddNetworkDistanceDirectionOk returns a tuple with the AddNetworkDistanceDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetAddNetworkDistanceDirectionOk() (*string, bool) {
	if o == nil || o.AddNetworkDistanceDirection == nil {
		return nil, false
	}
	return o.AddNetworkDistanceDirection, true
}

// HasAddNetworkDistanceDirection returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasAddNetworkDistanceDirection() bool {
	if o != nil && o.AddNetworkDistanceDirection != nil {
		return true
	}

	return false
}

// SetAddNetworkDistanceDirection gets a reference to the given string and assigns it to the AddNetworkDistanceDirection field.
func (o *CreateBGPPeerRequest) SetAddNetworkDistanceDirection(v string) {
	o.AddNetworkDistanceDirection = &v
}

// GetAddNetworkDistanceHops returns the AddNetworkDistanceHops field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetAddNetworkDistanceHops() int32 {
	if o == nil || o.AddNetworkDistanceHops == nil {
		var ret int32
		return ret
	}
	return *o.AddNetworkDistanceHops
}

// GetAddNetworkDistanceHopsOk returns a tuple with the AddNetworkDistanceHops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetAddNetworkDistanceHopsOk() (*int32, bool) {
	if o == nil || o.AddNetworkDistanceHops == nil {
		return nil, false
	}
	return o.AddNetworkDistanceHops, true
}

// HasAddNetworkDistanceHops returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasAddNetworkDistanceHops() bool {
	if o != nil && o.AddNetworkDistanceHops != nil {
		return true
	}

	return false
}

// SetAddNetworkDistanceHops gets a reference to the given int32 and assigns it to the AddNetworkDistanceHops field.
func (o *CreateBGPPeerRequest) SetAddNetworkDistanceHops(v int32) {
	o.AddNetworkDistanceHops = &v
}

// GetHoldTime returns the HoldTime field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetHoldTime() int32 {
	if o == nil || o.HoldTime == nil {
		var ret int32
		return ret
	}
	return *o.HoldTime
}

// GetHoldTimeOk returns a tuple with the HoldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetHoldTimeOk() (*int32, bool) {
	if o == nil || o.HoldTime == nil {
		return nil, false
	}
	return o.HoldTime, true
}

// HasHoldTime returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasHoldTime() bool {
	if o != nil && o.HoldTime != nil {
		return true
	}

	return false
}

// SetHoldTime gets a reference to the given int32 and assigns it to the HoldTime field.
func (o *CreateBGPPeerRequest) SetHoldTime(v int32) {
	o.HoldTime = &v
}

// GetKeepaliveInterval returns the KeepaliveInterval field value if set, zero value otherwise.
func (o *CreateBGPPeerRequest) GetKeepaliveInterval() int32 {
	if o == nil || o.KeepaliveInterval == nil {
		var ret int32
		return ret
	}
	return *o.KeepaliveInterval
}

// GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBGPPeerRequest) GetKeepaliveIntervalOk() (*int32, bool) {
	if o == nil || o.KeepaliveInterval == nil {
		return nil, false
	}
	return o.KeepaliveInterval, true
}

// HasKeepaliveInterval returns a boolean if a field has been set.
func (o *CreateBGPPeerRequest) HasKeepaliveInterval() bool {
	if o != nil && o.KeepaliveInterval != nil {
		return true
	}

	return false
}

// SetKeepaliveInterval gets a reference to the given int32 and assigns it to the KeepaliveInterval field.
func (o *CreateBGPPeerRequest) SetKeepaliveInterval(v int32) {
	o.KeepaliveInterval = &v
}

func (o CreateBGPPeerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if true {
		toSerialize["asn"] = o.Asn
	}
	if o.LocalAsnAlias != nil {
		toSerialize["local_asn_alias"] = o.LocalAsnAlias
	}
	if o.AccessList != nil {
		toSerialize["access_list"] = o.AccessList
	}
	if o.BgpPassword != nil {
		toSerialize["bgp_password"] = o.BgpPassword
	}
	if o.AddNetworkDistance != nil {
		toSerialize["add_network_distance"] = o.AddNetworkDistance
	}
	if o.AddNetworkDistanceDirection != nil {
		toSerialize["add_network_distance_direction"] = o.AddNetworkDistanceDirection
	}
	if o.AddNetworkDistanceHops != nil {
		toSerialize["add_network_distance_hops"] = o.AddNetworkDistanceHops
	}
	if o.HoldTime != nil {
		toSerialize["hold_time"] = o.HoldTime
	}
	if o.KeepaliveInterval != nil {
		toSerialize["keepalive_interval"] = o.KeepaliveInterval
	}
	return json.Marshal(toSerialize)
}

type NullableCreateBGPPeerRequest struct {
	value *CreateBGPPeerRequest
	isSet bool
}

func (v NullableCreateBGPPeerRequest) Get() *CreateBGPPeerRequest {
	return v.value
}

func (v *NullableCreateBGPPeerRequest) Set(val *CreateBGPPeerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBGPPeerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBGPPeerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBGPPeerRequest(val *CreateBGPPeerRequest) *NullableCreateBGPPeerRequest {
	return &NullableCreateBGPPeerRequest{value: val, isSet: true}
}

func (v NullableCreateBGPPeerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBGPPeerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


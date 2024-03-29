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

// FirewallV1SubgroupListResponse struct for FirewallV1SubgroupListResponse
type FirewallV1SubgroupListResponse struct {
	Response *FirewallSubgroupsData `json:"response,omitempty"`
}

// NewFirewallV1SubgroupListResponse instantiates a new FirewallV1SubgroupListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallV1SubgroupListResponse() *FirewallV1SubgroupListResponse {
	this := FirewallV1SubgroupListResponse{}
	return &this
}

// NewFirewallV1SubgroupListResponseWithDefaults instantiates a new FirewallV1SubgroupListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallV1SubgroupListResponseWithDefaults() *FirewallV1SubgroupListResponse {
	this := FirewallV1SubgroupListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *FirewallV1SubgroupListResponse) GetResponse() FirewallSubgroupsData {
	if o == nil || o.Response == nil {
		var ret FirewallSubgroupsData
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallV1SubgroupListResponse) GetResponseOk() (*FirewallSubgroupsData, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *FirewallV1SubgroupListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given FirewallSubgroupsData and assigns it to the Response field.
func (o *FirewallV1SubgroupListResponse) SetResponse(v FirewallSubgroupsData) {
	o.Response = &v
}

func (o FirewallV1SubgroupListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallV1SubgroupListResponse struct {
	value *FirewallV1SubgroupListResponse
	isSet bool
}

func (v NullableFirewallV1SubgroupListResponse) Get() *FirewallV1SubgroupListResponse {
	return v.value
}

func (v *NullableFirewallV1SubgroupListResponse) Set(val *FirewallV1SubgroupListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallV1SubgroupListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallV1SubgroupListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallV1SubgroupListResponse(val *FirewallV1SubgroupListResponse) *NullableFirewallV1SubgroupListResponse {
	return &NullableFirewallV1SubgroupListResponse{value: val, isSet: true}
}

func (v NullableFirewallV1SubgroupListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallV1SubgroupListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



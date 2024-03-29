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

// ClientpackStatus struct for ClientpackStatus
type ClientpackStatus struct {
	Disconnecting *string `json:"disconnecting,omitempty"`
	OverlayIpaddress *string `json:"overlay_ipaddress,omitempty"`
}

// NewClientpackStatus instantiates a new ClientpackStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientpackStatus() *ClientpackStatus {
	this := ClientpackStatus{}
	return &this
}

// NewClientpackStatusWithDefaults instantiates a new ClientpackStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientpackStatusWithDefaults() *ClientpackStatus {
	this := ClientpackStatus{}
	return &this
}

// GetDisconnecting returns the Disconnecting field value if set, zero value otherwise.
func (o *ClientpackStatus) GetDisconnecting() string {
	if o == nil || o.Disconnecting == nil {
		var ret string
		return ret
	}
	return *o.Disconnecting
}

// GetDisconnectingOk returns a tuple with the Disconnecting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackStatus) GetDisconnectingOk() (*string, bool) {
	if o == nil || o.Disconnecting == nil {
		return nil, false
	}
	return o.Disconnecting, true
}

// HasDisconnecting returns a boolean if a field has been set.
func (o *ClientpackStatus) HasDisconnecting() bool {
	if o != nil && o.Disconnecting != nil {
		return true
	}

	return false
}

// SetDisconnecting gets a reference to the given string and assigns it to the Disconnecting field.
func (o *ClientpackStatus) SetDisconnecting(v string) {
	o.Disconnecting = &v
}

// GetOverlayIpaddress returns the OverlayIpaddress field value if set, zero value otherwise.
func (o *ClientpackStatus) GetOverlayIpaddress() string {
	if o == nil || o.OverlayIpaddress == nil {
		var ret string
		return ret
	}
	return *o.OverlayIpaddress
}

// GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientpackStatus) GetOverlayIpaddressOk() (*string, bool) {
	if o == nil || o.OverlayIpaddress == nil {
		return nil, false
	}
	return o.OverlayIpaddress, true
}

// HasOverlayIpaddress returns a boolean if a field has been set.
func (o *ClientpackStatus) HasOverlayIpaddress() bool {
	if o != nil && o.OverlayIpaddress != nil {
		return true
	}

	return false
}

// SetOverlayIpaddress gets a reference to the given string and assigns it to the OverlayIpaddress field.
func (o *ClientpackStatus) SetOverlayIpaddress(v string) {
	o.OverlayIpaddress = &v
}

func (o ClientpackStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disconnecting != nil {
		toSerialize["disconnecting"] = o.Disconnecting
	}
	if o.OverlayIpaddress != nil {
		toSerialize["overlay_ipaddress"] = o.OverlayIpaddress
	}
	return json.Marshal(toSerialize)
}

type NullableClientpackStatus struct {
	value *ClientpackStatus
	isSet bool
}

func (v NullableClientpackStatus) Get() *ClientpackStatus {
	return v.value
}

func (v *NullableClientpackStatus) Set(val *ClientpackStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClientpackStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClientpackStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientpackStatus(val *ClientpackStatus) *NullableClientpackStatus {
	return &NullableClientpackStatus{value: val, isSet: true}
}

func (v NullableClientpackStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientpackStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



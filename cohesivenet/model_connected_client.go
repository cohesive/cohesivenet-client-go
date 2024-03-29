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

// ConnectedClient struct for ConnectedClient
type ConnectedClient struct {
	Managerid *int32 `json:"managerid,omitempty"`
	OverlayIpaddress *string `json:"overlay_ipaddress,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	// Key, value object of tags
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewConnectedClient instantiates a new ConnectedClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedClient() *ConnectedClient {
	this := ConnectedClient{}
	return &this
}

// NewConnectedClientWithDefaults instantiates a new ConnectedClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedClientWithDefaults() *ConnectedClient {
	this := ConnectedClient{}
	return &this
}

// GetManagerid returns the Managerid field value if set, zero value otherwise.
func (o *ConnectedClient) GetManagerid() int32 {
	if o == nil || o.Managerid == nil {
		var ret int32
		return ret
	}
	return *o.Managerid
}

// GetManageridOk returns a tuple with the Managerid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedClient) GetManageridOk() (*int32, bool) {
	if o == nil || o.Managerid == nil {
		return nil, false
	}
	return o.Managerid, true
}

// HasManagerid returns a boolean if a field has been set.
func (o *ConnectedClient) HasManagerid() bool {
	if o != nil && o.Managerid != nil {
		return true
	}

	return false
}

// SetManagerid gets a reference to the given int32 and assigns it to the Managerid field.
func (o *ConnectedClient) SetManagerid(v int32) {
	o.Managerid = &v
}

// GetOverlayIpaddress returns the OverlayIpaddress field value if set, zero value otherwise.
func (o *ConnectedClient) GetOverlayIpaddress() string {
	if o == nil || o.OverlayIpaddress == nil {
		var ret string
		return ret
	}
	return *o.OverlayIpaddress
}

// GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedClient) GetOverlayIpaddressOk() (*string, bool) {
	if o == nil || o.OverlayIpaddress == nil {
		return nil, false
	}
	return o.OverlayIpaddress, true
}

// HasOverlayIpaddress returns a boolean if a field has been set.
func (o *ConnectedClient) HasOverlayIpaddress() bool {
	if o != nil && o.OverlayIpaddress != nil {
		return true
	}

	return false
}

// SetOverlayIpaddress gets a reference to the given string and assigns it to the OverlayIpaddress field.
func (o *ConnectedClient) SetOverlayIpaddress(v string) {
	o.OverlayIpaddress = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *ConnectedClient) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedClient) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *ConnectedClient) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *ConnectedClient) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConnectedClient) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedClient) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConnectedClient) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ConnectedClient) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o ConnectedClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Managerid != nil {
		toSerialize["managerid"] = o.Managerid
	}
	if o.OverlayIpaddress != nil {
		toSerialize["overlay_ipaddress"] = o.OverlayIpaddress
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedClient struct {
	value *ConnectedClient
	isSet bool
}

func (v NullableConnectedClient) Get() *ConnectedClient {
	return v.value
}

func (v *NullableConnectedClient) Set(val *ConnectedClient) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedClient) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedClient(val *ConnectedClient) *NullableConnectedClient {
	return &NullableConnectedClient{value: val, isSet: true}
}

func (v NullableConnectedClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



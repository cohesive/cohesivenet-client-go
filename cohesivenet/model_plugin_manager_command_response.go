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

// PluginManagerCommandResponse struct for PluginManagerCommandResponse
type PluginManagerCommandResponse struct {
	Response *PluginManagerCommandOutput `json:"response,omitempty"`
}

// NewPluginManagerCommandResponse instantiates a new PluginManagerCommandResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerCommandResponse() *PluginManagerCommandResponse {
	this := PluginManagerCommandResponse{}
	return &this
}

// NewPluginManagerCommandResponseWithDefaults instantiates a new PluginManagerCommandResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerCommandResponseWithDefaults() *PluginManagerCommandResponse {
	this := PluginManagerCommandResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PluginManagerCommandResponse) GetResponse() PluginManagerCommandOutput {
	if o == nil || o.Response == nil {
		var ret PluginManagerCommandOutput
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerCommandResponse) GetResponseOk() (*PluginManagerCommandOutput, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PluginManagerCommandResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given PluginManagerCommandOutput and assigns it to the Response field.
func (o *PluginManagerCommandResponse) SetResponse(v PluginManagerCommandOutput) {
	o.Response = &v
}

func (o PluginManagerCommandResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerCommandResponse struct {
	value *PluginManagerCommandResponse
	isSet bool
}

func (v NullablePluginManagerCommandResponse) Get() *PluginManagerCommandResponse {
	return v.value
}

func (v *NullablePluginManagerCommandResponse) Set(val *PluginManagerCommandResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerCommandResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerCommandResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerCommandResponse(val *PluginManagerCommandResponse) *NullablePluginManagerCommandResponse {
	return &NullablePluginManagerCommandResponse{value: val, isSet: true}
}

func (v NullablePluginManagerCommandResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerCommandResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// PluginExportsListResponse struct for PluginExportsListResponse
type PluginExportsListResponse struct {
	Response []PluginExport `json:"response,omitempty"`
}

// NewPluginExportsListResponse instantiates a new PluginExportsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginExportsListResponse() *PluginExportsListResponse {
	this := PluginExportsListResponse{}
	return &this
}

// NewPluginExportsListResponseWithDefaults instantiates a new PluginExportsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginExportsListResponseWithDefaults() *PluginExportsListResponse {
	this := PluginExportsListResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PluginExportsListResponse) GetResponse() []PluginExport {
	if o == nil || o.Response == nil {
		var ret []PluginExport
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginExportsListResponse) GetResponseOk() ([]PluginExport, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PluginExportsListResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []PluginExport and assigns it to the Response field.
func (o *PluginExportsListResponse) SetResponse(v []PluginExport) {
	o.Response = v
}

func (o PluginExportsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullablePluginExportsListResponse struct {
	value *PluginExportsListResponse
	isSet bool
}

func (v NullablePluginExportsListResponse) Get() *PluginExportsListResponse {
	return v.value
}

func (v *NullablePluginExportsListResponse) Set(val *PluginExportsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginExportsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginExportsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginExportsListResponse(val *PluginExportsListResponse) *NullablePluginExportsListResponse {
	return &NullablePluginExportsListResponse{value: val, isSet: true}
}

func (v NullablePluginExportsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginExportsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



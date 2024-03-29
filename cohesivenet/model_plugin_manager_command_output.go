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

// PluginManagerCommandOutput struct for PluginManagerCommandOutput
type PluginManagerCommandOutput struct {
	Output *string `json:"output,omitempty"`
	Timeout *bool `json:"timeout,omitempty"`
	Failed *bool `json:"failed,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewPluginManagerCommandOutput instantiates a new PluginManagerCommandOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerCommandOutput() *PluginManagerCommandOutput {
	this := PluginManagerCommandOutput{}
	return &this
}

// NewPluginManagerCommandOutputWithDefaults instantiates a new PluginManagerCommandOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerCommandOutputWithDefaults() *PluginManagerCommandOutput {
	this := PluginManagerCommandOutput{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *PluginManagerCommandOutput) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerCommandOutput) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *PluginManagerCommandOutput) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *PluginManagerCommandOutput) SetOutput(v string) {
	o.Output = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PluginManagerCommandOutput) GetTimeout() bool {
	if o == nil || o.Timeout == nil {
		var ret bool
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerCommandOutput) GetTimeoutOk() (*bool, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PluginManagerCommandOutput) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given bool and assigns it to the Timeout field.
func (o *PluginManagerCommandOutput) SetTimeout(v bool) {
	o.Timeout = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *PluginManagerCommandOutput) GetFailed() bool {
	if o == nil || o.Failed == nil {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerCommandOutput) GetFailedOk() (*bool, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *PluginManagerCommandOutput) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *PluginManagerCommandOutput) SetFailed(v bool) {
	o.Failed = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PluginManagerCommandOutput) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerCommandOutput) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PluginManagerCommandOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *PluginManagerCommandOutput) SetError(v string) {
	o.Error = &v
}

func (o PluginManagerCommandOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullablePluginManagerCommandOutput struct {
	value *PluginManagerCommandOutput
	isSet bool
}

func (v NullablePluginManagerCommandOutput) Get() *PluginManagerCommandOutput {
	return v.value
}

func (v *NullablePluginManagerCommandOutput) Set(val *PluginManagerCommandOutput) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerCommandOutput) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerCommandOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerCommandOutput(val *PluginManagerCommandOutput) *NullablePluginManagerCommandOutput {
	return &NullablePluginManagerCommandOutput{value: val, isSet: true}
}

func (v NullablePluginManagerCommandOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerCommandOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// CreatePluginInstanceFirewallRuleRequest struct for CreatePluginInstanceFirewallRuleRequest
type CreatePluginInstanceFirewallRuleRequest struct {
	// One of ssh, internet or port_map
	Preset string `json:"preset"`
	// VNS3 port. Required for preset \"port_map\"
	HostPort *int32 `json:"host_port,omitempty"`
	// Plugin port to map VNS3 port to. Required for preset \"port_map\"
	ContainerPort *int32 `json:"container_port,omitempty"`
	// Protocol for port map. Required for preset \"port_map\"
	Protocol *string `json:"protocol,omitempty"`
}

// NewCreatePluginInstanceFirewallRuleRequest instantiates a new CreatePluginInstanceFirewallRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePluginInstanceFirewallRuleRequest(preset string) *CreatePluginInstanceFirewallRuleRequest {
	this := CreatePluginInstanceFirewallRuleRequest{}
	this.Preset = preset
	return &this
}

// NewCreatePluginInstanceFirewallRuleRequestWithDefaults instantiates a new CreatePluginInstanceFirewallRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePluginInstanceFirewallRuleRequestWithDefaults() *CreatePluginInstanceFirewallRuleRequest {
	this := CreatePluginInstanceFirewallRuleRequest{}
	return &this
}

// GetPreset returns the Preset field value
func (o *CreatePluginInstanceFirewallRuleRequest) GetPreset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Preset
}

// GetPresetOk returns a tuple with the Preset field value
// and a boolean to check if the value has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) GetPresetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preset, true
}

// SetPreset sets field value
func (o *CreatePluginInstanceFirewallRuleRequest) SetPreset(v string) {
	o.Preset = v
}

// GetHostPort returns the HostPort field value if set, zero value otherwise.
func (o *CreatePluginInstanceFirewallRuleRequest) GetHostPort() int32 {
	if o == nil || o.HostPort == nil {
		var ret int32
		return ret
	}
	return *o.HostPort
}

// GetHostPortOk returns a tuple with the HostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) GetHostPortOk() (*int32, bool) {
	if o == nil || o.HostPort == nil {
		return nil, false
	}
	return o.HostPort, true
}

// HasHostPort returns a boolean if a field has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) HasHostPort() bool {
	if o != nil && o.HostPort != nil {
		return true
	}

	return false
}

// SetHostPort gets a reference to the given int32 and assigns it to the HostPort field.
func (o *CreatePluginInstanceFirewallRuleRequest) SetHostPort(v int32) {
	o.HostPort = &v
}

// GetContainerPort returns the ContainerPort field value if set, zero value otherwise.
func (o *CreatePluginInstanceFirewallRuleRequest) GetContainerPort() int32 {
	if o == nil || o.ContainerPort == nil {
		var ret int32
		return ret
	}
	return *o.ContainerPort
}

// GetContainerPortOk returns a tuple with the ContainerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) GetContainerPortOk() (*int32, bool) {
	if o == nil || o.ContainerPort == nil {
		return nil, false
	}
	return o.ContainerPort, true
}

// HasContainerPort returns a boolean if a field has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) HasContainerPort() bool {
	if o != nil && o.ContainerPort != nil {
		return true
	}

	return false
}

// SetContainerPort gets a reference to the given int32 and assigns it to the ContainerPort field.
func (o *CreatePluginInstanceFirewallRuleRequest) SetContainerPort(v int32) {
	o.ContainerPort = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CreatePluginInstanceFirewallRuleRequest) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *CreatePluginInstanceFirewallRuleRequest) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *CreatePluginInstanceFirewallRuleRequest) SetProtocol(v string) {
	o.Protocol = &v
}

func (o CreatePluginInstanceFirewallRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preset"] = o.Preset
	}
	if o.HostPort != nil {
		toSerialize["host_port"] = o.HostPort
	}
	if o.ContainerPort != nil {
		toSerialize["container_port"] = o.ContainerPort
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePluginInstanceFirewallRuleRequest struct {
	value *CreatePluginInstanceFirewallRuleRequest
	isSet bool
}

func (v NullableCreatePluginInstanceFirewallRuleRequest) Get() *CreatePluginInstanceFirewallRuleRequest {
	return v.value
}

func (v *NullableCreatePluginInstanceFirewallRuleRequest) Set(val *CreatePluginInstanceFirewallRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePluginInstanceFirewallRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePluginInstanceFirewallRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePluginInstanceFirewallRuleRequest(val *CreatePluginInstanceFirewallRuleRequest) *NullableCreatePluginInstanceFirewallRuleRequest {
	return &NullableCreatePluginInstanceFirewallRuleRequest{value: val, isSet: true}
}

func (v NullableCreatePluginInstanceFirewallRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePluginInstanceFirewallRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


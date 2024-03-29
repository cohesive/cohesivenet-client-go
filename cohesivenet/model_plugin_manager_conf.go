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

// PluginManagerConf struct for PluginManagerConf
type PluginManagerConf struct {
	LogFiles []PluginManagerLogFile `json:"log_files,omitempty"`
	ConfigurationFiles []PluginManagerConfigFile `json:"configuration_files,omitempty"`
	Ports []PluginManagerPort `json:"ports,omitempty"`
	ProcessManager *PluginManagerProcessManager `json:"process_manager,omitempty"`
	Executables []PluginManagerExecutable `json:"executables,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PluginManagerConf PluginManagerConf

// NewPluginManagerConf instantiates a new PluginManagerConf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginManagerConf() *PluginManagerConf {
	this := PluginManagerConf{}
	return &this
}

// NewPluginManagerConfWithDefaults instantiates a new PluginManagerConf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginManagerConfWithDefaults() *PluginManagerConf {
	this := PluginManagerConf{}
	return &this
}

// GetLogFiles returns the LogFiles field value if set, zero value otherwise.
func (o *PluginManagerConf) GetLogFiles() []PluginManagerLogFile {
	if o == nil || o.LogFiles == nil {
		var ret []PluginManagerLogFile
		return ret
	}
	return o.LogFiles
}

// GetLogFilesOk returns a tuple with the LogFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConf) GetLogFilesOk() ([]PluginManagerLogFile, bool) {
	if o == nil || o.LogFiles == nil {
		return nil, false
	}
	return o.LogFiles, true
}

// HasLogFiles returns a boolean if a field has been set.
func (o *PluginManagerConf) HasLogFiles() bool {
	if o != nil && o.LogFiles != nil {
		return true
	}

	return false
}

// SetLogFiles gets a reference to the given []PluginManagerConfLogFilesInner and assigns it to the LogFiles field.
func (o *PluginManagerConf) SetLogFiles(v []PluginManagerLogFile) {
	o.LogFiles = v
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise.
func (o *PluginManagerConf) GetConfigurationFiles() []PluginManagerConfigFile {
	if o == nil || o.ConfigurationFiles == nil {
		var ret []PluginManagerConfigFile
		return ret
	}
	return o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConf) GetConfigurationFilesOk() ([]PluginManagerConfigFile, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *PluginManagerConf) HasConfigurationFiles() bool {
	if o != nil && o.ConfigurationFiles != nil {
		return true
	}

	return false
}

// SetConfigurationFiles gets a reference to the given []PluginManagerConfigFile and assigns it to the ConfigurationFiles field.
func (o *PluginManagerConf) SetConfigurationFiles(v []PluginManagerConfigFile) {
	o.ConfigurationFiles = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *PluginManagerConf) GetPorts() []PluginManagerPort {
	if o == nil || o.Ports == nil {
		var ret []PluginManagerPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConf) GetPortsOk() ([]PluginManagerPort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *PluginManagerConf) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []PluginManagerPort and assigns it to the Ports field.
func (o *PluginManagerConf) SetPorts(v []PluginManagerPort) {
	o.Ports = v
}

// GetProcessManager returns the ProcessManager field value if set, zero value otherwise.
func (o *PluginManagerConf) GetProcessManager() PluginManagerProcessManager {
	if o == nil || o.ProcessManager == nil {
		var ret PluginManagerProcessManager
		return ret
	}
	return *o.ProcessManager
}

// GetProcessManagerOk returns a tuple with the ProcessManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConf) GetProcessManagerOk() (*PluginManagerProcessManager, bool) {
	if o == nil || o.ProcessManager == nil {
		return nil, false
	}
	return o.ProcessManager, true
}

// HasProcessManager returns a boolean if a field has been set.
func (o *PluginManagerConf) HasProcessManager() bool {
	if o != nil && o.ProcessManager != nil {
		return true
	}

	return false
}

// SetProcessManager gets a reference to the given PluginManagerProcessManager and assigns it to the ProcessManager field.
func (o *PluginManagerConf) SetProcessManager(v PluginManagerProcessManager) {
	o.ProcessManager = &v
}

// GetExecutables returns the Executables field value if set, zero value otherwise.
func (o *PluginManagerConf) GetExecutables() []PluginManagerExecutable {
	if o == nil || o.Executables == nil {
		var ret []PluginManagerExecutable
		return ret
	}
	return o.Executables
}

// GetExecutablesOk returns a tuple with the Executables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginManagerConf) GetExecutablesOk() ([]PluginManagerExecutable, bool) {
	if o == nil || o.Executables == nil {
		return nil, false
	}
	return o.Executables, true
}

// HasExecutables returns a boolean if a field has been set.
func (o *PluginManagerConf) HasExecutables() bool {
	if o != nil && o.Executables != nil {
		return true
	}

	return false
}

// SetExecutables gets a reference to the given []PluginManagerExecutable and assigns it to the Executables field.
func (o *PluginManagerConf) SetExecutables(v []PluginManagerExecutable) {
	o.Executables = v
}

func (o PluginManagerConf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogFiles != nil {
		toSerialize["log_files"] = o.LogFiles
	}
	if o.ConfigurationFiles != nil {
		toSerialize["configuration_files"] = o.ConfigurationFiles
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.ProcessManager != nil {
		toSerialize["process_manager"] = o.ProcessManager
	}
	if o.Executables != nil {
		toSerialize["executables"] = o.Executables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PluginManagerConf) UnmarshalJSON(bytes []byte) (err error) {
	varPluginManagerConf := _PluginManagerConf{}

	if err = json.Unmarshal(bytes, &varPluginManagerConf); err == nil {
		*o = PluginManagerConf(varPluginManagerConf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "log_files")
		delete(additionalProperties, "configuration_files")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "process_manager")
		delete(additionalProperties, "executables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePluginManagerConf struct {
	value *PluginManagerConf
	isSet bool
}

func (v NullablePluginManagerConf) Get() *PluginManagerConf {
	return v.value
}

func (v *NullablePluginManagerConf) Set(val *PluginManagerConf) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginManagerConf) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginManagerConf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginManagerConf(val *PluginManagerConf) *NullablePluginManagerConf {
	return &NullablePluginManagerConf{value: val, isSet: true}
}

func (v NullablePluginManagerConf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginManagerConf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



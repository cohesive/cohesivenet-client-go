# PluginManagerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogFiles** | Pointer to [**[]PluginManagerConfigLogFilesInner**](PluginManagerConfigLogFilesInner.md) |  | [optional] 
**ConfigurationFiles** | Pointer to [**[]PluginManagerConfigConfigFile**](PluginManagerConfigConfigFile.md) |  | [optional] 
**Ports** | Pointer to [**[]PluginManagerConfigPort**](PluginManagerConfigPort.md) |  | [optional] 
**ProcessManager** | Pointer to [**PluginManagerConfigProcessManager**](PluginManagerConfigProcessManager.md) |  | [optional] 
**Executables** | Pointer to [**[]PluginManagerConfigExecutable**](PluginManagerConfigExecutable.md) |  | [optional] 

## Methods

### NewPluginManagerConfig

`func NewPluginManagerConfig() *PluginManagerConfig`

NewPluginManagerConfig instantiates a new PluginManagerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginManagerConfigWithDefaults

`func NewPluginManagerConfigWithDefaults() *PluginManagerConfig`

NewPluginManagerConfigWithDefaults instantiates a new PluginManagerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogFiles

`func (o *PluginManagerConfig) GetLogFiles() []PluginManagerConfigLogFilesInner`

GetLogFiles returns the LogFiles field if non-nil, zero value otherwise.

### GetLogFilesOk

`func (o *PluginManagerConfig) GetLogFilesOk() (*[]PluginManagerConfigLogFilesInner, bool)`

GetLogFilesOk returns a tuple with the LogFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFiles

`func (o *PluginManagerConfig) SetLogFiles(v []PluginManagerConfigLogFilesInner)`

SetLogFiles sets LogFiles field to given value.

### HasLogFiles

`func (o *PluginManagerConfig) HasLogFiles() bool`

HasLogFiles returns a boolean if a field has been set.

### GetConfigurationFiles

`func (o *PluginManagerConfig) GetConfigurationFiles() []PluginManagerConfigConfigFile`

GetConfigurationFiles returns the ConfigurationFiles field if non-nil, zero value otherwise.

### GetConfigurationFilesOk

`func (o *PluginManagerConfig) GetConfigurationFilesOk() (*[]PluginManagerConfigConfigFile, bool)`

GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFiles

`func (o *PluginManagerConfig) SetConfigurationFiles(v []PluginManagerConfigConfigFile)`

SetConfigurationFiles sets ConfigurationFiles field to given value.

### HasConfigurationFiles

`func (o *PluginManagerConfig) HasConfigurationFiles() bool`

HasConfigurationFiles returns a boolean if a field has been set.

### GetPorts

`func (o *PluginManagerConfig) GetPorts() []PluginManagerConfigPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PluginManagerConfig) GetPortsOk() (*[]PluginManagerConfigPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PluginManagerConfig) SetPorts(v []PluginManagerConfigPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PluginManagerConfig) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcessManager

`func (o *PluginManagerConfig) GetProcessManager() PluginManagerConfigProcessManager`

GetProcessManager returns the ProcessManager field if non-nil, zero value otherwise.

### GetProcessManagerOk

`func (o *PluginManagerConfig) GetProcessManagerOk() (*PluginManagerConfigProcessManager, bool)`

GetProcessManagerOk returns a tuple with the ProcessManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessManager

`func (o *PluginManagerConfig) SetProcessManager(v PluginManagerConfigProcessManager)`

SetProcessManager sets ProcessManager field to given value.

### HasProcessManager

`func (o *PluginManagerConfig) HasProcessManager() bool`

HasProcessManager returns a boolean if a field has been set.

### GetExecutables

`func (o *PluginManagerConfig) GetExecutables() []PluginManagerConfigExecutable`

GetExecutables returns the Executables field if non-nil, zero value otherwise.

### GetExecutablesOk

`func (o *PluginManagerConfig) GetExecutablesOk() (*[]PluginManagerConfigExecutable, bool)`

GetExecutablesOk returns a tuple with the Executables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutables

`func (o *PluginManagerConfig) SetExecutables(v []PluginManagerConfigExecutable)`

SetExecutables sets Executables field to given value.

### HasExecutables

`func (o *PluginManagerConfig) HasExecutables() bool`

HasExecutables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PluginManagerConfigExecutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Commands** | Pointer to **map[string]string** | Map of \&quot;name\&quot; to the command for the underlying executable. For example, start -&gt; execute would pass \&quot;execute\&quot; to the executable for the \&quot;start\&quot; command.  | [optional] 

## Methods

### NewPluginManagerConfigExecutable

`func NewPluginManagerConfigExecutable() *PluginManagerConfigExecutable`

NewPluginManagerConfigExecutable instantiates a new PluginManagerConfigExecutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginManagerConfigExecutableWithDefaults

`func NewPluginManagerConfigExecutableWithDefaults() *PluginManagerConfigExecutable`

NewPluginManagerConfigExecutableWithDefaults instantiates a new PluginManagerConfigExecutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PluginManagerConfigExecutable) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PluginManagerConfigExecutable) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PluginManagerConfigExecutable) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PluginManagerConfigExecutable) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *PluginManagerConfigExecutable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginManagerConfigExecutable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginManagerConfigExecutable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginManagerConfigExecutable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommands

`func (o *PluginManagerConfigExecutable) GetCommands() map[string]string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *PluginManagerConfigExecutable) GetCommandsOk() (*map[string]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *PluginManagerConfigExecutable) SetCommands(v map[string]string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *PluginManagerConfigExecutable) HasCommands() bool`

HasCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



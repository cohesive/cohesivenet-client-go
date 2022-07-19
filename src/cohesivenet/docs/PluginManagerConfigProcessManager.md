# PluginManagerConfigProcessManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of process manager such as supervisor. Currently  we support commands for supervisor and service.  | [optional] 
**Subprocesses** | Pointer to **[]string** | Name of programs, services or units managed | [optional] 

## Methods

### NewPluginManagerConfigProcessManager

`func NewPluginManagerConfigProcessManager() *PluginManagerConfigProcessManager`

NewPluginManagerConfigProcessManager instantiates a new PluginManagerConfigProcessManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginManagerConfigProcessManagerWithDefaults

`func NewPluginManagerConfigProcessManagerWithDefaults() *PluginManagerConfigProcessManager`

NewPluginManagerConfigProcessManagerWithDefaults instantiates a new PluginManagerConfigProcessManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginManagerConfigProcessManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginManagerConfigProcessManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginManagerConfigProcessManager) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginManagerConfigProcessManager) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubprocesses

`func (o *PluginManagerConfigProcessManager) GetSubprocesses() []string`

GetSubprocesses returns the Subprocesses field if non-nil, zero value otherwise.

### GetSubprocessesOk

`func (o *PluginManagerConfigProcessManager) GetSubprocessesOk() (*[]string, bool)`

GetSubprocessesOk returns a tuple with the Subprocesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocesses

`func (o *PluginManagerConfigProcessManager) SetSubprocesses(v []string)`

SetSubprocesses sets Subprocesses field to given value.

### HasSubprocesses

`func (o *PluginManagerConfigProcessManager) HasSubprocesses() bool`

HasSubprocesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



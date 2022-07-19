# StartPluginInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new plugin instance | 
**PluginId** | **int32** | ID of plugin | 
**Description** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** | IP to use from plugin network | [optional] 
**Command** | Pointer to **string** | Optionally override the plugins command. Required if plugin command not defined. | [optional] 
**ManagerConfig** | Pointer to [**PluginManagerConfig**](PluginManagerConfig.md) |  | [optional] 
**Environment** | Pointer to [**[]StartPluginInstanceRequestEnvironmentInner**](StartPluginInstanceRequestEnvironmentInner.md) |  | [optional] 

## Methods

### NewStartPluginInstanceRequest

`func NewStartPluginInstanceRequest(name string, pluginId int32, ) *StartPluginInstanceRequest`

NewStartPluginInstanceRequest instantiates a new StartPluginInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartPluginInstanceRequestWithDefaults

`func NewStartPluginInstanceRequestWithDefaults() *StartPluginInstanceRequest`

NewStartPluginInstanceRequestWithDefaults instantiates a new StartPluginInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StartPluginInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartPluginInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartPluginInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPluginId

`func (o *StartPluginInstanceRequest) GetPluginId() int32`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *StartPluginInstanceRequest) GetPluginIdOk() (*int32, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *StartPluginInstanceRequest) SetPluginId(v int32)`

SetPluginId sets PluginId field to given value.


### GetDescription

`func (o *StartPluginInstanceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StartPluginInstanceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StartPluginInstanceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StartPluginInstanceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *StartPluginInstanceRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StartPluginInstanceRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StartPluginInstanceRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StartPluginInstanceRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCommand

`func (o *StartPluginInstanceRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *StartPluginInstanceRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *StartPluginInstanceRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *StartPluginInstanceRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetManagerConfig

`func (o *StartPluginInstanceRequest) GetManagerConfig() PluginManagerConfig`

GetManagerConfig returns the ManagerConfig field if non-nil, zero value otherwise.

### GetManagerConfigOk

`func (o *StartPluginInstanceRequest) GetManagerConfigOk() (*PluginManagerConfig, bool)`

GetManagerConfigOk returns a tuple with the ManagerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerConfig

`func (o *StartPluginInstanceRequest) SetManagerConfig(v PluginManagerConfig)`

SetManagerConfig sets ManagerConfig field to given value.

### HasManagerConfig

`func (o *StartPluginInstanceRequest) HasManagerConfig() bool`

HasManagerConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *StartPluginInstanceRequest) GetEnvironment() []StartPluginInstanceRequestEnvironmentInner`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StartPluginInstanceRequest) GetEnvironmentOk() (*[]StartPluginInstanceRequestEnvironmentInner, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StartPluginInstanceRequest) SetEnvironment(v []StartPluginInstanceRequestEnvironmentInner)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StartPluginInstanceRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



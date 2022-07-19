# CreatePluginInstancePresetFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preset** | **string** | One of ssh, internet or port_map | 
**HostPort** | Pointer to **int32** | VNS3 port. Required for preset \&quot;port_map\&quot; | [optional] 
**ContainerPort** | Pointer to **int32** | Plugin port to map VNS3 port to. Required for preset \&quot;port_map\&quot; | [optional] 
**Protocol** | Pointer to **string** | Protocol for port map. Required for preset \&quot;port_map\&quot; | [optional] 

## Methods

### NewCreatePluginInstancePresetFirewallRule

`func NewCreatePluginInstancePresetFirewallRule(preset string, ) *CreatePluginInstancePresetFirewallRule`

NewCreatePluginInstancePresetFirewallRule instantiates a new CreatePluginInstancePresetFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePluginInstancePresetFirewallRuleWithDefaults

`func NewCreatePluginInstancePresetFirewallRuleWithDefaults() *CreatePluginInstancePresetFirewallRule`

NewCreatePluginInstancePresetFirewallRuleWithDefaults instantiates a new CreatePluginInstancePresetFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreset

`func (o *CreatePluginInstancePresetFirewallRule) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *CreatePluginInstancePresetFirewallRule) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *CreatePluginInstancePresetFirewallRule) SetPreset(v string)`

SetPreset sets Preset field to given value.


### GetHostPort

`func (o *CreatePluginInstancePresetFirewallRule) GetHostPort() int32`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *CreatePluginInstancePresetFirewallRule) GetHostPortOk() (*int32, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *CreatePluginInstancePresetFirewallRule) SetHostPort(v int32)`

SetHostPort sets HostPort field to given value.

### HasHostPort

`func (o *CreatePluginInstancePresetFirewallRule) HasHostPort() bool`

HasHostPort returns a boolean if a field has been set.

### GetContainerPort

`func (o *CreatePluginInstancePresetFirewallRule) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *CreatePluginInstancePresetFirewallRule) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *CreatePluginInstancePresetFirewallRule) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *CreatePluginInstancePresetFirewallRule) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.

### GetProtocol

`func (o *CreatePluginInstancePresetFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreatePluginInstancePresetFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreatePluginInstancePresetFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CreatePluginInstancePresetFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



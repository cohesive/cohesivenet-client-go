# PluginInstanceFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**PluginPort** | Pointer to **int32** |  | [optional] 
**DestinationPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewPluginInstanceFirewallRule

`func NewPluginInstanceFirewallRule() *PluginInstanceFirewallRule`

NewPluginInstanceFirewallRule instantiates a new PluginInstanceFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInstanceFirewallRuleWithDefaults

`func NewPluginInstanceFirewallRuleWithDefaults() *PluginInstanceFirewallRule`

NewPluginInstanceFirewallRuleWithDefaults instantiates a new PluginInstanceFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *PluginInstanceFirewallRule) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PluginInstanceFirewallRule) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PluginInstanceFirewallRule) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PluginInstanceFirewallRule) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetTags

`func (o *PluginInstanceFirewallRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PluginInstanceFirewallRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PluginInstanceFirewallRule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PluginInstanceFirewallRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPluginPort

`func (o *PluginInstanceFirewallRule) GetPluginPort() int32`

GetPluginPort returns the PluginPort field if non-nil, zero value otherwise.

### GetPluginPortOk

`func (o *PluginInstanceFirewallRule) GetPluginPortOk() (*int32, bool)`

GetPluginPortOk returns a tuple with the PluginPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginPort

`func (o *PluginInstanceFirewallRule) SetPluginPort(v int32)`

SetPluginPort sets PluginPort field to given value.

### HasPluginPort

`func (o *PluginInstanceFirewallRule) HasPluginPort() bool`

HasPluginPort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *PluginInstanceFirewallRule) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *PluginInstanceFirewallRule) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *PluginInstanceFirewallRule) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *PluginInstanceFirewallRule) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



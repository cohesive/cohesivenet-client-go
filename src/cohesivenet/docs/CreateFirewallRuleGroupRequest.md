# CreateFirewallRuleGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of group | 
**RuleIds** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFirewallRuleGroupRequest

`func NewCreateFirewallRuleGroupRequest(name string, ) *CreateFirewallRuleGroupRequest`

NewCreateFirewallRuleGroupRequest instantiates a new CreateFirewallRuleGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRuleGroupRequestWithDefaults

`func NewCreateFirewallRuleGroupRequestWithDefaults() *CreateFirewallRuleGroupRequest`

NewCreateFirewallRuleGroupRequestWithDefaults instantiates a new CreateFirewallRuleGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirewallRuleGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirewallRuleGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirewallRuleGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRuleIds

`func (o *CreateFirewallRuleGroupRequest) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *CreateFirewallRuleGroupRequest) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *CreateFirewallRuleGroupRequest) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *CreateFirewallRuleGroupRequest) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.

### GetDescription

`func (o *CreateFirewallRuleGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirewallRuleGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirewallRuleGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirewallRuleGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



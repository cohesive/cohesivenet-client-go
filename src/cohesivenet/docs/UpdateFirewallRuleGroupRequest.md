# UpdateFirewallRuleGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for rule group. This will change the URI as the name is the unique ID | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RuleIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateFirewallRuleGroupRequest

`func NewUpdateFirewallRuleGroupRequest() *UpdateFirewallRuleGroupRequest`

NewUpdateFirewallRuleGroupRequest instantiates a new UpdateFirewallRuleGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallRuleGroupRequestWithDefaults

`func NewUpdateFirewallRuleGroupRequestWithDefaults() *UpdateFirewallRuleGroupRequest`

NewUpdateFirewallRuleGroupRequestWithDefaults instantiates a new UpdateFirewallRuleGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFirewallRuleGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirewallRuleGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirewallRuleGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFirewallRuleGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateFirewallRuleGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFirewallRuleGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFirewallRuleGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFirewallRuleGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleIds

`func (o *UpdateFirewallRuleGroupRequest) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *UpdateFirewallRuleGroupRequest) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *UpdateFirewallRuleGroupRequest) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *UpdateFirewallRuleGroupRequest) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



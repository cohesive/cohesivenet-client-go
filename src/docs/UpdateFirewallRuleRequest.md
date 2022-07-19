# UpdateFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** | firewall rule string | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** | if true, rule will be disabled | [optional] 

## Methods

### NewUpdateFirewallRuleRequest

`func NewUpdateFirewallRuleRequest() *UpdateFirewallRuleRequest`

NewUpdateFirewallRuleRequest instantiates a new UpdateFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallRuleRequestWithDefaults

`func NewUpdateFirewallRuleRequestWithDefaults() *UpdateFirewallRuleRequest`

NewUpdateFirewallRuleRequestWithDefaults instantiates a new UpdateFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *UpdateFirewallRuleRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UpdateFirewallRuleRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UpdateFirewallRuleRequest) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *UpdateFirewallRuleRequest) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetComment

`func (o *UpdateFirewallRuleRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateFirewallRuleRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateFirewallRuleRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateFirewallRuleRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroups

`func (o *UpdateFirewallRuleRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateFirewallRuleRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateFirewallRuleRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateFirewallRuleRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDisabled

`func (o *UpdateFirewallRuleRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateFirewallRuleRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateFirewallRuleRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateFirewallRuleRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFirewallRuleRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** | firewall rule string | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to **[]string** | list of rules to create | [optional] 
**Position** | Pointer to **int32** | starting position for the rule or rules. -1 indicates end of firewall | [optional] [default to -1]
**Groups** | Pointer to **[]string** | List of groups to add this rule to | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateFirewallRuleRequestV2

`func NewCreateFirewallRuleRequestV2() *CreateFirewallRuleRequestV2`

NewCreateFirewallRuleRequestV2 instantiates a new CreateFirewallRuleRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRuleRequestV2WithDefaults

`func NewCreateFirewallRuleRequestV2WithDefaults() *CreateFirewallRuleRequestV2`

NewCreateFirewallRuleRequestV2WithDefaults instantiates a new CreateFirewallRuleRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *CreateFirewallRuleRequestV2) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateFirewallRuleRequestV2) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateFirewallRuleRequestV2) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CreateFirewallRuleRequestV2) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetComment

`func (o *CreateFirewallRuleRequestV2) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateFirewallRuleRequestV2) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateFirewallRuleRequestV2) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateFirewallRuleRequestV2) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRules

`func (o *CreateFirewallRuleRequestV2) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateFirewallRuleRequestV2) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateFirewallRuleRequestV2) SetRules(v []string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateFirewallRuleRequestV2) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetPosition

`func (o *CreateFirewallRuleRequestV2) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateFirewallRuleRequestV2) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateFirewallRuleRequestV2) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateFirewallRuleRequestV2) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetGroups

`func (o *CreateFirewallRuleRequestV2) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateFirewallRuleRequestV2) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateFirewallRuleRequestV2) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateFirewallRuleRequestV2) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDisabled

`func (o *CreateFirewallRuleRequestV2) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateFirewallRuleRequestV2) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateFirewallRuleRequestV2) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateFirewallRuleRequestV2) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



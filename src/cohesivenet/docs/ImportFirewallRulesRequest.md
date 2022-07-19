# ImportFirewallRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | **string** | Rules as string. This can be valid json for a rule object or rules delimited by newlines | 
**Position** | Pointer to **int32** | Starting position for imported rules. -1 indicates end of firewall. | [optional] [default to -1]
**Disabled** | Pointer to **bool** | Import all rules as immediately disabled | [optional] 
**Groups** | Pointer to **[]string** | Groups to add these imported rules to | [optional] 

## Methods

### NewImportFirewallRulesRequest

`func NewImportFirewallRulesRequest(rules string, ) *ImportFirewallRulesRequest`

NewImportFirewallRulesRequest instantiates a new ImportFirewallRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFirewallRulesRequestWithDefaults

`func NewImportFirewallRulesRequestWithDefaults() *ImportFirewallRulesRequest`

NewImportFirewallRulesRequestWithDefaults instantiates a new ImportFirewallRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ImportFirewallRulesRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ImportFirewallRulesRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ImportFirewallRulesRequest) SetRules(v string)`

SetRules sets Rules field to given value.


### GetPosition

`func (o *ImportFirewallRulesRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ImportFirewallRulesRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ImportFirewallRulesRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ImportFirewallRulesRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetDisabled

`func (o *ImportFirewallRulesRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ImportFirewallRulesRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ImportFirewallRulesRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ImportFirewallRulesRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGroups

`func (o *ImportFirewallRulesRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ImportFirewallRulesRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ImportFirewallRulesRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ImportFirewallRulesRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FirewallRuleV2OsRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Table** | Pointer to **string** | OS level table. This will differ from the VNS3 level table | [optional] 
**RuleType** | Pointer to **string** | Indicates where this rule exists in OS | [optional] 

## Methods

### NewFirewallRuleV2OsRulesInner

`func NewFirewallRuleV2OsRulesInner() *FirewallRuleV2OsRulesInner`

NewFirewallRuleV2OsRulesInner instantiates a new FirewallRuleV2OsRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleV2OsRulesInnerWithDefaults

`func NewFirewallRuleV2OsRulesInnerWithDefaults() *FirewallRuleV2OsRulesInner`

NewFirewallRuleV2OsRulesInnerWithDefaults instantiates a new FirewallRuleV2OsRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *FirewallRuleV2OsRulesInner) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallRuleV2OsRulesInner) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallRuleV2OsRulesInner) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FirewallRuleV2OsRulesInner) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallRuleV2OsRulesInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallRuleV2OsRulesInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallRuleV2OsRulesInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallRuleV2OsRulesInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTable

`func (o *FirewallRuleV2OsRulesInner) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *FirewallRuleV2OsRulesInner) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *FirewallRuleV2OsRulesInner) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *FirewallRuleV2OsRulesInner) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetRuleType

`func (o *FirewallRuleV2OsRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *FirewallRuleV2OsRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *FirewallRuleV2OsRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *FirewallRuleV2OsRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



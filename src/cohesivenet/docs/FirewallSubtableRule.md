# FirewallSubtableRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** |  | [optional] 
**RuleResolved** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**OsRules** | Pointer to [**[]FirewallSubtableRuleOsRulesInner**](FirewallSubtableRuleOsRulesInner.md) |  | [optional] 

## Methods

### NewFirewallSubtableRule

`func NewFirewallSubtableRule() *FirewallSubtableRule`

NewFirewallSubtableRule instantiates a new FirewallSubtableRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubtableRuleWithDefaults

`func NewFirewallSubtableRuleWithDefaults() *FirewallSubtableRule`

NewFirewallSubtableRuleWithDefaults instantiates a new FirewallSubtableRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *FirewallSubtableRule) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallSubtableRule) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallSubtableRule) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FirewallSubtableRule) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetRuleResolved

`func (o *FirewallSubtableRule) GetRuleResolved() string`

GetRuleResolved returns the RuleResolved field if non-nil, zero value otherwise.

### GetRuleResolvedOk

`func (o *FirewallSubtableRule) GetRuleResolvedOk() (*string, bool)`

GetRuleResolvedOk returns a tuple with the RuleResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleResolved

`func (o *FirewallSubtableRule) SetRuleResolved(v string)`

SetRuleResolved sets RuleResolved field to given value.

### HasRuleResolved

`func (o *FirewallSubtableRule) HasRuleResolved() bool`

HasRuleResolved returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallSubtableRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallSubtableRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallSubtableRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallSubtableRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComment

`func (o *FirewallSubtableRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FirewallSubtableRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FirewallSubtableRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FirewallSubtableRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *FirewallSubtableRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallSubtableRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallSubtableRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallSubtableRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisabled

`func (o *FirewallSubtableRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FirewallSubtableRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FirewallSubtableRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FirewallSubtableRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOsRules

`func (o *FirewallSubtableRule) GetOsRules() []FirewallSubtableRuleOsRulesInner`

GetOsRules returns the OsRules field if non-nil, zero value otherwise.

### GetOsRulesOk

`func (o *FirewallSubtableRule) GetOsRulesOk() (*[]FirewallSubtableRuleOsRulesInner, bool)`

GetOsRulesOk returns a tuple with the OsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRules

`func (o *FirewallSubtableRule) SetOsRules(v []FirewallSubtableRuleOsRulesInner)`

SetOsRules sets OsRules field to given value.

### HasOsRules

`func (o *FirewallSubtableRule) HasOsRules() bool`

HasOsRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



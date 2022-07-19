# FirewallSubtableCreateResponseResponseAllOfErrorsInner

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
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallSubtableCreateResponseResponseAllOfErrorsInner

`func NewFirewallSubtableCreateResponseResponseAllOfErrorsInner() *FirewallSubtableCreateResponseResponseAllOfErrorsInner`

NewFirewallSubtableCreateResponseResponseAllOfErrorsInner instantiates a new FirewallSubtableCreateResponseResponseAllOfErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubtableCreateResponseResponseAllOfErrorsInnerWithDefaults

`func NewFirewallSubtableCreateResponseResponseAllOfErrorsInnerWithDefaults() *FirewallSubtableCreateResponseResponseAllOfErrorsInner`

NewFirewallSubtableCreateResponseResponseAllOfErrorsInnerWithDefaults instantiates a new FirewallSubtableCreateResponseResponseAllOfErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetRuleResolved

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetRuleResolved() string`

GetRuleResolved returns the RuleResolved field if non-nil, zero value otherwise.

### GetRuleResolvedOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetRuleResolvedOk() (*string, bool)`

GetRuleResolvedOk returns a tuple with the RuleResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleResolved

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetRuleResolved(v string)`

SetRuleResolved sets RuleResolved field to given value.

### HasRuleResolved

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasRuleResolved() bool`

HasRuleResolved returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComment

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisabled

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOsRules

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetOsRules() []FirewallSubtableRuleOsRulesInner`

GetOsRules returns the OsRules field if non-nil, zero value otherwise.

### GetOsRulesOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetOsRulesOk() (*[]FirewallSubtableRuleOsRulesInner, bool)`

GetOsRulesOk returns a tuple with the OsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRules

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetOsRules(v []FirewallSubtableRuleOsRulesInner)`

SetOsRules sets OsRules field to given value.

### HasOsRules

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasOsRules() bool`

HasOsRules returns a boolean if a field has been set.

### GetError

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FirewallSubtableCreateResponseResponseAllOfErrorsInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



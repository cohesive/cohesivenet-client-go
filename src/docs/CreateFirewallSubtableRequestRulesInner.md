# CreateFirewallSubtableRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** |  | 
**Position** | **int32** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateFirewallSubtableRequestRulesInner

`func NewCreateFirewallSubtableRequestRulesInner(rule string, position int32, ) *CreateFirewallSubtableRequestRulesInner`

NewCreateFirewallSubtableRequestRulesInner instantiates a new CreateFirewallSubtableRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallSubtableRequestRulesInnerWithDefaults

`func NewCreateFirewallSubtableRequestRulesInnerWithDefaults() *CreateFirewallSubtableRequestRulesInner`

NewCreateFirewallSubtableRequestRulesInnerWithDefaults instantiates a new CreateFirewallSubtableRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *CreateFirewallSubtableRequestRulesInner) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateFirewallSubtableRequestRulesInner) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateFirewallSubtableRequestRulesInner) SetRule(v string)`

SetRule sets Rule field to given value.


### GetPosition

`func (o *CreateFirewallSubtableRequestRulesInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateFirewallSubtableRequestRulesInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateFirewallSubtableRequestRulesInner) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetComment

`func (o *CreateFirewallSubtableRequestRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateFirewallSubtableRequestRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateFirewallSubtableRequestRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateFirewallSubtableRequestRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *CreateFirewallSubtableRequestRulesInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateFirewallSubtableRequestRulesInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateFirewallSubtableRequestRulesInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateFirewallSubtableRequestRulesInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



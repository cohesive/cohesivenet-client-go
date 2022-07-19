# AddFirewallRuleToSubtableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** |  | 
**Position** | Pointer to **int32** | Position of rule in subtable. -1 indicates the end of the subtable | [optional] [default to -1]
**Comment** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAddFirewallRuleToSubtableRequest

`func NewAddFirewallRuleToSubtableRequest(rule string, ) *AddFirewallRuleToSubtableRequest`

NewAddFirewallRuleToSubtableRequest instantiates a new AddFirewallRuleToSubtableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFirewallRuleToSubtableRequestWithDefaults

`func NewAddFirewallRuleToSubtableRequestWithDefaults() *AddFirewallRuleToSubtableRequest`

NewAddFirewallRuleToSubtableRequestWithDefaults instantiates a new AddFirewallRuleToSubtableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *AddFirewallRuleToSubtableRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *AddFirewallRuleToSubtableRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *AddFirewallRuleToSubtableRequest) SetRule(v string)`

SetRule sets Rule field to given value.


### GetPosition

`func (o *AddFirewallRuleToSubtableRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AddFirewallRuleToSubtableRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AddFirewallRuleToSubtableRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AddFirewallRuleToSubtableRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComment

`func (o *AddFirewallRuleToSubtableRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddFirewallRuleToSubtableRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddFirewallRuleToSubtableRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddFirewallRuleToSubtableRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *AddFirewallRuleToSubtableRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AddFirewallRuleToSubtableRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AddFirewallRuleToSubtableRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AddFirewallRuleToSubtableRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



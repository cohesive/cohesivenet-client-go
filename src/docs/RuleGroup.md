# RuleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to **[]string** | List of rule IDs in group | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuleGroup

`func NewRuleGroup() *RuleGroup`

NewRuleGroup instantiates a new RuleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleGroupWithDefaults

`func NewRuleGroupWithDefaults() *RuleGroup`

NewRuleGroupWithDefaults instantiates a new RuleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RuleGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *RuleGroup) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RuleGroup) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RuleGroup) SetRules(v []string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RuleGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSize

`func (o *RuleGroup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RuleGroup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RuleGroup) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RuleGroup) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



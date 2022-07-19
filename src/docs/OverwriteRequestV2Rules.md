# OverwriteRequestV2Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** |  | 
**Position** | Pointer to **int32** | position in resulting firewall. default position is position in list | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | optional ID to preserve groups and map errors | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewOverwriteRequestV2Rules

`func NewOverwriteRequestV2Rules(rule string, ) *OverwriteRequestV2Rules`

NewOverwriteRequestV2Rules instantiates a new OverwriteRequestV2Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverwriteRequestV2RulesWithDefaults

`func NewOverwriteRequestV2RulesWithDefaults() *OverwriteRequestV2Rules`

NewOverwriteRequestV2RulesWithDefaults instantiates a new OverwriteRequestV2Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *OverwriteRequestV2Rules) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *OverwriteRequestV2Rules) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *OverwriteRequestV2Rules) SetRule(v string)`

SetRule sets Rule field to given value.


### GetPosition

`func (o *OverwriteRequestV2Rules) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *OverwriteRequestV2Rules) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *OverwriteRequestV2Rules) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *OverwriteRequestV2Rules) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComment

`func (o *OverwriteRequestV2Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OverwriteRequestV2Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OverwriteRequestV2Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OverwriteRequestV2Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *OverwriteRequestV2Rules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OverwriteRequestV2Rules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OverwriteRequestV2Rules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OverwriteRequestV2Rules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroups

`func (o *OverwriteRequestV2Rules) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OverwriteRequestV2Rules) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OverwriteRequestV2Rules) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *OverwriteRequestV2Rules) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDisabled

`func (o *OverwriteRequestV2Rules) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *OverwriteRequestV2Rules) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *OverwriteRequestV2Rules) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *OverwriteRequestV2Rules) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



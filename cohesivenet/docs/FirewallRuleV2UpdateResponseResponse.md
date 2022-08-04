# FirewallRuleV2UpdateResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unique random id | [optional] 
**Rule** | Pointer to **string** | Unresolved rule | [optional] 
**RuleResolved** | Pointer to **string** | Rule with variables resolved | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**LastResolved** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Groups** | Pointer to **[]string** | List of groups that this rule is in | [optional] 
**OsRules** | Pointer to [**[]FirewallRuleV2OsRulesInner**](FirewallRuleV2OsRulesInner.md) | The actuall firewall rules enforced at the operating system level | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirewallRuleV2UpdateResponseResponse

`func NewFirewallRuleV2UpdateResponseResponse() *FirewallRuleV2UpdateResponseResponse`

NewFirewallRuleV2UpdateResponseResponse instantiates a new FirewallRuleV2UpdateResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleV2UpdateResponseResponseWithDefaults

`func NewFirewallRuleV2UpdateResponseResponseWithDefaults() *FirewallRuleV2UpdateResponseResponse`

NewFirewallRuleV2UpdateResponseResponseWithDefaults instantiates a new FirewallRuleV2UpdateResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRuleV2UpdateResponseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRuleV2UpdateResponseResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallRuleV2UpdateResponseResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRule

`func (o *FirewallRuleV2UpdateResponseResponse) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallRuleV2UpdateResponseResponse) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FirewallRuleV2UpdateResponseResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetRuleResolved

`func (o *FirewallRuleV2UpdateResponseResponse) GetRuleResolved() string`

GetRuleResolved returns the RuleResolved field if non-nil, zero value otherwise.

### GetRuleResolvedOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetRuleResolvedOk() (*string, bool)`

GetRuleResolvedOk returns a tuple with the RuleResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleResolved

`func (o *FirewallRuleV2UpdateResponseResponse) SetRuleResolved(v string)`

SetRuleResolved sets RuleResolved field to given value.

### HasRuleResolved

`func (o *FirewallRuleV2UpdateResponseResponse) HasRuleResolved() bool`

HasRuleResolved returns a boolean if a field has been set.

### GetTable

`func (o *FirewallRuleV2UpdateResponseResponse) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *FirewallRuleV2UpdateResponseResponse) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *FirewallRuleV2UpdateResponseResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallRuleV2UpdateResponseResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallRuleV2UpdateResponseResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallRuleV2UpdateResponseResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComment

`func (o *FirewallRuleV2UpdateResponseResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FirewallRuleV2UpdateResponseResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FirewallRuleV2UpdateResponseResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLastResolved

`func (o *FirewallRuleV2UpdateResponseResponse) GetLastResolved() string`

GetLastResolved returns the LastResolved field if non-nil, zero value otherwise.

### GetLastResolvedOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetLastResolvedOk() (*string, bool)`

GetLastResolvedOk returns a tuple with the LastResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolved

`func (o *FirewallRuleV2UpdateResponseResponse) SetLastResolved(v string)`

SetLastResolved sets LastResolved field to given value.

### HasLastResolved

`func (o *FirewallRuleV2UpdateResponseResponse) HasLastResolved() bool`

HasLastResolved returns a boolean if a field has been set.

### GetDisabled

`func (o *FirewallRuleV2UpdateResponseResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FirewallRuleV2UpdateResponseResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FirewallRuleV2UpdateResponseResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallRuleV2UpdateResponseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallRuleV2UpdateResponseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallRuleV2UpdateResponseResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGroups

`func (o *FirewallRuleV2UpdateResponseResponse) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FirewallRuleV2UpdateResponseResponse) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FirewallRuleV2UpdateResponseResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetOsRules

`func (o *FirewallRuleV2UpdateResponseResponse) GetOsRules() []FirewallRuleV2OsRulesInner`

GetOsRules returns the OsRules field if non-nil, zero value otherwise.

### GetOsRulesOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetOsRulesOk() (*[]FirewallRuleV2OsRulesInner, bool)`

GetOsRulesOk returns a tuple with the OsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRules

`func (o *FirewallRuleV2UpdateResponseResponse) SetOsRules(v []FirewallRuleV2OsRulesInner)`

SetOsRules sets OsRules field to given value.

### HasOsRules

`func (o *FirewallRuleV2UpdateResponseResponse) HasOsRules() bool`

HasOsRules returns a boolean if a field has been set.

### GetErrors

`func (o *FirewallRuleV2UpdateResponseResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FirewallRuleV2UpdateResponseResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FirewallRuleV2UpdateResponseResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FirewallRuleV2UpdateResponseResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



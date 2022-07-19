# FirewallRuleV2CreateMultipleResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]FirewallRuleV2**](FirewallRuleV2.md) |  | [optional] 
**Errors** | Pointer to **[]string** | List of errors for any rules that failed | [optional] 

## Methods

### NewFirewallRuleV2CreateMultipleResponseResponse

`func NewFirewallRuleV2CreateMultipleResponseResponse() *FirewallRuleV2CreateMultipleResponseResponse`

NewFirewallRuleV2CreateMultipleResponseResponse instantiates a new FirewallRuleV2CreateMultipleResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleV2CreateMultipleResponseResponseWithDefaults

`func NewFirewallRuleV2CreateMultipleResponseResponseWithDefaults() *FirewallRuleV2CreateMultipleResponseResponse`

NewFirewallRuleV2CreateMultipleResponseResponseWithDefaults instantiates a new FirewallRuleV2CreateMultipleResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FirewallRuleV2CreateMultipleResponseResponse) GetRules() []FirewallRuleV2`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallRuleV2CreateMultipleResponseResponse) GetRulesOk() (*[]FirewallRuleV2, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallRuleV2CreateMultipleResponseResponse) SetRules(v []FirewallRuleV2)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallRuleV2CreateMultipleResponseResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetErrors

`func (o *FirewallRuleV2CreateMultipleResponseResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FirewallRuleV2CreateMultipleResponseResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FirewallRuleV2CreateMultipleResponseResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FirewallRuleV2CreateMultipleResponseResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



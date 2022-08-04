# FirewallBulkWriteResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]FirewallRuleV2**](FirewallRuleV2.md) |  | [optional] 
**Errors** | Pointer to [**[]FirewallBulkWriteResponseResponseErrorsInner**](FirewallBulkWriteResponseResponseErrorsInner.md) | List of rules that errored | [optional] 

## Methods

### NewFirewallBulkWriteResponseResponse

`func NewFirewallBulkWriteResponseResponse() *FirewallBulkWriteResponseResponse`

NewFirewallBulkWriteResponseResponse instantiates a new FirewallBulkWriteResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBulkWriteResponseResponseWithDefaults

`func NewFirewallBulkWriteResponseResponseWithDefaults() *FirewallBulkWriteResponseResponse`

NewFirewallBulkWriteResponseResponseWithDefaults instantiates a new FirewallBulkWriteResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FirewallBulkWriteResponseResponse) GetRules() []FirewallRuleV2`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallBulkWriteResponseResponse) GetRulesOk() (*[]FirewallRuleV2, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallBulkWriteResponseResponse) SetRules(v []FirewallRuleV2)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallBulkWriteResponseResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetErrors

`func (o *FirewallBulkWriteResponseResponse) GetErrors() []FirewallBulkWriteResponseResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FirewallBulkWriteResponseResponse) GetErrorsOk() (*[]FirewallBulkWriteResponseResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FirewallBulkWriteResponseResponse) SetErrors(v []FirewallBulkWriteResponseResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FirewallBulkWriteResponseResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



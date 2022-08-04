# FirewallRuleOperationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Rule** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** | Task token | [optional] 

## Methods

### NewFirewallRuleOperationData

`func NewFirewallRuleOperationData() *FirewallRuleOperationData`

NewFirewallRuleOperationData instantiates a new FirewallRuleOperationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleOperationDataWithDefaults

`func NewFirewallRuleOperationDataWithDefaults() *FirewallRuleOperationData`

NewFirewallRuleOperationDataWithDefaults instantiates a new FirewallRuleOperationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FirewallRuleOperationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRuleOperationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRuleOperationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallRuleOperationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRule

`func (o *FirewallRuleOperationData) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FirewallRuleOperationData) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FirewallRuleOperationData) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FirewallRuleOperationData) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallRuleOperationData) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallRuleOperationData) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallRuleOperationData) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallRuleOperationData) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetToken

`func (o *FirewallRuleOperationData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FirewallRuleOperationData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FirewallRuleOperationData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FirewallRuleOperationData) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



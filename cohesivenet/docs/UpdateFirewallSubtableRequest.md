# UpdateFirewallSubtableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]CreateFirewallSubtableRequestRulesInner**](CreateFirewallSubtableRequestRulesInner.md) |  | [optional] 

## Methods

### NewUpdateFirewallSubtableRequest

`func NewUpdateFirewallSubtableRequest() *UpdateFirewallSubtableRequest`

NewUpdateFirewallSubtableRequest instantiates a new UpdateFirewallSubtableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallSubtableRequestWithDefaults

`func NewUpdateFirewallSubtableRequestWithDefaults() *UpdateFirewallSubtableRequest`

NewUpdateFirewallSubtableRequestWithDefaults instantiates a new UpdateFirewallSubtableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateFirewallSubtableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFirewallSubtableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFirewallSubtableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFirewallSubtableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *UpdateFirewallSubtableRequest) GetRules() []CreateFirewallSubtableRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateFirewallSubtableRequest) GetRulesOk() (*[]CreateFirewallSubtableRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateFirewallSubtableRequest) SetRules(v []CreateFirewallSubtableRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateFirewallSubtableRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



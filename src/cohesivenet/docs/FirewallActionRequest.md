# FirewallActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to take. Currently only reset_connection_tracking supported | 

## Methods

### NewFirewallActionRequest

`func NewFirewallActionRequest(action string, ) *FirewallActionRequest`

NewFirewallActionRequest instantiates a new FirewallActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallActionRequestWithDefaults

`func NewFirewallActionRequestWithDefaults() *FirewallActionRequest`

NewFirewallActionRequestWithDefaults instantiates a new FirewallActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FirewallActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallActionRequest) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



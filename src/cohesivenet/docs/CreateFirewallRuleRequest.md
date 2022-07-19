# CreateFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** | New firewall rule string that needs to be compatible with a Linux \&quot;iptables\&quot; statement | 
**Position** | Pointer to **int32** | Position which the rule will be inserted in the list of Firewall rules.  Default is -1, which will post as the next rule in the list  | [optional] [default to -1]

## Methods

### NewCreateFirewallRuleRequest

`func NewCreateFirewallRuleRequest(rule string, ) *CreateFirewallRuleRequest`

NewCreateFirewallRuleRequest instantiates a new CreateFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRuleRequestWithDefaults

`func NewCreateFirewallRuleRequestWithDefaults() *CreateFirewallRuleRequest`

NewCreateFirewallRuleRequestWithDefaults instantiates a new CreateFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *CreateFirewallRuleRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateFirewallRuleRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateFirewallRuleRequest) SetRule(v string)`

SetRule sets Rule field to given value.


### GetPosition

`func (o *CreateFirewallRuleRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateFirewallRuleRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateFirewallRuleRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateFirewallRuleRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



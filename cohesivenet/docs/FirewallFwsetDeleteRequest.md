# FirewallFwsetDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **string** | Chained firewall rules seperated by \&quot;\\n\&quot; | [optional] 
**Name** | Pointer to **string** | Name of the Fwset. Must be valid chain that begins with one of the following: NETS_, PORTS_, LIST_.  | [optional] 

## Methods

### NewFirewallFwsetDeleteRequest

`func NewFirewallFwsetDeleteRequest() *FirewallFwsetDeleteRequest`

NewFirewallFwsetDeleteRequest instantiates a new FirewallFwsetDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFwsetDeleteRequestWithDefaults

`func NewFirewallFwsetDeleteRequestWithDefaults() *FirewallFwsetDeleteRequest`

NewFirewallFwsetDeleteRequestWithDefaults instantiates a new FirewallFwsetDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FirewallFwsetDeleteRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallFwsetDeleteRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallFwsetDeleteRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallFwsetDeleteRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetName

`func (o *FirewallFwsetDeleteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallFwsetDeleteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallFwsetDeleteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallFwsetDeleteRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



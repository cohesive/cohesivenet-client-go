# FirewallSubgroupDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **string** | Chained firewall rules seperated by \&quot;\\n\&quot; | [optional] 
**Name** | Pointer to **string** | Name of the subgroup chain. Must be valid chain name. | [optional] 

## Methods

### NewFirewallSubgroupDeleteRequest

`func NewFirewallSubgroupDeleteRequest() *FirewallSubgroupDeleteRequest`

NewFirewallSubgroupDeleteRequest instantiates a new FirewallSubgroupDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubgroupDeleteRequestWithDefaults

`func NewFirewallSubgroupDeleteRequestWithDefaults() *FirewallSubgroupDeleteRequest`

NewFirewallSubgroupDeleteRequestWithDefaults instantiates a new FirewallSubgroupDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FirewallSubgroupDeleteRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallSubgroupDeleteRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallSubgroupDeleteRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallSubgroupDeleteRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetName

`func (o *FirewallSubgroupDeleteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallSubgroupDeleteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallSubgroupDeleteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallSubgroupDeleteRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



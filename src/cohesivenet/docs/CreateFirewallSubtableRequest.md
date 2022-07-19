# CreateFirewallSubtableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of subtable. Alphanumeric, _, - characters allowed. No spaces. | 
**Type** | **string** | Valid table. This determines which main tables can route to this subtable. | 
**Description** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]CreateFirewallSubtableRequestRulesInner**](CreateFirewallSubtableRequestRulesInner.md) |  | [optional] 

## Methods

### NewCreateFirewallSubtableRequest

`func NewCreateFirewallSubtableRequest(name string, type_ string, ) *CreateFirewallSubtableRequest`

NewCreateFirewallSubtableRequest instantiates a new CreateFirewallSubtableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallSubtableRequestWithDefaults

`func NewCreateFirewallSubtableRequestWithDefaults() *CreateFirewallSubtableRequest`

NewCreateFirewallSubtableRequestWithDefaults instantiates a new CreateFirewallSubtableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirewallSubtableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirewallSubtableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirewallSubtableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateFirewallSubtableRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFirewallSubtableRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFirewallSubtableRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CreateFirewallSubtableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirewallSubtableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirewallSubtableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirewallSubtableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *CreateFirewallSubtableRequest) GetRules() []CreateFirewallSubtableRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateFirewallSubtableRequest) GetRulesOk() (*[]CreateFirewallSubtableRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateFirewallSubtableRequest) SetRules(v []CreateFirewallSubtableRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateFirewallSubtableRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



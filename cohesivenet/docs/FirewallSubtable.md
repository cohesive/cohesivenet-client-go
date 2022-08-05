# FirewallSubtable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** | table this subtable can be routed to from such as prerouting | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Rules** | Pointer to [**[]FirewallSubtableRule**](FirewallSubtableRule.md) |  | [optional] 

## Methods

### NewFirewallSubtable

`func NewFirewallSubtable() *FirewallSubtable`

NewFirewallSubtable instantiates a new FirewallSubtable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubtableWithDefaults

`func NewFirewallSubtableWithDefaults() *FirewallSubtable`

NewFirewallSubtableWithDefaults instantiates a new FirewallSubtable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallSubtable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallSubtable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallSubtable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallSubtable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallSubtable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallSubtable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallSubtable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallSubtable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallSubtable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallSubtable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallSubtable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallSubtable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirewallSubtable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallSubtable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallSubtable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirewallSubtable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *FirewallSubtable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallSubtable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallSubtable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallSubtable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FirewallSubtable) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallSubtable) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallSubtable) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirewallSubtable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRules

`func (o *FirewallSubtable) GetRules() []FirewallSubtableRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallSubtable) GetRulesOk() (*[]FirewallSubtableRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallSubtable) SetRules(v []FirewallSubtableRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallSubtable) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



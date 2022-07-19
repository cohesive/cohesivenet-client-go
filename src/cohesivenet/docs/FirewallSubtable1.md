# FirewallSubtable1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** | table this subtable can be routed to from such as prerouting | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Rules** | Pointer to [**[]Object**](Object.md) |  | [optional] 

## Methods

### NewFirewallSubtable1

`func NewFirewallSubtable1() *FirewallSubtable1`

NewFirewallSubtable1 instantiates a new FirewallSubtable1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubtable1WithDefaults

`func NewFirewallSubtable1WithDefaults() *FirewallSubtable1`

NewFirewallSubtable1WithDefaults instantiates a new FirewallSubtable1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallSubtable1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallSubtable1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallSubtable1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallSubtable1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallSubtable1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallSubtable1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallSubtable1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallSubtable1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallSubtable1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallSubtable1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallSubtable1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallSubtable1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirewallSubtable1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallSubtable1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallSubtable1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirewallSubtable1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *FirewallSubtable1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallSubtable1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallSubtable1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallSubtable1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FirewallSubtable1) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallSubtable1) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallSubtable1) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirewallSubtable1) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRules

`func (o *FirewallSubtable1) GetRules() []Object`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallSubtable1) GetRulesOk() (*[]Object, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallSubtable1) SetRules(v []Object)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallSubtable1) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



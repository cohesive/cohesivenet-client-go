# FirewallFwSetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** | Valid ipset type. Supported are net, port, list and service | [optional] 
**Size** | Pointer to **int32** | number of entries in set | [optional] 
**Entries** | Pointer to [**[]Object**](Object.md) |  | [optional] 

## Methods

### NewFirewallFwSetV2

`func NewFirewallFwSetV2() *FirewallFwSetV2`

NewFirewallFwSetV2 instantiates a new FirewallFwSetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFwSetV2WithDefaults

`func NewFirewallFwSetV2WithDefaults() *FirewallFwSetV2`

NewFirewallFwSetV2WithDefaults instantiates a new FirewallFwSetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallFwSetV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallFwSetV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallFwSetV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallFwSetV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallFwSetV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallFwSetV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallFwSetV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallFwSetV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallFwSetV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallFwSetV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallFwSetV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallFwSetV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirewallFwSetV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallFwSetV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallFwSetV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirewallFwSetV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *FirewallFwSetV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallFwSetV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallFwSetV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallFwSetV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FirewallFwSetV2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallFwSetV2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallFwSetV2) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirewallFwSetV2) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetEntries

`func (o *FirewallFwSetV2) GetEntries() []Object`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *FirewallFwSetV2) GetEntriesOk() (*[]Object, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *FirewallFwSetV2) SetEntries(v []Object)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *FirewallFwSetV2) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FirewallSubtableCreateResponseResponse

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
**Errors** | Pointer to [**[]FirewallSubtableCreateResponseResponseAllOfErrorsInner**](FirewallSubtableCreateResponseResponseAllOfErrorsInner.md) | Rule objects that failed insertion into subtable | [optional] 

## Methods

### NewFirewallSubtableCreateResponseResponse

`func NewFirewallSubtableCreateResponseResponse() *FirewallSubtableCreateResponseResponse`

NewFirewallSubtableCreateResponseResponse instantiates a new FirewallSubtableCreateResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSubtableCreateResponseResponseWithDefaults

`func NewFirewallSubtableCreateResponseResponseWithDefaults() *FirewallSubtableCreateResponseResponse`

NewFirewallSubtableCreateResponseResponseWithDefaults instantiates a new FirewallSubtableCreateResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallSubtableCreateResponseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallSubtableCreateResponseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallSubtableCreateResponseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallSubtableCreateResponseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallSubtableCreateResponseResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallSubtableCreateResponseResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallSubtableCreateResponseResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallSubtableCreateResponseResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallSubtableCreateResponseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallSubtableCreateResponseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallSubtableCreateResponseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallSubtableCreateResponseResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirewallSubtableCreateResponseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallSubtableCreateResponseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallSubtableCreateResponseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirewallSubtableCreateResponseResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *FirewallSubtableCreateResponseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallSubtableCreateResponseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallSubtableCreateResponseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallSubtableCreateResponseResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FirewallSubtableCreateResponseResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallSubtableCreateResponseResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallSubtableCreateResponseResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirewallSubtableCreateResponseResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRules

`func (o *FirewallSubtableCreateResponseResponse) GetRules() []Object`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallSubtableCreateResponseResponse) GetRulesOk() (*[]Object, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallSubtableCreateResponseResponse) SetRules(v []Object)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallSubtableCreateResponseResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetErrors

`func (o *FirewallSubtableCreateResponseResponse) GetErrors() []FirewallSubtableCreateResponseResponseAllOfErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FirewallSubtableCreateResponseResponse) GetErrorsOk() (*[]FirewallSubtableCreateResponseResponseAllOfErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FirewallSubtableCreateResponseResponse) SetErrors(v []FirewallSubtableCreateResponseResponseAllOfErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FirewallSubtableCreateResponseResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



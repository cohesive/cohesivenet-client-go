# FirewallFwsetSaveResponseResponse

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
**Errors** | Pointer to [**[]FirewallFwsetSaveResponseResponseAllOfErrorsInner**](FirewallFwsetSaveResponseResponseAllOfErrorsInner.md) | entry objects that failed insertion into fwset | [optional] 

## Methods

### NewFirewallFwsetSaveResponseResponse

`func NewFirewallFwsetSaveResponseResponse() *FirewallFwsetSaveResponseResponse`

NewFirewallFwsetSaveResponseResponse instantiates a new FirewallFwsetSaveResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFwsetSaveResponseResponseWithDefaults

`func NewFirewallFwsetSaveResponseResponseWithDefaults() *FirewallFwsetSaveResponseResponse`

NewFirewallFwsetSaveResponseResponseWithDefaults instantiates a new FirewallFwsetSaveResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallFwsetSaveResponseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallFwsetSaveResponseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallFwsetSaveResponseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallFwsetSaveResponseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FirewallFwsetSaveResponseResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallFwsetSaveResponseResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallFwsetSaveResponseResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallFwsetSaveResponseResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallFwsetSaveResponseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallFwsetSaveResponseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallFwsetSaveResponseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallFwsetSaveResponseResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirewallFwsetSaveResponseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallFwsetSaveResponseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallFwsetSaveResponseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirewallFwsetSaveResponseResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *FirewallFwsetSaveResponseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallFwsetSaveResponseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallFwsetSaveResponseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallFwsetSaveResponseResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FirewallFwsetSaveResponseResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallFwsetSaveResponseResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallFwsetSaveResponseResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FirewallFwsetSaveResponseResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetEntries

`func (o *FirewallFwsetSaveResponseResponse) GetEntries() []Object`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *FirewallFwsetSaveResponseResponse) GetEntriesOk() (*[]Object, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *FirewallFwsetSaveResponseResponse) SetEntries(v []Object)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *FirewallFwsetSaveResponseResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetErrors

`func (o *FirewallFwsetSaveResponseResponse) GetErrors() []FirewallFwsetSaveResponseResponseAllOfErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FirewallFwsetSaveResponseResponse) GetErrorsOk() (*[]FirewallFwsetSaveResponseResponseAllOfErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FirewallFwsetSaveResponseResponse) SetErrors(v []FirewallFwsetSaveResponseResponseAllOfErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FirewallFwsetSaveResponseResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



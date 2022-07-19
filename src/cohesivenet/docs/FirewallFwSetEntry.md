# FirewallFwSetEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**EntryResolved** | Pointer to **string** |  | [optional] 
**LastResolved** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFirewallFwSetEntry

`func NewFirewallFwSetEntry() *FirewallFwSetEntry`

NewFirewallFwSetEntry instantiates a new FirewallFwSetEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallFwSetEntryWithDefaults

`func NewFirewallFwSetEntryWithDefaults() *FirewallFwSetEntry`

NewFirewallFwSetEntryWithDefaults instantiates a new FirewallFwSetEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *FirewallFwSetEntry) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *FirewallFwSetEntry) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *FirewallFwSetEntry) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *FirewallFwSetEntry) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetComment

`func (o *FirewallFwSetEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FirewallFwSetEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FirewallFwSetEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FirewallFwSetEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEntryResolved

`func (o *FirewallFwSetEntry) GetEntryResolved() string`

GetEntryResolved returns the EntryResolved field if non-nil, zero value otherwise.

### GetEntryResolvedOk

`func (o *FirewallFwSetEntry) GetEntryResolvedOk() (*string, bool)`

GetEntryResolvedOk returns a tuple with the EntryResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryResolved

`func (o *FirewallFwSetEntry) SetEntryResolved(v string)`

SetEntryResolved sets EntryResolved field to given value.

### HasEntryResolved

`func (o *FirewallFwSetEntry) HasEntryResolved() bool`

HasEntryResolved returns a boolean if a field has been set.

### GetLastResolved

`func (o *FirewallFwSetEntry) GetLastResolved() time.Time`

GetLastResolved returns the LastResolved field if non-nil, zero value otherwise.

### GetLastResolvedOk

`func (o *FirewallFwSetEntry) GetLastResolvedOk() (*time.Time, bool)`

GetLastResolvedOk returns a tuple with the LastResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolved

`func (o *FirewallFwSetEntry) SetLastResolved(v time.Time)`

SetLastResolved sets LastResolved field to given value.

### HasLastResolved

`func (o *FirewallFwSetEntry) HasLastResolved() bool`

HasLastResolved returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FirewallFwSetEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallFwSetEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallFwSetEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallFwSetEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddFirewallEntryToFwsetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | **string** | resolvable entry for fwset | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewAddFirewallEntryToFwsetRequest

`func NewAddFirewallEntryToFwsetRequest(entry string, ) *AddFirewallEntryToFwsetRequest`

NewAddFirewallEntryToFwsetRequest instantiates a new AddFirewallEntryToFwsetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFirewallEntryToFwsetRequestWithDefaults

`func NewAddFirewallEntryToFwsetRequestWithDefaults() *AddFirewallEntryToFwsetRequest`

NewAddFirewallEntryToFwsetRequestWithDefaults instantiates a new AddFirewallEntryToFwsetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *AddFirewallEntryToFwsetRequest) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *AddFirewallEntryToFwsetRequest) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *AddFirewallEntryToFwsetRequest) SetEntry(v string)`

SetEntry sets Entry field to given value.


### GetComment

`func (o *AddFirewallEntryToFwsetRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddFirewallEntryToFwsetRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddFirewallEntryToFwsetRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddFirewallEntryToFwsetRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



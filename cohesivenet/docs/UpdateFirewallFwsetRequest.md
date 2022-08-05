# UpdateFirewallFwsetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of fwset. Alphanumeric, _, - characters allowed. No spaces. This will change resource URI. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]CreateFirewallFwsetRequestEntriesOneOfInner**](CreateFirewallFwsetRequestEntriesOneOfInner.md) |  | [optional] 

## Methods

### NewUpdateFirewallFwsetRequest

`func NewUpdateFirewallFwsetRequest() *UpdateFirewallFwsetRequest`

NewUpdateFirewallFwsetRequest instantiates a new UpdateFirewallFwsetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallFwsetRequestWithDefaults

`func NewUpdateFirewallFwsetRequestWithDefaults() *UpdateFirewallFwsetRequest`

NewUpdateFirewallFwsetRequestWithDefaults instantiates a new UpdateFirewallFwsetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFirewallFwsetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirewallFwsetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirewallFwsetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFirewallFwsetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateFirewallFwsetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFirewallFwsetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFirewallFwsetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFirewallFwsetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *UpdateFirewallFwsetRequest) GetEntries() []CreateFirewallFwsetRequestEntriesOneOfInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *UpdateFirewallFwsetRequest) GetEntriesOk() (*[]CreateFirewallFwsetRequestEntriesOneOfInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *UpdateFirewallFwsetRequest) SetEntries(v []CreateFirewallFwsetRequestEntriesOneOfInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *UpdateFirewallFwsetRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFirewallFwsetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of fwset. Alphanumeric, _, - characters allowed. No spaces. | 
**Type** | **string** | Valid fwset type. This determines the type of data in the set. | 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**CreateFirewallFwsetRequestEntries**](CreateFirewallFwsetRequestEntries.md) |  | [optional] 

## Methods

### NewCreateFirewallFwsetRequest

`func NewCreateFirewallFwsetRequest(name string, type_ string, ) *CreateFirewallFwsetRequest`

NewCreateFirewallFwsetRequest instantiates a new CreateFirewallFwsetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallFwsetRequestWithDefaults

`func NewCreateFirewallFwsetRequestWithDefaults() *CreateFirewallFwsetRequest`

NewCreateFirewallFwsetRequestWithDefaults instantiates a new CreateFirewallFwsetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirewallFwsetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirewallFwsetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirewallFwsetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateFirewallFwsetRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFirewallFwsetRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFirewallFwsetRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CreateFirewallFwsetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirewallFwsetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirewallFwsetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirewallFwsetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *CreateFirewallFwsetRequest) GetEntries() CreateFirewallFwsetRequestEntries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CreateFirewallFwsetRequest) GetEntriesOk() (*CreateFirewallFwsetRequestEntries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CreateFirewallFwsetRequest) SetEntries(v CreateFirewallFwsetRequestEntries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *CreateFirewallFwsetRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



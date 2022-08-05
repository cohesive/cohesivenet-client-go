# UpdateClientpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the clientpack (IP snake case) | 
**Enabled** | Pointer to **bool** | Enable or disable clientpack. | [optional] 
**CheckedOut** | Pointer to **bool** | Update whether clientpack is checked out and thus unavailable | [optional] 
**Regenerate** | Pointer to **bool** | Regenerate clientpack file. Returns a task token in the response. | [optional] 
**Compression** | Pointer to **string** | Turn compression on or off for all. Can be \&quot;on\&quot; or \&quot;off\&quot; currently. | [optional] 
**AuthOn** | Pointer to **bool** | Enable authentication for clientpack. | [optional] 
**PskOn** | Pointer to **bool** | Enable PSK for clientpack. Wireguard only. | [optional] 

## Methods

### NewUpdateClientpackRequest

`func NewUpdateClientpackRequest(name string, ) *UpdateClientpackRequest`

NewUpdateClientpackRequest instantiates a new UpdateClientpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClientpackRequestWithDefaults

`func NewUpdateClientpackRequestWithDefaults() *UpdateClientpackRequest`

NewUpdateClientpackRequestWithDefaults instantiates a new UpdateClientpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClientpackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClientpackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClientpackRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UpdateClientpackRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateClientpackRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateClientpackRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateClientpackRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCheckedOut

`func (o *UpdateClientpackRequest) GetCheckedOut() bool`

GetCheckedOut returns the CheckedOut field if non-nil, zero value otherwise.

### GetCheckedOutOk

`func (o *UpdateClientpackRequest) GetCheckedOutOk() (*bool, bool)`

GetCheckedOutOk returns a tuple with the CheckedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedOut

`func (o *UpdateClientpackRequest) SetCheckedOut(v bool)`

SetCheckedOut sets CheckedOut field to given value.

### HasCheckedOut

`func (o *UpdateClientpackRequest) HasCheckedOut() bool`

HasCheckedOut returns a boolean if a field has been set.

### GetRegenerate

`func (o *UpdateClientpackRequest) GetRegenerate() bool`

GetRegenerate returns the Regenerate field if non-nil, zero value otherwise.

### GetRegenerateOk

`func (o *UpdateClientpackRequest) GetRegenerateOk() (*bool, bool)`

GetRegenerateOk returns a tuple with the Regenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerate

`func (o *UpdateClientpackRequest) SetRegenerate(v bool)`

SetRegenerate sets Regenerate field to given value.

### HasRegenerate

`func (o *UpdateClientpackRequest) HasRegenerate() bool`

HasRegenerate returns a boolean if a field has been set.

### GetCompression

`func (o *UpdateClientpackRequest) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *UpdateClientpackRequest) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *UpdateClientpackRequest) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *UpdateClientpackRequest) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetAuthOn

`func (o *UpdateClientpackRequest) GetAuthOn() bool`

GetAuthOn returns the AuthOn field if non-nil, zero value otherwise.

### GetAuthOnOk

`func (o *UpdateClientpackRequest) GetAuthOnOk() (*bool, bool)`

GetAuthOnOk returns a tuple with the AuthOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOn

`func (o *UpdateClientpackRequest) SetAuthOn(v bool)`

SetAuthOn sets AuthOn field to given value.

### HasAuthOn

`func (o *UpdateClientpackRequest) HasAuthOn() bool`

HasAuthOn returns a boolean if a field has been set.

### GetPskOn

`func (o *UpdateClientpackRequest) GetPskOn() bool`

GetPskOn returns the PskOn field if non-nil, zero value otherwise.

### GetPskOnOk

`func (o *UpdateClientpackRequest) GetPskOnOk() (*bool, bool)`

GetPskOnOk returns a tuple with the PskOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskOn

`func (o *UpdateClientpackRequest) SetPskOn(v bool)`

SetPskOn sets PskOn field to given value.

### HasPskOn

`func (o *UpdateClientpackRequest) HasPskOn() bool`

HasPskOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



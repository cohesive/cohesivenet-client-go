# LinkHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remote** | Pointer to **string** |  | [optional] 
**Local** | Pointer to **string** |  | [optional] 
**Tunnelid** | Pointer to **int32** |  | [optional] 
**History** | Pointer to [**[]LinkEvent**](LinkEvent.md) |  | [optional] 

## Methods

### NewLinkHistory

`func NewLinkHistory() *LinkHistory`

NewLinkHistory instantiates a new LinkHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkHistoryWithDefaults

`func NewLinkHistoryWithDefaults() *LinkHistory`

NewLinkHistoryWithDefaults instantiates a new LinkHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemote

`func (o *LinkHistory) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *LinkHistory) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *LinkHistory) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *LinkHistory) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetLocal

`func (o *LinkHistory) GetLocal() string`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *LinkHistory) GetLocalOk() (*string, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *LinkHistory) SetLocal(v string)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *LinkHistory) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetTunnelid

`func (o *LinkHistory) GetTunnelid() int32`

GetTunnelid returns the Tunnelid field if non-nil, zero value otherwise.

### GetTunnelidOk

`func (o *LinkHistory) GetTunnelidOk() (*int32, bool)`

GetTunnelidOk returns a tuple with the Tunnelid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelid

`func (o *LinkHistory) SetTunnelid(v int32)`

SetTunnelid sets Tunnelid field to given value.

### HasTunnelid

`func (o *LinkHistory) HasTunnelid() bool`

HasTunnelid returns a boolean if a field has been set.

### GetHistory

`func (o *LinkHistory) GetHistory() []LinkEvent`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *LinkHistory) GetHistoryOk() (*[]LinkEvent, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *LinkHistory) SetHistory(v []LinkEvent)`

SetHistory sets History field to given value.

### HasHistory

`func (o *LinkHistory) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



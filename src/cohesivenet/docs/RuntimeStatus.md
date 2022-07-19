# RuntimeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedClients** | Pointer to [**map[string]RuntimeStatusConnectedClientsValue**](RuntimeStatusConnectedClientsValue.md) | clients keyed by ip address | [optional] 
**ConnectedSubnets** | Pointer to **[][]string** | Array of arrays with each element of length 2 representing [network, subnet mask] | [optional] 
**Ipsec** | Pointer to [**map[string]RuntimeStatusIpsecValue**](RuntimeStatusIpsecValue.md) | IPSEC tunnels keyed by tunnel ID | [optional] 

## Methods

### NewRuntimeStatus

`func NewRuntimeStatus() *RuntimeStatus`

NewRuntimeStatus instantiates a new RuntimeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeStatusWithDefaults

`func NewRuntimeStatusWithDefaults() *RuntimeStatus`

NewRuntimeStatusWithDefaults instantiates a new RuntimeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedClients

`func (o *RuntimeStatus) GetConnectedClients() map[string]RuntimeStatusConnectedClientsValue`

GetConnectedClients returns the ConnectedClients field if non-nil, zero value otherwise.

### GetConnectedClientsOk

`func (o *RuntimeStatus) GetConnectedClientsOk() (*map[string]RuntimeStatusConnectedClientsValue, bool)`

GetConnectedClientsOk returns a tuple with the ConnectedClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedClients

`func (o *RuntimeStatus) SetConnectedClients(v map[string]RuntimeStatusConnectedClientsValue)`

SetConnectedClients sets ConnectedClients field to given value.

### HasConnectedClients

`func (o *RuntimeStatus) HasConnectedClients() bool`

HasConnectedClients returns a boolean if a field has been set.

### GetConnectedSubnets

`func (o *RuntimeStatus) GetConnectedSubnets() [][]string`

GetConnectedSubnets returns the ConnectedSubnets field if non-nil, zero value otherwise.

### GetConnectedSubnetsOk

`func (o *RuntimeStatus) GetConnectedSubnetsOk() (*[][]string, bool)`

GetConnectedSubnetsOk returns a tuple with the ConnectedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSubnets

`func (o *RuntimeStatus) SetConnectedSubnets(v [][]string)`

SetConnectedSubnets sets ConnectedSubnets field to given value.

### HasConnectedSubnets

`func (o *RuntimeStatus) HasConnectedSubnets() bool`

HasConnectedSubnets returns a boolean if a field has been set.

### GetIpsec

`func (o *RuntimeStatus) GetIpsec() map[string]RuntimeStatusIpsecValue`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *RuntimeStatus) GetIpsecOk() (*map[string]RuntimeStatusIpsecValue, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *RuntimeStatus) SetIpsec(v map[string]RuntimeStatusIpsecValue)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *RuntimeStatus) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



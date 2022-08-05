# UpdateIpsecTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounce** | Pointer to **bool** | Resets the IPsec connection for this specific tunnel | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**RemoteSubnet** | Pointer to **string** | Remote subnet for tunnel in CIDR notation | [optional] 
**LocalSubnet** | Pointer to **string** | Local subnet for tunnel in CIDR notation | [optional] 
**PingIpaddress** | Pointer to **string** | Exo Ping feature - remote IP destination of ping | [optional] 
**PingInterval** | Pointer to **int32** | Exo Ping feature - periodicy of the ping in seconds | [optional] 
**PingInterface** | Pointer to **string** | Exo Ping feature - what network interface IP of the VNS3 controller to use as the source of ping | [optional] 
**Enabled** | Pointer to **bool** | Disables tunnel if set to false | [optional] 

## Methods

### NewUpdateIpsecTunnelRequest

`func NewUpdateIpsecTunnelRequest() *UpdateIpsecTunnelRequest`

NewUpdateIpsecTunnelRequest instantiates a new UpdateIpsecTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpsecTunnelRequestWithDefaults

`func NewUpdateIpsecTunnelRequestWithDefaults() *UpdateIpsecTunnelRequest`

NewUpdateIpsecTunnelRequestWithDefaults instantiates a new UpdateIpsecTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounce

`func (o *UpdateIpsecTunnelRequest) GetBounce() bool`

GetBounce returns the Bounce field if non-nil, zero value otherwise.

### GetBounceOk

`func (o *UpdateIpsecTunnelRequest) GetBounceOk() (*bool, bool)`

GetBounceOk returns a tuple with the Bounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounce

`func (o *UpdateIpsecTunnelRequest) SetBounce(v bool)`

SetBounce sets Bounce field to given value.

### HasBounce

`func (o *UpdateIpsecTunnelRequest) HasBounce() bool`

HasBounce returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIpsecTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIpsecTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIpsecTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIpsecTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *UpdateIpsecTunnelRequest) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *UpdateIpsecTunnelRequest) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *UpdateIpsecTunnelRequest) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *UpdateIpsecTunnelRequest) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetLocalSubnet

`func (o *UpdateIpsecTunnelRequest) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *UpdateIpsecTunnelRequest) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *UpdateIpsecTunnelRequest) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *UpdateIpsecTunnelRequest) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetPingIpaddress

`func (o *UpdateIpsecTunnelRequest) GetPingIpaddress() string`

GetPingIpaddress returns the PingIpaddress field if non-nil, zero value otherwise.

### GetPingIpaddressOk

`func (o *UpdateIpsecTunnelRequest) GetPingIpaddressOk() (*string, bool)`

GetPingIpaddressOk returns a tuple with the PingIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIpaddress

`func (o *UpdateIpsecTunnelRequest) SetPingIpaddress(v string)`

SetPingIpaddress sets PingIpaddress field to given value.

### HasPingIpaddress

`func (o *UpdateIpsecTunnelRequest) HasPingIpaddress() bool`

HasPingIpaddress returns a boolean if a field has been set.

### GetPingInterval

`func (o *UpdateIpsecTunnelRequest) GetPingInterval() int32`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *UpdateIpsecTunnelRequest) GetPingIntervalOk() (*int32, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *UpdateIpsecTunnelRequest) SetPingInterval(v int32)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *UpdateIpsecTunnelRequest) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetPingInterface

`func (o *UpdateIpsecTunnelRequest) GetPingInterface() string`

GetPingInterface returns the PingInterface field if non-nil, zero value otherwise.

### GetPingInterfaceOk

`func (o *UpdateIpsecTunnelRequest) GetPingInterfaceOk() (*string, bool)`

GetPingInterfaceOk returns a tuple with the PingInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterface

`func (o *UpdateIpsecTunnelRequest) SetPingInterface(v string)`

SetPingInterface sets PingInterface field to given value.

### HasPingInterface

`func (o *UpdateIpsecTunnelRequest) HasPingInterface() bool`

HasPingInterface returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateIpsecTunnelRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateIpsecTunnelRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateIpsecTunnelRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateIpsecTunnelRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



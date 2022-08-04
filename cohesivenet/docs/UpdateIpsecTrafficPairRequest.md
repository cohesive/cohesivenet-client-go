# UpdateIpsecTrafficPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteSubnet** | Pointer to **string** | Remote subnet for tunnel in CIDR notation | [optional] 
**LocalSubnet** | Pointer to **string** | Local subnet for tunnel in CIDR notation | [optional] 
**PingIpaddress** | Pointer to **string** | Exo Ping feature - remote IP destination of ping | [optional] 
**PingInterval** | Pointer to **int32** | Exo Ping feature - periodicy of the ping in seconds | [optional] 
**PingInterface** | Pointer to **string** | Exo Ping feature - what network interface IP of the VNS3 controller to use as the source of ping | [optional] 
**Enabled** | Pointer to **bool** | Disables tunnel if set to false | [optional] [default to true]
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateIpsecTrafficPairRequest

`func NewUpdateIpsecTrafficPairRequest() *UpdateIpsecTrafficPairRequest`

NewUpdateIpsecTrafficPairRequest instantiates a new UpdateIpsecTrafficPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpsecTrafficPairRequestWithDefaults

`func NewUpdateIpsecTrafficPairRequestWithDefaults() *UpdateIpsecTrafficPairRequest`

NewUpdateIpsecTrafficPairRequestWithDefaults instantiates a new UpdateIpsecTrafficPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteSubnet

`func (o *UpdateIpsecTrafficPairRequest) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *UpdateIpsecTrafficPairRequest) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *UpdateIpsecTrafficPairRequest) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *UpdateIpsecTrafficPairRequest) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetLocalSubnet

`func (o *UpdateIpsecTrafficPairRequest) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *UpdateIpsecTrafficPairRequest) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *UpdateIpsecTrafficPairRequest) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *UpdateIpsecTrafficPairRequest) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetPingIpaddress

`func (o *UpdateIpsecTrafficPairRequest) GetPingIpaddress() string`

GetPingIpaddress returns the PingIpaddress field if non-nil, zero value otherwise.

### GetPingIpaddressOk

`func (o *UpdateIpsecTrafficPairRequest) GetPingIpaddressOk() (*string, bool)`

GetPingIpaddressOk returns a tuple with the PingIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIpaddress

`func (o *UpdateIpsecTrafficPairRequest) SetPingIpaddress(v string)`

SetPingIpaddress sets PingIpaddress field to given value.

### HasPingIpaddress

`func (o *UpdateIpsecTrafficPairRequest) HasPingIpaddress() bool`

HasPingIpaddress returns a boolean if a field has been set.

### GetPingInterval

`func (o *UpdateIpsecTrafficPairRequest) GetPingInterval() int32`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *UpdateIpsecTrafficPairRequest) GetPingIntervalOk() (*int32, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *UpdateIpsecTrafficPairRequest) SetPingInterval(v int32)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *UpdateIpsecTrafficPairRequest) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetPingInterface

`func (o *UpdateIpsecTrafficPairRequest) GetPingInterface() string`

GetPingInterface returns the PingInterface field if non-nil, zero value otherwise.

### GetPingInterfaceOk

`func (o *UpdateIpsecTrafficPairRequest) GetPingInterfaceOk() (*string, bool)`

GetPingInterfaceOk returns a tuple with the PingInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterface

`func (o *UpdateIpsecTrafficPairRequest) SetPingInterface(v string)`

SetPingInterface sets PingInterface field to given value.

### HasPingInterface

`func (o *UpdateIpsecTrafficPairRequest) HasPingInterface() bool`

HasPingInterface returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateIpsecTrafficPairRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateIpsecTrafficPairRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateIpsecTrafficPairRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateIpsecTrafficPairRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIpsecTrafficPairRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIpsecTrafficPairRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIpsecTrafficPairRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIpsecTrafficPairRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



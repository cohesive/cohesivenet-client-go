# CreateIpsecTrafficPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteSubnet** | **string** | Remote subnet for tunnel in CIDR notation | 
**LocalSubnet** | **string** | Local subnet for tunnel in CIDR notation | 
**PingIpaddress** | Pointer to **string** | Exo Ping feature - remote IP destination of ping | [optional] 
**PingInterval** | Pointer to **int32** | Exo Ping feature - periodicy of the ping in seconds | [optional] 
**PingInterface** | Pointer to **string** | Exo Ping feature - what network interface IP of the VNS3 controller to use as the source of ping | [optional] 
**Enabled** | Pointer to **bool** | Disables tunnel if set to false | [optional] [default to true]
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateIpsecTrafficPairRequest

`func NewCreateIpsecTrafficPairRequest(remoteSubnet string, localSubnet string, ) *CreateIpsecTrafficPairRequest`

NewCreateIpsecTrafficPairRequest instantiates a new CreateIpsecTrafficPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpsecTrafficPairRequestWithDefaults

`func NewCreateIpsecTrafficPairRequestWithDefaults() *CreateIpsecTrafficPairRequest`

NewCreateIpsecTrafficPairRequestWithDefaults instantiates a new CreateIpsecTrafficPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteSubnet

`func (o *CreateIpsecTrafficPairRequest) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *CreateIpsecTrafficPairRequest) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *CreateIpsecTrafficPairRequest) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.


### GetLocalSubnet

`func (o *CreateIpsecTrafficPairRequest) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *CreateIpsecTrafficPairRequest) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *CreateIpsecTrafficPairRequest) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.


### GetPingIpaddress

`func (o *CreateIpsecTrafficPairRequest) GetPingIpaddress() string`

GetPingIpaddress returns the PingIpaddress field if non-nil, zero value otherwise.

### GetPingIpaddressOk

`func (o *CreateIpsecTrafficPairRequest) GetPingIpaddressOk() (*string, bool)`

GetPingIpaddressOk returns a tuple with the PingIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIpaddress

`func (o *CreateIpsecTrafficPairRequest) SetPingIpaddress(v string)`

SetPingIpaddress sets PingIpaddress field to given value.

### HasPingIpaddress

`func (o *CreateIpsecTrafficPairRequest) HasPingIpaddress() bool`

HasPingIpaddress returns a boolean if a field has been set.

### GetPingInterval

`func (o *CreateIpsecTrafficPairRequest) GetPingInterval() int32`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *CreateIpsecTrafficPairRequest) GetPingIntervalOk() (*int32, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *CreateIpsecTrafficPairRequest) SetPingInterval(v int32)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *CreateIpsecTrafficPairRequest) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetPingInterface

`func (o *CreateIpsecTrafficPairRequest) GetPingInterface() string`

GetPingInterface returns the PingInterface field if non-nil, zero value otherwise.

### GetPingInterfaceOk

`func (o *CreateIpsecTrafficPairRequest) GetPingInterfaceOk() (*string, bool)`

GetPingInterfaceOk returns a tuple with the PingInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterface

`func (o *CreateIpsecTrafficPairRequest) SetPingInterface(v string)`

SetPingInterface sets PingInterface field to given value.

### HasPingInterface

`func (o *CreateIpsecTrafficPairRequest) HasPingInterface() bool`

HasPingInterface returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateIpsecTrafficPairRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateIpsecTrafficPairRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateIpsecTrafficPairRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateIpsecTrafficPairRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *CreateIpsecTrafficPairRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIpsecTrafficPairRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIpsecTrafficPairRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIpsecTrafficPairRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



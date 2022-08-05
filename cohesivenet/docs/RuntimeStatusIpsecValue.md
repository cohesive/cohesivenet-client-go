# RuntimeStatusIpsecValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LocalSubnet** | Pointer to **string** |  | [optional] 
**RemoteSubnet** | Pointer to **string** |  | [optional] 
**Endpointid** | Pointer to **int32** |  | [optional] 
**EndpointId** | Pointer to **int32** |  | [optional] 
**EndpointName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Bounce** | Pointer to **bool** | True if tunnel was just bounced | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**PingInterface** | Pointer to **string** |  | [optional] 
**PingInterval** | Pointer to **NullableInt32** | Interval for ping in seconds | [optional] 
**PingIpaddress** | Pointer to **string** |  | [optional] 
**TunnelParams** | Pointer to [**IpsecTunnelParams**](IpsecTunnelParams.md) |  | [optional] 

## Methods

### NewRuntimeStatusIpsecValue

`func NewRuntimeStatusIpsecValue() *RuntimeStatusIpsecValue`

NewRuntimeStatusIpsecValue instantiates a new RuntimeStatusIpsecValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeStatusIpsecValueWithDefaults

`func NewRuntimeStatusIpsecValueWithDefaults() *RuntimeStatusIpsecValue`

NewRuntimeStatusIpsecValueWithDefaults instantiates a new RuntimeStatusIpsecValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuntimeStatusIpsecValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuntimeStatusIpsecValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuntimeStatusIpsecValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RuntimeStatusIpsecValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalSubnet

`func (o *RuntimeStatusIpsecValue) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *RuntimeStatusIpsecValue) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *RuntimeStatusIpsecValue) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *RuntimeStatusIpsecValue) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *RuntimeStatusIpsecValue) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *RuntimeStatusIpsecValue) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *RuntimeStatusIpsecValue) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *RuntimeStatusIpsecValue) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetEndpointid

`func (o *RuntimeStatusIpsecValue) GetEndpointid() int32`

GetEndpointid returns the Endpointid field if non-nil, zero value otherwise.

### GetEndpointidOk

`func (o *RuntimeStatusIpsecValue) GetEndpointidOk() (*int32, bool)`

GetEndpointidOk returns a tuple with the Endpointid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointid

`func (o *RuntimeStatusIpsecValue) SetEndpointid(v int32)`

SetEndpointid sets Endpointid field to given value.

### HasEndpointid

`func (o *RuntimeStatusIpsecValue) HasEndpointid() bool`

HasEndpointid returns a boolean if a field has been set.

### GetEndpointId

`func (o *RuntimeStatusIpsecValue) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *RuntimeStatusIpsecValue) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *RuntimeStatusIpsecValue) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *RuntimeStatusIpsecValue) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointName

`func (o *RuntimeStatusIpsecValue) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *RuntimeStatusIpsecValue) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *RuntimeStatusIpsecValue) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *RuntimeStatusIpsecValue) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetEnabled

`func (o *RuntimeStatusIpsecValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuntimeStatusIpsecValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuntimeStatusIpsecValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuntimeStatusIpsecValue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActive

`func (o *RuntimeStatusIpsecValue) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuntimeStatusIpsecValue) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuntimeStatusIpsecValue) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuntimeStatusIpsecValue) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *RuntimeStatusIpsecValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuntimeStatusIpsecValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuntimeStatusIpsecValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuntimeStatusIpsecValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RuntimeStatusIpsecValue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RuntimeStatusIpsecValue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBounce

`func (o *RuntimeStatusIpsecValue) GetBounce() bool`

GetBounce returns the Bounce field if non-nil, zero value otherwise.

### GetBounceOk

`func (o *RuntimeStatusIpsecValue) GetBounceOk() (*bool, bool)`

GetBounceOk returns a tuple with the Bounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounce

`func (o *RuntimeStatusIpsecValue) SetBounce(v bool)`

SetBounce sets Bounce field to given value.

### HasBounce

`func (o *RuntimeStatusIpsecValue) HasBounce() bool`

HasBounce returns a boolean if a field has been set.

### GetConnected

`func (o *RuntimeStatusIpsecValue) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *RuntimeStatusIpsecValue) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *RuntimeStatusIpsecValue) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *RuntimeStatusIpsecValue) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetPingInterface

`func (o *RuntimeStatusIpsecValue) GetPingInterface() string`

GetPingInterface returns the PingInterface field if non-nil, zero value otherwise.

### GetPingInterfaceOk

`func (o *RuntimeStatusIpsecValue) GetPingInterfaceOk() (*string, bool)`

GetPingInterfaceOk returns a tuple with the PingInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterface

`func (o *RuntimeStatusIpsecValue) SetPingInterface(v string)`

SetPingInterface sets PingInterface field to given value.

### HasPingInterface

`func (o *RuntimeStatusIpsecValue) HasPingInterface() bool`

HasPingInterface returns a boolean if a field has been set.

### GetPingInterval

`func (o *RuntimeStatusIpsecValue) GetPingInterval() int32`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *RuntimeStatusIpsecValue) GetPingIntervalOk() (*int32, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *RuntimeStatusIpsecValue) SetPingInterval(v int32)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *RuntimeStatusIpsecValue) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### SetPingIntervalNil

`func (o *RuntimeStatusIpsecValue) SetPingIntervalNil(b bool)`

 SetPingIntervalNil sets the value for PingInterval to be an explicit nil

### UnsetPingInterval
`func (o *RuntimeStatusIpsecValue) UnsetPingInterval()`

UnsetPingInterval ensures that no value is present for PingInterval, not even an explicit nil
### GetPingIpaddress

`func (o *RuntimeStatusIpsecValue) GetPingIpaddress() string`

GetPingIpaddress returns the PingIpaddress field if non-nil, zero value otherwise.

### GetPingIpaddressOk

`func (o *RuntimeStatusIpsecValue) GetPingIpaddressOk() (*string, bool)`

GetPingIpaddressOk returns a tuple with the PingIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIpaddress

`func (o *RuntimeStatusIpsecValue) SetPingIpaddress(v string)`

SetPingIpaddress sets PingIpaddress field to given value.

### HasPingIpaddress

`func (o *RuntimeStatusIpsecValue) HasPingIpaddress() bool`

HasPingIpaddress returns a boolean if a field has been set.

### GetTunnelParams

`func (o *RuntimeStatusIpsecValue) GetTunnelParams() IpsecTunnelParams`

GetTunnelParams returns the TunnelParams field if non-nil, zero value otherwise.

### GetTunnelParamsOk

`func (o *RuntimeStatusIpsecValue) GetTunnelParamsOk() (*IpsecTunnelParams, bool)`

GetTunnelParamsOk returns a tuple with the TunnelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelParams

`func (o *RuntimeStatusIpsecValue) SetTunnelParams(v IpsecTunnelParams)`

SetTunnelParams sets TunnelParams field to given value.

### HasTunnelParams

`func (o *RuntimeStatusIpsecValue) HasTunnelParams() bool`

HasTunnelParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IpsecSystemDetailResponseRemoteEndpointsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NatTEnabled** | Pointer to **bool** |  | [optional] 
**IkeVersion** | Pointer to **int32** |  | [optional] 
**Pfs** | Pointer to **bool** | Perfect forward secrecy enabled | [optional] 
**PrivateIpaddress** | Pointer to **string** |  | [optional] 
**ExtraConfig** | Pointer to **[]string** |  | [optional] 
**Tunnels** | Pointer to [**map[string]RuntimeStatusIpsecValue**](RuntimeStatusIpsecValue.md) |  | [optional] 
**TrafficPairs** | Pointer to [**map[string]IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue**](IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue.md) |  | [optional] 
**BgpPeers** | Pointer to [**map[string]IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue**](IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue.md) |  | [optional] 
**Type** | Pointer to **string** | Indicating Ipsec or GRE over ipsec | [optional] 
**VpnType** | Pointer to **string** |  | [optional] 
**GreInterfaceAddress** | Pointer to **string** |  | [optional] 
**RouteBasedIntAddress** | Pointer to **string** |  | [optional] 
**RouteBasedLocal** | Pointer to **string** |  | [optional] 
**RouteBasedRemote** | Pointer to **string** |  | [optional] 
**Psk** | Pointer to **string** |  | [optional] 

## Methods

### NewIpsecSystemDetailResponseRemoteEndpointsValue

`func NewIpsecSystemDetailResponseRemoteEndpointsValue() *IpsecSystemDetailResponseRemoteEndpointsValue`

NewIpsecSystemDetailResponseRemoteEndpointsValue instantiates a new IpsecSystemDetailResponseRemoteEndpointsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecSystemDetailResponseRemoteEndpointsValueWithDefaults

`func NewIpsecSystemDetailResponseRemoteEndpointsValueWithDefaults() *IpsecSystemDetailResponseRemoteEndpointsValue`

NewIpsecSystemDetailResponseRemoteEndpointsValueWithDefaults instantiates a new IpsecSystemDetailResponseRemoteEndpointsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNatTEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetNatTEnabled() bool`

GetNatTEnabled returns the NatTEnabled field if non-nil, zero value otherwise.

### GetNatTEnabledOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetNatTEnabledOk() (*bool, bool)`

GetNatTEnabledOk returns a tuple with the NatTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetNatTEnabled(v bool)`

SetNatTEnabled sets NatTEnabled field to given value.

### HasNatTEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasNatTEnabled() bool`

HasNatTEnabled returns a boolean if a field has been set.

### GetIkeVersion

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetPfs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPfs() bool`

GetPfs returns the Pfs field if non-nil, zero value otherwise.

### GetPfsOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPfsOk() (*bool, bool)`

GetPfsOk returns a tuple with the Pfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetPfs(v bool)`

SetPfs sets Pfs field to given value.

### HasPfs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasPfs() bool`

HasPfs returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetExtraConfig

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetExtraConfig() []string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetExtraConfigOk() (*[]string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetExtraConfig(v []string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### GetTunnels

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetTunnels() map[string]RuntimeStatusIpsecValue`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetTunnelsOk() (*map[string]RuntimeStatusIpsecValue, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetTunnels(v map[string]RuntimeStatusIpsecValue)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetTrafficPairs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetTrafficPairs() map[string]IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue`

GetTrafficPairs returns the TrafficPairs field if non-nil, zero value otherwise.

### GetTrafficPairsOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetTrafficPairsOk() (*map[string]IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue, bool)`

GetTrafficPairsOk returns a tuple with the TrafficPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPairs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetTrafficPairs(v map[string]IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue)`

SetTrafficPairs sets TrafficPairs field to given value.

### HasTrafficPairs

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasTrafficPairs() bool`

HasTrafficPairs returns a boolean if a field has been set.

### GetBgpPeers

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetBgpPeers() map[string]IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue`

GetBgpPeers returns the BgpPeers field if non-nil, zero value otherwise.

### GetBgpPeersOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetBgpPeersOk() (*map[string]IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue, bool)`

GetBgpPeersOk returns a tuple with the BgpPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpPeers

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetBgpPeers(v map[string]IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue)`

SetBgpPeers sets BgpPeers field to given value.

### HasBgpPeers

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasBgpPeers() bool`

HasBgpPeers returns a boolean if a field has been set.

### GetType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVpnType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetVpnType() string`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetVpnTypeOk() (*string, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetVpnType(v string)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetGreInterfaceAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetGreInterfaceAddress() string`

GetGreInterfaceAddress returns the GreInterfaceAddress field if non-nil, zero value otherwise.

### GetGreInterfaceAddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetGreInterfaceAddressOk() (*string, bool)`

GetGreInterfaceAddressOk returns a tuple with the GreInterfaceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreInterfaceAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetGreInterfaceAddress(v string)`

SetGreInterfaceAddress sets GreInterfaceAddress field to given value.

### HasGreInterfaceAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasGreInterfaceAddress() bool`

HasGreInterfaceAddress returns a boolean if a field has been set.

### GetRouteBasedIntAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedIntAddress() string`

GetRouteBasedIntAddress returns the RouteBasedIntAddress field if non-nil, zero value otherwise.

### GetRouteBasedIntAddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedIntAddressOk() (*string, bool)`

GetRouteBasedIntAddressOk returns a tuple with the RouteBasedIntAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedIntAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetRouteBasedIntAddress(v string)`

SetRouteBasedIntAddress sets RouteBasedIntAddress field to given value.

### HasRouteBasedIntAddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasRouteBasedIntAddress() bool`

HasRouteBasedIntAddress returns a boolean if a field has been set.

### GetRouteBasedLocal

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedLocal() string`

GetRouteBasedLocal returns the RouteBasedLocal field if non-nil, zero value otherwise.

### GetRouteBasedLocalOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedLocalOk() (*string, bool)`

GetRouteBasedLocalOk returns a tuple with the RouteBasedLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedLocal

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetRouteBasedLocal(v string)`

SetRouteBasedLocal sets RouteBasedLocal field to given value.

### HasRouteBasedLocal

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasRouteBasedLocal() bool`

HasRouteBasedLocal returns a boolean if a field has been set.

### GetRouteBasedRemote

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedRemote() string`

GetRouteBasedRemote returns the RouteBasedRemote field if non-nil, zero value otherwise.

### GetRouteBasedRemoteOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetRouteBasedRemoteOk() (*string, bool)`

GetRouteBasedRemoteOk returns a tuple with the RouteBasedRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedRemote

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetRouteBasedRemote(v string)`

SetRouteBasedRemote sets RouteBasedRemote field to given value.

### HasRouteBasedRemote

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasRouteBasedRemote() bool`

HasRouteBasedRemote returns a boolean if a field has been set.

### GetPsk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValue) HasPsk() bool`

HasPsk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



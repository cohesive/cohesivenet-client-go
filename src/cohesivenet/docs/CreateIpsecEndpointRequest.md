# CreateIpsecEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for the connection. | 
**Description** | Pointer to **string** | Description of this IPsec endpoint | [optional] 
**Ipaddress** | **string** | IP of the remote gateway | 
**Secret** | **string** | Pre-shared key | 
**Pfs** | Pointer to **bool** | Perfect Forward Secrecy if true, disables if false. | [optional] [default to true]
**IkeVersion** | Pointer to **int32** | Version for IKE algorithm | [optional] [default to 1]
**NatTEnabled** | Pointer to **bool** | True if you want encapsulated IPsec protocol to this gateway | [optional] [default to true]
**ExtraConfig** | Pointer to **string** | Additional optionals for connection such as &#39;phase1&#x3D;aes256_gcm-sha2_256-dh14 phase2&#x3D;aes256_gcm&#39; | [optional] 
**PrivateIpaddress** | Pointer to **string** | Internal NAT address of the remote gateway | [optional] 
**Gre** | Pointer to **bool** | True if GRE is being used for the specific endpoint | [optional] 
**GreInterfaceAddress** | Pointer to **string** | Interface for GRE in /30 format | [optional] 
**VpnType** | Pointer to **string** | policy, gre, vti | [optional] [default to "policy"]
**RouteBasedIntAddress** | Pointer to **string** |  | [optional] 
**RouteBasedLocal** | Pointer to **string** |  | [optional] 
**RouteBasedRemote** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateIpsecEndpointRequest

`func NewCreateIpsecEndpointRequest(name string, ipaddress string, secret string, ) *CreateIpsecEndpointRequest`

NewCreateIpsecEndpointRequest instantiates a new CreateIpsecEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpsecEndpointRequestWithDefaults

`func NewCreateIpsecEndpointRequestWithDefaults() *CreateIpsecEndpointRequest`

NewCreateIpsecEndpointRequestWithDefaults instantiates a new CreateIpsecEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIpsecEndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIpsecEndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIpsecEndpointRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateIpsecEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIpsecEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIpsecEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIpsecEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpaddress

`func (o *CreateIpsecEndpointRequest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *CreateIpsecEndpointRequest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *CreateIpsecEndpointRequest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.


### GetSecret

`func (o *CreateIpsecEndpointRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateIpsecEndpointRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateIpsecEndpointRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetPfs

`func (o *CreateIpsecEndpointRequest) GetPfs() bool`

GetPfs returns the Pfs field if non-nil, zero value otherwise.

### GetPfsOk

`func (o *CreateIpsecEndpointRequest) GetPfsOk() (*bool, bool)`

GetPfsOk returns a tuple with the Pfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfs

`func (o *CreateIpsecEndpointRequest) SetPfs(v bool)`

SetPfs sets Pfs field to given value.

### HasPfs

`func (o *CreateIpsecEndpointRequest) HasPfs() bool`

HasPfs returns a boolean if a field has been set.

### GetIkeVersion

`func (o *CreateIpsecEndpointRequest) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *CreateIpsecEndpointRequest) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *CreateIpsecEndpointRequest) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *CreateIpsecEndpointRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNatTEnabled

`func (o *CreateIpsecEndpointRequest) GetNatTEnabled() bool`

GetNatTEnabled returns the NatTEnabled field if non-nil, zero value otherwise.

### GetNatTEnabledOk

`func (o *CreateIpsecEndpointRequest) GetNatTEnabledOk() (*bool, bool)`

GetNatTEnabledOk returns a tuple with the NatTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTEnabled

`func (o *CreateIpsecEndpointRequest) SetNatTEnabled(v bool)`

SetNatTEnabled sets NatTEnabled field to given value.

### HasNatTEnabled

`func (o *CreateIpsecEndpointRequest) HasNatTEnabled() bool`

HasNatTEnabled returns a boolean if a field has been set.

### GetExtraConfig

`func (o *CreateIpsecEndpointRequest) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *CreateIpsecEndpointRequest) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *CreateIpsecEndpointRequest) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *CreateIpsecEndpointRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *CreateIpsecEndpointRequest) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *CreateIpsecEndpointRequest) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *CreateIpsecEndpointRequest) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *CreateIpsecEndpointRequest) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetGre

`func (o *CreateIpsecEndpointRequest) GetGre() bool`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *CreateIpsecEndpointRequest) GetGreOk() (*bool, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *CreateIpsecEndpointRequest) SetGre(v bool)`

SetGre sets Gre field to given value.

### HasGre

`func (o *CreateIpsecEndpointRequest) HasGre() bool`

HasGre returns a boolean if a field has been set.

### GetGreInterfaceAddress

`func (o *CreateIpsecEndpointRequest) GetGreInterfaceAddress() string`

GetGreInterfaceAddress returns the GreInterfaceAddress field if non-nil, zero value otherwise.

### GetGreInterfaceAddressOk

`func (o *CreateIpsecEndpointRequest) GetGreInterfaceAddressOk() (*string, bool)`

GetGreInterfaceAddressOk returns a tuple with the GreInterfaceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreInterfaceAddress

`func (o *CreateIpsecEndpointRequest) SetGreInterfaceAddress(v string)`

SetGreInterfaceAddress sets GreInterfaceAddress field to given value.

### HasGreInterfaceAddress

`func (o *CreateIpsecEndpointRequest) HasGreInterfaceAddress() bool`

HasGreInterfaceAddress returns a boolean if a field has been set.

### GetVpnType

`func (o *CreateIpsecEndpointRequest) GetVpnType() string`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *CreateIpsecEndpointRequest) GetVpnTypeOk() (*string, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *CreateIpsecEndpointRequest) SetVpnType(v string)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *CreateIpsecEndpointRequest) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetRouteBasedIntAddress

`func (o *CreateIpsecEndpointRequest) GetRouteBasedIntAddress() string`

GetRouteBasedIntAddress returns the RouteBasedIntAddress field if non-nil, zero value otherwise.

### GetRouteBasedIntAddressOk

`func (o *CreateIpsecEndpointRequest) GetRouteBasedIntAddressOk() (*string, bool)`

GetRouteBasedIntAddressOk returns a tuple with the RouteBasedIntAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedIntAddress

`func (o *CreateIpsecEndpointRequest) SetRouteBasedIntAddress(v string)`

SetRouteBasedIntAddress sets RouteBasedIntAddress field to given value.

### HasRouteBasedIntAddress

`func (o *CreateIpsecEndpointRequest) HasRouteBasedIntAddress() bool`

HasRouteBasedIntAddress returns a boolean if a field has been set.

### GetRouteBasedLocal

`func (o *CreateIpsecEndpointRequest) GetRouteBasedLocal() string`

GetRouteBasedLocal returns the RouteBasedLocal field if non-nil, zero value otherwise.

### GetRouteBasedLocalOk

`func (o *CreateIpsecEndpointRequest) GetRouteBasedLocalOk() (*string, bool)`

GetRouteBasedLocalOk returns a tuple with the RouteBasedLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedLocal

`func (o *CreateIpsecEndpointRequest) SetRouteBasedLocal(v string)`

SetRouteBasedLocal sets RouteBasedLocal field to given value.

### HasRouteBasedLocal

`func (o *CreateIpsecEndpointRequest) HasRouteBasedLocal() bool`

HasRouteBasedLocal returns a boolean if a field has been set.

### GetRouteBasedRemote

`func (o *CreateIpsecEndpointRequest) GetRouteBasedRemote() string`

GetRouteBasedRemote returns the RouteBasedRemote field if non-nil, zero value otherwise.

### GetRouteBasedRemoteOk

`func (o *CreateIpsecEndpointRequest) GetRouteBasedRemoteOk() (*string, bool)`

GetRouteBasedRemoteOk returns a tuple with the RouteBasedRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedRemote

`func (o *CreateIpsecEndpointRequest) SetRouteBasedRemote(v string)`

SetRouteBasedRemote sets RouteBasedRemote field to given value.

### HasRouteBasedRemote

`func (o *CreateIpsecEndpointRequest) HasRouteBasedRemote() bool`

HasRouteBasedRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



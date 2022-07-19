# UpdateIpsecEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the endpoint. | [optional] 
**Description** | Pointer to **string** | Description of this IPsec endpoint | [optional] 
**Ipaddress** | Pointer to **string** | IP of the remote gateway | [optional] 
**Secret** | Pointer to **string** | Pre-shared key | [optional] 
**Pfs** | Pointer to **bool** | Perfect Forward Secrecy if true, disables if false. | [optional] 
**IkeVersion** | Pointer to **int32** | Version for IKE algorithm | [optional] 
**NatTEnabled** | Pointer to **bool** | True if you want encapsulated IPsec protocol to this gateway | [optional] 
**ExtraConfig** | Pointer to **string** | Additional optionals for connection such as &#39;phase1&#x3D;aes256_gcm-sha2_256-dh14 phase2&#x3D;aes256_gcm&#39; | [optional] 
**PrivateIpaddress** | Pointer to **string** | Internal NAT address of the remote gateway | [optional] 
**Gre** | Pointer to **bool** | True if GRE is being used for the specific endpoint | [optional] 
**GreInterfaceAddress** | Pointer to **string** | Interface address for GRE | [optional] 
**VpnType** | Pointer to **string** | policy, gre, vti | [optional] 
**RouteBasedIntAddress** | Pointer to **string** |  | [optional] 
**RouteBasedLocal** | Pointer to **string** |  | [optional] 
**RouteBasedRemote** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateIpsecEndpointRequest

`func NewUpdateIpsecEndpointRequest() *UpdateIpsecEndpointRequest`

NewUpdateIpsecEndpointRequest instantiates a new UpdateIpsecEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpsecEndpointRequestWithDefaults

`func NewUpdateIpsecEndpointRequestWithDefaults() *UpdateIpsecEndpointRequest`

NewUpdateIpsecEndpointRequestWithDefaults instantiates a new UpdateIpsecEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIpsecEndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIpsecEndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIpsecEndpointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIpsecEndpointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIpsecEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIpsecEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIpsecEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIpsecEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpaddress

`func (o *UpdateIpsecEndpointRequest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *UpdateIpsecEndpointRequest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *UpdateIpsecEndpointRequest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *UpdateIpsecEndpointRequest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateIpsecEndpointRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateIpsecEndpointRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateIpsecEndpointRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateIpsecEndpointRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPfs

`func (o *UpdateIpsecEndpointRequest) GetPfs() bool`

GetPfs returns the Pfs field if non-nil, zero value otherwise.

### GetPfsOk

`func (o *UpdateIpsecEndpointRequest) GetPfsOk() (*bool, bool)`

GetPfsOk returns a tuple with the Pfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfs

`func (o *UpdateIpsecEndpointRequest) SetPfs(v bool)`

SetPfs sets Pfs field to given value.

### HasPfs

`func (o *UpdateIpsecEndpointRequest) HasPfs() bool`

HasPfs returns a boolean if a field has been set.

### GetIkeVersion

`func (o *UpdateIpsecEndpointRequest) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *UpdateIpsecEndpointRequest) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *UpdateIpsecEndpointRequest) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *UpdateIpsecEndpointRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNatTEnabled

`func (o *UpdateIpsecEndpointRequest) GetNatTEnabled() bool`

GetNatTEnabled returns the NatTEnabled field if non-nil, zero value otherwise.

### GetNatTEnabledOk

`func (o *UpdateIpsecEndpointRequest) GetNatTEnabledOk() (*bool, bool)`

GetNatTEnabledOk returns a tuple with the NatTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTEnabled

`func (o *UpdateIpsecEndpointRequest) SetNatTEnabled(v bool)`

SetNatTEnabled sets NatTEnabled field to given value.

### HasNatTEnabled

`func (o *UpdateIpsecEndpointRequest) HasNatTEnabled() bool`

HasNatTEnabled returns a boolean if a field has been set.

### GetExtraConfig

`func (o *UpdateIpsecEndpointRequest) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *UpdateIpsecEndpointRequest) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *UpdateIpsecEndpointRequest) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *UpdateIpsecEndpointRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *UpdateIpsecEndpointRequest) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *UpdateIpsecEndpointRequest) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *UpdateIpsecEndpointRequest) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *UpdateIpsecEndpointRequest) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetGre

`func (o *UpdateIpsecEndpointRequest) GetGre() bool`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *UpdateIpsecEndpointRequest) GetGreOk() (*bool, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *UpdateIpsecEndpointRequest) SetGre(v bool)`

SetGre sets Gre field to given value.

### HasGre

`func (o *UpdateIpsecEndpointRequest) HasGre() bool`

HasGre returns a boolean if a field has been set.

### GetGreInterfaceAddress

`func (o *UpdateIpsecEndpointRequest) GetGreInterfaceAddress() string`

GetGreInterfaceAddress returns the GreInterfaceAddress field if non-nil, zero value otherwise.

### GetGreInterfaceAddressOk

`func (o *UpdateIpsecEndpointRequest) GetGreInterfaceAddressOk() (*string, bool)`

GetGreInterfaceAddressOk returns a tuple with the GreInterfaceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreInterfaceAddress

`func (o *UpdateIpsecEndpointRequest) SetGreInterfaceAddress(v string)`

SetGreInterfaceAddress sets GreInterfaceAddress field to given value.

### HasGreInterfaceAddress

`func (o *UpdateIpsecEndpointRequest) HasGreInterfaceAddress() bool`

HasGreInterfaceAddress returns a boolean if a field has been set.

### GetVpnType

`func (o *UpdateIpsecEndpointRequest) GetVpnType() string`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *UpdateIpsecEndpointRequest) GetVpnTypeOk() (*string, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *UpdateIpsecEndpointRequest) SetVpnType(v string)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *UpdateIpsecEndpointRequest) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetRouteBasedIntAddress

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedIntAddress() string`

GetRouteBasedIntAddress returns the RouteBasedIntAddress field if non-nil, zero value otherwise.

### GetRouteBasedIntAddressOk

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedIntAddressOk() (*string, bool)`

GetRouteBasedIntAddressOk returns a tuple with the RouteBasedIntAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedIntAddress

`func (o *UpdateIpsecEndpointRequest) SetRouteBasedIntAddress(v string)`

SetRouteBasedIntAddress sets RouteBasedIntAddress field to given value.

### HasRouteBasedIntAddress

`func (o *UpdateIpsecEndpointRequest) HasRouteBasedIntAddress() bool`

HasRouteBasedIntAddress returns a boolean if a field has been set.

### GetRouteBasedLocal

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedLocal() string`

GetRouteBasedLocal returns the RouteBasedLocal field if non-nil, zero value otherwise.

### GetRouteBasedLocalOk

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedLocalOk() (*string, bool)`

GetRouteBasedLocalOk returns a tuple with the RouteBasedLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedLocal

`func (o *UpdateIpsecEndpointRequest) SetRouteBasedLocal(v string)`

SetRouteBasedLocal sets RouteBasedLocal field to given value.

### HasRouteBasedLocal

`func (o *UpdateIpsecEndpointRequest) HasRouteBasedLocal() bool`

HasRouteBasedLocal returns a boolean if a field has been set.

### GetRouteBasedRemote

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedRemote() string`

GetRouteBasedRemote returns the RouteBasedRemote field if non-nil, zero value otherwise.

### GetRouteBasedRemoteOk

`func (o *UpdateIpsecEndpointRequest) GetRouteBasedRemoteOk() (*string, bool)`

GetRouteBasedRemoteOk returns a tuple with the RouteBasedRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedRemote

`func (o *UpdateIpsecEndpointRequest) SetRouteBasedRemote(v string)`

SetRouteBasedRemote sets RouteBasedRemote field to given value.

### HasRouteBasedRemote

`func (o *UpdateIpsecEndpointRequest) HasRouteBasedRemote() bool`

HasRouteBasedRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



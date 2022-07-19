# RuntimeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** | Autonomous system number for controller if peered | [optional] 
**TopologyName** | Pointer to **string** |  | [optional] 
**ControllerName** | Pointer to **string** |  | [optional] 
**TopologyChecksum** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **int32** | This managers ID in peered topology | [optional] 
**NtpHosts** | Pointer to **string** | NTP host endpoints, whitespace delimited | [optional] 
**Vns3Version** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **bool** |  | [optional] 
**OverlayIpaddress** | Pointer to **string** | This managers overlay IP in peered topology | [optional] 
**Peered** | Pointer to **bool** |  | [optional] 
**PublicIpaddress** | Pointer to **string** |  | [optional] 
**SubnetGateway** | Pointer to **string** |  | [optional] 
**PrivateIpaddress** | Pointer to **string** |  | [optional] 
**SecurityToken** | Pointer to **string** | Security token in peered topology | [optional] 

## Methods

### NewRuntimeConfig

`func NewRuntimeConfig() *RuntimeConfig`

NewRuntimeConfig instantiates a new RuntimeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeConfigWithDefaults

`func NewRuntimeConfigWithDefaults() *RuntimeConfig`

NewRuntimeConfigWithDefaults instantiates a new RuntimeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *RuntimeConfig) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *RuntimeConfig) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *RuntimeConfig) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *RuntimeConfig) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetTopologyName

`func (o *RuntimeConfig) GetTopologyName() string`

GetTopologyName returns the TopologyName field if non-nil, zero value otherwise.

### GetTopologyNameOk

`func (o *RuntimeConfig) GetTopologyNameOk() (*string, bool)`

GetTopologyNameOk returns a tuple with the TopologyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyName

`func (o *RuntimeConfig) SetTopologyName(v string)`

SetTopologyName sets TopologyName field to given value.

### HasTopologyName

`func (o *RuntimeConfig) HasTopologyName() bool`

HasTopologyName returns a boolean if a field has been set.

### GetControllerName

`func (o *RuntimeConfig) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *RuntimeConfig) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *RuntimeConfig) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *RuntimeConfig) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetTopologyChecksum

`func (o *RuntimeConfig) GetTopologyChecksum() string`

GetTopologyChecksum returns the TopologyChecksum field if non-nil, zero value otherwise.

### GetTopologyChecksumOk

`func (o *RuntimeConfig) GetTopologyChecksumOk() (*string, bool)`

GetTopologyChecksumOk returns a tuple with the TopologyChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyChecksum

`func (o *RuntimeConfig) SetTopologyChecksum(v string)`

SetTopologyChecksum sets TopologyChecksum field to given value.

### HasTopologyChecksum

`func (o *RuntimeConfig) HasTopologyChecksum() bool`

HasTopologyChecksum returns a boolean if a field has been set.

### GetManagerId

`func (o *RuntimeConfig) GetManagerId() int32`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *RuntimeConfig) GetManagerIdOk() (*int32, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *RuntimeConfig) SetManagerId(v int32)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *RuntimeConfig) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetNtpHosts

`func (o *RuntimeConfig) GetNtpHosts() string`

GetNtpHosts returns the NtpHosts field if non-nil, zero value otherwise.

### GetNtpHostsOk

`func (o *RuntimeConfig) GetNtpHostsOk() (*string, bool)`

GetNtpHostsOk returns a tuple with the NtpHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpHosts

`func (o *RuntimeConfig) SetNtpHosts(v string)`

SetNtpHosts sets NtpHosts field to given value.

### HasNtpHosts

`func (o *RuntimeConfig) HasNtpHosts() bool`

HasNtpHosts returns a boolean if a field has been set.

### GetVns3Version

`func (o *RuntimeConfig) GetVns3Version() string`

GetVns3Version returns the Vns3Version field if non-nil, zero value otherwise.

### GetVns3VersionOk

`func (o *RuntimeConfig) GetVns3VersionOk() (*string, bool)`

GetVns3VersionOk returns a tuple with the Vns3Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVns3Version

`func (o *RuntimeConfig) SetVns3Version(v string)`

SetVns3Version sets Vns3Version field to given value.

### HasVns3Version

`func (o *RuntimeConfig) HasVns3Version() bool`

HasVns3Version returns a boolean if a field has been set.

### GetLicensed

`func (o *RuntimeConfig) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *RuntimeConfig) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *RuntimeConfig) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *RuntimeConfig) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetOverlayIpaddress

`func (o *RuntimeConfig) GetOverlayIpaddress() string`

GetOverlayIpaddress returns the OverlayIpaddress field if non-nil, zero value otherwise.

### GetOverlayIpaddressOk

`func (o *RuntimeConfig) GetOverlayIpaddressOk() (*string, bool)`

GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayIpaddress

`func (o *RuntimeConfig) SetOverlayIpaddress(v string)`

SetOverlayIpaddress sets OverlayIpaddress field to given value.

### HasOverlayIpaddress

`func (o *RuntimeConfig) HasOverlayIpaddress() bool`

HasOverlayIpaddress returns a boolean if a field has been set.

### GetPeered

`func (o *RuntimeConfig) GetPeered() bool`

GetPeered returns the Peered field if non-nil, zero value otherwise.

### GetPeeredOk

`func (o *RuntimeConfig) GetPeeredOk() (*bool, bool)`

GetPeeredOk returns a tuple with the Peered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeered

`func (o *RuntimeConfig) SetPeered(v bool)`

SetPeered sets Peered field to given value.

### HasPeered

`func (o *RuntimeConfig) HasPeered() bool`

HasPeered returns a boolean if a field has been set.

### GetPublicIpaddress

`func (o *RuntimeConfig) GetPublicIpaddress() string`

GetPublicIpaddress returns the PublicIpaddress field if non-nil, zero value otherwise.

### GetPublicIpaddressOk

`func (o *RuntimeConfig) GetPublicIpaddressOk() (*string, bool)`

GetPublicIpaddressOk returns a tuple with the PublicIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpaddress

`func (o *RuntimeConfig) SetPublicIpaddress(v string)`

SetPublicIpaddress sets PublicIpaddress field to given value.

### HasPublicIpaddress

`func (o *RuntimeConfig) HasPublicIpaddress() bool`

HasPublicIpaddress returns a boolean if a field has been set.

### GetSubnetGateway

`func (o *RuntimeConfig) GetSubnetGateway() string`

GetSubnetGateway returns the SubnetGateway field if non-nil, zero value otherwise.

### GetSubnetGatewayOk

`func (o *RuntimeConfig) GetSubnetGatewayOk() (*string, bool)`

GetSubnetGatewayOk returns a tuple with the SubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGateway

`func (o *RuntimeConfig) SetSubnetGateway(v string)`

SetSubnetGateway sets SubnetGateway field to given value.

### HasSubnetGateway

`func (o *RuntimeConfig) HasSubnetGateway() bool`

HasSubnetGateway returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *RuntimeConfig) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *RuntimeConfig) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *RuntimeConfig) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *RuntimeConfig) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetSecurityToken

`func (o *RuntimeConfig) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *RuntimeConfig) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *RuntimeConfig) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *RuntimeConfig) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



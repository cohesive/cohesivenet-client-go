# IpsecLocalEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatTraversal** | Pointer to **bool** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**OverlaySubnet** | Pointer to **string** |  | [optional] 
**PrivateIpaddress** | Pointer to **string** |  | [optional] 
**IpsecLocalIpaddress** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to **int32** |  | [optional] 

## Methods

### NewIpsecLocalEndpoint

`func NewIpsecLocalEndpoint() *IpsecLocalEndpoint`

NewIpsecLocalEndpoint instantiates a new IpsecLocalEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecLocalEndpointWithDefaults

`func NewIpsecLocalEndpointWithDefaults() *IpsecLocalEndpoint`

NewIpsecLocalEndpointWithDefaults instantiates a new IpsecLocalEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatTraversal

`func (o *IpsecLocalEndpoint) GetNatTraversal() bool`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *IpsecLocalEndpoint) GetNatTraversalOk() (*bool, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *IpsecLocalEndpoint) SetNatTraversal(v bool)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *IpsecLocalEndpoint) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetIpaddress

`func (o *IpsecLocalEndpoint) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *IpsecLocalEndpoint) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *IpsecLocalEndpoint) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *IpsecLocalEndpoint) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetOverlaySubnet

`func (o *IpsecLocalEndpoint) GetOverlaySubnet() string`

GetOverlaySubnet returns the OverlaySubnet field if non-nil, zero value otherwise.

### GetOverlaySubnetOk

`func (o *IpsecLocalEndpoint) GetOverlaySubnetOk() (*string, bool)`

GetOverlaySubnetOk returns a tuple with the OverlaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlaySubnet

`func (o *IpsecLocalEndpoint) SetOverlaySubnet(v string)`

SetOverlaySubnet sets OverlaySubnet field to given value.

### HasOverlaySubnet

`func (o *IpsecLocalEndpoint) HasOverlaySubnet() bool`

HasOverlaySubnet returns a boolean if a field has been set.

### GetPrivateIpaddress

`func (o *IpsecLocalEndpoint) GetPrivateIpaddress() string`

GetPrivateIpaddress returns the PrivateIpaddress field if non-nil, zero value otherwise.

### GetPrivateIpaddressOk

`func (o *IpsecLocalEndpoint) GetPrivateIpaddressOk() (*string, bool)`

GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpaddress

`func (o *IpsecLocalEndpoint) SetPrivateIpaddress(v string)`

SetPrivateIpaddress sets PrivateIpaddress field to given value.

### HasPrivateIpaddress

`func (o *IpsecLocalEndpoint) HasPrivateIpaddress() bool`

HasPrivateIpaddress returns a boolean if a field has been set.

### GetIpsecLocalIpaddress

`func (o *IpsecLocalEndpoint) GetIpsecLocalIpaddress() string`

GetIpsecLocalIpaddress returns the IpsecLocalIpaddress field if non-nil, zero value otherwise.

### GetIpsecLocalIpaddressOk

`func (o *IpsecLocalEndpoint) GetIpsecLocalIpaddressOk() (*string, bool)`

GetIpsecLocalIpaddressOk returns a tuple with the IpsecLocalIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLocalIpaddress

`func (o *IpsecLocalEndpoint) SetIpsecLocalIpaddress(v string)`

SetIpsecLocalIpaddress sets IpsecLocalIpaddress field to given value.

### HasIpsecLocalIpaddress

`func (o *IpsecLocalEndpoint) HasIpsecLocalIpaddress() bool`

HasIpsecLocalIpaddress returns a boolean if a field has been set.

### GetAsn

`func (o *IpsecLocalEndpoint) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IpsecLocalEndpoint) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IpsecLocalEndpoint) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IpsecLocalEndpoint) HasAsn() bool`

HasAsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



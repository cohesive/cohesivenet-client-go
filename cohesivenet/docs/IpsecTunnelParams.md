# IpsecTunnelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase2** | Pointer to **string** |  | [optional] 
**OutboundSpi** | Pointer to **string** |  | [optional] 
**InboundSpi** | Pointer to **string** |  | [optional] 
**BytesIn** | Pointer to **string** |  | [optional] 
**BytesOut** | Pointer to **string** |  | [optional] 
**EspTimeRemaining** | Pointer to **string** |  | [optional] 
**EspPort** | Pointer to **string** |  | [optional] 
**Phase2Algo** | Pointer to **string** |  | [optional] 
**Phase2Hash** | Pointer to **string** |  | [optional] 
**NatT** | Pointer to **string** |  | [optional] 
**Dpd** | Pointer to **string** |  | [optional] 
**PfsDhGroup** | Pointer to **NullableString** |  | [optional] 
**Phase1** | Pointer to **string** |  | [optional] 
**IsakmpPort** | Pointer to **string** |  | [optional] 
**IsakmpTimeRemaining** | Pointer to **string** |  | [optional] 
**LastDpd** | Pointer to **string** |  | [optional] 
**Phase1Cipher** | Pointer to **NullableString** |  | [optional] 
**Phase1Prf** | Pointer to **NullableString** |  | [optional] 
**Phase1DhGroup** | Pointer to **NullableString** |  | [optional] 
**IkeVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewIpsecTunnelParams

`func NewIpsecTunnelParams() *IpsecTunnelParams`

NewIpsecTunnelParams instantiates a new IpsecTunnelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecTunnelParamsWithDefaults

`func NewIpsecTunnelParamsWithDefaults() *IpsecTunnelParams`

NewIpsecTunnelParamsWithDefaults instantiates a new IpsecTunnelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase2

`func (o *IpsecTunnelParams) GetPhase2() string`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IpsecTunnelParams) GetPhase2Ok() (*string, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IpsecTunnelParams) SetPhase2(v string)`

SetPhase2 sets Phase2 field to given value.

### HasPhase2

`func (o *IpsecTunnelParams) HasPhase2() bool`

HasPhase2 returns a boolean if a field has been set.

### GetOutboundSpi

`func (o *IpsecTunnelParams) GetOutboundSpi() string`

GetOutboundSpi returns the OutboundSpi field if non-nil, zero value otherwise.

### GetOutboundSpiOk

`func (o *IpsecTunnelParams) GetOutboundSpiOk() (*string, bool)`

GetOutboundSpiOk returns a tuple with the OutboundSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSpi

`func (o *IpsecTunnelParams) SetOutboundSpi(v string)`

SetOutboundSpi sets OutboundSpi field to given value.

### HasOutboundSpi

`func (o *IpsecTunnelParams) HasOutboundSpi() bool`

HasOutboundSpi returns a boolean if a field has been set.

### GetInboundSpi

`func (o *IpsecTunnelParams) GetInboundSpi() string`

GetInboundSpi returns the InboundSpi field if non-nil, zero value otherwise.

### GetInboundSpiOk

`func (o *IpsecTunnelParams) GetInboundSpiOk() (*string, bool)`

GetInboundSpiOk returns a tuple with the InboundSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSpi

`func (o *IpsecTunnelParams) SetInboundSpi(v string)`

SetInboundSpi sets InboundSpi field to given value.

### HasInboundSpi

`func (o *IpsecTunnelParams) HasInboundSpi() bool`

HasInboundSpi returns a boolean if a field has been set.

### GetBytesIn

`func (o *IpsecTunnelParams) GetBytesIn() string`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *IpsecTunnelParams) GetBytesInOk() (*string, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *IpsecTunnelParams) SetBytesIn(v string)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *IpsecTunnelParams) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *IpsecTunnelParams) GetBytesOut() string`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *IpsecTunnelParams) GetBytesOutOk() (*string, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *IpsecTunnelParams) SetBytesOut(v string)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *IpsecTunnelParams) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetEspTimeRemaining

`func (o *IpsecTunnelParams) GetEspTimeRemaining() string`

GetEspTimeRemaining returns the EspTimeRemaining field if non-nil, zero value otherwise.

### GetEspTimeRemainingOk

`func (o *IpsecTunnelParams) GetEspTimeRemainingOk() (*string, bool)`

GetEspTimeRemainingOk returns a tuple with the EspTimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspTimeRemaining

`func (o *IpsecTunnelParams) SetEspTimeRemaining(v string)`

SetEspTimeRemaining sets EspTimeRemaining field to given value.

### HasEspTimeRemaining

`func (o *IpsecTunnelParams) HasEspTimeRemaining() bool`

HasEspTimeRemaining returns a boolean if a field has been set.

### GetEspPort

`func (o *IpsecTunnelParams) GetEspPort() string`

GetEspPort returns the EspPort field if non-nil, zero value otherwise.

### GetEspPortOk

`func (o *IpsecTunnelParams) GetEspPortOk() (*string, bool)`

GetEspPortOk returns a tuple with the EspPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspPort

`func (o *IpsecTunnelParams) SetEspPort(v string)`

SetEspPort sets EspPort field to given value.

### HasEspPort

`func (o *IpsecTunnelParams) HasEspPort() bool`

HasEspPort returns a boolean if a field has been set.

### GetPhase2Algo

`func (o *IpsecTunnelParams) GetPhase2Algo() string`

GetPhase2Algo returns the Phase2Algo field if non-nil, zero value otherwise.

### GetPhase2AlgoOk

`func (o *IpsecTunnelParams) GetPhase2AlgoOk() (*string, bool)`

GetPhase2AlgoOk returns a tuple with the Phase2Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Algo

`func (o *IpsecTunnelParams) SetPhase2Algo(v string)`

SetPhase2Algo sets Phase2Algo field to given value.

### HasPhase2Algo

`func (o *IpsecTunnelParams) HasPhase2Algo() bool`

HasPhase2Algo returns a boolean if a field has been set.

### GetPhase2Hash

`func (o *IpsecTunnelParams) GetPhase2Hash() string`

GetPhase2Hash returns the Phase2Hash field if non-nil, zero value otherwise.

### GetPhase2HashOk

`func (o *IpsecTunnelParams) GetPhase2HashOk() (*string, bool)`

GetPhase2HashOk returns a tuple with the Phase2Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Hash

`func (o *IpsecTunnelParams) SetPhase2Hash(v string)`

SetPhase2Hash sets Phase2Hash field to given value.

### HasPhase2Hash

`func (o *IpsecTunnelParams) HasPhase2Hash() bool`

HasPhase2Hash returns a boolean if a field has been set.

### GetNatT

`func (o *IpsecTunnelParams) GetNatT() string`

GetNatT returns the NatT field if non-nil, zero value otherwise.

### GetNatTOk

`func (o *IpsecTunnelParams) GetNatTOk() (*string, bool)`

GetNatTOk returns a tuple with the NatT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatT

`func (o *IpsecTunnelParams) SetNatT(v string)`

SetNatT sets NatT field to given value.

### HasNatT

`func (o *IpsecTunnelParams) HasNatT() bool`

HasNatT returns a boolean if a field has been set.

### GetDpd

`func (o *IpsecTunnelParams) GetDpd() string`

GetDpd returns the Dpd field if non-nil, zero value otherwise.

### GetDpdOk

`func (o *IpsecTunnelParams) GetDpdOk() (*string, bool)`

GetDpdOk returns a tuple with the Dpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpd

`func (o *IpsecTunnelParams) SetDpd(v string)`

SetDpd sets Dpd field to given value.

### HasDpd

`func (o *IpsecTunnelParams) HasDpd() bool`

HasDpd returns a boolean if a field has been set.

### GetPfsDhGroup

`func (o *IpsecTunnelParams) GetPfsDhGroup() string`

GetPfsDhGroup returns the PfsDhGroup field if non-nil, zero value otherwise.

### GetPfsDhGroupOk

`func (o *IpsecTunnelParams) GetPfsDhGroupOk() (*string, bool)`

GetPfsDhGroupOk returns a tuple with the PfsDhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsDhGroup

`func (o *IpsecTunnelParams) SetPfsDhGroup(v string)`

SetPfsDhGroup sets PfsDhGroup field to given value.

### HasPfsDhGroup

`func (o *IpsecTunnelParams) HasPfsDhGroup() bool`

HasPfsDhGroup returns a boolean if a field has been set.

### SetPfsDhGroupNil

`func (o *IpsecTunnelParams) SetPfsDhGroupNil(b bool)`

 SetPfsDhGroupNil sets the value for PfsDhGroup to be an explicit nil

### UnsetPfsDhGroup
`func (o *IpsecTunnelParams) UnsetPfsDhGroup()`

UnsetPfsDhGroup ensures that no value is present for PfsDhGroup, not even an explicit nil
### GetPhase1

`func (o *IpsecTunnelParams) GetPhase1() string`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IpsecTunnelParams) GetPhase1Ok() (*string, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IpsecTunnelParams) SetPhase1(v string)`

SetPhase1 sets Phase1 field to given value.

### HasPhase1

`func (o *IpsecTunnelParams) HasPhase1() bool`

HasPhase1 returns a boolean if a field has been set.

### GetIsakmpPort

`func (o *IpsecTunnelParams) GetIsakmpPort() string`

GetIsakmpPort returns the IsakmpPort field if non-nil, zero value otherwise.

### GetIsakmpPortOk

`func (o *IpsecTunnelParams) GetIsakmpPortOk() (*string, bool)`

GetIsakmpPortOk returns a tuple with the IsakmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsakmpPort

`func (o *IpsecTunnelParams) SetIsakmpPort(v string)`

SetIsakmpPort sets IsakmpPort field to given value.

### HasIsakmpPort

`func (o *IpsecTunnelParams) HasIsakmpPort() bool`

HasIsakmpPort returns a boolean if a field has been set.

### GetIsakmpTimeRemaining

`func (o *IpsecTunnelParams) GetIsakmpTimeRemaining() string`

GetIsakmpTimeRemaining returns the IsakmpTimeRemaining field if non-nil, zero value otherwise.

### GetIsakmpTimeRemainingOk

`func (o *IpsecTunnelParams) GetIsakmpTimeRemainingOk() (*string, bool)`

GetIsakmpTimeRemainingOk returns a tuple with the IsakmpTimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsakmpTimeRemaining

`func (o *IpsecTunnelParams) SetIsakmpTimeRemaining(v string)`

SetIsakmpTimeRemaining sets IsakmpTimeRemaining field to given value.

### HasIsakmpTimeRemaining

`func (o *IpsecTunnelParams) HasIsakmpTimeRemaining() bool`

HasIsakmpTimeRemaining returns a boolean if a field has been set.

### GetLastDpd

`func (o *IpsecTunnelParams) GetLastDpd() string`

GetLastDpd returns the LastDpd field if non-nil, zero value otherwise.

### GetLastDpdOk

`func (o *IpsecTunnelParams) GetLastDpdOk() (*string, bool)`

GetLastDpdOk returns a tuple with the LastDpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDpd

`func (o *IpsecTunnelParams) SetLastDpd(v string)`

SetLastDpd sets LastDpd field to given value.

### HasLastDpd

`func (o *IpsecTunnelParams) HasLastDpd() bool`

HasLastDpd returns a boolean if a field has been set.

### GetPhase1Cipher

`func (o *IpsecTunnelParams) GetPhase1Cipher() string`

GetPhase1Cipher returns the Phase1Cipher field if non-nil, zero value otherwise.

### GetPhase1CipherOk

`func (o *IpsecTunnelParams) GetPhase1CipherOk() (*string, bool)`

GetPhase1CipherOk returns a tuple with the Phase1Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Cipher

`func (o *IpsecTunnelParams) SetPhase1Cipher(v string)`

SetPhase1Cipher sets Phase1Cipher field to given value.

### HasPhase1Cipher

`func (o *IpsecTunnelParams) HasPhase1Cipher() bool`

HasPhase1Cipher returns a boolean if a field has been set.

### SetPhase1CipherNil

`func (o *IpsecTunnelParams) SetPhase1CipherNil(b bool)`

 SetPhase1CipherNil sets the value for Phase1Cipher to be an explicit nil

### UnsetPhase1Cipher
`func (o *IpsecTunnelParams) UnsetPhase1Cipher()`

UnsetPhase1Cipher ensures that no value is present for Phase1Cipher, not even an explicit nil
### GetPhase1Prf

`func (o *IpsecTunnelParams) GetPhase1Prf() string`

GetPhase1Prf returns the Phase1Prf field if non-nil, zero value otherwise.

### GetPhase1PrfOk

`func (o *IpsecTunnelParams) GetPhase1PrfOk() (*string, bool)`

GetPhase1PrfOk returns a tuple with the Phase1Prf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Prf

`func (o *IpsecTunnelParams) SetPhase1Prf(v string)`

SetPhase1Prf sets Phase1Prf field to given value.

### HasPhase1Prf

`func (o *IpsecTunnelParams) HasPhase1Prf() bool`

HasPhase1Prf returns a boolean if a field has been set.

### SetPhase1PrfNil

`func (o *IpsecTunnelParams) SetPhase1PrfNil(b bool)`

 SetPhase1PrfNil sets the value for Phase1Prf to be an explicit nil

### UnsetPhase1Prf
`func (o *IpsecTunnelParams) UnsetPhase1Prf()`

UnsetPhase1Prf ensures that no value is present for Phase1Prf, not even an explicit nil
### GetPhase1DhGroup

`func (o *IpsecTunnelParams) GetPhase1DhGroup() string`

GetPhase1DhGroup returns the Phase1DhGroup field if non-nil, zero value otherwise.

### GetPhase1DhGroupOk

`func (o *IpsecTunnelParams) GetPhase1DhGroupOk() (*string, bool)`

GetPhase1DhGroupOk returns a tuple with the Phase1DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroup

`func (o *IpsecTunnelParams) SetPhase1DhGroup(v string)`

SetPhase1DhGroup sets Phase1DhGroup field to given value.

### HasPhase1DhGroup

`func (o *IpsecTunnelParams) HasPhase1DhGroup() bool`

HasPhase1DhGroup returns a boolean if a field has been set.

### SetPhase1DhGroupNil

`func (o *IpsecTunnelParams) SetPhase1DhGroupNil(b bool)`

 SetPhase1DhGroupNil sets the value for Phase1DhGroup to be an explicit nil

### UnsetPhase1DhGroup
`func (o *IpsecTunnelParams) UnsetPhase1DhGroup()`

UnsetPhase1DhGroup ensures that no value is present for Phase1DhGroup, not even an explicit nil
### GetIkeVersion

`func (o *IpsecTunnelParams) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *IpsecTunnelParams) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *IpsecTunnelParams) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *IpsecTunnelParams) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



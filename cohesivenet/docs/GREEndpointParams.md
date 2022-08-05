# GREEndpointParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IpInternal** | Pointer to **NullableString** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**MaskBits** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**SystemDefault** | Pointer to **bool** |  | [optional] [default to false]
**LocalConnectionIp** | Pointer to **NullableString** |  | [optional] 
**RemoteConnectionIp** | Pointer to **NullableString** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] [default to 255]

## Methods

### NewGREEndpointParams

`func NewGREEndpointParams() *GREEndpointParams`

NewGREEndpointParams instantiates a new GREEndpointParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGREEndpointParamsWithDefaults

`func NewGREEndpointParamsWithDefaults() *GREEndpointParams`

NewGREEndpointParamsWithDefaults instantiates a new GREEndpointParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *GREEndpointParams) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *GREEndpointParams) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *GREEndpointParams) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *GREEndpointParams) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetDescription

`func (o *GREEndpointParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GREEndpointParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GREEndpointParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GREEndpointParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GREEndpointParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GREEndpointParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpInternal

`func (o *GREEndpointParams) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *GREEndpointParams) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *GREEndpointParams) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *GREEndpointParams) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### SetIpInternalNil

`func (o *GREEndpointParams) SetIpInternalNil(b bool)`

 SetIpInternalNil sets the value for IpInternal to be an explicit nil

### UnsetIpInternal
`func (o *GREEndpointParams) UnsetIpInternal()`

UnsetIpInternal ensures that no value is present for IpInternal, not even an explicit nil
### GetMtu

`func (o *GREEndpointParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *GREEndpointParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *GREEndpointParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *GREEndpointParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *GREEndpointParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GREEndpointParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GREEndpointParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GREEndpointParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaskBits

`func (o *GREEndpointParams) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *GREEndpointParams) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *GREEndpointParams) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *GREEndpointParams) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### SetMaskBitsNil

`func (o *GREEndpointParams) SetMaskBitsNil(b bool)`

 SetMaskBitsNil sets the value for MaskBits to be an explicit nil

### UnsetMaskBits
`func (o *GREEndpointParams) UnsetMaskBits()`

UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil
### GetGateway

`func (o *GREEndpointParams) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GREEndpointParams) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GREEndpointParams) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GREEndpointParams) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GREEndpointParams) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GREEndpointParams) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSystemDefault

`func (o *GREEndpointParams) GetSystemDefault() bool`

GetSystemDefault returns the SystemDefault field if non-nil, zero value otherwise.

### GetSystemDefaultOk

`func (o *GREEndpointParams) GetSystemDefaultOk() (*bool, bool)`

GetSystemDefaultOk returns a tuple with the SystemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefault

`func (o *GREEndpointParams) SetSystemDefault(v bool)`

SetSystemDefault sets SystemDefault field to given value.

### HasSystemDefault

`func (o *GREEndpointParams) HasSystemDefault() bool`

HasSystemDefault returns a boolean if a field has been set.

### GetLocalConnectionIp

`func (o *GREEndpointParams) GetLocalConnectionIp() string`

GetLocalConnectionIp returns the LocalConnectionIp field if non-nil, zero value otherwise.

### GetLocalConnectionIpOk

`func (o *GREEndpointParams) GetLocalConnectionIpOk() (*string, bool)`

GetLocalConnectionIpOk returns a tuple with the LocalConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionIp

`func (o *GREEndpointParams) SetLocalConnectionIp(v string)`

SetLocalConnectionIp sets LocalConnectionIp field to given value.

### HasLocalConnectionIp

`func (o *GREEndpointParams) HasLocalConnectionIp() bool`

HasLocalConnectionIp returns a boolean if a field has been set.

### SetLocalConnectionIpNil

`func (o *GREEndpointParams) SetLocalConnectionIpNil(b bool)`

 SetLocalConnectionIpNil sets the value for LocalConnectionIp to be an explicit nil

### UnsetLocalConnectionIp
`func (o *GREEndpointParams) UnsetLocalConnectionIp()`

UnsetLocalConnectionIp ensures that no value is present for LocalConnectionIp, not even an explicit nil
### GetRemoteConnectionIp

`func (o *GREEndpointParams) GetRemoteConnectionIp() string`

GetRemoteConnectionIp returns the RemoteConnectionIp field if non-nil, zero value otherwise.

### GetRemoteConnectionIpOk

`func (o *GREEndpointParams) GetRemoteConnectionIpOk() (*string, bool)`

GetRemoteConnectionIpOk returns a tuple with the RemoteConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionIp

`func (o *GREEndpointParams) SetRemoteConnectionIp(v string)`

SetRemoteConnectionIp sets RemoteConnectionIp field to given value.

### HasRemoteConnectionIp

`func (o *GREEndpointParams) HasRemoteConnectionIp() bool`

HasRemoteConnectionIp returns a boolean if a field has been set.

### SetRemoteConnectionIpNil

`func (o *GREEndpointParams) SetRemoteConnectionIpNil(b bool)`

 SetRemoteConnectionIpNil sets the value for RemoteConnectionIp to be an explicit nil

### UnsetRemoteConnectionIp
`func (o *GREEndpointParams) UnsetRemoteConnectionIp()`

UnsetRemoteConnectionIp ensures that no value is present for RemoteConnectionIp, not even an explicit nil
### GetTtl

`func (o *GREEndpointParams) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GREEndpointParams) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GREEndpointParams) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GREEndpointParams) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PutUpdateSystemInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpInternal** | Pointer to **string** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**MaskBits** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPutUpdateSystemInterfaceRequest

`func NewPutUpdateSystemInterfaceRequest() *PutUpdateSystemInterfaceRequest`

NewPutUpdateSystemInterfaceRequest instantiates a new PutUpdateSystemInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUpdateSystemInterfaceRequestWithDefaults

`func NewPutUpdateSystemInterfaceRequestWithDefaults() *PutUpdateSystemInterfaceRequest`

NewPutUpdateSystemInterfaceRequestWithDefaults instantiates a new PutUpdateSystemInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutUpdateSystemInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutUpdateSystemInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutUpdateSystemInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutUpdateSystemInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PutUpdateSystemInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutUpdateSystemInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutUpdateSystemInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutUpdateSystemInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpInternal

`func (o *PutUpdateSystemInterfaceRequest) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *PutUpdateSystemInterfaceRequest) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *PutUpdateSystemInterfaceRequest) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *PutUpdateSystemInterfaceRequest) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### GetIpExternal

`func (o *PutUpdateSystemInterfaceRequest) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *PutUpdateSystemInterfaceRequest) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *PutUpdateSystemInterfaceRequest) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *PutUpdateSystemInterfaceRequest) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *PutUpdateSystemInterfaceRequest) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *PutUpdateSystemInterfaceRequest) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetMtu

`func (o *PutUpdateSystemInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PutUpdateSystemInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PutUpdateSystemInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PutUpdateSystemInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *PutUpdateSystemInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutUpdateSystemInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutUpdateSystemInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PutUpdateSystemInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaskBits

`func (o *PutUpdateSystemInterfaceRequest) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *PutUpdateSystemInterfaceRequest) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *PutUpdateSystemInterfaceRequest) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *PutUpdateSystemInterfaceRequest) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### GetGateway

`func (o *PutUpdateSystemInterfaceRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PutUpdateSystemInterfaceRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PutUpdateSystemInterfaceRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PutUpdateSystemInterfaceRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PutUpdateSystemInterfaceRequest) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PutUpdateSystemInterfaceRequest) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



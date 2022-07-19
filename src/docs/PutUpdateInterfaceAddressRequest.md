# PutUpdateInterfaceAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpInternal** | Pointer to **string** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**MaskBits** | Pointer to **int32** | between 1 and 32 inclusive | [optional] [default to 32]
**Gateway** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPutUpdateInterfaceAddressRequest

`func NewPutUpdateInterfaceAddressRequest() *PutUpdateInterfaceAddressRequest`

NewPutUpdateInterfaceAddressRequest instantiates a new PutUpdateInterfaceAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUpdateInterfaceAddressRequestWithDefaults

`func NewPutUpdateInterfaceAddressRequestWithDefaults() *PutUpdateInterfaceAddressRequest`

NewPutUpdateInterfaceAddressRequestWithDefaults instantiates a new PutUpdateInterfaceAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PutUpdateInterfaceAddressRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutUpdateInterfaceAddressRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutUpdateInterfaceAddressRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutUpdateInterfaceAddressRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PutUpdateInterfaceAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutUpdateInterfaceAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutUpdateInterfaceAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutUpdateInterfaceAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpInternal

`func (o *PutUpdateInterfaceAddressRequest) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *PutUpdateInterfaceAddressRequest) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *PutUpdateInterfaceAddressRequest) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *PutUpdateInterfaceAddressRequest) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### GetIpExternal

`func (o *PutUpdateInterfaceAddressRequest) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *PutUpdateInterfaceAddressRequest) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *PutUpdateInterfaceAddressRequest) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *PutUpdateInterfaceAddressRequest) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *PutUpdateInterfaceAddressRequest) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *PutUpdateInterfaceAddressRequest) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetMaskBits

`func (o *PutUpdateInterfaceAddressRequest) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *PutUpdateInterfaceAddressRequest) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *PutUpdateInterfaceAddressRequest) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *PutUpdateInterfaceAddressRequest) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### GetGateway

`func (o *PutUpdateInterfaceAddressRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PutUpdateInterfaceAddressRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PutUpdateInterfaceAddressRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PutUpdateInterfaceAddressRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PutUpdateInterfaceAddressRequest) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PutUpdateInterfaceAddressRequest) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



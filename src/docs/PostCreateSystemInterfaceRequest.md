# PostCreateSystemInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IpInternal** | Pointer to **string** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**MaskBits** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostCreateSystemInterfaceRequest

`func NewPostCreateSystemInterfaceRequest(name string, ) *PostCreateSystemInterfaceRequest`

NewPostCreateSystemInterfaceRequest instantiates a new PostCreateSystemInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateSystemInterfaceRequestWithDefaults

`func NewPostCreateSystemInterfaceRequestWithDefaults() *PostCreateSystemInterfaceRequest`

NewPostCreateSystemInterfaceRequestWithDefaults instantiates a new PostCreateSystemInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostCreateSystemInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCreateSystemInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCreateSystemInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostCreateSystemInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostCreateSystemInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostCreateSystemInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostCreateSystemInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpInternal

`func (o *PostCreateSystemInterfaceRequest) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *PostCreateSystemInterfaceRequest) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *PostCreateSystemInterfaceRequest) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *PostCreateSystemInterfaceRequest) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### GetIpExternal

`func (o *PostCreateSystemInterfaceRequest) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *PostCreateSystemInterfaceRequest) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *PostCreateSystemInterfaceRequest) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *PostCreateSystemInterfaceRequest) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *PostCreateSystemInterfaceRequest) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *PostCreateSystemInterfaceRequest) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetMtu

`func (o *PostCreateSystemInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PostCreateSystemInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PostCreateSystemInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PostCreateSystemInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *PostCreateSystemInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostCreateSystemInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostCreateSystemInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostCreateSystemInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaskBits

`func (o *PostCreateSystemInterfaceRequest) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *PostCreateSystemInterfaceRequest) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *PostCreateSystemInterfaceRequest) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *PostCreateSystemInterfaceRequest) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### GetGateway

`func (o *PostCreateSystemInterfaceRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PostCreateSystemInterfaceRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PostCreateSystemInterfaceRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PostCreateSystemInterfaceRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PostCreateSystemInterfaceRequest) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PostCreateSystemInterfaceRequest) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



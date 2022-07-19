# PostCreateGreInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpInternal** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**MaskBits** | Pointer to **string** |  | [optional] 
**LocalConnectionIp** | Pointer to **string** |  | [optional] 
**RemoteConnectionIp** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] [default to 255]

## Methods

### NewPostCreateGreInterfaceRequest

`func NewPostCreateGreInterfaceRequest() *PostCreateGreInterfaceRequest`

NewPostCreateGreInterfaceRequest instantiates a new PostCreateGreInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateGreInterfaceRequestWithDefaults

`func NewPostCreateGreInterfaceRequestWithDefaults() *PostCreateGreInterfaceRequest`

NewPostCreateGreInterfaceRequestWithDefaults instantiates a new PostCreateGreInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *PostCreateGreInterfaceRequest) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *PostCreateGreInterfaceRequest) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *PostCreateGreInterfaceRequest) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *PostCreateGreInterfaceRequest) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetDescription

`func (o *PostCreateGreInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostCreateGreInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostCreateGreInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostCreateGreInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpInternal

`func (o *PostCreateGreInterfaceRequest) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *PostCreateGreInterfaceRequest) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *PostCreateGreInterfaceRequest) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *PostCreateGreInterfaceRequest) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### GetMtu

`func (o *PostCreateGreInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PostCreateGreInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PostCreateGreInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PostCreateGreInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *PostCreateGreInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostCreateGreInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostCreateGreInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostCreateGreInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaskBits

`func (o *PostCreateGreInterfaceRequest) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *PostCreateGreInterfaceRequest) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *PostCreateGreInterfaceRequest) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *PostCreateGreInterfaceRequest) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### GetLocalConnectionIp

`func (o *PostCreateGreInterfaceRequest) GetLocalConnectionIp() string`

GetLocalConnectionIp returns the LocalConnectionIp field if non-nil, zero value otherwise.

### GetLocalConnectionIpOk

`func (o *PostCreateGreInterfaceRequest) GetLocalConnectionIpOk() (*string, bool)`

GetLocalConnectionIpOk returns a tuple with the LocalConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionIp

`func (o *PostCreateGreInterfaceRequest) SetLocalConnectionIp(v string)`

SetLocalConnectionIp sets LocalConnectionIp field to given value.

### HasLocalConnectionIp

`func (o *PostCreateGreInterfaceRequest) HasLocalConnectionIp() bool`

HasLocalConnectionIp returns a boolean if a field has been set.

### GetRemoteConnectionIp

`func (o *PostCreateGreInterfaceRequest) GetRemoteConnectionIp() string`

GetRemoteConnectionIp returns the RemoteConnectionIp field if non-nil, zero value otherwise.

### GetRemoteConnectionIpOk

`func (o *PostCreateGreInterfaceRequest) GetRemoteConnectionIpOk() (*string, bool)`

GetRemoteConnectionIpOk returns a tuple with the RemoteConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionIp

`func (o *PostCreateGreInterfaceRequest) SetRemoteConnectionIp(v string)`

SetRemoteConnectionIp sets RemoteConnectionIp field to given value.

### HasRemoteConnectionIp

`func (o *PostCreateGreInterfaceRequest) HasRemoteConnectionIp() bool`

HasRemoteConnectionIp returns a boolean if a field has been set.

### GetTtl

`func (o *PostCreateGreInterfaceRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostCreateGreInterfaceRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostCreateGreInterfaceRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostCreateGreInterfaceRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GREEndpoint

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
**Id** | Pointer to **int32** |  | [optional] 
**InterfaceId** | Pointer to **int32** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InterfaceType** | Pointer to **string** | system or system_virtual | [optional] 
**Status** | Pointer to **string** | Availability of interface, Up or Down | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGREEndpoint

`func NewGREEndpoint() *GREEndpoint`

NewGREEndpoint instantiates a new GREEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGREEndpointWithDefaults

`func NewGREEndpointWithDefaults() *GREEndpoint`

NewGREEndpointWithDefaults instantiates a new GREEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *GREEndpoint) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *GREEndpoint) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *GREEndpoint) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *GREEndpoint) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetDescription

`func (o *GREEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GREEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GREEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GREEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GREEndpoint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GREEndpoint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpInternal

`func (o *GREEndpoint) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *GREEndpoint) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *GREEndpoint) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *GREEndpoint) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### SetIpInternalNil

`func (o *GREEndpoint) SetIpInternalNil(b bool)`

 SetIpInternalNil sets the value for IpInternal to be an explicit nil

### UnsetIpInternal
`func (o *GREEndpoint) UnsetIpInternal()`

UnsetIpInternal ensures that no value is present for IpInternal, not even an explicit nil
### GetMtu

`func (o *GREEndpoint) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *GREEndpoint) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *GREEndpoint) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *GREEndpoint) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *GREEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GREEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GREEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GREEndpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaskBits

`func (o *GREEndpoint) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *GREEndpoint) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *GREEndpoint) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *GREEndpoint) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### SetMaskBitsNil

`func (o *GREEndpoint) SetMaskBitsNil(b bool)`

 SetMaskBitsNil sets the value for MaskBits to be an explicit nil

### UnsetMaskBits
`func (o *GREEndpoint) UnsetMaskBits()`

UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil
### GetGateway

`func (o *GREEndpoint) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GREEndpoint) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GREEndpoint) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GREEndpoint) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GREEndpoint) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GREEndpoint) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSystemDefault

`func (o *GREEndpoint) GetSystemDefault() bool`

GetSystemDefault returns the SystemDefault field if non-nil, zero value otherwise.

### GetSystemDefaultOk

`func (o *GREEndpoint) GetSystemDefaultOk() (*bool, bool)`

GetSystemDefaultOk returns a tuple with the SystemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefault

`func (o *GREEndpoint) SetSystemDefault(v bool)`

SetSystemDefault sets SystemDefault field to given value.

### HasSystemDefault

`func (o *GREEndpoint) HasSystemDefault() bool`

HasSystemDefault returns a boolean if a field has been set.

### GetLocalConnectionIp

`func (o *GREEndpoint) GetLocalConnectionIp() string`

GetLocalConnectionIp returns the LocalConnectionIp field if non-nil, zero value otherwise.

### GetLocalConnectionIpOk

`func (o *GREEndpoint) GetLocalConnectionIpOk() (*string, bool)`

GetLocalConnectionIpOk returns a tuple with the LocalConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionIp

`func (o *GREEndpoint) SetLocalConnectionIp(v string)`

SetLocalConnectionIp sets LocalConnectionIp field to given value.

### HasLocalConnectionIp

`func (o *GREEndpoint) HasLocalConnectionIp() bool`

HasLocalConnectionIp returns a boolean if a field has been set.

### SetLocalConnectionIpNil

`func (o *GREEndpoint) SetLocalConnectionIpNil(b bool)`

 SetLocalConnectionIpNil sets the value for LocalConnectionIp to be an explicit nil

### UnsetLocalConnectionIp
`func (o *GREEndpoint) UnsetLocalConnectionIp()`

UnsetLocalConnectionIp ensures that no value is present for LocalConnectionIp, not even an explicit nil
### GetRemoteConnectionIp

`func (o *GREEndpoint) GetRemoteConnectionIp() string`

GetRemoteConnectionIp returns the RemoteConnectionIp field if non-nil, zero value otherwise.

### GetRemoteConnectionIpOk

`func (o *GREEndpoint) GetRemoteConnectionIpOk() (*string, bool)`

GetRemoteConnectionIpOk returns a tuple with the RemoteConnectionIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionIp

`func (o *GREEndpoint) SetRemoteConnectionIp(v string)`

SetRemoteConnectionIp sets RemoteConnectionIp field to given value.

### HasRemoteConnectionIp

`func (o *GREEndpoint) HasRemoteConnectionIp() bool`

HasRemoteConnectionIp returns a boolean if a field has been set.

### SetRemoteConnectionIpNil

`func (o *GREEndpoint) SetRemoteConnectionIpNil(b bool)`

 SetRemoteConnectionIpNil sets the value for RemoteConnectionIp to be an explicit nil

### UnsetRemoteConnectionIp
`func (o *GREEndpoint) UnsetRemoteConnectionIp()`

UnsetRemoteConnectionIp ensures that no value is present for RemoteConnectionIp, not even an explicit nil
### GetTtl

`func (o *GREEndpoint) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GREEndpoint) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GREEndpoint) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GREEndpoint) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetId

`func (o *GREEndpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GREEndpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GREEndpoint) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GREEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *GREEndpoint) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GREEndpoint) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GREEndpoint) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GREEndpoint) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetIpExternal

`func (o *GREEndpoint) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *GREEndpoint) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *GREEndpoint) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *GREEndpoint) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *GREEndpoint) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *GREEndpoint) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetName

`func (o *GREEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GREEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GREEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GREEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *GREEndpoint) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *GREEndpoint) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *GREEndpoint) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *GREEndpoint) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetStatus

`func (o *GREEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GREEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GREEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GREEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GREEndpoint) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GREEndpoint) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GREEndpoint) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GREEndpoint) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



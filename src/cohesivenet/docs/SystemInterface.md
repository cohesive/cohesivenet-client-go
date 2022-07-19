# SystemInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InterfaceType** | Pointer to **string** | system or system_virtual | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpInternal** | Pointer to **NullableString** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | Availability of interface, Up or Down | [optional] 
**MaskBits** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**SystemDefault** | Pointer to **bool** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSystemInterface

`func NewSystemInterface() *SystemInterface`

NewSystemInterface instantiates a new SystemInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInterfaceWithDefaults

`func NewSystemInterfaceWithDefaults() *SystemInterface`

NewSystemInterfaceWithDefaults instantiates a new SystemInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemInterface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SystemInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SystemInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *SystemInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *SystemInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *SystemInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *SystemInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetDescription

`func (o *SystemInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpInternal

`func (o *SystemInterface) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *SystemInterface) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *SystemInterface) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *SystemInterface) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### SetIpInternalNil

`func (o *SystemInterface) SetIpInternalNil(b bool)`

 SetIpInternalNil sets the value for IpInternal to be an explicit nil

### UnsetIpInternal
`func (o *SystemInterface) UnsetIpInternal()`

UnsetIpInternal ensures that no value is present for IpInternal, not even an explicit nil
### GetMtu

`func (o *SystemInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SystemInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SystemInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SystemInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *SystemInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SystemInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SystemInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SystemInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *SystemInterface) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemInterface) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemInterface) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemInterface) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaskBits

`func (o *SystemInterface) GetMaskBits() string`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *SystemInterface) GetMaskBitsOk() (*string, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *SystemInterface) SetMaskBits(v string)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *SystemInterface) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### SetMaskBitsNil

`func (o *SystemInterface) SetMaskBitsNil(b bool)`

 SetMaskBitsNil sets the value for MaskBits to be an explicit nil

### UnsetMaskBits
`func (o *SystemInterface) UnsetMaskBits()`

UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil
### GetGateway

`func (o *SystemInterface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SystemInterface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SystemInterface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SystemInterface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *SystemInterface) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *SystemInterface) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSystemDefault

`func (o *SystemInterface) GetSystemDefault() bool`

GetSystemDefault returns the SystemDefault field if non-nil, zero value otherwise.

### GetSystemDefaultOk

`func (o *SystemInterface) GetSystemDefaultOk() (*bool, bool)`

GetSystemDefaultOk returns a tuple with the SystemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefault

`func (o *SystemInterface) SetSystemDefault(v bool)`

SetSystemDefault sets SystemDefault field to given value.

### HasSystemDefault

`func (o *SystemInterface) HasSystemDefault() bool`

HasSystemDefault returns a boolean if a field has been set.

### GetIpExternal

`func (o *SystemInterface) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *SystemInterface) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *SystemInterface) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *SystemInterface) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *SystemInterface) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *SystemInterface) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetTags

`func (o *SystemInterface) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SystemInterface) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SystemInterface) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SystemInterface) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



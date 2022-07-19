# InterfaceAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpInternal** | Pointer to **string** |  | [optional] 
**IpExternal** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**MaskBits** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInterfaceAddress

`func NewInterfaceAddress() *InterfaceAddress`

NewInterfaceAddress instantiates a new InterfaceAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceAddressWithDefaults

`func NewInterfaceAddressWithDefaults() *InterfaceAddress`

NewInterfaceAddressWithDefaults instantiates a new InterfaceAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpInternal

`func (o *InterfaceAddress) GetIpInternal() string`

GetIpInternal returns the IpInternal field if non-nil, zero value otherwise.

### GetIpInternalOk

`func (o *InterfaceAddress) GetIpInternalOk() (*string, bool)`

GetIpInternalOk returns a tuple with the IpInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInternal

`func (o *InterfaceAddress) SetIpInternal(v string)`

SetIpInternal sets IpInternal field to given value.

### HasIpInternal

`func (o *InterfaceAddress) HasIpInternal() bool`

HasIpInternal returns a boolean if a field has been set.

### GetIpExternal

`func (o *InterfaceAddress) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *InterfaceAddress) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *InterfaceAddress) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *InterfaceAddress) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGateway

`func (o *InterfaceAddress) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InterfaceAddress) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InterfaceAddress) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InterfaceAddress) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMaskBits

`func (o *InterfaceAddress) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *InterfaceAddress) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *InterfaceAddress) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *InterfaceAddress) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InterfaceAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InterfaceAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InterfaceAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InterfaceAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InterfaceAddress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InterfaceAddress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InterfaceAddress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InterfaceAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



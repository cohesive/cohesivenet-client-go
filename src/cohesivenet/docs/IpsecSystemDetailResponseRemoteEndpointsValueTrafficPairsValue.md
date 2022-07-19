# IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RemoteSubnet** | Pointer to **string** |  | [optional] 
**LocalSubnet** | Pointer to **string** |  | [optional] 
**PingIpaddress** | Pointer to **NullableString** |  | [optional] 
**PingInterface** | Pointer to **string** |  | [optional] 
**PingInterval** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IpsecEndpointId** | Pointer to **int32** |  | [optional] 
**EndpointId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue

`func NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue() *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue`

NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValueWithDefaults

`func NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValueWithDefaults() *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue`

NewIpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValueWithDefaults instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetLocalSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetPingIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingIpaddress() string`

GetPingIpaddress returns the PingIpaddress field if non-nil, zero value otherwise.

### GetPingIpaddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingIpaddressOk() (*string, bool)`

GetPingIpaddressOk returns a tuple with the PingIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetPingIpaddress(v string)`

SetPingIpaddress sets PingIpaddress field to given value.

### HasPingIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasPingIpaddress() bool`

HasPingIpaddress returns a boolean if a field has been set.

### SetPingIpaddressNil

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetPingIpaddressNil(b bool)`

 SetPingIpaddressNil sets the value for PingIpaddress to be an explicit nil

### UnsetPingIpaddress
`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) UnsetPingIpaddress()`

UnsetPingIpaddress ensures that no value is present for PingIpaddress, not even an explicit nil
### GetPingInterface

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingInterface() string`

GetPingInterface returns the PingInterface field if non-nil, zero value otherwise.

### GetPingInterfaceOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingInterfaceOk() (*string, bool)`

GetPingInterfaceOk returns a tuple with the PingInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterface

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetPingInterface(v string)`

SetPingInterface sets PingInterface field to given value.

### HasPingInterface

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasPingInterface() bool`

HasPingInterface returns a boolean if a field has been set.

### GetPingInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingInterval() int32`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetPingIntervalOk() (*int32, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetPingInterval(v int32)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpsecEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetIpsecEndpointId() int32`

GetIpsecEndpointId returns the IpsecEndpointId field if non-nil, zero value otherwise.

### GetIpsecEndpointIdOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetIpsecEndpointIdOk() (*int32, bool)`

GetIpsecEndpointIdOk returns a tuple with the IpsecEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetIpsecEndpointId(v int32)`

SetIpsecEndpointId sets IpsecEndpointId field to given value.

### HasIpsecEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasIpsecEndpointId() bool`

HasIpsecEndpointId returns a boolean if a field has been set.

### GetEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



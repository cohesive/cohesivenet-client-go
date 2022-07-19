# ContainerNetworkSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** |  | [optional] 
**SandboxID** | Pointer to **string** |  | [optional] 
**HairpinMode** | Pointer to **bool** |  | [optional] 
**Bridge** | Pointer to **string** |  | [optional] 
**LinkLocalIPv6Address** | Pointer to **string** |  | [optional] 
**LinkLocalIPv6PrefixLen** | Pointer to **int32** |  | [optional] 
**SandboxKey** | Pointer to **string** |  | [optional] 
**SecondaryIPAddresses** | Pointer to **NullableString** |  | [optional] 
**SecondaryIPv6Addresses** | Pointer to **NullableString** |  | [optional] 
**EndpointID** | Pointer to **string** |  | [optional] 
**GlobalIPv6Address** | Pointer to **string** |  | [optional] 
**GlobalIPv6PrefixLen** | Pointer to **int32** |  | [optional] 
**IPAddress** | Pointer to **string** |  | [optional] 
**IPPrefixLen** | Pointer to **int32** |  | [optional] 
**IPv6Gateway** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**PortMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewContainerNetworkSettings

`func NewContainerNetworkSettings() *ContainerNetworkSettings`

NewContainerNetworkSettings instantiates a new ContainerNetworkSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerNetworkSettingsWithDefaults

`func NewContainerNetworkSettingsWithDefaults() *ContainerNetworkSettings`

NewContainerNetworkSettingsWithDefaults instantiates a new ContainerNetworkSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *ContainerNetworkSettings) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ContainerNetworkSettings) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ContainerNetworkSettings) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ContainerNetworkSettings) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetSandboxID

`func (o *ContainerNetworkSettings) GetSandboxID() string`

GetSandboxID returns the SandboxID field if non-nil, zero value otherwise.

### GetSandboxIDOk

`func (o *ContainerNetworkSettings) GetSandboxIDOk() (*string, bool)`

GetSandboxIDOk returns a tuple with the SandboxID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxID

`func (o *ContainerNetworkSettings) SetSandboxID(v string)`

SetSandboxID sets SandboxID field to given value.

### HasSandboxID

`func (o *ContainerNetworkSettings) HasSandboxID() bool`

HasSandboxID returns a boolean if a field has been set.

### GetHairpinMode

`func (o *ContainerNetworkSettings) GetHairpinMode() bool`

GetHairpinMode returns the HairpinMode field if non-nil, zero value otherwise.

### GetHairpinModeOk

`func (o *ContainerNetworkSettings) GetHairpinModeOk() (*bool, bool)`

GetHairpinModeOk returns a tuple with the HairpinMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHairpinMode

`func (o *ContainerNetworkSettings) SetHairpinMode(v bool)`

SetHairpinMode sets HairpinMode field to given value.

### HasHairpinMode

`func (o *ContainerNetworkSettings) HasHairpinMode() bool`

HasHairpinMode returns a boolean if a field has been set.

### GetBridge

`func (o *ContainerNetworkSettings) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *ContainerNetworkSettings) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *ContainerNetworkSettings) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *ContainerNetworkSettings) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetLinkLocalIPv6Address

`func (o *ContainerNetworkSettings) GetLinkLocalIPv6Address() string`

GetLinkLocalIPv6Address returns the LinkLocalIPv6Address field if non-nil, zero value otherwise.

### GetLinkLocalIPv6AddressOk

`func (o *ContainerNetworkSettings) GetLinkLocalIPv6AddressOk() (*string, bool)`

GetLinkLocalIPv6AddressOk returns a tuple with the LinkLocalIPv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkLocalIPv6Address

`func (o *ContainerNetworkSettings) SetLinkLocalIPv6Address(v string)`

SetLinkLocalIPv6Address sets LinkLocalIPv6Address field to given value.

### HasLinkLocalIPv6Address

`func (o *ContainerNetworkSettings) HasLinkLocalIPv6Address() bool`

HasLinkLocalIPv6Address returns a boolean if a field has been set.

### GetLinkLocalIPv6PrefixLen

`func (o *ContainerNetworkSettings) GetLinkLocalIPv6PrefixLen() int32`

GetLinkLocalIPv6PrefixLen returns the LinkLocalIPv6PrefixLen field if non-nil, zero value otherwise.

### GetLinkLocalIPv6PrefixLenOk

`func (o *ContainerNetworkSettings) GetLinkLocalIPv6PrefixLenOk() (*int32, bool)`

GetLinkLocalIPv6PrefixLenOk returns a tuple with the LinkLocalIPv6PrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkLocalIPv6PrefixLen

`func (o *ContainerNetworkSettings) SetLinkLocalIPv6PrefixLen(v int32)`

SetLinkLocalIPv6PrefixLen sets LinkLocalIPv6PrefixLen field to given value.

### HasLinkLocalIPv6PrefixLen

`func (o *ContainerNetworkSettings) HasLinkLocalIPv6PrefixLen() bool`

HasLinkLocalIPv6PrefixLen returns a boolean if a field has been set.

### GetSandboxKey

`func (o *ContainerNetworkSettings) GetSandboxKey() string`

GetSandboxKey returns the SandboxKey field if non-nil, zero value otherwise.

### GetSandboxKeyOk

`func (o *ContainerNetworkSettings) GetSandboxKeyOk() (*string, bool)`

GetSandboxKeyOk returns a tuple with the SandboxKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxKey

`func (o *ContainerNetworkSettings) SetSandboxKey(v string)`

SetSandboxKey sets SandboxKey field to given value.

### HasSandboxKey

`func (o *ContainerNetworkSettings) HasSandboxKey() bool`

HasSandboxKey returns a boolean if a field has been set.

### GetSecondaryIPAddresses

`func (o *ContainerNetworkSettings) GetSecondaryIPAddresses() string`

GetSecondaryIPAddresses returns the SecondaryIPAddresses field if non-nil, zero value otherwise.

### GetSecondaryIPAddressesOk

`func (o *ContainerNetworkSettings) GetSecondaryIPAddressesOk() (*string, bool)`

GetSecondaryIPAddressesOk returns a tuple with the SecondaryIPAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIPAddresses

`func (o *ContainerNetworkSettings) SetSecondaryIPAddresses(v string)`

SetSecondaryIPAddresses sets SecondaryIPAddresses field to given value.

### HasSecondaryIPAddresses

`func (o *ContainerNetworkSettings) HasSecondaryIPAddresses() bool`

HasSecondaryIPAddresses returns a boolean if a field has been set.

### SetSecondaryIPAddressesNil

`func (o *ContainerNetworkSettings) SetSecondaryIPAddressesNil(b bool)`

 SetSecondaryIPAddressesNil sets the value for SecondaryIPAddresses to be an explicit nil

### UnsetSecondaryIPAddresses
`func (o *ContainerNetworkSettings) UnsetSecondaryIPAddresses()`

UnsetSecondaryIPAddresses ensures that no value is present for SecondaryIPAddresses, not even an explicit nil
### GetSecondaryIPv6Addresses

`func (o *ContainerNetworkSettings) GetSecondaryIPv6Addresses() string`

GetSecondaryIPv6Addresses returns the SecondaryIPv6Addresses field if non-nil, zero value otherwise.

### GetSecondaryIPv6AddressesOk

`func (o *ContainerNetworkSettings) GetSecondaryIPv6AddressesOk() (*string, bool)`

GetSecondaryIPv6AddressesOk returns a tuple with the SecondaryIPv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIPv6Addresses

`func (o *ContainerNetworkSettings) SetSecondaryIPv6Addresses(v string)`

SetSecondaryIPv6Addresses sets SecondaryIPv6Addresses field to given value.

### HasSecondaryIPv6Addresses

`func (o *ContainerNetworkSettings) HasSecondaryIPv6Addresses() bool`

HasSecondaryIPv6Addresses returns a boolean if a field has been set.

### SetSecondaryIPv6AddressesNil

`func (o *ContainerNetworkSettings) SetSecondaryIPv6AddressesNil(b bool)`

 SetSecondaryIPv6AddressesNil sets the value for SecondaryIPv6Addresses to be an explicit nil

### UnsetSecondaryIPv6Addresses
`func (o *ContainerNetworkSettings) UnsetSecondaryIPv6Addresses()`

UnsetSecondaryIPv6Addresses ensures that no value is present for SecondaryIPv6Addresses, not even an explicit nil
### GetEndpointID

`func (o *ContainerNetworkSettings) GetEndpointID() string`

GetEndpointID returns the EndpointID field if non-nil, zero value otherwise.

### GetEndpointIDOk

`func (o *ContainerNetworkSettings) GetEndpointIDOk() (*string, bool)`

GetEndpointIDOk returns a tuple with the EndpointID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointID

`func (o *ContainerNetworkSettings) SetEndpointID(v string)`

SetEndpointID sets EndpointID field to given value.

### HasEndpointID

`func (o *ContainerNetworkSettings) HasEndpointID() bool`

HasEndpointID returns a boolean if a field has been set.

### GetGlobalIPv6Address

`func (o *ContainerNetworkSettings) GetGlobalIPv6Address() string`

GetGlobalIPv6Address returns the GlobalIPv6Address field if non-nil, zero value otherwise.

### GetGlobalIPv6AddressOk

`func (o *ContainerNetworkSettings) GetGlobalIPv6AddressOk() (*string, bool)`

GetGlobalIPv6AddressOk returns a tuple with the GlobalIPv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIPv6Address

`func (o *ContainerNetworkSettings) SetGlobalIPv6Address(v string)`

SetGlobalIPv6Address sets GlobalIPv6Address field to given value.

### HasGlobalIPv6Address

`func (o *ContainerNetworkSettings) HasGlobalIPv6Address() bool`

HasGlobalIPv6Address returns a boolean if a field has been set.

### GetGlobalIPv6PrefixLen

`func (o *ContainerNetworkSettings) GetGlobalIPv6PrefixLen() int32`

GetGlobalIPv6PrefixLen returns the GlobalIPv6PrefixLen field if non-nil, zero value otherwise.

### GetGlobalIPv6PrefixLenOk

`func (o *ContainerNetworkSettings) GetGlobalIPv6PrefixLenOk() (*int32, bool)`

GetGlobalIPv6PrefixLenOk returns a tuple with the GlobalIPv6PrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIPv6PrefixLen

`func (o *ContainerNetworkSettings) SetGlobalIPv6PrefixLen(v int32)`

SetGlobalIPv6PrefixLen sets GlobalIPv6PrefixLen field to given value.

### HasGlobalIPv6PrefixLen

`func (o *ContainerNetworkSettings) HasGlobalIPv6PrefixLen() bool`

HasGlobalIPv6PrefixLen returns a boolean if a field has been set.

### GetIPAddress

`func (o *ContainerNetworkSettings) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *ContainerNetworkSettings) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *ContainerNetworkSettings) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *ContainerNetworkSettings) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetIPPrefixLen

`func (o *ContainerNetworkSettings) GetIPPrefixLen() int32`

GetIPPrefixLen returns the IPPrefixLen field if non-nil, zero value otherwise.

### GetIPPrefixLenOk

`func (o *ContainerNetworkSettings) GetIPPrefixLenOk() (*int32, bool)`

GetIPPrefixLenOk returns a tuple with the IPPrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPPrefixLen

`func (o *ContainerNetworkSettings) SetIPPrefixLen(v int32)`

SetIPPrefixLen sets IPPrefixLen field to given value.

### HasIPPrefixLen

`func (o *ContainerNetworkSettings) HasIPPrefixLen() bool`

HasIPPrefixLen returns a boolean if a field has been set.

### GetIPv6Gateway

`func (o *ContainerNetworkSettings) GetIPv6Gateway() string`

GetIPv6Gateway returns the IPv6Gateway field if non-nil, zero value otherwise.

### GetIPv6GatewayOk

`func (o *ContainerNetworkSettings) GetIPv6GatewayOk() (*string, bool)`

GetIPv6GatewayOk returns a tuple with the IPv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPv6Gateway

`func (o *ContainerNetworkSettings) SetIPv6Gateway(v string)`

SetIPv6Gateway sets IPv6Gateway field to given value.

### HasIPv6Gateway

`func (o *ContainerNetworkSettings) HasIPv6Gateway() bool`

HasIPv6Gateway returns a boolean if a field has been set.

### GetMacAddress

`func (o *ContainerNetworkSettings) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ContainerNetworkSettings) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ContainerNetworkSettings) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ContainerNetworkSettings) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPortMapping

`func (o *ContainerNetworkSettings) GetPortMapping() map[string]interface{}`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *ContainerNetworkSettings) GetPortMappingOk() (*map[string]interface{}, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *ContainerNetworkSettings) SetPortMapping(v map[string]interface{})`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *ContainerNetworkSettings) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### SetPortMappingNil

`func (o *ContainerNetworkSettings) SetPortMappingNil(b bool)`

 SetPortMappingNil sets the value for PortMapping to be an explicit nil

### UnsetPortMapping
`func (o *ContainerNetworkSettings) UnsetPortMapping()`

UnsetPortMapping ensures that no value is present for PortMapping, not even an explicit nil
### GetPorts

`func (o *ContainerNetworkSettings) GetPorts() map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ContainerNetworkSettings) GetPortsOk() (*map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ContainerNetworkSettings) SetPorts(v map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ContainerNetworkSettings) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *ContainerNetworkSettings) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *ContainerNetworkSettings) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



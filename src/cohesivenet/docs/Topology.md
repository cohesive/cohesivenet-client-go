# Topology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]TopologyClientsInner**](TopologyClientsInner.md) | IPs for clientpacks | [optional] 
**Managers** | Pointer to [**[]VNS3Controller**](VNS3Controller.md) |  | [optional] 
**TotalClients** | Pointer to **int32** |  | [optional] 
**IpsecMaxSubnets** | Pointer to **int32** |  | [optional] 
**IpsecMaxEndpoints** | Pointer to **int32** |  | [optional] 
**LicenseUpgrades** | Pointer to **[]string** |  | [optional] 
**OverlayMaxClients** | Pointer to **int32** |  | [optional] 
**OverlaySubnet** | Pointer to **string** | CIDR for overlay clients | [optional] 

## Methods

### NewTopology

`func NewTopology() *Topology`

NewTopology instantiates a new Topology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyWithDefaults

`func NewTopologyWithDefaults() *Topology`

NewTopologyWithDefaults instantiates a new Topology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *Topology) GetClients() []TopologyClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *Topology) GetClientsOk() (*[]TopologyClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *Topology) SetClients(v []TopologyClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *Topology) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetManagers

`func (o *Topology) GetManagers() []VNS3Controller`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *Topology) GetManagersOk() (*[]VNS3Controller, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *Topology) SetManagers(v []VNS3Controller)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *Topology) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetTotalClients

`func (o *Topology) GetTotalClients() int32`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *Topology) GetTotalClientsOk() (*int32, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *Topology) SetTotalClients(v int32)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *Topology) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetIpsecMaxSubnets

`func (o *Topology) GetIpsecMaxSubnets() int32`

GetIpsecMaxSubnets returns the IpsecMaxSubnets field if non-nil, zero value otherwise.

### GetIpsecMaxSubnetsOk

`func (o *Topology) GetIpsecMaxSubnetsOk() (*int32, bool)`

GetIpsecMaxSubnetsOk returns a tuple with the IpsecMaxSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecMaxSubnets

`func (o *Topology) SetIpsecMaxSubnets(v int32)`

SetIpsecMaxSubnets sets IpsecMaxSubnets field to given value.

### HasIpsecMaxSubnets

`func (o *Topology) HasIpsecMaxSubnets() bool`

HasIpsecMaxSubnets returns a boolean if a field has been set.

### GetIpsecMaxEndpoints

`func (o *Topology) GetIpsecMaxEndpoints() int32`

GetIpsecMaxEndpoints returns the IpsecMaxEndpoints field if non-nil, zero value otherwise.

### GetIpsecMaxEndpointsOk

`func (o *Topology) GetIpsecMaxEndpointsOk() (*int32, bool)`

GetIpsecMaxEndpointsOk returns a tuple with the IpsecMaxEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecMaxEndpoints

`func (o *Topology) SetIpsecMaxEndpoints(v int32)`

SetIpsecMaxEndpoints sets IpsecMaxEndpoints field to given value.

### HasIpsecMaxEndpoints

`func (o *Topology) HasIpsecMaxEndpoints() bool`

HasIpsecMaxEndpoints returns a boolean if a field has been set.

### GetLicenseUpgrades

`func (o *Topology) GetLicenseUpgrades() []string`

GetLicenseUpgrades returns the LicenseUpgrades field if non-nil, zero value otherwise.

### GetLicenseUpgradesOk

`func (o *Topology) GetLicenseUpgradesOk() (*[]string, bool)`

GetLicenseUpgradesOk returns a tuple with the LicenseUpgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUpgrades

`func (o *Topology) SetLicenseUpgrades(v []string)`

SetLicenseUpgrades sets LicenseUpgrades field to given value.

### HasLicenseUpgrades

`func (o *Topology) HasLicenseUpgrades() bool`

HasLicenseUpgrades returns a boolean if a field has been set.

### GetOverlayMaxClients

`func (o *Topology) GetOverlayMaxClients() int32`

GetOverlayMaxClients returns the OverlayMaxClients field if non-nil, zero value otherwise.

### GetOverlayMaxClientsOk

`func (o *Topology) GetOverlayMaxClientsOk() (*int32, bool)`

GetOverlayMaxClientsOk returns a tuple with the OverlayMaxClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayMaxClients

`func (o *Topology) SetOverlayMaxClients(v int32)`

SetOverlayMaxClients sets OverlayMaxClients field to given value.

### HasOverlayMaxClients

`func (o *Topology) HasOverlayMaxClients() bool`

HasOverlayMaxClients returns a boolean if a field has been set.

### GetOverlaySubnet

`func (o *Topology) GetOverlaySubnet() string`

GetOverlaySubnet returns the OverlaySubnet field if non-nil, zero value otherwise.

### GetOverlaySubnetOk

`func (o *Topology) GetOverlaySubnetOk() (*string, bool)`

GetOverlaySubnetOk returns a tuple with the OverlaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlaySubnet

`func (o *Topology) SetOverlaySubnet(v string)`

SetOverlaySubnet sets OverlaySubnet field to given value.

### HasOverlaySubnet

`func (o *Topology) HasOverlaySubnet() bool`

HasOverlaySubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



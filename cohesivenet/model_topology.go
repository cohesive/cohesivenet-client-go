/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesivenet

import (
	"encoding/json"
)

// Topology struct for Topology
type Topology struct {
	// IPs for clientpacks
	Clients []TopologyClient `json:"clients,omitempty"`
	Managers []VNS3Controller `json:"managers,omitempty"`
	TotalClients *int32 `json:"total_clients,omitempty"`
	IpsecMaxSubnets *int32 `json:"ipsec_max_subnets,omitempty"`
	IpsecMaxEndpoints *int32 `json:"ipsec_max_endpoints,omitempty"`
	LicenseUpgrades []string `json:"license_upgrades,omitempty"`
	OverlayMaxClients *int32 `json:"overlay_max_clients,omitempty"`
	// CIDR for overlay clients
	OverlaySubnet *string `json:"overlay_subnet,omitempty"`
}

// NewTopology instantiates a new Topology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopology() *Topology {
	this := Topology{}
	return &this
}

// NewTopologyWithDefaults instantiates a new Topology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyWithDefaults() *Topology {
	this := Topology{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *Topology) GetClients() []TopologyClient {
	if o == nil || o.Clients == nil {
		var ret []TopologyClient
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetClientsOk() ([]TopologyClient, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *Topology) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given []TopologyClient and assigns it to the Clients field.
func (o *Topology) SetClients(v []TopologyClient) {
	o.Clients = v
}

// GetManagers returns the Managers field value if set, zero value otherwise.
func (o *Topology) GetManagers() []VNS3Controller {
	if o == nil || o.Managers == nil {
		var ret []VNS3Controller
		return ret
	}
	return o.Managers
}

// GetManagersOk returns a tuple with the Managers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetManagersOk() ([]VNS3Controller, bool) {
	if o == nil || o.Managers == nil {
		return nil, false
	}
	return o.Managers, true
}

// HasManagers returns a boolean if a field has been set.
func (o *Topology) HasManagers() bool {
	if o != nil && o.Managers != nil {
		return true
	}

	return false
}

// SetManagers gets a reference to the given []VNS3Controller and assigns it to the Managers field.
func (o *Topology) SetManagers(v []VNS3Controller) {
	o.Managers = v
}

// GetTotalClients returns the TotalClients field value if set, zero value otherwise.
func (o *Topology) GetTotalClients() int32 {
	if o == nil || o.TotalClients == nil {
		var ret int32
		return ret
	}
	return *o.TotalClients
}

// GetTotalClientsOk returns a tuple with the TotalClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetTotalClientsOk() (*int32, bool) {
	if o == nil || o.TotalClients == nil {
		return nil, false
	}
	return o.TotalClients, true
}

// HasTotalClients returns a boolean if a field has been set.
func (o *Topology) HasTotalClients() bool {
	if o != nil && o.TotalClients != nil {
		return true
	}

	return false
}

// SetTotalClients gets a reference to the given int32 and assigns it to the TotalClients field.
func (o *Topology) SetTotalClients(v int32) {
	o.TotalClients = &v
}

// GetIpsecMaxSubnets returns the IpsecMaxSubnets field value if set, zero value otherwise.
func (o *Topology) GetIpsecMaxSubnets() int32 {
	if o == nil || o.IpsecMaxSubnets == nil {
		var ret int32
		return ret
	}
	return *o.IpsecMaxSubnets
}

// GetIpsecMaxSubnetsOk returns a tuple with the IpsecMaxSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetIpsecMaxSubnetsOk() (*int32, bool) {
	if o == nil || o.IpsecMaxSubnets == nil {
		return nil, false
	}
	return o.IpsecMaxSubnets, true
}

// HasIpsecMaxSubnets returns a boolean if a field has been set.
func (o *Topology) HasIpsecMaxSubnets() bool {
	if o != nil && o.IpsecMaxSubnets != nil {
		return true
	}

	return false
}

// SetIpsecMaxSubnets gets a reference to the given int32 and assigns it to the IpsecMaxSubnets field.
func (o *Topology) SetIpsecMaxSubnets(v int32) {
	o.IpsecMaxSubnets = &v
}

// GetIpsecMaxEndpoints returns the IpsecMaxEndpoints field value if set, zero value otherwise.
func (o *Topology) GetIpsecMaxEndpoints() int32 {
	if o == nil || o.IpsecMaxEndpoints == nil {
		var ret int32
		return ret
	}
	return *o.IpsecMaxEndpoints
}

// GetIpsecMaxEndpointsOk returns a tuple with the IpsecMaxEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetIpsecMaxEndpointsOk() (*int32, bool) {
	if o == nil || o.IpsecMaxEndpoints == nil {
		return nil, false
	}
	return o.IpsecMaxEndpoints, true
}

// HasIpsecMaxEndpoints returns a boolean if a field has been set.
func (o *Topology) HasIpsecMaxEndpoints() bool {
	if o != nil && o.IpsecMaxEndpoints != nil {
		return true
	}

	return false
}

// SetIpsecMaxEndpoints gets a reference to the given int32 and assigns it to the IpsecMaxEndpoints field.
func (o *Topology) SetIpsecMaxEndpoints(v int32) {
	o.IpsecMaxEndpoints = &v
}

// GetLicenseUpgrades returns the LicenseUpgrades field value if set, zero value otherwise.
func (o *Topology) GetLicenseUpgrades() []string {
	if o == nil || o.LicenseUpgrades == nil {
		var ret []string
		return ret
	}
	return o.LicenseUpgrades
}

// GetLicenseUpgradesOk returns a tuple with the LicenseUpgrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetLicenseUpgradesOk() ([]string, bool) {
	if o == nil || o.LicenseUpgrades == nil {
		return nil, false
	}
	return o.LicenseUpgrades, true
}

// HasLicenseUpgrades returns a boolean if a field has been set.
func (o *Topology) HasLicenseUpgrades() bool {
	if o != nil && o.LicenseUpgrades != nil {
		return true
	}

	return false
}

// SetLicenseUpgrades gets a reference to the given []string and assigns it to the LicenseUpgrades field.
func (o *Topology) SetLicenseUpgrades(v []string) {
	o.LicenseUpgrades = v
}

// GetOverlayMaxClients returns the OverlayMaxClients field value if set, zero value otherwise.
func (o *Topology) GetOverlayMaxClients() int32 {
	if o == nil || o.OverlayMaxClients == nil {
		var ret int32
		return ret
	}
	return *o.OverlayMaxClients
}

// GetOverlayMaxClientsOk returns a tuple with the OverlayMaxClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetOverlayMaxClientsOk() (*int32, bool) {
	if o == nil || o.OverlayMaxClients == nil {
		return nil, false
	}
	return o.OverlayMaxClients, true
}

// HasOverlayMaxClients returns a boolean if a field has been set.
func (o *Topology) HasOverlayMaxClients() bool {
	if o != nil && o.OverlayMaxClients != nil {
		return true
	}

	return false
}

// SetOverlayMaxClients gets a reference to the given int32 and assigns it to the OverlayMaxClients field.
func (o *Topology) SetOverlayMaxClients(v int32) {
	o.OverlayMaxClients = &v
}

// GetOverlaySubnet returns the OverlaySubnet field value if set, zero value otherwise.
func (o *Topology) GetOverlaySubnet() string {
	if o == nil || o.OverlaySubnet == nil {
		var ret string
		return ret
	}
	return *o.OverlaySubnet
}

// GetOverlaySubnetOk returns a tuple with the OverlaySubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topology) GetOverlaySubnetOk() (*string, bool) {
	if o == nil || o.OverlaySubnet == nil {
		return nil, false
	}
	return o.OverlaySubnet, true
}

// HasOverlaySubnet returns a boolean if a field has been set.
func (o *Topology) HasOverlaySubnet() bool {
	if o != nil && o.OverlaySubnet != nil {
		return true
	}

	return false
}

// SetOverlaySubnet gets a reference to the given string and assigns it to the OverlaySubnet field.
func (o *Topology) SetOverlaySubnet(v string) {
	o.OverlaySubnet = &v
}

func (o Topology) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	if o.Managers != nil {
		toSerialize["managers"] = o.Managers
	}
	if o.TotalClients != nil {
		toSerialize["total_clients"] = o.TotalClients
	}
	if o.IpsecMaxSubnets != nil {
		toSerialize["ipsec_max_subnets"] = o.IpsecMaxSubnets
	}
	if o.IpsecMaxEndpoints != nil {
		toSerialize["ipsec_max_endpoints"] = o.IpsecMaxEndpoints
	}
	if o.LicenseUpgrades != nil {
		toSerialize["license_upgrades"] = o.LicenseUpgrades
	}
	if o.OverlayMaxClients != nil {
		toSerialize["overlay_max_clients"] = o.OverlayMaxClients
	}
	if o.OverlaySubnet != nil {
		toSerialize["overlay_subnet"] = o.OverlaySubnet
	}
	return json.Marshal(toSerialize)
}

type NullableTopology struct {
	value *Topology
	isSet bool
}

func (v NullableTopology) Get() *Topology {
	return v.value
}

func (v *NullableTopology) Set(val *Topology) {
	v.value = val
	v.isSet = true
}

func (v NullableTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopology(val *Topology) *NullableTopology {
	return &NullableTopology{value: val, isSet: true}
}

func (v NullableTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// UpdateIpsecEndpointRequest struct for UpdateIpsecEndpointRequest
type UpdateIpsecEndpointRequest struct {
	// Name for the endpoint.
	Name *string `json:"name,omitempty"`
	// Description of this IPsec endpoint
	Description *string `json:"description,omitempty"`
	// IP of the remote gateway
	Ipaddress *string `json:"ipaddress,omitempty"`
	// Pre-shared key
	Secret *string `json:"secret,omitempty"`
	// Perfect Forward Secrecy if true, disables if false.
	Pfs *bool `json:"pfs,omitempty"`
	// Version for IKE algorithm
	IkeVersion *int32 `json:"ike_version,omitempty"`
	// True if you want encapsulated IPsec protocol to this gateway
	NatTEnabled *bool `json:"nat_t_enabled,omitempty"`
	// Additional optionals for connection such as 'phase1=aes256_gcm-sha2_256-dh14 phase2=aes256_gcm'
	ExtraConfig *string `json:"extra_config,omitempty"`
	// Internal NAT address of the remote gateway
	PrivateIpaddress *string `json:"private_ipaddress,omitempty"`
	// True if GRE is being used for the specific endpoint
	Gre *bool `json:"gre,omitempty"`
	// Interface address for GRE
	GreInterfaceAddress *string `json:"gre_interface_address,omitempty"`
	// policy, gre, vti
	VpnType *string `json:"vpn_type,omitempty"`
	RouteBasedIntAddress *string `json:"route_based_int_address,omitempty"`
	RouteBasedLocal *string `json:"route_based_local,omitempty"`
	RouteBasedRemote *string `json:"route_based_remote,omitempty"`
}

// NewUpdateIpsecEndpointRequest instantiates a new UpdateIpsecEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpsecEndpointRequest() *UpdateIpsecEndpointRequest {
	this := UpdateIpsecEndpointRequest{}
	return &this
}

// NewUpdateIpsecEndpointRequestWithDefaults instantiates a new UpdateIpsecEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpsecEndpointRequestWithDefaults() *UpdateIpsecEndpointRequest {
	this := UpdateIpsecEndpointRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateIpsecEndpointRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateIpsecEndpointRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *UpdateIpsecEndpointRequest) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UpdateIpsecEndpointRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPfs returns the Pfs field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetPfs() bool {
	if o == nil || o.Pfs == nil {
		var ret bool
		return ret
	}
	return *o.Pfs
}

// GetPfsOk returns a tuple with the Pfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetPfsOk() (*bool, bool) {
	if o == nil || o.Pfs == nil {
		return nil, false
	}
	return o.Pfs, true
}

// HasPfs returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasPfs() bool {
	if o != nil && o.Pfs != nil {
		return true
	}

	return false
}

// SetPfs gets a reference to the given bool and assigns it to the Pfs field.
func (o *UpdateIpsecEndpointRequest) SetPfs(v bool) {
	o.Pfs = &v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetIkeVersion() int32 {
	if o == nil || o.IkeVersion == nil {
		var ret int32
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetIkeVersionOk() (*int32, bool) {
	if o == nil || o.IkeVersion == nil {
		return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasIkeVersion() bool {
	if o != nil && o.IkeVersion != nil {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given int32 and assigns it to the IkeVersion field.
func (o *UpdateIpsecEndpointRequest) SetIkeVersion(v int32) {
	o.IkeVersion = &v
}

// GetNatTEnabled returns the NatTEnabled field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetNatTEnabled() bool {
	if o == nil || o.NatTEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NatTEnabled
}

// GetNatTEnabledOk returns a tuple with the NatTEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetNatTEnabledOk() (*bool, bool) {
	if o == nil || o.NatTEnabled == nil {
		return nil, false
	}
	return o.NatTEnabled, true
}

// HasNatTEnabled returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasNatTEnabled() bool {
	if o != nil && o.NatTEnabled != nil {
		return true
	}

	return false
}

// SetNatTEnabled gets a reference to the given bool and assigns it to the NatTEnabled field.
func (o *UpdateIpsecEndpointRequest) SetNatTEnabled(v bool) {
	o.NatTEnabled = &v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetExtraConfig() string {
	if o == nil || o.ExtraConfig == nil {
		var ret string
		return ret
	}
	return *o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetExtraConfigOk() (*string, bool) {
	if o == nil || o.ExtraConfig == nil {
		return nil, false
	}
	return o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasExtraConfig() bool {
	if o != nil && o.ExtraConfig != nil {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given string and assigns it to the ExtraConfig field.
func (o *UpdateIpsecEndpointRequest) SetExtraConfig(v string) {
	o.ExtraConfig = &v
}

// GetPrivateIpaddress returns the PrivateIpaddress field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetPrivateIpaddress() string {
	if o == nil || o.PrivateIpaddress == nil {
		var ret string
		return ret
	}
	return *o.PrivateIpaddress
}

// GetPrivateIpaddressOk returns a tuple with the PrivateIpaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetPrivateIpaddressOk() (*string, bool) {
	if o == nil || o.PrivateIpaddress == nil {
		return nil, false
	}
	return o.PrivateIpaddress, true
}

// HasPrivateIpaddress returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasPrivateIpaddress() bool {
	if o != nil && o.PrivateIpaddress != nil {
		return true
	}

	return false
}

// SetPrivateIpaddress gets a reference to the given string and assigns it to the PrivateIpaddress field.
func (o *UpdateIpsecEndpointRequest) SetPrivateIpaddress(v string) {
	o.PrivateIpaddress = &v
}

// GetGre returns the Gre field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetGre() bool {
	if o == nil || o.Gre == nil {
		var ret bool
		return ret
	}
	return *o.Gre
}

// GetGreOk returns a tuple with the Gre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetGreOk() (*bool, bool) {
	if o == nil || o.Gre == nil {
		return nil, false
	}
	return o.Gre, true
}

// HasGre returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasGre() bool {
	if o != nil && o.Gre != nil {
		return true
	}

	return false
}

// SetGre gets a reference to the given bool and assigns it to the Gre field.
func (o *UpdateIpsecEndpointRequest) SetGre(v bool) {
	o.Gre = &v
}

// GetGreInterfaceAddress returns the GreInterfaceAddress field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetGreInterfaceAddress() string {
	if o == nil || o.GreInterfaceAddress == nil {
		var ret string
		return ret
	}
	return *o.GreInterfaceAddress
}

// GetGreInterfaceAddressOk returns a tuple with the GreInterfaceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetGreInterfaceAddressOk() (*string, bool) {
	if o == nil || o.GreInterfaceAddress == nil {
		return nil, false
	}
	return o.GreInterfaceAddress, true
}

// HasGreInterfaceAddress returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasGreInterfaceAddress() bool {
	if o != nil && o.GreInterfaceAddress != nil {
		return true
	}

	return false
}

// SetGreInterfaceAddress gets a reference to the given string and assigns it to the GreInterfaceAddress field.
func (o *UpdateIpsecEndpointRequest) SetGreInterfaceAddress(v string) {
	o.GreInterfaceAddress = &v
}

// GetVpnType returns the VpnType field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetVpnType() string {
	if o == nil || o.VpnType == nil {
		var ret string
		return ret
	}
	return *o.VpnType
}

// GetVpnTypeOk returns a tuple with the VpnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetVpnTypeOk() (*string, bool) {
	if o == nil || o.VpnType == nil {
		return nil, false
	}
	return o.VpnType, true
}

// HasVpnType returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasVpnType() bool {
	if o != nil && o.VpnType != nil {
		return true
	}

	return false
}

// SetVpnType gets a reference to the given string and assigns it to the VpnType field.
func (o *UpdateIpsecEndpointRequest) SetVpnType(v string) {
	o.VpnType = &v
}

// GetRouteBasedIntAddress returns the RouteBasedIntAddress field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedIntAddress() string {
	if o == nil || o.RouteBasedIntAddress == nil {
		var ret string
		return ret
	}
	return *o.RouteBasedIntAddress
}

// GetRouteBasedIntAddressOk returns a tuple with the RouteBasedIntAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedIntAddressOk() (*string, bool) {
	if o == nil || o.RouteBasedIntAddress == nil {
		return nil, false
	}
	return o.RouteBasedIntAddress, true
}

// HasRouteBasedIntAddress returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasRouteBasedIntAddress() bool {
	if o != nil && o.RouteBasedIntAddress != nil {
		return true
	}

	return false
}

// SetRouteBasedIntAddress gets a reference to the given string and assigns it to the RouteBasedIntAddress field.
func (o *UpdateIpsecEndpointRequest) SetRouteBasedIntAddress(v string) {
	o.RouteBasedIntAddress = &v
}

// GetRouteBasedLocal returns the RouteBasedLocal field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedLocal() string {
	if o == nil || o.RouteBasedLocal == nil {
		var ret string
		return ret
	}
	return *o.RouteBasedLocal
}

// GetRouteBasedLocalOk returns a tuple with the RouteBasedLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedLocalOk() (*string, bool) {
	if o == nil || o.RouteBasedLocal == nil {
		return nil, false
	}
	return o.RouteBasedLocal, true
}

// HasRouteBasedLocal returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasRouteBasedLocal() bool {
	if o != nil && o.RouteBasedLocal != nil {
		return true
	}

	return false
}

// SetRouteBasedLocal gets a reference to the given string and assigns it to the RouteBasedLocal field.
func (o *UpdateIpsecEndpointRequest) SetRouteBasedLocal(v string) {
	o.RouteBasedLocal = &v
}

// GetRouteBasedRemote returns the RouteBasedRemote field value if set, zero value otherwise.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedRemote() string {
	if o == nil || o.RouteBasedRemote == nil {
		var ret string
		return ret
	}
	return *o.RouteBasedRemote
}

// GetRouteBasedRemoteOk returns a tuple with the RouteBasedRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpsecEndpointRequest) GetRouteBasedRemoteOk() (*string, bool) {
	if o == nil || o.RouteBasedRemote == nil {
		return nil, false
	}
	return o.RouteBasedRemote, true
}

// HasRouteBasedRemote returns a boolean if a field has been set.
func (o *UpdateIpsecEndpointRequest) HasRouteBasedRemote() bool {
	if o != nil && o.RouteBasedRemote != nil {
		return true
	}

	return false
}

// SetRouteBasedRemote gets a reference to the given string and assigns it to the RouteBasedRemote field.
func (o *UpdateIpsecEndpointRequest) SetRouteBasedRemote(v string) {
	o.RouteBasedRemote = &v
}

func (o UpdateIpsecEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Pfs != nil {
		toSerialize["pfs"] = o.Pfs
	}
	if o.IkeVersion != nil {
		toSerialize["ike_version"] = o.IkeVersion
	}
	if o.NatTEnabled != nil {
		toSerialize["nat_t_enabled"] = o.NatTEnabled
	}
	if o.ExtraConfig != nil {
		toSerialize["extra_config"] = o.ExtraConfig
	}
	if o.PrivateIpaddress != nil {
		toSerialize["private_ipaddress"] = o.PrivateIpaddress
	}
	if o.Gre != nil {
		toSerialize["gre"] = o.Gre
	}
	if o.GreInterfaceAddress != nil {
		toSerialize["gre_interface_address"] = o.GreInterfaceAddress
	}
	if o.VpnType != nil {
		toSerialize["vpn_type"] = o.VpnType
	}
	if o.RouteBasedIntAddress != nil {
		toSerialize["route_based_int_address"] = o.RouteBasedIntAddress
	}
	if o.RouteBasedLocal != nil {
		toSerialize["route_based_local"] = o.RouteBasedLocal
	}
	if o.RouteBasedRemote != nil {
		toSerialize["route_based_remote"] = o.RouteBasedRemote
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateIpsecEndpointRequest struct {
	value *UpdateIpsecEndpointRequest
	isSet bool
}

func (v NullableUpdateIpsecEndpointRequest) Get() *UpdateIpsecEndpointRequest {
	return v.value
}

func (v *NullableUpdateIpsecEndpointRequest) Set(val *UpdateIpsecEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpsecEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpsecEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpsecEndpointRequest(val *UpdateIpsecEndpointRequest) *NullableUpdateIpsecEndpointRequest {
	return &NullableUpdateIpsecEndpointRequest{value: val, isSet: true}
}

func (v NullableUpdateIpsecEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpsecEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


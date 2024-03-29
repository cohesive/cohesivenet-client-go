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

// CreateRouteRequest struct for CreateRouteRequest
type CreateRouteRequest struct {
	// CIDR of a route that the VNS3 Controller has access  to that it wants to publish throughout the  Routing tables of the overlay network 
	Cidr string `json:"cidr"`
	Description *string `json:"description,omitempty"`
	// Sets the interface where this route will be advertised.
	Interface *string `json:"interface,omitempty"`
	// If interface is set, a specific gateway address reachable from that interface
	Gateway *string `json:"gateway,omitempty"`
	// Table to create route for
	Table *string `json:"table,omitempty"`
	// numerical reference for the GRE endpoint id (must provide either tunnel OR interface)
	Tunnel *int32 `json:"tunnel,omitempty"`
	// advertise route to overlay network
	Advertise *bool `json:"advertise,omitempty"`
	// weight for route
	Metric *int32 `json:"metric,omitempty"`
}

// NewCreateRouteRequest instantiates a new CreateRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteRequest(cidr string) *CreateRouteRequest {
	this := CreateRouteRequest{}
	this.Cidr = cidr
	var table string = "main"
	this.Table = &table
	return &this
}

// NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteRequestWithDefaults() *CreateRouteRequest {
	this := CreateRouteRequest{}
	var table string = "main"
	this.Table = &table
	return &this
}

// GetCidr returns the Cidr field value
func (o *CreateRouteRequest) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *CreateRouteRequest) SetCidr(v string) {
	o.Cidr = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRouteRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetInterface() string {
	if o == nil || o.Interface == nil {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetInterfaceOk() (*string, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *CreateRouteRequest) SetInterface(v string) {
	o.Interface = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CreateRouteRequest) SetGateway(v string) {
	o.Gateway = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasTable() bool {
	if o != nil && o.Table != nil {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *CreateRouteRequest) SetTable(v string) {
	o.Table = &v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetTunnel() int32 {
	if o == nil || o.Tunnel == nil {
		var ret int32
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetTunnelOk() (*int32, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasTunnel() bool {
	if o != nil && o.Tunnel != nil {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given int32 and assigns it to the Tunnel field.
func (o *CreateRouteRequest) SetTunnel(v int32) {
	o.Tunnel = &v
}

// GetAdvertise returns the Advertise field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetAdvertise() bool {
	if o == nil || o.Advertise == nil {
		var ret bool
		return ret
	}
	return *o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetAdvertiseOk() (*bool, bool) {
	if o == nil || o.Advertise == nil {
		return nil, false
	}
	return o.Advertise, true
}

// HasAdvertise returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasAdvertise() bool {
	if o != nil && o.Advertise != nil {
		return true
	}

	return false
}

// SetAdvertise gets a reference to the given bool and assigns it to the Advertise field.
func (o *CreateRouteRequest) SetAdvertise(v bool) {
	o.Advertise = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetMetric() int32 {
	if o == nil || o.Metric == nil {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetMetricOk() (*int32, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *CreateRouteRequest) SetMetric(v int32) {
	o.Metric = &v
}

func (o CreateRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.Tunnel != nil {
		toSerialize["tunnel"] = o.Tunnel
	}
	if o.Advertise != nil {
		toSerialize["advertise"] = o.Advertise
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRouteRequest struct {
	value *CreateRouteRequest
	isSet bool
}

func (v NullableCreateRouteRequest) Get() *CreateRouteRequest {
	return v.value
}

func (v *NullableCreateRouteRequest) Set(val *CreateRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteRequest(val *CreateRouteRequest) *NullableCreateRouteRequest {
	return &NullableCreateRouteRequest{value: val, isSet: true}
}

func (v NullableCreateRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



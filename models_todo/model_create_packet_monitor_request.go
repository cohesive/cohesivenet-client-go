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

// CreatePacketMonitorRequest struct for CreatePacketMonitorRequest
type CreatePacketMonitorRequest struct {
	// Name of packet monitor. Must be conntrack for type=conntrack as only one conntrack monitor can run at a time
	Name string `json:"name"`
	// conntrack, netflow or pcap
	Type string `json:"type"`
	Interface string `json:"interface"`
	// filter strings are particular to the type of packet monitor. For instance, \"-p 8000\" for pcap. Can be empty string.
	Filter NullableString `json:"filter"`
	// Indicates length of time to run capture for. Can be forever or some string parsable by the Linux date command
	Duration string `json:"duration"`
	// must be file if pcap or conntrack. Otherwise a host should be specified with the prefix \"host\". E.g. \"host:10.0.3.2:4000\"
	Destination string `json:"destination"`
}

// NewCreatePacketMonitorRequest instantiates a new CreatePacketMonitorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePacketMonitorRequest(name string, type_ string, interface_ string, filter NullableString, duration string, destination string) *CreatePacketMonitorRequest {
	this := CreatePacketMonitorRequest{}
	this.Name = name
	this.Type = type_
	this.Interface = interface_
	this.Filter = filter
	this.Duration = duration
	this.Destination = destination
	return &this
}

// NewCreatePacketMonitorRequestWithDefaults instantiates a new CreatePacketMonitorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePacketMonitorRequestWithDefaults() *CreatePacketMonitorRequest {
	this := CreatePacketMonitorRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePacketMonitorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePacketMonitorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePacketMonitorRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreatePacketMonitorRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePacketMonitorRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreatePacketMonitorRequest) SetType(v string) {
	o.Type = v
}

// GetInterface returns the Interface field value
func (o *CreatePacketMonitorRequest) GetInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *CreatePacketMonitorRequest) GetInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *CreatePacketMonitorRequest) SetInterface(v string) {
	o.Interface = v
}

// GetFilter returns the Filter field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreatePacketMonitorRequest) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}

	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePacketMonitorRequest) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// SetFilter sets field value
func (o *CreatePacketMonitorRequest) SetFilter(v string) {
	o.Filter.Set(&v)
}

// GetDuration returns the Duration field value
func (o *CreatePacketMonitorRequest) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CreatePacketMonitorRequest) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CreatePacketMonitorRequest) SetDuration(v string) {
	o.Duration = v
}

// GetDestination returns the Destination field value
func (o *CreatePacketMonitorRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *CreatePacketMonitorRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *CreatePacketMonitorRequest) SetDestination(v string) {
	o.Destination = v
}

func (o CreatePacketMonitorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["interface"] = o.Interface
	}
	if true {
		toSerialize["filter"] = o.Filter.Get()
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePacketMonitorRequest struct {
	value *CreatePacketMonitorRequest
	isSet bool
}

func (v NullableCreatePacketMonitorRequest) Get() *CreatePacketMonitorRequest {
	return v.value
}

func (v *NullableCreatePacketMonitorRequest) Set(val *CreatePacketMonitorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePacketMonitorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePacketMonitorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePacketMonitorRequest(val *CreatePacketMonitorRequest) *NullableCreatePacketMonitorRequest {
	return &NullableCreatePacketMonitorRequest{value: val, isSet: true}
}

func (v NullableCreatePacketMonitorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePacketMonitorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


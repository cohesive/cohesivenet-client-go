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

// CheckoutNextClientpackRequest struct for CheckoutNextClientpackRequest
type CheckoutNextClientpackRequest struct {
	// Set the lower bound for the resulting IP
	LowIp *string `json:"low_ip,omitempty"`
	// Set the upper bound for the resulting IP
	HighIp *string `json:"high_ip,omitempty"`
	// Allows clientpack allocation from the disabled pool, for workflows where all clientpacks are disabled at the start.
	IncludeDisabled *bool `json:"include_disabled,omitempty"`
}

// NewCheckoutNextClientpackRequest instantiates a new CheckoutNextClientpackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutNextClientpackRequest() *CheckoutNextClientpackRequest {
	this := CheckoutNextClientpackRequest{}
	var includeDisabled bool = false
	this.IncludeDisabled = &includeDisabled
	return &this
}

// NewCheckoutNextClientpackRequestWithDefaults instantiates a new CheckoutNextClientpackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutNextClientpackRequestWithDefaults() *CheckoutNextClientpackRequest {
	this := CheckoutNextClientpackRequest{}
	var includeDisabled bool = false
	this.IncludeDisabled = &includeDisabled
	return &this
}

// GetLowIp returns the LowIp field value if set, zero value otherwise.
func (o *CheckoutNextClientpackRequest) GetLowIp() string {
	if o == nil || o.LowIp == nil {
		var ret string
		return ret
	}
	return *o.LowIp
}

// GetLowIpOk returns a tuple with the LowIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNextClientpackRequest) GetLowIpOk() (*string, bool) {
	if o == nil || o.LowIp == nil {
		return nil, false
	}
	return o.LowIp, true
}

// HasLowIp returns a boolean if a field has been set.
func (o *CheckoutNextClientpackRequest) HasLowIp() bool {
	if o != nil && o.LowIp != nil {
		return true
	}

	return false
}

// SetLowIp gets a reference to the given string and assigns it to the LowIp field.
func (o *CheckoutNextClientpackRequest) SetLowIp(v string) {
	o.LowIp = &v
}

// GetHighIp returns the HighIp field value if set, zero value otherwise.
func (o *CheckoutNextClientpackRequest) GetHighIp() string {
	if o == nil || o.HighIp == nil {
		var ret string
		return ret
	}
	return *o.HighIp
}

// GetHighIpOk returns a tuple with the HighIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNextClientpackRequest) GetHighIpOk() (*string, bool) {
	if o == nil || o.HighIp == nil {
		return nil, false
	}
	return o.HighIp, true
}

// HasHighIp returns a boolean if a field has been set.
func (o *CheckoutNextClientpackRequest) HasHighIp() bool {
	if o != nil && o.HighIp != nil {
		return true
	}

	return false
}

// SetHighIp gets a reference to the given string and assigns it to the HighIp field.
func (o *CheckoutNextClientpackRequest) SetHighIp(v string) {
	o.HighIp = &v
}

// GetIncludeDisabled returns the IncludeDisabled field value if set, zero value otherwise.
func (o *CheckoutNextClientpackRequest) GetIncludeDisabled() bool {
	if o == nil || o.IncludeDisabled == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDisabled
}

// GetIncludeDisabledOk returns a tuple with the IncludeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNextClientpackRequest) GetIncludeDisabledOk() (*bool, bool) {
	if o == nil || o.IncludeDisabled == nil {
		return nil, false
	}
	return o.IncludeDisabled, true
}

// HasIncludeDisabled returns a boolean if a field has been set.
func (o *CheckoutNextClientpackRequest) HasIncludeDisabled() bool {
	if o != nil && o.IncludeDisabled != nil {
		return true
	}

	return false
}

// SetIncludeDisabled gets a reference to the given bool and assigns it to the IncludeDisabled field.
func (o *CheckoutNextClientpackRequest) SetIncludeDisabled(v bool) {
	o.IncludeDisabled = &v
}

func (o CheckoutNextClientpackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LowIp != nil {
		toSerialize["low_ip"] = o.LowIp
	}
	if o.HighIp != nil {
		toSerialize["high_ip"] = o.HighIp
	}
	if o.IncludeDisabled != nil {
		toSerialize["include_disabled"] = o.IncludeDisabled
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutNextClientpackRequest struct {
	value *CheckoutNextClientpackRequest
	isSet bool
}

func (v NullableCheckoutNextClientpackRequest) Get() *CheckoutNextClientpackRequest {
	return v.value
}

func (v *NullableCheckoutNextClientpackRequest) Set(val *CheckoutNextClientpackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutNextClientpackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutNextClientpackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutNextClientpackRequest(val *CheckoutNextClientpackRequest) *NullableCheckoutNextClientpackRequest {
	return &NullableCheckoutNextClientpackRequest{value: val, isSet: true}
}

func (v NullableCheckoutNextClientpackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutNextClientpackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


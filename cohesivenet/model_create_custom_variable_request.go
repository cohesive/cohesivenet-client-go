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

// CreateCustomVariableRequest struct for CreateCustomVariableRequest
type CreateCustomVariableRequest struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Description *string `json:"description,omitempty"`
}

// NewCreateCustomVariableRequest instantiates a new CreateCustomVariableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomVariableRequest(name string, value string) *CreateCustomVariableRequest {
	this := CreateCustomVariableRequest{}
	this.Name = name
	this.Value = value
	return &this
}

// NewCreateCustomVariableRequestWithDefaults instantiates a new CreateCustomVariableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomVariableRequestWithDefaults() *CreateCustomVariableRequest {
	this := CreateCustomVariableRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustomVariableRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomVariableRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomVariableRequest) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *CreateCustomVariableRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateCustomVariableRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateCustomVariableRequest) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCustomVariableRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomVariableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCustomVariableRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCustomVariableRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateCustomVariableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCustomVariableRequest struct {
	value *CreateCustomVariableRequest
	isSet bool
}

func (v NullableCreateCustomVariableRequest) Get() *CreateCustomVariableRequest {
	return v.value
}

func (v *NullableCreateCustomVariableRequest) Set(val *CreateCustomVariableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomVariableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomVariableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomVariableRequest(val *CreateCustomVariableRequest) *NullableCreateCustomVariableRequest {
	return &NullableCreateCustomVariableRequest{value: val, isSet: true}
}

func (v NullableCreateCustomVariableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomVariableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



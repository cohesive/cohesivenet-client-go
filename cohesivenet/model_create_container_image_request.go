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

// CreateContainerImageRequest struct for CreateContainerImageRequest
type CreateContainerImageRequest struct {
	// Name of the new plugin instance
	Name string `json:"name"`
	// URL of the image file to be imported
	Url *string `json:"url,omitempty"`
	// TODO: add fields buildfile, imagefile
	// URL of a dockerfile that will be used to build the image
	BuildUrl *string `json:"buildurl,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewCreateContainerImageRequest instantiates a new CreateContainerImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerImageRequest(name string) *CreateContainerImageRequest {
	this := CreateContainerImageRequest{}
	this.Name = name
	return &this
}

// NewCreateContainerImageRequestWithDefaults instantiates a new CreateContainerImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerImageRequestWithDefaults() *CreateContainerImageRequest {
	this := CreateContainerImageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateContainerImageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateContainerImageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateContainerImageRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *CreateContainerImageRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateContainerImageRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *CreateContainerImageRequest) SetUrl(v string) {
	o.Url = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateContainerImageRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerImageRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateContainerImageRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateContainerImageRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBuildUrl returns the BuildUrl field value if set, zero value otherwise.
func (o *CreateContainerImageRequest) GetBuildUrl() string {
	if o == nil || o.BuildUrl == nil {
		var ret string
		return ret
	}
	return *o.BuildUrl
}

// GetBuildUrlOk returns a tuple with the BuildUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerImageRequest) GetBuildUrlOk() (*string, bool) {
	if o == nil || o.BuildUrl == nil {
		return nil, false
	}
	return o.BuildUrl, true
}

// HasBuildUrl returns a boolean if a field has been set.
func (o *CreateContainerImageRequest) HasBuildUrl() bool {
	if o != nil && o.BuildUrl != nil {
		return true
	}

	return false
}

// SetBuildUrl gets a reference to the given string and assigns it to the BuildUrl field.
func (o *CreateContainerImageRequest) SetBuildUrl(v string) {
	o.BuildUrl = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateContainerImageRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerImageRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateContainerImageRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateContainerImageRequest) SetVersion(v string) {
	o.Version = &v
}

func (o CreateContainerImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.BuildUrl != nil {
		toSerialize["buildurl"] = o.BuildUrl
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCreateContainerImageRequest struct {
	value *CreateContainerImageRequest
	isSet bool
}

func (v NullableCreateContainerImageRequest) Get() *CreateContainerImageRequest {
	return v.value
}

func (v *NullableCreateContainerImageRequest) Set(val *CreateContainerImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerImageRequest(val *CreateContainerImageRequest) *NullableCreateContainerImageRequest {
	return &NullableCreateContainerImageRequest{value: val, isSet: true}
}

func (v NullableCreateContainerImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


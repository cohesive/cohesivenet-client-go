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

// CloudInfoGCE Metadata returned from GCE metadata call.
type CloudInfoGCE struct {
	ProjectId *string `json:"projectId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudInfoGCE CloudInfoGCE

// NewCloudInfoGCE instantiates a new CloudInfoGCE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInfoGCE() *CloudInfoGCE {
	this := CloudInfoGCE{}
	return &this
}

// NewCloudInfoGCEWithDefaults instantiates a new CloudInfoGCE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInfoGCEWithDefaults() *CloudInfoGCE {
	this := CloudInfoGCE{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CloudInfoGCE) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInfoGCE) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CloudInfoGCE) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CloudInfoGCE) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o CloudInfoGCE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudInfoGCE) UnmarshalJSON(bytes []byte) (err error) {
	varCloudInfoGCE := _CloudInfoGCE{}

	if err = json.Unmarshal(bytes, &varCloudInfoGCE); err == nil {
		*o = CloudInfoGCE(varCloudInfoGCE)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudInfoGCE struct {
	value *CloudInfoGCE
	isSet bool
}

func (v NullableCloudInfoGCE) Get() *CloudInfoGCE {
	return v.value
}

func (v *NullableCloudInfoGCE) Set(val *CloudInfoGCE) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInfoGCE) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInfoGCE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInfoGCE(val *CloudInfoGCE) *NullableCloudInfoGCE {
	return &NullableCloudInfoGCE{value: val, isSet: true}
}

func (v NullableCloudInfoGCE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInfoGCE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


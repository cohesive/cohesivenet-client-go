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
	"fmt"
)

// PutIdentityControllerSettingsRequest - struct for PutIdentityControllerSettingsRequest
type PutIdentityControllerSettingsRequest struct {
	ERRORUNKNOWN *ERRORUNKNOWN
}

// ERRORUNKNOWNAsPutIdentityControllerSettingsRequest is a convenience function that returns ERRORUNKNOWN wrapped in PutIdentityControllerSettingsRequest
func ERRORUNKNOWNAsPutIdentityControllerSettingsRequest(v *ERRORUNKNOWN) PutIdentityControllerSettingsRequest {
	return PutIdentityControllerSettingsRequest{
		ERRORUNKNOWN: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutIdentityControllerSettingsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ERRORUNKNOWN
	err = newStrictDecoder(data).Decode(&dst.ERRORUNKNOWN)
	if err == nil {
		jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
		if string(jsonERRORUNKNOWN) == "{}" { // empty struct
			dst.ERRORUNKNOWN = nil
		} else {
			match++
		}
	} else {
		dst.ERRORUNKNOWN = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ERRORUNKNOWN = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PutIdentityControllerSettingsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PutIdentityControllerSettingsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutIdentityControllerSettingsRequest) MarshalJSON() ([]byte, error) {
	if src.ERRORUNKNOWN != nil {
		return json.Marshal(&src.ERRORUNKNOWN)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutIdentityControllerSettingsRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ERRORUNKNOWN != nil {
		return obj.ERRORUNKNOWN
	}

	// all schemas are nil
	return nil
}

type NullablePutIdentityControllerSettingsRequest struct {
	value *PutIdentityControllerSettingsRequest
	isSet bool
}

func (v NullablePutIdentityControllerSettingsRequest) Get() *PutIdentityControllerSettingsRequest {
	return v.value
}

func (v *NullablePutIdentityControllerSettingsRequest) Set(val *PutIdentityControllerSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutIdentityControllerSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutIdentityControllerSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutIdentityControllerSettingsRequest(val *PutIdentityControllerSettingsRequest) *NullablePutIdentityControllerSettingsRequest {
	return &NullablePutIdentityControllerSettingsRequest{value: val, isSet: true}
}

func (v NullablePutIdentityControllerSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutIdentityControllerSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


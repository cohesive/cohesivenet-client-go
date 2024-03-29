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

// DeleteAccessURLRequest - struct for DeleteAccessURLRequest
type DeleteAccessURLSearchRequest struct {
	AccessURLId *int32 `json:"access_url_id,omitempty"`
	
	AccessURL *string `json:"access_url,omitempty"`
}

// // interface{}AsDeleteAccessURLRequest is a convenience function that returns interface{} wrapped in DeleteAccessURLRequest
// func Interface{}AsDeleteAccessURLRequest(v *interface{}) DeleteAccessURLRequest {
// 	return DeleteAccessURLRequest{
// 		Interface{}: v,
// 	}
// }

// NewCreateAPITokenRequest instantiates a new CreateAPITokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAccessURLSearchRequest() *DeleteAccessURLSearchRequest {
	this := DeleteAccessURLSearchRequest{}
	return &this
}

func (o DeleteAccessURLSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessURLId != nil {
		toSerialize["access_url_id"] = o.AccessURLId
	}
	if o.AccessURL != nil {
		toSerialize["access_url"] = o.AccessURL
	}
	return json.Marshal(toSerialize)
}


type NullableDeleteAccessURLSearchRequest struct {
	value *DeleteAccessURLSearchRequest
	isSet bool
}

func (v NullableDeleteAccessURLSearchRequest) Get() *DeleteAccessURLSearchRequest {
	return v.value
}

func (v *NullableDeleteAccessURLSearchRequest) Set(val *DeleteAccessURLSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAccessURLSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAccessURLSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAccessURLSearchRequest(val *DeleteAccessURLSearchRequest) *NullableDeleteAccessURLSearchRequest {
	return &NullableDeleteAccessURLSearchRequest{value: val, isSet: true}
}

func (v NullableDeleteAccessURLSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAccessURLSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


// // Unmarshal JSON data into one of the pointers in the struct
// func (dst *DeleteAccessURLRequest) UnmarshalJSON(data []byte) error {
// 	var err error
// 	match := 0
// 	// try to unmarshal data into Interface{}
// 	err = newStrictDecoder(data).Decode(&dst.Interface{})
// 	if err == nil {
// 		jsonInterface{}, _ := json.Marshal(dst.Interface{})
// 		if string(jsonInterface{}) == "{}" { // empty struct
// 			dst.Interface{} = nil
// 		} else {
// 			match++
// 		}
// 	} else {
// 		dst.Interface{} = nil
// 	}

// 	if match > 1 { // more than 1 match
// 		// reset to nil
// 		dst.Interface{} = nil

// 		return fmt.Errorf("Data matches more than one schema in oneOf(DeleteAccessURLRequest)")
// 	} else if match == 1 {
// 		return nil // exactly one match
// 	} else { // no match
// 		return fmt.Errorf("Data failed to match schemas in oneOf(DeleteAccessURLRequest)")
// 	}
// }

// // Marshal data from the first non-nil pointers in the struct to JSON
// func (src DeleteAccessURLRequest) MarshalJSON() ([]byte, error) {
// 	if src.Interface{} != nil {
// 		return json.Marshal(&src.Interface{})
// 	}

// 	return nil, nil // no data in oneOf schemas
// }

// // Get the actual instance
// func (obj *DeleteAccessURLRequest) GetActualInstance() (interface{}) {
// 	if obj == nil {
// 		return nil
// 	}
// 	if obj.Interface{} != nil {
// 		return obj.Interface{}
// 	}

// 	// all schemas are nil
// 	return nil
// }


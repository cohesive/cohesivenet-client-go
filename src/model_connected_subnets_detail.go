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

// ConnectedSubnetsDetail - struct for ConnectedSubnetsDetail
type ConnectedSubnetsDetail struct {
	ArrayOfConnectedSubnet *[]ConnectedSubnet
	ArrayOfArrayOfString *[][]string
}

// []ConnectedSubnetAsConnectedSubnetsDetail is a convenience function that returns []ConnectedSubnet wrapped in ConnectedSubnetsDetail
func ArrayOfConnectedSubnetAsConnectedSubnetsDetail(v *[]ConnectedSubnet) ConnectedSubnetsDetail {
	return ConnectedSubnetsDetail{
		ArrayOfConnectedSubnet: v,
	}
}

// [][]stringAsConnectedSubnetsDetail is a convenience function that returns [][]string wrapped in ConnectedSubnetsDetail
func ArrayOfArrayOfStringAsConnectedSubnetsDetail(v *[][]string) ConnectedSubnetsDetail {
	return ConnectedSubnetsDetail{
		ArrayOfArrayOfString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConnectedSubnetsDetail) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfConnectedSubnet
	err = newStrictDecoder(data).Decode(&dst.ArrayOfConnectedSubnet)
	if err == nil {
		jsonArrayOfConnectedSubnet, _ := json.Marshal(dst.ArrayOfConnectedSubnet)
		if string(jsonArrayOfConnectedSubnet) == "{}" { // empty struct
			dst.ArrayOfConnectedSubnet = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfConnectedSubnet = nil
	}

	// try to unmarshal data into ArrayOfArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfArrayOfString)
	if err == nil {
		jsonArrayOfArrayOfString, _ := json.Marshal(dst.ArrayOfArrayOfString)
		if string(jsonArrayOfArrayOfString) == "{}" { // empty struct
			dst.ArrayOfArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfArrayOfString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfConnectedSubnet = nil
		dst.ArrayOfArrayOfString = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ConnectedSubnetsDetail)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ConnectedSubnetsDetail)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConnectedSubnetsDetail) MarshalJSON() ([]byte, error) {
	if src.ArrayOfConnectedSubnet != nil {
		return json.Marshal(&src.ArrayOfConnectedSubnet)
	}

	if src.ArrayOfArrayOfString != nil {
		return json.Marshal(&src.ArrayOfArrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConnectedSubnetsDetail) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfConnectedSubnet != nil {
		return obj.ArrayOfConnectedSubnet
	}

	if obj.ArrayOfArrayOfString != nil {
		return obj.ArrayOfArrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableConnectedSubnetsDetail struct {
	value *ConnectedSubnetsDetail
	isSet bool
}

func (v NullableConnectedSubnetsDetail) Get() *ConnectedSubnetsDetail {
	return v.value
}

func (v *NullableConnectedSubnetsDetail) Set(val *ConnectedSubnetsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedSubnetsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedSubnetsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedSubnetsDetail(val *ConnectedSubnetsDetail) *NullableConnectedSubnetsDetail {
	return &NullableConnectedSubnetsDetail{value: val, isSet: true}
}

func (v NullableConnectedSubnetsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedSubnetsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


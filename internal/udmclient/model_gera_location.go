/*
Nudm_EE

Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package udmclient

import (
	"encoding/json"
	"time"
	"fmt"
)

// GeraLocation - Exactly one of cgi, sai or lai shall be present.
type GeraLocation struct {
	Interface{} *interface{}
}

// interface{}AsGeraLocation is a convenience function that returns interface{} wrapped in GeraLocation
func Interface{}AsGeraLocation(v *interface{}) GeraLocation {
	return GeraLocation{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GeraLocation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GeraLocation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GeraLocation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GeraLocation) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GeraLocation) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableGeraLocation struct {
	value *GeraLocation
	isSet bool
}

func (v NullableGeraLocation) Get() *GeraLocation {
	return v.value
}

func (v *NullableGeraLocation) Set(val *GeraLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeraLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGeraLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeraLocation(val *GeraLocation) *NullableGeraLocation {
	return &NullableGeraLocation{value: val, isSet: true}
}

func (v NullableGeraLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeraLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nudm_EE

Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package udmclient

import (
	"encoding/json"
	"fmt"
)

// UeReachabilityAnyOf the model 'UeReachabilityAnyOf'
type UeReachabilityAnyOf string

// List of UeReachability_anyOf
const (
	UNREACHABLE UeReachabilityAnyOf = "UNREACHABLE"
	REACHABLE UeReachabilityAnyOf = "REACHABLE"
	REGULATORY_ONLY UeReachabilityAnyOf = "REGULATORY_ONLY"
)

// All allowed values of UeReachabilityAnyOf enum
var AllowedUeReachabilityAnyOfEnumValues = []UeReachabilityAnyOf{
	"UNREACHABLE",
	"REACHABLE",
	"REGULATORY_ONLY",
}

func (v *UeReachabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UeReachabilityAnyOf(value)
	for _, existing := range AllowedUeReachabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UeReachabilityAnyOf", value)
}

// NewUeReachabilityAnyOfFromValue returns a pointer to a valid UeReachabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUeReachabilityAnyOfFromValue(v string) (*UeReachabilityAnyOf, error) {
	ev := UeReachabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UeReachabilityAnyOf: valid values are %v", v, AllowedUeReachabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UeReachabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedUeReachabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UeReachability_anyOf value
func (v UeReachabilityAnyOf) Ptr() *UeReachabilityAnyOf {
	return &v
}

type NullableUeReachabilityAnyOf struct {
	value *UeReachabilityAnyOf
	isSet bool
}

func (v NullableUeReachabilityAnyOf) Get() *UeReachabilityAnyOf {
	return v.value
}

func (v *NullableUeReachabilityAnyOf) Set(val *UeReachabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachabilityAnyOf(val *UeReachabilityAnyOf) *NullableUeReachabilityAnyOf {
	return &NullableUeReachabilityAnyOf{value: val, isSet: true}
}

func (v NullableUeReachabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


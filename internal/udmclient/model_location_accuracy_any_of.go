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

// LocationAccuracyAnyOf the model 'LocationAccuracyAnyOf'
type LocationAccuracyAnyOf string

// List of LocationAccuracy_anyOf
const (
	CELL_LEVEL LocationAccuracyAnyOf = "CELL_LEVEL"
	RAN_NODE_LEVEL LocationAccuracyAnyOf = "RAN_NODE_LEVEL"
	TA_LEVEL LocationAccuracyAnyOf = "TA_LEVEL"
	N3_IWF_LEVEL LocationAccuracyAnyOf = "N3IWF_LEVEL"
	UE_IP LocationAccuracyAnyOf = "UE_IP"
	UE_PORT LocationAccuracyAnyOf = "UE_PORT"
)

// All allowed values of LocationAccuracyAnyOf enum
var AllowedLocationAccuracyAnyOfEnumValues = []LocationAccuracyAnyOf{
	"CELL_LEVEL",
	"RAN_NODE_LEVEL",
	"TA_LEVEL",
	"N3IWF_LEVEL",
	"UE_IP",
	"UE_PORT",
}

func (v *LocationAccuracyAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationAccuracyAnyOf(value)
	for _, existing := range AllowedLocationAccuracyAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationAccuracyAnyOf", value)
}

// NewLocationAccuracyAnyOfFromValue returns a pointer to a valid LocationAccuracyAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationAccuracyAnyOfFromValue(v string) (*LocationAccuracyAnyOf, error) {
	ev := LocationAccuracyAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationAccuracyAnyOf: valid values are %v", v, AllowedLocationAccuracyAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationAccuracyAnyOf) IsValid() bool {
	for _, existing := range AllowedLocationAccuracyAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationAccuracy_anyOf value
func (v LocationAccuracyAnyOf) Ptr() *LocationAccuracyAnyOf {
	return &v
}

type NullableLocationAccuracyAnyOf struct {
	value *LocationAccuracyAnyOf
	isSet bool
}

func (v NullableLocationAccuracyAnyOf) Get() *LocationAccuracyAnyOf {
	return v.value
}

func (v *NullableLocationAccuracyAnyOf) Set(val *LocationAccuracyAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAccuracyAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAccuracyAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAccuracyAnyOf(val *LocationAccuracyAnyOf) *NullableLocationAccuracyAnyOf {
	return &NullableLocationAccuracyAnyOf{value: val, isSet: true}
}

func (v NullableLocationAccuracyAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAccuracyAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


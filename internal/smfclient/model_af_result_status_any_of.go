/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
	"fmt"
)

// AfResultStatusAnyOf the model 'AfResultStatusAnyOf'
type AfResultStatusAnyOf string

// List of AfResultStatus_anyOf
const (
	AFRESULTSTATUSANYOF_SUCCESS AfResultStatusAnyOf = "SUCCESS"
	AFRESULTSTATUSANYOF_TEMPORARY_CONGESTION AfResultStatusAnyOf = "TEMPORARY_CONGESTION"
	AFRESULTSTATUSANYOF_RELOC_NO_ALLOWED AfResultStatusAnyOf = "RELOC_NO_ALLOWED"
	AFRESULTSTATUSANYOF_OTHER AfResultStatusAnyOf = "OTHER"
)

// All allowed values of AfResultStatusAnyOf enum
var AllowedAfResultStatusAnyOfEnumValues = []AfResultStatusAnyOf{
	"SUCCESS",
	"TEMPORARY_CONGESTION",
	"RELOC_NO_ALLOWED",
	"OTHER",
}

func (v *AfResultStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AfResultStatusAnyOf(value)
	for _, existing := range AllowedAfResultStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AfResultStatusAnyOf", value)
}

// NewAfResultStatusAnyOfFromValue returns a pointer to a valid AfResultStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAfResultStatusAnyOfFromValue(v string) (*AfResultStatusAnyOf, error) {
	ev := AfResultStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AfResultStatusAnyOf: valid values are %v", v, AllowedAfResultStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AfResultStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedAfResultStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AfResultStatus_anyOf value
func (v AfResultStatusAnyOf) Ptr() *AfResultStatusAnyOf {
	return &v
}

type NullableAfResultStatusAnyOf struct {
	value *AfResultStatusAnyOf
	isSet bool
}

func (v NullableAfResultStatusAnyOf) Get() *AfResultStatusAnyOf {
	return v.value
}

func (v *NullableAfResultStatusAnyOf) Set(val *AfResultStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAfResultStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAfResultStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfResultStatusAnyOf(val *AfResultStatusAnyOf) *NullableAfResultStatusAnyOf {
	return &NullableAfResultStatusAnyOf{value: val, isSet: true}
}

func (v NullableAfResultStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfResultStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// ReachabilityForDataReportConfigAnyOf the model 'ReachabilityForDataReportConfigAnyOf'
type ReachabilityForDataReportConfigAnyOf string

// List of ReachabilityForDataReportConfig_anyOf
const (
	DIRECT_REPORT ReachabilityForDataReportConfigAnyOf = "DIRECT_REPORT"
	INDIRECT_REPORT ReachabilityForDataReportConfigAnyOf = "INDIRECT_REPORT"
)

// All allowed values of ReachabilityForDataReportConfigAnyOf enum
var AllowedReachabilityForDataReportConfigAnyOfEnumValues = []ReachabilityForDataReportConfigAnyOf{
	"DIRECT_REPORT",
	"INDIRECT_REPORT",
}

func (v *ReachabilityForDataReportConfigAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReachabilityForDataReportConfigAnyOf(value)
	for _, existing := range AllowedReachabilityForDataReportConfigAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReachabilityForDataReportConfigAnyOf", value)
}

// NewReachabilityForDataReportConfigAnyOfFromValue returns a pointer to a valid ReachabilityForDataReportConfigAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReachabilityForDataReportConfigAnyOfFromValue(v string) (*ReachabilityForDataReportConfigAnyOf, error) {
	ev := ReachabilityForDataReportConfigAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReachabilityForDataReportConfigAnyOf: valid values are %v", v, AllowedReachabilityForDataReportConfigAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReachabilityForDataReportConfigAnyOf) IsValid() bool {
	for _, existing := range AllowedReachabilityForDataReportConfigAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReachabilityForDataReportConfig_anyOf value
func (v ReachabilityForDataReportConfigAnyOf) Ptr() *ReachabilityForDataReportConfigAnyOf {
	return &v
}

type NullableReachabilityForDataReportConfigAnyOf struct {
	value *ReachabilityForDataReportConfigAnyOf
	isSet bool
}

func (v NullableReachabilityForDataReportConfigAnyOf) Get() *ReachabilityForDataReportConfigAnyOf {
	return v.value
}

func (v *NullableReachabilityForDataReportConfigAnyOf) Set(val *ReachabilityForDataReportConfigAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForDataReportConfigAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForDataReportConfigAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForDataReportConfigAnyOf(val *ReachabilityForDataReportConfigAnyOf) *NullableReachabilityForDataReportConfigAnyOf {
	return &NullableReachabilityForDataReportConfigAnyOf{value: val, isSet: true}
}

func (v NullableReachabilityForDataReportConfigAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForDataReportConfigAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// PduSessionType PduSessionType indicates the type of a PDU session. It shall comply with the provisions defined in table 5.4.3.3-1.  
type PduSessionType struct {
	PduSessionTypeAnyOf *PduSessionTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PduSessionTypeAnyOf
	err = json.Unmarshal(data, &dst.PduSessionTypeAnyOf);
	if err == nil {
		jsonPduSessionTypeAnyOf, _ := json.Marshal(dst.PduSessionTypeAnyOf)
		if string(jsonPduSessionTypeAnyOf) == "{}" { // empty struct
			dst.PduSessionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.PduSessionTypeAnyOf, return on the first match
		}
	} else {
		dst.PduSessionTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(PduSessionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionType) MarshalJSON() ([]byte, error) {
	if src.PduSessionTypeAnyOf != nil {
		return json.Marshal(&src.PduSessionTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionType struct {
	value *PduSessionType
	isSet bool
}

func (v NullablePduSessionType) Get() *PduSessionType {
	return v.value
}

func (v *NullablePduSessionType) Set(val *PduSessionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionType(val *PduSessionType) *NullablePduSessionType {
	return &NullablePduSessionType{value: val, isSet: true}
}

func (v NullablePduSessionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


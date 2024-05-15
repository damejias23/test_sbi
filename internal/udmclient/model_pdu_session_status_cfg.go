/*
Nudm_EE

Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package udmclient

import (
	"encoding/json"
)

// PduSessionStatusCfg struct for PduSessionStatusCfg
type PduSessionStatusCfg struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
}

// NewPduSessionStatusCfg instantiates a new PduSessionStatusCfg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionStatusCfg() *PduSessionStatusCfg {
	this := PduSessionStatusCfg{}
	return &this
}

// NewPduSessionStatusCfgWithDefaults instantiates a new PduSessionStatusCfg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionStatusCfgWithDefaults() *PduSessionStatusCfg {
	this := PduSessionStatusCfg{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PduSessionStatusCfg) GetDnn() string {
	if o == nil || o.Dnn == nil {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionStatusCfg) GetDnnOk() (*string, bool) {
	if o == nil || o.Dnn == nil {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PduSessionStatusCfg) HasDnn() bool {
	if o != nil && o.Dnn != nil {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PduSessionStatusCfg) SetDnn(v string) {
	o.Dnn = &v
}

func (o PduSessionStatusCfg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dnn != nil {
		toSerialize["dnn"] = o.Dnn
	}
	return json.Marshal(toSerialize)
}

type NullablePduSessionStatusCfg struct {
	value *PduSessionStatusCfg
	isSet bool
}

func (v NullablePduSessionStatusCfg) Get() *PduSessionStatusCfg {
	return v.value
}

func (v *NullablePduSessionStatusCfg) Set(val *PduSessionStatusCfg) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionStatusCfg) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionStatusCfg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionStatusCfg(val *PduSessionStatusCfg) *NullablePduSessionStatusCfg {
	return &NullablePduSessionStatusCfg{value: val, isSet: true}
}

func (v NullablePduSessionStatusCfg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionStatusCfg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



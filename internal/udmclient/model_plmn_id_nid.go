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

// PlmnIdNid Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN. 
type PlmnIdNid struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewPlmnIdNid instantiates a new PlmnIdNid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnIdNid(mcc string, mnc string) *PlmnIdNid {
	this := PlmnIdNid{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnIdNidWithDefaults instantiates a new PlmnIdNid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdNidWithDefaults() *PlmnIdNid {
	this := PlmnIdNid{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnIdNid) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnIdNid) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnIdNid) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnIdNid) SetMnc(v string) {
	o.Mnc = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlmnIdNid) GetNid() string {
	if o == nil || o.Nid == nil {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetNidOk() (*string, bool) {
	if o == nil || o.Nid == nil {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlmnIdNid) HasNid() bool {
	if o != nil && o.Nid != nil {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlmnIdNid) SetNid(v string) {
	o.Nid = &v
}

func (o PlmnIdNid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mcc"] = o.Mcc
	}
	if true {
		toSerialize["mnc"] = o.Mnc
	}
	if o.Nid != nil {
		toSerialize["nid"] = o.Nid
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnIdNid struct {
	value *PlmnIdNid
	isSet bool
}

func (v NullablePlmnIdNid) Get() *PlmnIdNid {
	return v.value
}

func (v *NullablePlmnIdNid) Set(val *PlmnIdNid) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnIdNid) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnIdNid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnIdNid(val *PlmnIdNid) *NullablePlmnIdNid {
	return &NullablePlmnIdNid{value: val, isSet: true}
}

func (v NullablePlmnIdNid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnIdNid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// RoutingAreaId Contains a Routing Area Identification as defined in 3GPP TS 23.003, clause 4.2.
type RoutingAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	// Location Area Code
	Lac string `json:"lac"`
	// Routing Area Code
	Rac string `json:"rac"`
}

// NewRoutingAreaId instantiates a new RoutingAreaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAreaId(plmnId PlmnId, lac string, rac string) *RoutingAreaId {
	this := RoutingAreaId{}
	this.PlmnId = plmnId
	this.Lac = lac
	this.Rac = rac
	return &this
}

// NewRoutingAreaIdWithDefaults instantiates a new RoutingAreaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAreaIdWithDefaults() *RoutingAreaId {
	this := RoutingAreaId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *RoutingAreaId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *RoutingAreaId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *RoutingAreaId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetLacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *RoutingAreaId) SetLac(v string) {
	o.Lac = v
}

// GetRac returns the Rac field value
func (o *RoutingAreaId) GetRac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rac
}

// GetRacOk returns a tuple with the Rac field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetRacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rac, true
}

// SetRac sets field value
func (o *RoutingAreaId) SetRac(v string) {
	o.Rac = v
}

func (o RoutingAreaId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["lac"] = o.Lac
	}
	if true {
		toSerialize["rac"] = o.Rac
	}
	return json.Marshal(toSerialize)
}

type NullableRoutingAreaId struct {
	value *RoutingAreaId
	isSet bool
}

func (v NullableRoutingAreaId) Get() *RoutingAreaId {
	return v.value
}

func (v *NullableRoutingAreaId) Set(val *RoutingAreaId) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAreaId) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAreaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAreaId(val *RoutingAreaId) *NullableRoutingAreaId {
	return &NullableRoutingAreaId{value: val, isSet: true}
}

func (v NullableRoutingAreaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAreaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



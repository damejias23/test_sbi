/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
)

// EasIpReplacementInfo Contains EAS IP replacement information for a Source and a Target EAS.
type EasIpReplacementInfo struct {
	Source EasServerAddress `json:"source"`
	Target EasServerAddress `json:"target"`
}

// NewEasIpReplacementInfo instantiates a new EasIpReplacementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasIpReplacementInfo(source EasServerAddress, target EasServerAddress) *EasIpReplacementInfo {
	this := EasIpReplacementInfo{}
	this.Source = source
	this.Target = target
	return &this
}

// NewEasIpReplacementInfoWithDefaults instantiates a new EasIpReplacementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasIpReplacementInfoWithDefaults() *EasIpReplacementInfo {
	this := EasIpReplacementInfo{}
	return &this
}

// GetSource returns the Source field value
func (o *EasIpReplacementInfo) GetSource() EasServerAddress {
	if o == nil {
		var ret EasServerAddress
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EasIpReplacementInfo) GetSourceOk() (*EasServerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EasIpReplacementInfo) SetSource(v EasServerAddress) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *EasIpReplacementInfo) GetTarget() EasServerAddress {
	if o == nil {
		var ret EasServerAddress
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EasIpReplacementInfo) GetTargetOk() (*EasServerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EasIpReplacementInfo) SetTarget(v EasServerAddress) {
	o.Target = v
}

func (o EasIpReplacementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableEasIpReplacementInfo struct {
	value *EasIpReplacementInfo
	isSet bool
}

func (v NullableEasIpReplacementInfo) Get() *EasIpReplacementInfo {
	return v.value
}

func (v *NullableEasIpReplacementInfo) Set(val *EasIpReplacementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEasIpReplacementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEasIpReplacementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasIpReplacementInfo(val *EasIpReplacementInfo) *NullableEasIpReplacementInfo {
	return &NullableEasIpReplacementInfo{value: val, isSet: true}
}

func (v NullableEasIpReplacementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasIpReplacementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



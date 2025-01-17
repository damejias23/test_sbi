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

// IdentityRange A range of GPSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type IdentityRange struct {
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}

// NewIdentityRange instantiates a new IdentityRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRange() *IdentityRange {
	this := IdentityRange{}
	return &this
}

// NewIdentityRangeWithDefaults instantiates a new IdentityRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRangeWithDefaults() *IdentityRange {
	this := IdentityRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *IdentityRange) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *IdentityRange) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *IdentityRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *IdentityRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *IdentityRange) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *IdentityRange) SetEnd(v string) {
	o.End = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *IdentityRange) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *IdentityRange) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *IdentityRange) SetPattern(v string) {
	o.Pattern = &v
}

func (o IdentityRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRange struct {
	value *IdentityRange
	isSet bool
}

func (v NullableIdentityRange) Get() *IdentityRange {
	return v.value
}

func (v *NullableIdentityRange) Set(val *IdentityRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRange(val *IdentityRange) *NullableIdentityRange {
	return &NullableIdentityRange{value: val, isSet: true}
}

func (v NullableIdentityRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



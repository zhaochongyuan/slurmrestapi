/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.11.5&openapi/slurmctld&openapi/slurmdbd
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0042QosLimitsMaxWallClock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042QosLimitsMaxWallClock{}

// V0042QosLimitsMaxWallClock struct for V0042QosLimitsMaxWallClock
type V0042QosLimitsMaxWallClock struct {
	Per *V0042QosLimitsMaxWallClockPer `json:"per,omitempty"`
}

// NewV0042QosLimitsMaxWallClock instantiates a new V0042QosLimitsMaxWallClock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042QosLimitsMaxWallClock() *V0042QosLimitsMaxWallClock {
	this := V0042QosLimitsMaxWallClock{}
	return &this
}

// NewV0042QosLimitsMaxWallClockWithDefaults instantiates a new V0042QosLimitsMaxWallClock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042QosLimitsMaxWallClockWithDefaults() *V0042QosLimitsMaxWallClock {
	this := V0042QosLimitsMaxWallClock{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0042QosLimitsMaxWallClock) GetPer() V0042QosLimitsMaxWallClockPer {
	if o == nil || IsNil(o.Per) {
		var ret V0042QosLimitsMaxWallClockPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042QosLimitsMaxWallClock) GetPerOk() (*V0042QosLimitsMaxWallClockPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0042QosLimitsMaxWallClock) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0042QosLimitsMaxWallClockPer and assigns it to the Per field.
func (o *V0042QosLimitsMaxWallClock) SetPer(v V0042QosLimitsMaxWallClockPer) {
	o.Per = &v
}

func (o V0042QosLimitsMaxWallClock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042QosLimitsMaxWallClock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0042QosLimitsMaxWallClock struct {
	value *V0042QosLimitsMaxWallClock
	isSet bool
}

func (v NullableV0042QosLimitsMaxWallClock) Get() *V0042QosLimitsMaxWallClock {
	return v.value
}

func (v *NullableV0042QosLimitsMaxWallClock) Set(val *V0042QosLimitsMaxWallClock) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042QosLimitsMaxWallClock) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042QosLimitsMaxWallClock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042QosLimitsMaxWallClock(val *V0042QosLimitsMaxWallClock) *NullableV0042QosLimitsMaxWallClock {
	return &NullableV0042QosLimitsMaxWallClock{value: val, isSet: true}
}

func (v NullableV0042QosLimitsMaxWallClock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042QosLimitsMaxWallClock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



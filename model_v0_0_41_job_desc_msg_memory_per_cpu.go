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

// checks if the V0041JobDescMsgMemoryPerCpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobDescMsgMemoryPerCpu{}

// V0041JobDescMsgMemoryPerCpu Minimum memory in megabytes per allocated CPU
type V0041JobDescMsgMemoryPerCpu struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int64 `json:"number,omitempty"`
}

// NewV0041JobDescMsgMemoryPerCpu instantiates a new V0041JobDescMsgMemoryPerCpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobDescMsgMemoryPerCpu() *V0041JobDescMsgMemoryPerCpu {
	this := V0041JobDescMsgMemoryPerCpu{}
	return &this
}

// NewV0041JobDescMsgMemoryPerCpuWithDefaults instantiates a new V0041JobDescMsgMemoryPerCpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobDescMsgMemoryPerCpuWithDefaults() *V0041JobDescMsgMemoryPerCpu {
	this := V0041JobDescMsgMemoryPerCpu{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *V0041JobDescMsgMemoryPerCpu) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgMemoryPerCpu) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *V0041JobDescMsgMemoryPerCpu) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *V0041JobDescMsgMemoryPerCpu) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *V0041JobDescMsgMemoryPerCpu) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgMemoryPerCpu) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *V0041JobDescMsgMemoryPerCpu) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *V0041JobDescMsgMemoryPerCpu) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *V0041JobDescMsgMemoryPerCpu) GetNumber() int64 {
	if o == nil || IsNil(o.Number) {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgMemoryPerCpu) GetNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *V0041JobDescMsgMemoryPerCpu) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *V0041JobDescMsgMemoryPerCpu) SetNumber(v int64) {
	o.Number = &v
}

func (o V0041JobDescMsgMemoryPerCpu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobDescMsgMemoryPerCpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.Infinite) {
		toSerialize["infinite"] = o.Infinite
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableV0041JobDescMsgMemoryPerCpu struct {
	value *V0041JobDescMsgMemoryPerCpu
	isSet bool
}

func (v NullableV0041JobDescMsgMemoryPerCpu) Get() *V0041JobDescMsgMemoryPerCpu {
	return v.value
}

func (v *NullableV0041JobDescMsgMemoryPerCpu) Set(val *V0041JobDescMsgMemoryPerCpu) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobDescMsgMemoryPerCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobDescMsgMemoryPerCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobDescMsgMemoryPerCpu(val *V0041JobDescMsgMemoryPerCpu) *NullableV0041JobDescMsgMemoryPerCpu {
	return &NullableV0041JobDescMsgMemoryPerCpu{value: val, isSet: true}
}

func (v NullableV0041JobDescMsgMemoryPerCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobDescMsgMemoryPerCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the V0041UpdateNodeMsgResumeAfter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041UpdateNodeMsgResumeAfter{}

// V0041UpdateNodeMsgResumeAfter Number of seconds after which to automatically resume DOWN or DRAINED node
type V0041UpdateNodeMsgResumeAfter struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int32 `json:"number,omitempty"`
}

// NewV0041UpdateNodeMsgResumeAfter instantiates a new V0041UpdateNodeMsgResumeAfter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041UpdateNodeMsgResumeAfter() *V0041UpdateNodeMsgResumeAfter {
	this := V0041UpdateNodeMsgResumeAfter{}
	return &this
}

// NewV0041UpdateNodeMsgResumeAfterWithDefaults instantiates a new V0041UpdateNodeMsgResumeAfter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041UpdateNodeMsgResumeAfterWithDefaults() *V0041UpdateNodeMsgResumeAfter {
	this := V0041UpdateNodeMsgResumeAfter{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *V0041UpdateNodeMsgResumeAfter) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041UpdateNodeMsgResumeAfter) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *V0041UpdateNodeMsgResumeAfter) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *V0041UpdateNodeMsgResumeAfter) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *V0041UpdateNodeMsgResumeAfter) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041UpdateNodeMsgResumeAfter) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *V0041UpdateNodeMsgResumeAfter) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *V0041UpdateNodeMsgResumeAfter) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *V0041UpdateNodeMsgResumeAfter) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041UpdateNodeMsgResumeAfter) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *V0041UpdateNodeMsgResumeAfter) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *V0041UpdateNodeMsgResumeAfter) SetNumber(v int32) {
	o.Number = &v
}

func (o V0041UpdateNodeMsgResumeAfter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041UpdateNodeMsgResumeAfter) ToMap() (map[string]interface{}, error) {
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

type NullableV0041UpdateNodeMsgResumeAfter struct {
	value *V0041UpdateNodeMsgResumeAfter
	isSet bool
}

func (v NullableV0041UpdateNodeMsgResumeAfter) Get() *V0041UpdateNodeMsgResumeAfter {
	return v.value
}

func (v *NullableV0041UpdateNodeMsgResumeAfter) Set(val *V0041UpdateNodeMsgResumeAfter) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041UpdateNodeMsgResumeAfter) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041UpdateNodeMsgResumeAfter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041UpdateNodeMsgResumeAfter(val *V0041UpdateNodeMsgResumeAfter) *NullableV0041UpdateNodeMsgResumeAfter {
	return &NullableV0041UpdateNodeMsgResumeAfter{value: val, isSet: true}
}

func (v NullableV0041UpdateNodeMsgResumeAfter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041UpdateNodeMsgResumeAfter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



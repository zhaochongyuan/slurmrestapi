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

// checks if the V0040JobTimeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040JobTimeUser{}

// V0040JobTimeUser struct for V0040JobTimeUser
type V0040JobTimeUser struct {
	// User CPU time used by the job in seconds
	Seconds *int64 `json:"seconds,omitempty"`
	// User CPU time used by the job in microseconds
	Microseconds *int64 `json:"microseconds,omitempty"`
}

// NewV0040JobTimeUser instantiates a new V0040JobTimeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040JobTimeUser() *V0040JobTimeUser {
	this := V0040JobTimeUser{}
	return &this
}

// NewV0040JobTimeUserWithDefaults instantiates a new V0040JobTimeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040JobTimeUserWithDefaults() *V0040JobTimeUser {
	this := V0040JobTimeUser{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0040JobTimeUser) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTimeUser) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0040JobTimeUser) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0040JobTimeUser) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0040JobTimeUser) GetMicroseconds() int64 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int64
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040JobTimeUser) GetMicrosecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0040JobTimeUser) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int64 and assigns it to the Microseconds field.
func (o *V0040JobTimeUser) SetMicroseconds(v int64) {
	o.Microseconds = &v
}

func (o V0040JobTimeUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040JobTimeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0040JobTimeUser struct {
	value *V0040JobTimeUser
	isSet bool
}

func (v NullableV0040JobTimeUser) Get() *V0040JobTimeUser {
	return v.value
}

func (v *NullableV0040JobTimeUser) Set(val *V0040JobTimeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040JobTimeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040JobTimeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040JobTimeUser(val *V0040JobTimeUser) *NullableV0040JobTimeUser {
	return &NullableV0040JobTimeUser{value: val, isSet: true}
}

func (v NullableV0040JobTimeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040JobTimeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



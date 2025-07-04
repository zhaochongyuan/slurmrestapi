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

// checks if the V0040ReservationInfoPurgeCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ReservationInfoPurgeCompleted{}

// V0040ReservationInfoPurgeCompleted struct for V0040ReservationInfoPurgeCompleted
type V0040ReservationInfoPurgeCompleted struct {
	Time *V0040Uint32NoVal `json:"time,omitempty"`
}

// NewV0040ReservationInfoPurgeCompleted instantiates a new V0040ReservationInfoPurgeCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ReservationInfoPurgeCompleted() *V0040ReservationInfoPurgeCompleted {
	this := V0040ReservationInfoPurgeCompleted{}
	return &this
}

// NewV0040ReservationInfoPurgeCompletedWithDefaults instantiates a new V0040ReservationInfoPurgeCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ReservationInfoPurgeCompletedWithDefaults() *V0040ReservationInfoPurgeCompleted {
	this := V0040ReservationInfoPurgeCompleted{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0040ReservationInfoPurgeCompleted) GetTime() V0040Uint32NoVal {
	if o == nil || IsNil(o.Time) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ReservationInfoPurgeCompleted) GetTimeOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0040ReservationInfoPurgeCompleted) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0040Uint32NoVal and assigns it to the Time field.
func (o *V0040ReservationInfoPurgeCompleted) SetTime(v V0040Uint32NoVal) {
	o.Time = &v
}

func (o V0040ReservationInfoPurgeCompleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ReservationInfoPurgeCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableV0040ReservationInfoPurgeCompleted struct {
	value *V0040ReservationInfoPurgeCompleted
	isSet bool
}

func (v NullableV0040ReservationInfoPurgeCompleted) Get() *V0040ReservationInfoPurgeCompleted {
	return v.value
}

func (v *NullableV0040ReservationInfoPurgeCompleted) Set(val *V0040ReservationInfoPurgeCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ReservationInfoPurgeCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ReservationInfoPurgeCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ReservationInfoPurgeCompleted(val *V0040ReservationInfoPurgeCompleted) *NullableV0040ReservationInfoPurgeCompleted {
	return &NullableV0040ReservationInfoPurgeCompleted{value: val, isSet: true}
}

func (v NullableV0040ReservationInfoPurgeCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ReservationInfoPurgeCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



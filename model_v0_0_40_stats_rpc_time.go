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

// checks if the V0040StatsRpcTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StatsRpcTime{}

// V0040StatsRpcTime struct for V0040StatsRpcTime
type V0040StatsRpcTime struct {
	// Average RPC processing time in microseconds
	Average *int64 `json:"average,omitempty"`
	// Total RPC processing time in microseconds
	Total *int64 `json:"total,omitempty"`
}

// NewV0040StatsRpcTime instantiates a new V0040StatsRpcTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StatsRpcTime() *V0040StatsRpcTime {
	this := V0040StatsRpcTime{}
	return &this
}

// NewV0040StatsRpcTimeWithDefaults instantiates a new V0040StatsRpcTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StatsRpcTimeWithDefaults() *V0040StatsRpcTime {
	this := V0040StatsRpcTime{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *V0040StatsRpcTime) GetAverage() int64 {
	if o == nil || IsNil(o.Average) {
		var ret int64
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsRpcTime) GetAverageOk() (*int64, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *V0040StatsRpcTime) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given int64 and assigns it to the Average field.
func (o *V0040StatsRpcTime) SetAverage(v int64) {
	o.Average = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040StatsRpcTime) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsRpcTime) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040StatsRpcTime) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *V0040StatsRpcTime) SetTotal(v int64) {
	o.Total = &v
}

func (o V0040StatsRpcTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StatsRpcTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0040StatsRpcTime struct {
	value *V0040StatsRpcTime
	isSet bool
}

func (v NullableV0040StatsRpcTime) Get() *V0040StatsRpcTime {
	return v.value
}

func (v *NullableV0040StatsRpcTime) Set(val *V0040StatsRpcTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StatsRpcTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StatsRpcTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StatsRpcTime(val *V0040StatsRpcTime) *NullableV0040StatsRpcTime {
	return &NullableV0040StatsRpcTime{value: val, isSet: true}
}

func (v NullableV0040StatsRpcTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StatsRpcTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



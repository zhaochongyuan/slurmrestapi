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

// checks if the V0040StatsRpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StatsRpc{}

// V0040StatsRpc struct for V0040StatsRpc
type V0040StatsRpc struct {
	// RPC type
	Rpc *string `json:"rpc,omitempty"`
	// Number of RPCs processed
	Count *int32 `json:"count,omitempty"`
	Time *V0040StatsRpcTime `json:"time,omitempty"`
}

// NewV0040StatsRpc instantiates a new V0040StatsRpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StatsRpc() *V0040StatsRpc {
	this := V0040StatsRpc{}
	return &this
}

// NewV0040StatsRpcWithDefaults instantiates a new V0040StatsRpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StatsRpcWithDefaults() *V0040StatsRpc {
	this := V0040StatsRpc{}
	return &this
}

// GetRpc returns the Rpc field value if set, zero value otherwise.
func (o *V0040StatsRpc) GetRpc() string {
	if o == nil || IsNil(o.Rpc) {
		var ret string
		return ret
	}
	return *o.Rpc
}

// GetRpcOk returns a tuple with the Rpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsRpc) GetRpcOk() (*string, bool) {
	if o == nil || IsNil(o.Rpc) {
		return nil, false
	}
	return o.Rpc, true
}

// HasRpc returns a boolean if a field has been set.
func (o *V0040StatsRpc) HasRpc() bool {
	if o != nil && !IsNil(o.Rpc) {
		return true
	}

	return false
}

// SetRpc gets a reference to the given string and assigns it to the Rpc field.
func (o *V0040StatsRpc) SetRpc(v string) {
	o.Rpc = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040StatsRpc) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsRpc) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040StatsRpc) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0040StatsRpc) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0040StatsRpc) GetTime() V0040StatsRpcTime {
	if o == nil || IsNil(o.Time) {
		var ret V0040StatsRpcTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsRpc) GetTimeOk() (*V0040StatsRpcTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0040StatsRpc) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0040StatsRpcTime and assigns it to the Time field.
func (o *V0040StatsRpc) SetTime(v V0040StatsRpcTime) {
	o.Time = &v
}

func (o V0040StatsRpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StatsRpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rpc) {
		toSerialize["rpc"] = o.Rpc
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableV0040StatsRpc struct {
	value *V0040StatsRpc
	isSet bool
}

func (v NullableV0040StatsRpc) Get() *V0040StatsRpc {
	return v.value
}

func (v *NullableV0040StatsRpc) Set(val *V0040StatsRpc) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StatsRpc) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StatsRpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StatsRpc(val *V0040StatsRpc) *NullableV0040StatsRpc {
	return &NullableV0040StatsRpc{value: val, isSet: true}
}

func (v NullableV0040StatsRpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StatsRpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



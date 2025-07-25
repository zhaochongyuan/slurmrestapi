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

// checks if the V0042StatsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042StatsUser{}

// V0042StatsUser struct for V0042StatsUser
type V0042StatsUser struct {
	// User ID
	User *string `json:"user,omitempty"`
	// Number of RPCs processed
	Count *int32 `json:"count,omitempty"`
	Time *V0040StatsRpcTime `json:"time,omitempty"`
}

// NewV0042StatsUser instantiates a new V0042StatsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042StatsUser() *V0042StatsUser {
	this := V0042StatsUser{}
	return &this
}

// NewV0042StatsUserWithDefaults instantiates a new V0042StatsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042StatsUserWithDefaults() *V0042StatsUser {
	this := V0042StatsUser{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0042StatsUser) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042StatsUser) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0042StatsUser) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V0042StatsUser) SetUser(v string) {
	o.User = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0042StatsUser) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042StatsUser) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0042StatsUser) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0042StatsUser) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0042StatsUser) GetTime() V0040StatsRpcTime {
	if o == nil || IsNil(o.Time) {
		var ret V0040StatsRpcTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042StatsUser) GetTimeOk() (*V0040StatsRpcTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0042StatsUser) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0040StatsRpcTime and assigns it to the Time field.
func (o *V0042StatsUser) SetTime(v V0040StatsRpcTime) {
	o.Time = &v
}

func (o V0042StatsUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042StatsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableV0042StatsUser struct {
	value *V0042StatsUser
	isSet bool
}

func (v NullableV0042StatsUser) Get() *V0042StatsUser {
	return v.value
}

func (v *NullableV0042StatsUser) Set(val *V0042StatsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042StatsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042StatsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042StatsUser(val *V0042StatsUser) *NullableV0042StatsUser {
	return &NullableV0042StatsUser{value: val, isSet: true}
}

func (v NullableV0042StatsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042StatsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



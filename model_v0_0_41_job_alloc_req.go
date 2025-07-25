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

// checks if the V0041JobAllocReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobAllocReq{}

// V0041JobAllocReq struct for V0041JobAllocReq
type V0041JobAllocReq struct {
	// HetJob description
	Hetjob []V0041JobDescMsg `json:"hetjob,omitempty"`
	Job *V0041JobDescMsg `json:"job,omitempty"`
}

// NewV0041JobAllocReq instantiates a new V0041JobAllocReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobAllocReq() *V0041JobAllocReq {
	this := V0041JobAllocReq{}
	return &this
}

// NewV0041JobAllocReqWithDefaults instantiates a new V0041JobAllocReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobAllocReqWithDefaults() *V0041JobAllocReq {
	this := V0041JobAllocReq{}
	return &this
}

// GetHetjob returns the Hetjob field value if set, zero value otherwise.
func (o *V0041JobAllocReq) GetHetjob() []V0041JobDescMsg {
	if o == nil || IsNil(o.Hetjob) {
		var ret []V0041JobDescMsg
		return ret
	}
	return o.Hetjob
}

// GetHetjobOk returns a tuple with the Hetjob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobAllocReq) GetHetjobOk() ([]V0041JobDescMsg, bool) {
	if o == nil || IsNil(o.Hetjob) {
		return nil, false
	}
	return o.Hetjob, true
}

// HasHetjob returns a boolean if a field has been set.
func (o *V0041JobAllocReq) HasHetjob() bool {
	if o != nil && !IsNil(o.Hetjob) {
		return true
	}

	return false
}

// SetHetjob gets a reference to the given []V0041JobDescMsg and assigns it to the Hetjob field.
func (o *V0041JobAllocReq) SetHetjob(v []V0041JobDescMsg) {
	o.Hetjob = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041JobAllocReq) GetJob() V0041JobDescMsg {
	if o == nil || IsNil(o.Job) {
		var ret V0041JobDescMsg
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobAllocReq) GetJobOk() (*V0041JobDescMsg, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041JobAllocReq) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0041JobDescMsg and assigns it to the Job field.
func (o *V0041JobAllocReq) SetJob(v V0041JobDescMsg) {
	o.Job = &v
}

func (o V0041JobAllocReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobAllocReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hetjob) {
		toSerialize["hetjob"] = o.Hetjob
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0041JobAllocReq struct {
	value *V0041JobAllocReq
	isSet bool
}

func (v NullableV0041JobAllocReq) Get() *V0041JobAllocReq {
	return v.value
}

func (v *NullableV0041JobAllocReq) Set(val *V0041JobAllocReq) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobAllocReq) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobAllocReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobAllocReq(val *V0041JobAllocReq) *NullableV0041JobAllocReq {
	return &NullableV0041JobAllocReq{value: val, isSet: true}
}

func (v NullableV0041JobAllocReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobAllocReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



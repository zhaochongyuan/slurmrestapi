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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer struct {
	Qos *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos `json:"qos,omitempty"`
	Job *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob `json:"job,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) GetQos() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos {
	if o == nil || IsNil(o.Qos) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) GetQosOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos and assigns it to the Qos field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) SetQos(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerQos) {
	o.Qos = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) GetJob() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob {
	if o == nil || IsNil(o.Job) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) GetJobOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob and assigns it to the Job field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) SetJob(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob) {
	o.Job = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



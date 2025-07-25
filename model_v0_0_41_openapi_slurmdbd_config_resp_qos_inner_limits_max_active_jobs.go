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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs struct {
	Accruing *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing `json:"accruing,omitempty"`
	Count *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount `json:"count,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs{}
	return &this
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) GetAccruing() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing {
	if o == nil || IsNil(o.Accruing) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) GetAccruingOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing and assigns it to the Accruing field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) SetAccruing(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing) {
	o.Accruing = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) GetCount() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount {
	if o == nil || IsNil(o.Count) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) GetCountOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount and assigns it to the Count field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) SetCount(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount) {
	o.Count = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer struct {
	Count *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount `json:"count,omitempty"`
	Accruing *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing `json:"accruing,omitempty"`
	Submitted *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted `json:"submitted,omitempty"`
	WallClock *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob `json:"wall_clock,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetCount() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount {
	if o == nil || IsNil(o.Count) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetCountOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount and assigns it to the Count field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) SetCount(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsCount) {
	o.Count = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetAccruing() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing {
	if o == nil || IsNil(o.Accruing) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetAccruingOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing and assigns it to the Accruing field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) SetAccruing(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxActiveJobsAccruing) {
	o.Accruing = &v
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetSubmitted() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted {
	if o == nil || IsNil(o.Submitted) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted
		return ret
	}
	return *o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetSubmittedOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) HasSubmitted() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted and assigns it to the Submitted field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) SetSubmitted(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPerSubmitted) {
	o.Submitted = &v
}

// GetWallClock returns the WallClock field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetWallClock() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob {
	if o == nil || IsNil(o.WallClock) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob
		return ret
	}
	return *o.WallClock
}

// GetWallClockOk returns a tuple with the WallClock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) GetWallClockOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob, bool) {
	if o == nil || IsNil(o.WallClock) {
		return nil, false
	}
	return o.WallClock, true
}

// HasWallClock returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) HasWallClock() bool {
	if o != nil && !IsNil(o.WallClock) {
		return true
	}

	return false
}

// SetWallClock gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob and assigns it to the WallClock field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) SetWallClock(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxWallClockPerJob) {
	o.WallClock = &v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Submitted) {
		toSerialize["submitted"] = o.Submitted
	}
	if !IsNil(o.WallClock) {
		toSerialize["wall_clock"] = o.WallClock
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



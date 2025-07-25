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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin struct {
	PriorityThreshold *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold `json:"priority_threshold,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMinWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMinWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) GetPriorityThreshold() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) GetPriorityThresholdOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold and assigns it to the PriorityThreshold field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) SetPriorityThreshold(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinPriorityThreshold) {
	o.PriorityThreshold = &v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



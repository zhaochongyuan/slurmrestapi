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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres struct {
	// GrpTRES
	Total []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"total,omitempty"`
	Group *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup `json:"group,omitempty"`
	Minutes *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes `json:"minutes,omitempty"`
	Per *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer `json:"per,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetTotal() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.Total) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetTotalOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the Total field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) SetTotal(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.Total = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetGroup() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup {
	if o == nil || IsNil(o.Group) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetGroupOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup and assigns it to the Group field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) SetGroup(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) {
	o.Group = &v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetMinutes() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes {
	if o == nil || IsNil(o.Minutes) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetMinutesOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes and assigns it to the Minutes field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) SetMinutes(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutes) {
	o.Minutes = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetPer() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer {
	if o == nil || IsNil(o.Per) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) GetPerOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer and assigns it to the Per field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) SetPer(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresPer) {
	o.Per = &v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



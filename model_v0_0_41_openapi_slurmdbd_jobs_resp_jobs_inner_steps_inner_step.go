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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep struct for V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep struct {
	// Step ID
	Id *string `json:"id,omitempty"`
	// Step name
	Name *string `json:"name,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStepWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStepWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) SetName(v string) {
	o.Name = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



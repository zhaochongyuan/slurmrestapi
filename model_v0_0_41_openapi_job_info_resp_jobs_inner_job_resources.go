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
	"bytes"
	"fmt"
)

// checks if the V0041OpenapiJobInfoRespJobsInnerJobResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobInfoRespJobsInnerJobResources{}

// V0041OpenapiJobInfoRespJobsInnerJobResources Resources used by the job
type V0041OpenapiJobInfoRespJobsInnerJobResources struct {
	// Scheduler consumable resource selection type
	SelectType []string `json:"select_type"`
	Nodes *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes `json:"nodes,omitempty"`
	// Number of allocated CPUs
	Cpus int32 `json:"cpus"`
	ThreadsPerCore V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore `json:"threads_per_core"`
}

type _V0041OpenapiJobInfoRespJobsInnerJobResources V0041OpenapiJobInfoRespJobsInnerJobResources

// NewV0041OpenapiJobInfoRespJobsInnerJobResources instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobInfoRespJobsInnerJobResources(selectType []string, cpus int32, threadsPerCore V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore) *V0041OpenapiJobInfoRespJobsInnerJobResources {
	this := V0041OpenapiJobInfoRespJobsInnerJobResources{}
	this.SelectType = selectType
	this.Cpus = cpus
	this.ThreadsPerCore = threadsPerCore
	return &this
}

// NewV0041OpenapiJobInfoRespJobsInnerJobResourcesWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResources {
	this := V0041OpenapiJobInfoRespJobsInnerJobResources{}
	return &this
}

// GetSelectType returns the SelectType field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetSelectType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SelectType
}

// GetSelectTypeOk returns a tuple with the SelectType field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetSelectTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectType, true
}

// SetSelectType sets field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetSelectType(v []string) {
	o.SelectType = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetNodes() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes {
	if o == nil || IsNil(o.Nodes) {
		var ret V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetNodesOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes and assigns it to the Nodes field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetNodes(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) {
	o.Nodes = &v
}

// GetCpus returns the Cpus field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetCpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetCpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpus, true
}

// SetCpus sets field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetCpus(v int32) {
	o.Cpus = v
}

// GetThreadsPerCore returns the ThreadsPerCore field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetThreadsPerCore() V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore {
	if o == nil {
		var ret V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore
		return ret
	}

	return o.ThreadsPerCore
}

// GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetThreadsPerCoreOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadsPerCore, true
}

// SetThreadsPerCore sets field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetThreadsPerCore(v V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore) {
	o.ThreadsPerCore = v
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["select_type"] = o.SelectType
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	toSerialize["cpus"] = o.Cpus
	toSerialize["threads_per_core"] = o.ThreadsPerCore
	return toSerialize, nil
}

func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"select_type",
		"cpus",
		"threads_per_core",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0041OpenapiJobInfoRespJobsInnerJobResources := _V0041OpenapiJobInfoRespJobsInnerJobResources{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiJobInfoRespJobsInnerJobResources)

	if err != nil {
		return err
	}

	*o = V0041OpenapiJobInfoRespJobsInnerJobResources(varV0041OpenapiJobInfoRespJobsInnerJobResources)

	return err
}

type NullableV0041OpenapiJobInfoRespJobsInnerJobResources struct {
	value *V0041OpenapiJobInfoRespJobsInnerJobResources
	isSet bool
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResources) Get() *V0041OpenapiJobInfoRespJobsInnerJobResources {
	return v.value
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResources) Set(val *V0041OpenapiJobInfoRespJobsInnerJobResources) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResources) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobInfoRespJobsInnerJobResources(val *V0041OpenapiJobInfoRespJobsInnerJobResources) *NullableV0041OpenapiJobInfoRespJobsInnerJobResources {
	return &NullableV0041OpenapiJobInfoRespJobsInnerJobResources{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



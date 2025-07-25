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

// checks if the V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner{}

// V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner struct for V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner
type V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner struct {
	// Core index
	Index int32 `json:"index"`
	// Core status
	Status []string `json:"status"`
}

type _V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner

// NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner(index int32, status []string) *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner {
	this := V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner{}
	this.Index = index
	this.Status = status
	return &this
}

// NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInnerWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInnerWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner {
	this := V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner{}
	return &this
}

// GetIndex returns the Index field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) SetIndex(v int32) {
	o.Index = v
}

// GetStatus returns the Status field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) GetStatus() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) GetStatusOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) SetStatus(v []string) {
	o.Status = v
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"status",
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

	varV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner := _V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner)

	if err != nil {
		return err
	}

	*o = V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner(varV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner)

	return err
}

type NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner struct {
	value *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner
	isSet bool
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) Get() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner {
	return v.value
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) Set(val *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner(val *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner {
	return &NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInnerCoresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



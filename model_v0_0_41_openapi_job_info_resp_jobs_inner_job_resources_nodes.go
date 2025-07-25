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

// checks if the V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes{}

// V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes struct for V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes
type V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes struct {
	// Number of allocated nodes
	Count *int32 `json:"count,omitempty"`
	// Node scheduling selection method
	SelectType []string `json:"select_type,omitempty"`
	// Node(s) allocated to the job
	List *string `json:"list,omitempty"`
	// Whether whole nodes were allocated
	Whole *bool `json:"whole,omitempty"`
	// Allocated node resources
	Allocation []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner `json:"allocation,omitempty"`
}

// NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes {
	this := V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes{}
	return &this
}

// NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes {
	this := V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetCount(v int32) {
	o.Count = &v
}

// GetSelectType returns the SelectType field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetSelectType() []string {
	if o == nil || IsNil(o.SelectType) {
		var ret []string
		return ret
	}
	return o.SelectType
}

// GetSelectTypeOk returns a tuple with the SelectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetSelectTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.SelectType) {
		return nil, false
	}
	return o.SelectType, true
}

// HasSelectType returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasSelectType() bool {
	if o != nil && !IsNil(o.SelectType) {
		return true
	}

	return false
}

// SetSelectType gets a reference to the given []string and assigns it to the SelectType field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetSelectType(v []string) {
	o.SelectType = v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetList() string {
	if o == nil || IsNil(o.List) {
		var ret string
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetListOk() (*string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given string and assigns it to the List field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetList(v string) {
	o.List = &v
}

// GetWhole returns the Whole field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetWhole() bool {
	if o == nil || IsNil(o.Whole) {
		var ret bool
		return ret
	}
	return *o.Whole
}

// GetWholeOk returns a tuple with the Whole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetWholeOk() (*bool, bool) {
	if o == nil || IsNil(o.Whole) {
		return nil, false
	}
	return o.Whole, true
}

// HasWhole returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasWhole() bool {
	if o != nil && !IsNil(o.Whole) {
		return true
	}

	return false
}

// SetWhole gets a reference to the given bool and assigns it to the Whole field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetWhole(v bool) {
	o.Whole = &v
}

// GetAllocation returns the Allocation field value if set, zero value otherwise.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetAllocation() []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner {
	if o == nil || IsNil(o.Allocation) {
		var ret []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner
		return ret
	}
	return o.Allocation
}

// GetAllocationOk returns a tuple with the Allocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetAllocationOk() ([]V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner, bool) {
	if o == nil || IsNil(o.Allocation) {
		return nil, false
	}
	return o.Allocation, true
}

// HasAllocation returns a boolean if a field has been set.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasAllocation() bool {
	if o != nil && !IsNil(o.Allocation) {
		return true
	}

	return false
}

// SetAllocation gets a reference to the given []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner and assigns it to the Allocation field.
func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetAllocation(v []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) {
	o.Allocation = v
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.SelectType) {
		toSerialize["select_type"] = o.SelectType
	}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Whole) {
		toSerialize["whole"] = o.Whole
	}
	if !IsNil(o.Allocation) {
		toSerialize["allocation"] = o.Allocation
	}
	return toSerialize, nil
}

type NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes struct {
	value *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes
	isSet bool
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) Get() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes {
	return v.value
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) Set(val *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes(val *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes {
	return &NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



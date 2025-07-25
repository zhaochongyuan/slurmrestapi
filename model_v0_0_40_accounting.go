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

// checks if the V0040Accounting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040Accounting{}

// V0040Accounting struct for V0040Accounting
type V0040Accounting struct {
	Allocated *V0040AccountingAllocated `json:"allocated,omitempty"`
	// Association ID or Workload characterization key ID
	Id *int32 `json:"id,omitempty"`
	// When the record was started
	Start *int64 `json:"start,omitempty"`
	TRES *V0040Tres `json:"TRES,omitempty"`
}

// NewV0040Accounting instantiates a new V0040Accounting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040Accounting() *V0040Accounting {
	this := V0040Accounting{}
	return &this
}

// NewV0040AccountingWithDefaults instantiates a new V0040Accounting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AccountingWithDefaults() *V0040Accounting {
	this := V0040Accounting{}
	return &this
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *V0040Accounting) GetAllocated() V0040AccountingAllocated {
	if o == nil || IsNil(o.Allocated) {
		var ret V0040AccountingAllocated
		return ret
	}
	return *o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Accounting) GetAllocatedOk() (*V0040AccountingAllocated, bool) {
	if o == nil || IsNil(o.Allocated) {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *V0040Accounting) HasAllocated() bool {
	if o != nil && !IsNil(o.Allocated) {
		return true
	}

	return false
}

// SetAllocated gets a reference to the given V0040AccountingAllocated and assigns it to the Allocated field.
func (o *V0040Accounting) SetAllocated(v V0040AccountingAllocated) {
	o.Allocated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0040Accounting) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Accounting) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0040Accounting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0040Accounting) SetId(v int32) {
	o.Id = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *V0040Accounting) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Accounting) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *V0040Accounting) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *V0040Accounting) SetStart(v int64) {
	o.Start = &v
}

// GetTRES returns the TRES field value if set, zero value otherwise.
func (o *V0040Accounting) GetTRES() V0040Tres {
	if o == nil || IsNil(o.TRES) {
		var ret V0040Tres
		return ret
	}
	return *o.TRES
}

// GetTRESOk returns a tuple with the TRES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Accounting) GetTRESOk() (*V0040Tres, bool) {
	if o == nil || IsNil(o.TRES) {
		return nil, false
	}
	return o.TRES, true
}

// HasTRES returns a boolean if a field has been set.
func (o *V0040Accounting) HasTRES() bool {
	if o != nil && !IsNil(o.TRES) {
		return true
	}

	return false
}

// SetTRES gets a reference to the given V0040Tres and assigns it to the TRES field.
func (o *V0040Accounting) SetTRES(v V0040Tres) {
	o.TRES = &v
}

func (o V0040Accounting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040Accounting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allocated) {
		toSerialize["allocated"] = o.Allocated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.TRES) {
		toSerialize["TRES"] = o.TRES
	}
	return toSerialize, nil
}

type NullableV0040Accounting struct {
	value *V0040Accounting
	isSet bool
}

func (v NullableV0040Accounting) Get() *V0040Accounting {
	return v.value
}

func (v *NullableV0040Accounting) Set(val *V0040Accounting) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040Accounting) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040Accounting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040Accounting(val *V0040Accounting) *NullableV0040Accounting {
	return &NullableV0040Accounting{value: val, isSet: true}
}

func (v NullableV0040Accounting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040Accounting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



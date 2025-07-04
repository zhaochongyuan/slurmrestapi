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

// checks if the V0040OpenapiSlurmdbdStatsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiSlurmdbdStatsResp{}

// V0040OpenapiSlurmdbdStatsResp struct for V0040OpenapiSlurmdbdStatsResp
type V0040OpenapiSlurmdbdStatsResp struct {
	Statistics V0040StatsRec `json:"statistics"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiSlurmdbdStatsResp V0040OpenapiSlurmdbdStatsResp

// NewV0040OpenapiSlurmdbdStatsResp instantiates a new V0040OpenapiSlurmdbdStatsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiSlurmdbdStatsResp(statistics V0040StatsRec) *V0040OpenapiSlurmdbdStatsResp {
	this := V0040OpenapiSlurmdbdStatsResp{}
	this.Statistics = statistics
	return &this
}

// NewV0040OpenapiSlurmdbdStatsRespWithDefaults instantiates a new V0040OpenapiSlurmdbdStatsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiSlurmdbdStatsRespWithDefaults() *V0040OpenapiSlurmdbdStatsResp {
	this := V0040OpenapiSlurmdbdStatsResp{}
	return &this
}

// GetStatistics returns the Statistics field value
func (o *V0040OpenapiSlurmdbdStatsResp) GetStatistics() V0040StatsRec {
	if o == nil {
		var ret V0040StatsRec
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) GetStatisticsOk() (*V0040StatsRec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// SetStatistics sets field value
func (o *V0040OpenapiSlurmdbdStatsResp) SetStatistics(v V0040StatsRec) {
	o.Statistics = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdStatsResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiSlurmdbdStatsResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdStatsResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiSlurmdbdStatsResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdStatsResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdStatsResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiSlurmdbdStatsResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiSlurmdbdStatsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiSlurmdbdStatsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statistics"] = o.Statistics
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0040OpenapiSlurmdbdStatsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statistics",
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

	varV0040OpenapiSlurmdbdStatsResp := _V0040OpenapiSlurmdbdStatsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiSlurmdbdStatsResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiSlurmdbdStatsResp(varV0040OpenapiSlurmdbdStatsResp)

	return err
}

type NullableV0040OpenapiSlurmdbdStatsResp struct {
	value *V0040OpenapiSlurmdbdStatsResp
	isSet bool
}

func (v NullableV0040OpenapiSlurmdbdStatsResp) Get() *V0040OpenapiSlurmdbdStatsResp {
	return v.value
}

func (v *NullableV0040OpenapiSlurmdbdStatsResp) Set(val *V0040OpenapiSlurmdbdStatsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiSlurmdbdStatsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiSlurmdbdStatsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiSlurmdbdStatsResp(val *V0040OpenapiSlurmdbdStatsResp) *NullableV0040OpenapiSlurmdbdStatsResp {
	return &NullableV0040OpenapiSlurmdbdStatsResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiSlurmdbdStatsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiSlurmdbdStatsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



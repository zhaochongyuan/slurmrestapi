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

// checks if the V0042OpenapiAssocsRemovedResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042OpenapiAssocsRemovedResp{}

// V0042OpenapiAssocsRemovedResp struct for V0042OpenapiAssocsRemovedResp
type V0042OpenapiAssocsRemovedResp struct {
	RemovedAssociations []string `json:"removed_associations"`
	Meta *V0042OpenapiMeta `json:"meta,omitempty"`
	Errors []V0042OpenapiError `json:"errors,omitempty"`
	Warnings []V0042OpenapiWarning `json:"warnings,omitempty"`
}

type _V0042OpenapiAssocsRemovedResp V0042OpenapiAssocsRemovedResp

// NewV0042OpenapiAssocsRemovedResp instantiates a new V0042OpenapiAssocsRemovedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042OpenapiAssocsRemovedResp(removedAssociations []string) *V0042OpenapiAssocsRemovedResp {
	this := V0042OpenapiAssocsRemovedResp{}
	this.RemovedAssociations = removedAssociations
	return &this
}

// NewV0042OpenapiAssocsRemovedRespWithDefaults instantiates a new V0042OpenapiAssocsRemovedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042OpenapiAssocsRemovedRespWithDefaults() *V0042OpenapiAssocsRemovedResp {
	this := V0042OpenapiAssocsRemovedResp{}
	return &this
}

// GetRemovedAssociations returns the RemovedAssociations field value
func (o *V0042OpenapiAssocsRemovedResp) GetRemovedAssociations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedAssociations
}

// GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field value
// and a boolean to check if the value has been set.
func (o *V0042OpenapiAssocsRemovedResp) GetRemovedAssociationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedAssociations, true
}

// SetRemovedAssociations sets field value
func (o *V0042OpenapiAssocsRemovedResp) SetRemovedAssociations(v []string) {
	o.RemovedAssociations = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0042OpenapiAssocsRemovedResp) GetMeta() V0042OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0042OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiAssocsRemovedResp) GetMetaOk() (*V0042OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0042OpenapiAssocsRemovedResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0042OpenapiMeta and assigns it to the Meta field.
func (o *V0042OpenapiAssocsRemovedResp) SetMeta(v V0042OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0042OpenapiAssocsRemovedResp) GetErrors() []V0042OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0042OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiAssocsRemovedResp) GetErrorsOk() ([]V0042OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0042OpenapiAssocsRemovedResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0042OpenapiError and assigns it to the Errors field.
func (o *V0042OpenapiAssocsRemovedResp) SetErrors(v []V0042OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0042OpenapiAssocsRemovedResp) GetWarnings() []V0042OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0042OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiAssocsRemovedResp) GetWarningsOk() ([]V0042OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0042OpenapiAssocsRemovedResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0042OpenapiWarning and assigns it to the Warnings field.
func (o *V0042OpenapiAssocsRemovedResp) SetWarnings(v []V0042OpenapiWarning) {
	o.Warnings = v
}

func (o V0042OpenapiAssocsRemovedResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042OpenapiAssocsRemovedResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removed_associations"] = o.RemovedAssociations
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

func (o *V0042OpenapiAssocsRemovedResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removed_associations",
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

	varV0042OpenapiAssocsRemovedResp := _V0042OpenapiAssocsRemovedResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0042OpenapiAssocsRemovedResp)

	if err != nil {
		return err
	}

	*o = V0042OpenapiAssocsRemovedResp(varV0042OpenapiAssocsRemovedResp)

	return err
}

type NullableV0042OpenapiAssocsRemovedResp struct {
	value *V0042OpenapiAssocsRemovedResp
	isSet bool
}

func (v NullableV0042OpenapiAssocsRemovedResp) Get() *V0042OpenapiAssocsRemovedResp {
	return v.value
}

func (v *NullableV0042OpenapiAssocsRemovedResp) Set(val *V0042OpenapiAssocsRemovedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042OpenapiAssocsRemovedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042OpenapiAssocsRemovedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042OpenapiAssocsRemovedResp(val *V0042OpenapiAssocsRemovedResp) *NullableV0042OpenapiAssocsRemovedResp {
	return &NullableV0042OpenapiAssocsRemovedResp{value: val, isSet: true}
}

func (v NullableV0042OpenapiAssocsRemovedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042OpenapiAssocsRemovedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



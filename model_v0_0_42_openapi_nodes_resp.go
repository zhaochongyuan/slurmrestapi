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

// checks if the V0042OpenapiNodesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042OpenapiNodesResp{}

// V0042OpenapiNodesResp struct for V0042OpenapiNodesResp
type V0042OpenapiNodesResp struct {
	Nodes []V0042Node `json:"nodes"`
	LastUpdate V0042Uint64NoValStruct `json:"last_update"`
	Meta *V0042OpenapiMeta `json:"meta,omitempty"`
	Errors []V0042OpenapiError `json:"errors,omitempty"`
	Warnings []V0042OpenapiWarning `json:"warnings,omitempty"`
}

type _V0042OpenapiNodesResp V0042OpenapiNodesResp

// NewV0042OpenapiNodesResp instantiates a new V0042OpenapiNodesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042OpenapiNodesResp(nodes []V0042Node, lastUpdate V0042Uint64NoValStruct) *V0042OpenapiNodesResp {
	this := V0042OpenapiNodesResp{}
	this.Nodes = nodes
	this.LastUpdate = lastUpdate
	return &this
}

// NewV0042OpenapiNodesRespWithDefaults instantiates a new V0042OpenapiNodesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042OpenapiNodesRespWithDefaults() *V0042OpenapiNodesResp {
	this := V0042OpenapiNodesResp{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *V0042OpenapiNodesResp) GetNodes() []V0042Node {
	if o == nil {
		var ret []V0042Node
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *V0042OpenapiNodesResp) GetNodesOk() ([]V0042Node, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *V0042OpenapiNodesResp) SetNodes(v []V0042Node) {
	o.Nodes = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *V0042OpenapiNodesResp) GetLastUpdate() V0042Uint64NoValStruct {
	if o == nil {
		var ret V0042Uint64NoValStruct
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *V0042OpenapiNodesResp) GetLastUpdateOk() (*V0042Uint64NoValStruct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *V0042OpenapiNodesResp) SetLastUpdate(v V0042Uint64NoValStruct) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0042OpenapiNodesResp) GetMeta() V0042OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0042OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiNodesResp) GetMetaOk() (*V0042OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0042OpenapiNodesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0042OpenapiMeta and assigns it to the Meta field.
func (o *V0042OpenapiNodesResp) SetMeta(v V0042OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0042OpenapiNodesResp) GetErrors() []V0042OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0042OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiNodesResp) GetErrorsOk() ([]V0042OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0042OpenapiNodesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0042OpenapiError and assigns it to the Errors field.
func (o *V0042OpenapiNodesResp) SetErrors(v []V0042OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0042OpenapiNodesResp) GetWarnings() []V0042OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0042OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiNodesResp) GetWarningsOk() ([]V0042OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0042OpenapiNodesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0042OpenapiWarning and assigns it to the Warnings field.
func (o *V0042OpenapiNodesResp) SetWarnings(v []V0042OpenapiWarning) {
	o.Warnings = v
}

func (o V0042OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042OpenapiNodesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
	toSerialize["last_update"] = o.LastUpdate
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

func (o *V0042OpenapiNodesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodes",
		"last_update",
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

	varV0042OpenapiNodesResp := _V0042OpenapiNodesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0042OpenapiNodesResp)

	if err != nil {
		return err
	}

	*o = V0042OpenapiNodesResp(varV0042OpenapiNodesResp)

	return err
}

type NullableV0042OpenapiNodesResp struct {
	value *V0042OpenapiNodesResp
	isSet bool
}

func (v NullableV0042OpenapiNodesResp) Get() *V0042OpenapiNodesResp {
	return v.value
}

func (v *NullableV0042OpenapiNodesResp) Set(val *V0042OpenapiNodesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042OpenapiNodesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042OpenapiNodesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042OpenapiNodesResp(val *V0042OpenapiNodesResp) *NullableV0042OpenapiNodesResp {
	return &NullableV0042OpenapiNodesResp{value: val, isSet: true}
}

func (v NullableV0042OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042OpenapiNodesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



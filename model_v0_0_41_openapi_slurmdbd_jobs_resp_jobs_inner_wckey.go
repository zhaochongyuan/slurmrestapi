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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerWckey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerWckey{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerWckey Workload characterization key
type V0041OpenapiSlurmdbdJobsRespJobsInnerWckey struct {
	// WCKey name
	Wckey string `json:"wckey"`
	// Active flags
	Flags []string `json:"flags"`
}

type _V0041OpenapiSlurmdbdJobsRespJobsInnerWckey V0041OpenapiSlurmdbdJobsRespJobsInnerWckey

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerWckey instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerWckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerWckey(wckey string, flags []string) *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerWckey{}
	this.Wckey = wckey
	this.Flags = flags
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerWckeyWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerWckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerWckeyWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerWckey{}
	return &this
}

// GetWckey returns the Wckey field value
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) GetWckey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) GetWckeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wckey, true
}

// SetWckey sets field value
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) SetWckey(v string) {
	o.Wckey = v
}

// GetFlags returns the Flags field value
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) GetFlags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) GetFlagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) SetFlags(v []string) {
	o.Flags = v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wckey"] = o.Wckey
	toSerialize["flags"] = o.Flags
	return toSerialize, nil
}

func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wckey",
		"flags",
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

	varV0041OpenapiSlurmdbdJobsRespJobsInnerWckey := _V0041OpenapiSlurmdbdJobsRespJobsInnerWckey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiSlurmdbdJobsRespJobsInnerWckey)

	if err != nil {
		return err
	}

	*o = V0041OpenapiSlurmdbdJobsRespJobsInnerWckey(varV0041OpenapiSlurmdbdJobsRespJobsInnerWckey)

	return err
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey(val *V0041OpenapiSlurmdbdJobsRespJobsInnerWckey) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerWckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



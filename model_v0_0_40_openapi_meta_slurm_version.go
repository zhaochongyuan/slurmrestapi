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

// checks if the V0040OpenapiMetaSlurmVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiMetaSlurmVersion{}

// V0040OpenapiMetaSlurmVersion struct for V0040OpenapiMetaSlurmVersion
type V0040OpenapiMetaSlurmVersion struct {
	// Slurm release major version
	Major *string `json:"major,omitempty"`
	// Slurm release micro version
	Micro *string `json:"micro,omitempty"`
	// Slurm release minor version
	Minor *string `json:"minor,omitempty"`
}

// NewV0040OpenapiMetaSlurmVersion instantiates a new V0040OpenapiMetaSlurmVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiMetaSlurmVersion() *V0040OpenapiMetaSlurmVersion {
	this := V0040OpenapiMetaSlurmVersion{}
	return &this
}

// NewV0040OpenapiMetaSlurmVersionWithDefaults instantiates a new V0040OpenapiMetaSlurmVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiMetaSlurmVersionWithDefaults() *V0040OpenapiMetaSlurmVersion {
	this := V0040OpenapiMetaSlurmVersion{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *V0040OpenapiMetaSlurmVersion) GetMajor() string {
	if o == nil || IsNil(o.Major) {
		var ret string
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMetaSlurmVersion) GetMajorOk() (*string, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *V0040OpenapiMetaSlurmVersion) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given string and assigns it to the Major field.
func (o *V0040OpenapiMetaSlurmVersion) SetMajor(v string) {
	o.Major = &v
}

// GetMicro returns the Micro field value if set, zero value otherwise.
func (o *V0040OpenapiMetaSlurmVersion) GetMicro() string {
	if o == nil || IsNil(o.Micro) {
		var ret string
		return ret
	}
	return *o.Micro
}

// GetMicroOk returns a tuple with the Micro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMetaSlurmVersion) GetMicroOk() (*string, bool) {
	if o == nil || IsNil(o.Micro) {
		return nil, false
	}
	return o.Micro, true
}

// HasMicro returns a boolean if a field has been set.
func (o *V0040OpenapiMetaSlurmVersion) HasMicro() bool {
	if o != nil && !IsNil(o.Micro) {
		return true
	}

	return false
}

// SetMicro gets a reference to the given string and assigns it to the Micro field.
func (o *V0040OpenapiMetaSlurmVersion) SetMicro(v string) {
	o.Micro = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *V0040OpenapiMetaSlurmVersion) GetMinor() string {
	if o == nil || IsNil(o.Minor) {
		var ret string
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiMetaSlurmVersion) GetMinorOk() (*string, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *V0040OpenapiMetaSlurmVersion) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given string and assigns it to the Minor field.
func (o *V0040OpenapiMetaSlurmVersion) SetMinor(v string) {
	o.Minor = &v
}

func (o V0040OpenapiMetaSlurmVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiMetaSlurmVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Micro) {
		toSerialize["micro"] = o.Micro
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return toSerialize, nil
}

type NullableV0040OpenapiMetaSlurmVersion struct {
	value *V0040OpenapiMetaSlurmVersion
	isSet bool
}

func (v NullableV0040OpenapiMetaSlurmVersion) Get() *V0040OpenapiMetaSlurmVersion {
	return v.value
}

func (v *NullableV0040OpenapiMetaSlurmVersion) Set(val *V0040OpenapiMetaSlurmVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiMetaSlurmVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiMetaSlurmVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiMetaSlurmVersion(val *V0040OpenapiMetaSlurmVersion) *NullableV0040OpenapiMetaSlurmVersion {
	return &NullableV0040OpenapiMetaSlurmVersion{value: val, isSet: true}
}

func (v NullableV0040OpenapiMetaSlurmVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiMetaSlurmVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



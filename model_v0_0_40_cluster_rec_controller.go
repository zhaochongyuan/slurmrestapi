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

// checks if the V0040ClusterRecController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ClusterRecController{}

// V0040ClusterRecController struct for V0040ClusterRecController
type V0040ClusterRecController struct {
	// ControlHost
	Host *string `json:"host,omitempty"`
	// ControlPort
	Port *int32 `json:"port,omitempty"`
}

// NewV0040ClusterRecController instantiates a new V0040ClusterRecController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ClusterRecController() *V0040ClusterRecController {
	this := V0040ClusterRecController{}
	return &this
}

// NewV0040ClusterRecControllerWithDefaults instantiates a new V0040ClusterRecController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ClusterRecControllerWithDefaults() *V0040ClusterRecController {
	this := V0040ClusterRecController{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V0040ClusterRecController) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRecController) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V0040ClusterRecController) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V0040ClusterRecController) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V0040ClusterRecController) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ClusterRecController) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V0040ClusterRecController) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *V0040ClusterRecController) SetPort(v int32) {
	o.Port = &v
}

func (o V0040ClusterRecController) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ClusterRecController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableV0040ClusterRecController struct {
	value *V0040ClusterRecController
	isSet bool
}

func (v NullableV0040ClusterRecController) Get() *V0040ClusterRecController {
	return v.value
}

func (v *NullableV0040ClusterRecController) Set(val *V0040ClusterRecController) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ClusterRecController) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ClusterRecController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ClusterRecController(val *V0040ClusterRecController) *NullableV0040ClusterRecController {
	return &NullableV0040ClusterRecController{value: val, isSet: true}
}

func (v NullableV0040ClusterRecController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ClusterRecController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



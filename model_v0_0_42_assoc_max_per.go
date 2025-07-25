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

// checks if the V0042AssocMaxPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042AssocMaxPer{}

// V0042AssocMaxPer struct for V0042AssocMaxPer
type V0042AssocMaxPer struct {
	Account *V0042AssocMaxPerAccount `json:"account,omitempty"`
}

// NewV0042AssocMaxPer instantiates a new V0042AssocMaxPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042AssocMaxPer() *V0042AssocMaxPer {
	this := V0042AssocMaxPer{}
	return &this
}

// NewV0042AssocMaxPerWithDefaults instantiates a new V0042AssocMaxPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042AssocMaxPerWithDefaults() *V0042AssocMaxPer {
	this := V0042AssocMaxPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0042AssocMaxPer) GetAccount() V0042AssocMaxPerAccount {
	if o == nil || IsNil(o.Account) {
		var ret V0042AssocMaxPerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042AssocMaxPer) GetAccountOk() (*V0042AssocMaxPerAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0042AssocMaxPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0042AssocMaxPerAccount and assigns it to the Account field.
func (o *V0042AssocMaxPer) SetAccount(v V0042AssocMaxPerAccount) {
	o.Account = &v
}

func (o V0042AssocMaxPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042AssocMaxPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableV0042AssocMaxPer struct {
	value *V0042AssocMaxPer
	isSet bool
}

func (v NullableV0042AssocMaxPer) Get() *V0042AssocMaxPer {
	return v.value
}

func (v *NullableV0042AssocMaxPer) Set(val *V0042AssocMaxPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042AssocMaxPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042AssocMaxPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042AssocMaxPer(val *V0042AssocMaxPer) *NullableV0042AssocMaxPer {
	return &NullableV0042AssocMaxPer{value: val, isSet: true}
}

func (v NullableV0042AssocMaxPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042AssocMaxPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the V0040OpenapiAccountsAddCondResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiAccountsAddCondResp{}

// V0040OpenapiAccountsAddCondResp struct for V0040OpenapiAccountsAddCondResp
type V0040OpenapiAccountsAddCondResp struct {
	AssociationCondition *V0040AccountsAddCond `json:"association_condition,omitempty"`
	Account *V0040AccountShort `json:"account,omitempty"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

// NewV0040OpenapiAccountsAddCondResp instantiates a new V0040OpenapiAccountsAddCondResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiAccountsAddCondResp() *V0040OpenapiAccountsAddCondResp {
	this := V0040OpenapiAccountsAddCondResp{}
	return &this
}

// NewV0040OpenapiAccountsAddCondRespWithDefaults instantiates a new V0040OpenapiAccountsAddCondResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiAccountsAddCondRespWithDefaults() *V0040OpenapiAccountsAddCondResp {
	this := V0040OpenapiAccountsAddCondResp{}
	return &this
}

// GetAssociationCondition returns the AssociationCondition field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondResp) GetAssociationCondition() V0040AccountsAddCond {
	if o == nil || IsNil(o.AssociationCondition) {
		var ret V0040AccountsAddCond
		return ret
	}
	return *o.AssociationCondition
}

// GetAssociationConditionOk returns a tuple with the AssociationCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0040AccountsAddCond, bool) {
	if o == nil || IsNil(o.AssociationCondition) {
		return nil, false
	}
	return o.AssociationCondition, true
}

// HasAssociationCondition returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondResp) HasAssociationCondition() bool {
	if o != nil && !IsNil(o.AssociationCondition) {
		return true
	}

	return false
}

// SetAssociationCondition gets a reference to the given V0040AccountsAddCond and assigns it to the AssociationCondition field.
func (o *V0040OpenapiAccountsAddCondResp) SetAssociationCondition(v V0040AccountsAddCond) {
	o.AssociationCondition = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondResp) GetAccount() V0040AccountShort {
	if o == nil || IsNil(o.Account) {
		var ret V0040AccountShort
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondResp) GetAccountOk() (*V0040AccountShort, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondResp) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0040AccountShort and assigns it to the Account field.
func (o *V0040OpenapiAccountsAddCondResp) SetAccount(v V0040AccountShort) {
	o.Account = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiAccountsAddCondResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiAccountsAddCondResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiAccountsAddCondResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiAccountsAddCondResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiAccountsAddCondResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociationCondition) {
		toSerialize["association_condition"] = o.AssociationCondition
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
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

type NullableV0040OpenapiAccountsAddCondResp struct {
	value *V0040OpenapiAccountsAddCondResp
	isSet bool
}

func (v NullableV0040OpenapiAccountsAddCondResp) Get() *V0040OpenapiAccountsAddCondResp {
	return v.value
}

func (v *NullableV0040OpenapiAccountsAddCondResp) Set(val *V0040OpenapiAccountsAddCondResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiAccountsAddCondResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiAccountsAddCondResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiAccountsAddCondResp(val *V0040OpenapiAccountsAddCondResp) *NullableV0040OpenapiAccountsAddCondResp {
	return &NullableV0040OpenapiAccountsAddCondResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiAccountsAddCondResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiAccountsAddCondResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



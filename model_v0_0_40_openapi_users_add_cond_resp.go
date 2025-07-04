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

// checks if the V0040OpenapiUsersAddCondResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiUsersAddCondResp{}

// V0040OpenapiUsersAddCondResp struct for V0040OpenapiUsersAddCondResp
type V0040OpenapiUsersAddCondResp struct {
	AssociationCondition V0040UsersAddCond `json:"association_condition"`
	User V0040UserShort `json:"user"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiUsersAddCondResp V0040OpenapiUsersAddCondResp

// NewV0040OpenapiUsersAddCondResp instantiates a new V0040OpenapiUsersAddCondResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiUsersAddCondResp(associationCondition V0040UsersAddCond, user V0040UserShort) *V0040OpenapiUsersAddCondResp {
	this := V0040OpenapiUsersAddCondResp{}
	this.AssociationCondition = associationCondition
	this.User = user
	return &this
}

// NewV0040OpenapiUsersAddCondRespWithDefaults instantiates a new V0040OpenapiUsersAddCondResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiUsersAddCondRespWithDefaults() *V0040OpenapiUsersAddCondResp {
	this := V0040OpenapiUsersAddCondResp{}
	return &this
}

// GetAssociationCondition returns the AssociationCondition field value
func (o *V0040OpenapiUsersAddCondResp) GetAssociationCondition() V0040UsersAddCond {
	if o == nil {
		var ret V0040UsersAddCond
		return ret
	}

	return o.AssociationCondition
}

// GetAssociationConditionOk returns a tuple with the AssociationCondition field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0040UsersAddCond, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationCondition, true
}

// SetAssociationCondition sets field value
func (o *V0040OpenapiUsersAddCondResp) SetAssociationCondition(v V0040UsersAddCond) {
	o.AssociationCondition = v
}

// GetUser returns the User field value
func (o *V0040OpenapiUsersAddCondResp) GetUser() V0040UserShort {
	if o == nil {
		var ret V0040UserShort
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiUsersAddCondResp) GetUserOk() (*V0040UserShort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0040OpenapiUsersAddCondResp) SetUser(v V0040UserShort) {
	o.User = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiUsersAddCondResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiUsersAddCondResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiUsersAddCondResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiUsersAddCondResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiUsersAddCondResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiUsersAddCondResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiUsersAddCondResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiUsersAddCondResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiUsersAddCondResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiUsersAddCondResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiUsersAddCondResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiUsersAddCondResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiUsersAddCondResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiUsersAddCondResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["association_condition"] = o.AssociationCondition
	toSerialize["user"] = o.User
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

func (o *V0040OpenapiUsersAddCondResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"association_condition",
		"user",
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

	varV0040OpenapiUsersAddCondResp := _V0040OpenapiUsersAddCondResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiUsersAddCondResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiUsersAddCondResp(varV0040OpenapiUsersAddCondResp)

	return err
}

type NullableV0040OpenapiUsersAddCondResp struct {
	value *V0040OpenapiUsersAddCondResp
	isSet bool
}

func (v NullableV0040OpenapiUsersAddCondResp) Get() *V0040OpenapiUsersAddCondResp {
	return v.value
}

func (v *NullableV0040OpenapiUsersAddCondResp) Set(val *V0040OpenapiUsersAddCondResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiUsersAddCondResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiUsersAddCondResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiUsersAddCondResp(val *V0040OpenapiUsersAddCondResp) *NullableV0040OpenapiUsersAddCondResp {
	return &NullableV0040OpenapiUsersAddCondResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiUsersAddCondResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiUsersAddCondResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



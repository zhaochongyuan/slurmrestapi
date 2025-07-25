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

// checks if the V0042AccountsAddCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042AccountsAddCond{}

// V0042AccountsAddCond struct for V0042AccountsAddCond
type V0042AccountsAddCond struct {
	Accounts []string `json:"accounts"`
	Association *V0042AssocRecSet `json:"association,omitempty"`
	Clusters []string `json:"clusters,omitempty"`
}

type _V0042AccountsAddCond V0042AccountsAddCond

// NewV0042AccountsAddCond instantiates a new V0042AccountsAddCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042AccountsAddCond(accounts []string) *V0042AccountsAddCond {
	this := V0042AccountsAddCond{}
	this.Accounts = accounts
	return &this
}

// NewV0042AccountsAddCondWithDefaults instantiates a new V0042AccountsAddCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042AccountsAddCondWithDefaults() *V0042AccountsAddCond {
	this := V0042AccountsAddCond{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *V0042AccountsAddCond) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *V0042AccountsAddCond) GetAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *V0042AccountsAddCond) SetAccounts(v []string) {
	o.Accounts = v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *V0042AccountsAddCond) GetAssociation() V0042AssocRecSet {
	if o == nil || IsNil(o.Association) {
		var ret V0042AssocRecSet
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042AccountsAddCond) GetAssociationOk() (*V0042AssocRecSet, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *V0042AccountsAddCond) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given V0042AssocRecSet and assigns it to the Association field.
func (o *V0042AccountsAddCond) SetAssociation(v V0042AssocRecSet) {
	o.Association = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V0042AccountsAddCond) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042AccountsAddCond) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V0042AccountsAddCond) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *V0042AccountsAddCond) SetClusters(v []string) {
	o.Clusters = v
}

func (o V0042AccountsAddCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042AccountsAddCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

func (o *V0042AccountsAddCond) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accounts",
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

	varV0042AccountsAddCond := _V0042AccountsAddCond{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0042AccountsAddCond)

	if err != nil {
		return err
	}

	*o = V0042AccountsAddCond(varV0042AccountsAddCond)

	return err
}

type NullableV0042AccountsAddCond struct {
	value *V0042AccountsAddCond
	isSet bool
}

func (v NullableV0042AccountsAddCond) Get() *V0042AccountsAddCond {
	return v.value
}

func (v *NullableV0042AccountsAddCond) Set(val *V0042AccountsAddCond) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042AccountsAddCond) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042AccountsAddCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042AccountsAddCond(val *V0042AccountsAddCond) *NullableV0042AccountsAddCond {
	return &NullableV0042AccountsAddCond{value: val, isSet: true}
}

func (v NullableV0042AccountsAddCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042AccountsAddCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



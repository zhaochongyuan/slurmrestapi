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

// checks if the V0040AccountsAddCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AccountsAddCond{}

// V0040AccountsAddCond struct for V0040AccountsAddCond
type V0040AccountsAddCond struct {
	Accounts []string `json:"accounts"`
	Association *V0040AssocRecSet `json:"association,omitempty"`
	Clusters []string `json:"clusters,omitempty"`
}

type _V0040AccountsAddCond V0040AccountsAddCond

// NewV0040AccountsAddCond instantiates a new V0040AccountsAddCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AccountsAddCond(accounts []string) *V0040AccountsAddCond {
	this := V0040AccountsAddCond{}
	this.Accounts = accounts
	return &this
}

// NewV0040AccountsAddCondWithDefaults instantiates a new V0040AccountsAddCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AccountsAddCondWithDefaults() *V0040AccountsAddCond {
	this := V0040AccountsAddCond{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *V0040AccountsAddCond) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *V0040AccountsAddCond) GetAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *V0040AccountsAddCond) SetAccounts(v []string) {
	o.Accounts = v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *V0040AccountsAddCond) GetAssociation() V0040AssocRecSet {
	if o == nil || IsNil(o.Association) {
		var ret V0040AssocRecSet
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AccountsAddCond) GetAssociationOk() (*V0040AssocRecSet, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *V0040AccountsAddCond) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given V0040AssocRecSet and assigns it to the Association field.
func (o *V0040AccountsAddCond) SetAssociation(v V0040AssocRecSet) {
	o.Association = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V0040AccountsAddCond) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AccountsAddCond) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V0040AccountsAddCond) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *V0040AccountsAddCond) SetClusters(v []string) {
	o.Clusters = v
}

func (o V0040AccountsAddCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AccountsAddCond) ToMap() (map[string]interface{}, error) {
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

func (o *V0040AccountsAddCond) UnmarshalJSON(data []byte) (err error) {
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

	varV0040AccountsAddCond := _V0040AccountsAddCond{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040AccountsAddCond)

	if err != nil {
		return err
	}

	*o = V0040AccountsAddCond(varV0040AccountsAddCond)

	return err
}

type NullableV0040AccountsAddCond struct {
	value *V0040AccountsAddCond
	isSet bool
}

func (v NullableV0040AccountsAddCond) Get() *V0040AccountsAddCond {
	return v.value
}

func (v *NullableV0040AccountsAddCond) Set(val *V0040AccountsAddCond) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AccountsAddCond) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AccountsAddCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AccountsAddCond(val *V0040AccountsAddCond) *NullableV0040AccountsAddCond {
	return &NullableV0040AccountsAddCond{value: val, isSet: true}
}

func (v NullableV0040AccountsAddCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AccountsAddCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



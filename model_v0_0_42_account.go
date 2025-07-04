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

// checks if the V0042Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042Account{}

// V0042Account struct for V0042Account
type V0042Account struct {
	Associations []V0042AssocShort `json:"associations,omitempty"`
	Coordinators []V0042Coord `json:"coordinators,omitempty"`
	// Arbitrary string describing the account
	Description string `json:"description"`
	// Account name
	Name string `json:"name"`
	// Organization to which the account belongs
	Organization string `json:"organization"`
	Flags []string `json:"flags,omitempty"`
}

type _V0042Account V0042Account

// NewV0042Account instantiates a new V0042Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042Account(description string, name string, organization string) *V0042Account {
	this := V0042Account{}
	this.Description = description
	this.Name = name
	this.Organization = organization
	return &this
}

// NewV0042AccountWithDefaults instantiates a new V0042Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042AccountWithDefaults() *V0042Account {
	this := V0042Account{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0042Account) GetAssociations() []V0042AssocShort {
	if o == nil || IsNil(o.Associations) {
		var ret []V0042AssocShort
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Account) GetAssociationsOk() ([]V0042AssocShort, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0042Account) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0042AssocShort and assigns it to the Associations field.
func (o *V0042Account) SetAssociations(v []V0042AssocShort) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *V0042Account) GetCoordinators() []V0042Coord {
	if o == nil || IsNil(o.Coordinators) {
		var ret []V0042Coord
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Account) GetCoordinatorsOk() ([]V0042Coord, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *V0042Account) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []V0042Coord and assigns it to the Coordinators field.
func (o *V0042Account) SetCoordinators(v []V0042Coord) {
	o.Coordinators = v
}

// GetDescription returns the Description field value
func (o *V0042Account) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *V0042Account) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *V0042Account) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *V0042Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0042Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0042Account) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value
func (o *V0042Account) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *V0042Account) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *V0042Account) SetOrganization(v string) {
	o.Organization = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0042Account) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Account) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0042Account) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0042Account) SetFlags(v []string) {
	o.Flags = v
}

func (o V0042Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["organization"] = o.Organization
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *V0042Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"organization",
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

	varV0042Account := _V0042Account{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0042Account)

	if err != nil {
		return err
	}

	*o = V0042Account(varV0042Account)

	return err
}

type NullableV0042Account struct {
	value *V0042Account
	isSet bool
}

func (v NullableV0042Account) Get() *V0042Account {
	return v.value
}

func (v *NullableV0042Account) Set(val *V0042Account) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042Account) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042Account) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042Account(val *V0042Account) *NullableV0042Account {
	return &NullableV0042Account{value: val, isSet: true}
}

func (v NullableV0042Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042Account) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



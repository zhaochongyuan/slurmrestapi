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

// checks if the V0041OpenapiSlurmdbdConfigRespUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespUsersInner{}

// V0041OpenapiSlurmdbdConfigRespUsersInner struct for V0041OpenapiSlurmdbdConfigRespUsersInner
type V0041OpenapiSlurmdbdConfigRespUsersInner struct {
	// AdminLevel granted to the user
	AdministratorLevel []string `json:"administrator_level,omitempty"`
	// Associations created for this user
	Associations []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner `json:"associations,omitempty"`
	// Accounts this user is a coordinator for
	Coordinators []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner `json:"coordinators,omitempty"`
	Default *V0040UserDefault `json:"default,omitempty"`
	// Flags associated with user
	Flags []string `json:"flags,omitempty"`
	// User name
	Name string `json:"name"`
	// Previous user name
	OldName *string `json:"old_name,omitempty"`
	// List of available WCKeys
	Wckeys []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner `json:"wckeys,omitempty"`
}

type _V0041OpenapiSlurmdbdConfigRespUsersInner V0041OpenapiSlurmdbdConfigRespUsersInner

// NewV0041OpenapiSlurmdbdConfigRespUsersInner instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespUsersInner(name string) *V0041OpenapiSlurmdbdConfigRespUsersInner {
	this := V0041OpenapiSlurmdbdConfigRespUsersInner{}
	this.Name = name
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespUsersInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespUsersInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespUsersInner {
	this := V0041OpenapiSlurmdbdConfigRespUsersInner{}
	return &this
}

// GetAdministratorLevel returns the AdministratorLevel field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAdministratorLevel() []string {
	if o == nil || IsNil(o.AdministratorLevel) {
		var ret []string
		return ret
	}
	return o.AdministratorLevel
}

// GetAdministratorLevelOk returns a tuple with the AdministratorLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAdministratorLevelOk() ([]string, bool) {
	if o == nil || IsNil(o.AdministratorLevel) {
		return nil, false
	}
	return o.AdministratorLevel, true
}

// HasAdministratorLevel returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasAdministratorLevel() bool {
	if o != nil && !IsNil(o.AdministratorLevel) {
		return true
	}

	return false
}

// SetAdministratorLevel gets a reference to the given []string and assigns it to the AdministratorLevel field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetAdministratorLevel(v []string) {
	o.AdministratorLevel = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAssociations() []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner {
	if o == nil || IsNil(o.Associations) {
		var ret []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetAssociationsOk() ([]V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner and assigns it to the Associations field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetAssociations(v []V0041OpenapiSlurmdbdConfigRespAccountsInnerAssociationsInner) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetCoordinators() []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner {
	if o == nil || IsNil(o.Coordinators) {
		var ret []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetCoordinatorsOk() ([]V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner and assigns it to the Coordinators field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetCoordinators(v []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner) {
	o.Coordinators = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetDefault() V0040UserDefault {
	if o == nil || IsNil(o.Default) {
		var ret V0040UserDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetDefaultOk() (*V0040UserDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given V0040UserDefault and assigns it to the Default field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetDefault(v V0040UserDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetName(v string) {
	o.Name = v
}

// GetOldName returns the OldName field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetOldName() string {
	if o == nil || IsNil(o.OldName) {
		var ret string
		return ret
	}
	return *o.OldName
}

// GetOldNameOk returns a tuple with the OldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetOldNameOk() (*string, bool) {
	if o == nil || IsNil(o.OldName) {
		return nil, false
	}
	return o.OldName, true
}

// HasOldName returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasOldName() bool {
	if o != nil && !IsNil(o.OldName) {
		return true
	}

	return false
}

// SetOldName gets a reference to the given string and assigns it to the OldName field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetOldName(v string) {
	o.OldName = &v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetWckeys() []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner {
	if o == nil || IsNil(o.Wckeys) {
		var ret []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) GetWckeysOk() ([]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner and assigns it to the Wckeys field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) SetWckeys(v []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) {
	o.Wckeys = v
}

func (o V0041OpenapiSlurmdbdConfigRespUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdministratorLevel) {
		toSerialize["administrator_level"] = o.AdministratorLevel
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OldName) {
		toSerialize["old_name"] = o.OldName
	}
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

func (o *V0041OpenapiSlurmdbdConfigRespUsersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varV0041OpenapiSlurmdbdConfigRespUsersInner := _V0041OpenapiSlurmdbdConfigRespUsersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiSlurmdbdConfigRespUsersInner)

	if err != nil {
		return err
	}

	*o = V0041OpenapiSlurmdbdConfigRespUsersInner(varV0041OpenapiSlurmdbdConfigRespUsersInner)

	return err
}

type NullableV0041OpenapiSlurmdbdConfigRespUsersInner struct {
	value *V0041OpenapiSlurmdbdConfigRespUsersInner
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInner) Get() *V0041OpenapiSlurmdbdConfigRespUsersInner {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInner) Set(val *V0041OpenapiSlurmdbdConfigRespUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespUsersInner(val *V0041OpenapiSlurmdbdConfigRespUsersInner) *NullableV0041OpenapiSlurmdbdConfigRespUsersInner {
	return &NullableV0041OpenapiSlurmdbdConfigRespUsersInner{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



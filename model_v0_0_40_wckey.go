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

// checks if the V0040Wckey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040Wckey{}

// V0040Wckey struct for V0040Wckey
type V0040Wckey struct {
	Accounting []V0040Accounting `json:"accounting,omitempty"`
	// Cluster name
	Cluster string `json:"cluster"`
	// Unique ID for this user-cluster-wckey combination
	Id *int32 `json:"id,omitempty"`
	// WCKey name
	Name string `json:"name"`
	// User name
	User string `json:"user"`
	Flags []string `json:"flags,omitempty"`
}

type _V0040Wckey V0040Wckey

// NewV0040Wckey instantiates a new V0040Wckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040Wckey(cluster string, name string, user string) *V0040Wckey {
	this := V0040Wckey{}
	this.Cluster = cluster
	this.Name = name
	this.User = user
	return &this
}

// NewV0040WckeyWithDefaults instantiates a new V0040Wckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040WckeyWithDefaults() *V0040Wckey {
	this := V0040Wckey{}
	return &this
}

// GetAccounting returns the Accounting field value if set, zero value otherwise.
func (o *V0040Wckey) GetAccounting() []V0040Accounting {
	if o == nil || IsNil(o.Accounting) {
		var ret []V0040Accounting
		return ret
	}
	return o.Accounting
}

// GetAccountingOk returns a tuple with the Accounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetAccountingOk() ([]V0040Accounting, bool) {
	if o == nil || IsNil(o.Accounting) {
		return nil, false
	}
	return o.Accounting, true
}

// HasAccounting returns a boolean if a field has been set.
func (o *V0040Wckey) HasAccounting() bool {
	if o != nil && !IsNil(o.Accounting) {
		return true
	}

	return false
}

// SetAccounting gets a reference to the given []V0040Accounting and assigns it to the Accounting field.
func (o *V0040Wckey) SetAccounting(v []V0040Accounting) {
	o.Accounting = v
}

// GetCluster returns the Cluster field value
func (o *V0040Wckey) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *V0040Wckey) SetCluster(v string) {
	o.Cluster = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0040Wckey) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0040Wckey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0040Wckey) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *V0040Wckey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0040Wckey) SetName(v string) {
	o.Name = v
}

// GetUser returns the User field value
func (o *V0040Wckey) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0040Wckey) SetUser(v string) {
	o.User = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0040Wckey) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040Wckey) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0040Wckey) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0040Wckey) SetFlags(v []string) {
	o.Flags = v
}

func (o V0040Wckey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040Wckey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounting) {
		toSerialize["accounting"] = o.Accounting
	}
	toSerialize["cluster"] = o.Cluster
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["user"] = o.User
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *V0040Wckey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster",
		"name",
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

	varV0040Wckey := _V0040Wckey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040Wckey)

	if err != nil {
		return err
	}

	*o = V0040Wckey(varV0040Wckey)

	return err
}

type NullableV0040Wckey struct {
	value *V0040Wckey
	isSet bool
}

func (v NullableV0040Wckey) Get() *V0040Wckey {
	return v.value
}

func (v *NullableV0040Wckey) Set(val *V0040Wckey) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040Wckey) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040Wckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040Wckey(val *V0040Wckey) *NullableV0040Wckey {
	return &NullableV0040Wckey{value: val, isSet: true}
}

func (v NullableV0040Wckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040Wckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the V0042Assoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042Assoc{}

// V0042Assoc struct for V0042Assoc
type V0042Assoc struct {
	Accounting []V0042Accounting `json:"accounting,omitempty"`
	// Account name
	Account *string `json:"account,omitempty"`
	// Cluster name
	Cluster *string `json:"cluster,omitempty"`
	// Arbitrary comment
	Comment *string `json:"comment,omitempty"`
	Default *V0040AssocDefault `json:"default,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Max *V0042AssocMax `json:"max,omitempty"`
	// Unique ID
	Id *int32 `json:"id,omitempty"`
	// Is default association for user
	IsDefault *bool `json:"is_default,omitempty"`
	// Complete path up the hierarchy to the root association
	Lineage *string `json:"lineage,omitempty"`
	Min *V0042AssocMin `json:"min,omitempty"`
	// Name of parent account
	ParentAccount *string `json:"parent_account,omitempty"`
	// Partition name
	Partition *string `json:"partition,omitempty"`
	Priority *V0042Uint32NoValStruct `json:"priority,omitempty"`
	// List of QOS names
	Qos []string `json:"qos,omitempty"`
	// Allocated shares used for fairshare calculation
	SharesRaw *int32 `json:"shares_raw,omitempty"`
	// User name
	User string `json:"user"`
}

type _V0042Assoc V0042Assoc

// NewV0042Assoc instantiates a new V0042Assoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042Assoc(user string) *V0042Assoc {
	this := V0042Assoc{}
	this.User = user
	return &this
}

// NewV0042AssocWithDefaults instantiates a new V0042Assoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042AssocWithDefaults() *V0042Assoc {
	this := V0042Assoc{}
	return &this
}

// GetAccounting returns the Accounting field value if set, zero value otherwise.
func (o *V0042Assoc) GetAccounting() []V0042Accounting {
	if o == nil || IsNil(o.Accounting) {
		var ret []V0042Accounting
		return ret
	}
	return o.Accounting
}

// GetAccountingOk returns a tuple with the Accounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetAccountingOk() ([]V0042Accounting, bool) {
	if o == nil || IsNil(o.Accounting) {
		return nil, false
	}
	return o.Accounting, true
}

// HasAccounting returns a boolean if a field has been set.
func (o *V0042Assoc) HasAccounting() bool {
	if o != nil && !IsNil(o.Accounting) {
		return true
	}

	return false
}

// SetAccounting gets a reference to the given []V0042Accounting and assigns it to the Accounting field.
func (o *V0042Assoc) SetAccounting(v []V0042Accounting) {
	o.Accounting = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0042Assoc) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0042Assoc) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *V0042Assoc) SetAccount(v string) {
	o.Account = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V0042Assoc) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V0042Assoc) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *V0042Assoc) SetCluster(v string) {
	o.Cluster = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V0042Assoc) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V0042Assoc) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V0042Assoc) SetComment(v string) {
	o.Comment = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V0042Assoc) GetDefault() V0040AssocDefault {
	if o == nil || IsNil(o.Default) {
		var ret V0040AssocDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetDefaultOk() (*V0040AssocDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V0042Assoc) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given V0040AssocDefault and assigns it to the Default field.
func (o *V0042Assoc) SetDefault(v V0040AssocDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0042Assoc) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0042Assoc) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0042Assoc) SetFlags(v []string) {
	o.Flags = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V0042Assoc) GetMax() V0042AssocMax {
	if o == nil || IsNil(o.Max) {
		var ret V0042AssocMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetMaxOk() (*V0042AssocMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V0042Assoc) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given V0042AssocMax and assigns it to the Max field.
func (o *V0042Assoc) SetMax(v V0042AssocMax) {
	o.Max = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0042Assoc) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0042Assoc) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0042Assoc) SetId(v int32) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *V0042Assoc) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *V0042Assoc) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *V0042Assoc) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *V0042Assoc) GetLineage() string {
	if o == nil || IsNil(o.Lineage) {
		var ret string
		return ret
	}
	return *o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetLineageOk() (*string, bool) {
	if o == nil || IsNil(o.Lineage) {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *V0042Assoc) HasLineage() bool {
	if o != nil && !IsNil(o.Lineage) {
		return true
	}

	return false
}

// SetLineage gets a reference to the given string and assigns it to the Lineage field.
func (o *V0042Assoc) SetLineage(v string) {
	o.Lineage = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0042Assoc) GetMin() V0042AssocMin {
	if o == nil || IsNil(o.Min) {
		var ret V0042AssocMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetMinOk() (*V0042AssocMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0042Assoc) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given V0042AssocMin and assigns it to the Min field.
func (o *V0042Assoc) SetMin(v V0042AssocMin) {
	o.Min = &v
}

// GetParentAccount returns the ParentAccount field value if set, zero value otherwise.
func (o *V0042Assoc) GetParentAccount() string {
	if o == nil || IsNil(o.ParentAccount) {
		var ret string
		return ret
	}
	return *o.ParentAccount
}

// GetParentAccountOk returns a tuple with the ParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetParentAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ParentAccount) {
		return nil, false
	}
	return o.ParentAccount, true
}

// HasParentAccount returns a boolean if a field has been set.
func (o *V0042Assoc) HasParentAccount() bool {
	if o != nil && !IsNil(o.ParentAccount) {
		return true
	}

	return false
}

// SetParentAccount gets a reference to the given string and assigns it to the ParentAccount field.
func (o *V0042Assoc) SetParentAccount(v string) {
	o.ParentAccount = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0042Assoc) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0042Assoc) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0042Assoc) SetPartition(v string) {
	o.Partition = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V0042Assoc) GetPriority() V0042Uint32NoValStruct {
	if o == nil || IsNil(o.Priority) {
		var ret V0042Uint32NoValStruct
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetPriorityOk() (*V0042Uint32NoValStruct, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V0042Assoc) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given V0042Uint32NoValStruct and assigns it to the Priority field.
func (o *V0042Assoc) SetPriority(v V0042Uint32NoValStruct) {
	o.Priority = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0042Assoc) GetQos() []string {
	if o == nil || IsNil(o.Qos) {
		var ret []string
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetQosOk() ([]string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0042Assoc) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []string and assigns it to the Qos field.
func (o *V0042Assoc) SetQos(v []string) {
	o.Qos = v
}

// GetSharesRaw returns the SharesRaw field value if set, zero value otherwise.
func (o *V0042Assoc) GetSharesRaw() int32 {
	if o == nil || IsNil(o.SharesRaw) {
		var ret int32
		return ret
	}
	return *o.SharesRaw
}

// GetSharesRawOk returns a tuple with the SharesRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetSharesRawOk() (*int32, bool) {
	if o == nil || IsNil(o.SharesRaw) {
		return nil, false
	}
	return o.SharesRaw, true
}

// HasSharesRaw returns a boolean if a field has been set.
func (o *V0042Assoc) HasSharesRaw() bool {
	if o != nil && !IsNil(o.SharesRaw) {
		return true
	}

	return false
}

// SetSharesRaw gets a reference to the given int32 and assigns it to the SharesRaw field.
func (o *V0042Assoc) SetSharesRaw(v int32) {
	o.SharesRaw = &v
}

// GetUser returns the User field value
func (o *V0042Assoc) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0042Assoc) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0042Assoc) SetUser(v string) {
	o.User = v
}

func (o V0042Assoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042Assoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounting) {
		toSerialize["accounting"] = o.Accounting
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.Lineage) {
		toSerialize["lineage"] = o.Lineage
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.ParentAccount) {
		toSerialize["parent_account"] = o.ParentAccount
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.SharesRaw) {
		toSerialize["shares_raw"] = o.SharesRaw
	}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *V0042Assoc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varV0042Assoc := _V0042Assoc{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0042Assoc)

	if err != nil {
		return err
	}

	*o = V0042Assoc(varV0042Assoc)

	return err
}

type NullableV0042Assoc struct {
	value *V0042Assoc
	isSet bool
}

func (v NullableV0042Assoc) Get() *V0042Assoc {
	return v.value
}

func (v *NullableV0042Assoc) Set(val *V0042Assoc) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042Assoc) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042Assoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042Assoc(val *V0042Assoc) *NullableV0042Assoc {
	return &NullableV0042Assoc{value: val, isSet: true}
}

func (v NullableV0042Assoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042Assoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



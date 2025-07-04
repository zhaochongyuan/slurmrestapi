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

// checks if the V0041OpenapiJobAllocResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobAllocResp{}

// V0041OpenapiJobAllocResp struct for V0041OpenapiJobAllocResp
type V0041OpenapiJobAllocResp struct {
	// Submitted Job ID
	JobId *int32 `json:"job_id,omitempty"`
	// Job submission user message
	JobSubmitUserMsg *string `json:"job_submit_user_msg,omitempty"`
	Meta *V0041OpenapiSharesRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSharesRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSharesRespWarningsInner `json:"warnings,omitempty"`
}

// NewV0041OpenapiJobAllocResp instantiates a new V0041OpenapiJobAllocResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobAllocResp() *V0041OpenapiJobAllocResp {
	this := V0041OpenapiJobAllocResp{}
	return &this
}

// NewV0041OpenapiJobAllocRespWithDefaults instantiates a new V0041OpenapiJobAllocResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobAllocRespWithDefaults() *V0041OpenapiJobAllocResp {
	this := V0041OpenapiJobAllocResp{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0041OpenapiJobAllocResp) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobAllocResp) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0041OpenapiJobAllocResp) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0041OpenapiJobAllocResp) SetJobId(v int32) {
	o.JobId = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *V0041OpenapiJobAllocResp) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobAllocResp) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *V0041OpenapiJobAllocResp) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *V0041OpenapiJobAllocResp) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiJobAllocResp) GetMeta() V0041OpenapiSharesRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSharesRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobAllocResp) GetMetaOk() (*V0041OpenapiSharesRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiJobAllocResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSharesRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiJobAllocResp) SetMeta(v V0041OpenapiSharesRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiJobAllocResp) GetErrors() []V0041OpenapiSharesRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSharesRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobAllocResp) GetErrorsOk() ([]V0041OpenapiSharesRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiJobAllocResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSharesRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiJobAllocResp) SetErrors(v []V0041OpenapiSharesRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiJobAllocResp) GetWarnings() []V0041OpenapiSharesRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSharesRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobAllocResp) GetWarningsOk() ([]V0041OpenapiSharesRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiJobAllocResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSharesRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiJobAllocResp) SetWarnings(v []V0041OpenapiSharesRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiJobAllocResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobAllocResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
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

type NullableV0041OpenapiJobAllocResp struct {
	value *V0041OpenapiJobAllocResp
	isSet bool
}

func (v NullableV0041OpenapiJobAllocResp) Get() *V0041OpenapiJobAllocResp {
	return v.value
}

func (v *NullableV0041OpenapiJobAllocResp) Set(val *V0041OpenapiJobAllocResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobAllocResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobAllocResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobAllocResp(val *V0041OpenapiJobAllocResp) *NullableV0041OpenapiJobAllocResp {
	return &NullableV0041OpenapiJobAllocResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobAllocResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobAllocResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



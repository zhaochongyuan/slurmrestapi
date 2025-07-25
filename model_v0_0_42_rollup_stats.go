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

// checks if the V0042RollupStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042RollupStats{}

// V0042RollupStats struct for V0042RollupStats
type V0042RollupStats struct {
	Hourly *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly `json:"hourly,omitempty"`
	Daily *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily `json:"daily,omitempty"`
	Monthly *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly `json:"monthly,omitempty"`
}

// NewV0042RollupStats instantiates a new V0042RollupStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042RollupStats() *V0042RollupStats {
	this := V0042RollupStats{}
	return &this
}

// NewV0042RollupStatsWithDefaults instantiates a new V0042RollupStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042RollupStatsWithDefaults() *V0042RollupStats {
	this := V0042RollupStats{}
	return &this
}

// GetHourly returns the Hourly field value if set, zero value otherwise.
func (o *V0042RollupStats) GetHourly() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly {
	if o == nil || IsNil(o.Hourly) {
		var ret V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly
		return ret
	}
	return *o.Hourly
}

// GetHourlyOk returns a tuple with the Hourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042RollupStats) GetHourlyOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly, bool) {
	if o == nil || IsNil(o.Hourly) {
		return nil, false
	}
	return o.Hourly, true
}

// HasHourly returns a boolean if a field has been set.
func (o *V0042RollupStats) HasHourly() bool {
	if o != nil && !IsNil(o.Hourly) {
		return true
	}

	return false
}

// SetHourly gets a reference to the given V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly and assigns it to the Hourly field.
func (o *V0042RollupStats) SetHourly(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) {
	o.Hourly = &v
}

// GetDaily returns the Daily field value if set, zero value otherwise.
func (o *V0042RollupStats) GetDaily() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily {
	if o == nil || IsNil(o.Daily) {
		var ret V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily
		return ret
	}
	return *o.Daily
}

// GetDailyOk returns a tuple with the Daily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042RollupStats) GetDailyOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily, bool) {
	if o == nil || IsNil(o.Daily) {
		return nil, false
	}
	return o.Daily, true
}

// HasDaily returns a boolean if a field has been set.
func (o *V0042RollupStats) HasDaily() bool {
	if o != nil && !IsNil(o.Daily) {
		return true
	}

	return false
}

// SetDaily gets a reference to the given V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily and assigns it to the Daily field.
func (o *V0042RollupStats) SetDaily(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) {
	o.Daily = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *V0042RollupStats) GetMonthly() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly {
	if o == nil || IsNil(o.Monthly) {
		var ret V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042RollupStats) GetMonthlyOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *V0042RollupStats) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly and assigns it to the Monthly field.
func (o *V0042RollupStats) SetMonthly(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) {
	o.Monthly = &v
}

func (o V0042RollupStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042RollupStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hourly) {
		toSerialize["hourly"] = o.Hourly
	}
	if !IsNil(o.Daily) {
		toSerialize["daily"] = o.Daily
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableV0042RollupStats struct {
	value *V0042RollupStats
	isSet bool
}

func (v NullableV0042RollupStats) Get() *V0042RollupStats {
	return v.value
}

func (v *NullableV0042RollupStats) Set(val *V0042RollupStats) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042RollupStats) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042RollupStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042RollupStats(val *V0042RollupStats) *NullableV0042RollupStats {
	return &NullableV0042RollupStats{value: val, isSet: true}
}

func (v NullableV0042RollupStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042RollupStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



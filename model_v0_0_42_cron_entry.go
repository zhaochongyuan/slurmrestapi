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

// checks if the V0042CronEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042CronEntry{}

// V0042CronEntry struct for V0042CronEntry
type V0042CronEntry struct {
	Flags []string `json:"flags,omitempty"`
	// Ranged string specifying eligible minute values (e.g. 0-10,50)
	Minute *string `json:"minute,omitempty"`
	// Ranged string specifying eligible hour values (e.g. 0-5,23)
	Hour *string `json:"hour,omitempty"`
	// Ranged string specifying eligible day of month values (e.g. 0-10,29)
	DayOfMonth *string `json:"day_of_month,omitempty"`
	// Ranged string specifying eligible month values (e.g. 0-5,12)
	Month *string `json:"month,omitempty"`
	// Ranged string specifying eligible day of week values (e.g.0-3,7)
	DayOfWeek *string `json:"day_of_week,omitempty"`
	// Complete time specification (* means valid for all allowed values) - minute hour day_of_month month day_of_week
	Specification *string `json:"specification,omitempty"`
	// Command to run
	Command *string `json:"command,omitempty"`
	Line *V0040CronEntryLine `json:"line,omitempty"`
}

// NewV0042CronEntry instantiates a new V0042CronEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042CronEntry() *V0042CronEntry {
	this := V0042CronEntry{}
	return &this
}

// NewV0042CronEntryWithDefaults instantiates a new V0042CronEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042CronEntryWithDefaults() *V0042CronEntry {
	this := V0042CronEntry{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0042CronEntry) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0042CronEntry) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0042CronEntry) SetFlags(v []string) {
	o.Flags = v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *V0042CronEntry) GetMinute() string {
	if o == nil || IsNil(o.Minute) {
		var ret string
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetMinuteOk() (*string, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *V0042CronEntry) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given string and assigns it to the Minute field.
func (o *V0042CronEntry) SetMinute(v string) {
	o.Minute = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *V0042CronEntry) GetHour() string {
	if o == nil || IsNil(o.Hour) {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetHourOk() (*string, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *V0042CronEntry) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *V0042CronEntry) SetHour(v string) {
	o.Hour = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *V0042CronEntry) GetDayOfMonth() string {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret string
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetDayOfMonthOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *V0042CronEntry) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given string and assigns it to the DayOfMonth field.
func (o *V0042CronEntry) SetDayOfMonth(v string) {
	o.DayOfMonth = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *V0042CronEntry) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *V0042CronEntry) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *V0042CronEntry) SetMonth(v string) {
	o.Month = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *V0042CronEntry) GetDayOfWeek() string {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *V0042CronEntry) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *V0042CronEntry) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *V0042CronEntry) GetSpecification() string {
	if o == nil || IsNil(o.Specification) {
		var ret string
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetSpecificationOk() (*string, bool) {
	if o == nil || IsNil(o.Specification) {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *V0042CronEntry) HasSpecification() bool {
	if o != nil && !IsNil(o.Specification) {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given string and assigns it to the Specification field.
func (o *V0042CronEntry) SetSpecification(v string) {
	o.Specification = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V0042CronEntry) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V0042CronEntry) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *V0042CronEntry) SetCommand(v string) {
	o.Command = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *V0042CronEntry) GetLine() V0040CronEntryLine {
	if o == nil || IsNil(o.Line) {
		var ret V0040CronEntryLine
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042CronEntry) GetLineOk() (*V0040CronEntryLine, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *V0042CronEntry) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given V0040CronEntryLine and assigns it to the Line field.
func (o *V0042CronEntry) SetLine(v V0040CronEntryLine) {
	o.Line = &v
}

func (o V0042CronEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042CronEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.DayOfMonth) {
		toSerialize["day_of_month"] = o.DayOfMonth
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.DayOfWeek) {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if !IsNil(o.Specification) {
		toSerialize["specification"] = o.Specification
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	return toSerialize, nil
}

type NullableV0042CronEntry struct {
	value *V0042CronEntry
	isSet bool
}

func (v NullableV0042CronEntry) Get() *V0042CronEntry {
	return v.value
}

func (v *NullableV0042CronEntry) Set(val *V0042CronEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042CronEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042CronEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042CronEntry(val *V0042CronEntry) *NullableV0042CronEntry {
	return &NullableV0042CronEntry{value: val, isSet: true}
}

func (v NullableV0042CronEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042CronEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



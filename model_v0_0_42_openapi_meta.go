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

// checks if the V0042OpenapiMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0042OpenapiMeta{}

// V0042OpenapiMeta struct for V0042OpenapiMeta
type V0042OpenapiMeta struct {
	Plugin *V0040OpenapiMetaPlugin `json:"plugin,omitempty"`
	Client *V0040OpenapiMetaClient `json:"client,omitempty"`
	Command []string `json:"command,omitempty"`
	Slurm *V0040OpenapiMetaSlurm `json:"slurm,omitempty"`
}

// NewV0042OpenapiMeta instantiates a new V0042OpenapiMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0042OpenapiMeta() *V0042OpenapiMeta {
	this := V0042OpenapiMeta{}
	return &this
}

// NewV0042OpenapiMetaWithDefaults instantiates a new V0042OpenapiMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0042OpenapiMetaWithDefaults() *V0042OpenapiMeta {
	this := V0042OpenapiMeta{}
	return &this
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *V0042OpenapiMeta) GetPlugin() V0040OpenapiMetaPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret V0040OpenapiMetaPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiMeta) GetPluginOk() (*V0040OpenapiMetaPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *V0042OpenapiMeta) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given V0040OpenapiMetaPlugin and assigns it to the Plugin field.
func (o *V0042OpenapiMeta) SetPlugin(v V0040OpenapiMetaPlugin) {
	o.Plugin = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *V0042OpenapiMeta) GetClient() V0040OpenapiMetaClient {
	if o == nil || IsNil(o.Client) {
		var ret V0040OpenapiMetaClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiMeta) GetClientOk() (*V0040OpenapiMetaClient, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *V0042OpenapiMeta) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given V0040OpenapiMetaClient and assigns it to the Client field.
func (o *V0042OpenapiMeta) SetClient(v V0040OpenapiMetaClient) {
	o.Client = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V0042OpenapiMeta) GetCommand() []string {
	if o == nil || IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiMeta) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V0042OpenapiMeta) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *V0042OpenapiMeta) SetCommand(v []string) {
	o.Command = v
}

// GetSlurm returns the Slurm field value if set, zero value otherwise.
func (o *V0042OpenapiMeta) GetSlurm() V0040OpenapiMetaSlurm {
	if o == nil || IsNil(o.Slurm) {
		var ret V0040OpenapiMetaSlurm
		return ret
	}
	return *o.Slurm
}

// GetSlurmOk returns a tuple with the Slurm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0042OpenapiMeta) GetSlurmOk() (*V0040OpenapiMetaSlurm, bool) {
	if o == nil || IsNil(o.Slurm) {
		return nil, false
	}
	return o.Slurm, true
}

// HasSlurm returns a boolean if a field has been set.
func (o *V0042OpenapiMeta) HasSlurm() bool {
	if o != nil && !IsNil(o.Slurm) {
		return true
	}

	return false
}

// SetSlurm gets a reference to the given V0040OpenapiMetaSlurm and assigns it to the Slurm field.
func (o *V0042OpenapiMeta) SetSlurm(v V0040OpenapiMetaSlurm) {
	o.Slurm = &v
}

func (o V0042OpenapiMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0042OpenapiMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Slurm) {
		toSerialize["slurm"] = o.Slurm
	}
	return toSerialize, nil
}

type NullableV0042OpenapiMeta struct {
	value *V0042OpenapiMeta
	isSet bool
}

func (v NullableV0042OpenapiMeta) Get() *V0042OpenapiMeta {
	return v.value
}

func (v *NullableV0042OpenapiMeta) Set(val *V0042OpenapiMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableV0042OpenapiMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableV0042OpenapiMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0042OpenapiMeta(val *V0042OpenapiMeta) *NullableV0042OpenapiMeta {
	return &NullableV0042OpenapiMeta{value: val, isSet: true}
}

func (v NullableV0042OpenapiMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0042OpenapiMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.8-14-gbb23763d3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the OutputID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputID{}

// OutputID struct for OutputID
type OutputID struct {
	// The output ID
	OutputId string `json:"outputId"`
}

// NewOutputID instantiates a new OutputID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputID(outputId string) *OutputID {
	this := OutputID{}
	this.OutputId = outputId
	return &this
}

// NewOutputIDWithDefaults instantiates a new OutputID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputIDWithDefaults() *OutputID {
	this := OutputID{}
	return &this
}

// GetOutputId returns the OutputId field value
func (o *OutputID) GetOutputId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputId
}

// GetOutputIdOk returns a tuple with the OutputId field value
// and a boolean to check if the value has been set.
func (o *OutputID) GetOutputIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputId, true
}

// SetOutputId sets field value
func (o *OutputID) SetOutputId(v string) {
	o.OutputId = v
}

func (o OutputID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["outputId"] = o.OutputId
	return toSerialize, nil
}

type NullableOutputID struct {
	value *OutputID
	isSet bool
}

func (v NullableOutputID) Get() *OutputID {
	return v.value
}

func (v *NullableOutputID) Set(val *OutputID) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputID) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputID(val *OutputID) *NullableOutputID {
	return &NullableOutputID{value: val, isSet: true}
}

func (v NullableOutputID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



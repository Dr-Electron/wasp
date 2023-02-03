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

// checks if the PeeringNodeIdentityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeeringNodeIdentityResponse{}

// PeeringNodeIdentityResponse struct for PeeringNodeIdentityResponse
type PeeringNodeIdentityResponse struct {
	IsTrusted bool `json:"isTrusted"`
	// The NetID of the peer
	NetId string `json:"netId"`
	// The peers public key encoded in Hex
	PublicKey string `json:"publicKey"`
}

// NewPeeringNodeIdentityResponse instantiates a new PeeringNodeIdentityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeeringNodeIdentityResponse(isTrusted bool, netId string, publicKey string) *PeeringNodeIdentityResponse {
	this := PeeringNodeIdentityResponse{}
	this.IsTrusted = isTrusted
	this.NetId = netId
	this.PublicKey = publicKey
	return &this
}

// NewPeeringNodeIdentityResponseWithDefaults instantiates a new PeeringNodeIdentityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeeringNodeIdentityResponseWithDefaults() *PeeringNodeIdentityResponse {
	this := PeeringNodeIdentityResponse{}
	return &this
}

// GetIsTrusted returns the IsTrusted field value
func (o *PeeringNodeIdentityResponse) GetIsTrusted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTrusted
}

// GetIsTrustedOk returns a tuple with the IsTrusted field value
// and a boolean to check if the value has been set.
func (o *PeeringNodeIdentityResponse) GetIsTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTrusted, true
}

// SetIsTrusted sets field value
func (o *PeeringNodeIdentityResponse) SetIsTrusted(v bool) {
	o.IsTrusted = v
}

// GetNetId returns the NetId field value
func (o *PeeringNodeIdentityResponse) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *PeeringNodeIdentityResponse) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *PeeringNodeIdentityResponse) SetNetId(v string) {
	o.NetId = v
}

// GetPublicKey returns the PublicKey field value
func (o *PeeringNodeIdentityResponse) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *PeeringNodeIdentityResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *PeeringNodeIdentityResponse) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o PeeringNodeIdentityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeeringNodeIdentityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isTrusted"] = o.IsTrusted
	toSerialize["netId"] = o.NetId
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

type NullablePeeringNodeIdentityResponse struct {
	value *PeeringNodeIdentityResponse
	isSet bool
}

func (v NullablePeeringNodeIdentityResponse) Get() *PeeringNodeIdentityResponse {
	return v.value
}

func (v *NullablePeeringNodeIdentityResponse) Set(val *PeeringNodeIdentityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePeeringNodeIdentityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePeeringNodeIdentityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeeringNodeIdentityResponse(val *PeeringNodeIdentityResponse) *NullablePeeringNodeIdentityResponse {
	return &NullablePeeringNodeIdentityResponse{value: val, isSet: true}
}

func (v NullablePeeringNodeIdentityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeeringNodeIdentityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



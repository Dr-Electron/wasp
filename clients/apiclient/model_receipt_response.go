/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ReceiptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiptResponse{}

// ReceiptResponse struct for ReceiptResponse
type ReceiptResponse struct {
	BlockIndex uint32 `json:"blockIndex"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The gas budget (uint64 as string)
	GasBudget string `json:"gasBudget"`
	GasBurnLog []BurnRecord `json:"gasBurnLog"`
	// The burned gas (uint64 as string)
	GasBurned string `json:"gasBurned"`
	// The charged gas fee (uint64 as string)
	GasFeeCharged string `json:"gasFeeCharged"`
	RawError *UnresolvedVMErrorJSON `json:"rawError,omitempty"`
	Request string `json:"request"`
	RequestIndex uint32 `json:"requestIndex"`
}

// NewReceiptResponse instantiates a new ReceiptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptResponse(blockIndex uint32, gasBudget string, gasBurnLog []BurnRecord, gasBurned string, gasFeeCharged string, request string, requestIndex uint32) *ReceiptResponse {
	this := ReceiptResponse{}
	this.BlockIndex = blockIndex
	this.GasBudget = gasBudget
	this.GasBurnLog = gasBurnLog
	this.GasBurned = gasBurned
	this.GasFeeCharged = gasFeeCharged
	this.Request = request
	this.RequestIndex = requestIndex
	return &this
}

// NewReceiptResponseWithDefaults instantiates a new ReceiptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptResponseWithDefaults() *ReceiptResponse {
	this := ReceiptResponse{}
	return &this
}

// GetBlockIndex returns the BlockIndex field value
func (o *ReceiptResponse) GetBlockIndex() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.BlockIndex
}

// GetBlockIndexOk returns a tuple with the BlockIndex field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetBlockIndexOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIndex, true
}

// SetBlockIndex sets field value
func (o *ReceiptResponse) SetBlockIndex(v uint32) {
	o.BlockIndex = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ReceiptResponse) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ReceiptResponse) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ReceiptResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetGasBudget returns the GasBudget field value
func (o *ReceiptResponse) GetGasBudget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasBudget
}

// GetGasBudgetOk returns a tuple with the GasBudget field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetGasBudgetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasBudget, true
}

// SetGasBudget sets field value
func (o *ReceiptResponse) SetGasBudget(v string) {
	o.GasBudget = v
}

// GetGasBurnLog returns the GasBurnLog field value
func (o *ReceiptResponse) GetGasBurnLog() []BurnRecord {
	if o == nil {
		var ret []BurnRecord
		return ret
	}

	return o.GasBurnLog
}

// GetGasBurnLogOk returns a tuple with the GasBurnLog field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetGasBurnLogOk() ([]BurnRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasBurnLog, true
}

// SetGasBurnLog sets field value
func (o *ReceiptResponse) SetGasBurnLog(v []BurnRecord) {
	o.GasBurnLog = v
}

// GetGasBurned returns the GasBurned field value
func (o *ReceiptResponse) GetGasBurned() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasBurned
}

// GetGasBurnedOk returns a tuple with the GasBurned field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetGasBurnedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasBurned, true
}

// SetGasBurned sets field value
func (o *ReceiptResponse) SetGasBurned(v string) {
	o.GasBurned = v
}

// GetGasFeeCharged returns the GasFeeCharged field value
func (o *ReceiptResponse) GetGasFeeCharged() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasFeeCharged
}

// GetGasFeeChargedOk returns a tuple with the GasFeeCharged field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetGasFeeChargedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasFeeCharged, true
}

// SetGasFeeCharged sets field value
func (o *ReceiptResponse) SetGasFeeCharged(v string) {
	o.GasFeeCharged = v
}

// GetRawError returns the RawError field value if set, zero value otherwise.
func (o *ReceiptResponse) GetRawError() UnresolvedVMErrorJSON {
	if o == nil || isNil(o.RawError) {
		var ret UnresolvedVMErrorJSON
		return ret
	}
	return *o.RawError
}

// GetRawErrorOk returns a tuple with the RawError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetRawErrorOk() (*UnresolvedVMErrorJSON, bool) {
	if o == nil || isNil(o.RawError) {
		return nil, false
	}
	return o.RawError, true
}

// HasRawError returns a boolean if a field has been set.
func (o *ReceiptResponse) HasRawError() bool {
	if o != nil && !isNil(o.RawError) {
		return true
	}

	return false
}

// SetRawError gets a reference to the given UnresolvedVMErrorJSON and assigns it to the RawError field.
func (o *ReceiptResponse) SetRawError(v UnresolvedVMErrorJSON) {
	o.RawError = &v
}

// GetRequest returns the Request field value
func (o *ReceiptResponse) GetRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *ReceiptResponse) SetRequest(v string) {
	o.Request = v
}

// GetRequestIndex returns the RequestIndex field value
func (o *ReceiptResponse) GetRequestIndex() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.RequestIndex
}

// GetRequestIndexOk returns a tuple with the RequestIndex field value
// and a boolean to check if the value has been set.
func (o *ReceiptResponse) GetRequestIndexOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestIndex, true
}

// SetRequestIndex sets field value
func (o *ReceiptResponse) SetRequestIndex(v uint32) {
	o.RequestIndex = v
}

func (o ReceiptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blockIndex"] = o.BlockIndex
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	toSerialize["gasBudget"] = o.GasBudget
	toSerialize["gasBurnLog"] = o.GasBurnLog
	toSerialize["gasBurned"] = o.GasBurned
	toSerialize["gasFeeCharged"] = o.GasFeeCharged
	if !isNil(o.RawError) {
		toSerialize["rawError"] = o.RawError
	}
	toSerialize["request"] = o.Request
	toSerialize["requestIndex"] = o.RequestIndex
	return toSerialize, nil
}

type NullableReceiptResponse struct {
	value *ReceiptResponse
	isSet bool
}

func (v NullableReceiptResponse) Get() *ReceiptResponse {
	return v.value
}

func (v *NullableReceiptResponse) Set(val *ReceiptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptResponse(val *ReceiptResponse) *NullableReceiptResponse {
	return &NullableReceiptResponse{value: val, isSet: true}
}

func (v NullableReceiptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



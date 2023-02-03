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

// checks if the EventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsResponse{}

// EventsResponse struct for EventsResponse
type EventsResponse struct {
	Events []string `json:"events"`
}

// NewEventsResponse instantiates a new EventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsResponse(events []string) *EventsResponse {
	this := EventsResponse{}
	this.Events = events
	return &this
}

// NewEventsResponseWithDefaults instantiates a new EventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsResponseWithDefaults() *EventsResponse {
	this := EventsResponse{}
	return &this
}

// GetEvents returns the Events field value
func (o *EventsResponse) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsResponse) SetEvents(v []string) {
	o.Events = v
}

func (o EventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

type NullableEventsResponse struct {
	value *EventsResponse
	isSet bool
}

func (v NullableEventsResponse) Get() *EventsResponse {
	return v.value
}

func (v *NullableEventsResponse) Set(val *EventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsResponse(val *EventsResponse) *NullableEventsResponse {
	return &NullableEventsResponse{value: val, isSet: true}
}

func (v NullableEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
DCSA Track and Trace API

API specifications for the Track and Trace interface standard issued by DCSA.org.

API version: 1.2.0
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Events List of events for shipment journey.
type Events struct {
	Events []interface{} `json:"events"`
}

// NewEvents instantiates a new Events object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvents(events []interface{}) *Events {
	this := Events{}
	this.Events = events
	return &this
}

// NewEventsWithDefaults instantiates a new Events object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsWithDefaults() *Events {
	this := Events{}
	return &this
}

// GetEvents returns the Events field value
func (o *Events) GetEvents() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Events) GetEventsOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *Events) SetEvents(v []interface{}) {
	o.Events = v
}

func (o Events) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableEvents struct {
	value *Events
	isSet bool
}

func (v NullableEvents) Get() *Events {
	return v.value
}

func (v *NullableEvents) Set(val *Events) {
	v.value = val
	v.isSet = true
}

func (v NullableEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvents(val *Events) *NullableEvents {
	return &NullableEvents{value: val, isSet: true}
}

func (v NullableEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

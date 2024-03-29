/*
DCSA OpenAPI specification for Track & Trace

Provides equipment actual milestones along with Estimated Time of Arrival following DCSA standards

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BaseShipmentEventAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseShipmentEventAllOf1{}

// BaseShipmentEventAllOf1 struct for BaseShipmentEventAllOf1
type BaseShipmentEventAllOf1 struct {
	// Value for eventDateTime must be the same value as eventCreatedDateTime
	EventDateTime interface{} `json:"eventDateTime,omitempty"`
}

// NewBaseShipmentEventAllOf1 instantiates a new BaseShipmentEventAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEventAllOf1() *BaseShipmentEventAllOf1 {
	this := BaseShipmentEventAllOf1{}
	return &this
}

// NewBaseShipmentEventAllOf1WithDefaults instantiates a new BaseShipmentEventAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventAllOf1WithDefaults() *BaseShipmentEventAllOf1 {
	this := BaseShipmentEventAllOf1{}
	return &this
}

// GetEventDateTime returns the EventDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseShipmentEventAllOf1) GetEventDateTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseShipmentEventAllOf1) GetEventDateTimeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EventDateTime) {
		return nil, false
	}
	return &o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *BaseShipmentEventAllOf1) HasEventDateTime() bool {
	if o != nil && IsNil(o.EventDateTime) {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given interface{} and assigns it to the EventDateTime field.
func (o *BaseShipmentEventAllOf1) SetEventDateTime(v interface{}) {
	o.EventDateTime = v
}

func (o BaseShipmentEventAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseShipmentEventAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EventDateTime != nil {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	return toSerialize, nil
}

type NullableBaseShipmentEventAllOf1 struct {
	value *BaseShipmentEventAllOf1
	isSet bool
}

func (v NullableBaseShipmentEventAllOf1) Get() *BaseShipmentEventAllOf1 {
	return v.value
}

func (v *NullableBaseShipmentEventAllOf1) Set(val *BaseShipmentEventAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEventAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEventAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEventAllOf1(val *BaseShipmentEventAllOf1) *NullableBaseShipmentEventAllOf1 {
	return &NullableBaseShipmentEventAllOf1{value: val, isSet: true}
}

func (v NullableBaseShipmentEventAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEventAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

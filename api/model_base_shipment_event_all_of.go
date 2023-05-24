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

// checks if the BaseShipmentEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseShipmentEventAllOf{}

// BaseShipmentEventAllOf struct for BaseShipmentEventAllOf
type BaseShipmentEventAllOf struct {
	EventType *string `json:"eventType,omitempty"`
}

// NewBaseShipmentEventAllOf instantiates a new BaseShipmentEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEventAllOf() *BaseShipmentEventAllOf {
	this := BaseShipmentEventAllOf{}
	return &this
}

// NewBaseShipmentEventAllOfWithDefaults instantiates a new BaseShipmentEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventAllOfWithDefaults() *BaseShipmentEventAllOf {
	this := BaseShipmentEventAllOf{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseShipmentEventAllOf) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseShipmentEventAllOf) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseShipmentEventAllOf) SetEventType(v string) {
	o.EventType = &v
}

func (o BaseShipmentEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseShipmentEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return toSerialize, nil
}

type NullableBaseShipmentEventAllOf struct {
	value *BaseShipmentEventAllOf
	isSet bool
}

func (v NullableBaseShipmentEventAllOf) Get() *BaseShipmentEventAllOf {
	return v.value
}

func (v *NullableBaseShipmentEventAllOf) Set(val *BaseShipmentEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEventAllOf(val *BaseShipmentEventAllOf) *NullableBaseShipmentEventAllOf {
	return &NullableBaseShipmentEventAllOf{value: val, isSet: true}
}

func (v NullableBaseShipmentEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
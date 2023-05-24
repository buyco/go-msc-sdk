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

// checks if the BaseEventBodyAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEventBodyAllOf1{}

// BaseEventBodyAllOf1 struct for BaseEventBodyAllOf1
type BaseEventBodyAllOf1 struct {
	// Code for the event classifier. Values can vary depending on eventType
	EventClassifierCode *string `json:"eventClassifierCode,omitempty"`
}

// NewBaseEventBodyAllOf1 instantiates a new BaseEventBodyAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventBodyAllOf1() *BaseEventBodyAllOf1 {
	this := BaseEventBodyAllOf1{}
	return &this
}

// NewBaseEventBodyAllOf1WithDefaults instantiates a new BaseEventBodyAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventBodyAllOf1WithDefaults() *BaseEventBodyAllOf1 {
	this := BaseEventBodyAllOf1{}
	return &this
}

// GetEventClassifierCode returns the EventClassifierCode field value if set, zero value otherwise.
func (o *BaseEventBodyAllOf1) GetEventClassifierCode() string {
	if o == nil || IsNil(o.EventClassifierCode) {
		var ret string
		return ret
	}
	return *o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventBodyAllOf1) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EventClassifierCode) {
		return nil, false
	}
	return o.EventClassifierCode, true
}

// HasEventClassifierCode returns a boolean if a field has been set.
func (o *BaseEventBodyAllOf1) HasEventClassifierCode() bool {
	if o != nil && !IsNil(o.EventClassifierCode) {
		return true
	}

	return false
}

// SetEventClassifierCode gets a reference to the given string and assigns it to the EventClassifierCode field.
func (o *BaseEventBodyAllOf1) SetEventClassifierCode(v string) {
	o.EventClassifierCode = &v
}

func (o BaseEventBodyAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEventBodyAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventClassifierCode) {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	return toSerialize, nil
}

type NullableBaseEventBodyAllOf1 struct {
	value *BaseEventBodyAllOf1
	isSet bool
}

func (v NullableBaseEventBodyAllOf1) Get() *BaseEventBodyAllOf1 {
	return v.value
}

func (v *NullableBaseEventBodyAllOf1) Set(val *BaseEventBodyAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventBodyAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventBodyAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventBodyAllOf1(val *BaseEventBodyAllOf1) *NullableBaseEventBodyAllOf1 {
	return &NullableBaseEventBodyAllOf1{value: val, isSet: true}
}

func (v NullableBaseEventBodyAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventBodyAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// checks if the ReferenceAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceAllOf1{}

// ReferenceAllOf1 struct for ReferenceAllOf1
type ReferenceAllOf1 struct {
	// The actual value of the reference.
	ReferenceValue *string `json:"referenceValue,omitempty"`
}

// NewReferenceAllOf1 instantiates a new ReferenceAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceAllOf1() *ReferenceAllOf1 {
	this := ReferenceAllOf1{}
	return &this
}

// NewReferenceAllOf1WithDefaults instantiates a new ReferenceAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceAllOf1WithDefaults() *ReferenceAllOf1 {
	this := ReferenceAllOf1{}
	return &this
}

// GetReferenceValue returns the ReferenceValue field value if set, zero value otherwise.
func (o *ReferenceAllOf1) GetReferenceValue() string {
	if o == nil || IsNil(o.ReferenceValue) {
		var ret string
		return ret
	}
	return *o.ReferenceValue
}

// GetReferenceValueOk returns a tuple with the ReferenceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceAllOf1) GetReferenceValueOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceValue) {
		return nil, false
	}
	return o.ReferenceValue, true
}

// HasReferenceValue returns a boolean if a field has been set.
func (o *ReferenceAllOf1) HasReferenceValue() bool {
	if o != nil && !IsNil(o.ReferenceValue) {
		return true
	}

	return false
}

// SetReferenceValue gets a reference to the given string and assigns it to the ReferenceValue field.
func (o *ReferenceAllOf1) SetReferenceValue(v string) {
	o.ReferenceValue = &v
}

func (o ReferenceAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceValue) {
		toSerialize["referenceValue"] = o.ReferenceValue
	}
	return toSerialize, nil
}

type NullableReferenceAllOf1 struct {
	value *ReferenceAllOf1
	isSet bool
}

func (v NullableReferenceAllOf1) Get() *ReferenceAllOf1 {
	return v.value
}

func (v *NullableReferenceAllOf1) Set(val *ReferenceAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceAllOf1(val *ReferenceAllOf1) *NullableReferenceAllOf1 {
	return &NullableReferenceAllOf1{value: val, isSet: true}
}

func (v NullableReferenceAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

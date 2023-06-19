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

// checks if the Seals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Seals{}

// Seals struct for Seals
type Seals struct {
	Seals []Seal `json:"seals,omitempty"`
}

// NewSeals instantiates a new Seals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeals() *Seals {
	this := Seals{}
	return &this
}

// NewSealsWithDefaults instantiates a new Seals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSealsWithDefaults() *Seals {
	this := Seals{}
	return &this
}

// GetSeals returns the Seals field value if set, zero value otherwise.
func (o *Seals) GetSeals() []Seal {
	if o == nil || IsNil(o.Seals) {
		var ret []Seal
		return ret
	}
	return o.Seals
}

// GetSealsOk returns a tuple with the Seals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seals) GetSealsOk() ([]Seal, bool) {
	if o == nil || IsNil(o.Seals) {
		return nil, false
	}
	return o.Seals, true
}

// HasSeals returns a boolean if a field has been set.
func (o *Seals) HasSeals() bool {
	if o != nil && !IsNil(o.Seals) {
		return true
	}

	return false
}

// SetSeals gets a reference to the given []Seal and assigns it to the Seals field.
func (o *Seals) SetSeals(v []Seal) {
	o.Seals = v
}

func (o Seals) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Seals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seals) {
		toSerialize["seals"] = o.Seals
	}
	return toSerialize, nil
}

type NullableSeals struct {
	value *Seals
	isSet bool
}

func (v NullableSeals) Get() *Seals {
	return v.value
}

func (v *NullableSeals) Set(val *Seals) {
	v.value = val
	v.isSet = true
}

func (v NullableSeals) IsSet() bool {
	return v.isSet
}

func (v *NullableSeals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeals(val *Seals) *NullableSeals {
	return &NullableSeals{value: val, isSet: true}
}

func (v NullableSeals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

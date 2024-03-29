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

// checks if the LocationAllOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationAllOf3{}

// LocationAllOf3 struct for LocationAllOf3
type LocationAllOf3 struct {
	Address *LocationAllOf3Address `json:"address,omitempty"`
}

// NewLocationAllOf3 instantiates a new LocationAllOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAllOf3() *LocationAllOf3 {
	this := LocationAllOf3{}
	return &this
}

// NewLocationAllOf3WithDefaults instantiates a new LocationAllOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAllOf3WithDefaults() *LocationAllOf3 {
	this := LocationAllOf3{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LocationAllOf3) GetAddress() LocationAllOf3Address {
	if o == nil || IsNil(o.Address) {
		var ret LocationAllOf3Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAllOf3) GetAddressOk() (*LocationAllOf3Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LocationAllOf3) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given LocationAllOf3Address and assigns it to the Address field.
func (o *LocationAllOf3) SetAddress(v LocationAllOf3Address) {
	o.Address = &v
}

func (o LocationAllOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationAllOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableLocationAllOf3 struct {
	value *LocationAllOf3
	isSet bool
}

func (v NullableLocationAllOf3) Get() *LocationAllOf3 {
	return v.value
}

func (v *NullableLocationAllOf3) Set(val *LocationAllOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAllOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAllOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAllOf3(val *LocationAllOf3) *NullableLocationAllOf3 {
	return &NullableLocationAllOf3{value: val, isSet: true}
}

func (v NullableLocationAllOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAllOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

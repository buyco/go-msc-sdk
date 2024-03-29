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

// checks if the AddressAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressAllOf1{}

// AddressAllOf1 struct for AddressAllOf1
type AddressAllOf1 struct {
	// The name of the street of the party’s address.
	Street *string `json:"street,omitempty"`
}

// NewAddressAllOf1 instantiates a new AddressAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf1() *AddressAllOf1 {
	this := AddressAllOf1{}
	return &this
}

// NewAddressAllOf1WithDefaults instantiates a new AddressAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf1WithDefaults() *AddressAllOf1 {
	this := AddressAllOf1{}
	return &this
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *AddressAllOf1) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf1) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *AddressAllOf1) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *AddressAllOf1) SetStreet(v string) {
	o.Street = &v
}

func (o AddressAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	return toSerialize, nil
}

type NullableAddressAllOf1 struct {
	value *AddressAllOf1
	isSet bool
}

func (v NullableAddressAllOf1) Get() *AddressAllOf1 {
	return v.value
}

func (v *NullableAddressAllOf1) Set(val *AddressAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf1(val *AddressAllOf1) *NullableAddressAllOf1 {
	return &NullableAddressAllOf1{value: val, isSet: true}
}

func (v NullableAddressAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

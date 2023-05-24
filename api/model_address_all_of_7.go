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

// checks if the AddressAllOf7 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressAllOf7{}

// AddressAllOf7 struct for AddressAllOf7
type AddressAllOf7 struct {
	// The country of the party’s address.
	Country *string `json:"country,omitempty"`
}

// NewAddressAllOf7 instantiates a new AddressAllOf7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf7() *AddressAllOf7 {
	this := AddressAllOf7{}
	return &this
}

// NewAddressAllOf7WithDefaults instantiates a new AddressAllOf7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf7WithDefaults() *AddressAllOf7 {
	this := AddressAllOf7{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressAllOf7) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf7) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressAllOf7) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressAllOf7) SetCountry(v string) {
	o.Country = &v
}

func (o AddressAllOf7) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressAllOf7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableAddressAllOf7 struct {
	value *AddressAllOf7
	isSet bool
}

func (v NullableAddressAllOf7) Get() *AddressAllOf7 {
	return v.value
}

func (v *NullableAddressAllOf7) Set(val *AddressAllOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf7(val *AddressAllOf7) *NullableAddressAllOf7 {
	return &NullableAddressAllOf7{value: val, isSet: true}
}

func (v NullableAddressAllOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
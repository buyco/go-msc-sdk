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

// checks if the AddressAllOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressAllOf4{}

// AddressAllOf4 struct for AddressAllOf4
type AddressAllOf4 struct {
	// The post code of the party’s address.
	PostCode *string `json:"postCode,omitempty"`
}

// NewAddressAllOf4 instantiates a new AddressAllOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf4() *AddressAllOf4 {
	this := AddressAllOf4{}
	return &this
}

// NewAddressAllOf4WithDefaults instantiates a new AddressAllOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf4WithDefaults() *AddressAllOf4 {
	this := AddressAllOf4{}
	return &this
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *AddressAllOf4) GetPostCode() string {
	if o == nil || IsNil(o.PostCode) {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf4) GetPostCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *AddressAllOf4) HasPostCode() bool {
	if o != nil && !IsNil(o.PostCode) {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *AddressAllOf4) SetPostCode(v string) {
	o.PostCode = &v
}

func (o AddressAllOf4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressAllOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostCode) {
		toSerialize["postCode"] = o.PostCode
	}
	return toSerialize, nil
}

type NullableAddressAllOf4 struct {
	value *AddressAllOf4
	isSet bool
}

func (v NullableAddressAllOf4) Get() *AddressAllOf4 {
	return v.value
}

func (v *NullableAddressAllOf4) Set(val *AddressAllOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf4(val *AddressAllOf4) *NullableAddressAllOf4 {
	return &NullableAddressAllOf4{value: val, isSet: true}
}

func (v NullableAddressAllOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

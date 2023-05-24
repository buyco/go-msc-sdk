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

// checks if the LocationAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationAllOf1{}

// LocationAllOf1 struct for LocationAllOf1
type LocationAllOf1 struct {
	// Geographic coordinate that specifies the north–south position of a point on the Earth&apos;s surface.
	Latitude *string `json:"latitude,omitempty"`
}

// NewLocationAllOf1 instantiates a new LocationAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAllOf1() *LocationAllOf1 {
	this := LocationAllOf1{}
	return &this
}

// NewLocationAllOf1WithDefaults instantiates a new LocationAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAllOf1WithDefaults() *LocationAllOf1 {
	this := LocationAllOf1{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *LocationAllOf1) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAllOf1) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *LocationAllOf1) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *LocationAllOf1) SetLatitude(v string) {
	o.Latitude = &v
}

func (o LocationAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	return toSerialize, nil
}

type NullableLocationAllOf1 struct {
	value *LocationAllOf1
	isSet bool
}

func (v NullableLocationAllOf1) Get() *LocationAllOf1 {
	return v.value
}

func (v *NullableLocationAllOf1) Set(val *LocationAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAllOf1(val *LocationAllOf1) *NullableLocationAllOf1 {
	return &NullableLocationAllOf1{value: val, isSet: true}
}

func (v NullableLocationAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
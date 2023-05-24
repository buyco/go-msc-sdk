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

// checks if the TransportCallAllOf5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportCallAllOf5{}

// TransportCallAllOf5 struct for TransportCallAllOf5
type TransportCallAllOf5 struct {
	// The vessel operator-specific identifier of the import Voyage.
	ImportVoyageNumber *string `json:"importVoyageNumber,omitempty"`
}

// NewTransportCallAllOf5 instantiates a new TransportCallAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf5() *TransportCallAllOf5 {
	this := TransportCallAllOf5{}
	return &this
}

// NewTransportCallAllOf5WithDefaults instantiates a new TransportCallAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf5WithDefaults() *TransportCallAllOf5 {
	this := TransportCallAllOf5{}
	return &this
}

// GetImportVoyageNumber returns the ImportVoyageNumber field value if set, zero value otherwise.
func (o *TransportCallAllOf5) GetImportVoyageNumber() string {
	if o == nil || IsNil(o.ImportVoyageNumber) {
		var ret string
		return ret
	}
	return *o.ImportVoyageNumber
}

// GetImportVoyageNumberOk returns a tuple with the ImportVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf5) GetImportVoyageNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ImportVoyageNumber) {
		return nil, false
	}
	return o.ImportVoyageNumber, true
}

// HasImportVoyageNumber returns a boolean if a field has been set.
func (o *TransportCallAllOf5) HasImportVoyageNumber() bool {
	if o != nil && !IsNil(o.ImportVoyageNumber) {
		return true
	}

	return false
}

// SetImportVoyageNumber gets a reference to the given string and assigns it to the ImportVoyageNumber field.
func (o *TransportCallAllOf5) SetImportVoyageNumber(v string) {
	o.ImportVoyageNumber = &v
}

func (o TransportCallAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportCallAllOf5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImportVoyageNumber) {
		toSerialize["importVoyageNumber"] = o.ImportVoyageNumber
	}
	return toSerialize, nil
}

type NullableTransportCallAllOf5 struct {
	value *TransportCallAllOf5
	isSet bool
}

func (v NullableTransportCallAllOf5) Get() *TransportCallAllOf5 {
	return v.value
}

func (v *NullableTransportCallAllOf5) Set(val *TransportCallAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf5(val *TransportCallAllOf5) *NullableTransportCallAllOf5 {
	return &NullableTransportCallAllOf5{value: val, isSet: true}
}

func (v NullableTransportCallAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

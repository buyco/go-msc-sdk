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

// checks if the TransportCallAllOf12 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportCallAllOf12{}

// TransportCallAllOf12 struct for TransportCallAllOf12
type TransportCallAllOf12 struct {
	ModeOfTransport *ModeOfTransport `json:"modeOfTransport,omitempty"`
}

// NewTransportCallAllOf12 instantiates a new TransportCallAllOf12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf12() *TransportCallAllOf12 {
	this := TransportCallAllOf12{}
	return &this
}

// NewTransportCallAllOf12WithDefaults instantiates a new TransportCallAllOf12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf12WithDefaults() *TransportCallAllOf12 {
	this := TransportCallAllOf12{}
	return &this
}

// GetModeOfTransport returns the ModeOfTransport field value if set, zero value otherwise.
func (o *TransportCallAllOf12) GetModeOfTransport() ModeOfTransport {
	if o == nil || IsNil(o.ModeOfTransport) {
		var ret ModeOfTransport
		return ret
	}
	return *o.ModeOfTransport
}

// GetModeOfTransportOk returns a tuple with the ModeOfTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf12) GetModeOfTransportOk() (*ModeOfTransport, bool) {
	if o == nil || IsNil(o.ModeOfTransport) {
		return nil, false
	}
	return o.ModeOfTransport, true
}

// HasModeOfTransport returns a boolean if a field has been set.
func (o *TransportCallAllOf12) HasModeOfTransport() bool {
	if o != nil && !IsNil(o.ModeOfTransport) {
		return true
	}

	return false
}

// SetModeOfTransport gets a reference to the given ModeOfTransport and assigns it to the ModeOfTransport field.
func (o *TransportCallAllOf12) SetModeOfTransport(v ModeOfTransport) {
	o.ModeOfTransport = &v
}

func (o TransportCallAllOf12) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportCallAllOf12) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModeOfTransport) {
		toSerialize["modeOfTransport"] = o.ModeOfTransport
	}
	return toSerialize, nil
}

type NullableTransportCallAllOf12 struct {
	value *TransportCallAllOf12
	isSet bool
}

func (v NullableTransportCallAllOf12) Get() *TransportCallAllOf12 {
	return v.value
}

func (v *NullableTransportCallAllOf12) Set(val *TransportCallAllOf12) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf12) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf12(val *TransportCallAllOf12) *NullableTransportCallAllOf12 {
	return &NullableTransportCallAllOf12{value: val, isSet: true}
}

func (v NullableTransportCallAllOf12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

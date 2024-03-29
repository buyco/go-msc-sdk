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

// checks if the TransportCallAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportCallAllOf{}

// TransportCallAllOf struct for TransportCallAllOf
type TransportCallAllOf struct {
	// The unique identifier for a transport call
	TransportCallID *string `json:"transportCallID,omitempty"`
}

// NewTransportCallAllOf instantiates a new TransportCallAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf() *TransportCallAllOf {
	this := TransportCallAllOf{}
	return &this
}

// NewTransportCallAllOfWithDefaults instantiates a new TransportCallAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOfWithDefaults() *TransportCallAllOf {
	this := TransportCallAllOf{}
	return &this
}

// GetTransportCallID returns the TransportCallID field value if set, zero value otherwise.
func (o *TransportCallAllOf) GetTransportCallID() string {
	if o == nil || IsNil(o.TransportCallID) {
		var ret string
		return ret
	}
	return *o.TransportCallID
}

// GetTransportCallIDOk returns a tuple with the TransportCallID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf) GetTransportCallIDOk() (*string, bool) {
	if o == nil || IsNil(o.TransportCallID) {
		return nil, false
	}
	return o.TransportCallID, true
}

// HasTransportCallID returns a boolean if a field has been set.
func (o *TransportCallAllOf) HasTransportCallID() bool {
	if o != nil && !IsNil(o.TransportCallID) {
		return true
	}

	return false
}

// SetTransportCallID gets a reference to the given string and assigns it to the TransportCallID field.
func (o *TransportCallAllOf) SetTransportCallID(v string) {
	o.TransportCallID = &v
}

func (o TransportCallAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportCallAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransportCallID) {
		toSerialize["transportCallID"] = o.TransportCallID
	}
	return toSerialize, nil
}

type NullableTransportCallAllOf struct {
	value *TransportCallAllOf
	isSet bool
}

func (v NullableTransportCallAllOf) Get() *TransportCallAllOf {
	return v.value
}

func (v *NullableTransportCallAllOf) Set(val *TransportCallAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf(val *TransportCallAllOf) *NullableTransportCallAllOf {
	return &NullableTransportCallAllOf{value: val, isSet: true}
}

func (v NullableTransportCallAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

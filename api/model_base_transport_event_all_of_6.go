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

// checks if the BaseTransportEventAllOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTransportEventAllOf6{}

// BaseTransportEventAllOf6 struct for BaseTransportEventAllOf6
type BaseTransportEventAllOf6 struct {
	TransportCallID *string `json:"transportCallID,omitempty"`
}

// NewBaseTransportEventAllOf6 instantiates a new BaseTransportEventAllOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTransportEventAllOf6() *BaseTransportEventAllOf6 {
	this := BaseTransportEventAllOf6{}
	return &this
}

// NewBaseTransportEventAllOf6WithDefaults instantiates a new BaseTransportEventAllOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTransportEventAllOf6WithDefaults() *BaseTransportEventAllOf6 {
	this := BaseTransportEventAllOf6{}
	return &this
}

// GetTransportCallID returns the TransportCallID field value if set, zero value otherwise.
func (o *BaseTransportEventAllOf6) GetTransportCallID() string {
	if o == nil || IsNil(o.TransportCallID) {
		var ret string
		return ret
	}
	return *o.TransportCallID
}

// GetTransportCallIDOk returns a tuple with the TransportCallID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEventAllOf6) GetTransportCallIDOk() (*string, bool) {
	if o == nil || IsNil(o.TransportCallID) {
		return nil, false
	}
	return o.TransportCallID, true
}

// HasTransportCallID returns a boolean if a field has been set.
func (o *BaseTransportEventAllOf6) HasTransportCallID() bool {
	if o != nil && !IsNil(o.TransportCallID) {
		return true
	}

	return false
}

// SetTransportCallID gets a reference to the given string and assigns it to the TransportCallID field.
func (o *BaseTransportEventAllOf6) SetTransportCallID(v string) {
	o.TransportCallID = &v
}

func (o BaseTransportEventAllOf6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTransportEventAllOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransportCallID) {
		toSerialize["transportCallID"] = o.TransportCallID
	}
	return toSerialize, nil
}

type NullableBaseTransportEventAllOf6 struct {
	value *BaseTransportEventAllOf6
	isSet bool
}

func (v NullableBaseTransportEventAllOf6) Get() *BaseTransportEventAllOf6 {
	return v.value
}

func (v *NullableBaseTransportEventAllOf6) Set(val *BaseTransportEventAllOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTransportEventAllOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTransportEventAllOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTransportEventAllOf6(val *BaseTransportEventAllOf6) *NullableBaseTransportEventAllOf6 {
	return &NullableBaseTransportEventAllOf6{value: val, isSet: true}
}

func (v NullableBaseTransportEventAllOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTransportEventAllOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

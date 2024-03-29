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

// checks if the BaseEquipmentEventAllOf5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEquipmentEventAllOf5{}

// BaseEquipmentEventAllOf5 struct for BaseEquipmentEventAllOf5
type BaseEquipmentEventAllOf5 struct {
	EmptyIndicatorCode *EmptyIndicatorCode `json:"emptyIndicatorCode,omitempty"`
}

// NewBaseEquipmentEventAllOf5 instantiates a new BaseEquipmentEventAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEventAllOf5() *BaseEquipmentEventAllOf5 {
	this := BaseEquipmentEventAllOf5{}
	return &this
}

// NewBaseEquipmentEventAllOf5WithDefaults instantiates a new BaseEquipmentEventAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventAllOf5WithDefaults() *BaseEquipmentEventAllOf5 {
	this := BaseEquipmentEventAllOf5{}
	return &this
}

// GetEmptyIndicatorCode returns the EmptyIndicatorCode field value if set, zero value otherwise.
func (o *BaseEquipmentEventAllOf5) GetEmptyIndicatorCode() EmptyIndicatorCode {
	if o == nil || IsNil(o.EmptyIndicatorCode) {
		var ret EmptyIndicatorCode
		return ret
	}
	return *o.EmptyIndicatorCode
}

// GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEventAllOf5) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool) {
	if o == nil || IsNil(o.EmptyIndicatorCode) {
		return nil, false
	}
	return o.EmptyIndicatorCode, true
}

// HasEmptyIndicatorCode returns a boolean if a field has been set.
func (o *BaseEquipmentEventAllOf5) HasEmptyIndicatorCode() bool {
	if o != nil && !IsNil(o.EmptyIndicatorCode) {
		return true
	}

	return false
}

// SetEmptyIndicatorCode gets a reference to the given EmptyIndicatorCode and assigns it to the EmptyIndicatorCode field.
func (o *BaseEquipmentEventAllOf5) SetEmptyIndicatorCode(v EmptyIndicatorCode) {
	o.EmptyIndicatorCode = &v
}

func (o BaseEquipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEquipmentEventAllOf5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmptyIndicatorCode) {
		toSerialize["emptyIndicatorCode"] = o.EmptyIndicatorCode
	}
	return toSerialize, nil
}

type NullableBaseEquipmentEventAllOf5 struct {
	value *BaseEquipmentEventAllOf5
	isSet bool
}

func (v NullableBaseEquipmentEventAllOf5) Get() *BaseEquipmentEventAllOf5 {
	return v.value
}

func (v *NullableBaseEquipmentEventAllOf5) Set(val *BaseEquipmentEventAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEventAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEventAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEventAllOf5(val *BaseEquipmentEventAllOf5) *NullableBaseEquipmentEventAllOf5 {
	return &NullableBaseEquipmentEventAllOf5{value: val, isSet: true}
}

func (v NullableBaseEquipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEventAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

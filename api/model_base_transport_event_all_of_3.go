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

// checks if the BaseTransportEventAllOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTransportEventAllOf3{}

// BaseTransportEventAllOf3 struct for BaseTransportEventAllOf3
type BaseTransportEventAllOf3 struct {
	// Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/
	DelayReasonCode *string `json:"delayReasonCode,omitempty"`
}

// NewBaseTransportEventAllOf3 instantiates a new BaseTransportEventAllOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTransportEventAllOf3() *BaseTransportEventAllOf3 {
	this := BaseTransportEventAllOf3{}
	return &this
}

// NewBaseTransportEventAllOf3WithDefaults instantiates a new BaseTransportEventAllOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTransportEventAllOf3WithDefaults() *BaseTransportEventAllOf3 {
	this := BaseTransportEventAllOf3{}
	return &this
}

// GetDelayReasonCode returns the DelayReasonCode field value if set, zero value otherwise.
func (o *BaseTransportEventAllOf3) GetDelayReasonCode() string {
	if o == nil || IsNil(o.DelayReasonCode) {
		var ret string
		return ret
	}
	return *o.DelayReasonCode
}

// GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEventAllOf3) GetDelayReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DelayReasonCode) {
		return nil, false
	}
	return o.DelayReasonCode, true
}

// HasDelayReasonCode returns a boolean if a field has been set.
func (o *BaseTransportEventAllOf3) HasDelayReasonCode() bool {
	if o != nil && !IsNil(o.DelayReasonCode) {
		return true
	}

	return false
}

// SetDelayReasonCode gets a reference to the given string and assigns it to the DelayReasonCode field.
func (o *BaseTransportEventAllOf3) SetDelayReasonCode(v string) {
	o.DelayReasonCode = &v
}

func (o BaseTransportEventAllOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTransportEventAllOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayReasonCode) {
		toSerialize["delayReasonCode"] = o.DelayReasonCode
	}
	return toSerialize, nil
}

type NullableBaseTransportEventAllOf3 struct {
	value *BaseTransportEventAllOf3
	isSet bool
}

func (v NullableBaseTransportEventAllOf3) Get() *BaseTransportEventAllOf3 {
	return v.value
}

func (v *NullableBaseTransportEventAllOf3) Set(val *BaseTransportEventAllOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTransportEventAllOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTransportEventAllOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTransportEventAllOf3(val *BaseTransportEventAllOf3) *NullableBaseTransportEventAllOf3 {
	return &NullableBaseTransportEventAllOf3{value: val, isSet: true}
}

func (v NullableBaseTransportEventAllOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTransportEventAllOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

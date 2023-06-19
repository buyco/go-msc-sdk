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

// checks if the References type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &References{}

// References struct for References
type References struct {
	References []Reference `json:"references,omitempty"`
}

// NewReferences instantiates a new References object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferences() *References {
	this := References{}
	return &this
}

// NewReferencesWithDefaults instantiates a new References object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencesWithDefaults() *References {
	this := References{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *References) GetReferences() []Reference {
	if o == nil || IsNil(o.References) {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *References) GetReferencesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *References) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *References) SetReferences(v []Reference) {
	o.References = v
}

func (o References) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o References) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	return toSerialize, nil
}

type NullableReferences struct {
	value *References
	isSet bool
}

func (v NullableReferences) Get() *References {
	return v.value
}

func (v *NullableReferences) Set(val *References) {
	v.value = val
	v.isSet = true
}

func (v NullableReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferences(val *References) *NullableReferences {
	return &NullableReferences{value: val, isSet: true}
}

func (v NullableReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the DocumentReferencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentReferencesInner{}

// DocumentReferencesInner struct for DocumentReferencesInner
type DocumentReferencesInner struct {
	// Describes where the documentReferenceValue is pointing to
	DocumentReferenceType *string `json:"documentReferenceType,omitempty"`
	// The value of the identifier the documentReferenceType is describing
	DocumentReferenceValue *string `json:"documentReferenceValue,omitempty"`
}

// NewDocumentReferencesInner instantiates a new DocumentReferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentReferencesInner() *DocumentReferencesInner {
	this := DocumentReferencesInner{}
	return &this
}

// NewDocumentReferencesInnerWithDefaults instantiates a new DocumentReferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentReferencesInnerWithDefaults() *DocumentReferencesInner {
	this := DocumentReferencesInner{}
	return &this
}

// GetDocumentReferenceType returns the DocumentReferenceType field value if set, zero value otherwise.
func (o *DocumentReferencesInner) GetDocumentReferenceType() string {
	if o == nil || IsNil(o.DocumentReferenceType) {
		var ret string
		return ret
	}
	return *o.DocumentReferenceType
}

// GetDocumentReferenceTypeOk returns a tuple with the DocumentReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferencesInner) GetDocumentReferenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentReferenceType) {
		return nil, false
	}
	return o.DocumentReferenceType, true
}

// HasDocumentReferenceType returns a boolean if a field has been set.
func (o *DocumentReferencesInner) HasDocumentReferenceType() bool {
	if o != nil && !IsNil(o.DocumentReferenceType) {
		return true
	}

	return false
}

// SetDocumentReferenceType gets a reference to the given string and assigns it to the DocumentReferenceType field.
func (o *DocumentReferencesInner) SetDocumentReferenceType(v string) {
	o.DocumentReferenceType = &v
}

// GetDocumentReferenceValue returns the DocumentReferenceValue field value if set, zero value otherwise.
func (o *DocumentReferencesInner) GetDocumentReferenceValue() string {
	if o == nil || IsNil(o.DocumentReferenceValue) {
		var ret string
		return ret
	}
	return *o.DocumentReferenceValue
}

// GetDocumentReferenceValueOk returns a tuple with the DocumentReferenceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferencesInner) GetDocumentReferenceValueOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentReferenceValue) {
		return nil, false
	}
	return o.DocumentReferenceValue, true
}

// HasDocumentReferenceValue returns a boolean if a field has been set.
func (o *DocumentReferencesInner) HasDocumentReferenceValue() bool {
	if o != nil && !IsNil(o.DocumentReferenceValue) {
		return true
	}

	return false
}

// SetDocumentReferenceValue gets a reference to the given string and assigns it to the DocumentReferenceValue field.
func (o *DocumentReferencesInner) SetDocumentReferenceValue(v string) {
	o.DocumentReferenceValue = &v
}

func (o DocumentReferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentReferencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentReferenceType) {
		toSerialize["documentReferenceType"] = o.DocumentReferenceType
	}
	if !IsNil(o.DocumentReferenceValue) {
		toSerialize["documentReferenceValue"] = o.DocumentReferenceValue
	}
	return toSerialize, nil
}

type NullableDocumentReferencesInner struct {
	value *DocumentReferencesInner
	isSet bool
}

func (v NullableDocumentReferencesInner) Get() *DocumentReferencesInner {
	return v.value
}

func (v *NullableDocumentReferencesInner) Set(val *DocumentReferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentReferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentReferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentReferencesInner(val *DocumentReferencesInner) *NullableDocumentReferencesInner {
	return &NullableDocumentReferencesInner{value: val, isSet: true}
}

func (v NullableDocumentReferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentReferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

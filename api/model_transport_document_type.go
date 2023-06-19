/*
DCSA OpenAPI specification for Track & Trace

Provides equipment actual milestones along with Estimated Time of Arrival following DCSA standards

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TransportDocumentType Specifies the type of the transport document (a Bill of Lading (BOL) or a Sea Waybill (SWB)).
type TransportDocumentType string

// List of transportDocumentType
const (
	BOL TransportDocumentType = "BOL"
	SWB TransportDocumentType = "SWB"
)

// All allowed values of TransportDocumentType enum
var AllowedTransportDocumentTypeEnumValues = []TransportDocumentType{
	"BOL",
	"SWB",
}

func (v *TransportDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportDocumentType(value)
	for _, existing := range AllowedTransportDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportDocumentType", value)
}

// NewTransportDocumentTypeFromValue returns a pointer to a valid TransportDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportDocumentTypeFromValue(v string) (*TransportDocumentType, error) {
	ev := TransportDocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportDocumentType: valid values are %v", v, AllowedTransportDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportDocumentType) IsValid() bool {
	for _, existing := range AllowedTransportDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to transportDocumentType value
func (v TransportDocumentType) Ptr() *TransportDocumentType {
	return &v
}

type NullableTransportDocumentType struct {
	value *TransportDocumentType
	isSet bool
}

func (v NullableTransportDocumentType) Get() *TransportDocumentType {
	return v.value
}

func (v *NullableTransportDocumentType) Set(val *TransportDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportDocumentType(val *TransportDocumentType) *NullableTransportDocumentType {
	return &NullableTransportDocumentType{value: val, isSet: true}
}

func (v NullableTransportDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

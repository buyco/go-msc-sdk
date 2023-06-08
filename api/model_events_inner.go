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

// EventsInner struct for EventsInner
type EventsInner struct {
	EquipmentEvent *EquipmentEvent
	ShipmentEvent  *ShipmentEvent
	TransportEvent *TransportEvent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EquipmentEvent
	err = json.Unmarshal(data, &dst.EquipmentEvent)
	if err == nil {
		jsonEquipmentEvent, _ := json.Marshal(dst.EquipmentEvent)
		if string(jsonEquipmentEvent) == "{}" { // empty struct
			dst.EquipmentEvent = nil
		} else {
			return nil // data stored in dst.EquipmentEvent, return on the first match
		}
	} else {
		dst.EquipmentEvent = nil
	}

	// try to unmarshal JSON data into ShipmentEvent
	err = json.Unmarshal(data, &dst.ShipmentEvent)
	if err == nil {
		jsonShipmentEvent, _ := json.Marshal(dst.ShipmentEvent)
		if string(jsonShipmentEvent) == "{}" { // empty struct
			dst.ShipmentEvent = nil
		} else {
			return nil // data stored in dst.ShipmentEvent, return on the first match
		}
	} else {
		dst.ShipmentEvent = nil
	}

	// try to unmarshal JSON data into TransportEvent
	err = json.Unmarshal(data, &dst.TransportEvent)
	if err == nil {
		jsonTransportEvent, _ := json.Marshal(dst.TransportEvent)
		if string(jsonTransportEvent) == "{}" { // empty struct
			dst.TransportEvent = nil
		} else {
			return nil // data stored in dst.TransportEvent, return on the first match
		}
	} else {
		dst.TransportEvent = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EventsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventsInner) MarshalJSON() ([]byte, error) {
	if src.EquipmentEvent != nil {
		return json.Marshal(&src.EquipmentEvent)
	}

	if src.ShipmentEvent != nil {
		return json.Marshal(&src.ShipmentEvent)
	}

	if src.TransportEvent != nil {
		return json.Marshal(&src.TransportEvent)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventsInner struct {
	value *EventsInner
	isSet bool
}

func (v NullableEventsInner) Get() *EventsInner {
	return v.value
}

func (v *NullableEventsInner) Set(val *EventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsInner(val *EventsInner) *NullableEventsInner {
	return &NullableEventsInner{value: val, isSet: true}
}

func (v NullableEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
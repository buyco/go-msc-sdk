package model

import "time"

type TrackingResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	EventType                   string    `json:"eventType"`
	EquipmentReference          string    `json:"equipmentReference,omitempty"`
	FacilityTypeCode            string    `json:"facilityTypeCode,omitempty"`
	UnLocationCode              string    `json:"unLocationCode,omitempty"`
	FacilityCode                string    `json:"facilityCode,omitempty"`
	OtherFacility               string    `json:"otherFacility,omitempty"`
	EmptyIndicatorCode          string    `json:"emptyIndicatorCode,omitempty"`
	Description                 string    `json:"description"`
	EventID                     string    `json:"eventId"`
	EventDateTime               time.Time `json:"eventDateTime"`
	EventClassifierCode         string    `json:"eventClassifierCode"`
	EventTypeCode               string    `json:"eventTypeCode"`
	ShipmentInformationTypeCode string    `json:"shipmentInformationTypeCode,omitempty"`
	TransportReference          string    `json:"transportReference,omitempty"`
	TransportLegReference       string    `json:"transportLegReference,omitempty"`
	ModeOfTransportCode         string    `json:"modeOfTransportCode,omitempty"`
}

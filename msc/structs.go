package msc

type Event struct {
	BookingNumber      string
	Description        string
	EmptyIndicatorCode string

	EquipmentReference string

	EventClassifierCode string
	EventDateTime       string
	EventId             string
	EventType           string
	EventTypeCode       string

	FacilityCode     string
	FacilityTypeCode string
	OtherFacility    string

	TransportLegReference string
	TransportMode         string
	TransportReference    string

	UnLocationCode string
}

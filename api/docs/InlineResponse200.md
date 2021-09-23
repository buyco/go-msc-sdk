# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID. | 
**EventType** | **string** |  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place, in ISO 8601 format. | 
**EventClassifierCode** | **string** | Code for the event classifier, either PLN, ACT or EST. | 
**EventTypeCode** | **string** | Unique identifier for Event Type Code. | 
**TransportReference** | **string** | The reference for the transport, e.g. when the mode of transport is a vessel, the transport reference will be the vessel IMO number. | 
**TransportLegReference** | **string** | The transport leg reference will be specific per mode of transport: - Vessel: Voyage number as specified by the vessel operator - Truck: Not yet specified - Rail: Not yet specified - Barge: Not yet specified.  | 
**FacilityTypeCode** | **string** | The code to identify the specific type of facility. | 
**UNLocationCode** | **string** | The UN Location Code identifies a location in the sense of a city/a town/a village, being the smaller administrative area existing as defined by the competent national authority in each country. | 
**FacilityCode** | **string** | The code used for identifying the specific facility. | 
**OtherFacility** | Pointer to **string** | An alternative way to capture the facility when no standardized DCSA facility code can be found. | [optional] 
**ModeOfTransportCode** | Pointer to **string** | A code specifying a type of transport mode. | [optional] 
**ShipmentInformationTypeCode** | **string** | The code to identify the type of information that is related to the shipment. | 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. If a container is not yet assigned to a shipment, the interface cannot return any information when an equipment reference is given as input. If a container is assigned to an (active) shipment, the interface returns information on the active shipment. | [optional] 
**EmptyIndicatorCode** | **string** | Code to denote whether the equipment is empty or laden. | 

## Methods

### NewInlineResponse200

`func NewInlineResponse200(eventID string, eventType string, eventDateTime time.Time, eventClassifierCode string, eventTypeCode string, transportReference string, transportLegReference string, facilityTypeCode string, uNLocationCode string, facilityCode string, shipmentInformationTypeCode string, emptyIndicatorCode string, ) *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *InlineResponse200) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *InlineResponse200) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *InlineResponse200) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetEventType

`func (o *InlineResponse200) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *InlineResponse200) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *InlineResponse200) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventDateTime

`func (o *InlineResponse200) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *InlineResponse200) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *InlineResponse200) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventClassifierCode

`func (o *InlineResponse200) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *InlineResponse200) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *InlineResponse200) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventTypeCode

`func (o *InlineResponse200) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *InlineResponse200) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *InlineResponse200) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.


### GetTransportReference

`func (o *InlineResponse200) GetTransportReference() string`

GetTransportReference returns the TransportReference field if non-nil, zero value otherwise.

### GetTransportReferenceOk

`func (o *InlineResponse200) GetTransportReferenceOk() (*string, bool)`

GetTransportReferenceOk returns a tuple with the TransportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReference

`func (o *InlineResponse200) SetTransportReference(v string)`

SetTransportReference sets TransportReference field to given value.


### GetTransportLegReference

`func (o *InlineResponse200) GetTransportLegReference() string`

GetTransportLegReference returns the TransportLegReference field if non-nil, zero value otherwise.

### GetTransportLegReferenceOk

`func (o *InlineResponse200) GetTransportLegReferenceOk() (*string, bool)`

GetTransportLegReferenceOk returns a tuple with the TransportLegReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportLegReference

`func (o *InlineResponse200) SetTransportLegReference(v string)`

SetTransportLegReference sets TransportLegReference field to given value.


### GetFacilityTypeCode

`func (o *InlineResponse200) GetFacilityTypeCode() string`

GetFacilityTypeCode returns the FacilityTypeCode field if non-nil, zero value otherwise.

### GetFacilityTypeCodeOk

`func (o *InlineResponse200) GetFacilityTypeCodeOk() (*string, bool)`

GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityTypeCode

`func (o *InlineResponse200) SetFacilityTypeCode(v string)`

SetFacilityTypeCode sets FacilityTypeCode field to given value.


### GetUNLocationCode

`func (o *InlineResponse200) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *InlineResponse200) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *InlineResponse200) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.


### GetFacilityCode

`func (o *InlineResponse200) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *InlineResponse200) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *InlineResponse200) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.


### GetOtherFacility

`func (o *InlineResponse200) GetOtherFacility() string`

GetOtherFacility returns the OtherFacility field if non-nil, zero value otherwise.

### GetOtherFacilityOk

`func (o *InlineResponse200) GetOtherFacilityOk() (*string, bool)`

GetOtherFacilityOk returns a tuple with the OtherFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFacility

`func (o *InlineResponse200) SetOtherFacility(v string)`

SetOtherFacility sets OtherFacility field to given value.

### HasOtherFacility

`func (o *InlineResponse200) HasOtherFacility() bool`

HasOtherFacility returns a boolean if a field has been set.

### GetModeOfTransportCode

`func (o *InlineResponse200) GetModeOfTransportCode() string`

GetModeOfTransportCode returns the ModeOfTransportCode field if non-nil, zero value otherwise.

### GetModeOfTransportCodeOk

`func (o *InlineResponse200) GetModeOfTransportCodeOk() (*string, bool)`

GetModeOfTransportCodeOk returns a tuple with the ModeOfTransportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransportCode

`func (o *InlineResponse200) SetModeOfTransportCode(v string)`

SetModeOfTransportCode sets ModeOfTransportCode field to given value.

### HasModeOfTransportCode

`func (o *InlineResponse200) HasModeOfTransportCode() bool`

HasModeOfTransportCode returns a boolean if a field has been set.

### GetShipmentInformationTypeCode

`func (o *InlineResponse200) GetShipmentInformationTypeCode() string`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *InlineResponse200) GetShipmentInformationTypeCodeOk() (*string, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *InlineResponse200) SetShipmentInformationTypeCode(v string)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.


### GetEquipmentReference

`func (o *InlineResponse200) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *InlineResponse200) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *InlineResponse200) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *InlineResponse200) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *InlineResponse200) GetEmptyIndicatorCode() string`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *InlineResponse200) GetEmptyIndicatorCodeOk() (*string, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *InlineResponse200) SetEmptyIndicatorCode(v string)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



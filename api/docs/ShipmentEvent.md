# ShipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID. | 
**EventType** | **string** |  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place, in ISO 8601 format. | 
**EventClassifierCode** | **string** | Code for the event classifier, either PLN, ACT or EST. | 
**EventTypeCode** | **string** | Unique identifier for Event Type Code. | 
**ShipmentInformationTypeCode** | **string** | The code to identify the type of information that is related to the shipment. | 

## Methods

### NewShipmentEvent

`func NewShipmentEvent(eventID string, eventType string, eventDateTime time.Time, eventClassifierCode string, eventTypeCode string, shipmentInformationTypeCode string, ) *ShipmentEvent`

NewShipmentEvent instantiates a new ShipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentEventWithDefaults

`func NewShipmentEventWithDefaults() *ShipmentEvent`

NewShipmentEventWithDefaults instantiates a new ShipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *ShipmentEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *ShipmentEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *ShipmentEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetEventType

`func (o *ShipmentEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ShipmentEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ShipmentEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventDateTime

`func (o *ShipmentEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *ShipmentEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *ShipmentEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventClassifierCode

`func (o *ShipmentEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *ShipmentEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *ShipmentEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventTypeCode

`func (o *ShipmentEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *ShipmentEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *ShipmentEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.


### GetShipmentInformationTypeCode

`func (o *ShipmentEvent) GetShipmentInformationTypeCode() string`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *ShipmentEvent) GetShipmentInformationTypeCodeOk() (*string, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *ShipmentEvent) SetShipmentInformationTypeCode(v string)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



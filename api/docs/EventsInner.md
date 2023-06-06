# EventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventType** | **string** |  | 
**EventClassifierCode** | **string** | Code for the event classifier can be - PLN (Planned) - ACT (Actual) - EST (Estimated)  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 
**TransportEventTypeCode** | [**TransportEventTypeCode**](TransportEventTypeCode.md) |  | 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/ | [optional] 
**VesselScheduleChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.  Deprecated - use changeRemark instead  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**TransportCallID** | Pointer to **string** |  | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code, for transport events this is either - LOAD (Loaded) - DISC (Discharged) - GTIN (Gated in) - GTOT (Gated out) - STUF (Stuffed) - STRP (Stripped)  Deprecated - use equipmentEventTypeCode instead  | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The &lt;b&gt;documentReferenceType&lt;/b&gt;-field is used to describe where the &lt;b&gt;documentReferenceValue&lt;/b&gt;-field is pointing to. | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**ShipmentEventTypeCode** | [**ShipmentEventTypeCode**](ShipmentEventTypeCode.md) |  | 
**DocumentID** | **string** | The id of the object defined by the documentTypeCode.  | 
**DocumentTypeCode** | [**DocumentTypeCode**](DocumentTypeCode.md) |  | 
**ShipmentInformationTypeCode** | Pointer to [**ShipmentInformationType**](ShipmentInformationType.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 
**ShipmentID** | Pointer to **interface{}** | ID uniquely identifying a shipment.  Deprecated - this is replaced by documentID which can contain different values depending on the value of the documentTypeCode field  | [optional] 
**EquipmentEventTypeCode** | [**EquipmentEventTypeCode**](EquipmentEventTypeCode.md) |  | 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | [optional] 
**ISOEquipmentCode** | Pointer to **string** | Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | [**EmptyIndicatorCode**](EmptyIndicatorCode.md) |  | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) |  | [optional] 

## Methods

### NewEventsInner

`func NewEventsInner(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall, shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode, equipmentEventTypeCode EquipmentEventTypeCode, emptyIndicatorCode EmptyIndicatorCode, ) *EventsInner`

NewEventsInner instantiates a new EventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsInnerWithDefaults

`func NewEventsInnerWithDefaults() *EventsInner`

NewEventsInnerWithDefaults instantiates a new EventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *EventsInner) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *EventsInner) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *EventsInner) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *EventsInner) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *EventsInner) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *EventsInner) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *EventsInner) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *EventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *EventsInner) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *EventsInner) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *EventsInner) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *EventsInner) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *EventsInner) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *EventsInner) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetTransportEventTypeCode

`func (o *EventsInner) GetTransportEventTypeCode() TransportEventTypeCode`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *EventsInner) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *EventsInner) SetTransportEventTypeCode(v TransportEventTypeCode)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *EventsInner) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *EventsInner) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *EventsInner) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *EventsInner) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetVesselScheduleChangeRemark

`func (o *EventsInner) GetVesselScheduleChangeRemark() string`

GetVesselScheduleChangeRemark returns the VesselScheduleChangeRemark field if non-nil, zero value otherwise.

### GetVesselScheduleChangeRemarkOk

`func (o *EventsInner) GetVesselScheduleChangeRemarkOk() (*string, bool)`

GetVesselScheduleChangeRemarkOk returns a tuple with the VesselScheduleChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselScheduleChangeRemark

`func (o *EventsInner) SetVesselScheduleChangeRemark(v string)`

SetVesselScheduleChangeRemark sets VesselScheduleChangeRemark field to given value.

### HasVesselScheduleChangeRemark

`func (o *EventsInner) HasVesselScheduleChangeRemark() bool`

HasVesselScheduleChangeRemark returns a boolean if a field has been set.

### GetChangeRemark

`func (o *EventsInner) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *EventsInner) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *EventsInner) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *EventsInner) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCallID

`func (o *EventsInner) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *EventsInner) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *EventsInner) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *EventsInner) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetTransportCall

`func (o *EventsInner) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *EventsInner) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *EventsInner) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetEventTypeCode

`func (o *EventsInner) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *EventsInner) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *EventsInner) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *EventsInner) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *EventsInner) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *EventsInner) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *EventsInner) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *EventsInner) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *EventsInner) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EventsInner) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EventsInner) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EventsInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *EventsInner) GetShipmentEventTypeCode() ShipmentEventTypeCode`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *EventsInner) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *EventsInner) SetShipmentEventTypeCode(v ShipmentEventTypeCode)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentID

`func (o *EventsInner) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *EventsInner) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *EventsInner) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetDocumentTypeCode

`func (o *EventsInner) GetDocumentTypeCode() DocumentTypeCode`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *EventsInner) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *EventsInner) SetDocumentTypeCode(v DocumentTypeCode)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetShipmentInformationTypeCode

`func (o *EventsInner) GetShipmentInformationTypeCode() ShipmentInformationType`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *EventsInner) GetShipmentInformationTypeCodeOk() (*ShipmentInformationType, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *EventsInner) SetShipmentInformationTypeCode(v ShipmentInformationType)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.

### HasShipmentInformationTypeCode

`func (o *EventsInner) HasShipmentInformationTypeCode() bool`

HasShipmentInformationTypeCode returns a boolean if a field has been set.

### GetReason

`func (o *EventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShipmentID

`func (o *EventsInner) GetShipmentID() interface{}`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *EventsInner) GetShipmentIDOk() (*interface{}, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *EventsInner) SetShipmentID(v interface{})`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *EventsInner) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *EventsInner) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *EventsInner) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetEquipmentEventTypeCode

`func (o *EventsInner) GetEquipmentEventTypeCode() EquipmentEventTypeCode`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *EventsInner) GetEquipmentEventTypeCodeOk() (*EquipmentEventTypeCode, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *EventsInner) SetEquipmentEventTypeCode(v EquipmentEventTypeCode)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *EventsInner) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *EventsInner) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *EventsInner) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *EventsInner) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *EventsInner) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *EventsInner) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *EventsInner) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *EventsInner) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *EventsInner) GetEmptyIndicatorCode() EmptyIndicatorCode`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *EventsInner) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *EventsInner) SetEmptyIndicatorCode(v EmptyIndicatorCode)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *EventsInner) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *EventsInner) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *EventsInner) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *EventsInner) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetSeals

`func (o *EventsInner) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *EventsInner) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *EventsInner) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *EventsInner) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



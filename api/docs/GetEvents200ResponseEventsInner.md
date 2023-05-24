# GetEvents200ResponseEventsInner

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

### NewGetEvents200ResponseEventsInner

`func NewGetEvents200ResponseEventsInner(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall, shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode, equipmentEventTypeCode EquipmentEventTypeCode, emptyIndicatorCode EmptyIndicatorCode, ) *GetEvents200ResponseEventsInner`

NewGetEvents200ResponseEventsInner instantiates a new GetEvents200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseEventsInnerWithDefaults

`func NewGetEvents200ResponseEventsInnerWithDefaults() *GetEvents200ResponseEventsInner`

NewGetEvents200ResponseEventsInnerWithDefaults instantiates a new GetEvents200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *GetEvents200ResponseEventsInner) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *GetEvents200ResponseEventsInner) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *GetEvents200ResponseEventsInner) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *GetEvents200ResponseEventsInner) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *GetEvents200ResponseEventsInner) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *GetEvents200ResponseEventsInner) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *GetEvents200ResponseEventsInner) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *GetEvents200ResponseEventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetEvents200ResponseEventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetEvents200ResponseEventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *GetEvents200ResponseEventsInner) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *GetEvents200ResponseEventsInner) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *GetEvents200ResponseEventsInner) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *GetEvents200ResponseEventsInner) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *GetEvents200ResponseEventsInner) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *GetEvents200ResponseEventsInner) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetTransportEventTypeCode

`func (o *GetEvents200ResponseEventsInner) GetTransportEventTypeCode() TransportEventTypeCode`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *GetEvents200ResponseEventsInner) SetTransportEventTypeCode(v TransportEventTypeCode)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *GetEvents200ResponseEventsInner) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *GetEvents200ResponseEventsInner) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *GetEvents200ResponseEventsInner) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *GetEvents200ResponseEventsInner) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetVesselScheduleChangeRemark

`func (o *GetEvents200ResponseEventsInner) GetVesselScheduleChangeRemark() string`

GetVesselScheduleChangeRemark returns the VesselScheduleChangeRemark field if non-nil, zero value otherwise.

### GetVesselScheduleChangeRemarkOk

`func (o *GetEvents200ResponseEventsInner) GetVesselScheduleChangeRemarkOk() (*string, bool)`

GetVesselScheduleChangeRemarkOk returns a tuple with the VesselScheduleChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselScheduleChangeRemark

`func (o *GetEvents200ResponseEventsInner) SetVesselScheduleChangeRemark(v string)`

SetVesselScheduleChangeRemark sets VesselScheduleChangeRemark field to given value.

### HasVesselScheduleChangeRemark

`func (o *GetEvents200ResponseEventsInner) HasVesselScheduleChangeRemark() bool`

HasVesselScheduleChangeRemark returns a boolean if a field has been set.

### GetChangeRemark

`func (o *GetEvents200ResponseEventsInner) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *GetEvents200ResponseEventsInner) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *GetEvents200ResponseEventsInner) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *GetEvents200ResponseEventsInner) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCallID

`func (o *GetEvents200ResponseEventsInner) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *GetEvents200ResponseEventsInner) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *GetEvents200ResponseEventsInner) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *GetEvents200ResponseEventsInner) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetTransportCall

`func (o *GetEvents200ResponseEventsInner) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *GetEvents200ResponseEventsInner) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *GetEvents200ResponseEventsInner) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetEventTypeCode

`func (o *GetEvents200ResponseEventsInner) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *GetEvents200ResponseEventsInner) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *GetEvents200ResponseEventsInner) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *GetEvents200ResponseEventsInner) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *GetEvents200ResponseEventsInner) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *GetEvents200ResponseEventsInner) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *GetEvents200ResponseEventsInner) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *GetEvents200ResponseEventsInner) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetEvents200ResponseEventsInner) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetEvents200ResponseEventsInner) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *GetEvents200ResponseEventsInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *GetEvents200ResponseEventsInner) GetShipmentEventTypeCode() ShipmentEventTypeCode`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *GetEvents200ResponseEventsInner) SetShipmentEventTypeCode(v ShipmentEventTypeCode)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentID

`func (o *GetEvents200ResponseEventsInner) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *GetEvents200ResponseEventsInner) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *GetEvents200ResponseEventsInner) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetDocumentTypeCode

`func (o *GetEvents200ResponseEventsInner) GetDocumentTypeCode() DocumentTypeCode`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *GetEvents200ResponseEventsInner) SetDocumentTypeCode(v DocumentTypeCode)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetShipmentInformationTypeCode

`func (o *GetEvents200ResponseEventsInner) GetShipmentInformationTypeCode() ShipmentInformationType`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetShipmentInformationTypeCodeOk() (*ShipmentInformationType, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *GetEvents200ResponseEventsInner) SetShipmentInformationTypeCode(v ShipmentInformationType)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.

### HasShipmentInformationTypeCode

`func (o *GetEvents200ResponseEventsInner) HasShipmentInformationTypeCode() bool`

HasShipmentInformationTypeCode returns a boolean if a field has been set.

### GetReason

`func (o *GetEvents200ResponseEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetEvents200ResponseEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetEvents200ResponseEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetEvents200ResponseEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShipmentID

`func (o *GetEvents200ResponseEventsInner) GetShipmentID() interface{}`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *GetEvents200ResponseEventsInner) GetShipmentIDOk() (*interface{}, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *GetEvents200ResponseEventsInner) SetShipmentID(v interface{})`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *GetEvents200ResponseEventsInner) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *GetEvents200ResponseEventsInner) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *GetEvents200ResponseEventsInner) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetEquipmentEventTypeCode

`func (o *GetEvents200ResponseEventsInner) GetEquipmentEventTypeCode() EquipmentEventTypeCode`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *GetEvents200ResponseEventsInner) GetEquipmentEventTypeCodeOk() (*EquipmentEventTypeCode, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *GetEvents200ResponseEventsInner) SetEquipmentEventTypeCode(v EquipmentEventTypeCode)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *GetEvents200ResponseEventsInner) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *GetEvents200ResponseEventsInner) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *GetEvents200ResponseEventsInner) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *GetEvents200ResponseEventsInner) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *GetEvents200ResponseEventsInner) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *GetEvents200ResponseEventsInner) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *GetEvents200ResponseEventsInner) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *GetEvents200ResponseEventsInner) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *GetEvents200ResponseEventsInner) GetEmptyIndicatorCode() EmptyIndicatorCode`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *GetEvents200ResponseEventsInner) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *GetEvents200ResponseEventsInner) SetEmptyIndicatorCode(v EmptyIndicatorCode)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *GetEvents200ResponseEventsInner) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *GetEvents200ResponseEventsInner) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *GetEvents200ResponseEventsInner) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *GetEvents200ResponseEventsInner) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetSeals

`func (o *GetEvents200ResponseEventsInner) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *GetEvents200ResponseEventsInner) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *GetEvents200ResponseEventsInner) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *GetEvents200ResponseEventsInner) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



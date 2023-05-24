# BaseEquipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**EventClassifierCode** | Pointer to **string** | Code for the event classifier can be - PLN (Planned) - ACT (Actual) - EST (Estimated)  | [optional] 
**EquipmentEventTypeCode** | [**EquipmentEventTypeCode**](EquipmentEventTypeCode.md) |  | 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | [optional] 
**ISOEquipmentCode** | Pointer to **string** | Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | [**EmptyIndicatorCode**](EmptyIndicatorCode.md) |  | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**TransportCallID** | Pointer to **string** |  | [optional] 
**TransportCall** | Pointer to [**TransportCall**](TransportCall.md) |  | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The &lt;b&gt;documentReferenceType&lt;/b&gt;-field is used to describe where the &lt;b&gt;documentReferenceValue&lt;/b&gt;-field is pointing to. | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) |  | [optional] 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code, for transport events this is either - LOAD (Loaded) - DISC (Discharged) - GTIN (Gated in) - GTOT (Gated out) - STUF (Stuffed) - STRP (Stripped)  Deprecated - use equipmentEventTypeCode instead  | [optional] 

## Methods

### NewBaseEquipmentEvent

`func NewBaseEquipmentEvent(equipmentEventTypeCode EquipmentEventTypeCode, emptyIndicatorCode EmptyIndicatorCode, ) *BaseEquipmentEvent`

NewBaseEquipmentEvent instantiates a new BaseEquipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEquipmentEventWithDefaults

`func NewBaseEquipmentEventWithDefaults() *BaseEquipmentEvent`

NewBaseEquipmentEventWithDefaults instantiates a new BaseEquipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BaseEquipmentEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEquipmentEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEquipmentEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseEquipmentEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventClassifierCode

`func (o *BaseEquipmentEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *BaseEquipmentEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *BaseEquipmentEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.

### HasEventClassifierCode

`func (o *BaseEquipmentEvent) HasEventClassifierCode() bool`

HasEventClassifierCode returns a boolean if a field has been set.

### GetEquipmentEventTypeCode

`func (o *BaseEquipmentEvent) GetEquipmentEventTypeCode() EquipmentEventTypeCode`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *BaseEquipmentEvent) GetEquipmentEventTypeCodeOk() (*EquipmentEventTypeCode, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *BaseEquipmentEvent) SetEquipmentEventTypeCode(v EquipmentEventTypeCode)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *BaseEquipmentEvent) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *BaseEquipmentEvent) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *BaseEquipmentEvent) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *BaseEquipmentEvent) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *BaseEquipmentEvent) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *BaseEquipmentEvent) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *BaseEquipmentEvent) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *BaseEquipmentEvent) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *BaseEquipmentEvent) GetEmptyIndicatorCode() EmptyIndicatorCode`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *BaseEquipmentEvent) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *BaseEquipmentEvent) SetEmptyIndicatorCode(v EmptyIndicatorCode)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *BaseEquipmentEvent) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *BaseEquipmentEvent) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *BaseEquipmentEvent) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *BaseEquipmentEvent) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetTransportCallID

`func (o *BaseEquipmentEvent) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *BaseEquipmentEvent) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *BaseEquipmentEvent) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *BaseEquipmentEvent) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetTransportCall

`func (o *BaseEquipmentEvent) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *BaseEquipmentEvent) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *BaseEquipmentEvent) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.

### HasTransportCall

`func (o *BaseEquipmentEvent) HasTransportCall() bool`

HasTransportCall returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *BaseEquipmentEvent) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *BaseEquipmentEvent) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *BaseEquipmentEvent) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *BaseEquipmentEvent) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *BaseEquipmentEvent) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BaseEquipmentEvent) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BaseEquipmentEvent) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BaseEquipmentEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeals

`func (o *BaseEquipmentEvent) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *BaseEquipmentEvent) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *BaseEquipmentEvent) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *BaseEquipmentEvent) HasSeals() bool`

HasSeals returns a boolean if a field has been set.

### GetEventTypeCode

`func (o *BaseEquipmentEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseEquipmentEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseEquipmentEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseEquipmentEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ShipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventType** | **string** |  | 
**EventClassifierCode** | **string** | Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)  | 
**EventDateTime** | **interface{}** | Value for eventDateTime must be the same value as eventCreatedDateTime  | 
**ShipmentEventTypeCode** | [**ShipmentEventTypeCode**](ShipmentEventTypeCode.md) |  | 
**DocumentID** | **string** | The id of the object defined by the documentTypeCode.  | 
**DocumentTypeCode** | [**DocumentTypeCode**](DocumentTypeCode.md) |  | 
**ShipmentInformationTypeCode** | Pointer to [**ShipmentInformationType**](ShipmentInformationType.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code. For shipment events this can be - RECE (Received) - CONF (Confirmed) - ISSU (Issued) - APPR (Approved) - SUBM (Submitted) - SURR (Surrendered) - REJE (Rejected) - PENA (Pending approval)  Deprecated - use shipmentEventTypeCode instead  | [optional] 
**ShipmentID** | Pointer to **interface{}** | ID uniquely identifying a shipment.  Deprecated - this is replaced by documentID which can contain different values depending on the value of the documentTypeCode field  | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 

## Methods

### NewShipmentEvent

`func NewShipmentEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime interface{}, shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode, ) *ShipmentEvent`

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

### HasEventID

`func (o *ShipmentEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *ShipmentEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *ShipmentEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *ShipmentEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


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


### GetEventDateTime

`func (o *ShipmentEvent) GetEventDateTime() interface{}`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *ShipmentEvent) GetEventDateTimeOk() (*interface{}, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *ShipmentEvent) SetEventDateTime(v interface{})`

SetEventDateTime sets EventDateTime field to given value.


### SetEventDateTimeNil

`func (o *ShipmentEvent) SetEventDateTimeNil(b bool)`

 SetEventDateTimeNil sets the value for EventDateTime to be an explicit nil

### UnsetEventDateTime
`func (o *ShipmentEvent) UnsetEventDateTime()`

UnsetEventDateTime ensures that no value is present for EventDateTime, not even an explicit nil
### GetShipmentEventTypeCode

`func (o *ShipmentEvent) GetShipmentEventTypeCode() ShipmentEventTypeCode`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *ShipmentEvent) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *ShipmentEvent) SetShipmentEventTypeCode(v ShipmentEventTypeCode)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentID

`func (o *ShipmentEvent) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *ShipmentEvent) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *ShipmentEvent) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetDocumentTypeCode

`func (o *ShipmentEvent) GetDocumentTypeCode() DocumentTypeCode`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *ShipmentEvent) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *ShipmentEvent) SetDocumentTypeCode(v DocumentTypeCode)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetShipmentInformationTypeCode

`func (o *ShipmentEvent) GetShipmentInformationTypeCode() ShipmentInformationType`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *ShipmentEvent) GetShipmentInformationTypeCodeOk() (*ShipmentInformationType, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *ShipmentEvent) SetShipmentInformationTypeCode(v ShipmentInformationType)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.

### HasShipmentInformationTypeCode

`func (o *ShipmentEvent) HasShipmentInformationTypeCode() bool`

HasShipmentInformationTypeCode returns a boolean if a field has been set.

### GetReason

`func (o *ShipmentEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ShipmentEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ShipmentEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ShipmentEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

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

### HasEventTypeCode

`func (o *ShipmentEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetShipmentID

`func (o *ShipmentEvent) GetShipmentID() interface{}`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ShipmentEvent) GetShipmentIDOk() (*interface{}, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ShipmentEvent) SetShipmentID(v interface{})`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *ShipmentEvent) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *ShipmentEvent) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *ShipmentEvent) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetReferences

`func (o *ShipmentEvent) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentEvent) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentEvent) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



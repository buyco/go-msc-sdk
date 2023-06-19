# BaseShipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**EventDateTime** | Pointer to **interface{}** | Value for eventDateTime must be the same value as eventCreatedDateTime  | [optional] 
**EventClassifierCode** | Pointer to **string** | Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)  | [optional] 
**ShipmentEventTypeCode** | [**ShipmentEventTypeCode**](ShipmentEventTypeCode.md) |  | 
**DocumentID** | **string** | The id of the object defined by the documentTypeCode.  | 
**DocumentTypeCode** | [**DocumentTypeCode**](DocumentTypeCode.md) |  | 
**ShipmentInformationTypeCode** | Pointer to [**ShipmentInformationType**](ShipmentInformationType.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code. For shipment events this can be - RECE (Received) - CONF (Confirmed) - ISSU (Issued) - APPR (Approved) - SUBM (Submitted) - SURR (Surrendered) - REJE (Rejected) - PENA (Pending approval)  Deprecated - use shipmentEventTypeCode instead  | [optional] 
**ShipmentID** | Pointer to **interface{}** | ID uniquely identifying a shipment.  Deprecated - this is replaced by documentID which can contain different values depending on the value of the documentTypeCode field  | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 

## Methods

### NewBaseShipmentEvent

`func NewBaseShipmentEvent(shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode, ) *BaseShipmentEvent`

NewBaseShipmentEvent instantiates a new BaseShipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseShipmentEventWithDefaults

`func NewBaseShipmentEventWithDefaults() *BaseShipmentEvent`

NewBaseShipmentEventWithDefaults instantiates a new BaseShipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BaseShipmentEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseShipmentEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseShipmentEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseShipmentEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventDateTime

`func (o *BaseShipmentEvent) GetEventDateTime() interface{}`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *BaseShipmentEvent) GetEventDateTimeOk() (*interface{}, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *BaseShipmentEvent) SetEventDateTime(v interface{})`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *BaseShipmentEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTimeNil

`func (o *BaseShipmentEvent) SetEventDateTimeNil(b bool)`

 SetEventDateTimeNil sets the value for EventDateTime to be an explicit nil

### UnsetEventDateTime
`func (o *BaseShipmentEvent) UnsetEventDateTime()`

UnsetEventDateTime ensures that no value is present for EventDateTime, not even an explicit nil
### GetEventClassifierCode

`func (o *BaseShipmentEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *BaseShipmentEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *BaseShipmentEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.

### HasEventClassifierCode

`func (o *BaseShipmentEvent) HasEventClassifierCode() bool`

HasEventClassifierCode returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *BaseShipmentEvent) GetShipmentEventTypeCode() ShipmentEventTypeCode`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *BaseShipmentEvent) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *BaseShipmentEvent) SetShipmentEventTypeCode(v ShipmentEventTypeCode)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentID

`func (o *BaseShipmentEvent) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *BaseShipmentEvent) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *BaseShipmentEvent) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetDocumentTypeCode

`func (o *BaseShipmentEvent) GetDocumentTypeCode() DocumentTypeCode`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *BaseShipmentEvent) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *BaseShipmentEvent) SetDocumentTypeCode(v DocumentTypeCode)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetShipmentInformationTypeCode

`func (o *BaseShipmentEvent) GetShipmentInformationTypeCode() ShipmentInformationType`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *BaseShipmentEvent) GetShipmentInformationTypeCodeOk() (*ShipmentInformationType, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *BaseShipmentEvent) SetShipmentInformationTypeCode(v ShipmentInformationType)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.

### HasShipmentInformationTypeCode

`func (o *BaseShipmentEvent) HasShipmentInformationTypeCode() bool`

HasShipmentInformationTypeCode returns a boolean if a field has been set.

### GetReason

`func (o *BaseShipmentEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BaseShipmentEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BaseShipmentEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BaseShipmentEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetEventTypeCode

`func (o *BaseShipmentEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseShipmentEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseShipmentEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseShipmentEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetShipmentID

`func (o *BaseShipmentEvent) GetShipmentID() interface{}`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *BaseShipmentEvent) GetShipmentIDOk() (*interface{}, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *BaseShipmentEvent) SetShipmentID(v interface{})`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *BaseShipmentEvent) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *BaseShipmentEvent) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *BaseShipmentEvent) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetReferences

`func (o *BaseShipmentEvent) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BaseShipmentEvent) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BaseShipmentEvent) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BaseShipmentEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



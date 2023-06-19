# OperationsTransportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventType** | **string** |  | 
**EventClassifierCode** | **string** | Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 
**TransportEventTypeCode** | [**TransportEventTypeCode**](TransportEventTypeCode.md) |  | 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/ | [optional] 
**VesselScheduleChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.  Deprecated - use changeRemark instead  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**TransportCallID** | Pointer to **string** |  | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code, for transport events this is either - ARRI (Arrival) - DEPA (Departure)  Deprecated - use transportEventTypeCode instead  | [optional] 

## Methods

### NewOperationsTransportEvent

`func NewOperationsTransportEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall, ) *OperationsTransportEvent`

NewOperationsTransportEvent instantiates a new OperationsTransportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationsTransportEventWithDefaults

`func NewOperationsTransportEventWithDefaults() *OperationsTransportEvent`

NewOperationsTransportEventWithDefaults instantiates a new OperationsTransportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *OperationsTransportEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *OperationsTransportEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *OperationsTransportEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *OperationsTransportEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *OperationsTransportEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *OperationsTransportEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *OperationsTransportEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *OperationsTransportEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OperationsTransportEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OperationsTransportEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *OperationsTransportEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *OperationsTransportEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *OperationsTransportEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *OperationsTransportEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *OperationsTransportEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *OperationsTransportEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetTransportEventTypeCode

`func (o *OperationsTransportEvent) GetTransportEventTypeCode() TransportEventTypeCode`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *OperationsTransportEvent) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *OperationsTransportEvent) SetTransportEventTypeCode(v TransportEventTypeCode)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *OperationsTransportEvent) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *OperationsTransportEvent) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *OperationsTransportEvent) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *OperationsTransportEvent) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetVesselScheduleChangeRemark

`func (o *OperationsTransportEvent) GetVesselScheduleChangeRemark() string`

GetVesselScheduleChangeRemark returns the VesselScheduleChangeRemark field if non-nil, zero value otherwise.

### GetVesselScheduleChangeRemarkOk

`func (o *OperationsTransportEvent) GetVesselScheduleChangeRemarkOk() (*string, bool)`

GetVesselScheduleChangeRemarkOk returns a tuple with the VesselScheduleChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselScheduleChangeRemark

`func (o *OperationsTransportEvent) SetVesselScheduleChangeRemark(v string)`

SetVesselScheduleChangeRemark sets VesselScheduleChangeRemark field to given value.

### HasVesselScheduleChangeRemark

`func (o *OperationsTransportEvent) HasVesselScheduleChangeRemark() bool`

HasVesselScheduleChangeRemark returns a boolean if a field has been set.

### GetChangeRemark

`func (o *OperationsTransportEvent) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *OperationsTransportEvent) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *OperationsTransportEvent) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *OperationsTransportEvent) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCallID

`func (o *OperationsTransportEvent) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *OperationsTransportEvent) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *OperationsTransportEvent) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *OperationsTransportEvent) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetTransportCall

`func (o *OperationsTransportEvent) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *OperationsTransportEvent) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *OperationsTransportEvent) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetEventTypeCode

`func (o *OperationsTransportEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *OperationsTransportEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *OperationsTransportEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *OperationsTransportEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



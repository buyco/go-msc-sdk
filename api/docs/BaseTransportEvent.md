# BaseTransportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**EventClassifierCode** | Pointer to **string** | Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)  | [optional] 
**TransportEventTypeCode** | [**TransportEventTypeCode**](TransportEventTypeCode.md) |  | 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/ | [optional] 
**VesselScheduleChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.  Deprecated - use changeRemark instead  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**TransportCallID** | Pointer to **string** |  | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code, for transport events this is either - ARRI (Arrival) - DEPA (Departure)  Deprecated - use transportEventTypeCode instead  | [optional] 

## Methods

### NewBaseTransportEvent

`func NewBaseTransportEvent(transportEventTypeCode TransportEventTypeCode, transportCall TransportCall, ) *BaseTransportEvent`

NewBaseTransportEvent instantiates a new BaseTransportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTransportEventWithDefaults

`func NewBaseTransportEventWithDefaults() *BaseTransportEvent`

NewBaseTransportEventWithDefaults instantiates a new BaseTransportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BaseTransportEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseTransportEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseTransportEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseTransportEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventClassifierCode

`func (o *BaseTransportEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *BaseTransportEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *BaseTransportEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.

### HasEventClassifierCode

`func (o *BaseTransportEvent) HasEventClassifierCode() bool`

HasEventClassifierCode returns a boolean if a field has been set.

### GetTransportEventTypeCode

`func (o *BaseTransportEvent) GetTransportEventTypeCode() TransportEventTypeCode`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *BaseTransportEvent) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *BaseTransportEvent) SetTransportEventTypeCode(v TransportEventTypeCode)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *BaseTransportEvent) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *BaseTransportEvent) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *BaseTransportEvent) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *BaseTransportEvent) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetVesselScheduleChangeRemark

`func (o *BaseTransportEvent) GetVesselScheduleChangeRemark() string`

GetVesselScheduleChangeRemark returns the VesselScheduleChangeRemark field if non-nil, zero value otherwise.

### GetVesselScheduleChangeRemarkOk

`func (o *BaseTransportEvent) GetVesselScheduleChangeRemarkOk() (*string, bool)`

GetVesselScheduleChangeRemarkOk returns a tuple with the VesselScheduleChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselScheduleChangeRemark

`func (o *BaseTransportEvent) SetVesselScheduleChangeRemark(v string)`

SetVesselScheduleChangeRemark sets VesselScheduleChangeRemark field to given value.

### HasVesselScheduleChangeRemark

`func (o *BaseTransportEvent) HasVesselScheduleChangeRemark() bool`

HasVesselScheduleChangeRemark returns a boolean if a field has been set.

### GetChangeRemark

`func (o *BaseTransportEvent) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *BaseTransportEvent) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *BaseTransportEvent) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *BaseTransportEvent) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCallID

`func (o *BaseTransportEvent) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *BaseTransportEvent) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *BaseTransportEvent) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *BaseTransportEvent) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetTransportCall

`func (o *BaseTransportEvent) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *BaseTransportEvent) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *BaseTransportEvent) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetEventTypeCode

`func (o *BaseTransportEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseTransportEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseTransportEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseTransportEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BaseEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventType** | **string** | The Event Type of the object - to be used as a discriminator.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventClassifierCode** | **string** | Code for the event classifier. Values can vary depending on eventType  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 

## Methods

### NewBaseEvent

`func NewBaseEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, ) *BaseEvent`

NewBaseEvent instantiates a new BaseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventWithDefaults

`func NewBaseEventWithDefaults() *BaseEvent`

NewBaseEventWithDefaults instantiates a new BaseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *BaseEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *BaseEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *BaseEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *BaseEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *BaseEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *BaseEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *BaseEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *BaseEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *BaseEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *BaseEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *BaseEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *BaseEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *BaseEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *BaseEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



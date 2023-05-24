# BaseEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The Event Type of the object - to be used as a discriminator.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | 
**EventClassifierCode** | **string** | Code for the event classifier. Values can vary depending on eventType  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 

## Methods

### NewBaseEventBody

`func NewBaseEventBody(eventType string, eventClassifierCode string, eventDateTime time.Time, ) *BaseEventBody`

NewBaseEventBody instantiates a new BaseEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventBodyWithDefaults

`func NewBaseEventBodyWithDefaults() *BaseEventBody`

NewBaseEventBodyWithDefaults instantiates a new BaseEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BaseEventBody) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEventBody) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEventBody) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *BaseEventBody) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *BaseEventBody) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *BaseEventBody) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *BaseEventBody) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *BaseEventBody) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *BaseEventBody) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



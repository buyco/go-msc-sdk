# BaseEventBodyAllOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDateTime** | Pointer to **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | [optional] 

## Methods

### NewBaseEventBodyAllOf2

`func NewBaseEventBodyAllOf2() *BaseEventBodyAllOf2`

NewBaseEventBodyAllOf2 instantiates a new BaseEventBodyAllOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventBodyAllOf2WithDefaults

`func NewBaseEventBodyAllOf2WithDefaults() *BaseEventBodyAllOf2`

NewBaseEventBodyAllOf2WithDefaults instantiates a new BaseEventBodyAllOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDateTime

`func (o *BaseEventBodyAllOf2) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *BaseEventBodyAllOf2) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *BaseEventBodyAllOf2) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *BaseEventBodyAllOf2) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



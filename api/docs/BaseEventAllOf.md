# BaseEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**EventCreatedDateTime** | Pointer to **time.Time** | The timestamp of when the event was created.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseEventAllOf

`func NewBaseEventAllOf() *BaseEventAllOf`

NewBaseEventAllOf instantiates a new BaseEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventAllOfWithDefaults

`func NewBaseEventAllOfWithDefaults() *BaseEventAllOf`

NewBaseEventAllOfWithDefaults instantiates a new BaseEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *BaseEventAllOf) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *BaseEventAllOf) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *BaseEventAllOf) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *BaseEventAllOf) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *BaseEventAllOf) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *BaseEventAllOf) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *BaseEventAllOf) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.

### HasEventCreatedDateTime

`func (o *BaseEventAllOf) HasEventCreatedDateTime() bool`

HasEventCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *BaseEventAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseEventAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseEventAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseEventAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BaseEventBodyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The Event Type of the object - to be used as a discriminator.  &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata  | [optional] 

## Methods

### NewBaseEventBodyAllOf

`func NewBaseEventBodyAllOf() *BaseEventBodyAllOf`

NewBaseEventBodyAllOf instantiates a new BaseEventBodyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventBodyAllOfWithDefaults

`func NewBaseEventBodyAllOfWithDefaults() *BaseEventBodyAllOf`

NewBaseEventBodyAllOfWithDefaults instantiates a new BaseEventBodyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BaseEventBodyAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEventBodyAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEventBodyAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseEventBodyAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



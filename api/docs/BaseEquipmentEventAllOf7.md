# BaseEquipmentEventAllOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code, for transport events this is either - LOAD (Loaded) - DISC (Discharged) - GTIN (Gated in) - GTOT (Gated out) - STUF (Stuffed) - STRP (Stripped)  Deprecated - use equipmentEventTypeCode instead  | [optional] 

## Methods

### NewBaseEquipmentEventAllOf7

`func NewBaseEquipmentEventAllOf7() *BaseEquipmentEventAllOf7`

NewBaseEquipmentEventAllOf7 instantiates a new BaseEquipmentEventAllOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEquipmentEventAllOf7WithDefaults

`func NewBaseEquipmentEventAllOf7WithDefaults() *BaseEquipmentEventAllOf7`

NewBaseEquipmentEventAllOf7WithDefaults instantiates a new BaseEquipmentEventAllOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeCode

`func (o *BaseEquipmentEventAllOf7) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseEquipmentEventAllOf7) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseEquipmentEventAllOf7) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseEquipmentEventAllOf7) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



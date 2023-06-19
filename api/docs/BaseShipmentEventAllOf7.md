# BaseShipmentEventAllOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeCode** | Pointer to **string** | Unique identifier for Event Type Code. For shipment events this can be - RECE (Received) - CONF (Confirmed) - ISSU (Issued) - APPR (Approved) - SUBM (Submitted) - SURR (Surrendered) - REJE (Rejected) - PENA (Pending approval)  Deprecated - use shipmentEventTypeCode instead  | [optional] 

## Methods

### NewBaseShipmentEventAllOf7

`func NewBaseShipmentEventAllOf7() *BaseShipmentEventAllOf7`

NewBaseShipmentEventAllOf7 instantiates a new BaseShipmentEventAllOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseShipmentEventAllOf7WithDefaults

`func NewBaseShipmentEventAllOf7WithDefaults() *BaseShipmentEventAllOf7`

NewBaseShipmentEventAllOf7WithDefaults instantiates a new BaseShipmentEventAllOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeCode

`func (o *BaseShipmentEventAllOf7) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseShipmentEventAllOf7) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseShipmentEventAllOf7) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseShipmentEventAllOf7) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EventSubscriptionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | Your server, where the API server should send the event data | 
**EventType** | Pointer to [**[]EventType**](EventType.md) | A list of event types. Default is all events | [optional] [default to ["EQUIPMENT","SHIPMENT","TRANSPORT","TRANSPORTEQUIPMENT"]]
**BookingReference** | Pointer to **string** | The identifier for a shipment, which is issued by and unique within each of the carriers. | [optional] 
**BillOfLadingNumber** | Pointer to **string** | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier&#39;s receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. | [optional] 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. If a container is not yet assigned to a shipment, the interface cannot return any information when an equipment reference is given as input. If a container is assigned to an (active) shipment, the interface returns information on the active shipment. | [optional] 

## Methods

### NewEventSubscriptionBody

`func NewEventSubscriptionBody(callbackUrl string, ) *EventSubscriptionBody`

NewEventSubscriptionBody instantiates a new EventSubscriptionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionBodyWithDefaults

`func NewEventSubscriptionBodyWithDefaults() *EventSubscriptionBody`

NewEventSubscriptionBodyWithDefaults instantiates a new EventSubscriptionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *EventSubscriptionBody) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *EventSubscriptionBody) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *EventSubscriptionBody) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetEventType

`func (o *EventSubscriptionBody) GetEventType() []EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventSubscriptionBody) GetEventTypeOk() (*[]EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventSubscriptionBody) SetEventType(v []EventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventSubscriptionBody) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetBookingReference

`func (o *EventSubscriptionBody) GetBookingReference() string`

GetBookingReference returns the BookingReference field if non-nil, zero value otherwise.

### GetBookingReferenceOk

`func (o *EventSubscriptionBody) GetBookingReferenceOk() (*string, bool)`

GetBookingReferenceOk returns a tuple with the BookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingReference

`func (o *EventSubscriptionBody) SetBookingReference(v string)`

SetBookingReference sets BookingReference field to given value.

### HasBookingReference

`func (o *EventSubscriptionBody) HasBookingReference() bool`

HasBookingReference returns a boolean if a field has been set.

### GetBillOfLadingNumber

`func (o *EventSubscriptionBody) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *EventSubscriptionBody) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *EventSubscriptionBody) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *EventSubscriptionBody) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### GetEquipmentReference

`func (o *EventSubscriptionBody) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *EventSubscriptionBody) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *EventSubscriptionBody) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *EventSubscriptionBody) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



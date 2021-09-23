# ShipmentEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**ShipmentInformationTypeCode** | **string** | The code to identify the type of information that is related to the shipment. | 

## Methods

### NewShipmentEventAllOf

`func NewShipmentEventAllOf(shipmentInformationTypeCode string, ) *ShipmentEventAllOf`

NewShipmentEventAllOf instantiates a new ShipmentEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentEventAllOfWithDefaults

`func NewShipmentEventAllOfWithDefaults() *ShipmentEventAllOf`

NewShipmentEventAllOfWithDefaults instantiates a new ShipmentEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ShipmentEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ShipmentEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ShipmentEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ShipmentEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetShipmentInformationTypeCode

`func (o *ShipmentEventAllOf) GetShipmentInformationTypeCode() string`

GetShipmentInformationTypeCode returns the ShipmentInformationTypeCode field if non-nil, zero value otherwise.

### GetShipmentInformationTypeCodeOk

`func (o *ShipmentEventAllOf) GetShipmentInformationTypeCodeOk() (*string, bool)`

GetShipmentInformationTypeCodeOk returns a tuple with the ShipmentInformationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentInformationTypeCode

`func (o *ShipmentEventAllOf) SetShipmentInformationTypeCode(v string)`

SetShipmentInformationTypeCode sets ShipmentInformationTypeCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



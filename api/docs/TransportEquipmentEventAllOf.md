# TransportEquipmentEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**TransportReference** | **string** | The reference for the transport, e.g. when the mode of transport is a vessel, the transport reference will be the vessel IMO number. | 
**TransportLegReference** | **string** | The transport leg reference will be specific per mode of transport: - Vessel: Voyage number as specified by the vessel operator - Truck: Not yet specified - Rail: Not yet specified - Barge: Not yet specified.  | 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. If a container is not yet assigned to a shipment, the interface cannot return any information when an equipment reference is given as input. If a container is assigned to an (active) shipment, the interface returns information on the active shipment. | [optional] 
**FacilityTypeCode** | **string** | The code to identify the specific type of facility. | 
**UNLocationCode** | **string** | The UN Location Code identifies a location in the sense of a city/a town/a village, being the smaller administrative area existing as defined by the competent national authority in each country. | 
**FacilityCode** | **string** | The code used for identifying the specific facility. | 
**OtherFacility** | Pointer to **string** | An alternative way to capture the facility when no standardized DCSA facility code can be found. | [optional] 
**EmptyIndicatorCode** | **string** | Code to denote whether the equipment is empty or laden. | 
**ModeOfTransportCode** | Pointer to **string** | A code specifying a type of transport mode. | [optional] 

## Methods

### NewTransportEquipmentEventAllOf

`func NewTransportEquipmentEventAllOf(transportReference string, transportLegReference string, facilityTypeCode string, uNLocationCode string, facilityCode string, emptyIndicatorCode string, ) *TransportEquipmentEventAllOf`

NewTransportEquipmentEventAllOf instantiates a new TransportEquipmentEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEquipmentEventAllOfWithDefaults

`func NewTransportEquipmentEventAllOfWithDefaults() *TransportEquipmentEventAllOf`

NewTransportEquipmentEventAllOfWithDefaults instantiates a new TransportEquipmentEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TransportEquipmentEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransportEquipmentEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransportEquipmentEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TransportEquipmentEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTransportReference

`func (o *TransportEquipmentEventAllOf) GetTransportReference() string`

GetTransportReference returns the TransportReference field if non-nil, zero value otherwise.

### GetTransportReferenceOk

`func (o *TransportEquipmentEventAllOf) GetTransportReferenceOk() (*string, bool)`

GetTransportReferenceOk returns a tuple with the TransportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReference

`func (o *TransportEquipmentEventAllOf) SetTransportReference(v string)`

SetTransportReference sets TransportReference field to given value.


### GetTransportLegReference

`func (o *TransportEquipmentEventAllOf) GetTransportLegReference() string`

GetTransportLegReference returns the TransportLegReference field if non-nil, zero value otherwise.

### GetTransportLegReferenceOk

`func (o *TransportEquipmentEventAllOf) GetTransportLegReferenceOk() (*string, bool)`

GetTransportLegReferenceOk returns a tuple with the TransportLegReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportLegReference

`func (o *TransportEquipmentEventAllOf) SetTransportLegReference(v string)`

SetTransportLegReference sets TransportLegReference field to given value.


### GetEquipmentReference

`func (o *TransportEquipmentEventAllOf) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *TransportEquipmentEventAllOf) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *TransportEquipmentEventAllOf) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *TransportEquipmentEventAllOf) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetFacilityTypeCode

`func (o *TransportEquipmentEventAllOf) GetFacilityTypeCode() string`

GetFacilityTypeCode returns the FacilityTypeCode field if non-nil, zero value otherwise.

### GetFacilityTypeCodeOk

`func (o *TransportEquipmentEventAllOf) GetFacilityTypeCodeOk() (*string, bool)`

GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityTypeCode

`func (o *TransportEquipmentEventAllOf) SetFacilityTypeCode(v string)`

SetFacilityTypeCode sets FacilityTypeCode field to given value.


### GetUNLocationCode

`func (o *TransportEquipmentEventAllOf) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *TransportEquipmentEventAllOf) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *TransportEquipmentEventAllOf) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.


### GetFacilityCode

`func (o *TransportEquipmentEventAllOf) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *TransportEquipmentEventAllOf) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *TransportEquipmentEventAllOf) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.


### GetOtherFacility

`func (o *TransportEquipmentEventAllOf) GetOtherFacility() string`

GetOtherFacility returns the OtherFacility field if non-nil, zero value otherwise.

### GetOtherFacilityOk

`func (o *TransportEquipmentEventAllOf) GetOtherFacilityOk() (*string, bool)`

GetOtherFacilityOk returns a tuple with the OtherFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFacility

`func (o *TransportEquipmentEventAllOf) SetOtherFacility(v string)`

SetOtherFacility sets OtherFacility field to given value.

### HasOtherFacility

`func (o *TransportEquipmentEventAllOf) HasOtherFacility() bool`

HasOtherFacility returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *TransportEquipmentEventAllOf) GetEmptyIndicatorCode() string`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *TransportEquipmentEventAllOf) GetEmptyIndicatorCodeOk() (*string, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *TransportEquipmentEventAllOf) SetEmptyIndicatorCode(v string)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetModeOfTransportCode

`func (o *TransportEquipmentEventAllOf) GetModeOfTransportCode() string`

GetModeOfTransportCode returns the ModeOfTransportCode field if non-nil, zero value otherwise.

### GetModeOfTransportCodeOk

`func (o *TransportEquipmentEventAllOf) GetModeOfTransportCodeOk() (*string, bool)`

GetModeOfTransportCodeOk returns a tuple with the ModeOfTransportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransportCode

`func (o *TransportEquipmentEventAllOf) SetModeOfTransportCode(v string)`

SetModeOfTransportCode sets ModeOfTransportCode field to given value.

### HasModeOfTransportCode

`func (o *TransportEquipmentEventAllOf) HasModeOfTransportCode() bool`

HasModeOfTransportCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



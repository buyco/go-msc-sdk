# TransportEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**TransportReference** | **string** | The reference for the transport, e.g. when the mode of transport is a vessel, the transport reference will be the vessel IMO number. | 
**TransportLegReference** | **string** | The transport leg reference will be specific per mode of transport: - Vessel: Voyage number as specified by the vessel operator - Truck: Not yet specified - Rail: Not yet specified - Barge: Not yet specified.  | 
**FacilityTypeCode** | **string** | The code to identify the specific type of facility. | 
**UNLocationCode** | **string** | The UN Location Code identifies a location in the sense of a city/a town/a village, being the smaller administrative area existing as defined by the competent national authority in each country. | 
**FacilityCode** | **string** | The code used for identifying the specific facility. | 
**OtherFacility** | Pointer to **string** | An alternative way to capture the facility when no standardized DCSA facility code can be found. | [optional] 
**ModeOfTransportCode** | Pointer to **string** | A code specifying a type of transport mode. | [optional] 

## Methods

### NewTransportEventAllOf

`func NewTransportEventAllOf(transportReference string, transportLegReference string, facilityTypeCode string, uNLocationCode string, facilityCode string, ) *TransportEventAllOf`

NewTransportEventAllOf instantiates a new TransportEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEventAllOfWithDefaults

`func NewTransportEventAllOfWithDefaults() *TransportEventAllOf`

NewTransportEventAllOfWithDefaults instantiates a new TransportEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TransportEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransportEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransportEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TransportEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTransportReference

`func (o *TransportEventAllOf) GetTransportReference() string`

GetTransportReference returns the TransportReference field if non-nil, zero value otherwise.

### GetTransportReferenceOk

`func (o *TransportEventAllOf) GetTransportReferenceOk() (*string, bool)`

GetTransportReferenceOk returns a tuple with the TransportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReference

`func (o *TransportEventAllOf) SetTransportReference(v string)`

SetTransportReference sets TransportReference field to given value.


### GetTransportLegReference

`func (o *TransportEventAllOf) GetTransportLegReference() string`

GetTransportLegReference returns the TransportLegReference field if non-nil, zero value otherwise.

### GetTransportLegReferenceOk

`func (o *TransportEventAllOf) GetTransportLegReferenceOk() (*string, bool)`

GetTransportLegReferenceOk returns a tuple with the TransportLegReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportLegReference

`func (o *TransportEventAllOf) SetTransportLegReference(v string)`

SetTransportLegReference sets TransportLegReference field to given value.


### GetFacilityTypeCode

`func (o *TransportEventAllOf) GetFacilityTypeCode() string`

GetFacilityTypeCode returns the FacilityTypeCode field if non-nil, zero value otherwise.

### GetFacilityTypeCodeOk

`func (o *TransportEventAllOf) GetFacilityTypeCodeOk() (*string, bool)`

GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityTypeCode

`func (o *TransportEventAllOf) SetFacilityTypeCode(v string)`

SetFacilityTypeCode sets FacilityTypeCode field to given value.


### GetUNLocationCode

`func (o *TransportEventAllOf) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *TransportEventAllOf) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *TransportEventAllOf) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.


### GetFacilityCode

`func (o *TransportEventAllOf) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *TransportEventAllOf) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *TransportEventAllOf) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.


### GetOtherFacility

`func (o *TransportEventAllOf) GetOtherFacility() string`

GetOtherFacility returns the OtherFacility field if non-nil, zero value otherwise.

### GetOtherFacilityOk

`func (o *TransportEventAllOf) GetOtherFacilityOk() (*string, bool)`

GetOtherFacilityOk returns a tuple with the OtherFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFacility

`func (o *TransportEventAllOf) SetOtherFacility(v string)`

SetOtherFacility sets OtherFacility field to given value.

### HasOtherFacility

`func (o *TransportEventAllOf) HasOtherFacility() bool`

HasOtherFacility returns a boolean if a field has been set.

### GetModeOfTransportCode

`func (o *TransportEventAllOf) GetModeOfTransportCode() string`

GetModeOfTransportCode returns the ModeOfTransportCode field if non-nil, zero value otherwise.

### GetModeOfTransportCodeOk

`func (o *TransportEventAllOf) GetModeOfTransportCodeOk() (*string, bool)`

GetModeOfTransportCodeOk returns a tuple with the ModeOfTransportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransportCode

`func (o *TransportEventAllOf) SetModeOfTransportCode(v string)`

SetModeOfTransportCode sets ModeOfTransportCode field to given value.

### HasModeOfTransportCode

`func (o *TransportEventAllOf) HasModeOfTransportCode() bool`

HasModeOfTransportCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



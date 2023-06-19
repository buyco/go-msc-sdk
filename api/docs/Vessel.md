# Vessel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VesselIMONumber** | **string** | The unique reference for a registered Vessel. The reference is the International Maritime Organisation (IMO) number, also sometimes known as the Lloyd&amp;apos;s register code, which does not change during the lifetime of the vessel  | 
**VesselName** | Pointer to **string** | The name of the Vessel given by the Vessel Operator and registered with IMO.  | [optional] 
**VesselFlag** | Pointer to **string** | The flag of the nation whose laws the vessel is registered under. This is the ISO 3166 two-letter country code  | [optional] 
**VesselCallSignNumber** | Pointer to **string** | A unique alphanumeric identity that belongs to the vessel and is assigned by the International Telecommunication Union (ITU). It consists of a threeletter alphanumeric prefix that indicates nationality, followed by one to four characters to identify the individual vessel. For instance, vessels registered under Denmark are assigned the prefix ranges 5PA-5QZ, OUAOZZ, and XPA-XPZ. The Call Sign changes whenever a vessel changes its flag.  | [optional] 
**VesselOperatorCarrierCode** | Pointer to **string** | The carrier who is in charge of the vessel operation based on either the SMDG or SCAC code lists  | [optional] 
**VesselOperatorCarrierCodeListProvider** | Pointer to [**VesselOperatorCarrierCodeListProvider**](VesselOperatorCarrierCodeListProvider.md) |  | [optional] 

## Methods

### NewVessel

`func NewVessel(vesselIMONumber string, ) *Vessel`

NewVessel instantiates a new Vessel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVesselWithDefaults

`func NewVesselWithDefaults() *Vessel`

NewVesselWithDefaults instantiates a new Vessel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVesselIMONumber

`func (o *Vessel) GetVesselIMONumber() string`

GetVesselIMONumber returns the VesselIMONumber field if non-nil, zero value otherwise.

### GetVesselIMONumberOk

`func (o *Vessel) GetVesselIMONumberOk() (*string, bool)`

GetVesselIMONumberOk returns a tuple with the VesselIMONumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselIMONumber

`func (o *Vessel) SetVesselIMONumber(v string)`

SetVesselIMONumber sets VesselIMONumber field to given value.


### GetVesselName

`func (o *Vessel) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *Vessel) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *Vessel) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *Vessel) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### GetVesselFlag

`func (o *Vessel) GetVesselFlag() string`

GetVesselFlag returns the VesselFlag field if non-nil, zero value otherwise.

### GetVesselFlagOk

`func (o *Vessel) GetVesselFlagOk() (*string, bool)`

GetVesselFlagOk returns a tuple with the VesselFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselFlag

`func (o *Vessel) SetVesselFlag(v string)`

SetVesselFlag sets VesselFlag field to given value.

### HasVesselFlag

`func (o *Vessel) HasVesselFlag() bool`

HasVesselFlag returns a boolean if a field has been set.

### GetVesselCallSignNumber

`func (o *Vessel) GetVesselCallSignNumber() string`

GetVesselCallSignNumber returns the VesselCallSignNumber field if non-nil, zero value otherwise.

### GetVesselCallSignNumberOk

`func (o *Vessel) GetVesselCallSignNumberOk() (*string, bool)`

GetVesselCallSignNumberOk returns a tuple with the VesselCallSignNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselCallSignNumber

`func (o *Vessel) SetVesselCallSignNumber(v string)`

SetVesselCallSignNumber sets VesselCallSignNumber field to given value.

### HasVesselCallSignNumber

`func (o *Vessel) HasVesselCallSignNumber() bool`

HasVesselCallSignNumber returns a boolean if a field has been set.

### GetVesselOperatorCarrierCode

`func (o *Vessel) GetVesselOperatorCarrierCode() string`

GetVesselOperatorCarrierCode returns the VesselOperatorCarrierCode field if non-nil, zero value otherwise.

### GetVesselOperatorCarrierCodeOk

`func (o *Vessel) GetVesselOperatorCarrierCodeOk() (*string, bool)`

GetVesselOperatorCarrierCodeOk returns a tuple with the VesselOperatorCarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselOperatorCarrierCode

`func (o *Vessel) SetVesselOperatorCarrierCode(v string)`

SetVesselOperatorCarrierCode sets VesselOperatorCarrierCode field to given value.

### HasVesselOperatorCarrierCode

`func (o *Vessel) HasVesselOperatorCarrierCode() bool`

HasVesselOperatorCarrierCode returns a boolean if a field has been set.

### GetVesselOperatorCarrierCodeListProvider

`func (o *Vessel) GetVesselOperatorCarrierCodeListProvider() VesselOperatorCarrierCodeListProvider`

GetVesselOperatorCarrierCodeListProvider returns the VesselOperatorCarrierCodeListProvider field if non-nil, zero value otherwise.

### GetVesselOperatorCarrierCodeListProviderOk

`func (o *Vessel) GetVesselOperatorCarrierCodeListProviderOk() (*VesselOperatorCarrierCodeListProvider, bool)`

GetVesselOperatorCarrierCodeListProviderOk returns a tuple with the VesselOperatorCarrierCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselOperatorCarrierCodeListProvider

`func (o *Vessel) SetVesselOperatorCarrierCodeListProvider(v VesselOperatorCarrierCodeListProvider)`

SetVesselOperatorCarrierCodeListProvider sets VesselOperatorCarrierCodeListProvider field to given value.

### HasVesselOperatorCarrierCodeListProvider

`func (o *Vessel) HasVesselOperatorCarrierCodeListProvider() bool`

HasVesselOperatorCarrierCodeListProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



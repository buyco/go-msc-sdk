# TransportCallAllOf13Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationName** | Pointer to **string** | The name of the location. | [optional] 
**Latitude** | Pointer to **string** | Geographic coordinate that specifies the north–south position of a point on the Earth&amp;apos;s surface. | [optional] 
**Longitude** | Pointer to **string** | Geographic coordinate that specifies the east–west position of a point on the Earth&amp;apos;s surface. | [optional] 
**UNLocationCode** | Pointer to **string** | The UN Location code specifying where the place is located. | [optional] 
**FacilityCode** | Pointer to **string** | The code used for identifying the specific facility. This code does &lt;b&gt;not&lt;/b&gt; include the UN Location Code.  | [optional] 
**FacilityCodeListProvider** | Pointer to [**FacilityCodeListProvider**](FacilityCodeListProvider.md) |  | [optional] 
**Address** | Pointer to [**LocationAllOf3Address**](LocationAllOf3Address.md) |  | [optional] 

## Methods

### NewTransportCallAllOf13Location

`func NewTransportCallAllOf13Location() *TransportCallAllOf13Location`

NewTransportCallAllOf13Location instantiates a new TransportCallAllOf13Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportCallAllOf13LocationWithDefaults

`func NewTransportCallAllOf13LocationWithDefaults() *TransportCallAllOf13Location`

NewTransportCallAllOf13LocationWithDefaults instantiates a new TransportCallAllOf13Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationName

`func (o *TransportCallAllOf13Location) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *TransportCallAllOf13Location) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *TransportCallAllOf13Location) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *TransportCallAllOf13Location) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetLatitude

`func (o *TransportCallAllOf13Location) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TransportCallAllOf13Location) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TransportCallAllOf13Location) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TransportCallAllOf13Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *TransportCallAllOf13Location) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TransportCallAllOf13Location) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TransportCallAllOf13Location) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TransportCallAllOf13Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetUNLocationCode

`func (o *TransportCallAllOf13Location) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *TransportCallAllOf13Location) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *TransportCallAllOf13Location) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.

### HasUNLocationCode

`func (o *TransportCallAllOf13Location) HasUNLocationCode() bool`

HasUNLocationCode returns a boolean if a field has been set.

### GetFacilityCode

`func (o *TransportCallAllOf13Location) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *TransportCallAllOf13Location) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *TransportCallAllOf13Location) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *TransportCallAllOf13Location) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityCodeListProvider

`func (o *TransportCallAllOf13Location) GetFacilityCodeListProvider() FacilityCodeListProvider`

GetFacilityCodeListProvider returns the FacilityCodeListProvider field if non-nil, zero value otherwise.

### GetFacilityCodeListProviderOk

`func (o *TransportCallAllOf13Location) GetFacilityCodeListProviderOk() (*FacilityCodeListProvider, bool)`

GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCodeListProvider

`func (o *TransportCallAllOf13Location) SetFacilityCodeListProvider(v FacilityCodeListProvider)`

SetFacilityCodeListProvider sets FacilityCodeListProvider field to given value.

### HasFacilityCodeListProvider

`func (o *TransportCallAllOf13Location) HasFacilityCodeListProvider() bool`

HasFacilityCodeListProvider returns a boolean if a field has been set.

### GetAddress

`func (o *TransportCallAllOf13Location) GetAddress() LocationAllOf3Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransportCallAllOf13Location) GetAddressOk() (*LocationAllOf3Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransportCallAllOf13Location) SetAddress(v LocationAllOf3Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransportCallAllOf13Location) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



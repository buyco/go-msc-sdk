/*
DCSA OpenAPI specification for Track & Trace

Provides equipment actual milestones along with Estimated Time of Arrival following DCSA standards

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address An object for storing address related information
type Address struct {
	// Name of the address
	Name *string `json:"name,omitempty"`
	// The name of the street of the party’s address.
	Street *string `json:"street,omitempty"`
	// The number of the street of the party’s address.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// The floor of the party’s street number.
	Floor *string `json:"floor,omitempty"`
	// The post code of the party’s address.
	PostCode *string `json:"postCode,omitempty"`
	// The city name of the party’s address.
	City *string `json:"city,omitempty"`
	// The state/region of the party’s address.
	StateRegion *string `json:"stateRegion,omitempty"`
	// The country of the party’s address.
	Country *string `json:"country,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Address) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Address) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Address) SetName(v string) {
	o.Name = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *Address) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *Address) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *Address) SetStreet(v string) {
	o.Street = &v
}

// GetStreetNumber returns the StreetNumber field value if set, zero value otherwise.
func (o *Address) GetStreetNumber() string {
	if o == nil || IsNil(o.StreetNumber) {
		var ret string
		return ret
	}
	return *o.StreetNumber
}

// GetStreetNumberOk returns a tuple with the StreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.StreetNumber) {
		return nil, false
	}
	return o.StreetNumber, true
}

// HasStreetNumber returns a boolean if a field has been set.
func (o *Address) HasStreetNumber() bool {
	if o != nil && !IsNil(o.StreetNumber) {
		return true
	}

	return false
}

// SetStreetNumber gets a reference to the given string and assigns it to the StreetNumber field.
func (o *Address) SetStreetNumber(v string) {
	o.StreetNumber = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *Address) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *Address) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *Address) SetFloor(v string) {
	o.Floor = &v
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *Address) GetPostCode() string {
	if o == nil || IsNil(o.PostCode) {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *Address) HasPostCode() bool {
	if o != nil && !IsNil(o.PostCode) {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *Address) SetPostCode(v string) {
	o.PostCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City = &v
}

// GetStateRegion returns the StateRegion field value if set, zero value otherwise.
func (o *Address) GetStateRegion() string {
	if o == nil || IsNil(o.StateRegion) {
		var ret string
		return ret
	}
	return *o.StateRegion
}

// GetStateRegionOk returns a tuple with the StateRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateRegionOk() (*string, bool) {
	if o == nil || IsNil(o.StateRegion) {
		return nil, false
	}
	return o.StateRegion, true
}

// HasStateRegion returns a boolean if a field has been set.
func (o *Address) HasStateRegion() bool {
	if o != nil && !IsNil(o.StateRegion) {
		return true
	}

	return false
}

// SetStateRegion gets a reference to the given string and assigns it to the StateRegion field.
func (o *Address) SetStateRegion(v string) {
	o.StateRegion = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Address) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.StreetNumber) {
		toSerialize["streetNumber"] = o.StreetNumber
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.PostCode) {
		toSerialize["postCode"] = o.PostCode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.StateRegion) {
		toSerialize["stateRegion"] = o.StateRegion
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

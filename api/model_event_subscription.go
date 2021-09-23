/*
DCSA Track and Trace API

API specifications for the Track and Trace interface standard issued by DCSA.org.

API version: 1.2.0
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventSubscription struct for EventSubscription
type EventSubscription struct {
	SubscriptionID *string `json:"subscriptionID,omitempty"`
	// Your server, where the API server should send the event data
	CallbackUrl string `json:"callbackUrl"`
	// A list of event types. Default is all events
	EventType *[]EventType `json:"eventType,omitempty"`
	// The identifier for a shipment, which is issued by and unique within each of the carriers.
	BookingReference *string `json:"bookingReference,omitempty"`
	// Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier's receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery.
	BillOfLadingNumber *string `json:"billOfLadingNumber,omitempty"`
	// The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. If a container is not yet assigned to a shipment, the interface cannot return any information when an equipment reference is given as input. If a container is assigned to an (active) shipment, the interface returns information on the active shipment.
	EquipmentReference *string `json:"equipmentReference,omitempty"`
}

// NewEventSubscription instantiates a new EventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscription(callbackUrl string) *EventSubscription {
	this := EventSubscription{}
	this.CallbackUrl = callbackUrl
	return &this
}

// NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscriptionWithDefaults() *EventSubscription {
	this := EventSubscription{}
	return &this
}

// GetSubscriptionID returns the SubscriptionID field value if set, zero value otherwise.
func (o *EventSubscription) GetSubscriptionID() string {
	if o == nil || o.SubscriptionID == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionID
}

// GetSubscriptionIDOk returns a tuple with the SubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetSubscriptionIDOk() (*string, bool) {
	if o == nil || o.SubscriptionID == nil {
		return nil, false
	}
	return o.SubscriptionID, true
}

// HasSubscriptionID returns a boolean if a field has been set.
func (o *EventSubscription) HasSubscriptionID() bool {
	if o != nil && o.SubscriptionID != nil {
		return true
	}

	return false
}

// SetSubscriptionID gets a reference to the given string and assigns it to the SubscriptionID field.
func (o *EventSubscription) SetSubscriptionID(v string) {
	o.SubscriptionID = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *EventSubscription) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *EventSubscription) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventSubscription) GetEventType() []EventType {
	if o == nil || o.EventType == nil {
		var ret []EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEventTypeOk() (*[]EventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventSubscription) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []EventType and assigns it to the EventType field.
func (o *EventSubscription) SetEventType(v []EventType) {
	o.EventType = &v
}

// GetBookingReference returns the BookingReference field value if set, zero value otherwise.
func (o *EventSubscription) GetBookingReference() string {
	if o == nil || o.BookingReference == nil {
		var ret string
		return ret
	}
	return *o.BookingReference
}

// GetBookingReferenceOk returns a tuple with the BookingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetBookingReferenceOk() (*string, bool) {
	if o == nil || o.BookingReference == nil {
		return nil, false
	}
	return o.BookingReference, true
}

// HasBookingReference returns a boolean if a field has been set.
func (o *EventSubscription) HasBookingReference() bool {
	if o != nil && o.BookingReference != nil {
		return true
	}

	return false
}

// SetBookingReference gets a reference to the given string and assigns it to the BookingReference field.
func (o *EventSubscription) SetBookingReference(v string) {
	o.BookingReference = &v
}

// GetBillOfLadingNumber returns the BillOfLadingNumber field value if set, zero value otherwise.
func (o *EventSubscription) GetBillOfLadingNumber() string {
	if o == nil || o.BillOfLadingNumber == nil {
		var ret string
		return ret
	}
	return *o.BillOfLadingNumber
}

// GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetBillOfLadingNumberOk() (*string, bool) {
	if o == nil || o.BillOfLadingNumber == nil {
		return nil, false
	}
	return o.BillOfLadingNumber, true
}

// HasBillOfLadingNumber returns a boolean if a field has been set.
func (o *EventSubscription) HasBillOfLadingNumber() bool {
	if o != nil && o.BillOfLadingNumber != nil {
		return true
	}

	return false
}

// SetBillOfLadingNumber gets a reference to the given string and assigns it to the BillOfLadingNumber field.
func (o *EventSubscription) SetBillOfLadingNumber(v string) {
	o.BillOfLadingNumber = &v
}

// GetEquipmentReference returns the EquipmentReference field value if set, zero value otherwise.
func (o *EventSubscription) GetEquipmentReference() string {
	if o == nil || o.EquipmentReference == nil {
		var ret string
		return ret
	}
	return *o.EquipmentReference
}

// GetEquipmentReferenceOk returns a tuple with the EquipmentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEquipmentReferenceOk() (*string, bool) {
	if o == nil || o.EquipmentReference == nil {
		return nil, false
	}
	return o.EquipmentReference, true
}

// HasEquipmentReference returns a boolean if a field has been set.
func (o *EventSubscription) HasEquipmentReference() bool {
	if o != nil && o.EquipmentReference != nil {
		return true
	}

	return false
}

// SetEquipmentReference gets a reference to the given string and assigns it to the EquipmentReference field.
func (o *EventSubscription) SetEquipmentReference(v string) {
	o.EquipmentReference = &v
}

func (o EventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionID != nil {
		toSerialize["subscriptionID"] = o.SubscriptionID
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.BookingReference != nil {
		toSerialize["bookingReference"] = o.BookingReference
	}
	if o.BillOfLadingNumber != nil {
		toSerialize["billOfLadingNumber"] = o.BillOfLadingNumber
	}
	if o.EquipmentReference != nil {
		toSerialize["equipmentReference"] = o.EquipmentReference
	}
	return json.Marshal(toSerialize)
}

type NullableEventSubscription struct {
	value *EventSubscription
	isSet bool
}

func (v NullableEventSubscription) Get() *EventSubscription {
	return v.value
}

func (v *NullableEventSubscription) Set(val *EventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscription(val *EventSubscription) *NullableEventSubscription {
	return &NullableEventSubscription{value: val, isSet: true}
}

func (v NullableEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

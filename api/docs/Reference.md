# Reference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | [**ReferenceType**](ReferenceType.md) |  | 
**ReferenceValue** | **string** | The actual value of the reference. | 

## Methods

### NewReference

`func NewReference(referenceType ReferenceType, referenceValue string, ) *Reference`

NewReference instantiates a new Reference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceWithDefaults

`func NewReferenceWithDefaults() *Reference`

NewReferenceWithDefaults instantiates a new Reference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *Reference) GetReferenceType() ReferenceType`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *Reference) GetReferenceTypeOk() (*ReferenceType, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *Reference) SetReferenceType(v ReferenceType)`

SetReferenceType sets ReferenceType field to given value.


### GetReferenceValue

`func (o *Reference) GetReferenceValue() string`

GetReferenceValue returns the ReferenceValue field if non-nil, zero value otherwise.

### GetReferenceValueOk

`func (o *Reference) GetReferenceValueOk() (*string, bool)`

GetReferenceValueOk returns a tuple with the ReferenceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceValue

`func (o *Reference) SetReferenceValue(v string)`

SetReferenceValue sets ReferenceValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



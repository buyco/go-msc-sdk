# TransportEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The &lt;b&gt;documentReferenceType&lt;/b&gt;-field is used to describe where the &lt;b&gt;documentReferenceValue&lt;/b&gt;-field is pointing to. | [optional] 

## Methods

### NewTransportEventAllOf

`func NewTransportEventAllOf() *TransportEventAllOf`

NewTransportEventAllOf instantiates a new TransportEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEventAllOfWithDefaults

`func NewTransportEventAllOfWithDefaults() *TransportEventAllOf`

NewTransportEventAllOfWithDefaults instantiates a new TransportEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentReferences

`func (o *TransportEventAllOf) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *TransportEventAllOf) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *TransportEventAllOf) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *TransportEventAllOf) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



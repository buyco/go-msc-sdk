# InlineResponseDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethod** | **interface{}** |  | 
**RequestUri** | **interface{}** |  | 
**Errors** | [**[]SubErrorsInner**](SubErrorsInner.md) |  | 
**StatusCode** | **int32** | The HTTP status code | 
**StatusCodeText** | **string** | The textual representation of the response status. | 
**ErrorDateTime** | **string** | The date and time (in ISO 8601 format) the error occurred. | 

## Methods

### NewInlineResponseDefault

`func NewInlineResponseDefault(httpMethod interface{}, requestUri interface{}, errors []SubErrorsInner, statusCode int32, statusCodeText string, errorDateTime string, ) *InlineResponseDefault`

NewInlineResponseDefault instantiates a new InlineResponseDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponseDefaultWithDefaults

`func NewInlineResponseDefaultWithDefaults() *InlineResponseDefault`

NewInlineResponseDefaultWithDefaults instantiates a new InlineResponseDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpMethod

`func (o *InlineResponseDefault) GetHttpMethod() interface{}`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *InlineResponseDefault) GetHttpMethodOk() (*interface{}, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *InlineResponseDefault) SetHttpMethod(v interface{})`

SetHttpMethod sets HttpMethod field to given value.


### SetHttpMethodNil

`func (o *InlineResponseDefault) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *InlineResponseDefault) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetRequestUri

`func (o *InlineResponseDefault) GetRequestUri() interface{}`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *InlineResponseDefault) GetRequestUriOk() (*interface{}, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *InlineResponseDefault) SetRequestUri(v interface{})`

SetRequestUri sets RequestUri field to given value.


### SetRequestUriNil

`func (o *InlineResponseDefault) SetRequestUriNil(b bool)`

 SetRequestUriNil sets the value for RequestUri to be an explicit nil

### UnsetRequestUri
`func (o *InlineResponseDefault) UnsetRequestUri()`

UnsetRequestUri ensures that no value is present for RequestUri, not even an explicit nil
### GetErrors

`func (o *InlineResponseDefault) GetErrors() []SubErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponseDefault) GetErrorsOk() (*[]SubErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponseDefault) SetErrors(v []SubErrorsInner)`

SetErrors sets Errors field to given value.


### GetStatusCode

`func (o *InlineResponseDefault) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InlineResponseDefault) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InlineResponseDefault) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusCodeText

`func (o *InlineResponseDefault) GetStatusCodeText() string`

GetStatusCodeText returns the StatusCodeText field if non-nil, zero value otherwise.

### GetStatusCodeTextOk

`func (o *InlineResponseDefault) GetStatusCodeTextOk() (*string, bool)`

GetStatusCodeTextOk returns a tuple with the StatusCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeText

`func (o *InlineResponseDefault) SetStatusCodeText(v string)`

SetStatusCodeText sets StatusCodeText field to given value.


### GetErrorDateTime

`func (o *InlineResponseDefault) GetErrorDateTime() string`

GetErrorDateTime returns the ErrorDateTime field if non-nil, zero value otherwise.

### GetErrorDateTimeOk

`func (o *InlineResponseDefault) GetErrorDateTimeOk() (*string, bool)`

GetErrorDateTimeOk returns a tuple with the ErrorDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDateTime

`func (o *InlineResponseDefault) SetErrorDateTime(v string)`

SetErrorDateTime sets ErrorDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



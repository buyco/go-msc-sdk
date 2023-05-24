# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethod** | **string** | The HTTP request method type | 
**RequestUri** | **string** | The request URI. | 
**Errors** | [**[]SubErrorsInner**](SubErrorsInner.md) |  | 
**StatusCode** | **int32** | The HTTP status code | 
**StatusCodeText** | **string** | The textual representation of the response status. | 
**ErrorDateTime** | **string** | The date and time (in ISO 8601 format) the error occurred. | 

## Methods

### NewModelError

`func NewModelError(httpMethod string, requestUri string, errors []SubErrorsInner, statusCode int32, statusCodeText string, errorDateTime string, ) *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpMethod

`func (o *ModelError) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ModelError) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ModelError) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### GetRequestUri

`func (o *ModelError) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *ModelError) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *ModelError) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.


### GetErrors

`func (o *ModelError) GetErrors() []SubErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelError) GetErrorsOk() (*[]SubErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelError) SetErrors(v []SubErrorsInner)`

SetErrors sets Errors field to given value.


### GetStatusCode

`func (o *ModelError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ModelError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ModelError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusCodeText

`func (o *ModelError) GetStatusCodeText() string`

GetStatusCodeText returns the StatusCodeText field if non-nil, zero value otherwise.

### GetStatusCodeTextOk

`func (o *ModelError) GetStatusCodeTextOk() (*string, bool)`

GetStatusCodeTextOk returns a tuple with the StatusCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeText

`func (o *ModelError) SetStatusCodeText(v string)`

SetStatusCodeText sets StatusCodeText field to given value.


### GetErrorDateTime

`func (o *ModelError) GetErrorDateTime() string`

GetErrorDateTime returns the ErrorDateTime field if non-nil, zero value otherwise.

### GetErrorDateTimeOk

`func (o *ModelError) GetErrorDateTimeOk() (*string, bool)`

GetErrorDateTimeOk returns a tuple with the ErrorDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDateTime

`func (o *ModelError) SetErrorDateTime(v string)`

SetErrorDateTime sets ErrorDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



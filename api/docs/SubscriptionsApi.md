# \SubscriptionsApi

All URIs are relative to *https://virtserver.swaggerhub.com/dcsaorg/DCSA_OAS/1.2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventSubscriptionsGet**](SubscriptionsApi.md#EventSubscriptionsGet) | **Get** /event-subscriptions | Receive a list of your active subscriptionIDs
[**EventSubscriptionsPost**](SubscriptionsApi.md#EventSubscriptionsPost) | **Post** /event-subscriptions | Create an event subscription
[**EventSubscriptionsSubscriptionIDDelete**](SubscriptionsApi.md#EventSubscriptionsSubscriptionIDDelete) | **Delete** /event-subscriptions/{subscriptionID} | Stop an event subscription, using the subscription ID
[**EventSubscriptionsSubscriptionIDGet**](SubscriptionsApi.md#EventSubscriptionsSubscriptionIDGet) | **Get** /event-subscriptions/{subscriptionID} | Find event subscription by subscription ID
[**EventSubscriptionsSubscriptionIDPut**](SubscriptionsApi.md#EventSubscriptionsSubscriptionIDPut) | **Put** /event-subscriptions/{subscriptionID} | Alter an event subscription



## EventSubscriptionsGet

> []EventSubscription EventSubscriptionsGet(ctx).Execute()

Receive a list of your active subscriptionIDs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.EventSubscriptionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.EventSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventSubscriptionsGet`: []EventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.EventSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventSubscriptionsGetRequest struct via the builder pattern


### Return type

[**[]EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventSubscriptionsPost

> EventSubscription EventSubscriptionsPost(ctx).EventSubscriptionBody(eventSubscriptionBody).Execute()

Create an event subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSubscriptionBody := *openapiclient.NewEventSubscriptionBody("https://myserver.com/send/callback/here") // EventSubscriptionBody | Parameters used to configure the subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.EventSubscriptionsPost(context.Background()).EventSubscriptionBody(eventSubscriptionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.EventSubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventSubscriptionsPost`: EventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.EventSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSubscriptionBody** | [**EventSubscriptionBody**](EventSubscriptionBody.md) | Parameters used to configure the subscription | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventSubscriptionsSubscriptionIDDelete

> EventSubscriptionsSubscriptionIDDelete(ctx, subscriptionID).Execute()

Stop an event subscription, using the subscription ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.EventSubscriptionsSubscriptionIDDelete(context.Background(), subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.EventSubscriptionsSubscriptionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventSubscriptionsSubscriptionIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventSubscriptionsSubscriptionIDGet

> EventSubscription EventSubscriptionsSubscriptionIDGet(ctx, subscriptionID).Execute()

Find event subscription by subscription ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionID := TODO // string | The universal unique ID of the subscription to receive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.EventSubscriptionsSubscriptionIDGet(context.Background(), subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.EventSubscriptionsSubscriptionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventSubscriptionsSubscriptionIDGet`: EventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.EventSubscriptionsSubscriptionIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | [**string**](.md) | The universal unique ID of the subscription to receive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventSubscriptionsSubscriptionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventSubscriptionsSubscriptionIDPut

> EventSubscription EventSubscriptionsSubscriptionIDPut(ctx, subscriptionID).EventSubscription(eventSubscription).Execute()

Alter an event subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionID := TODO // string | 
    eventSubscription := *openapiclient.NewEventSubscription("https://myserver.com/send/callback/here") // EventSubscription | Parameters used to configure the subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.EventSubscriptionsSubscriptionIDPut(context.Background(), subscriptionID).EventSubscription(eventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.EventSubscriptionsSubscriptionIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventSubscriptionsSubscriptionIDPut`: EventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.EventSubscriptionsSubscriptionIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventSubscriptionsSubscriptionIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventSubscription** | [**EventSubscription**](EventSubscription.md) | Parameters used to configure the subscription | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


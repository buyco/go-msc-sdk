# \EventsApi

All URIs are relative to *https://virtserver.swaggerhub.com/dcsaorg/DCSA_OAS/1.2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsEventIDGet**](EventsApi.md#EventsEventIDGet) | **Get** /events/{eventID} | Find events by eventID.
[**EventsGet**](EventsApi.md#EventsGet) | **Get** /events | Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



## EventsEventIDGet

> InlineResponse200 EventsEventIDGet(ctx, eventID).Execute()

Find events by eventID.



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
    eventID := "eventID_example" // string | The ID of the event to receive

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.EventsEventIDGet(context.Background(), eventID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventIDGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventID** | **string** | The ID of the event to receive | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsGet

> Events EventsGet(ctx).EventType(eventType).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()

Find events by type, Booking Reference, Bill of Lading or Equipment Reference.



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
    eventType := []openapiclient.EventType{openapiclient.eventType("EQUIPMENT")} // []EventType | The type of event(s) to filter by. (optional) (default to ["EQUIPMENT","SHIPMENT","TRANSPORT","TRANSPORTEQUIPMENT"])
    bookingReference := "bookingReference_example" // string | The identifier for a shipment, which is issued by and unique within each of the carriers. (optional)
    billOfLadingNumber := "billOfLadingNumber_example" // string | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier's receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. (optional)
    equipmentReference := "equipmentReference_example" // string | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.EventsGet(context.Background()).EventType(eventType).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsGet`: Events
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventType** | [**[]EventType**](EventType.md) | The type of event(s) to filter by. | [default to [&quot;EQUIPMENT&quot;,&quot;SHIPMENT&quot;,&quot;TRANSPORT&quot;,&quot;TRANSPORTEQUIPMENT&quot;]]
 **bookingReference** | **string** | The identifier for a shipment, which is issued by and unique within each of the carriers. | 
 **billOfLadingNumber** | **string** | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier&#39;s receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. | 
 **equipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. | 

### Return type

[**Events**](Events.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \EventsApi

All URIs are relative to *https://tst.portal.tech.msc.com/trackandtrace*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /v2.2/events | Find events.



## GetEvents

> []EventsInner GetEvents(ctx).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).Execute()

Find events.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/buyco/go-msc-sdk"
)

func main() {
    carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
    transportDocumentReference := "transportDocumentReference_example" // string | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.  Specifying this filter will only return events related to this particular transportDocumentReference  (optional)
    equipmentReference := "equipmentReference_example" // string | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible.  Specifying this filter will only return events related to this particular equipmentReference  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvents(context.Background()).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: []EventsInner
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **transportDocumentReference** | **string** | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.  Specifying this filter will only return events related to this particular transportDocumentReference  | 
 **equipmentReference** | **string** | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible.  Specifying this filter will only return events related to this particular equipmentReference  | 

### Return type

[**[]EventsInner**](EventsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


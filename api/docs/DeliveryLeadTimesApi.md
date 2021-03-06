# \DeliveryLeadTimesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#DELETEDeliveryLeadTimesDeliveryLeadTimeId) | **Delete** /delivery_lead_times/{deliveryLeadTimeId} | Delete a delivery lead time
[**GETDeliveryLeadTimes**](DeliveryLeadTimesApi.md#GETDeliveryLeadTimes) | **Get** /delivery_lead_times | List all delivery lead times
[**GETDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#GETDeliveryLeadTimesDeliveryLeadTimeId) | **Get** /delivery_lead_times/{deliveryLeadTimeId} | Retrieve a delivery lead time
[**GETShipmentIdDeliveryLeadTime**](DeliveryLeadTimesApi.md#GETShipmentIdDeliveryLeadTime) | **Get** /shipments/{shipmentId}/delivery_lead_time | Retrieve the delivery lead time associated to the shipment
[**GETShippingMethodIdDeliveryLeadTimeForShipment**](DeliveryLeadTimesApi.md#GETShippingMethodIdDeliveryLeadTimeForShipment) | **Get** /shipping_methods/{shippingMethodId}/delivery_lead_time_for_shipment | Retrieve the delivery lead time for shipment associated to the shipping method
[**GETSkuIdDeliveryLeadTimes**](DeliveryLeadTimesApi.md#GETSkuIdDeliveryLeadTimes) | **Get** /skus/{skuId}/delivery_lead_times | Retrieve the delivery lead times associated to the SKU
[**PATCHDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#PATCHDeliveryLeadTimesDeliveryLeadTimeId) | **Patch** /delivery_lead_times/{deliveryLeadTimeId} | Update a delivery lead time
[**POSTDeliveryLeadTimes**](DeliveryLeadTimesApi.md#POSTDeliveryLeadTimes) | **Post** /delivery_lead_times | Create a delivery lead time



## DELETEDeliveryLeadTimesDeliveryLeadTimeId

> DELETEDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).Execute()

Delete a delivery lead time



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
    deliveryLeadTimeId := "deliveryLeadTimeId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.DELETEDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.DELETEDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETDeliveryLeadTimes

> GETDeliveryLeadTimes(ctx).Execute()

List all delivery lead times



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETDeliveryLeadTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETDeliveryLeadTimesDeliveryLeadTimeId

> DeliveryLeadTime GETDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).Execute()

Retrieve a delivery lead time



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
    deliveryLeadTimeId := "deliveryLeadTimeId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETDeliveryLeadTimesDeliveryLeadTimeId`: DeliveryLeadTime
    fmt.Fprintf(os.Stdout, "Response from `DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeliveryLeadTime**](DeliveryLeadTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdDeliveryLeadTime

> GETShipmentIdDeliveryLeadTime(ctx, shipmentId).Execute()

Retrieve the delivery lead time associated to the shipment



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
    shipmentId := "shipmentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETShipmentIdDeliveryLeadTime(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETShipmentIdDeliveryLeadTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentIdDeliveryLeadTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingMethodIdDeliveryLeadTimeForShipment

> GETShippingMethodIdDeliveryLeadTimeForShipment(ctx, shippingMethodId).Execute()

Retrieve the delivery lead time for shipment associated to the shipping method



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
    shippingMethodId := "shippingMethodId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETShippingMethodIdDeliveryLeadTimeForShipment(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETShippingMethodIdDeliveryLeadTimeForShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuIdDeliveryLeadTimes

> GETSkuIdDeliveryLeadTimes(ctx, skuId).Execute()

Retrieve the delivery lead times associated to the SKU



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
    skuId := "skuId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETSkuIdDeliveryLeadTimes(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETSkuIdDeliveryLeadTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuIdDeliveryLeadTimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHDeliveryLeadTimesDeliveryLeadTimeId

> PATCHDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).DeliveryLeadTimeUpdate(deliveryLeadTimeUpdate).Execute()

Update a delivery lead time



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
    deliveryLeadTimeUpdate := *openapiclient.NewDeliveryLeadTimeUpdate(*openapiclient.NewDeliveryLeadTimeUpdateData("delivery_lead_times", "XGZwpOSrWL", *openapiclient.NewDeliveryLeadTimeUpdateDataAttributes())) // DeliveryLeadTimeUpdate | 
    deliveryLeadTimeId := "deliveryLeadTimeId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).DeliveryLeadTimeUpdate(deliveryLeadTimeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryLeadTimeUpdate** | [**DeliveryLeadTimeUpdate**](DeliveryLeadTimeUpdate.md) |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTDeliveryLeadTimes

> POSTDeliveryLeadTimes(ctx).DeliveryLeadTimeCreate(deliveryLeadTimeCreate).Execute()

Create a delivery lead time



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
    deliveryLeadTimeCreate := *openapiclient.NewDeliveryLeadTimeCreate(*openapiclient.NewDeliveryLeadTimeCreateData("delivery_lead_times", *openapiclient.NewDeliveryLeadTimeCreateDataAttributes(int32(48), int32(72)))) // DeliveryLeadTimeCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.POSTDeliveryLeadTimes(context.Background()).DeliveryLeadTimeCreate(deliveryLeadTimeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.POSTDeliveryLeadTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTDeliveryLeadTimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryLeadTimeCreate** | [**DeliveryLeadTimeCreate**](DeliveryLeadTimeCreate.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


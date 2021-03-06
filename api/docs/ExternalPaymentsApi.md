# \ExternalPaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#DELETEExternalPaymentsExternalPaymentId) | **Delete** /external_payments/{externalPaymentId} | Delete an external payment
[**GETExternalGatewayIdExternalPayments**](ExternalPaymentsApi.md#GETExternalGatewayIdExternalPayments) | **Get** /external_gateways/{externalGatewayId}/external_payments | Retrieve the external payments associated to the external gateway
[**GETExternalPayments**](ExternalPaymentsApi.md#GETExternalPayments) | **Get** /external_payments | List all external payments
[**GETExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#GETExternalPaymentsExternalPaymentId) | **Get** /external_payments/{externalPaymentId} | Retrieve an external payment
[**PATCHExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#PATCHExternalPaymentsExternalPaymentId) | **Patch** /external_payments/{externalPaymentId} | Update an external payment
[**POSTExternalPayments**](ExternalPaymentsApi.md#POSTExternalPayments) | **Post** /external_payments | Create an external payment



## DELETEExternalPaymentsExternalPaymentId

> DELETEExternalPaymentsExternalPaymentId(ctx, externalPaymentId).Execute()

Delete an external payment



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
    externalPaymentId := "externalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.DELETEExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.DELETEExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


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


## GETExternalGatewayIdExternalPayments

> GETExternalGatewayIdExternalPayments(ctx, externalGatewayId).Execute()

Retrieve the external payments associated to the external gateway



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
    externalGatewayId := "externalGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.GETExternalGatewayIdExternalPayments(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalGatewayIdExternalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewayIdExternalPaymentsRequest struct via the builder pattern


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


## GETExternalPayments

> GETExternalPayments(ctx).Execute()

List all external payments



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
    resp, r, err := apiClient.ExternalPaymentsApi.GETExternalPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPaymentsRequest struct via the builder pattern


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


## GETExternalPaymentsExternalPaymentId

> ExternalPayment GETExternalPaymentsExternalPaymentId(ctx, externalPaymentId).Execute()

Retrieve an external payment



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
    externalPaymentId := "externalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPaymentsExternalPaymentId`: ExternalPayment
    fmt.Fprintf(os.Stdout, "Response from `ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalPayment**](ExternalPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalPaymentsExternalPaymentId

> PATCHExternalPaymentsExternalPaymentId(ctx, externalPaymentId).ExternalPaymentUpdate(externalPaymentUpdate).Execute()

Update an external payment



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
    externalPaymentUpdate := *openapiclient.NewExternalPaymentUpdate(*openapiclient.NewExternalPaymentUpdateData("external_payments", "XGZwpOSrWL", *openapiclient.NewExternalPaymentUpdateDataAttributes())) // ExternalPaymentUpdate | 
    externalPaymentId := "externalPaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).ExternalPaymentUpdate(externalPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPaymentUpdate** | [**ExternalPaymentUpdate**](ExternalPaymentUpdate.md) |  | 


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


## POSTExternalPayments

> POSTExternalPayments(ctx).ExternalPaymentCreate(externalPaymentCreate).Execute()

Create an external payment



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
    externalPaymentCreate := *openapiclient.NewExternalPaymentCreate(*openapiclient.NewExternalPaymentCreateData("external_payments", *openapiclient.NewExternalPaymentCreateDataAttributes("xxxx.yyyy.zzzz"))) // ExternalPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.POSTExternalPayments(context.Background()).ExternalPaymentCreate(externalPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.POSTExternalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPaymentCreate** | [**ExternalPaymentCreate**](ExternalPaymentCreate.md) |  | 

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


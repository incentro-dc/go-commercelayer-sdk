# \BraintreePaymentsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#DELETEBraintreePaymentsBraintreePaymentId) | **Delete** /braintree_payments/{braintreePaymentId} | Delete a braintree payment
[**GETBraintreeGatewayIdBraintreePayments**](BraintreePaymentsApi.md#GETBraintreeGatewayIdBraintreePayments) | **Get** /braintree_gateways/{braintreeGatewayId}/braintree_payments | Retrieve the braintree payments associated to the braintree gateway
[**GETBraintreePayments**](BraintreePaymentsApi.md#GETBraintreePayments) | **Get** /braintree_payments | List all braintree payments
[**GETBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#GETBraintreePaymentsBraintreePaymentId) | **Get** /braintree_payments/{braintreePaymentId} | Retrieve a braintree payment
[**PATCHBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#PATCHBraintreePaymentsBraintreePaymentId) | **Patch** /braintree_payments/{braintreePaymentId} | Update a braintree payment
[**POSTBraintreePayments**](BraintreePaymentsApi.md#POSTBraintreePayments) | **Post** /braintree_payments | Create a braintree payment



## DELETEBraintreePaymentsBraintreePaymentId

> DELETEBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).Execute()

Delete a braintree payment



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
    braintreePaymentId := "braintreePaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.DELETEBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.DELETEBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


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


## GETBraintreeGatewayIdBraintreePayments

> GETBraintreeGatewayIdBraintreePayments(ctx, braintreeGatewayId).Execute()

Retrieve the braintree payments associated to the braintree gateway



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
    braintreeGatewayId := "braintreeGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.GETBraintreeGatewayIdBraintreePayments(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreeGatewayIdBraintreePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewayIdBraintreePaymentsRequest struct via the builder pattern


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


## GETBraintreePayments

> GETBraintreePayments(ctx).Execute()

List all braintree payments



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
    resp, r, err := apiClient.BraintreePaymentsApi.GETBraintreePayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreePaymentsRequest struct via the builder pattern


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


## GETBraintreePaymentsBraintreePaymentId

> BraintreePayment GETBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).Execute()

Retrieve a braintree payment



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
    braintreePaymentId := "braintreePaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreePaymentsBraintreePaymentId`: BraintreePayment
    fmt.Fprintf(os.Stdout, "Response from `BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BraintreePayment**](BraintreePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBraintreePaymentsBraintreePaymentId

> PATCHBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).BraintreePaymentUpdate(braintreePaymentUpdate).Execute()

Update a braintree payment



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
    braintreePaymentUpdate := *openapiclient.NewBraintreePaymentUpdate(*openapiclient.NewBraintreePaymentUpdateData("braintree_payments", "XGZwpOSrWL", *openapiclient.NewBraintreePaymentUpdateDataAttributes())) // BraintreePaymentUpdate | 
    braintreePaymentId := "braintreePaymentId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).BraintreePaymentUpdate(braintreePaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreePaymentUpdate** | [**BraintreePaymentUpdate**](BraintreePaymentUpdate.md) |  | 


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


## POSTBraintreePayments

> POSTBraintreePayments(ctx).BraintreePaymentCreate(braintreePaymentCreate).Execute()

Create a braintree payment



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
    braintreePaymentCreate := *openapiclient.NewBraintreePaymentCreate(*openapiclient.NewBraintreePaymentCreateData("braintree_payments", *openapiclient.NewBraintreePaymentCreateDataAttributes())) // BraintreePaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.POSTBraintreePayments(context.Background()).BraintreePaymentCreate(braintreePaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.POSTBraintreePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBraintreePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreePaymentCreate** | [**BraintreePaymentCreate**](BraintreePaymentCreate.md) |  | 

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


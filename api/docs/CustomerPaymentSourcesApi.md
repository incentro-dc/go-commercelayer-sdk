# \CustomerPaymentSourcesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#DELETECustomerPaymentSourcesCustomerPaymentSourceId) | **Delete** /customer_payment_sources/{customerPaymentSourceId} | Delete a customer payment source
[**GETCustomerIdCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETCustomerIdCustomerPaymentSources) | **Get** /customers/{customerId}/customer_payment_sources | Retrieve the customer payment sources associated to the customer
[**GETCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETCustomerPaymentSources) | **Get** /customer_payment_sources | List all customer payment sources
[**GETCustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#GETCustomerPaymentSourcesCustomerPaymentSourceId) | **Get** /customer_payment_sources/{customerPaymentSourceId} | Retrieve a customer payment source
[**GETExternalPaymentIdWallet**](CustomerPaymentSourcesApi.md#GETExternalPaymentIdWallet) | **Get** /external_payments/{externalPaymentId}/wallet | Retrieve the wallet associated to the external payment
[**GETOrderIdAvailableCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETOrderIdAvailableCustomerPaymentSources) | **Get** /orders/{orderId}/available_customer_payment_sources | Retrieve the available customer payment sources associated to the order
[**PATCHCustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#PATCHCustomerPaymentSourcesCustomerPaymentSourceId) | **Patch** /customer_payment_sources/{customerPaymentSourceId} | Update a customer payment source
[**POSTCustomerPaymentSources**](CustomerPaymentSourcesApi.md#POSTCustomerPaymentSources) | **Post** /customer_payment_sources | Create a customer payment source



## DELETECustomerPaymentSourcesCustomerPaymentSourceId

> DELETECustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).Execute()

Delete a customer payment source



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
    customerPaymentSourceId := "customerPaymentSourceId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.DELETECustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.DELETECustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


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


## GETCustomerIdCustomerPaymentSources

> GETCustomerIdCustomerPaymentSources(ctx, customerId).Execute()

Retrieve the customer payment sources associated to the customer



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
    customerId := "customerId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerIdCustomerPaymentSources(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerIdCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerPaymentSourcesRequest struct via the builder pattern


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


## GETCustomerPaymentSources

> GETCustomerPaymentSources(ctx).Execute()

List all customer payment sources



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
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerPaymentSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourcesRequest struct via the builder pattern


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


## GETCustomerPaymentSourcesCustomerPaymentSourceId

> CustomerPaymentSource GETCustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).Execute()

Retrieve a customer payment source



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
    customerPaymentSourceId := "customerPaymentSourceId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPaymentSourcesCustomerPaymentSourceId`: CustomerPaymentSource
    fmt.Fprintf(os.Stdout, "Response from `CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerPaymentSource**](CustomerPaymentSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPaymentIdWallet

> GETExternalPaymentIdWallet(ctx, externalPaymentId).Execute()

Retrieve the wallet associated to the external payment



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
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETExternalPaymentIdWallet(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETExternalPaymentIdWallet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPaymentIdWalletRequest struct via the builder pattern


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


## GETOrderIdAvailableCustomerPaymentSources

> GETOrderIdAvailableCustomerPaymentSources(ctx, orderId).Execute()

Retrieve the available customer payment sources associated to the order



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
    orderId := "orderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETOrderIdAvailableCustomerPaymentSources(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETOrderIdAvailableCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderIdAvailableCustomerPaymentSourcesRequest struct via the builder pattern


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


## PATCHCustomerPaymentSourcesCustomerPaymentSourceId

> PATCHCustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).CustomerPaymentSourceUpdate(customerPaymentSourceUpdate).Execute()

Update a customer payment source



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
    customerPaymentSourceUpdate := *openapiclient.NewCustomerPaymentSourceUpdate(*openapiclient.NewCustomerPaymentSourceUpdateData("customer_payment_sources", "XGZwpOSrWL", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CustomerPaymentSourceUpdate | 
    customerPaymentSourceId := "customerPaymentSourceId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.PATCHCustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).CustomerPaymentSourceUpdate(customerPaymentSourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.PATCHCustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPaymentSourceUpdate** | [**CustomerPaymentSourceUpdate**](CustomerPaymentSourceUpdate.md) |  | 


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


## POSTCustomerPaymentSources

> POSTCustomerPaymentSources(ctx).CustomerPaymentSourceCreate(customerPaymentSourceCreate).Execute()

Create a customer payment source



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
    customerPaymentSourceCreate := *openapiclient.NewCustomerPaymentSourceCreate(*openapiclient.NewCustomerPaymentSourceCreateData("customer_payment_sources", *openapiclient.NewAdyenPaymentCreateDataAttributes())) // CustomerPaymentSourceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.POSTCustomerPaymentSources(context.Background()).CustomerPaymentSourceCreate(customerPaymentSourceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.POSTCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerPaymentSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPaymentSourceCreate** | [**CustomerPaymentSourceCreate**](CustomerPaymentSourceCreate.md) |  | 

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


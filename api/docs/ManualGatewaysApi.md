# \ManualGatewaysApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEManualGatewaysManualGatewayId**](ManualGatewaysApi.md#DELETEManualGatewaysManualGatewayId) | **Delete** /manual_gateways/{manualGatewayId} | Delete a manual gateway
[**GETManualGateways**](ManualGatewaysApi.md#GETManualGateways) | **Get** /manual_gateways | List all manual gateways
[**GETManualGatewaysManualGatewayId**](ManualGatewaysApi.md#GETManualGatewaysManualGatewayId) | **Get** /manual_gateways/{manualGatewayId} | Retrieve a manual gateway
[**PATCHManualGatewaysManualGatewayId**](ManualGatewaysApi.md#PATCHManualGatewaysManualGatewayId) | **Patch** /manual_gateways/{manualGatewayId} | Update a manual gateway
[**POSTManualGateways**](ManualGatewaysApi.md#POSTManualGateways) | **Post** /manual_gateways | Create a manual gateway



## DELETEManualGatewaysManualGatewayId

> DELETEManualGatewaysManualGatewayId(ctx, manualGatewayId).Execute()

Delete a manual gateway



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
    manualGatewayId := "manualGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualGatewaysApi.DELETEManualGatewaysManualGatewayId(context.Background(), manualGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualGatewaysApi.DELETEManualGatewaysManualGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEManualGatewaysManualGatewayIdRequest struct via the builder pattern


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


## GETManualGateways

> GETManualGateways(ctx).Execute()

List all manual gateways



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
    resp, r, err := apiClient.ManualGatewaysApi.GETManualGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualGatewaysApi.GETManualGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualGatewaysRequest struct via the builder pattern


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


## GETManualGatewaysManualGatewayId

> ManualGateway GETManualGatewaysManualGatewayId(ctx, manualGatewayId).Execute()

Retrieve a manual gateway



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
    manualGatewayId := "manualGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualGatewaysApi.GETManualGatewaysManualGatewayId(context.Background(), manualGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualGatewaysApi.GETManualGatewaysManualGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETManualGatewaysManualGatewayId`: ManualGateway
    fmt.Fprintf(os.Stdout, "Response from `ManualGatewaysApi.GETManualGatewaysManualGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualGatewaysManualGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManualGateway**](ManualGateway.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHManualGatewaysManualGatewayId

> PATCHManualGatewaysManualGatewayId(ctx, manualGatewayId).ManualGatewayUpdate(manualGatewayUpdate).Execute()

Update a manual gateway



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
    manualGatewayUpdate := *openapiclient.NewManualGatewayUpdate(*openapiclient.NewManualGatewayUpdateData("manual_gateways", "XGZwpOSrWL", *openapiclient.NewManualGatewayUpdateDataAttributes())) // ManualGatewayUpdate | 
    manualGatewayId := "manualGatewayId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualGatewaysApi.PATCHManualGatewaysManualGatewayId(context.Background(), manualGatewayId).ManualGatewayUpdate(manualGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualGatewaysApi.PATCHManualGatewaysManualGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualGatewayId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHManualGatewaysManualGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualGatewayUpdate** | [**ManualGatewayUpdate**](ManualGatewayUpdate.md) |  | 


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


## POSTManualGateways

> POSTManualGateways(ctx).ManualGatewayCreate(manualGatewayCreate).Execute()

Create a manual gateway



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
    manualGatewayCreate := *openapiclient.NewManualGatewayCreate(*openapiclient.NewManualGatewayCreateData("manual_gateways", *openapiclient.NewManualGatewayCreateDataAttributes("US payment gateway"))) // ManualGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualGatewaysApi.POSTManualGateways(context.Background()).ManualGatewayCreate(manualGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualGatewaysApi.POSTManualGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTManualGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualGatewayCreate** | [**ManualGatewayCreate**](ManualGatewayCreate.md) |  | 

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


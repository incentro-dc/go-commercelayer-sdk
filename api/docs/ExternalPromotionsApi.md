# \ExternalPromotionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#DELETEExternalPromotionsExternalPromotionId) | **Delete** /external_promotions/{externalPromotionId} | Delete an external promotion
[**GETExternalPromotions**](ExternalPromotionsApi.md#GETExternalPromotions) | **Get** /external_promotions | List all external promotions
[**GETExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#GETExternalPromotionsExternalPromotionId) | **Get** /external_promotions/{externalPromotionId} | Retrieve an external promotion
[**PATCHExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#PATCHExternalPromotionsExternalPromotionId) | **Patch** /external_promotions/{externalPromotionId} | Update an external promotion
[**POSTExternalPromotions**](ExternalPromotionsApi.md#POSTExternalPromotions) | **Post** /external_promotions | Create an external promotion



## DELETEExternalPromotionsExternalPromotionId

> DELETEExternalPromotionsExternalPromotionId(ctx, externalPromotionId).Execute()

Delete an external promotion



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
    externalPromotionId := "externalPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.DELETEExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.DELETEExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


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


## GETExternalPromotions

> GETExternalPromotions(ctx).Execute()

List all external promotions



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
    resp, r, err := apiClient.ExternalPromotionsApi.GETExternalPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.GETExternalPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionsRequest struct via the builder pattern


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


## GETExternalPromotionsExternalPromotionId

> ExternalPromotion GETExternalPromotionsExternalPromotionId(ctx, externalPromotionId).Execute()

Retrieve an external promotion



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
    externalPromotionId := "externalPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPromotionsExternalPromotionId`: ExternalPromotion
    fmt.Fprintf(os.Stdout, "Response from `ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalPromotion**](ExternalPromotion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalPromotionsExternalPromotionId

> PATCHExternalPromotionsExternalPromotionId(ctx, externalPromotionId).ExternalPromotionUpdate(externalPromotionUpdate).Execute()

Update an external promotion



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
    externalPromotionUpdate := *openapiclient.NewExternalPromotionUpdate(*openapiclient.NewExternalPromotionUpdateData("external_promotions", "XGZwpOSrWL", *openapiclient.NewExternalPromotionUpdateDataAttributes())) // ExternalPromotionUpdate | 
    externalPromotionId := "externalPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.PATCHExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).ExternalPromotionUpdate(externalPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.PATCHExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPromotionUpdate** | [**ExternalPromotionUpdate**](ExternalPromotionUpdate.md) |  | 


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


## POSTExternalPromotions

> POSTExternalPromotions(ctx).ExternalPromotionCreate(externalPromotionCreate).Execute()

Create an external promotion



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
    externalPromotionCreate := *openapiclient.NewExternalPromotionCreate(*openapiclient.NewExternalPromotionCreateData("external_promotions", *openapiclient.NewExternalPromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5), "https://external_promotion.yourbrand.com"))) // ExternalPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.POSTExternalPromotions(context.Background()).ExternalPromotionCreate(externalPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.POSTExternalPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPromotionCreate** | [**ExternalPromotionCreate**](ExternalPromotionCreate.md) |  | 

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


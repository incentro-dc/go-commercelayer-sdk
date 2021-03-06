# \FixedPricePromotionsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#DELETEFixedPricePromotionsFixedPricePromotionId) | **Delete** /fixed_price_promotions/{fixedPricePromotionId} | Delete a fixed price promotion
[**GETFixedPricePromotions**](FixedPricePromotionsApi.md#GETFixedPricePromotions) | **Get** /fixed_price_promotions | List all fixed price promotions
[**GETFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#GETFixedPricePromotionsFixedPricePromotionId) | **Get** /fixed_price_promotions/{fixedPricePromotionId} | Retrieve a fixed price promotion
[**PATCHFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#PATCHFixedPricePromotionsFixedPricePromotionId) | **Patch** /fixed_price_promotions/{fixedPricePromotionId} | Update a fixed price promotion
[**POSTFixedPricePromotions**](FixedPricePromotionsApi.md#POSTFixedPricePromotions) | **Post** /fixed_price_promotions | Create a fixed price promotion



## DELETEFixedPricePromotionsFixedPricePromotionId

> DELETEFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).Execute()

Delete a fixed price promotion



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
    fixedPricePromotionId := "fixedPricePromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.DELETEFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.DELETEFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


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


## GETFixedPricePromotions

> GETFixedPricePromotions(ctx).Execute()

List all fixed price promotions



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
    resp, r, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.GETFixedPricePromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionsRequest struct via the builder pattern


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


## GETFixedPricePromotionsFixedPricePromotionId

> FixedPricePromotion GETFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).Execute()

Retrieve a fixed price promotion



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
    fixedPricePromotionId := "fixedPricePromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedPricePromotionsFixedPricePromotionId`: FixedPricePromotion
    fmt.Fprintf(os.Stdout, "Response from `FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FixedPricePromotion**](FixedPricePromotion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFixedPricePromotionsFixedPricePromotionId

> PATCHFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).FixedPricePromotionUpdate(fixedPricePromotionUpdate).Execute()

Update a fixed price promotion



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
    fixedPricePromotionUpdate := *openapiclient.NewFixedPricePromotionUpdate(*openapiclient.NewFixedPricePromotionUpdateData("fixed_price_promotions", "XGZwpOSrWL", *openapiclient.NewFixedPricePromotionUpdateDataAttributes())) // FixedPricePromotionUpdate | 
    fixedPricePromotionId := "fixedPricePromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).FixedPricePromotionUpdate(fixedPricePromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedPricePromotionUpdate** | [**FixedPricePromotionUpdate**](FixedPricePromotionUpdate.md) |  | 


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


## POSTFixedPricePromotions

> POSTFixedPricePromotions(ctx).FixedPricePromotionCreate(fixedPricePromotionCreate).Execute()

Create a fixed price promotion



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
    fixedPricePromotionCreate := *openapiclient.NewFixedPricePromotionCreate(*openapiclient.NewFixedPricePromotionCreateData("fixed_price_promotions", *openapiclient.NewFixedPricePromotionCreateDataAttributes("Personal promotion", "2018-01-01T12:00:00.000Z", "2018-01-02T12:00:00.000Z", int32(5), int32(1000)))) // FixedPricePromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.POSTFixedPricePromotions(context.Background()).FixedPricePromotionCreate(fixedPricePromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.POSTFixedPricePromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFixedPricePromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedPricePromotionCreate** | [**FixedPricePromotionCreate**](FixedPricePromotionCreate.md) |  | 

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


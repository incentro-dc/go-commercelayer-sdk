# \BingGeocodersApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBingGeocodersBingGeocoderId**](BingGeocodersApi.md#DELETEBingGeocodersBingGeocoderId) | **Delete** /bing_geocoders/{bingGeocoderId} | Delete a bing geocoder
[**GETBingGeocoders**](BingGeocodersApi.md#GETBingGeocoders) | **Get** /bing_geocoders | List all bing geocoders
[**GETBingGeocodersBingGeocoderId**](BingGeocodersApi.md#GETBingGeocodersBingGeocoderId) | **Get** /bing_geocoders/{bingGeocoderId} | Retrieve a bing geocoder
[**PATCHBingGeocodersBingGeocoderId**](BingGeocodersApi.md#PATCHBingGeocodersBingGeocoderId) | **Patch** /bing_geocoders/{bingGeocoderId} | Update a bing geocoder
[**POSTBingGeocoders**](BingGeocodersApi.md#POSTBingGeocoders) | **Post** /bing_geocoders | Create a bing geocoder



## DELETEBingGeocodersBingGeocoderId

> DELETEBingGeocodersBingGeocoderId(ctx, bingGeocoderId).Execute()

Delete a bing geocoder



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
    bingGeocoderId := "bingGeocoderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.DELETEBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.DELETEBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBingGeocodersBingGeocoderIdRequest struct via the builder pattern


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


## GETBingGeocoders

> GETBingGeocoders(ctx).Execute()

List all bing geocoders



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
    resp, r, err := apiClient.BingGeocodersApi.GETBingGeocoders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.GETBingGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBingGeocodersRequest struct via the builder pattern


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


## GETBingGeocodersBingGeocoderId

> BingGeocoder GETBingGeocodersBingGeocoderId(ctx, bingGeocoderId).Execute()

Retrieve a bing geocoder



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
    bingGeocoderId := "bingGeocoderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.GETBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.GETBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBingGeocodersBingGeocoderId`: BingGeocoder
    fmt.Fprintf(os.Stdout, "Response from `BingGeocodersApi.GETBingGeocodersBingGeocoderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBingGeocodersBingGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BingGeocoder**](BingGeocoder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBingGeocodersBingGeocoderId

> PATCHBingGeocodersBingGeocoderId(ctx, bingGeocoderId).BingGeocoderUpdate(bingGeocoderUpdate).Execute()

Update a bing geocoder



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
    bingGeocoderUpdate := *openapiclient.NewBingGeocoderUpdate(*openapiclient.NewBingGeocoderUpdateData("bing_geocoders", "XGZwpOSrWL", *openapiclient.NewBingGeocoderUpdateDataAttributes())) // BingGeocoderUpdate | 
    bingGeocoderId := "bingGeocoderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.PATCHBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).BingGeocoderUpdate(bingGeocoderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.PATCHBingGeocodersBingGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bingGeocoderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBingGeocodersBingGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bingGeocoderUpdate** | [**BingGeocoderUpdate**](BingGeocoderUpdate.md) |  | 


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


## POSTBingGeocoders

> POSTBingGeocoders(ctx).BingGeocoderCreate(bingGeocoderCreate).Execute()

Create a bing geocoder



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
    bingGeocoderCreate := *openapiclient.NewBingGeocoderCreate(*openapiclient.NewBingGeocoderCreateData("bing_geocoders", *openapiclient.NewBingGeocoderCreateDataAttributes("Default geocoder", "xxxx-yyyy-zzzz"))) // BingGeocoderCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BingGeocodersApi.POSTBingGeocoders(context.Background()).BingGeocoderCreate(bingGeocoderCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BingGeocodersApi.POSTBingGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBingGeocodersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bingGeocoderCreate** | [**BingGeocoderCreate**](BingGeocoderCreate.md) |  | 

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


# \ReturnLineItemsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#DELETEReturnLineItemsReturnLineItemId) | **Delete** /return_line_items/{returnLineItemId} | Delete a return line item
[**GETReturnIdReturnLineItems**](ReturnLineItemsApi.md#GETReturnIdReturnLineItems) | **Get** /returns/{returnId}/return_line_items | Retrieve the return line items associated to the return
[**GETReturnLineItems**](ReturnLineItemsApi.md#GETReturnLineItems) | **Get** /return_line_items | List all return line items
[**GETReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#GETReturnLineItemsReturnLineItemId) | **Get** /return_line_items/{returnLineItemId} | Retrieve a return line item
[**PATCHReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#PATCHReturnLineItemsReturnLineItemId) | **Patch** /return_line_items/{returnLineItemId} | Update a return line item
[**POSTReturnLineItems**](ReturnLineItemsApi.md#POSTReturnLineItems) | **Post** /return_line_items | Create a return line item



## DELETEReturnLineItemsReturnLineItemId

> DELETEReturnLineItemsReturnLineItemId(ctx, returnLineItemId).Execute()

Delete a return line item



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
    returnLineItemId := "returnLineItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.DELETEReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.DELETEReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


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


## GETReturnIdReturnLineItems

> GETReturnIdReturnLineItems(ctx, returnId).Execute()

Retrieve the return line items associated to the return



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
    returnId := "returnId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.GETReturnIdReturnLineItems(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnIdReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnIdReturnLineItemsRequest struct via the builder pattern


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


## GETReturnLineItems

> GETReturnLineItems(ctx).Execute()

List all return line items



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
    resp, r, err := apiClient.ReturnLineItemsApi.GETReturnLineItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnLineItemsRequest struct via the builder pattern


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


## GETReturnLineItemsReturnLineItemId

> ReturnLineItem GETReturnLineItemsReturnLineItemId(ctx, returnLineItemId).Execute()

Retrieve a return line item



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
    returnLineItemId := "returnLineItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETReturnLineItemsReturnLineItemId`: ReturnLineItem
    fmt.Fprintf(os.Stdout, "Response from `ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReturnLineItem**](ReturnLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHReturnLineItemsReturnLineItemId

> PATCHReturnLineItemsReturnLineItemId(ctx, returnLineItemId).ReturnLineItemUpdate(returnLineItemUpdate).Execute()

Update a return line item



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
    returnLineItemUpdate := *openapiclient.NewReturnLineItemUpdate(*openapiclient.NewReturnLineItemUpdateData("return_line_items", "XGZwpOSrWL", *openapiclient.NewReturnLineItemUpdateDataAttributes())) // ReturnLineItemUpdate | 
    returnLineItemId := "returnLineItemId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.PATCHReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).ReturnLineItemUpdate(returnLineItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.PATCHReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnLineItemUpdate** | [**ReturnLineItemUpdate**](ReturnLineItemUpdate.md) |  | 


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


## POSTReturnLineItems

> POSTReturnLineItems(ctx).ReturnLineItemCreate(returnLineItemCreate).Execute()

Create a return line item



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
    returnLineItemCreate := *openapiclient.NewReturnLineItemCreate(*openapiclient.NewReturnLineItemCreateData("return_line_items", *openapiclient.NewReturnLineItemCreateDataAttributes(int32(4)))) // ReturnLineItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.POSTReturnLineItems(context.Background()).ReturnLineItemCreate(returnLineItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.POSTReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTReturnLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnLineItemCreate** | [**ReturnLineItemCreate**](ReturnLineItemCreate.md) |  | 

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


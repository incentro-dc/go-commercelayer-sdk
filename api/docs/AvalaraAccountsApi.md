# \AvalaraAccountsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#DELETEAvalaraAccountsAvalaraAccountId) | **Delete** /avalara_accounts/{avalaraAccountId} | Delete an avalara account
[**GETAvalaraAccounts**](AvalaraAccountsApi.md#GETAvalaraAccounts) | **Get** /avalara_accounts | List all avalara accounts
[**GETAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#GETAvalaraAccountsAvalaraAccountId) | **Get** /avalara_accounts/{avalaraAccountId} | Retrieve an avalara account
[**PATCHAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#PATCHAvalaraAccountsAvalaraAccountId) | **Patch** /avalara_accounts/{avalaraAccountId} | Update an avalara account
[**POSTAvalaraAccounts**](AvalaraAccountsApi.md#POSTAvalaraAccounts) | **Post** /avalara_accounts | Create an avalara account



## DELETEAvalaraAccountsAvalaraAccountId

> DELETEAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).Execute()

Delete an avalara account



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
    avalaraAccountId := "avalaraAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.DELETEAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.DELETEAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


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


## GETAvalaraAccounts

> GETAvalaraAccounts(ctx).Execute()

List all avalara accounts



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
    resp, r, err := apiClient.AvalaraAccountsApi.GETAvalaraAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.GETAvalaraAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAvalaraAccountsRequest struct via the builder pattern


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


## GETAvalaraAccountsAvalaraAccountId

> AvalaraAccount GETAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).Execute()

Retrieve an avalara account



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
    avalaraAccountId := "avalaraAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAvalaraAccountsAvalaraAccountId`: AvalaraAccount
    fmt.Fprintf(os.Stdout, "Response from `AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvalaraAccount**](AvalaraAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAvalaraAccountsAvalaraAccountId

> PATCHAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).AvalaraAccountUpdate(avalaraAccountUpdate).Execute()

Update an avalara account



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
    avalaraAccountUpdate := *openapiclient.NewAvalaraAccountUpdate(*openapiclient.NewAvalaraAccountUpdateData("avalara_accounts", "XGZwpOSrWL", *openapiclient.NewAvalaraAccountUpdateDataAttributes())) // AvalaraAccountUpdate | 
    avalaraAccountId := "avalaraAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.PATCHAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).AvalaraAccountUpdate(avalaraAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.PATCHAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avalaraAccountUpdate** | [**AvalaraAccountUpdate**](AvalaraAccountUpdate.md) |  | 


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


## POSTAvalaraAccounts

> POSTAvalaraAccounts(ctx).AvalaraAccountCreate(avalaraAccountCreate).Execute()

Create an avalara account



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
    avalaraAccountCreate := *openapiclient.NewAvalaraAccountCreate(*openapiclient.NewAvalaraAccountCreateData("avalara_accounts", *openapiclient.NewAvalaraAccountCreateDataAttributes("Personal tax calculator", "user@mydomain.com", "secret", "MYCOMPANY"))) // AvalaraAccountCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.POSTAvalaraAccounts(context.Background()).AvalaraAccountCreate(avalaraAccountCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.POSTAvalaraAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAvalaraAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avalaraAccountCreate** | [**AvalaraAccountCreate**](AvalaraAccountCreate.md) |  | 

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


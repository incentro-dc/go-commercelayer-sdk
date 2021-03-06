# \SingletonApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETApplicationApplicationId**](SingletonApi.md#GETApplicationApplicationId) | **Get** /application | Retrieve the application
[**GETOrganizationOrganizationId**](SingletonApi.md#GETOrganizationOrganizationId) | **Get** /organization | Retrieve the organization



## GETApplicationApplicationId

> Application GETApplicationApplicationId(ctx).Execute()

Retrieve the application



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
    resp, r, err := apiClient.SingletonApi.GETApplicationApplicationId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SingletonApi.GETApplicationApplicationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETApplicationApplicationId`: Application
    fmt.Fprintf(os.Stdout, "Response from `SingletonApi.GETApplicationApplicationId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETApplicationApplicationIdRequest struct via the builder pattern


### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrganizationOrganizationId

> Organization GETOrganizationOrganizationId(ctx).Execute()

Retrieve the organization



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
    resp, r, err := apiClient.SingletonApi.GETOrganizationOrganizationId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SingletonApi.GETOrganizationOrganizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrganizationOrganizationId`: Organization
    fmt.Fprintf(os.Stdout, "Response from `SingletonApi.GETOrganizationOrganizationId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrganizationOrganizationIdRequest struct via the builder pattern


### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


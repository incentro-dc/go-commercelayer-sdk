# \OrderValidationRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETOrderValidationRules**](OrderValidationRulesApi.md#GETOrderValidationRules) | **Get** /order_validation_rules | List all order validation rules
[**GETOrderValidationRulesOrderValidationRuleId**](OrderValidationRulesApi.md#GETOrderValidationRulesOrderValidationRuleId) | **Get** /order_validation_rules/{orderValidationRuleId} | Retrieve an order validation rule



## GETOrderValidationRules

> GETOrderValidationRules200Response GETOrderValidationRules(ctx).Execute()

List all order validation rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderValidationRulesApi.GETOrderValidationRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderValidationRulesApi.GETOrderValidationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderValidationRules`: GETOrderValidationRules200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderValidationRulesApi.GETOrderValidationRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderValidationRulesRequest struct via the builder pattern


### Return type

[**GETOrderValidationRules200Response**](GETOrderValidationRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderValidationRulesOrderValidationRuleId

> GETOrderValidationRulesOrderValidationRuleId200Response GETOrderValidationRulesOrderValidationRuleId(ctx, orderValidationRuleId).Execute()

Retrieve an order validation rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
)

func main() {
    orderValidationRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderValidationRulesApi.GETOrderValidationRulesOrderValidationRuleId(context.Background(), orderValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderValidationRulesApi.GETOrderValidationRulesOrderValidationRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderValidationRulesOrderValidationRuleId`: GETOrderValidationRulesOrderValidationRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderValidationRulesApi.GETOrderValidationRulesOrderValidationRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderValidationRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderValidationRulesOrderValidationRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderValidationRulesOrderValidationRuleId200Response**](GETOrderValidationRulesOrderValidationRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


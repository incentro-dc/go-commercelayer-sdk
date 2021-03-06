# \TaxRulesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETaxRulesTaxRuleId**](TaxRulesApi.md#DELETETaxRulesTaxRuleId) | **Delete** /tax_rules/{taxRuleId} | Delete a tax rule
[**GETManualTaxCalculatorIdTaxRules**](TaxRulesApi.md#GETManualTaxCalculatorIdTaxRules) | **Get** /manual_tax_calculators/{manualTaxCalculatorId}/tax_rules | Retrieve the tax rules associated to the manual tax calculator
[**GETTaxRules**](TaxRulesApi.md#GETTaxRules) | **Get** /tax_rules | List all tax rules
[**GETTaxRulesTaxRuleId**](TaxRulesApi.md#GETTaxRulesTaxRuleId) | **Get** /tax_rules/{taxRuleId} | Retrieve a tax rule
[**PATCHTaxRulesTaxRuleId**](TaxRulesApi.md#PATCHTaxRulesTaxRuleId) | **Patch** /tax_rules/{taxRuleId} | Update a tax rule
[**POSTTaxRules**](TaxRulesApi.md#POSTTaxRules) | **Post** /tax_rules | Create a tax rule



## DELETETaxRulesTaxRuleId

> DELETETaxRulesTaxRuleId(ctx, taxRuleId).Execute()

Delete a tax rule



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
    taxRuleId := "taxRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.DELETETaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.DELETETaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETaxRulesTaxRuleIdRequest struct via the builder pattern


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


## GETManualTaxCalculatorIdTaxRules

> GETManualTaxCalculatorIdTaxRules(ctx, manualTaxCalculatorId).Execute()

Retrieve the tax rules associated to the manual tax calculator



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
    manualTaxCalculatorId := "manualTaxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.GETManualTaxCalculatorIdTaxRules(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETManualTaxCalculatorIdTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualTaxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorIdTaxRulesRequest struct via the builder pattern


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


## GETTaxRules

> GETTaxRules(ctx).Execute()

List all tax rules



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
    resp, r, err := apiClient.TaxRulesApi.GETTaxRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxRulesRequest struct via the builder pattern


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


## GETTaxRulesTaxRuleId

> TaxRule GETTaxRulesTaxRuleId(ctx, taxRuleId).Execute()

Retrieve a tax rule



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
    taxRuleId := "taxRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.GETTaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETTaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxRulesTaxRuleId`: TaxRule
    fmt.Fprintf(os.Stdout, "Response from `TaxRulesApi.GETTaxRulesTaxRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxRulesTaxRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxRule**](TaxRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHTaxRulesTaxRuleId

> PATCHTaxRulesTaxRuleId(ctx, taxRuleId).TaxRuleUpdate(taxRuleUpdate).Execute()

Update a tax rule



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
    taxRuleUpdate := *openapiclient.NewTaxRuleUpdate(*openapiclient.NewTaxRuleUpdateData("tax_rules", "XGZwpOSrWL", *openapiclient.NewTaxRuleUpdateDataAttributes())) // TaxRuleUpdate | 
    taxRuleId := "taxRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.PATCHTaxRulesTaxRuleId(context.Background(), taxRuleId).TaxRuleUpdate(taxRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.PATCHTaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHTaxRulesTaxRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxRuleUpdate** | [**TaxRuleUpdate**](TaxRuleUpdate.md) |  | 


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


## POSTTaxRules

> POSTTaxRules(ctx).TaxRuleCreate(taxRuleCreate).Execute()

Create a tax rule



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
    taxRuleCreate := *openapiclient.NewTaxRuleCreate(*openapiclient.NewTaxRuleCreateData("tax_rules", *openapiclient.NewTaxRuleCreateDataAttributes("Fixed 22%"))) // TaxRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.POSTTaxRules(context.Background()).TaxRuleCreate(taxRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.POSTTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTaxRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxRuleCreate** | [**TaxRuleCreate**](TaxRuleCreate.md) |  | 

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


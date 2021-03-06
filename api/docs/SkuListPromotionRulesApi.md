# \SkuListPromotionRulesApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#DELETESkuListPromotionRulesSkuListPromotionRuleId) | **Delete** /sku_list_promotion_rules/{skuListPromotionRuleId} | Delete a SKU list promotion rule
[**GETExternalPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETExternalPromotionIdSkuListPromotionRule) | **Get** /external_promotions/{externalPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFixedAmountPromotionIdSkuListPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFixedPricePromotionIdSkuListPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFreeGiftPromotionIdSkuListPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFreeShippingPromotionIdSkuListPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETPercentageDiscountPromotionIdSkuListPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the percentage discount promotion
[**GETPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETPromotionIdSkuListPromotionRule) | **Get** /promotions/{promotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the promotion
[**GETSkuListPromotionRules**](SkuListPromotionRulesApi.md#GETSkuListPromotionRules) | **Get** /sku_list_promotion_rules | List all SKU list promotion rules
[**GETSkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#GETSkuListPromotionRulesSkuListPromotionRuleId) | **Get** /sku_list_promotion_rules/{skuListPromotionRuleId} | Retrieve a SKU list promotion rule
[**PATCHSkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#PATCHSkuListPromotionRulesSkuListPromotionRuleId) | **Patch** /sku_list_promotion_rules/{skuListPromotionRuleId} | Update a SKU list promotion rule
[**POSTSkuListPromotionRules**](SkuListPromotionRulesApi.md#POSTSkuListPromotionRules) | **Post** /sku_list_promotion_rules | Create a SKU list promotion rule



## DELETESkuListPromotionRulesSkuListPromotionRuleId

> DELETESkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).Execute()

Delete a SKU list promotion rule



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
    skuListPromotionRuleId := "skuListPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.DELETESkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.DELETESkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


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


## GETExternalPromotionIdSkuListPromotionRule

> GETExternalPromotionIdSkuListPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the sku list promotion rule associated to the external promotion



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
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETExternalPromotionIdSkuListPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETExternalPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdSkuListPromotionRule

> GETFixedAmountPromotionIdSkuListPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the sku list promotion rule associated to the fixed amount promotion



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
    fixedAmountPromotionId := "fixedAmountPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETFixedAmountPromotionIdSkuListPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFixedAmountPromotionIdSkuListPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFixedPricePromotionIdSkuListPromotionRule

> GETFixedPricePromotionIdSkuListPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the sku list promotion rule associated to the fixed price promotion



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
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETFixedPricePromotionIdSkuListPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFixedPricePromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdSkuListPromotionRule

> GETFreeGiftPromotionIdSkuListPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the sku list promotion rule associated to the free gift promotion



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
    freeGiftPromotionId := "freeGiftPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETFreeGiftPromotionIdSkuListPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFreeGiftPromotionIdSkuListPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeGiftPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdSkuListPromotionRule

> GETFreeShippingPromotionIdSkuListPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the sku list promotion rule associated to the free shipping promotion



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
    freeShippingPromotionId := "freeShippingPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETFreeShippingPromotionIdSkuListPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFreeShippingPromotionIdSkuListPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdSkuListPromotionRule

> GETPercentageDiscountPromotionIdSkuListPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the sku list promotion rule associated to the percentage discount promotion



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
    percentageDiscountPromotionId := "percentageDiscountPromotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETPercentageDiscountPromotionIdSkuListPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETPercentageDiscountPromotionIdSkuListPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**percentageDiscountPromotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETPromotionIdSkuListPromotionRule

> GETPromotionIdSkuListPromotionRule(ctx, promotionId).Execute()

Retrieve the sku list promotion rule associated to the promotion



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
    promotionId := "promotionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETPromotionIdSkuListPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETPromotionIdSkuListPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETSkuListPromotionRules

> GETSkuListPromotionRules(ctx).Execute()

List all SKU list promotion rules



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
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETSkuListPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListPromotionRulesRequest struct via the builder pattern


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


## GETSkuListPromotionRulesSkuListPromotionRuleId

> SkuListPromotionRule GETSkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).Execute()

Retrieve a SKU list promotion rule



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
    skuListPromotionRuleId := "skuListPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListPromotionRulesSkuListPromotionRuleId`: SkuListPromotionRule
    fmt.Fprintf(os.Stdout, "Response from `SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkuListPromotionRule**](SkuListPromotionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSkuListPromotionRulesSkuListPromotionRuleId

> PATCHSkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).SkuListPromotionRuleUpdate(skuListPromotionRuleUpdate).Execute()

Update a SKU list promotion rule



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
    skuListPromotionRuleUpdate := *openapiclient.NewSkuListPromotionRuleUpdate(*openapiclient.NewSkuListPromotionRuleUpdateData("sku_list_promotion_rules", "XGZwpOSrWL", *openapiclient.NewSkuListPromotionRuleCreateDataAttributes())) // SkuListPromotionRuleUpdate | 
    skuListPromotionRuleId := "skuListPromotionRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).SkuListPromotionRuleUpdate(skuListPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListPromotionRuleUpdate** | [**SkuListPromotionRuleUpdate**](SkuListPromotionRuleUpdate.md) |  | 


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


## POSTSkuListPromotionRules

> POSTSkuListPromotionRules(ctx).SkuListPromotionRuleCreate(skuListPromotionRuleCreate).Execute()

Create a SKU list promotion rule



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
    skuListPromotionRuleCreate := *openapiclient.NewSkuListPromotionRuleCreate(*openapiclient.NewSkuListPromotionRuleCreateData("sku_list_promotion_rules", *openapiclient.NewSkuListPromotionRuleCreateDataAttributes())) // SkuListPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.POSTSkuListPromotionRules(context.Background()).SkuListPromotionRuleCreate(skuListPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.POSTSkuListPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSkuListPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListPromotionRuleCreate** | [**SkuListPromotionRuleCreate**](SkuListPromotionRuleCreate.md) |  | 

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


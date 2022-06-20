# \MarketsApi

All URIs are relative to *https://core.commercelayer.io/users/sign_in*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMarketsMarketId**](MarketsApi.md#DELETEMarketsMarketId) | **Delete** /markets/{marketId} | Delete a market
[**GETAvalaraAccountIdMarkets**](MarketsApi.md#GETAvalaraAccountIdMarkets) | **Get** /avalara_accounts/{avalaraAccountId}/markets | Retrieve the markets associated to the avalara account
[**GETBillingInfoValidationRuleIdMarket**](MarketsApi.md#GETBillingInfoValidationRuleIdMarket) | **Get** /billing_info_validation_rules/{billingInfoValidationRuleId}/market | Retrieve the market associated to the billing info validation rule
[**GETBundleIdMarket**](MarketsApi.md#GETBundleIdMarket) | **Get** /bundles/{bundleId}/market | Retrieve the market associated to the bundle
[**GETCarrierAccountIdMarket**](MarketsApi.md#GETCarrierAccountIdMarket) | **Get** /carrier_accounts/{carrierAccountId}/market | Retrieve the market associated to the carrier account
[**GETCustomerGroupIdMarkets**](MarketsApi.md#GETCustomerGroupIdMarkets) | **Get** /customer_groups/{customerGroupId}/markets | Retrieve the markets associated to the customer group
[**GETExternalPromotionIdMarket**](MarketsApi.md#GETExternalPromotionIdMarket) | **Get** /external_promotions/{externalPromotionId}/market | Retrieve the market associated to the external promotion
[**GETExternalTaxCalculatorIdMarkets**](MarketsApi.md#GETExternalTaxCalculatorIdMarkets) | **Get** /external_tax_calculators/{externalTaxCalculatorId}/markets | Retrieve the markets associated to the external tax calculator
[**GETFixedAmountPromotionIdMarket**](MarketsApi.md#GETFixedAmountPromotionIdMarket) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/market | Retrieve the market associated to the fixed amount promotion
[**GETFixedPricePromotionIdMarket**](MarketsApi.md#GETFixedPricePromotionIdMarket) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/market | Retrieve the market associated to the fixed price promotion
[**GETFreeGiftPromotionIdMarket**](MarketsApi.md#GETFreeGiftPromotionIdMarket) | **Get** /free_gift_promotions/{freeGiftPromotionId}/market | Retrieve the market associated to the free gift promotion
[**GETFreeShippingPromotionIdMarket**](MarketsApi.md#GETFreeShippingPromotionIdMarket) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/market | Retrieve the market associated to the free shipping promotion
[**GETGiftCardIdMarket**](MarketsApi.md#GETGiftCardIdMarket) | **Get** /gift_cards/{giftCardId}/market | Retrieve the market associated to the gift card
[**GETInStockSubscriptionIdMarket**](MarketsApi.md#GETInStockSubscriptionIdMarket) | **Get** /in_stock_subscriptions/{inStockSubscriptionId}/market | Retrieve the market associated to the in stock subscription
[**GETManualTaxCalculatorIdMarkets**](MarketsApi.md#GETManualTaxCalculatorIdMarkets) | **Get** /manual_tax_calculators/{manualTaxCalculatorId}/markets | Retrieve the markets associated to the manual tax calculator
[**GETMarkets**](MarketsApi.md#GETMarkets) | **Get** /markets | List all markets
[**GETMarketsMarketId**](MarketsApi.md#GETMarketsMarketId) | **Get** /markets/{marketId} | Retrieve a market
[**GETOrderIdMarket**](MarketsApi.md#GETOrderIdMarket) | **Get** /orders/{orderId}/market | Retrieve the market associated to the order
[**GETOrderSubscriptionIdMarket**](MarketsApi.md#GETOrderSubscriptionIdMarket) | **Get** /order_subscriptions/{orderSubscriptionId}/market | Retrieve the market associated to the order subscription
[**GETOrderValidationRuleIdMarket**](MarketsApi.md#GETOrderValidationRuleIdMarket) | **Get** /order_validation_rules/{orderValidationRuleId}/market | Retrieve the market associated to the order validation rule
[**GETPaymentMethodIdMarket**](MarketsApi.md#GETPaymentMethodIdMarket) | **Get** /payment_methods/{paymentMethodId}/market | Retrieve the market associated to the payment method
[**GETPercentageDiscountPromotionIdMarket**](MarketsApi.md#GETPercentageDiscountPromotionIdMarket) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/market | Retrieve the market associated to the percentage discount promotion
[**GETPromotionIdMarket**](MarketsApi.md#GETPromotionIdMarket) | **Get** /promotions/{promotionId}/market | Retrieve the market associated to the promotion
[**GETShippingMethodIdMarket**](MarketsApi.md#GETShippingMethodIdMarket) | **Get** /shipping_methods/{shippingMethodId}/market | Retrieve the market associated to the shipping method
[**GETSkuOptionIdMarket**](MarketsApi.md#GETSkuOptionIdMarket) | **Get** /sku_options/{skuOptionId}/market | Retrieve the market associated to the SKU option
[**GETTaxCalculatorIdMarkets**](MarketsApi.md#GETTaxCalculatorIdMarkets) | **Get** /tax_calculators/{taxCalculatorId}/markets | Retrieve the markets associated to the tax calculator
[**GETTaxjarAccountIdMarkets**](MarketsApi.md#GETTaxjarAccountIdMarkets) | **Get** /taxjar_accounts/{taxjarAccountId}/markets | Retrieve the markets associated to the taxjar account
[**PATCHMarketsMarketId**](MarketsApi.md#PATCHMarketsMarketId) | **Patch** /markets/{marketId} | Update a market
[**POSTMarkets**](MarketsApi.md#POSTMarkets) | **Post** /markets | Create a market



## DELETEMarketsMarketId

> DELETEMarketsMarketId(ctx, marketId).Execute()

Delete a market



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
    marketId := "marketId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.DELETEMarketsMarketId(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.DELETEMarketsMarketId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMarketsMarketIdRequest struct via the builder pattern


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


## GETAvalaraAccountIdMarkets

> GETAvalaraAccountIdMarkets(ctx, avalaraAccountId).Execute()

Retrieve the markets associated to the avalara account



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
    resp, r, err := apiClient.MarketsApi.GETAvalaraAccountIdMarkets(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETAvalaraAccountIdMarkets``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETAvalaraAccountIdMarketsRequest struct via the builder pattern


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


## GETBillingInfoValidationRuleIdMarket

> GETBillingInfoValidationRuleIdMarket(ctx, billingInfoValidationRuleId).Execute()

Retrieve the market associated to the billing info validation rule



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
    billingInfoValidationRuleId := "billingInfoValidationRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETBillingInfoValidationRuleIdMarket(context.Background(), billingInfoValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETBillingInfoValidationRuleIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingInfoValidationRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBillingInfoValidationRuleIdMarketRequest struct via the builder pattern


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


## GETBundleIdMarket

> GETBundleIdMarket(ctx, bundleId).Execute()

Retrieve the market associated to the bundle



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
    bundleId := "bundleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETBundleIdMarket(context.Background(), bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETBundleIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBundleIdMarketRequest struct via the builder pattern


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


## GETCarrierAccountIdMarket

> GETCarrierAccountIdMarket(ctx, carrierAccountId).Execute()

Retrieve the market associated to the carrier account



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
    carrierAccountId := "carrierAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETCarrierAccountIdMarket(context.Background(), carrierAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETCarrierAccountIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carrierAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCarrierAccountIdMarketRequest struct via the builder pattern


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


## GETCustomerGroupIdMarkets

> GETCustomerGroupIdMarkets(ctx, customerGroupId).Execute()

Retrieve the markets associated to the customer group



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
    customerGroupId := "customerGroupId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETCustomerGroupIdMarkets(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETCustomerGroupIdMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupIdMarketsRequest struct via the builder pattern


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


## GETExternalPromotionIdMarket

> GETExternalPromotionIdMarket(ctx, externalPromotionId).Execute()

Retrieve the market associated to the external promotion



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
    resp, r, err := apiClient.MarketsApi.GETExternalPromotionIdMarket(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETExternalPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdMarketRequest struct via the builder pattern


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


## GETExternalTaxCalculatorIdMarkets

> GETExternalTaxCalculatorIdMarkets(ctx, externalTaxCalculatorId).Execute()

Retrieve the markets associated to the external tax calculator



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
    externalTaxCalculatorId := "externalTaxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETExternalTaxCalculatorIdMarkets(context.Background(), externalTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETExternalTaxCalculatorIdMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalTaxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalTaxCalculatorIdMarketsRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdMarket

> GETFixedAmountPromotionIdMarket(ctx, fixedAmountPromotionId).Execute()

Retrieve the market associated to the fixed amount promotion



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
    resp, r, err := apiClient.MarketsApi.GETFixedAmountPromotionIdMarket(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETFixedAmountPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdMarketRequest struct via the builder pattern


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


## GETFixedPricePromotionIdMarket

> GETFixedPricePromotionIdMarket(ctx, fixedPricePromotionId).Execute()

Retrieve the market associated to the fixed price promotion



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
    resp, r, err := apiClient.MarketsApi.GETFixedPricePromotionIdMarket(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETFixedPricePromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdMarketRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdMarket

> GETFreeGiftPromotionIdMarket(ctx, freeGiftPromotionId).Execute()

Retrieve the market associated to the free gift promotion



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
    resp, r, err := apiClient.MarketsApi.GETFreeGiftPromotionIdMarket(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETFreeGiftPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdMarketRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdMarket

> GETFreeShippingPromotionIdMarket(ctx, freeShippingPromotionId).Execute()

Retrieve the market associated to the free shipping promotion



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
    resp, r, err := apiClient.MarketsApi.GETFreeShippingPromotionIdMarket(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETFreeShippingPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdMarketRequest struct via the builder pattern


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


## GETGiftCardIdMarket

> GETGiftCardIdMarket(ctx, giftCardId).Execute()

Retrieve the market associated to the gift card



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
    giftCardId := "giftCardId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETGiftCardIdMarket(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETGiftCardIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardIdMarketRequest struct via the builder pattern


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


## GETInStockSubscriptionIdMarket

> GETInStockSubscriptionIdMarket(ctx, inStockSubscriptionId).Execute()

Retrieve the market associated to the in stock subscription



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
    inStockSubscriptionId := "inStockSubscriptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETInStockSubscriptionIdMarket(context.Background(), inStockSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETInStockSubscriptionIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionIdMarketRequest struct via the builder pattern


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


## GETManualTaxCalculatorIdMarkets

> GETManualTaxCalculatorIdMarkets(ctx, manualTaxCalculatorId).Execute()

Retrieve the markets associated to the manual tax calculator



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
    resp, r, err := apiClient.MarketsApi.GETManualTaxCalculatorIdMarkets(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETManualTaxCalculatorIdMarkets``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorIdMarketsRequest struct via the builder pattern


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


## GETMarkets

> GETMarkets(ctx).Execute()

List all markets



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
    resp, r, err := apiClient.MarketsApi.GETMarkets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketsRequest struct via the builder pattern


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


## GETMarketsMarketId

> Market GETMarketsMarketId(ctx, marketId).Execute()

Retrieve a market



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
    marketId := "marketId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETMarketsMarketId(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETMarketsMarketId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMarketsMarketId`: Market
    fmt.Fprintf(os.Stdout, "Response from `MarketsApi.GETMarketsMarketId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketsMarketIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Market**](Market.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdMarket

> GETOrderIdMarket(ctx, orderId).Execute()

Retrieve the market associated to the order



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
    orderId := "orderId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETOrderIdMarket(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETOrderIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderIdMarketRequest struct via the builder pattern


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


## GETOrderSubscriptionIdMarket

> GETOrderSubscriptionIdMarket(ctx, orderSubscriptionId).Execute()

Retrieve the market associated to the order subscription



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
    orderSubscriptionId := "orderSubscriptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETOrderSubscriptionIdMarket(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETOrderSubscriptionIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdMarketRequest struct via the builder pattern


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


## GETOrderValidationRuleIdMarket

> GETOrderValidationRuleIdMarket(ctx, orderValidationRuleId).Execute()

Retrieve the market associated to the order validation rule



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
    orderValidationRuleId := "orderValidationRuleId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETOrderValidationRuleIdMarket(context.Background(), orderValidationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETOrderValidationRuleIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderValidationRuleId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderValidationRuleIdMarketRequest struct via the builder pattern


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


## GETPaymentMethodIdMarket

> GETPaymentMethodIdMarket(ctx, paymentMethodId).Execute()

Retrieve the market associated to the payment method



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
    paymentMethodId := "paymentMethodId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETPaymentMethodIdMarket(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETPaymentMethodIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentMethodIdMarketRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdMarket

> GETPercentageDiscountPromotionIdMarket(ctx, percentageDiscountPromotionId).Execute()

Retrieve the market associated to the percentage discount promotion



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
    resp, r, err := apiClient.MarketsApi.GETPercentageDiscountPromotionIdMarket(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETPercentageDiscountPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdMarketRequest struct via the builder pattern


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


## GETPromotionIdMarket

> GETPromotionIdMarket(ctx, promotionId).Execute()

Retrieve the market associated to the promotion



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
    resp, r, err := apiClient.MarketsApi.GETPromotionIdMarket(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETPromotionIdMarket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdMarketRequest struct via the builder pattern


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


## GETShippingMethodIdMarket

> GETShippingMethodIdMarket(ctx, shippingMethodId).Execute()

Retrieve the market associated to the shipping method



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
    shippingMethodId := "shippingMethodId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETShippingMethodIdMarket(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETShippingMethodIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodIdMarketRequest struct via the builder pattern


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


## GETSkuOptionIdMarket

> GETSkuOptionIdMarket(ctx, skuOptionId).Execute()

Retrieve the market associated to the SKU option



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
    skuOptionId := "skuOptionId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETSkuOptionIdMarket(context.Background(), skuOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETSkuOptionIdMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuOptionId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuOptionIdMarketRequest struct via the builder pattern


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


## GETTaxCalculatorIdMarkets

> GETTaxCalculatorIdMarkets(ctx, taxCalculatorId).Execute()

Retrieve the markets associated to the tax calculator



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
    taxCalculatorId := "taxCalculatorId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETTaxCalculatorIdMarkets(context.Background(), taxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETTaxCalculatorIdMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCalculatorId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCalculatorIdMarketsRequest struct via the builder pattern


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


## GETTaxjarAccountIdMarkets

> GETTaxjarAccountIdMarkets(ctx, taxjarAccountId).Execute()

Retrieve the markets associated to the taxjar account



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
    taxjarAccountId := "taxjarAccountId_example" // string | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.GETTaxjarAccountIdMarkets(context.Background(), taxjarAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GETTaxjarAccountIdMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxjarAccountId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxjarAccountIdMarketsRequest struct via the builder pattern


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


## PATCHMarketsMarketId

> PATCHMarketsMarketId(ctx, marketId).MarketUpdate(marketUpdate).Execute()

Update a market



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
    marketId := "marketId_example" // string | The resource's id
    marketUpdate := *openapiclient.NewMarketUpdate(*openapiclient.NewMarketUpdateData("markets", "XGZwpOSrWL", *openapiclient.NewMarketUpdateDataAttributes())) // MarketUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.PATCHMarketsMarketId(context.Background(), marketId).MarketUpdate(marketUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.PATCHMarketsMarketId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHMarketsMarketIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketUpdate** | [**MarketUpdate**](MarketUpdate.md) |  | 

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


## POSTMarkets

> POSTMarkets(ctx).MarketCreate(marketCreate).Execute()

Create a market



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
    marketCreate := *openapiclient.NewMarketCreate(*openapiclient.NewMarketCreateData("markets", *openapiclient.NewMarketCreateDataAttributes("EU Market"))) // MarketCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketsApi.POSTMarkets(context.Background()).MarketCreate(marketCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.POSTMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMarketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketCreate** | [**MarketCreate**](MarketCreate.md) |  | 

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

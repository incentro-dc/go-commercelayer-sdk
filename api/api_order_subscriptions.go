/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// OrderSubscriptionsApiService OrderSubscriptionsApi service
type OrderSubscriptionsApiService service

type OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest struct {
	ctx                 context.Context
	ApiService          *OrderSubscriptionsApiService
	orderSubscriptionId string
}

func (r OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEOrderSubscriptionsOrderSubscriptionIdExecute(r)
}

/*
DELETEOrderSubscriptionsOrderSubscriptionId Delete an order subscription

Delete an order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderSubscriptionId The resource's id
	@return OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest
*/
func (a *OrderSubscriptionsApiService) DELETEOrderSubscriptionsOrderSubscriptionId(ctx context.Context, orderSubscriptionId string) OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest {
	return OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		orderSubscriptionId: orderSubscriptionId,
	}
}

// Execute executes the request
func (a *OrderSubscriptionsApiService) DELETEOrderSubscriptionsOrderSubscriptionIdExecute(r OrderSubscriptionsApiDELETEOrderSubscriptionsOrderSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.DELETEOrderSubscriptionsOrderSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions/{orderSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderSubscriptionId"+"}", url.PathEscape(parameterToString(r.orderSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *OrderSubscriptionsApiService
	customerId string
}

func (r OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdOrderSubscriptionsExecute(r)
}

/*
GETCustomerIdOrderSubscriptions Retrieve the order subscriptions associated to the customer

Retrieve the order subscriptions associated to the customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId The resource's id
	@return OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest
*/
func (a *OrderSubscriptionsApiService) GETCustomerIdOrderSubscriptions(ctx context.Context, customerId string) OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest {
	return OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *OrderSubscriptionsApiService) GETCustomerIdOrderSubscriptionsExecute(r OrderSubscriptionsApiGETCustomerIdOrderSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.GETCustomerIdOrderSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/order_subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest struct {
	ctx         context.Context
	ApiService  *OrderSubscriptionsApiService
	orderCopyId string
}

func (r OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderCopyIdOrderSubscriptionExecute(r)
}

/*
GETOrderCopyIdOrderSubscription Retrieve the order subscription associated to the order copy

Retrieve the order subscription associated to the order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderCopyId The resource's id
	@return OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest
*/
func (a *OrderSubscriptionsApiService) GETOrderCopyIdOrderSubscription(ctx context.Context, orderCopyId string) OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest {
	return OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest{
		ApiService:  a,
		ctx:         ctx,
		orderCopyId: orderCopyId,
	}
}

// Execute executes the request
func (a *OrderSubscriptionsApiService) GETOrderCopyIdOrderSubscriptionExecute(r OrderSubscriptionsApiGETOrderCopyIdOrderSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.GETOrderCopyIdOrderSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies/{orderCopyId}/order_subscription"
	localVarPath = strings.Replace(localVarPath, "{"+"orderCopyId"+"}", url.PathEscape(parameterToString(r.orderCopyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *OrderSubscriptionsApiService
	orderId    string
}

func (r OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdOrderSubscriptionsExecute(r)
}

/*
GETOrderIdOrderSubscriptions Retrieve the order subscriptions associated to the order

Retrieve the order subscriptions associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest
*/
func (a *OrderSubscriptionsApiService) GETOrderIdOrderSubscriptions(ctx context.Context, orderId string) OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest {
	return OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *OrderSubscriptionsApiService) GETOrderIdOrderSubscriptionsExecute(r OrderSubscriptionsApiGETOrderIdOrderSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.GETOrderIdOrderSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/order_subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterToString(r.orderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OrderSubscriptionsApiGETOrderSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *OrderSubscriptionsApiService
}

func (r OrderSubscriptionsApiGETOrderSubscriptionsRequest) Execute() (*GETOrderSubscriptions200Response, *http.Response, error) {
	return r.ApiService.GETOrderSubscriptionsExecute(r)
}

/*
GETOrderSubscriptions List all order subscriptions

List all order subscriptions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OrderSubscriptionsApiGETOrderSubscriptionsRequest
*/
func (a *OrderSubscriptionsApiService) GETOrderSubscriptions(ctx context.Context) OrderSubscriptionsApiGETOrderSubscriptionsRequest {
	return OrderSubscriptionsApiGETOrderSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETOrderSubscriptions200Response
func (a *OrderSubscriptionsApiService) GETOrderSubscriptionsExecute(r OrderSubscriptionsApiGETOrderSubscriptionsRequest) (*GETOrderSubscriptions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETOrderSubscriptions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.GETOrderSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest struct {
	ctx                 context.Context
	ApiService          *OrderSubscriptionsApiService
	orderSubscriptionId string
}

func (r OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest) Execute() (*GETOrderSubscriptionsOrderSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.GETOrderSubscriptionsOrderSubscriptionIdExecute(r)
}

/*
GETOrderSubscriptionsOrderSubscriptionId Retrieve an order subscription

Retrieve an order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderSubscriptionId The resource's id
	@return OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest
*/
func (a *OrderSubscriptionsApiService) GETOrderSubscriptionsOrderSubscriptionId(ctx context.Context, orderSubscriptionId string) OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest {
	return OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		orderSubscriptionId: orderSubscriptionId,
	}
}

// Execute executes the request
//
//	@return GETOrderSubscriptionsOrderSubscriptionId200Response
func (a *OrderSubscriptionsApiService) GETOrderSubscriptionsOrderSubscriptionIdExecute(r OrderSubscriptionsApiGETOrderSubscriptionsOrderSubscriptionIdRequest) (*GETOrderSubscriptionsOrderSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETOrderSubscriptionsOrderSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.GETOrderSubscriptionsOrderSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions/{orderSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderSubscriptionId"+"}", url.PathEscape(parameterToString(r.orderSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest struct {
	ctx                     context.Context
	ApiService              *OrderSubscriptionsApiService
	orderSubscriptionUpdate *OrderSubscriptionUpdate
	orderSubscriptionId     string
}

func (r OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest) OrderSubscriptionUpdate(orderSubscriptionUpdate OrderSubscriptionUpdate) OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest {
	r.orderSubscriptionUpdate = &orderSubscriptionUpdate
	return r
}

func (r OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest) Execute() (*PATCHOrderSubscriptionsOrderSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.PATCHOrderSubscriptionsOrderSubscriptionIdExecute(r)
}

/*
PATCHOrderSubscriptionsOrderSubscriptionId Update an order subscription

Update an order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderSubscriptionId The resource's id
	@return OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest
*/
func (a *OrderSubscriptionsApiService) PATCHOrderSubscriptionsOrderSubscriptionId(ctx context.Context, orderSubscriptionId string) OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest {
	return OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		orderSubscriptionId: orderSubscriptionId,
	}
}

// Execute executes the request
//
//	@return PATCHOrderSubscriptionsOrderSubscriptionId200Response
func (a *OrderSubscriptionsApiService) PATCHOrderSubscriptionsOrderSubscriptionIdExecute(r OrderSubscriptionsApiPATCHOrderSubscriptionsOrderSubscriptionIdRequest) (*PATCHOrderSubscriptionsOrderSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHOrderSubscriptionsOrderSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.PATCHOrderSubscriptionsOrderSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions/{orderSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderSubscriptionId"+"}", url.PathEscape(parameterToString(r.orderSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderSubscriptionUpdate == nil {
		return localVarReturnValue, nil, reportError("orderSubscriptionUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orderSubscriptionUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrderSubscriptionsApiPOSTOrderSubscriptionsRequest struct {
	ctx                     context.Context
	ApiService              *OrderSubscriptionsApiService
	orderSubscriptionCreate *OrderSubscriptionCreate
}

func (r OrderSubscriptionsApiPOSTOrderSubscriptionsRequest) OrderSubscriptionCreate(orderSubscriptionCreate OrderSubscriptionCreate) OrderSubscriptionsApiPOSTOrderSubscriptionsRequest {
	r.orderSubscriptionCreate = &orderSubscriptionCreate
	return r
}

func (r OrderSubscriptionsApiPOSTOrderSubscriptionsRequest) Execute() (*POSTOrderSubscriptions201Response, *http.Response, error) {
	return r.ApiService.POSTOrderSubscriptionsExecute(r)
}

/*
POSTOrderSubscriptions Create an order subscription

Create an order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OrderSubscriptionsApiPOSTOrderSubscriptionsRequest
*/
func (a *OrderSubscriptionsApiService) POSTOrderSubscriptions(ctx context.Context) OrderSubscriptionsApiPOSTOrderSubscriptionsRequest {
	return OrderSubscriptionsApiPOSTOrderSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTOrderSubscriptions201Response
func (a *OrderSubscriptionsApiService) POSTOrderSubscriptionsExecute(r OrderSubscriptionsApiPOSTOrderSubscriptionsRequest) (*POSTOrderSubscriptions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTOrderSubscriptions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderSubscriptionsApiService.POSTOrderSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderSubscriptionCreate == nil {
		return localVarReturnValue, nil, reportError("orderSubscriptionCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orderSubscriptionCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

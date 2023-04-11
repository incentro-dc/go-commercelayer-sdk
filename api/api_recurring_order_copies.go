/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
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

// RecurringOrderCopiesApiService RecurringOrderCopiesApi service
type RecurringOrderCopiesApiService service

type RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest struct {
	ctx                  context.Context
	ApiService           *RecurringOrderCopiesApiService
	recurringOrderCopyId interface{}
}

func (r RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETERecurringOrderCopiesRecurringOrderCopyIdExecute(r)
}

/*
DELETERecurringOrderCopiesRecurringOrderCopyId Delete a recurring order copy

Delete a recurring order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recurringOrderCopyId The resource's id
	@return RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest
*/
func (a *RecurringOrderCopiesApiService) DELETERecurringOrderCopiesRecurringOrderCopyId(ctx context.Context, recurringOrderCopyId interface{}) RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest {
	return RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		recurringOrderCopyId: recurringOrderCopyId,
	}
}

// Execute executes the request
func (a *RecurringOrderCopiesApiService) DELETERecurringOrderCopiesRecurringOrderCopyIdExecute(r RecurringOrderCopiesApiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.DELETERecurringOrderCopiesRecurringOrderCopyId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/recurring_order_copies/{recurringOrderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recurringOrderCopyId"+"}", url.PathEscape(parameterToString(r.recurringOrderCopyId, "")), -1)

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

type RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest struct {
	ctx        context.Context
	ApiService *RecurringOrderCopiesApiService
	orderId    interface{}
}

func (r RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdRecurringOrderCopiesExecute(r)
}

/*
GETOrderIdRecurringOrderCopies Retrieve the recurring order copies associated to the order

Retrieve the recurring order copies associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest
*/
func (a *RecurringOrderCopiesApiService) GETOrderIdRecurringOrderCopies(ctx context.Context, orderId interface{}) RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest {
	return RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *RecurringOrderCopiesApiService) GETOrderIdRecurringOrderCopiesExecute(r RecurringOrderCopiesApiGETOrderIdRecurringOrderCopiesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.GETOrderIdRecurringOrderCopies")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/recurring_order_copies"
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

type RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest struct {
	ctx                 context.Context
	ApiService          *RecurringOrderCopiesApiService
	orderSubscriptionId interface{}
}

func (r RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderSubscriptionIdRecurringOrderCopiesExecute(r)
}

/*
GETOrderSubscriptionIdRecurringOrderCopies Retrieve the recurring order copies associated to the order subscription

Retrieve the recurring order copies associated to the order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderSubscriptionId The resource's id
	@return RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest
*/
func (a *RecurringOrderCopiesApiService) GETOrderSubscriptionIdRecurringOrderCopies(ctx context.Context, orderSubscriptionId interface{}) RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest {
	return RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest{
		ApiService:          a,
		ctx:                 ctx,
		orderSubscriptionId: orderSubscriptionId,
	}
}

// Execute executes the request
func (a *RecurringOrderCopiesApiService) GETOrderSubscriptionIdRecurringOrderCopiesExecute(r RecurringOrderCopiesApiGETOrderSubscriptionIdRecurringOrderCopiesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.GETOrderSubscriptionIdRecurringOrderCopies")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions/{orderSubscriptionId}/recurring_order_copies"
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

type RecurringOrderCopiesApiGETRecurringOrderCopiesRequest struct {
	ctx        context.Context
	ApiService *RecurringOrderCopiesApiService
}

func (r RecurringOrderCopiesApiGETRecurringOrderCopiesRequest) Execute() (*GETRecurringOrderCopies200Response, *http.Response, error) {
	return r.ApiService.GETRecurringOrderCopiesExecute(r)
}

/*
GETRecurringOrderCopies List all recurring order copies

List all recurring order copies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecurringOrderCopiesApiGETRecurringOrderCopiesRequest
*/
func (a *RecurringOrderCopiesApiService) GETRecurringOrderCopies(ctx context.Context) RecurringOrderCopiesApiGETRecurringOrderCopiesRequest {
	return RecurringOrderCopiesApiGETRecurringOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETRecurringOrderCopies200Response
func (a *RecurringOrderCopiesApiService) GETRecurringOrderCopiesExecute(r RecurringOrderCopiesApiGETRecurringOrderCopiesRequest) (*GETRecurringOrderCopies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETRecurringOrderCopies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.GETRecurringOrderCopies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/recurring_order_copies"

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

type RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest struct {
	ctx                  context.Context
	ApiService           *RecurringOrderCopiesApiService
	recurringOrderCopyId interface{}
}

func (r RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest) Execute() (*GETRecurringOrderCopiesRecurringOrderCopyId200Response, *http.Response, error) {
	return r.ApiService.GETRecurringOrderCopiesRecurringOrderCopyIdExecute(r)
}

/*
GETRecurringOrderCopiesRecurringOrderCopyId Retrieve a recurring order copy

Retrieve a recurring order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recurringOrderCopyId The resource's id
	@return RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest
*/
func (a *RecurringOrderCopiesApiService) GETRecurringOrderCopiesRecurringOrderCopyId(ctx context.Context, recurringOrderCopyId interface{}) RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest {
	return RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		recurringOrderCopyId: recurringOrderCopyId,
	}
}

// Execute executes the request
//
//	@return GETRecurringOrderCopiesRecurringOrderCopyId200Response
func (a *RecurringOrderCopiesApiService) GETRecurringOrderCopiesRecurringOrderCopyIdExecute(r RecurringOrderCopiesApiGETRecurringOrderCopiesRecurringOrderCopyIdRequest) (*GETRecurringOrderCopiesRecurringOrderCopyId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETRecurringOrderCopiesRecurringOrderCopyId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.GETRecurringOrderCopiesRecurringOrderCopyId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/recurring_order_copies/{recurringOrderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recurringOrderCopyId"+"}", url.PathEscape(parameterToString(r.recurringOrderCopyId, "")), -1)

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

type RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest struct {
	ctx                      context.Context
	ApiService               *RecurringOrderCopiesApiService
	recurringOrderCopyUpdate *RecurringOrderCopyUpdate
	recurringOrderCopyId     interface{}
}

func (r RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest) RecurringOrderCopyUpdate(recurringOrderCopyUpdate RecurringOrderCopyUpdate) RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest {
	r.recurringOrderCopyUpdate = &recurringOrderCopyUpdate
	return r
}

func (r RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest) Execute() (*PATCHRecurringOrderCopiesRecurringOrderCopyId200Response, *http.Response, error) {
	return r.ApiService.PATCHRecurringOrderCopiesRecurringOrderCopyIdExecute(r)
}

/*
PATCHRecurringOrderCopiesRecurringOrderCopyId Update a recurring order copy

Update a recurring order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recurringOrderCopyId The resource's id
	@return RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest
*/
func (a *RecurringOrderCopiesApiService) PATCHRecurringOrderCopiesRecurringOrderCopyId(ctx context.Context, recurringOrderCopyId interface{}) RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest {
	return RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		recurringOrderCopyId: recurringOrderCopyId,
	}
}

// Execute executes the request
//
//	@return PATCHRecurringOrderCopiesRecurringOrderCopyId200Response
func (a *RecurringOrderCopiesApiService) PATCHRecurringOrderCopiesRecurringOrderCopyIdExecute(r RecurringOrderCopiesApiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest) (*PATCHRecurringOrderCopiesRecurringOrderCopyId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.PATCHRecurringOrderCopiesRecurringOrderCopyId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/recurring_order_copies/{recurringOrderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recurringOrderCopyId"+"}", url.PathEscape(parameterToString(r.recurringOrderCopyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recurringOrderCopyUpdate == nil {
		return localVarReturnValue, nil, reportError("recurringOrderCopyUpdate is required and must be specified")
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
	localVarPostBody = r.recurringOrderCopyUpdate
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

type RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest struct {
	ctx                      context.Context
	ApiService               *RecurringOrderCopiesApiService
	recurringOrderCopyCreate *RecurringOrderCopyCreate
}

func (r RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest) RecurringOrderCopyCreate(recurringOrderCopyCreate RecurringOrderCopyCreate) RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest {
	r.recurringOrderCopyCreate = &recurringOrderCopyCreate
	return r
}

func (r RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest) Execute() (*POSTRecurringOrderCopies201Response, *http.Response, error) {
	return r.ApiService.POSTRecurringOrderCopiesExecute(r)
}

/*
POSTRecurringOrderCopies Create a recurring order copy

Create a recurring order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest
*/
func (a *RecurringOrderCopiesApiService) POSTRecurringOrderCopies(ctx context.Context) RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest {
	return RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTRecurringOrderCopies201Response
func (a *RecurringOrderCopiesApiService) POSTRecurringOrderCopiesExecute(r RecurringOrderCopiesApiPOSTRecurringOrderCopiesRequest) (*POSTRecurringOrderCopies201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTRecurringOrderCopies201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringOrderCopiesApiService.POSTRecurringOrderCopies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/recurring_order_copies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recurringOrderCopyCreate == nil {
		return localVarReturnValue, nil, reportError("recurringOrderCopyCreate is required and must be specified")
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
	localVarPostBody = r.recurringOrderCopyCreate
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

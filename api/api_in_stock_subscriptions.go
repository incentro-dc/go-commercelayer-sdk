/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
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

// InStockSubscriptionsApiService InStockSubscriptionsApi service
type InStockSubscriptionsApiService service

type ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                   context.Context
	ApiService            *InStockSubscriptionsApiService
	inStockSubscriptionId string
}

func (r ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
DELETEInStockSubscriptionsInStockSubscriptionId Delete an in stock subscription

Delete an in stock subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inStockSubscriptionId The resource's id
 @return ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) DELETEInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId string) ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest {
	return ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
func (a *InStockSubscriptionsApiService) DELETEInStockSubscriptionsInStockSubscriptionIdExecute(r ApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.DELETEInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterToString(r.inStockSubscriptionId, "")), -1)

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

type ApiGETInStockSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *InStockSubscriptionsApiService
}

func (r ApiGETInStockSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInStockSubscriptionsExecute(r)
}

/*
GETInStockSubscriptions List all in stock subscriptions

List all in stock subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETInStockSubscriptionsRequest
*/
func (a *InStockSubscriptionsApiService) GETInStockSubscriptions(ctx context.Context) ApiGETInStockSubscriptionsRequest {
	return ApiGETInStockSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsExecute(r ApiGETInStockSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.GETInStockSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions"

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

type ApiGETInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                   context.Context
	ApiService            *InStockSubscriptionsApiService
	inStockSubscriptionId string
}

func (r ApiGETInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*InStockSubscription, *http.Response, error) {
	return r.ApiService.GETInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
GETInStockSubscriptionsInStockSubscriptionId Retrieve an in stock subscription

Retrieve an in stock subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inStockSubscriptionId The resource's id
 @return ApiGETInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId string) ApiGETInStockSubscriptionsInStockSubscriptionIdRequest {
	return ApiGETInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
//  @return InStockSubscription
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsInStockSubscriptionIdExecute(r ApiGETInStockSubscriptionsInStockSubscriptionIdRequest) (*InStockSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InStockSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.GETInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterToString(r.inStockSubscriptionId, "")), -1)

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

type ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                       context.Context
	ApiService                *InStockSubscriptionsApiService
	inStockSubscriptionUpdate *InStockSubscriptionUpdate
	inStockSubscriptionId     string
}

func (r ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) InStockSubscriptionUpdate(inStockSubscriptionUpdate InStockSubscriptionUpdate) ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest {
	r.inStockSubscriptionUpdate = &inStockSubscriptionUpdate
	return r
}

func (r ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
PATCHInStockSubscriptionsInStockSubscriptionId Update an in stock subscription

Update an in stock subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inStockSubscriptionId The resource's id
 @return ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) PATCHInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId string) ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest {
	return ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
func (a *InStockSubscriptionsApiService) PATCHInStockSubscriptionsInStockSubscriptionIdExecute(r ApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.PATCHInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterToString(r.inStockSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inStockSubscriptionUpdate == nil {
		return nil, reportError("inStockSubscriptionUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.inStockSubscriptionUpdate
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

type ApiPOSTInStockSubscriptionsRequest struct {
	ctx                       context.Context
	ApiService                *InStockSubscriptionsApiService
	inStockSubscriptionCreate *InStockSubscriptionCreate
}

func (r ApiPOSTInStockSubscriptionsRequest) InStockSubscriptionCreate(inStockSubscriptionCreate InStockSubscriptionCreate) ApiPOSTInStockSubscriptionsRequest {
	r.inStockSubscriptionCreate = &inStockSubscriptionCreate
	return r
}

func (r ApiPOSTInStockSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTInStockSubscriptionsExecute(r)
}

/*
POSTInStockSubscriptions Create an in stock subscription

Create an in stock subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTInStockSubscriptionsRequest
*/
func (a *InStockSubscriptionsApiService) POSTInStockSubscriptions(ctx context.Context) ApiPOSTInStockSubscriptionsRequest {
	return ApiPOSTInStockSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *InStockSubscriptionsApiService) POSTInStockSubscriptionsExecute(r ApiPOSTInStockSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.POSTInStockSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inStockSubscriptionCreate == nil {
		return nil, reportError("inStockSubscriptionCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.inStockSubscriptionCreate
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

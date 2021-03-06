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

// BraintreePaymentsApiService BraintreePaymentsApi service
type BraintreePaymentsApiService service

type ApiDELETEBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreePaymentId string
}

func (r ApiDELETEBraintreePaymentsBraintreePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
DELETEBraintreePaymentsBraintreePaymentId Delete a braintree payment

Delete a braintree payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreePaymentId The resource's id
 @return ApiDELETEBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) DELETEBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId string) ApiDELETEBraintreePaymentsBraintreePaymentIdRequest {
	return ApiDELETEBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) DELETEBraintreePaymentsBraintreePaymentIdExecute(r ApiDELETEBraintreePaymentsBraintreePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.DELETEBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterToString(r.braintreePaymentId, "")), -1)

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

type ApiGETBraintreeGatewayIdBraintreePaymentsRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreeGatewayId string
}

func (r ApiGETBraintreeGatewayIdBraintreePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBraintreeGatewayIdBraintreePaymentsExecute(r)
}

/*
GETBraintreeGatewayIdBraintreePayments Retrieve the braintree payments associated to the braintree gateway

Retrieve the braintree payments associated to the braintree gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreeGatewayId The resource's id
 @return ApiGETBraintreeGatewayIdBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreeGatewayIdBraintreePayments(ctx context.Context, braintreeGatewayId string) ApiGETBraintreeGatewayIdBraintreePaymentsRequest {
	return ApiGETBraintreeGatewayIdBraintreePaymentsRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreeGatewayId: braintreeGatewayId,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) GETBraintreeGatewayIdBraintreePaymentsExecute(r ApiGETBraintreeGatewayIdBraintreePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreeGatewayIdBraintreePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways/{braintreeGatewayId}/braintree_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreeGatewayId"+"}", url.PathEscape(parameterToString(r.braintreeGatewayId, "")), -1)

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

type ApiGETBraintreePaymentsRequest struct {
	ctx        context.Context
	ApiService *BraintreePaymentsApiService
}

func (r ApiGETBraintreePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBraintreePaymentsExecute(r)
}

/*
GETBraintreePayments List all braintree payments

List all braintree payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreePayments(ctx context.Context) ApiGETBraintreePaymentsRequest {
	return ApiGETBraintreePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) GETBraintreePaymentsExecute(r ApiGETBraintreePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments"

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

type ApiGETBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreePaymentId string
}

func (r ApiGETBraintreePaymentsBraintreePaymentIdRequest) Execute() (*BraintreePayment, *http.Response, error) {
	return r.ApiService.GETBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
GETBraintreePaymentsBraintreePaymentId Retrieve a braintree payment

Retrieve a braintree payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreePaymentId The resource's id
 @return ApiGETBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId string) ApiGETBraintreePaymentsBraintreePaymentIdRequest {
	return ApiGETBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
//  @return BraintreePayment
func (a *BraintreePaymentsApiService) GETBraintreePaymentsBraintreePaymentIdExecute(r ApiGETBraintreePaymentsBraintreePaymentIdRequest) (*BraintreePayment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BraintreePayment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterToString(r.braintreePaymentId, "")), -1)

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

type ApiPATCHBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                    context.Context
	ApiService             *BraintreePaymentsApiService
	braintreePaymentUpdate *BraintreePaymentUpdate
	braintreePaymentId     string
}

func (r ApiPATCHBraintreePaymentsBraintreePaymentIdRequest) BraintreePaymentUpdate(braintreePaymentUpdate BraintreePaymentUpdate) ApiPATCHBraintreePaymentsBraintreePaymentIdRequest {
	r.braintreePaymentUpdate = &braintreePaymentUpdate
	return r
}

func (r ApiPATCHBraintreePaymentsBraintreePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
PATCHBraintreePaymentsBraintreePaymentId Update a braintree payment

Update a braintree payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreePaymentId The resource's id
 @return ApiPATCHBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) PATCHBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId string) ApiPATCHBraintreePaymentsBraintreePaymentIdRequest {
	return ApiPATCHBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) PATCHBraintreePaymentsBraintreePaymentIdExecute(r ApiPATCHBraintreePaymentsBraintreePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.PATCHBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterToString(r.braintreePaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreePaymentUpdate == nil {
		return nil, reportError("braintreePaymentUpdate is required and must be specified")
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
	localVarPostBody = r.braintreePaymentUpdate
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

type ApiPOSTBraintreePaymentsRequest struct {
	ctx                    context.Context
	ApiService             *BraintreePaymentsApiService
	braintreePaymentCreate *BraintreePaymentCreate
}

func (r ApiPOSTBraintreePaymentsRequest) BraintreePaymentCreate(braintreePaymentCreate BraintreePaymentCreate) ApiPOSTBraintreePaymentsRequest {
	r.braintreePaymentCreate = &braintreePaymentCreate
	return r
}

func (r ApiPOSTBraintreePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTBraintreePaymentsExecute(r)
}

/*
POSTBraintreePayments Create a braintree payment

Create a braintree payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) POSTBraintreePayments(ctx context.Context) ApiPOSTBraintreePaymentsRequest {
	return ApiPOSTBraintreePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) POSTBraintreePaymentsExecute(r ApiPOSTBraintreePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.POSTBraintreePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreePaymentCreate == nil {
		return nil, reportError("braintreePaymentCreate is required and must be specified")
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
	localVarPostBody = r.braintreePaymentCreate
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

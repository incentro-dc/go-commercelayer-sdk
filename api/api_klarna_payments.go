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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// KlarnaPaymentsApiService KlarnaPaymentsApi service
type KlarnaPaymentsApiService service

type KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest struct {
	ctx             context.Context
	ApiService      *KlarnaPaymentsApiService
	klarnaPaymentId interface{}
}

func (r KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEKlarnaPaymentsKlarnaPaymentIdExecute(r)
}

/*
DELETEKlarnaPaymentsKlarnaPaymentId Delete a klarna payment

Delete a klarna payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param klarnaPaymentId The resource's id
	@return KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest
*/
func (a *KlarnaPaymentsApiService) DELETEKlarnaPaymentsKlarnaPaymentId(ctx context.Context, klarnaPaymentId interface{}) KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest {
	return KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaPaymentId: klarnaPaymentId,
	}
}

// Execute executes the request
func (a *KlarnaPaymentsApiService) DELETEKlarnaPaymentsKlarnaPaymentIdExecute(r KlarnaPaymentsApiDELETEKlarnaPaymentsKlarnaPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.DELETEKlarnaPaymentsKlarnaPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments/{klarnaPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaPaymentId"+"}", url.PathEscape(parameterValueToString(r.klarnaPaymentId, "klarnaPaymentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest struct {
	ctx             context.Context
	ApiService      *KlarnaPaymentsApiService
	klarnaGatewayId interface{}
}

func (r KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETKlarnaGatewayIdKlarnaPaymentsExecute(r)
}

/*
GETKlarnaGatewayIdKlarnaPayments Retrieve the klarna payments associated to the klarna gateway

Retrieve the klarna payments associated to the klarna gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param klarnaGatewayId The resource's id
	@return KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest
*/
func (a *KlarnaPaymentsApiService) GETKlarnaGatewayIdKlarnaPayments(ctx context.Context, klarnaGatewayId interface{}) KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest {
	return KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaGatewayId: klarnaGatewayId,
	}
}

// Execute executes the request
func (a *KlarnaPaymentsApiService) GETKlarnaGatewayIdKlarnaPaymentsExecute(r KlarnaPaymentsApiGETKlarnaGatewayIdKlarnaPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.GETKlarnaGatewayIdKlarnaPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways/{klarnaGatewayId}/klarna_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaGatewayId"+"}", url.PathEscape(parameterValueToString(r.klarnaGatewayId, "klarnaGatewayId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type KlarnaPaymentsApiGETKlarnaPaymentsRequest struct {
	ctx        context.Context
	ApiService *KlarnaPaymentsApiService
}

func (r KlarnaPaymentsApiGETKlarnaPaymentsRequest) Execute() (*GETKlarnaPayments200Response, *http.Response, error) {
	return r.ApiService.GETKlarnaPaymentsExecute(r)
}

/*
GETKlarnaPayments List all klarna payments

List all klarna payments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return KlarnaPaymentsApiGETKlarnaPaymentsRequest
*/
func (a *KlarnaPaymentsApiService) GETKlarnaPayments(ctx context.Context) KlarnaPaymentsApiGETKlarnaPaymentsRequest {
	return KlarnaPaymentsApiGETKlarnaPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETKlarnaPayments200Response
func (a *KlarnaPaymentsApiService) GETKlarnaPaymentsExecute(r KlarnaPaymentsApiGETKlarnaPaymentsRequest) (*GETKlarnaPayments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETKlarnaPayments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.GETKlarnaPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest struct {
	ctx             context.Context
	ApiService      *KlarnaPaymentsApiService
	klarnaPaymentId interface{}
}

func (r KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest) Execute() (*GETKlarnaPaymentsKlarnaPaymentId200Response, *http.Response, error) {
	return r.ApiService.GETKlarnaPaymentsKlarnaPaymentIdExecute(r)
}

/*
GETKlarnaPaymentsKlarnaPaymentId Retrieve a klarna payment

Retrieve a klarna payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param klarnaPaymentId The resource's id
	@return KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest
*/
func (a *KlarnaPaymentsApiService) GETKlarnaPaymentsKlarnaPaymentId(ctx context.Context, klarnaPaymentId interface{}) KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest {
	return KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaPaymentId: klarnaPaymentId,
	}
}

// Execute executes the request
//
//	@return GETKlarnaPaymentsKlarnaPaymentId200Response
func (a *KlarnaPaymentsApiService) GETKlarnaPaymentsKlarnaPaymentIdExecute(r KlarnaPaymentsApiGETKlarnaPaymentsKlarnaPaymentIdRequest) (*GETKlarnaPaymentsKlarnaPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETKlarnaPaymentsKlarnaPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.GETKlarnaPaymentsKlarnaPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments/{klarnaPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaPaymentId"+"}", url.PathEscape(parameterValueToString(r.klarnaPaymentId, "klarnaPaymentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest struct {
	ctx                 context.Context
	ApiService          *KlarnaPaymentsApiService
	klarnaPaymentUpdate *KlarnaPaymentUpdate
	klarnaPaymentId     interface{}
}

func (r KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest) KlarnaPaymentUpdate(klarnaPaymentUpdate KlarnaPaymentUpdate) KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest {
	r.klarnaPaymentUpdate = &klarnaPaymentUpdate
	return r
}

func (r KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest) Execute() (*PATCHKlarnaPaymentsKlarnaPaymentId200Response, *http.Response, error) {
	return r.ApiService.PATCHKlarnaPaymentsKlarnaPaymentIdExecute(r)
}

/*
PATCHKlarnaPaymentsKlarnaPaymentId Update a klarna payment

Update a klarna payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param klarnaPaymentId The resource's id
	@return KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest
*/
func (a *KlarnaPaymentsApiService) PATCHKlarnaPaymentsKlarnaPaymentId(ctx context.Context, klarnaPaymentId interface{}) KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest {
	return KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaPaymentId: klarnaPaymentId,
	}
}

// Execute executes the request
//
//	@return PATCHKlarnaPaymentsKlarnaPaymentId200Response
func (a *KlarnaPaymentsApiService) PATCHKlarnaPaymentsKlarnaPaymentIdExecute(r KlarnaPaymentsApiPATCHKlarnaPaymentsKlarnaPaymentIdRequest) (*PATCHKlarnaPaymentsKlarnaPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHKlarnaPaymentsKlarnaPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.PATCHKlarnaPaymentsKlarnaPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments/{klarnaPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaPaymentId"+"}", url.PathEscape(parameterValueToString(r.klarnaPaymentId, "klarnaPaymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.klarnaPaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("klarnaPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.klarnaPaymentUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type KlarnaPaymentsApiPOSTKlarnaPaymentsRequest struct {
	ctx                 context.Context
	ApiService          *KlarnaPaymentsApiService
	klarnaPaymentCreate *KlarnaPaymentCreate
}

func (r KlarnaPaymentsApiPOSTKlarnaPaymentsRequest) KlarnaPaymentCreate(klarnaPaymentCreate KlarnaPaymentCreate) KlarnaPaymentsApiPOSTKlarnaPaymentsRequest {
	r.klarnaPaymentCreate = &klarnaPaymentCreate
	return r
}

func (r KlarnaPaymentsApiPOSTKlarnaPaymentsRequest) Execute() (*POSTKlarnaPayments201Response, *http.Response, error) {
	return r.ApiService.POSTKlarnaPaymentsExecute(r)
}

/*
POSTKlarnaPayments Create a klarna payment

Create a klarna payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return KlarnaPaymentsApiPOSTKlarnaPaymentsRequest
*/
func (a *KlarnaPaymentsApiService) POSTKlarnaPayments(ctx context.Context) KlarnaPaymentsApiPOSTKlarnaPaymentsRequest {
	return KlarnaPaymentsApiPOSTKlarnaPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTKlarnaPayments201Response
func (a *KlarnaPaymentsApiService) POSTKlarnaPaymentsExecute(r KlarnaPaymentsApiPOSTKlarnaPaymentsRequest) (*POSTKlarnaPayments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTKlarnaPayments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaPaymentsApiService.POSTKlarnaPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.klarnaPaymentCreate == nil {
		return localVarReturnValue, nil, reportError("klarnaPaymentCreate is required and must be specified")
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
	localVarPostBody = r.klarnaPaymentCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

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

// ExternalPaymentsApiService ExternalPaymentsApi service
type ExternalPaymentsApiService service

type ApiDELETEExternalPaymentsExternalPaymentIdRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalPaymentId string
}

func (r ApiDELETEExternalPaymentsExternalPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEExternalPaymentsExternalPaymentIdExecute(r)
}

/*
DELETEExternalPaymentsExternalPaymentId Delete an external payment

Delete an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ApiDELETEExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) DELETEExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ApiDELETEExternalPaymentsExternalPaymentIdRequest {
	return ApiDELETEExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) DELETEExternalPaymentsExternalPaymentIdExecute(r ApiDELETEExternalPaymentsExternalPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.DELETEExternalPaymentsExternalPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

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

type ApiGETExternalGatewayIdExternalPaymentsRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalGatewayId string
}

func (r ApiGETExternalGatewayIdExternalPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalGatewayIdExternalPaymentsExecute(r)
}

/*
GETExternalGatewayIdExternalPayments Retrieve the external payments associated to the external gateway

Retrieve the external payments associated to the external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGatewayId The resource's id
 @return ApiGETExternalGatewayIdExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) GETExternalGatewayIdExternalPayments(ctx context.Context, externalGatewayId string) ApiGETExternalGatewayIdExternalPaymentsRequest {
	return ApiGETExternalGatewayIdExternalPaymentsRequest{
		ApiService:        a,
		ctx:               ctx,
		externalGatewayId: externalGatewayId,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) GETExternalGatewayIdExternalPaymentsExecute(r ApiGETExternalGatewayIdExternalPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalGatewayIdExternalPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways/{externalGatewayId}/external_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGatewayId"+"}", url.PathEscape(parameterToString(r.externalGatewayId, "")), -1)

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

type ApiGETExternalPaymentsRequest struct {
	ctx        context.Context
	ApiService *ExternalPaymentsApiService
}

func (r ApiGETExternalPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalPaymentsExecute(r)
}

/*
GETExternalPayments List all external payments

List all external payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) GETExternalPayments(ctx context.Context) ApiGETExternalPaymentsRequest {
	return ApiGETExternalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) GETExternalPaymentsExecute(r ApiGETExternalPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments"

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

type ApiGETExternalPaymentsExternalPaymentIdRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalPaymentId string
}

func (r ApiGETExternalPaymentsExternalPaymentIdRequest) Execute() (*ExternalPayment, *http.Response, error) {
	return r.ApiService.GETExternalPaymentsExternalPaymentIdExecute(r)
}

/*
GETExternalPaymentsExternalPaymentId Retrieve an external payment

Retrieve an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ApiGETExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) GETExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ApiGETExternalPaymentsExternalPaymentIdRequest {
	return ApiGETExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
//  @return ExternalPayment
func (a *ExternalPaymentsApiService) GETExternalPaymentsExternalPaymentIdExecute(r ApiGETExternalPaymentsExternalPaymentIdRequest) (*ExternalPayment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPayment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalPaymentsExternalPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

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

type ApiPATCHExternalPaymentsExternalPaymentIdRequest struct {
	ctx                   context.Context
	ApiService            *ExternalPaymentsApiService
	externalPaymentUpdate *ExternalPaymentUpdate
	externalPaymentId     string
}

func (r ApiPATCHExternalPaymentsExternalPaymentIdRequest) ExternalPaymentUpdate(externalPaymentUpdate ExternalPaymentUpdate) ApiPATCHExternalPaymentsExternalPaymentIdRequest {
	r.externalPaymentUpdate = &externalPaymentUpdate
	return r
}

func (r ApiPATCHExternalPaymentsExternalPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHExternalPaymentsExternalPaymentIdExecute(r)
}

/*
PATCHExternalPaymentsExternalPaymentId Update an external payment

Update an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ApiPATCHExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) PATCHExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ApiPATCHExternalPaymentsExternalPaymentIdRequest {
	return ApiPATCHExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) PATCHExternalPaymentsExternalPaymentIdExecute(r ApiPATCHExternalPaymentsExternalPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.PATCHExternalPaymentsExternalPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPaymentUpdate == nil {
		return nil, reportError("externalPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.externalPaymentUpdate
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

type ApiPOSTExternalPaymentsRequest struct {
	ctx                   context.Context
	ApiService            *ExternalPaymentsApiService
	externalPaymentCreate *ExternalPaymentCreate
}

func (r ApiPOSTExternalPaymentsRequest) ExternalPaymentCreate(externalPaymentCreate ExternalPaymentCreate) ApiPOSTExternalPaymentsRequest {
	r.externalPaymentCreate = &externalPaymentCreate
	return r
}

func (r ApiPOSTExternalPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTExternalPaymentsExecute(r)
}

/*
POSTExternalPayments Create an external payment

Create an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) POSTExternalPayments(ctx context.Context) ApiPOSTExternalPaymentsRequest {
	return ApiPOSTExternalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) POSTExternalPaymentsExecute(r ApiPOSTExternalPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.POSTExternalPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPaymentCreate == nil {
		return nil, reportError("externalPaymentCreate is required and must be specified")
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
	localVarPostBody = r.externalPaymentCreate
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

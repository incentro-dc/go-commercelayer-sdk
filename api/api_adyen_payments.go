/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
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

type AdyenPaymentsApi interface {

	/*
		DELETEAdyenPaymentsAdyenPaymentId Delete an adyen payment

		Delete an adyen payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param adyenPaymentId The resource's id
		@return ApiDELETEAdyenPaymentsAdyenPaymentIdRequest
	*/
	DELETEAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiDELETEAdyenPaymentsAdyenPaymentIdRequest

	// DELETEAdyenPaymentsAdyenPaymentIdExecute executes the request
	DELETEAdyenPaymentsAdyenPaymentIdExecute(r ApiDELETEAdyenPaymentsAdyenPaymentIdRequest) (*http.Response, error)

	/*
		GETAdyenGatewayIdAdyenPayments Retrieve the adyen payments associated to the adyen gateway

		Retrieve the adyen payments associated to the adyen gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param adyenGatewayId The resource's id
		@return ApiGETAdyenGatewayIdAdyenPaymentsRequest
	*/
	GETAdyenGatewayIdAdyenPayments(ctx context.Context, adyenGatewayId string) ApiGETAdyenGatewayIdAdyenPaymentsRequest

	// GETAdyenGatewayIdAdyenPaymentsExecute executes the request
	GETAdyenGatewayIdAdyenPaymentsExecute(r ApiGETAdyenGatewayIdAdyenPaymentsRequest) (*http.Response, error)

	/*
		GETAdyenPayments List all adyen payments

		List all adyen payments

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETAdyenPaymentsRequest
	*/
	GETAdyenPayments(ctx context.Context) ApiGETAdyenPaymentsRequest

	// GETAdyenPaymentsExecute executes the request
	GETAdyenPaymentsExecute(r ApiGETAdyenPaymentsRequest) (*http.Response, error)

	/*
		GETAdyenPaymentsAdyenPaymentId Retrieve an adyen payment

		Retrieve an adyen payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param adyenPaymentId The resource's id
		@return ApiGETAdyenPaymentsAdyenPaymentIdRequest
	*/
	GETAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiGETAdyenPaymentsAdyenPaymentIdRequest

	// GETAdyenPaymentsAdyenPaymentIdExecute executes the request
	//  @return AdyenPayment
	GETAdyenPaymentsAdyenPaymentIdExecute(r ApiGETAdyenPaymentsAdyenPaymentIdRequest) (*AdyenPayment, *http.Response, error)

	/*
		PATCHAdyenPaymentsAdyenPaymentId Update an adyen payment

		Update an adyen payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param adyenPaymentId The resource's id
		@return ApiPATCHAdyenPaymentsAdyenPaymentIdRequest
	*/
	PATCHAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiPATCHAdyenPaymentsAdyenPaymentIdRequest

	// PATCHAdyenPaymentsAdyenPaymentIdExecute executes the request
	PATCHAdyenPaymentsAdyenPaymentIdExecute(r ApiPATCHAdyenPaymentsAdyenPaymentIdRequest) (*http.Response, error)

	/*
		POSTAdyenPayments Create an adyen payment

		Create an adyen payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTAdyenPaymentsRequest
	*/
	POSTAdyenPayments(ctx context.Context) ApiPOSTAdyenPaymentsRequest

	// POSTAdyenPaymentsExecute executes the request
	POSTAdyenPaymentsExecute(r ApiPOSTAdyenPaymentsRequest) (*http.Response, error)
}

// AdyenPaymentsApiService AdyenPaymentsApi service
type AdyenPaymentsApiService service

type ApiDELETEAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx            context.Context
	ApiService     AdyenPaymentsApi
	adyenPaymentId string
}

func (r ApiDELETEAdyenPaymentsAdyenPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
DELETEAdyenPaymentsAdyenPaymentId Delete an adyen payment

Delete an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return ApiDELETEAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) DELETEAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiDELETEAdyenPaymentsAdyenPaymentIdRequest {
	return ApiDELETEAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) DELETEAdyenPaymentsAdyenPaymentIdExecute(r ApiDELETEAdyenPaymentsAdyenPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.DELETEAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

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

type ApiGETAdyenGatewayIdAdyenPaymentsRequest struct {
	ctx            context.Context
	ApiService     AdyenPaymentsApi
	adyenGatewayId string
}

func (r ApiGETAdyenGatewayIdAdyenPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAdyenGatewayIdAdyenPaymentsExecute(r)
}

/*
GETAdyenGatewayIdAdyenPayments Retrieve the adyen payments associated to the adyen gateway

Retrieve the adyen payments associated to the adyen gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenGatewayId The resource's id
 @return ApiGETAdyenGatewayIdAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenGatewayIdAdyenPayments(ctx context.Context, adyenGatewayId string) ApiGETAdyenGatewayIdAdyenPaymentsRequest {
	return ApiGETAdyenGatewayIdAdyenPaymentsRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenGatewayId: adyenGatewayId,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) GETAdyenGatewayIdAdyenPaymentsExecute(r ApiGETAdyenGatewayIdAdyenPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenGatewayIdAdyenPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways/{adyenGatewayId}/adyen_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenGatewayId"+"}", url.PathEscape(parameterToString(r.adyenGatewayId, "")), -1)

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

type ApiGETAdyenPaymentsRequest struct {
	ctx        context.Context
	ApiService AdyenPaymentsApi
}

func (r ApiGETAdyenPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAdyenPaymentsExecute(r)
}

/*
GETAdyenPayments List all adyen payments

List all adyen payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenPayments(ctx context.Context) ApiGETAdyenPaymentsRequest {
	return ApiGETAdyenPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) GETAdyenPaymentsExecute(r ApiGETAdyenPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments"

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

type ApiGETAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx            context.Context
	ApiService     AdyenPaymentsApi
	adyenPaymentId string
}

func (r ApiGETAdyenPaymentsAdyenPaymentIdRequest) Execute() (*AdyenPayment, *http.Response, error) {
	return r.ApiService.GETAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
GETAdyenPaymentsAdyenPaymentId Retrieve an adyen payment

Retrieve an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return ApiGETAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiGETAdyenPaymentsAdyenPaymentIdRequest {
	return ApiGETAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
//  @return AdyenPayment
func (a *AdyenPaymentsApiService) GETAdyenPaymentsAdyenPaymentIdExecute(r ApiGETAdyenPaymentsAdyenPaymentIdRequest) (*AdyenPayment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdyenPayment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

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

type ApiPATCHAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx                context.Context
	ApiService         AdyenPaymentsApi
	adyenPaymentId     string
	adyenPaymentUpdate *AdyenPaymentUpdate
}

func (r ApiPATCHAdyenPaymentsAdyenPaymentIdRequest) AdyenPaymentUpdate(adyenPaymentUpdate AdyenPaymentUpdate) ApiPATCHAdyenPaymentsAdyenPaymentIdRequest {
	r.adyenPaymentUpdate = &adyenPaymentUpdate
	return r
}

func (r ApiPATCHAdyenPaymentsAdyenPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
PATCHAdyenPaymentsAdyenPaymentId Update an adyen payment

Update an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return ApiPATCHAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) PATCHAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) ApiPATCHAdyenPaymentsAdyenPaymentIdRequest {
	return ApiPATCHAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) PATCHAdyenPaymentsAdyenPaymentIdExecute(r ApiPATCHAdyenPaymentsAdyenPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.PATCHAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenPaymentUpdate == nil {
		return nil, reportError("adyenPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.adyenPaymentUpdate
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

type ApiPOSTAdyenPaymentsRequest struct {
	ctx                context.Context
	ApiService         AdyenPaymentsApi
	adyenPaymentCreate *AdyenPaymentCreate
}

func (r ApiPOSTAdyenPaymentsRequest) AdyenPaymentCreate(adyenPaymentCreate AdyenPaymentCreate) ApiPOSTAdyenPaymentsRequest {
	r.adyenPaymentCreate = &adyenPaymentCreate
	return r
}

func (r ApiPOSTAdyenPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTAdyenPaymentsExecute(r)
}

/*
POSTAdyenPayments Create an adyen payment

Create an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) POSTAdyenPayments(ctx context.Context) ApiPOSTAdyenPaymentsRequest {
	return ApiPOSTAdyenPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) POSTAdyenPaymentsExecute(r ApiPOSTAdyenPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.POSTAdyenPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenPaymentCreate == nil {
		return nil, reportError("adyenPaymentCreate is required and must be specified")
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
	localVarPostBody = r.adyenPaymentCreate
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

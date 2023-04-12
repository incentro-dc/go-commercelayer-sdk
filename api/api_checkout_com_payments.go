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

// CheckoutComPaymentsApiService CheckoutComPaymentsApi service
type CheckoutComPaymentsApiService service

type CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                  context.Context
	ApiService           *CheckoutComPaymentsApiService
	checkoutComPaymentId interface{}
}

func (r CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
DELETECheckoutComPaymentsCheckoutComPaymentId Delete a checkout.com payment

Delete a checkout.com payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param checkoutComPaymentId The resource's id
	@return CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) DELETECheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId interface{}) CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest {
	return CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) DELETECheckoutComPaymentsCheckoutComPaymentIdExecute(r CheckoutComPaymentsApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.DELETECheckoutComPaymentsCheckoutComPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterValueToString(r.checkoutComPaymentId, "checkoutComPaymentId")), -1)

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

type CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest struct {
	ctx                  context.Context
	ApiService           *CheckoutComPaymentsApiService
	checkoutComGatewayId interface{}
}

func (r CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCheckoutComGatewayIdCheckoutComPaymentsExecute(r)
}

/*
GETCheckoutComGatewayIdCheckoutComPayments Retrieve the checkout com payments associated to the checkout.com gateway

Retrieve the checkout com payments associated to the checkout.com gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param checkoutComGatewayId The resource's id
	@return CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComGatewayIdCheckoutComPayments(ctx context.Context, checkoutComGatewayId interface{}) CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest {
	return CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) GETCheckoutComGatewayIdCheckoutComPaymentsExecute(r CheckoutComPaymentsApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.GETCheckoutComGatewayIdCheckoutComPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}/checkout_com_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterValueToString(r.checkoutComGatewayId, "checkoutComGatewayId")), -1)

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

type CheckoutComPaymentsApiGETCheckoutComPaymentsRequest struct {
	ctx        context.Context
	ApiService *CheckoutComPaymentsApiService
}

func (r CheckoutComPaymentsApiGETCheckoutComPaymentsRequest) Execute() (*GETCheckoutComPayments200Response, *http.Response, error) {
	return r.ApiService.GETCheckoutComPaymentsExecute(r)
}

/*
GETCheckoutComPayments List all checkout.com payments

List all checkout.com payments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CheckoutComPaymentsApiGETCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComPayments(ctx context.Context) CheckoutComPaymentsApiGETCheckoutComPaymentsRequest {
	return CheckoutComPaymentsApiGETCheckoutComPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCheckoutComPayments200Response
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsExecute(r CheckoutComPaymentsApiGETCheckoutComPaymentsRequest) (*GETCheckoutComPayments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCheckoutComPayments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.GETCheckoutComPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments"

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

type CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                  context.Context
	ApiService           *CheckoutComPaymentsApiService
	checkoutComPaymentId interface{}
}

func (r CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*GETCheckoutComPaymentsCheckoutComPaymentId200Response, *http.Response, error) {
	return r.ApiService.GETCheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
GETCheckoutComPaymentsCheckoutComPaymentId Retrieve a checkout.com payment

Retrieve a checkout.com payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param checkoutComPaymentId The resource's id
	@return CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId interface{}) CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest {
	return CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
//
//	@return GETCheckoutComPaymentsCheckoutComPaymentId200Response
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsCheckoutComPaymentIdExecute(r CheckoutComPaymentsApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest) (*GETCheckoutComPaymentsCheckoutComPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCheckoutComPaymentsCheckoutComPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.GETCheckoutComPaymentsCheckoutComPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterValueToString(r.checkoutComPaymentId, "checkoutComPaymentId")), -1)

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

type CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                      context.Context
	ApiService               *CheckoutComPaymentsApiService
	checkoutComPaymentUpdate *CheckoutComPaymentUpdate
	checkoutComPaymentId     interface{}
}

func (r CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) CheckoutComPaymentUpdate(checkoutComPaymentUpdate CheckoutComPaymentUpdate) CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest {
	r.checkoutComPaymentUpdate = &checkoutComPaymentUpdate
	return r
}

func (r CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*PATCHCheckoutComPaymentsCheckoutComPaymentId200Response, *http.Response, error) {
	return r.ApiService.PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
PATCHCheckoutComPaymentsCheckoutComPaymentId Update a checkout.com payment

Update a checkout.com payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param checkoutComPaymentId The resource's id
	@return CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) PATCHCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId interface{}) CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest {
	return CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
//
//	@return PATCHCheckoutComPaymentsCheckoutComPaymentId200Response
func (a *CheckoutComPaymentsApiService) PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute(r CheckoutComPaymentsApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) (*PATCHCheckoutComPaymentsCheckoutComPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.PATCHCheckoutComPaymentsCheckoutComPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterValueToString(r.checkoutComPaymentId, "checkoutComPaymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComPaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("checkoutComPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.checkoutComPaymentUpdate
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

type CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest struct {
	ctx                      context.Context
	ApiService               *CheckoutComPaymentsApiService
	checkoutComPaymentCreate *CheckoutComPaymentCreate
}

func (r CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest) CheckoutComPaymentCreate(checkoutComPaymentCreate CheckoutComPaymentCreate) CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest {
	r.checkoutComPaymentCreate = &checkoutComPaymentCreate
	return r
}

func (r CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest) Execute() (*POSTCheckoutComPayments201Response, *http.Response, error) {
	return r.ApiService.POSTCheckoutComPaymentsExecute(r)
}

/*
POSTCheckoutComPayments Create a checkout.com payment

Create a checkout.com payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) POSTCheckoutComPayments(ctx context.Context) CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest {
	return CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTCheckoutComPayments201Response
func (a *CheckoutComPaymentsApiService) POSTCheckoutComPaymentsExecute(r CheckoutComPaymentsApiPOSTCheckoutComPaymentsRequest) (*POSTCheckoutComPayments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCheckoutComPayments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.POSTCheckoutComPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComPaymentCreate == nil {
		return localVarReturnValue, nil, reportError("checkoutComPaymentCreate is required and must be specified")
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
	localVarPostBody = r.checkoutComPaymentCreate
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

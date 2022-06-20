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

type CheckoutComPaymentsApi interface {

	/*
		DELETECheckoutComPaymentsCheckoutComPaymentId Delete a checkout.com payment

		Delete a checkout.com payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComPaymentId The resource's id
		@return ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest
	*/
	DELETECheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest

	// DELETECheckoutComPaymentsCheckoutComPaymentIdExecute executes the request
	DELETECheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest) (*http.Response, error)

	/*
		GETCheckoutComGatewayIdCheckoutComPayments Retrieve the checkout com payments associated to the checkout.com gateway

		Retrieve the checkout com payments associated to the checkout.com gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComGatewayId The resource's id
		@return ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest
	*/
	GETCheckoutComGatewayIdCheckoutComPayments(ctx context.Context, checkoutComGatewayId string) ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest

	// GETCheckoutComGatewayIdCheckoutComPaymentsExecute executes the request
	GETCheckoutComGatewayIdCheckoutComPaymentsExecute(r ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest) (*http.Response, error)

	/*
		GETCheckoutComPayments List all checkout.com payments

		List all checkout.com payments

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCheckoutComPaymentsRequest
	*/
	GETCheckoutComPayments(ctx context.Context) ApiGETCheckoutComPaymentsRequest

	// GETCheckoutComPaymentsExecute executes the request
	GETCheckoutComPaymentsExecute(r ApiGETCheckoutComPaymentsRequest) (*http.Response, error)

	/*
		GETCheckoutComPaymentsCheckoutComPaymentId Retrieve a checkout.com payment

		Retrieve a checkout.com payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComPaymentId The resource's id
		@return ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest
	*/
	GETCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest

	// GETCheckoutComPaymentsCheckoutComPaymentIdExecute executes the request
	//  @return CheckoutComPayment
	GETCheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest) (*CheckoutComPayment, *http.Response, error)

	/*
		PATCHCheckoutComPaymentsCheckoutComPaymentId Update a checkout.com payment

		Update a checkout.com payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComPaymentId The resource's id
		@return ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest
	*/
	PATCHCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest

	// PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute executes the request
	PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) (*http.Response, error)

	/*
		POSTCheckoutComPayments Create a checkout.com payment

		Create a checkout.com payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTCheckoutComPaymentsRequest
	*/
	POSTCheckoutComPayments(ctx context.Context) ApiPOSTCheckoutComPaymentsRequest

	// POSTCheckoutComPaymentsExecute executes the request
	POSTCheckoutComPaymentsExecute(r ApiPOSTCheckoutComPaymentsRequest) (*http.Response, error)
}

// CheckoutComPaymentsApiService CheckoutComPaymentsApi service
type CheckoutComPaymentsApiService service

type ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                  context.Context
	ApiService           CheckoutComPaymentsApi
	checkoutComPaymentId string
}

func (r ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
DELETECheckoutComPaymentsCheckoutComPaymentId Delete a checkout.com payment

Delete a checkout.com payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComPaymentId The resource's id
 @return ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) DELETECheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest {
	return ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) DELETECheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiDELETECheckoutComPaymentsCheckoutComPaymentIdRequest) (*http.Response, error) {
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
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterToString(r.checkoutComPaymentId, "")), -1)

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

type ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest struct {
	ctx                  context.Context
	ApiService           CheckoutComPaymentsApi
	checkoutComGatewayId string
}

func (r ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCheckoutComGatewayIdCheckoutComPaymentsExecute(r)
}

/*
GETCheckoutComGatewayIdCheckoutComPayments Retrieve the checkout com payments associated to the checkout.com gateway

Retrieve the checkout com payments associated to the checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComGatewayIdCheckoutComPayments(ctx context.Context, checkoutComGatewayId string) ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest {
	return ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) GETCheckoutComGatewayIdCheckoutComPaymentsExecute(r ApiGETCheckoutComGatewayIdCheckoutComPaymentsRequest) (*http.Response, error) {
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
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

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

type ApiGETCheckoutComPaymentsRequest struct {
	ctx        context.Context
	ApiService CheckoutComPaymentsApi
}

func (r ApiGETCheckoutComPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCheckoutComPaymentsExecute(r)
}

/*
GETCheckoutComPayments List all checkout.com payments

List all checkout.com payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComPayments(ctx context.Context) ApiGETCheckoutComPaymentsRequest {
	return ApiGETCheckoutComPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsExecute(r ApiGETCheckoutComPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.GETCheckoutComPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
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

type ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                  context.Context
	ApiService           CheckoutComPaymentsApi
	checkoutComPaymentId string
}

func (r ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*CheckoutComPayment, *http.Response, error) {
	return r.ApiService.GETCheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
GETCheckoutComPaymentsCheckoutComPaymentId Retrieve a checkout.com payment

Retrieve a checkout.com payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComPaymentId The resource's id
 @return ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest {
	return ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
//  @return CheckoutComPayment
func (a *CheckoutComPaymentsApiService) GETCheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiGETCheckoutComPaymentsCheckoutComPaymentIdRequest) (*CheckoutComPayment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CheckoutComPayment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.GETCheckoutComPaymentsCheckoutComPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterToString(r.checkoutComPaymentId, "")), -1)

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

type ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest struct {
	ctx                      context.Context
	ApiService               CheckoutComPaymentsApi
	checkoutComPaymentId     string
	checkoutComPaymentUpdate *CheckoutComPaymentUpdate
}

func (r ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) CheckoutComPaymentUpdate(checkoutComPaymentUpdate CheckoutComPaymentUpdate) ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest {
	r.checkoutComPaymentUpdate = &checkoutComPaymentUpdate
	return r
}

func (r ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute(r)
}

/*
PATCHCheckoutComPaymentsCheckoutComPaymentId Update a checkout.com payment

Update a checkout.com payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComPaymentId The resource's id
 @return ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest
*/
func (a *CheckoutComPaymentsApiService) PATCHCheckoutComPaymentsCheckoutComPaymentId(ctx context.Context, checkoutComPaymentId string) ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest {
	return ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) PATCHCheckoutComPaymentsCheckoutComPaymentIdExecute(r ApiPATCHCheckoutComPaymentsCheckoutComPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.PATCHCheckoutComPaymentsCheckoutComPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterToString(r.checkoutComPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComPaymentUpdate == nil {
		return nil, reportError("checkoutComPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.checkoutComPaymentUpdate
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

type ApiPOSTCheckoutComPaymentsRequest struct {
	ctx                      context.Context
	ApiService               CheckoutComPaymentsApi
	checkoutComPaymentCreate *CheckoutComPaymentCreate
}

func (r ApiPOSTCheckoutComPaymentsRequest) CheckoutComPaymentCreate(checkoutComPaymentCreate CheckoutComPaymentCreate) ApiPOSTCheckoutComPaymentsRequest {
	r.checkoutComPaymentCreate = &checkoutComPaymentCreate
	return r
}

func (r ApiPOSTCheckoutComPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCheckoutComPaymentsExecute(r)
}

/*
POSTCheckoutComPayments Create a checkout.com payment

Create a checkout.com payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCheckoutComPaymentsRequest
*/
func (a *CheckoutComPaymentsApiService) POSTCheckoutComPayments(ctx context.Context) ApiPOSTCheckoutComPaymentsRequest {
	return ApiPOSTCheckoutComPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CheckoutComPaymentsApiService) POSTCheckoutComPaymentsExecute(r ApiPOSTCheckoutComPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComPaymentsApiService.POSTCheckoutComPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComPaymentCreate == nil {
		return nil, reportError("checkoutComPaymentCreate is required and must be specified")
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
	localVarPostBody = r.checkoutComPaymentCreate
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

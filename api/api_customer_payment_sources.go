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

type CustomerPaymentSourcesApi interface {

	/*
		DELETECustomerPaymentSourcesCustomerPaymentSourceId Delete a customer payment source

		Delete a customer payment source

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPaymentSourceId The resource's id
		@return ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest
	*/
	DELETECustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest

	// DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute executes the request
	DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*http.Response, error)

	/*
		GETCustomerIdCustomerPaymentSources Retrieve the customer payment sources associated to the customer

		Retrieve the customer payment sources associated to the customer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerId The resource's id
		@return ApiGETCustomerIdCustomerPaymentSourcesRequest
	*/
	GETCustomerIdCustomerPaymentSources(ctx context.Context, customerId string) ApiGETCustomerIdCustomerPaymentSourcesRequest

	// GETCustomerIdCustomerPaymentSourcesExecute executes the request
	GETCustomerIdCustomerPaymentSourcesExecute(r ApiGETCustomerIdCustomerPaymentSourcesRequest) (*http.Response, error)

	/*
		GETCustomerPaymentSources List all customer payment sources

		List all customer payment sources

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCustomerPaymentSourcesRequest
	*/
	GETCustomerPaymentSources(ctx context.Context) ApiGETCustomerPaymentSourcesRequest

	// GETCustomerPaymentSourcesExecute executes the request
	GETCustomerPaymentSourcesExecute(r ApiGETCustomerPaymentSourcesRequest) (*http.Response, error)

	/*
		GETCustomerPaymentSourcesCustomerPaymentSourceId Retrieve a customer payment source

		Retrieve a customer payment source

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPaymentSourceId The resource's id
		@return ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest
	*/
	GETCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest

	// GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute executes the request
	//  @return CustomerPaymentSource
	GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*CustomerPaymentSource, *http.Response, error)

	/*
		GETExternalPaymentIdWallet Retrieve the wallet associated to the external payment

		Retrieve the wallet associated to the external payment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalPaymentId The resource's id
		@return ApiGETExternalPaymentIdWalletRequest
	*/
	GETExternalPaymentIdWallet(ctx context.Context, externalPaymentId string) ApiGETExternalPaymentIdWalletRequest

	// GETExternalPaymentIdWalletExecute executes the request
	GETExternalPaymentIdWalletExecute(r ApiGETExternalPaymentIdWalletRequest) (*http.Response, error)

	/*
		GETOrderIdAvailableCustomerPaymentSources Retrieve the available customer payment sources associated to the order

		Retrieve the available customer payment sources associated to the order

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orderId The resource's id
		@return ApiGETOrderIdAvailableCustomerPaymentSourcesRequest
	*/
	GETOrderIdAvailableCustomerPaymentSources(ctx context.Context, orderId string) ApiGETOrderIdAvailableCustomerPaymentSourcesRequest

	// GETOrderIdAvailableCustomerPaymentSourcesExecute executes the request
	GETOrderIdAvailableCustomerPaymentSourcesExecute(r ApiGETOrderIdAvailableCustomerPaymentSourcesRequest) (*http.Response, error)

	/*
		PATCHCustomerPaymentSourcesCustomerPaymentSourceId Update a customer payment source

		Update a customer payment source

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPaymentSourceId The resource's id
		@return ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest
	*/
	PATCHCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest

	// PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute executes the request
	PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*http.Response, error)

	/*
		POSTCustomerPaymentSources Create a customer payment source

		Create a customer payment source

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTCustomerPaymentSourcesRequest
	*/
	POSTCustomerPaymentSources(ctx context.Context) ApiPOSTCustomerPaymentSourcesRequest

	// POSTCustomerPaymentSourcesExecute executes the request
	POSTCustomerPaymentSourcesExecute(r ApiPOSTCustomerPaymentSourcesRequest) (*http.Response, error)
}

// CustomerPaymentSourcesApiService CustomerPaymentSourcesApi service
type CustomerPaymentSourcesApiService service

type ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                     context.Context
	ApiService              CustomerPaymentSourcesApi
	customerPaymentSourceId string
}

func (r ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
DELETECustomerPaymentSourcesCustomerPaymentSourceId Delete a customer payment source

Delete a customer payment source

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPaymentSourceId The resource's id
 @return ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) DELETECustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.DELETECustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

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

type ApiGETCustomerIdCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService CustomerPaymentSourcesApi
	customerId string
}

func (r ApiGETCustomerIdCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerPaymentSourcesExecute(r)
}

/*
GETCustomerIdCustomerPaymentSources Retrieve the customer payment sources associated to the customer

Retrieve the customer payment sources associated to the customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The resource's id
 @return ApiGETCustomerIdCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerIdCustomerPaymentSources(ctx context.Context, customerId string) ApiGETCustomerIdCustomerPaymentSourcesRequest {
	return ApiGETCustomerIdCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETCustomerIdCustomerPaymentSourcesExecute(r ApiGETCustomerIdCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerIdCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_payment_sources"
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

type ApiGETCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService CustomerPaymentSourcesApi
}

func (r ApiGETCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerPaymentSourcesExecute(r)
}

/*
GETCustomerPaymentSources List all customer payment sources

List all customer payment sources

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSources(ctx context.Context) ApiGETCustomerPaymentSourcesRequest {
	return ApiGETCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesExecute(r ApiGETCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources"

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

type ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                     context.Context
	ApiService              CustomerPaymentSourcesApi
	customerPaymentSourceId string
}

func (r ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*CustomerPaymentSource, *http.Response, error) {
	return r.ApiService.GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
GETCustomerPaymentSourcesCustomerPaymentSourceId Retrieve a customer payment source

Retrieve a customer payment source

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPaymentSourceId The resource's id
 @return ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
//  @return CustomerPaymentSource
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*CustomerPaymentSource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerPaymentSource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

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

type ApiGETExternalPaymentIdWalletRequest struct {
	ctx               context.Context
	ApiService        CustomerPaymentSourcesApi
	externalPaymentId string
}

func (r ApiGETExternalPaymentIdWalletRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalPaymentIdWalletExecute(r)
}

/*
GETExternalPaymentIdWallet Retrieve the wallet associated to the external payment

Retrieve the wallet associated to the external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ApiGETExternalPaymentIdWalletRequest
*/
func (a *CustomerPaymentSourcesApiService) GETExternalPaymentIdWallet(ctx context.Context, externalPaymentId string) ApiGETExternalPaymentIdWalletRequest {
	return ApiGETExternalPaymentIdWalletRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETExternalPaymentIdWalletExecute(r ApiGETExternalPaymentIdWalletRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETExternalPaymentIdWallet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}/wallet"
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

type ApiGETOrderIdAvailableCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService CustomerPaymentSourcesApi
	orderId    string
}

func (r ApiGETOrderIdAvailableCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdAvailableCustomerPaymentSourcesExecute(r)
}

/*
GETOrderIdAvailableCustomerPaymentSources Retrieve the available customer payment sources associated to the order

Retrieve the available customer payment sources associated to the order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderId The resource's id
 @return ApiGETOrderIdAvailableCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETOrderIdAvailableCustomerPaymentSources(ctx context.Context, orderId string) ApiGETOrderIdAvailableCustomerPaymentSourcesRequest {
	return ApiGETOrderIdAvailableCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETOrderIdAvailableCustomerPaymentSourcesExecute(r ApiGETOrderIdAvailableCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETOrderIdAvailableCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/available_customer_payment_sources"
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

type ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                         context.Context
	ApiService                  CustomerPaymentSourcesApi
	customerPaymentSourceId     string
	customerPaymentSourceUpdate *CustomerPaymentSourceUpdate
}

func (r ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) CustomerPaymentSourceUpdate(customerPaymentSourceUpdate CustomerPaymentSourceUpdate) ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	r.customerPaymentSourceUpdate = &customerPaymentSourceUpdate
	return r
}

func (r ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
PATCHCustomerPaymentSourcesCustomerPaymentSourceId Update a customer payment source

Update a customer payment source

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPaymentSourceId The resource's id
 @return ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) PATCHCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r ApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.PATCHCustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPaymentSourceUpdate == nil {
		return nil, reportError("customerPaymentSourceUpdate is required and must be specified")
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
	localVarPostBody = r.customerPaymentSourceUpdate
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

type ApiPOSTCustomerPaymentSourcesRequest struct {
	ctx                         context.Context
	ApiService                  CustomerPaymentSourcesApi
	customerPaymentSourceCreate *CustomerPaymentSourceCreate
}

func (r ApiPOSTCustomerPaymentSourcesRequest) CustomerPaymentSourceCreate(customerPaymentSourceCreate CustomerPaymentSourceCreate) ApiPOSTCustomerPaymentSourcesRequest {
	r.customerPaymentSourceCreate = &customerPaymentSourceCreate
	return r
}

func (r ApiPOSTCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCustomerPaymentSourcesExecute(r)
}

/*
POSTCustomerPaymentSources Create a customer payment source

Create a customer payment source

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) POSTCustomerPaymentSources(ctx context.Context) ApiPOSTCustomerPaymentSourcesRequest {
	return ApiPOSTCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) POSTCustomerPaymentSourcesExecute(r ApiPOSTCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.POSTCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPaymentSourceCreate == nil {
		return nil, reportError("customerPaymentSourceCreate is required and must be specified")
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
	localVarPostBody = r.customerPaymentSourceCreate
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

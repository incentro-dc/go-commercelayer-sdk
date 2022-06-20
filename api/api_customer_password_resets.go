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

type CustomerPasswordResetsApi interface {

	/*
		DELETECustomerPasswordResetsCustomerPasswordResetId Delete a customer password reset

		Delete a customer password reset

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPasswordResetId The resource's id
		@return ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest
	*/
	DELETECustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest

	// DELETECustomerPasswordResetsCustomerPasswordResetIdExecute executes the request
	DELETECustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest) (*http.Response, error)

	/*
		GETCustomerPasswordResets List all customer password resets

		List all customer password resets

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCustomerPasswordResetsRequest
	*/
	GETCustomerPasswordResets(ctx context.Context) ApiGETCustomerPasswordResetsRequest

	// GETCustomerPasswordResetsExecute executes the request
	GETCustomerPasswordResetsExecute(r ApiGETCustomerPasswordResetsRequest) (*http.Response, error)

	/*
		GETCustomerPasswordResetsCustomerPasswordResetId Retrieve a customer password reset

		Retrieve a customer password reset

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPasswordResetId The resource's id
		@return ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest
	*/
	GETCustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest

	// GETCustomerPasswordResetsCustomerPasswordResetIdExecute executes the request
	//  @return CustomerPasswordReset
	GETCustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest) (*CustomerPasswordReset, *http.Response, error)

	/*
		PATCHCustomerPasswordResetsCustomerPasswordResetId Update a customer password reset

		Update a customer password reset

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerPasswordResetId The resource's id
		@return ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest
	*/
	PATCHCustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest

	// PATCHCustomerPasswordResetsCustomerPasswordResetIdExecute executes the request
	PATCHCustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest) (*http.Response, error)

	/*
		POSTCustomerPasswordResets Create a customer password reset

		Create a customer password reset

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTCustomerPasswordResetsRequest
	*/
	POSTCustomerPasswordResets(ctx context.Context) ApiPOSTCustomerPasswordResetsRequest

	// POSTCustomerPasswordResetsExecute executes the request
	POSTCustomerPasswordResetsExecute(r ApiPOSTCustomerPasswordResetsRequest) (*http.Response, error)
}

// CustomerPasswordResetsApiService CustomerPasswordResetsApi service
type CustomerPasswordResetsApiService service

type ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest struct {
	ctx                     context.Context
	ApiService              CustomerPasswordResetsApi
	customerPasswordResetId string
}

func (r ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerPasswordResetsCustomerPasswordResetIdExecute(r)
}

/*
DELETECustomerPasswordResetsCustomerPasswordResetId Delete a customer password reset

Delete a customer password reset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPasswordResetId The resource's id
 @return ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest
*/
func (a *CustomerPasswordResetsApiService) DELETECustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest {
	return ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPasswordResetId: customerPasswordResetId,
	}
}

// Execute executes the request
func (a *CustomerPasswordResetsApiService) DELETECustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPasswordResetsApiService.DELETECustomerPasswordResetsCustomerPasswordResetId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_password_resets/{customerPasswordResetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPasswordResetId"+"}", url.PathEscape(parameterToString(r.customerPasswordResetId, "")), -1)

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

type ApiGETCustomerPasswordResetsRequest struct {
	ctx        context.Context
	ApiService CustomerPasswordResetsApi
}

func (r ApiGETCustomerPasswordResetsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerPasswordResetsExecute(r)
}

/*
GETCustomerPasswordResets List all customer password resets

List all customer password resets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCustomerPasswordResetsRequest
*/
func (a *CustomerPasswordResetsApiService) GETCustomerPasswordResets(ctx context.Context) ApiGETCustomerPasswordResetsRequest {
	return ApiGETCustomerPasswordResetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerPasswordResetsApiService) GETCustomerPasswordResetsExecute(r ApiGETCustomerPasswordResetsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPasswordResetsApiService.GETCustomerPasswordResets")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_password_resets"

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

type ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest struct {
	ctx                     context.Context
	ApiService              CustomerPasswordResetsApi
	customerPasswordResetId string
}

func (r ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest) Execute() (*CustomerPasswordReset, *http.Response, error) {
	return r.ApiService.GETCustomerPasswordResetsCustomerPasswordResetIdExecute(r)
}

/*
GETCustomerPasswordResetsCustomerPasswordResetId Retrieve a customer password reset

Retrieve a customer password reset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPasswordResetId The resource's id
 @return ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest
*/
func (a *CustomerPasswordResetsApiService) GETCustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest {
	return ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPasswordResetId: customerPasswordResetId,
	}
}

// Execute executes the request
//  @return CustomerPasswordReset
func (a *CustomerPasswordResetsApiService) GETCustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiGETCustomerPasswordResetsCustomerPasswordResetIdRequest) (*CustomerPasswordReset, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerPasswordReset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPasswordResetsApiService.GETCustomerPasswordResetsCustomerPasswordResetId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_password_resets/{customerPasswordResetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPasswordResetId"+"}", url.PathEscape(parameterToString(r.customerPasswordResetId, "")), -1)

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

type ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest struct {
	ctx                         context.Context
	ApiService                  CustomerPasswordResetsApi
	customerPasswordResetId     string
	customerPasswordResetUpdate *CustomerPasswordResetUpdate
}

func (r ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest) CustomerPasswordResetUpdate(customerPasswordResetUpdate CustomerPasswordResetUpdate) ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest {
	r.customerPasswordResetUpdate = &customerPasswordResetUpdate
	return r
}

func (r ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCustomerPasswordResetsCustomerPasswordResetIdExecute(r)
}

/*
PATCHCustomerPasswordResetsCustomerPasswordResetId Update a customer password reset

Update a customer password reset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerPasswordResetId The resource's id
 @return ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest
*/
func (a *CustomerPasswordResetsApiService) PATCHCustomerPasswordResetsCustomerPasswordResetId(ctx context.Context, customerPasswordResetId string) ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest {
	return ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPasswordResetId: customerPasswordResetId,
	}
}

// Execute executes the request
func (a *CustomerPasswordResetsApiService) PATCHCustomerPasswordResetsCustomerPasswordResetIdExecute(r ApiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPasswordResetsApiService.PATCHCustomerPasswordResetsCustomerPasswordResetId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_password_resets/{customerPasswordResetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPasswordResetId"+"}", url.PathEscape(parameterToString(r.customerPasswordResetId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPasswordResetUpdate == nil {
		return nil, reportError("customerPasswordResetUpdate is required and must be specified")
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
	localVarPostBody = r.customerPasswordResetUpdate
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

type ApiPOSTCustomerPasswordResetsRequest struct {
	ctx                         context.Context
	ApiService                  CustomerPasswordResetsApi
	customerPasswordResetCreate *CustomerPasswordResetCreate
}

func (r ApiPOSTCustomerPasswordResetsRequest) CustomerPasswordResetCreate(customerPasswordResetCreate CustomerPasswordResetCreate) ApiPOSTCustomerPasswordResetsRequest {
	r.customerPasswordResetCreate = &customerPasswordResetCreate
	return r
}

func (r ApiPOSTCustomerPasswordResetsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCustomerPasswordResetsExecute(r)
}

/*
POSTCustomerPasswordResets Create a customer password reset

Create a customer password reset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCustomerPasswordResetsRequest
*/
func (a *CustomerPasswordResetsApiService) POSTCustomerPasswordResets(ctx context.Context) ApiPOSTCustomerPasswordResetsRequest {
	return ApiPOSTCustomerPasswordResetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerPasswordResetsApiService) POSTCustomerPasswordResetsExecute(r ApiPOSTCustomerPasswordResetsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPasswordResetsApiService.POSTCustomerPasswordResets")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_password_resets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPasswordResetCreate == nil {
		return nil, reportError("customerPasswordResetCreate is required and must be specified")
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
	localVarPostBody = r.customerPasswordResetCreate
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

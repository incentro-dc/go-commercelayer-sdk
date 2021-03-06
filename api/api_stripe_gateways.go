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

// StripeGatewaysApiService StripeGatewaysApi service
type StripeGatewaysApiService service

type ApiDELETEStripeGatewaysStripeGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *StripeGatewaysApiService
	stripeGatewayId string
}

func (r ApiDELETEStripeGatewaysStripeGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStripeGatewaysStripeGatewayIdExecute(r)
}

/*
DELETEStripeGatewaysStripeGatewayId Delete a stripe gateway

Delete a stripe gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripeGatewayId The resource's id
 @return ApiDELETEStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) DELETEStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId string) ApiDELETEStripeGatewaysStripeGatewayIdRequest {
	return ApiDELETEStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
func (a *StripeGatewaysApiService) DELETEStripeGatewaysStripeGatewayIdExecute(r ApiDELETEStripeGatewaysStripeGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.DELETEStripeGatewaysStripeGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterToString(r.stripeGatewayId, "")), -1)

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

type ApiGETStripeGatewaysRequest struct {
	ctx        context.Context
	ApiService *StripeGatewaysApiService
}

func (r ApiGETStripeGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStripeGatewaysExecute(r)
}

/*
GETStripeGateways List all stripe gateways

List all stripe gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETStripeGatewaysRequest
*/
func (a *StripeGatewaysApiService) GETStripeGateways(ctx context.Context) ApiGETStripeGatewaysRequest {
	return ApiGETStripeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *StripeGatewaysApiService) GETStripeGatewaysExecute(r ApiGETStripeGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.GETStripeGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways"

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

type ApiGETStripeGatewaysStripeGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *StripeGatewaysApiService
	stripeGatewayId string
}

func (r ApiGETStripeGatewaysStripeGatewayIdRequest) Execute() (*StripeGateway, *http.Response, error) {
	return r.ApiService.GETStripeGatewaysStripeGatewayIdExecute(r)
}

/*
GETStripeGatewaysStripeGatewayId Retrieve a stripe gateway

Retrieve a stripe gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripeGatewayId The resource's id
 @return ApiGETStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) GETStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId string) ApiGETStripeGatewaysStripeGatewayIdRequest {
	return ApiGETStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
//  @return StripeGateway
func (a *StripeGatewaysApiService) GETStripeGatewaysStripeGatewayIdExecute(r ApiGETStripeGatewaysStripeGatewayIdRequest) (*StripeGateway, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StripeGateway
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.GETStripeGatewaysStripeGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterToString(r.stripeGatewayId, "")), -1)

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

type ApiPATCHStripeGatewaysStripeGatewayIdRequest struct {
	ctx                 context.Context
	ApiService          *StripeGatewaysApiService
	stripeGatewayUpdate *StripeGatewayUpdate
	stripeGatewayId     string
}

func (r ApiPATCHStripeGatewaysStripeGatewayIdRequest) StripeGatewayUpdate(stripeGatewayUpdate StripeGatewayUpdate) ApiPATCHStripeGatewaysStripeGatewayIdRequest {
	r.stripeGatewayUpdate = &stripeGatewayUpdate
	return r
}

func (r ApiPATCHStripeGatewaysStripeGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHStripeGatewaysStripeGatewayIdExecute(r)
}

/*
PATCHStripeGatewaysStripeGatewayId Update a stripe gateway

Update a stripe gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripeGatewayId The resource's id
 @return ApiPATCHStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) PATCHStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId string) ApiPATCHStripeGatewaysStripeGatewayIdRequest {
	return ApiPATCHStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
func (a *StripeGatewaysApiService) PATCHStripeGatewaysStripeGatewayIdExecute(r ApiPATCHStripeGatewaysStripeGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.PATCHStripeGatewaysStripeGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterToString(r.stripeGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripeGatewayUpdate == nil {
		return nil, reportError("stripeGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.stripeGatewayUpdate
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

type ApiPOSTStripeGatewaysRequest struct {
	ctx                 context.Context
	ApiService          *StripeGatewaysApiService
	stripeGatewayCreate *StripeGatewayCreate
}

func (r ApiPOSTStripeGatewaysRequest) StripeGatewayCreate(stripeGatewayCreate StripeGatewayCreate) ApiPOSTStripeGatewaysRequest {
	r.stripeGatewayCreate = &stripeGatewayCreate
	return r
}

func (r ApiPOSTStripeGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTStripeGatewaysExecute(r)
}

/*
POSTStripeGateways Create a stripe gateway

Create a stripe gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTStripeGatewaysRequest
*/
func (a *StripeGatewaysApiService) POSTStripeGateways(ctx context.Context) ApiPOSTStripeGatewaysRequest {
	return ApiPOSTStripeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *StripeGatewaysApiService) POSTStripeGatewaysExecute(r ApiPOSTStripeGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.POSTStripeGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripeGatewayCreate == nil {
		return nil, reportError("stripeGatewayCreate is required and must be specified")
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
	localVarPostBody = r.stripeGatewayCreate
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

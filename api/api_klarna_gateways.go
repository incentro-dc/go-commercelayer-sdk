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

// KlarnaGatewaysApiService KlarnaGatewaysApi service
type KlarnaGatewaysApiService service

type ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *KlarnaGatewaysApiService
	klarnaGatewayId string
}

func (r ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEKlarnaGatewaysKlarnaGatewayIdExecute(r)
}

/*
DELETEKlarnaGatewaysKlarnaGatewayId Delete a klarna gateway

Delete a klarna gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param klarnaGatewayId The resource's id
 @return ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest
*/
func (a *KlarnaGatewaysApiService) DELETEKlarnaGatewaysKlarnaGatewayId(ctx context.Context, klarnaGatewayId string) ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest {
	return ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaGatewayId: klarnaGatewayId,
	}
}

// Execute executes the request
func (a *KlarnaGatewaysApiService) DELETEKlarnaGatewaysKlarnaGatewayIdExecute(r ApiDELETEKlarnaGatewaysKlarnaGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaGatewaysApiService.DELETEKlarnaGatewaysKlarnaGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways/{klarnaGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaGatewayId"+"}", url.PathEscape(parameterToString(r.klarnaGatewayId, "")), -1)

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

type ApiGETKlarnaGatewaysRequest struct {
	ctx        context.Context
	ApiService *KlarnaGatewaysApiService
}

func (r ApiGETKlarnaGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETKlarnaGatewaysExecute(r)
}

/*
GETKlarnaGateways List all klarna gateways

List all klarna gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETKlarnaGatewaysRequest
*/
func (a *KlarnaGatewaysApiService) GETKlarnaGateways(ctx context.Context) ApiGETKlarnaGatewaysRequest {
	return ApiGETKlarnaGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *KlarnaGatewaysApiService) GETKlarnaGatewaysExecute(r ApiGETKlarnaGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaGatewaysApiService.GETKlarnaGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways"

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

type ApiGETKlarnaGatewaysKlarnaGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *KlarnaGatewaysApiService
	klarnaGatewayId string
}

func (r ApiGETKlarnaGatewaysKlarnaGatewayIdRequest) Execute() (*KlarnaGateway, *http.Response, error) {
	return r.ApiService.GETKlarnaGatewaysKlarnaGatewayIdExecute(r)
}

/*
GETKlarnaGatewaysKlarnaGatewayId Retrieve a klarna gateway

Retrieve a klarna gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param klarnaGatewayId The resource's id
 @return ApiGETKlarnaGatewaysKlarnaGatewayIdRequest
*/
func (a *KlarnaGatewaysApiService) GETKlarnaGatewaysKlarnaGatewayId(ctx context.Context, klarnaGatewayId string) ApiGETKlarnaGatewaysKlarnaGatewayIdRequest {
	return ApiGETKlarnaGatewaysKlarnaGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaGatewayId: klarnaGatewayId,
	}
}

// Execute executes the request
//  @return KlarnaGateway
func (a *KlarnaGatewaysApiService) GETKlarnaGatewaysKlarnaGatewayIdExecute(r ApiGETKlarnaGatewaysKlarnaGatewayIdRequest) (*KlarnaGateway, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KlarnaGateway
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaGatewaysApiService.GETKlarnaGatewaysKlarnaGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways/{klarnaGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaGatewayId"+"}", url.PathEscape(parameterToString(r.klarnaGatewayId, "")), -1)

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

type ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest struct {
	ctx                 context.Context
	ApiService          *KlarnaGatewaysApiService
	klarnaGatewayUpdate *KlarnaGatewayUpdate
	klarnaGatewayId     string
}

func (r ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest) KlarnaGatewayUpdate(klarnaGatewayUpdate KlarnaGatewayUpdate) ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest {
	r.klarnaGatewayUpdate = &klarnaGatewayUpdate
	return r
}

func (r ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHKlarnaGatewaysKlarnaGatewayIdExecute(r)
}

/*
PATCHKlarnaGatewaysKlarnaGatewayId Update a klarna gateway

Update a klarna gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param klarnaGatewayId The resource's id
 @return ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest
*/
func (a *KlarnaGatewaysApiService) PATCHKlarnaGatewaysKlarnaGatewayId(ctx context.Context, klarnaGatewayId string) ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest {
	return ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		klarnaGatewayId: klarnaGatewayId,
	}
}

// Execute executes the request
func (a *KlarnaGatewaysApiService) PATCHKlarnaGatewaysKlarnaGatewayIdExecute(r ApiPATCHKlarnaGatewaysKlarnaGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaGatewaysApiService.PATCHKlarnaGatewaysKlarnaGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways/{klarnaGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaGatewayId"+"}", url.PathEscape(parameterToString(r.klarnaGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.klarnaGatewayUpdate == nil {
		return nil, reportError("klarnaGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.klarnaGatewayUpdate
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

type ApiPOSTKlarnaGatewaysRequest struct {
	ctx                 context.Context
	ApiService          *KlarnaGatewaysApiService
	klarnaGatewayCreate *KlarnaGatewayCreate
}

func (r ApiPOSTKlarnaGatewaysRequest) KlarnaGatewayCreate(klarnaGatewayCreate KlarnaGatewayCreate) ApiPOSTKlarnaGatewaysRequest {
	r.klarnaGatewayCreate = &klarnaGatewayCreate
	return r
}

func (r ApiPOSTKlarnaGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTKlarnaGatewaysExecute(r)
}

/*
POSTKlarnaGateways Create a klarna gateway

Create a klarna gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTKlarnaGatewaysRequest
*/
func (a *KlarnaGatewaysApiService) POSTKlarnaGateways(ctx context.Context) ApiPOSTKlarnaGatewaysRequest {
	return ApiPOSTKlarnaGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *KlarnaGatewaysApiService) POSTKlarnaGatewaysExecute(r ApiPOSTKlarnaGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlarnaGatewaysApiService.POSTKlarnaGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.klarnaGatewayCreate == nil {
		return nil, reportError("klarnaGatewayCreate is required and must be specified")
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
	localVarPostBody = r.klarnaGatewayCreate
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

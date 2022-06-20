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

type ExternalGatewaysApi interface {

	/*
		DELETEExternalGatewaysExternalGatewayId Delete an external gateway

		Delete an external gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalGatewayId The resource's id
		@return ApiDELETEExternalGatewaysExternalGatewayIdRequest
	*/
	DELETEExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiDELETEExternalGatewaysExternalGatewayIdRequest

	// DELETEExternalGatewaysExternalGatewayIdExecute executes the request
	DELETEExternalGatewaysExternalGatewayIdExecute(r ApiDELETEExternalGatewaysExternalGatewayIdRequest) (*http.Response, error)

	/*
		GETExternalGateways List all external gateways

		List all external gateways

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETExternalGatewaysRequest
	*/
	GETExternalGateways(ctx context.Context) ApiGETExternalGatewaysRequest

	// GETExternalGatewaysExecute executes the request
	GETExternalGatewaysExecute(r ApiGETExternalGatewaysRequest) (*http.Response, error)

	/*
		GETExternalGatewaysExternalGatewayId Retrieve an external gateway

		Retrieve an external gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalGatewayId The resource's id
		@return ApiGETExternalGatewaysExternalGatewayIdRequest
	*/
	GETExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiGETExternalGatewaysExternalGatewayIdRequest

	// GETExternalGatewaysExternalGatewayIdExecute executes the request
	//  @return ExternalGateway
	GETExternalGatewaysExternalGatewayIdExecute(r ApiGETExternalGatewaysExternalGatewayIdRequest) (*ExternalGateway, *http.Response, error)

	/*
		PATCHExternalGatewaysExternalGatewayId Update an external gateway

		Update an external gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalGatewayId The resource's id
		@return ApiPATCHExternalGatewaysExternalGatewayIdRequest
	*/
	PATCHExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiPATCHExternalGatewaysExternalGatewayIdRequest

	// PATCHExternalGatewaysExternalGatewayIdExecute executes the request
	PATCHExternalGatewaysExternalGatewayIdExecute(r ApiPATCHExternalGatewaysExternalGatewayIdRequest) (*http.Response, error)

	/*
		POSTExternalGateways Create an external gateway

		Create an external gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTExternalGatewaysRequest
	*/
	POSTExternalGateways(ctx context.Context) ApiPOSTExternalGatewaysRequest

	// POSTExternalGatewaysExecute executes the request
	POSTExternalGatewaysExecute(r ApiPOSTExternalGatewaysRequest) (*http.Response, error)
}

// ExternalGatewaysApiService ExternalGatewaysApi service
type ExternalGatewaysApiService service

type ApiDELETEExternalGatewaysExternalGatewayIdRequest struct {
	ctx               context.Context
	ApiService        ExternalGatewaysApi
	externalGatewayId string
}

func (r ApiDELETEExternalGatewaysExternalGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEExternalGatewaysExternalGatewayIdExecute(r)
}

/*
DELETEExternalGatewaysExternalGatewayId Delete an external gateway

Delete an external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGatewayId The resource's id
 @return ApiDELETEExternalGatewaysExternalGatewayIdRequest
*/
func (a *ExternalGatewaysApiService) DELETEExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiDELETEExternalGatewaysExternalGatewayIdRequest {
	return ApiDELETEExternalGatewaysExternalGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalGatewayId: externalGatewayId,
	}
}

// Execute executes the request
func (a *ExternalGatewaysApiService) DELETEExternalGatewaysExternalGatewayIdExecute(r ApiDELETEExternalGatewaysExternalGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalGatewaysApiService.DELETEExternalGatewaysExternalGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways/{externalGatewayId}"
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

type ApiGETExternalGatewaysRequest struct {
	ctx        context.Context
	ApiService ExternalGatewaysApi
}

func (r ApiGETExternalGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalGatewaysExecute(r)
}

/*
GETExternalGateways List all external gateways

List all external gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETExternalGatewaysRequest
*/
func (a *ExternalGatewaysApiService) GETExternalGateways(ctx context.Context) ApiGETExternalGatewaysRequest {
	return ApiGETExternalGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalGatewaysApiService) GETExternalGatewaysExecute(r ApiGETExternalGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalGatewaysApiService.GETExternalGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways"

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

type ApiGETExternalGatewaysExternalGatewayIdRequest struct {
	ctx               context.Context
	ApiService        ExternalGatewaysApi
	externalGatewayId string
}

func (r ApiGETExternalGatewaysExternalGatewayIdRequest) Execute() (*ExternalGateway, *http.Response, error) {
	return r.ApiService.GETExternalGatewaysExternalGatewayIdExecute(r)
}

/*
GETExternalGatewaysExternalGatewayId Retrieve an external gateway

Retrieve an external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGatewayId The resource's id
 @return ApiGETExternalGatewaysExternalGatewayIdRequest
*/
func (a *ExternalGatewaysApiService) GETExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiGETExternalGatewaysExternalGatewayIdRequest {
	return ApiGETExternalGatewaysExternalGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalGatewayId: externalGatewayId,
	}
}

// Execute executes the request
//  @return ExternalGateway
func (a *ExternalGatewaysApiService) GETExternalGatewaysExternalGatewayIdExecute(r ApiGETExternalGatewaysExternalGatewayIdRequest) (*ExternalGateway, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalGateway
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalGatewaysApiService.GETExternalGatewaysExternalGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways/{externalGatewayId}"
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

type ApiPATCHExternalGatewaysExternalGatewayIdRequest struct {
	ctx                   context.Context
	ApiService            ExternalGatewaysApi
	externalGatewayId     string
	externalGatewayUpdate *ExternalGatewayUpdate
}

func (r ApiPATCHExternalGatewaysExternalGatewayIdRequest) ExternalGatewayUpdate(externalGatewayUpdate ExternalGatewayUpdate) ApiPATCHExternalGatewaysExternalGatewayIdRequest {
	r.externalGatewayUpdate = &externalGatewayUpdate
	return r
}

func (r ApiPATCHExternalGatewaysExternalGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHExternalGatewaysExternalGatewayIdExecute(r)
}

/*
PATCHExternalGatewaysExternalGatewayId Update an external gateway

Update an external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGatewayId The resource's id
 @return ApiPATCHExternalGatewaysExternalGatewayIdRequest
*/
func (a *ExternalGatewaysApiService) PATCHExternalGatewaysExternalGatewayId(ctx context.Context, externalGatewayId string) ApiPATCHExternalGatewaysExternalGatewayIdRequest {
	return ApiPATCHExternalGatewaysExternalGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalGatewayId: externalGatewayId,
	}
}

// Execute executes the request
func (a *ExternalGatewaysApiService) PATCHExternalGatewaysExternalGatewayIdExecute(r ApiPATCHExternalGatewaysExternalGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalGatewaysApiService.PATCHExternalGatewaysExternalGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways/{externalGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGatewayId"+"}", url.PathEscape(parameterToString(r.externalGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalGatewayUpdate == nil {
		return nil, reportError("externalGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.externalGatewayUpdate
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

type ApiPOSTExternalGatewaysRequest struct {
	ctx                   context.Context
	ApiService            ExternalGatewaysApi
	externalGatewayCreate *ExternalGatewayCreate
}

func (r ApiPOSTExternalGatewaysRequest) ExternalGatewayCreate(externalGatewayCreate ExternalGatewayCreate) ApiPOSTExternalGatewaysRequest {
	r.externalGatewayCreate = &externalGatewayCreate
	return r
}

func (r ApiPOSTExternalGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTExternalGatewaysExecute(r)
}

/*
POSTExternalGateways Create an external gateway

Create an external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTExternalGatewaysRequest
*/
func (a *ExternalGatewaysApiService) POSTExternalGateways(ctx context.Context) ApiPOSTExternalGatewaysRequest {
	return ApiPOSTExternalGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalGatewaysApiService) POSTExternalGatewaysExecute(r ApiPOSTExternalGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalGatewaysApiService.POSTExternalGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalGatewayCreate == nil {
		return nil, reportError("externalGatewayCreate is required and must be specified")
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
	localVarPostBody = r.externalGatewayCreate
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

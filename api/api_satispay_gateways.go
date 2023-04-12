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

// SatispayGatewaysApiService SatispayGatewaysApi service
type SatispayGatewaysApiService service

type SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest struct {
	ctx               context.Context
	ApiService        *SatispayGatewaysApiService
	satispayGatewayId interface{}
}

func (r SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETESatispayGatewaysSatispayGatewayIdExecute(r)
}

/*
DELETESatispayGatewaysSatispayGatewayId Delete a satispay gateway

Delete a satispay gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param satispayGatewayId The resource's id
	@return SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest
*/
func (a *SatispayGatewaysApiService) DELETESatispayGatewaysSatispayGatewayId(ctx context.Context, satispayGatewayId interface{}) SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest {
	return SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		satispayGatewayId: satispayGatewayId,
	}
}

// Execute executes the request
func (a *SatispayGatewaysApiService) DELETESatispayGatewaysSatispayGatewayIdExecute(r SatispayGatewaysApiDELETESatispayGatewaysSatispayGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SatispayGatewaysApiService.DELETESatispayGatewaysSatispayGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/satispay_gateways/{satispayGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"satispayGatewayId"+"}", url.PathEscape(parameterValueToString(r.satispayGatewayId, "satispayGatewayId")), -1)

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

type SatispayGatewaysApiGETSatispayGatewaysRequest struct {
	ctx        context.Context
	ApiService *SatispayGatewaysApiService
}

func (r SatispayGatewaysApiGETSatispayGatewaysRequest) Execute() (*GETSatispayGateways200Response, *http.Response, error) {
	return r.ApiService.GETSatispayGatewaysExecute(r)
}

/*
GETSatispayGateways List all satispay gateways

List all satispay gateways

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SatispayGatewaysApiGETSatispayGatewaysRequest
*/
func (a *SatispayGatewaysApiService) GETSatispayGateways(ctx context.Context) SatispayGatewaysApiGETSatispayGatewaysRequest {
	return SatispayGatewaysApiGETSatispayGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETSatispayGateways200Response
func (a *SatispayGatewaysApiService) GETSatispayGatewaysExecute(r SatispayGatewaysApiGETSatispayGatewaysRequest) (*GETSatispayGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSatispayGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SatispayGatewaysApiService.GETSatispayGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/satispay_gateways"

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

type SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest struct {
	ctx               context.Context
	ApiService        *SatispayGatewaysApiService
	satispayGatewayId interface{}
}

func (r SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest) Execute() (*GETSatispayGatewaysSatispayGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETSatispayGatewaysSatispayGatewayIdExecute(r)
}

/*
GETSatispayGatewaysSatispayGatewayId Retrieve a satispay gateway

Retrieve a satispay gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param satispayGatewayId The resource's id
	@return SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest
*/
func (a *SatispayGatewaysApiService) GETSatispayGatewaysSatispayGatewayId(ctx context.Context, satispayGatewayId interface{}) SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest {
	return SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		satispayGatewayId: satispayGatewayId,
	}
}

// Execute executes the request
//
//	@return GETSatispayGatewaysSatispayGatewayId200Response
func (a *SatispayGatewaysApiService) GETSatispayGatewaysSatispayGatewayIdExecute(r SatispayGatewaysApiGETSatispayGatewaysSatispayGatewayIdRequest) (*GETSatispayGatewaysSatispayGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSatispayGatewaysSatispayGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SatispayGatewaysApiService.GETSatispayGatewaysSatispayGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/satispay_gateways/{satispayGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"satispayGatewayId"+"}", url.PathEscape(parameterValueToString(r.satispayGatewayId, "satispayGatewayId")), -1)

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

type SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest struct {
	ctx                   context.Context
	ApiService            *SatispayGatewaysApiService
	satispayGatewayUpdate *SatispayGatewayUpdate
	satispayGatewayId     interface{}
}

func (r SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest) SatispayGatewayUpdate(satispayGatewayUpdate SatispayGatewayUpdate) SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest {
	r.satispayGatewayUpdate = &satispayGatewayUpdate
	return r
}

func (r SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest) Execute() (*PATCHSatispayGatewaysSatispayGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHSatispayGatewaysSatispayGatewayIdExecute(r)
}

/*
PATCHSatispayGatewaysSatispayGatewayId Update a satispay gateway

Update a satispay gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param satispayGatewayId The resource's id
	@return SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest
*/
func (a *SatispayGatewaysApiService) PATCHSatispayGatewaysSatispayGatewayId(ctx context.Context, satispayGatewayId interface{}) SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest {
	return SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest{
		ApiService:        a,
		ctx:               ctx,
		satispayGatewayId: satispayGatewayId,
	}
}

// Execute executes the request
//
//	@return PATCHSatispayGatewaysSatispayGatewayId200Response
func (a *SatispayGatewaysApiService) PATCHSatispayGatewaysSatispayGatewayIdExecute(r SatispayGatewaysApiPATCHSatispayGatewaysSatispayGatewayIdRequest) (*PATCHSatispayGatewaysSatispayGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHSatispayGatewaysSatispayGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SatispayGatewaysApiService.PATCHSatispayGatewaysSatispayGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/satispay_gateways/{satispayGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"satispayGatewayId"+"}", url.PathEscape(parameterValueToString(r.satispayGatewayId, "satispayGatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.satispayGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("satispayGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.satispayGatewayUpdate
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

type SatispayGatewaysApiPOSTSatispayGatewaysRequest struct {
	ctx                   context.Context
	ApiService            *SatispayGatewaysApiService
	satispayGatewayCreate *SatispayGatewayCreate
}

func (r SatispayGatewaysApiPOSTSatispayGatewaysRequest) SatispayGatewayCreate(satispayGatewayCreate SatispayGatewayCreate) SatispayGatewaysApiPOSTSatispayGatewaysRequest {
	r.satispayGatewayCreate = &satispayGatewayCreate
	return r
}

func (r SatispayGatewaysApiPOSTSatispayGatewaysRequest) Execute() (*POSTSatispayGateways201Response, *http.Response, error) {
	return r.ApiService.POSTSatispayGatewaysExecute(r)
}

/*
POSTSatispayGateways Create a satispay gateway

Create a satispay gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SatispayGatewaysApiPOSTSatispayGatewaysRequest
*/
func (a *SatispayGatewaysApiService) POSTSatispayGateways(ctx context.Context) SatispayGatewaysApiPOSTSatispayGatewaysRequest {
	return SatispayGatewaysApiPOSTSatispayGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTSatispayGateways201Response
func (a *SatispayGatewaysApiService) POSTSatispayGatewaysExecute(r SatispayGatewaysApiPOSTSatispayGatewaysRequest) (*POSTSatispayGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTSatispayGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SatispayGatewaysApiService.POSTSatispayGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/satispay_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.satispayGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("satispayGatewayCreate is required and must be specified")
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
	localVarPostBody = r.satispayGatewayCreate
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

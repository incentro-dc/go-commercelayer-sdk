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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// BingGeocodersApiService BingGeocodersApi service
type BingGeocodersApiService service

type BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest struct {
	ctx            context.Context
	ApiService     *BingGeocodersApiService
	bingGeocoderId interface{}
}

func (r BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBingGeocodersBingGeocoderIdExecute(r)
}

/*
DELETEBingGeocodersBingGeocoderId Delete a bing geocoder

Delete a bing geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bingGeocoderId The resource's id
	@return BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) DELETEBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId interface{}) BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest {
	return BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
func (a *BingGeocodersApiService) DELETEBingGeocodersBingGeocoderIdExecute(r BingGeocodersApiDELETEBingGeocodersBingGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.DELETEBingGeocodersBingGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

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

type BingGeocodersApiGETBingGeocodersRequest struct {
	ctx        context.Context
	ApiService *BingGeocodersApiService
}

func (r BingGeocodersApiGETBingGeocodersRequest) Execute() (*GETBingGeocoders200Response, *http.Response, error) {
	return r.ApiService.GETBingGeocodersExecute(r)
}

/*
GETBingGeocoders List all bing geocoders

List all bing geocoders

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BingGeocodersApiGETBingGeocodersRequest
*/
func (a *BingGeocodersApiService) GETBingGeocoders(ctx context.Context) BingGeocodersApiGETBingGeocodersRequest {
	return BingGeocodersApiGETBingGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETBingGeocoders200Response
func (a *BingGeocodersApiService) GETBingGeocodersExecute(r BingGeocodersApiGETBingGeocodersRequest) (*GETBingGeocoders200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBingGeocoders200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.GETBingGeocoders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders"

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

type BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest struct {
	ctx            context.Context
	ApiService     *BingGeocodersApiService
	bingGeocoderId interface{}
}

func (r BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest) Execute() (*GETBingGeocodersBingGeocoderId200Response, *http.Response, error) {
	return r.ApiService.GETBingGeocodersBingGeocoderIdExecute(r)
}

/*
GETBingGeocodersBingGeocoderId Retrieve a bing geocoder

Retrieve a bing geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bingGeocoderId The resource's id
	@return BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) GETBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId interface{}) BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest {
	return BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
//
//	@return GETBingGeocodersBingGeocoderId200Response
func (a *BingGeocodersApiService) GETBingGeocodersBingGeocoderIdExecute(r BingGeocodersApiGETBingGeocodersBingGeocoderIdRequest) (*GETBingGeocodersBingGeocoderId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBingGeocodersBingGeocoderId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.GETBingGeocodersBingGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

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

type BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest struct {
	ctx                context.Context
	ApiService         *BingGeocodersApiService
	bingGeocoderUpdate *BingGeocoderUpdate
	bingGeocoderId     interface{}
}

func (r BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest) BingGeocoderUpdate(bingGeocoderUpdate BingGeocoderUpdate) BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest {
	r.bingGeocoderUpdate = &bingGeocoderUpdate
	return r
}

func (r BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest) Execute() (*PATCHBingGeocodersBingGeocoderId200Response, *http.Response, error) {
	return r.ApiService.PATCHBingGeocodersBingGeocoderIdExecute(r)
}

/*
PATCHBingGeocodersBingGeocoderId Update a bing geocoder

Update a bing geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bingGeocoderId The resource's id
	@return BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) PATCHBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId interface{}) BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest {
	return BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
//
//	@return PATCHBingGeocodersBingGeocoderId200Response
func (a *BingGeocodersApiService) PATCHBingGeocodersBingGeocoderIdExecute(r BingGeocodersApiPATCHBingGeocodersBingGeocoderIdRequest) (*PATCHBingGeocodersBingGeocoderId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHBingGeocodersBingGeocoderId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.PATCHBingGeocodersBingGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bingGeocoderUpdate == nil {
		return localVarReturnValue, nil, reportError("bingGeocoderUpdate is required and must be specified")
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
	localVarPostBody = r.bingGeocoderUpdate
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

type BingGeocodersApiPOSTBingGeocodersRequest struct {
	ctx                context.Context
	ApiService         *BingGeocodersApiService
	bingGeocoderCreate *BingGeocoderCreate
}

func (r BingGeocodersApiPOSTBingGeocodersRequest) BingGeocoderCreate(bingGeocoderCreate BingGeocoderCreate) BingGeocodersApiPOSTBingGeocodersRequest {
	r.bingGeocoderCreate = &bingGeocoderCreate
	return r
}

func (r BingGeocodersApiPOSTBingGeocodersRequest) Execute() (*POSTBingGeocoders201Response, *http.Response, error) {
	return r.ApiService.POSTBingGeocodersExecute(r)
}

/*
POSTBingGeocoders Create a bing geocoder

Create a bing geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BingGeocodersApiPOSTBingGeocodersRequest
*/
func (a *BingGeocodersApiService) POSTBingGeocoders(ctx context.Context) BingGeocodersApiPOSTBingGeocodersRequest {
	return BingGeocodersApiPOSTBingGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTBingGeocoders201Response
func (a *BingGeocodersApiService) POSTBingGeocodersExecute(r BingGeocodersApiPOSTBingGeocodersRequest) (*POSTBingGeocoders201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTBingGeocoders201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.POSTBingGeocoders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bingGeocoderCreate == nil {
		return localVarReturnValue, nil, reportError("bingGeocoderCreate is required and must be specified")
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
	localVarPostBody = r.bingGeocoderCreate
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

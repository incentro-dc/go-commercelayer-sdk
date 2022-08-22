/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
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

// ShippingZonesApiService ShippingZonesApi service
type ShippingZonesApiService service

type ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest struct {
	ctx            context.Context
	ApiService     *ShippingZonesApiService
	shippingZoneId string
}

func (r ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEShippingZonesShippingZoneIdExecute(r)
}

/*
DELETEShippingZonesShippingZoneId Delete a shipping zone

Delete a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) DELETEShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest {
	return ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) DELETEShippingZonesShippingZoneIdExecute(r ShippingZonesApiDELETEShippingZonesShippingZoneIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.DELETEShippingZonesShippingZoneId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

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

type ShippingZonesApiGETShippingMethodIdShippingZoneRequest struct {
	ctx              context.Context
	ApiService       *ShippingZonesApiService
	shippingMethodId string
}

func (r ShippingZonesApiGETShippingMethodIdShippingZoneRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdShippingZoneExecute(r)
}

/*
GETShippingMethodIdShippingZone Retrieve the shipping zone associated to the shipping method

Retrieve the shipping zone associated to the shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ShippingZonesApiGETShippingMethodIdShippingZoneRequest
*/
func (a *ShippingZonesApiService) GETShippingMethodIdShippingZone(ctx context.Context, shippingMethodId string) ShippingZonesApiGETShippingMethodIdShippingZoneRequest {
	return ShippingZonesApiGETShippingMethodIdShippingZoneRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) GETShippingMethodIdShippingZoneExecute(r ShippingZonesApiGETShippingMethodIdShippingZoneRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingMethodIdShippingZone")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/shipping_zone"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

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

type ShippingZonesApiGETShippingZonesRequest struct {
	ctx        context.Context
	ApiService *ShippingZonesApiService
}

func (r ShippingZonesApiGETShippingZonesRequest) Execute() (*GETShippingZones200Response, *http.Response, error) {
	return r.ApiService.GETShippingZonesExecute(r)
}

/*
GETShippingZones List all shipping zones

List all shipping zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShippingZonesApiGETShippingZonesRequest
*/
func (a *ShippingZonesApiService) GETShippingZones(ctx context.Context) ShippingZonesApiGETShippingZonesRequest {
	return ShippingZonesApiGETShippingZonesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETShippingZones200Response
func (a *ShippingZonesApiService) GETShippingZonesExecute(r ShippingZonesApiGETShippingZonesRequest) (*GETShippingZones200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingZones200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingZones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones"

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

type ShippingZonesApiGETShippingZonesShippingZoneIdRequest struct {
	ctx            context.Context
	ApiService     *ShippingZonesApiService
	shippingZoneId string
}

func (r ShippingZonesApiGETShippingZonesShippingZoneIdRequest) Execute() (*GETShippingZonesShippingZoneId200Response, *http.Response, error) {
	return r.ApiService.GETShippingZonesShippingZoneIdExecute(r)
}

/*
GETShippingZonesShippingZoneId Retrieve a shipping zone

Retrieve a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ShippingZonesApiGETShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) GETShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ShippingZonesApiGETShippingZonesShippingZoneIdRequest {
	return ShippingZonesApiGETShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
//  @return GETShippingZonesShippingZoneId200Response
func (a *ShippingZonesApiService) GETShippingZonesShippingZoneIdExecute(r ShippingZonesApiGETShippingZonesShippingZoneIdRequest) (*GETShippingZonesShippingZoneId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingZonesShippingZoneId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingZonesShippingZoneId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

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

type ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest struct {
	ctx                context.Context
	ApiService         *ShippingZonesApiService
	shippingZoneUpdate *ShippingZoneUpdate
	shippingZoneId     string
}

func (r ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest) ShippingZoneUpdate(shippingZoneUpdate ShippingZoneUpdate) ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest {
	r.shippingZoneUpdate = &shippingZoneUpdate
	return r
}

func (r ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest) Execute() (*PATCHShippingZonesShippingZoneId200Response, *http.Response, error) {
	return r.ApiService.PATCHShippingZonesShippingZoneIdExecute(r)
}

/*
PATCHShippingZonesShippingZoneId Update a shipping zone

Update a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) PATCHShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest {
	return ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
//  @return PATCHShippingZonesShippingZoneId200Response
func (a *ShippingZonesApiService) PATCHShippingZonesShippingZoneIdExecute(r ShippingZonesApiPATCHShippingZonesShippingZoneIdRequest) (*PATCHShippingZonesShippingZoneId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHShippingZonesShippingZoneId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.PATCHShippingZonesShippingZoneId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingZoneUpdate == nil {
		return localVarReturnValue, nil, reportError("shippingZoneUpdate is required and must be specified")
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
	localVarPostBody = r.shippingZoneUpdate
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

type ShippingZonesApiPOSTShippingZonesRequest struct {
	ctx                context.Context
	ApiService         *ShippingZonesApiService
	shippingZoneCreate *ShippingZoneCreate
}

func (r ShippingZonesApiPOSTShippingZonesRequest) ShippingZoneCreate(shippingZoneCreate ShippingZoneCreate) ShippingZonesApiPOSTShippingZonesRequest {
	r.shippingZoneCreate = &shippingZoneCreate
	return r
}

func (r ShippingZonesApiPOSTShippingZonesRequest) Execute() (*POSTShippingZones201Response, *http.Response, error) {
	return r.ApiService.POSTShippingZonesExecute(r)
}

/*
POSTShippingZones Create a shipping zone

Create a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShippingZonesApiPOSTShippingZonesRequest
*/
func (a *ShippingZonesApiService) POSTShippingZones(ctx context.Context) ShippingZonesApiPOSTShippingZonesRequest {
	return ShippingZonesApiPOSTShippingZonesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTShippingZones201Response
func (a *ShippingZonesApiService) POSTShippingZonesExecute(r ShippingZonesApiPOSTShippingZonesRequest) (*POSTShippingZones201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTShippingZones201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.POSTShippingZones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingZoneCreate == nil {
		return localVarReturnValue, nil, reportError("shippingZoneCreate is required and must be specified")
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
	localVarPostBody = r.shippingZoneCreate
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

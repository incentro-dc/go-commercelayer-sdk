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

// ShippingWeightTiersApiService ShippingWeightTiersApi service
type ShippingWeightTiersApiService service

type ApiDELETEShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingWeightTiersApiService
	shippingWeightTierId string
}

func (r ApiDELETEShippingWeightTiersShippingWeightTierIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
DELETEShippingWeightTiersShippingWeightTierId Delete a shipping weight tier

Delete a shipping weight tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingWeightTierId The resource's id
 @return ApiDELETEShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) DELETEShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId string) ApiDELETEShippingWeightTiersShippingWeightTierIdRequest {
	return ApiDELETEShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) DELETEShippingWeightTiersShippingWeightTierIdExecute(r ApiDELETEShippingWeightTiersShippingWeightTierIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.DELETEShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterToString(r.shippingWeightTierId, "")), -1)

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

type ApiGETShippingMethodIdShippingWeightTiersRequest struct {
	ctx              context.Context
	ApiService       *ShippingWeightTiersApiService
	shippingMethodId string
}

func (r ApiGETShippingMethodIdShippingWeightTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdShippingWeightTiersExecute(r)
}

/*
GETShippingMethodIdShippingWeightTiers Retrieve the shipping weight tiers associated to the shipping method

Retrieve the shipping weight tiers associated to the shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ApiGETShippingMethodIdShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingMethodIdShippingWeightTiers(ctx context.Context, shippingMethodId string) ApiGETShippingMethodIdShippingWeightTiersRequest {
	return ApiGETShippingMethodIdShippingWeightTiersRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) GETShippingMethodIdShippingWeightTiersExecute(r ApiGETShippingMethodIdShippingWeightTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingMethodIdShippingWeightTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/shipping_weight_tiers"
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

type ApiGETShippingWeightTiersRequest struct {
	ctx        context.Context
	ApiService *ShippingWeightTiersApiService
}

func (r ApiGETShippingWeightTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingWeightTiersExecute(r)
}

/*
GETShippingWeightTiers List all shipping weight tiers

List all shipping weight tiers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingWeightTiers(ctx context.Context) ApiGETShippingWeightTiersRequest {
	return ApiGETShippingWeightTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersExecute(r ApiGETShippingWeightTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingWeightTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers"

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

type ApiGETShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingWeightTiersApiService
	shippingWeightTierId string
}

func (r ApiGETShippingWeightTiersShippingWeightTierIdRequest) Execute() (*ShippingWeightTier, *http.Response, error) {
	return r.ApiService.GETShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
GETShippingWeightTiersShippingWeightTierId Retrieve a shipping weight tier

Retrieve a shipping weight tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingWeightTierId The resource's id
 @return ApiGETShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId string) ApiGETShippingWeightTiersShippingWeightTierIdRequest {
	return ApiGETShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
//  @return ShippingWeightTier
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersShippingWeightTierIdExecute(r ApiGETShippingWeightTiersShippingWeightTierIdRequest) (*ShippingWeightTier, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingWeightTier
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterToString(r.shippingWeightTierId, "")), -1)

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

type ApiPATCHShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                      context.Context
	ApiService               *ShippingWeightTiersApiService
	shippingWeightTierUpdate *ShippingWeightTierUpdate
	shippingWeightTierId     string
}

func (r ApiPATCHShippingWeightTiersShippingWeightTierIdRequest) ShippingWeightTierUpdate(shippingWeightTierUpdate ShippingWeightTierUpdate) ApiPATCHShippingWeightTiersShippingWeightTierIdRequest {
	r.shippingWeightTierUpdate = &shippingWeightTierUpdate
	return r
}

func (r ApiPATCHShippingWeightTiersShippingWeightTierIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
PATCHShippingWeightTiersShippingWeightTierId Update a shipping weight tier

Update a shipping weight tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingWeightTierId The resource's id
 @return ApiPATCHShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) PATCHShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId string) ApiPATCHShippingWeightTiersShippingWeightTierIdRequest {
	return ApiPATCHShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) PATCHShippingWeightTiersShippingWeightTierIdExecute(r ApiPATCHShippingWeightTiersShippingWeightTierIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.PATCHShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterToString(r.shippingWeightTierId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingWeightTierUpdate == nil {
		return nil, reportError("shippingWeightTierUpdate is required and must be specified")
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
	localVarPostBody = r.shippingWeightTierUpdate
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

type ApiPOSTShippingWeightTiersRequest struct {
	ctx                      context.Context
	ApiService               *ShippingWeightTiersApiService
	shippingWeightTierCreate *ShippingWeightTierCreate
}

func (r ApiPOSTShippingWeightTiersRequest) ShippingWeightTierCreate(shippingWeightTierCreate ShippingWeightTierCreate) ApiPOSTShippingWeightTiersRequest {
	r.shippingWeightTierCreate = &shippingWeightTierCreate
	return r
}

func (r ApiPOSTShippingWeightTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTShippingWeightTiersExecute(r)
}

/*
POSTShippingWeightTiers Create a shipping weight tier

Create a shipping weight tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) POSTShippingWeightTiers(ctx context.Context) ApiPOSTShippingWeightTiersRequest {
	return ApiPOSTShippingWeightTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) POSTShippingWeightTiersExecute(r ApiPOSTShippingWeightTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.POSTShippingWeightTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingWeightTierCreate == nil {
		return nil, reportError("shippingWeightTierCreate is required and must be specified")
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
	localVarPostBody = r.shippingWeightTierCreate
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

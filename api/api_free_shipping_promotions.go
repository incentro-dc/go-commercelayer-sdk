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

// FreeShippingPromotionsApiService FreeShippingPromotionsApi service
type FreeShippingPromotionsApiService service

type FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest struct {
	ctx                     context.Context
	ApiService              *FreeShippingPromotionsApiService
	freeShippingPromotionId interface{}
}

func (r FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEFreeShippingPromotionsFreeShippingPromotionIdExecute(r)
}

/*
DELETEFreeShippingPromotionsFreeShippingPromotionId Delete a free shipping promotion

Delete a free shipping promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeShippingPromotionId The resource's id
	@return FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest
*/
func (a *FreeShippingPromotionsApiService) DELETEFreeShippingPromotionsFreeShippingPromotionId(ctx context.Context, freeShippingPromotionId interface{}) FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest {
	return FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		freeShippingPromotionId: freeShippingPromotionId,
	}
}

// Execute executes the request
func (a *FreeShippingPromotionsApiService) DELETEFreeShippingPromotionsFreeShippingPromotionIdExecute(r FreeShippingPromotionsApiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeShippingPromotionsApiService.DELETEFreeShippingPromotionsFreeShippingPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_shipping_promotions/{freeShippingPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeShippingPromotionId"+"}", url.PathEscape(parameterToString(r.freeShippingPromotionId, "")), -1)

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

type FreeShippingPromotionsApiGETFreeShippingPromotionsRequest struct {
	ctx        context.Context
	ApiService *FreeShippingPromotionsApiService
}

func (r FreeShippingPromotionsApiGETFreeShippingPromotionsRequest) Execute() (*GETFreeShippingPromotions200Response, *http.Response, error) {
	return r.ApiService.GETFreeShippingPromotionsExecute(r)
}

/*
GETFreeShippingPromotions List all free shipping promotions

List all free shipping promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FreeShippingPromotionsApiGETFreeShippingPromotionsRequest
*/
func (a *FreeShippingPromotionsApiService) GETFreeShippingPromotions(ctx context.Context) FreeShippingPromotionsApiGETFreeShippingPromotionsRequest {
	return FreeShippingPromotionsApiGETFreeShippingPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETFreeShippingPromotions200Response
func (a *FreeShippingPromotionsApiService) GETFreeShippingPromotionsExecute(r FreeShippingPromotionsApiGETFreeShippingPromotionsRequest) (*GETFreeShippingPromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFreeShippingPromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeShippingPromotionsApiService.GETFreeShippingPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_shipping_promotions"

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

type FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest struct {
	ctx                     context.Context
	ApiService              *FreeShippingPromotionsApiService
	freeShippingPromotionId interface{}
}

func (r FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest) Execute() (*GETFreeShippingPromotionsFreeShippingPromotionId200Response, *http.Response, error) {
	return r.ApiService.GETFreeShippingPromotionsFreeShippingPromotionIdExecute(r)
}

/*
GETFreeShippingPromotionsFreeShippingPromotionId Retrieve a free shipping promotion

Retrieve a free shipping promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeShippingPromotionId The resource's id
	@return FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest
*/
func (a *FreeShippingPromotionsApiService) GETFreeShippingPromotionsFreeShippingPromotionId(ctx context.Context, freeShippingPromotionId interface{}) FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest {
	return FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		freeShippingPromotionId: freeShippingPromotionId,
	}
}

// Execute executes the request
//
//	@return GETFreeShippingPromotionsFreeShippingPromotionId200Response
func (a *FreeShippingPromotionsApiService) GETFreeShippingPromotionsFreeShippingPromotionIdExecute(r FreeShippingPromotionsApiGETFreeShippingPromotionsFreeShippingPromotionIdRequest) (*GETFreeShippingPromotionsFreeShippingPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFreeShippingPromotionsFreeShippingPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeShippingPromotionsApiService.GETFreeShippingPromotionsFreeShippingPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_shipping_promotions/{freeShippingPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeShippingPromotionId"+"}", url.PathEscape(parameterToString(r.freeShippingPromotionId, "")), -1)

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

type FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest struct {
	ctx                         context.Context
	ApiService                  *FreeShippingPromotionsApiService
	freeShippingPromotionUpdate *FreeShippingPromotionUpdate
	freeShippingPromotionId     interface{}
}

func (r FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest) FreeShippingPromotionUpdate(freeShippingPromotionUpdate FreeShippingPromotionUpdate) FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest {
	r.freeShippingPromotionUpdate = &freeShippingPromotionUpdate
	return r
}

func (r FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest) Execute() (*PATCHFreeShippingPromotionsFreeShippingPromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHFreeShippingPromotionsFreeShippingPromotionIdExecute(r)
}

/*
PATCHFreeShippingPromotionsFreeShippingPromotionId Update a free shipping promotion

Update a free shipping promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeShippingPromotionId The resource's id
	@return FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest
*/
func (a *FreeShippingPromotionsApiService) PATCHFreeShippingPromotionsFreeShippingPromotionId(ctx context.Context, freeShippingPromotionId interface{}) FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest {
	return FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		freeShippingPromotionId: freeShippingPromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHFreeShippingPromotionsFreeShippingPromotionId200Response
func (a *FreeShippingPromotionsApiService) PATCHFreeShippingPromotionsFreeShippingPromotionIdExecute(r FreeShippingPromotionsApiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest) (*PATCHFreeShippingPromotionsFreeShippingPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeShippingPromotionsApiService.PATCHFreeShippingPromotionsFreeShippingPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_shipping_promotions/{freeShippingPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeShippingPromotionId"+"}", url.PathEscape(parameterToString(r.freeShippingPromotionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.freeShippingPromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("freeShippingPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.freeShippingPromotionUpdate
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

type FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest struct {
	ctx                         context.Context
	ApiService                  *FreeShippingPromotionsApiService
	freeShippingPromotionCreate *FreeShippingPromotionCreate
}

func (r FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest) FreeShippingPromotionCreate(freeShippingPromotionCreate FreeShippingPromotionCreate) FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest {
	r.freeShippingPromotionCreate = &freeShippingPromotionCreate
	return r
}

func (r FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest) Execute() (*POSTFreeShippingPromotions201Response, *http.Response, error) {
	return r.ApiService.POSTFreeShippingPromotionsExecute(r)
}

/*
POSTFreeShippingPromotions Create a free shipping promotion

Create a free shipping promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest
*/
func (a *FreeShippingPromotionsApiService) POSTFreeShippingPromotions(ctx context.Context) FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest {
	return FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTFreeShippingPromotions201Response
func (a *FreeShippingPromotionsApiService) POSTFreeShippingPromotionsExecute(r FreeShippingPromotionsApiPOSTFreeShippingPromotionsRequest) (*POSTFreeShippingPromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTFreeShippingPromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeShippingPromotionsApiService.POSTFreeShippingPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_shipping_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.freeShippingPromotionCreate == nil {
		return localVarReturnValue, nil, reportError("freeShippingPromotionCreate is required and must be specified")
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
	localVarPostBody = r.freeShippingPromotionCreate
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

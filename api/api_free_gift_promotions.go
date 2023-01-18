/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
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

// FreeGiftPromotionsApiService FreeGiftPromotionsApi service
type FreeGiftPromotionsApiService service

type FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest struct {
	ctx                 context.Context
	ApiService          *FreeGiftPromotionsApiService
	freeGiftPromotionId string
}

func (r FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEFreeGiftPromotionsFreeGiftPromotionIdExecute(r)
}

/*
DELETEFreeGiftPromotionsFreeGiftPromotionId Delete a free gift promotion

Delete a free gift promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeGiftPromotionId The resource's id
	@return FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest
*/
func (a *FreeGiftPromotionsApiService) DELETEFreeGiftPromotionsFreeGiftPromotionId(ctx context.Context, freeGiftPromotionId string) FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest {
	return FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		freeGiftPromotionId: freeGiftPromotionId,
	}
}

// Execute executes the request
func (a *FreeGiftPromotionsApiService) DELETEFreeGiftPromotionsFreeGiftPromotionIdExecute(r FreeGiftPromotionsApiDELETEFreeGiftPromotionsFreeGiftPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeGiftPromotionsApiService.DELETEFreeGiftPromotionsFreeGiftPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_gift_promotions/{freeGiftPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeGiftPromotionId"+"}", url.PathEscape(parameterToString(r.freeGiftPromotionId, "")), -1)

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

type FreeGiftPromotionsApiGETFreeGiftPromotionsRequest struct {
	ctx        context.Context
	ApiService *FreeGiftPromotionsApiService
}

func (r FreeGiftPromotionsApiGETFreeGiftPromotionsRequest) Execute() (*GETFreeGiftPromotions200Response, *http.Response, error) {
	return r.ApiService.GETFreeGiftPromotionsExecute(r)
}

/*
GETFreeGiftPromotions List all free gift promotions

List all free gift promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FreeGiftPromotionsApiGETFreeGiftPromotionsRequest
*/
func (a *FreeGiftPromotionsApiService) GETFreeGiftPromotions(ctx context.Context) FreeGiftPromotionsApiGETFreeGiftPromotionsRequest {
	return FreeGiftPromotionsApiGETFreeGiftPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETFreeGiftPromotions200Response
func (a *FreeGiftPromotionsApiService) GETFreeGiftPromotionsExecute(r FreeGiftPromotionsApiGETFreeGiftPromotionsRequest) (*GETFreeGiftPromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFreeGiftPromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeGiftPromotionsApiService.GETFreeGiftPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_gift_promotions"

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

type FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest struct {
	ctx                 context.Context
	ApiService          *FreeGiftPromotionsApiService
	freeGiftPromotionId string
}

func (r FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest) Execute() (*GETFreeGiftPromotionsFreeGiftPromotionId200Response, *http.Response, error) {
	return r.ApiService.GETFreeGiftPromotionsFreeGiftPromotionIdExecute(r)
}

/*
GETFreeGiftPromotionsFreeGiftPromotionId Retrieve a free gift promotion

Retrieve a free gift promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeGiftPromotionId The resource's id
	@return FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest
*/
func (a *FreeGiftPromotionsApiService) GETFreeGiftPromotionsFreeGiftPromotionId(ctx context.Context, freeGiftPromotionId string) FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest {
	return FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		freeGiftPromotionId: freeGiftPromotionId,
	}
}

// Execute executes the request
//
//	@return GETFreeGiftPromotionsFreeGiftPromotionId200Response
func (a *FreeGiftPromotionsApiService) GETFreeGiftPromotionsFreeGiftPromotionIdExecute(r FreeGiftPromotionsApiGETFreeGiftPromotionsFreeGiftPromotionIdRequest) (*GETFreeGiftPromotionsFreeGiftPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFreeGiftPromotionsFreeGiftPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeGiftPromotionsApiService.GETFreeGiftPromotionsFreeGiftPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_gift_promotions/{freeGiftPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeGiftPromotionId"+"}", url.PathEscape(parameterToString(r.freeGiftPromotionId, "")), -1)

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

type FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest struct {
	ctx                     context.Context
	ApiService              *FreeGiftPromotionsApiService
	freeGiftPromotionUpdate *FreeGiftPromotionUpdate
	freeGiftPromotionId     string
}

func (r FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) FreeGiftPromotionUpdate(freeGiftPromotionUpdate FreeGiftPromotionUpdate) FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	r.freeGiftPromotionUpdate = &freeGiftPromotionUpdate
	return r
}

func (r FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) Execute() (*PATCHFreeGiftPromotionsFreeGiftPromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHFreeGiftPromotionsFreeGiftPromotionIdExecute(r)
}

/*
PATCHFreeGiftPromotionsFreeGiftPromotionId Update a free gift promotion

Update a free gift promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param freeGiftPromotionId The resource's id
	@return FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest
*/
func (a *FreeGiftPromotionsApiService) PATCHFreeGiftPromotionsFreeGiftPromotionId(ctx context.Context, freeGiftPromotionId string) FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	return FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		freeGiftPromotionId: freeGiftPromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHFreeGiftPromotionsFreeGiftPromotionId200Response
func (a *FreeGiftPromotionsApiService) PATCHFreeGiftPromotionsFreeGiftPromotionIdExecute(r FreeGiftPromotionsApiPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) (*PATCHFreeGiftPromotionsFreeGiftPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeGiftPromotionsApiService.PATCHFreeGiftPromotionsFreeGiftPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_gift_promotions/{freeGiftPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"freeGiftPromotionId"+"}", url.PathEscape(parameterToString(r.freeGiftPromotionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.freeGiftPromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("freeGiftPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.freeGiftPromotionUpdate
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

type FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest struct {
	ctx                     context.Context
	ApiService              *FreeGiftPromotionsApiService
	freeGiftPromotionCreate *FreeGiftPromotionCreate
}

func (r FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest) FreeGiftPromotionCreate(freeGiftPromotionCreate FreeGiftPromotionCreate) FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest {
	r.freeGiftPromotionCreate = &freeGiftPromotionCreate
	return r
}

func (r FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest) Execute() (*POSTFreeGiftPromotions201Response, *http.Response, error) {
	return r.ApiService.POSTFreeGiftPromotionsExecute(r)
}

/*
POSTFreeGiftPromotions Create a free gift promotion

Create a free gift promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest
*/
func (a *FreeGiftPromotionsApiService) POSTFreeGiftPromotions(ctx context.Context) FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest {
	return FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTFreeGiftPromotions201Response
func (a *FreeGiftPromotionsApiService) POSTFreeGiftPromotionsExecute(r FreeGiftPromotionsApiPOSTFreeGiftPromotionsRequest) (*POSTFreeGiftPromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTFreeGiftPromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeGiftPromotionsApiService.POSTFreeGiftPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free_gift_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.freeGiftPromotionCreate == nil {
		return localVarReturnValue, nil, reportError("freeGiftPromotionCreate is required and must be specified")
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
	localVarPostBody = r.freeGiftPromotionCreate
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

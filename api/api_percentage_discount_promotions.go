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

// PercentageDiscountPromotionsApiService PercentageDiscountPromotionsApi service
type PercentageDiscountPromotionsApiService service

type PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct {
	ctx                           context.Context
	ApiService                    *PercentageDiscountPromotionsApiService
	percentageDiscountPromotionId interface{}
}

func (r PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r)
}

/*
DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId Delete a percentage discount promotion

Delete a percentage discount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param percentageDiscountPromotionId The resource's id
	@return PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest
*/
func (a *PercentageDiscountPromotionsApiService) DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx context.Context, percentageDiscountPromotionId interface{}) PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest {
	return PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest{
		ApiService:                    a,
		ctx:                           ctx,
		percentageDiscountPromotionId: percentageDiscountPromotionId,
	}
}

// Execute executes the request
func (a *PercentageDiscountPromotionsApiService) DELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r PercentageDiscountPromotionsApiDELETEPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PercentageDiscountPromotionsApiService.DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/percentage_discount_promotions/{percentageDiscountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"percentageDiscountPromotionId"+"}", url.PathEscape(parameterValueToString(r.percentageDiscountPromotionId, "percentageDiscountPromotionId")), -1)

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

type PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest struct {
	ctx        context.Context
	ApiService *PercentageDiscountPromotionsApiService
}

func (r PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest) Execute() (*GETPercentageDiscountPromotions200Response, *http.Response, error) {
	return r.ApiService.GETPercentageDiscountPromotionsExecute(r)
}

/*
GETPercentageDiscountPromotions List all percentage discount promotions

List all percentage discount promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest
*/
func (a *PercentageDiscountPromotionsApiService) GETPercentageDiscountPromotions(ctx context.Context) PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest {
	return PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETPercentageDiscountPromotions200Response
func (a *PercentageDiscountPromotionsApiService) GETPercentageDiscountPromotionsExecute(r PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsRequest) (*GETPercentageDiscountPromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPercentageDiscountPromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PercentageDiscountPromotionsApiService.GETPercentageDiscountPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/percentage_discount_promotions"

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

type PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct {
	ctx                           context.Context
	ApiService                    *PercentageDiscountPromotionsApiService
	percentageDiscountPromotionId interface{}
}

func (r PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) Execute() (*GETPercentageDiscountPromotionsPercentageDiscountPromotionId200Response, *http.Response, error) {
	return r.ApiService.GETPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r)
}

/*
GETPercentageDiscountPromotionsPercentageDiscountPromotionId Retrieve a percentage discount promotion

Retrieve a percentage discount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param percentageDiscountPromotionId The resource's id
	@return PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest
*/
func (a *PercentageDiscountPromotionsApiService) GETPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx context.Context, percentageDiscountPromotionId interface{}) PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest {
	return PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest{
		ApiService:                    a,
		ctx:                           ctx,
		percentageDiscountPromotionId: percentageDiscountPromotionId,
	}
}

// Execute executes the request
//
//	@return GETPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
func (a *PercentageDiscountPromotionsApiService) GETPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r PercentageDiscountPromotionsApiGETPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) (*GETPercentageDiscountPromotionsPercentageDiscountPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PercentageDiscountPromotionsApiService.GETPercentageDiscountPromotionsPercentageDiscountPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/percentage_discount_promotions/{percentageDiscountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"percentageDiscountPromotionId"+"}", url.PathEscape(parameterValueToString(r.percentageDiscountPromotionId, "percentageDiscountPromotionId")), -1)

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

type PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest struct {
	ctx                               context.Context
	ApiService                        *PercentageDiscountPromotionsApiService
	percentageDiscountPromotionUpdate *PercentageDiscountPromotionUpdate
	percentageDiscountPromotionId     interface{}
}

func (r PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) PercentageDiscountPromotionUpdate(percentageDiscountPromotionUpdate PercentageDiscountPromotionUpdate) PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest {
	r.percentageDiscountPromotionUpdate = &percentageDiscountPromotionUpdate
	return r
}

func (r PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) Execute() (*PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r)
}

/*
PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId Update a percentage discount promotion

Update a percentage discount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param percentageDiscountPromotionId The resource's id
	@return PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest
*/
func (a *PercentageDiscountPromotionsApiService) PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId(ctx context.Context, percentageDiscountPromotionId interface{}) PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest {
	return PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest{
		ApiService:                    a,
		ctx:                           ctx,
		percentageDiscountPromotionId: percentageDiscountPromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
func (a *PercentageDiscountPromotionsApiService) PATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdExecute(r PercentageDiscountPromotionsApiPATCHPercentageDiscountPromotionsPercentageDiscountPromotionIdRequest) (*PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PercentageDiscountPromotionsApiService.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/percentage_discount_promotions/{percentageDiscountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"percentageDiscountPromotionId"+"}", url.PathEscape(parameterValueToString(r.percentageDiscountPromotionId, "percentageDiscountPromotionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.percentageDiscountPromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("percentageDiscountPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.percentageDiscountPromotionUpdate
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

type PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest struct {
	ctx                               context.Context
	ApiService                        *PercentageDiscountPromotionsApiService
	percentageDiscountPromotionCreate *PercentageDiscountPromotionCreate
}

func (r PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest) PercentageDiscountPromotionCreate(percentageDiscountPromotionCreate PercentageDiscountPromotionCreate) PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest {
	r.percentageDiscountPromotionCreate = &percentageDiscountPromotionCreate
	return r
}

func (r PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest) Execute() (*POSTPercentageDiscountPromotions201Response, *http.Response, error) {
	return r.ApiService.POSTPercentageDiscountPromotionsExecute(r)
}

/*
POSTPercentageDiscountPromotions Create a percentage discount promotion

Create a percentage discount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest
*/
func (a *PercentageDiscountPromotionsApiService) POSTPercentageDiscountPromotions(ctx context.Context) PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest {
	return PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTPercentageDiscountPromotions201Response
func (a *PercentageDiscountPromotionsApiService) POSTPercentageDiscountPromotionsExecute(r PercentageDiscountPromotionsApiPOSTPercentageDiscountPromotionsRequest) (*POSTPercentageDiscountPromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPercentageDiscountPromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PercentageDiscountPromotionsApiService.POSTPercentageDiscountPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/percentage_discount_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.percentageDiscountPromotionCreate == nil {
		return localVarReturnValue, nil, reportError("percentageDiscountPromotionCreate is required and must be specified")
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
	localVarPostBody = r.percentageDiscountPromotionCreate
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

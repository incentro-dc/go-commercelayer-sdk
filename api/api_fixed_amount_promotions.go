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

// FixedAmountPromotionsApiService FixedAmountPromotionsApi service
type FixedAmountPromotionsApiService service

type FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest struct {
	ctx                    context.Context
	ApiService             *FixedAmountPromotionsApiService
	fixedAmountPromotionId string
}

func (r FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEFixedAmountPromotionsFixedAmountPromotionIdExecute(r)
}

/*
DELETEFixedAmountPromotionsFixedAmountPromotionId Delete a fixed amount promotion

Delete a fixed amount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedAmountPromotionId The resource's id
	@return FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest
*/
func (a *FixedAmountPromotionsApiService) DELETEFixedAmountPromotionsFixedAmountPromotionId(ctx context.Context, fixedAmountPromotionId string) FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest {
	return FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		fixedAmountPromotionId: fixedAmountPromotionId,
	}
}

// Execute executes the request
func (a *FixedAmountPromotionsApiService) DELETEFixedAmountPromotionsFixedAmountPromotionIdExecute(r FixedAmountPromotionsApiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedAmountPromotionsApiService.DELETEFixedAmountPromotionsFixedAmountPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_amount_promotions/{fixedAmountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedAmountPromotionId"+"}", url.PathEscape(parameterToString(r.fixedAmountPromotionId, "")), -1)

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

type FixedAmountPromotionsApiGETFixedAmountPromotionsRequest struct {
	ctx        context.Context
	ApiService *FixedAmountPromotionsApiService
}

func (r FixedAmountPromotionsApiGETFixedAmountPromotionsRequest) Execute() (*GETFixedAmountPromotions200Response, *http.Response, error) {
	return r.ApiService.GETFixedAmountPromotionsExecute(r)
}

/*
GETFixedAmountPromotions List all fixed amount promotions

List all fixed amount promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FixedAmountPromotionsApiGETFixedAmountPromotionsRequest
*/
func (a *FixedAmountPromotionsApiService) GETFixedAmountPromotions(ctx context.Context) FixedAmountPromotionsApiGETFixedAmountPromotionsRequest {
	return FixedAmountPromotionsApiGETFixedAmountPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETFixedAmountPromotions200Response
func (a *FixedAmountPromotionsApiService) GETFixedAmountPromotionsExecute(r FixedAmountPromotionsApiGETFixedAmountPromotionsRequest) (*GETFixedAmountPromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFixedAmountPromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedAmountPromotionsApiService.GETFixedAmountPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_amount_promotions"

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

type FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest struct {
	ctx                    context.Context
	ApiService             *FixedAmountPromotionsApiService
	fixedAmountPromotionId string
}

func (r FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest) Execute() (*GETFixedAmountPromotionsFixedAmountPromotionId200Response, *http.Response, error) {
	return r.ApiService.GETFixedAmountPromotionsFixedAmountPromotionIdExecute(r)
}

/*
GETFixedAmountPromotionsFixedAmountPromotionId Retrieve a fixed amount promotion

Retrieve a fixed amount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedAmountPromotionId The resource's id
	@return FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest
*/
func (a *FixedAmountPromotionsApiService) GETFixedAmountPromotionsFixedAmountPromotionId(ctx context.Context, fixedAmountPromotionId string) FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest {
	return FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		fixedAmountPromotionId: fixedAmountPromotionId,
	}
}

// Execute executes the request
//
//	@return GETFixedAmountPromotionsFixedAmountPromotionId200Response
func (a *FixedAmountPromotionsApiService) GETFixedAmountPromotionsFixedAmountPromotionIdExecute(r FixedAmountPromotionsApiGETFixedAmountPromotionsFixedAmountPromotionIdRequest) (*GETFixedAmountPromotionsFixedAmountPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFixedAmountPromotionsFixedAmountPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedAmountPromotionsApiService.GETFixedAmountPromotionsFixedAmountPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_amount_promotions/{fixedAmountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedAmountPromotionId"+"}", url.PathEscape(parameterToString(r.fixedAmountPromotionId, "")), -1)

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

type FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest struct {
	ctx                        context.Context
	ApiService                 *FixedAmountPromotionsApiService
	fixedAmountPromotionUpdate *FixedAmountPromotionUpdate
	fixedAmountPromotionId     string
}

func (r FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest) FixedAmountPromotionUpdate(fixedAmountPromotionUpdate FixedAmountPromotionUpdate) FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest {
	r.fixedAmountPromotionUpdate = &fixedAmountPromotionUpdate
	return r
}

func (r FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest) Execute() (*PATCHFixedAmountPromotionsFixedAmountPromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHFixedAmountPromotionsFixedAmountPromotionIdExecute(r)
}

/*
PATCHFixedAmountPromotionsFixedAmountPromotionId Update a fixed amount promotion

Update a fixed amount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedAmountPromotionId The resource's id
	@return FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest
*/
func (a *FixedAmountPromotionsApiService) PATCHFixedAmountPromotionsFixedAmountPromotionId(ctx context.Context, fixedAmountPromotionId string) FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest {
	return FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		fixedAmountPromotionId: fixedAmountPromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHFixedAmountPromotionsFixedAmountPromotionId200Response
func (a *FixedAmountPromotionsApiService) PATCHFixedAmountPromotionsFixedAmountPromotionIdExecute(r FixedAmountPromotionsApiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest) (*PATCHFixedAmountPromotionsFixedAmountPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedAmountPromotionsApiService.PATCHFixedAmountPromotionsFixedAmountPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_amount_promotions/{fixedAmountPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedAmountPromotionId"+"}", url.PathEscape(parameterToString(r.fixedAmountPromotionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fixedAmountPromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("fixedAmountPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.fixedAmountPromotionUpdate
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

type FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest struct {
	ctx                        context.Context
	ApiService                 *FixedAmountPromotionsApiService
	fixedAmountPromotionCreate *FixedAmountPromotionCreate
}

func (r FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest) FixedAmountPromotionCreate(fixedAmountPromotionCreate FixedAmountPromotionCreate) FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest {
	r.fixedAmountPromotionCreate = &fixedAmountPromotionCreate
	return r
}

func (r FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest) Execute() (*POSTFixedAmountPromotions201Response, *http.Response, error) {
	return r.ApiService.POSTFixedAmountPromotionsExecute(r)
}

/*
POSTFixedAmountPromotions Create a fixed amount promotion

Create a fixed amount promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest
*/
func (a *FixedAmountPromotionsApiService) POSTFixedAmountPromotions(ctx context.Context) FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest {
	return FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTFixedAmountPromotions201Response
func (a *FixedAmountPromotionsApiService) POSTFixedAmountPromotionsExecute(r FixedAmountPromotionsApiPOSTFixedAmountPromotionsRequest) (*POSTFixedAmountPromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTFixedAmountPromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedAmountPromotionsApiService.POSTFixedAmountPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_amount_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fixedAmountPromotionCreate == nil {
		return localVarReturnValue, nil, reportError("fixedAmountPromotionCreate is required and must be specified")
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
	localVarPostBody = r.fixedAmountPromotionCreate
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

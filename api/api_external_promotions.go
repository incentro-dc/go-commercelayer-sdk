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

type ExternalPromotionsApi interface {

	/*
		DELETEExternalPromotionsExternalPromotionId Delete an external promotion

		Delete an external promotion

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalPromotionId The resource's id
		@return ApiDELETEExternalPromotionsExternalPromotionIdRequest
	*/
	DELETEExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiDELETEExternalPromotionsExternalPromotionIdRequest

	// DELETEExternalPromotionsExternalPromotionIdExecute executes the request
	DELETEExternalPromotionsExternalPromotionIdExecute(r ApiDELETEExternalPromotionsExternalPromotionIdRequest) (*http.Response, error)

	/*
		GETExternalPromotions List all external promotions

		List all external promotions

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETExternalPromotionsRequest
	*/
	GETExternalPromotions(ctx context.Context) ApiGETExternalPromotionsRequest

	// GETExternalPromotionsExecute executes the request
	GETExternalPromotionsExecute(r ApiGETExternalPromotionsRequest) (*http.Response, error)

	/*
		GETExternalPromotionsExternalPromotionId Retrieve an external promotion

		Retrieve an external promotion

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalPromotionId The resource's id
		@return ApiGETExternalPromotionsExternalPromotionIdRequest
	*/
	GETExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiGETExternalPromotionsExternalPromotionIdRequest

	// GETExternalPromotionsExternalPromotionIdExecute executes the request
	//  @return ExternalPromotion
	GETExternalPromotionsExternalPromotionIdExecute(r ApiGETExternalPromotionsExternalPromotionIdRequest) (*ExternalPromotion, *http.Response, error)

	/*
		PATCHExternalPromotionsExternalPromotionId Update an external promotion

		Update an external promotion

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalPromotionId The resource's id
		@return ApiPATCHExternalPromotionsExternalPromotionIdRequest
	*/
	PATCHExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiPATCHExternalPromotionsExternalPromotionIdRequest

	// PATCHExternalPromotionsExternalPromotionIdExecute executes the request
	PATCHExternalPromotionsExternalPromotionIdExecute(r ApiPATCHExternalPromotionsExternalPromotionIdRequest) (*http.Response, error)

	/*
		POSTExternalPromotions Create an external promotion

		Create an external promotion

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTExternalPromotionsRequest
	*/
	POSTExternalPromotions(ctx context.Context) ApiPOSTExternalPromotionsRequest

	// POSTExternalPromotionsExecute executes the request
	POSTExternalPromotionsExecute(r ApiPOSTExternalPromotionsRequest) (*http.Response, error)
}

// ExternalPromotionsApiService ExternalPromotionsApi service
type ExternalPromotionsApiService service

type ApiDELETEExternalPromotionsExternalPromotionIdRequest struct {
	ctx                 context.Context
	ApiService          ExternalPromotionsApi
	externalPromotionId string
}

func (r ApiDELETEExternalPromotionsExternalPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEExternalPromotionsExternalPromotionIdExecute(r)
}

/*
DELETEExternalPromotionsExternalPromotionId Delete an external promotion

Delete an external promotion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPromotionId The resource's id
 @return ApiDELETEExternalPromotionsExternalPromotionIdRequest
*/
func (a *ExternalPromotionsApiService) DELETEExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiDELETEExternalPromotionsExternalPromotionIdRequest {
	return ApiDELETEExternalPromotionsExternalPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		externalPromotionId: externalPromotionId,
	}
}

// Execute executes the request
func (a *ExternalPromotionsApiService) DELETEExternalPromotionsExternalPromotionIdExecute(r ApiDELETEExternalPromotionsExternalPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPromotionsApiService.DELETEExternalPromotionsExternalPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_promotions/{externalPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPromotionId"+"}", url.PathEscape(parameterToString(r.externalPromotionId, "")), -1)

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

type ApiGETExternalPromotionsRequest struct {
	ctx        context.Context
	ApiService ExternalPromotionsApi
}

func (r ApiGETExternalPromotionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalPromotionsExecute(r)
}

/*
GETExternalPromotions List all external promotions

List all external promotions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETExternalPromotionsRequest
*/
func (a *ExternalPromotionsApiService) GETExternalPromotions(ctx context.Context) ApiGETExternalPromotionsRequest {
	return ApiGETExternalPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalPromotionsApiService) GETExternalPromotionsExecute(r ApiGETExternalPromotionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPromotionsApiService.GETExternalPromotions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_promotions"

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

type ApiGETExternalPromotionsExternalPromotionIdRequest struct {
	ctx                 context.Context
	ApiService          ExternalPromotionsApi
	externalPromotionId string
}

func (r ApiGETExternalPromotionsExternalPromotionIdRequest) Execute() (*ExternalPromotion, *http.Response, error) {
	return r.ApiService.GETExternalPromotionsExternalPromotionIdExecute(r)
}

/*
GETExternalPromotionsExternalPromotionId Retrieve an external promotion

Retrieve an external promotion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPromotionId The resource's id
 @return ApiGETExternalPromotionsExternalPromotionIdRequest
*/
func (a *ExternalPromotionsApiService) GETExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiGETExternalPromotionsExternalPromotionIdRequest {
	return ApiGETExternalPromotionsExternalPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		externalPromotionId: externalPromotionId,
	}
}

// Execute executes the request
//  @return ExternalPromotion
func (a *ExternalPromotionsApiService) GETExternalPromotionsExternalPromotionIdExecute(r ApiGETExternalPromotionsExternalPromotionIdRequest) (*ExternalPromotion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPromotion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPromotionsApiService.GETExternalPromotionsExternalPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_promotions/{externalPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPromotionId"+"}", url.PathEscape(parameterToString(r.externalPromotionId, "")), -1)

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

type ApiPATCHExternalPromotionsExternalPromotionIdRequest struct {
	ctx                     context.Context
	ApiService              ExternalPromotionsApi
	externalPromotionId     string
	externalPromotionUpdate *ExternalPromotionUpdate
}

func (r ApiPATCHExternalPromotionsExternalPromotionIdRequest) ExternalPromotionUpdate(externalPromotionUpdate ExternalPromotionUpdate) ApiPATCHExternalPromotionsExternalPromotionIdRequest {
	r.externalPromotionUpdate = &externalPromotionUpdate
	return r
}

func (r ApiPATCHExternalPromotionsExternalPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHExternalPromotionsExternalPromotionIdExecute(r)
}

/*
PATCHExternalPromotionsExternalPromotionId Update an external promotion

Update an external promotion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPromotionId The resource's id
 @return ApiPATCHExternalPromotionsExternalPromotionIdRequest
*/
func (a *ExternalPromotionsApiService) PATCHExternalPromotionsExternalPromotionId(ctx context.Context, externalPromotionId string) ApiPATCHExternalPromotionsExternalPromotionIdRequest {
	return ApiPATCHExternalPromotionsExternalPromotionIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		externalPromotionId: externalPromotionId,
	}
}

// Execute executes the request
func (a *ExternalPromotionsApiService) PATCHExternalPromotionsExternalPromotionIdExecute(r ApiPATCHExternalPromotionsExternalPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPromotionsApiService.PATCHExternalPromotionsExternalPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_promotions/{externalPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPromotionId"+"}", url.PathEscape(parameterToString(r.externalPromotionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPromotionUpdate == nil {
		return nil, reportError("externalPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.externalPromotionUpdate
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

type ApiPOSTExternalPromotionsRequest struct {
	ctx                     context.Context
	ApiService              ExternalPromotionsApi
	externalPromotionCreate *ExternalPromotionCreate
}

func (r ApiPOSTExternalPromotionsRequest) ExternalPromotionCreate(externalPromotionCreate ExternalPromotionCreate) ApiPOSTExternalPromotionsRequest {
	r.externalPromotionCreate = &externalPromotionCreate
	return r
}

func (r ApiPOSTExternalPromotionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTExternalPromotionsExecute(r)
}

/*
POSTExternalPromotions Create an external promotion

Create an external promotion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTExternalPromotionsRequest
*/
func (a *ExternalPromotionsApiService) POSTExternalPromotions(ctx context.Context) ApiPOSTExternalPromotionsRequest {
	return ApiPOSTExternalPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ExternalPromotionsApiService) POSTExternalPromotionsExecute(r ApiPOSTExternalPromotionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPromotionsApiService.POSTExternalPromotions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPromotionCreate == nil {
		return nil, reportError("externalPromotionCreate is required and must be specified")
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
	localVarPostBody = r.externalPromotionCreate
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

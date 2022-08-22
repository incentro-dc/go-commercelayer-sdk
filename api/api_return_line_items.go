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

// ReturnLineItemsApiService ReturnLineItemsApi service
type ReturnLineItemsApiService service

type ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest struct {
	ctx              context.Context
	ApiService       *ReturnLineItemsApiService
	returnLineItemId string
}

func (r ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEReturnLineItemsReturnLineItemIdExecute(r)
}

/*
DELETEReturnLineItemsReturnLineItemId Delete a return line item

Delete a return line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnLineItemId The resource's id
 @return ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest
*/
func (a *ReturnLineItemsApiService) DELETEReturnLineItemsReturnLineItemId(ctx context.Context, returnLineItemId string) ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest {
	return ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		returnLineItemId: returnLineItemId,
	}
}

// Execute executes the request
func (a *ReturnLineItemsApiService) DELETEReturnLineItemsReturnLineItemIdExecute(r ReturnLineItemsApiDELETEReturnLineItemsReturnLineItemIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.DELETEReturnLineItemsReturnLineItemId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items/{returnLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnLineItemId"+"}", url.PathEscape(parameterToString(r.returnLineItemId, "")), -1)

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

type ReturnLineItemsApiGETReturnIdReturnLineItemsRequest struct {
	ctx        context.Context
	ApiService *ReturnLineItemsApiService
	returnId   string
}

func (r ReturnLineItemsApiGETReturnIdReturnLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnIdReturnLineItemsExecute(r)
}

/*
GETReturnIdReturnLineItems Retrieve the return line items associated to the return

Retrieve the return line items associated to the return

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnId The resource's id
 @return ReturnLineItemsApiGETReturnIdReturnLineItemsRequest
*/
func (a *ReturnLineItemsApiService) GETReturnIdReturnLineItems(ctx context.Context, returnId string) ReturnLineItemsApiGETReturnIdReturnLineItemsRequest {
	return ReturnLineItemsApiGETReturnIdReturnLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
func (a *ReturnLineItemsApiService) GETReturnIdReturnLineItemsExecute(r ReturnLineItemsApiGETReturnIdReturnLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.GETReturnIdReturnLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}/return_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterToString(r.returnId, "")), -1)

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

type ReturnLineItemsApiGETReturnLineItemsRequest struct {
	ctx        context.Context
	ApiService *ReturnLineItemsApiService
}

func (r ReturnLineItemsApiGETReturnLineItemsRequest) Execute() (*GETReturnLineItems200Response, *http.Response, error) {
	return r.ApiService.GETReturnLineItemsExecute(r)
}

/*
GETReturnLineItems List all return line items

List all return line items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReturnLineItemsApiGETReturnLineItemsRequest
*/
func (a *ReturnLineItemsApiService) GETReturnLineItems(ctx context.Context) ReturnLineItemsApiGETReturnLineItemsRequest {
	return ReturnLineItemsApiGETReturnLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETReturnLineItems200Response
func (a *ReturnLineItemsApiService) GETReturnLineItemsExecute(r ReturnLineItemsApiGETReturnLineItemsRequest) (*GETReturnLineItems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETReturnLineItems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.GETReturnLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items"

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

type ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest struct {
	ctx              context.Context
	ApiService       *ReturnLineItemsApiService
	returnLineItemId string
}

func (r ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest) Execute() (*GETReturnLineItemsReturnLineItemId200Response, *http.Response, error) {
	return r.ApiService.GETReturnLineItemsReturnLineItemIdExecute(r)
}

/*
GETReturnLineItemsReturnLineItemId Retrieve a return line item

Retrieve a return line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnLineItemId The resource's id
 @return ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest
*/
func (a *ReturnLineItemsApiService) GETReturnLineItemsReturnLineItemId(ctx context.Context, returnLineItemId string) ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest {
	return ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		returnLineItemId: returnLineItemId,
	}
}

// Execute executes the request
//  @return GETReturnLineItemsReturnLineItemId200Response
func (a *ReturnLineItemsApiService) GETReturnLineItemsReturnLineItemIdExecute(r ReturnLineItemsApiGETReturnLineItemsReturnLineItemIdRequest) (*GETReturnLineItemsReturnLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETReturnLineItemsReturnLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.GETReturnLineItemsReturnLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items/{returnLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnLineItemId"+"}", url.PathEscape(parameterToString(r.returnLineItemId, "")), -1)

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

type ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest struct {
	ctx                  context.Context
	ApiService           *ReturnLineItemsApiService
	returnLineItemUpdate *ReturnLineItemUpdate
	returnLineItemId     string
}

func (r ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest) ReturnLineItemUpdate(returnLineItemUpdate ReturnLineItemUpdate) ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest {
	r.returnLineItemUpdate = &returnLineItemUpdate
	return r
}

func (r ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest) Execute() (*PATCHReturnLineItemsReturnLineItemId200Response, *http.Response, error) {
	return r.ApiService.PATCHReturnLineItemsReturnLineItemIdExecute(r)
}

/*
PATCHReturnLineItemsReturnLineItemId Update a return line item

Update a return line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnLineItemId The resource's id
 @return ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest
*/
func (a *ReturnLineItemsApiService) PATCHReturnLineItemsReturnLineItemId(ctx context.Context, returnLineItemId string) ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest {
	return ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		returnLineItemId: returnLineItemId,
	}
}

// Execute executes the request
//  @return PATCHReturnLineItemsReturnLineItemId200Response
func (a *ReturnLineItemsApiService) PATCHReturnLineItemsReturnLineItemIdExecute(r ReturnLineItemsApiPATCHReturnLineItemsReturnLineItemIdRequest) (*PATCHReturnLineItemsReturnLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHReturnLineItemsReturnLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.PATCHReturnLineItemsReturnLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items/{returnLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnLineItemId"+"}", url.PathEscape(parameterToString(r.returnLineItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.returnLineItemUpdate == nil {
		return localVarReturnValue, nil, reportError("returnLineItemUpdate is required and must be specified")
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
	localVarPostBody = r.returnLineItemUpdate
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

type ReturnLineItemsApiPOSTReturnLineItemsRequest struct {
	ctx                  context.Context
	ApiService           *ReturnLineItemsApiService
	returnLineItemCreate *ReturnLineItemCreate
}

func (r ReturnLineItemsApiPOSTReturnLineItemsRequest) ReturnLineItemCreate(returnLineItemCreate ReturnLineItemCreate) ReturnLineItemsApiPOSTReturnLineItemsRequest {
	r.returnLineItemCreate = &returnLineItemCreate
	return r
}

func (r ReturnLineItemsApiPOSTReturnLineItemsRequest) Execute() (*POSTReturnLineItems201Response, *http.Response, error) {
	return r.ApiService.POSTReturnLineItemsExecute(r)
}

/*
POSTReturnLineItems Create a return line item

Create a return line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReturnLineItemsApiPOSTReturnLineItemsRequest
*/
func (a *ReturnLineItemsApiService) POSTReturnLineItems(ctx context.Context) ReturnLineItemsApiPOSTReturnLineItemsRequest {
	return ReturnLineItemsApiPOSTReturnLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTReturnLineItems201Response
func (a *ReturnLineItemsApiService) POSTReturnLineItemsExecute(r ReturnLineItemsApiPOSTReturnLineItemsRequest) (*POSTReturnLineItems201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTReturnLineItems201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnLineItemsApiService.POSTReturnLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.returnLineItemCreate == nil {
		return localVarReturnValue, nil, reportError("returnLineItemCreate is required and must be specified")
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
	localVarPostBody = r.returnLineItemCreate
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

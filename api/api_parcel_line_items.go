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

// ParcelLineItemsApiService ParcelLineItemsApi service
type ParcelLineItemsApiService service

type ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest struct {
	ctx              context.Context
	ApiService       *ParcelLineItemsApiService
	parcelLineItemId interface{}
}

func (r ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEParcelLineItemsParcelLineItemIdExecute(r)
}

/*
DELETEParcelLineItemsParcelLineItemId Delete a parcel line item

Delete a parcel line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parcelLineItemId The resource's id
	@return ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest
*/
func (a *ParcelLineItemsApiService) DELETEParcelLineItemsParcelLineItemId(ctx context.Context, parcelLineItemId interface{}) ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest {
	return ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
func (a *ParcelLineItemsApiService) DELETEParcelLineItemsParcelLineItemIdExecute(r ParcelLineItemsApiDELETEParcelLineItemsParcelLineItemIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.DELETEParcelLineItemsParcelLineItemId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterToString(r.parcelLineItemId, "")), -1)

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

type ParcelLineItemsApiGETParcelIdParcelLineItemsRequest struct {
	ctx        context.Context
	ApiService *ParcelLineItemsApiService
	parcelId   interface{}
}

func (r ParcelLineItemsApiGETParcelIdParcelLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelIdParcelLineItemsExecute(r)
}

/*
GETParcelIdParcelLineItems Retrieve the parcel line items associated to the parcel

Retrieve the parcel line items associated to the parcel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parcelId The resource's id
	@return ParcelLineItemsApiGETParcelIdParcelLineItemsRequest
*/
func (a *ParcelLineItemsApiService) GETParcelIdParcelLineItems(ctx context.Context, parcelId interface{}) ParcelLineItemsApiGETParcelIdParcelLineItemsRequest {
	return ParcelLineItemsApiGETParcelIdParcelLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
func (a *ParcelLineItemsApiService) GETParcelIdParcelLineItemsExecute(r ParcelLineItemsApiGETParcelIdParcelLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.GETParcelIdParcelLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}/parcel_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelId"+"}", url.PathEscape(parameterToString(r.parcelId, "")), -1)

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

type ParcelLineItemsApiGETParcelLineItemsRequest struct {
	ctx        context.Context
	ApiService *ParcelLineItemsApiService
}

func (r ParcelLineItemsApiGETParcelLineItemsRequest) Execute() (*GETParcelLineItems200Response, *http.Response, error) {
	return r.ApiService.GETParcelLineItemsExecute(r)
}

/*
GETParcelLineItems List all parcel line items

List all parcel line items

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ParcelLineItemsApiGETParcelLineItemsRequest
*/
func (a *ParcelLineItemsApiService) GETParcelLineItems(ctx context.Context) ParcelLineItemsApiGETParcelLineItemsRequest {
	return ParcelLineItemsApiGETParcelLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETParcelLineItems200Response
func (a *ParcelLineItemsApiService) GETParcelLineItemsExecute(r ParcelLineItemsApiGETParcelLineItemsRequest) (*GETParcelLineItems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETParcelLineItems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.GETParcelLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items"

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

type ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest struct {
	ctx              context.Context
	ApiService       *ParcelLineItemsApiService
	parcelLineItemId interface{}
}

func (r ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest) Execute() (*GETParcelLineItemsParcelLineItemId200Response, *http.Response, error) {
	return r.ApiService.GETParcelLineItemsParcelLineItemIdExecute(r)
}

/*
GETParcelLineItemsParcelLineItemId Retrieve a parcel line item

Retrieve a parcel line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parcelLineItemId The resource's id
	@return ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest
*/
func (a *ParcelLineItemsApiService) GETParcelLineItemsParcelLineItemId(ctx context.Context, parcelLineItemId interface{}) ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest {
	return ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
//
//	@return GETParcelLineItemsParcelLineItemId200Response
func (a *ParcelLineItemsApiService) GETParcelLineItemsParcelLineItemIdExecute(r ParcelLineItemsApiGETParcelLineItemsParcelLineItemIdRequest) (*GETParcelLineItemsParcelLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETParcelLineItemsParcelLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.GETParcelLineItemsParcelLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterToString(r.parcelLineItemId, "")), -1)

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

type ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest struct {
	ctx                  context.Context
	ApiService           *ParcelLineItemsApiService
	parcelLineItemUpdate *ParcelLineItemUpdate
	parcelLineItemId     interface{}
}

func (r ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest) ParcelLineItemUpdate(parcelLineItemUpdate ParcelLineItemUpdate) ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest {
	r.parcelLineItemUpdate = &parcelLineItemUpdate
	return r
}

func (r ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest) Execute() (*PATCHParcelLineItemsParcelLineItemId200Response, *http.Response, error) {
	return r.ApiService.PATCHParcelLineItemsParcelLineItemIdExecute(r)
}

/*
PATCHParcelLineItemsParcelLineItemId Update a parcel line item

Update a parcel line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parcelLineItemId The resource's id
	@return ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest
*/
func (a *ParcelLineItemsApiService) PATCHParcelLineItemsParcelLineItemId(ctx context.Context, parcelLineItemId interface{}) ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest {
	return ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
//
//	@return PATCHParcelLineItemsParcelLineItemId200Response
func (a *ParcelLineItemsApiService) PATCHParcelLineItemsParcelLineItemIdExecute(r ParcelLineItemsApiPATCHParcelLineItemsParcelLineItemIdRequest) (*PATCHParcelLineItemsParcelLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHParcelLineItemsParcelLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.PATCHParcelLineItemsParcelLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterToString(r.parcelLineItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.parcelLineItemUpdate == nil {
		return localVarReturnValue, nil, reportError("parcelLineItemUpdate is required and must be specified")
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
	localVarPostBody = r.parcelLineItemUpdate
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

type ParcelLineItemsApiPOSTParcelLineItemsRequest struct {
	ctx                  context.Context
	ApiService           *ParcelLineItemsApiService
	parcelLineItemCreate *ParcelLineItemCreate
}

func (r ParcelLineItemsApiPOSTParcelLineItemsRequest) ParcelLineItemCreate(parcelLineItemCreate ParcelLineItemCreate) ParcelLineItemsApiPOSTParcelLineItemsRequest {
	r.parcelLineItemCreate = &parcelLineItemCreate
	return r
}

func (r ParcelLineItemsApiPOSTParcelLineItemsRequest) Execute() (*POSTParcelLineItems201Response, *http.Response, error) {
	return r.ApiService.POSTParcelLineItemsExecute(r)
}

/*
POSTParcelLineItems Create a parcel line item

Create a parcel line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ParcelLineItemsApiPOSTParcelLineItemsRequest
*/
func (a *ParcelLineItemsApiService) POSTParcelLineItems(ctx context.Context) ParcelLineItemsApiPOSTParcelLineItemsRequest {
	return ParcelLineItemsApiPOSTParcelLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTParcelLineItems201Response
func (a *ParcelLineItemsApiService) POSTParcelLineItemsExecute(r ParcelLineItemsApiPOSTParcelLineItemsRequest) (*POSTParcelLineItems201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTParcelLineItems201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelLineItemsApiService.POSTParcelLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.parcelLineItemCreate == nil {
		return localVarReturnValue, nil, reportError("parcelLineItemCreate is required and must be specified")
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
	localVarPostBody = r.parcelLineItemCreate
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

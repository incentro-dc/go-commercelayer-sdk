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

// InventoryModelsApiService InventoryModelsApi service
type InventoryModelsApiService service

type ApiDELETEInventoryModelsInventoryModelIdRequest struct {
	ctx              context.Context
	ApiService       *InventoryModelsApiService
	inventoryModelId string
}

func (r ApiDELETEInventoryModelsInventoryModelIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInventoryModelsInventoryModelIdExecute(r)
}

/*
DELETEInventoryModelsInventoryModelId Delete an inventory model

Delete an inventory model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryModelId The resource's id
 @return ApiDELETEInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) DELETEInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId string) ApiDELETEInventoryModelsInventoryModelIdRequest {
	return ApiDELETEInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) DELETEInventoryModelsInventoryModelIdExecute(r ApiDELETEInventoryModelsInventoryModelIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.DELETEInventoryModelsInventoryModelId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterToString(r.inventoryModelId, "")), -1)

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

type ApiGETInventoryModelsRequest struct {
	ctx        context.Context
	ApiService *InventoryModelsApiService
}

func (r ApiGETInventoryModelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryModelsExecute(r)
}

/*
GETInventoryModels List all inventory models

List all inventory models

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETInventoryModelsRequest
*/
func (a *InventoryModelsApiService) GETInventoryModels(ctx context.Context) ApiGETInventoryModelsRequest {
	return ApiGETInventoryModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETInventoryModelsExecute(r ApiGETInventoryModelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryModels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models"

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

type ApiGETInventoryModelsInventoryModelIdRequest struct {
	ctx              context.Context
	ApiService       *InventoryModelsApiService
	inventoryModelId string
}

func (r ApiGETInventoryModelsInventoryModelIdRequest) Execute() (*InventoryModel, *http.Response, error) {
	return r.ApiService.GETInventoryModelsInventoryModelIdExecute(r)
}

/*
GETInventoryModelsInventoryModelId Retrieve an inventory model

Retrieve an inventory model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryModelId The resource's id
 @return ApiGETInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) GETInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId string) ApiGETInventoryModelsInventoryModelIdRequest {
	return ApiGETInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
//  @return InventoryModel
func (a *InventoryModelsApiService) GETInventoryModelsInventoryModelIdExecute(r ApiGETInventoryModelsInventoryModelIdRequest) (*InventoryModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InventoryModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryModelsInventoryModelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterToString(r.inventoryModelId, "")), -1)

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

type ApiGETInventoryReturnLocationIdInventoryModelRequest struct {
	ctx                       context.Context
	ApiService                *InventoryModelsApiService
	inventoryReturnLocationId string
}

func (r ApiGETInventoryReturnLocationIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryReturnLocationIdInventoryModelExecute(r)
}

/*
GETInventoryReturnLocationIdInventoryModel Retrieve the inventory model associated to the inventory return location

Retrieve the inventory model associated to the inventory return location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryReturnLocationId The resource's id
 @return ApiGETInventoryReturnLocationIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETInventoryReturnLocationIdInventoryModel(ctx context.Context, inventoryReturnLocationId string) ApiGETInventoryReturnLocationIdInventoryModelRequest {
	return ApiGETInventoryReturnLocationIdInventoryModelRequest{
		ApiService:                a,
		ctx:                       ctx,
		inventoryReturnLocationId: inventoryReturnLocationId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETInventoryReturnLocationIdInventoryModelExecute(r ApiGETInventoryReturnLocationIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryReturnLocationIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations/{inventoryReturnLocationId}/inventory_model"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryReturnLocationId"+"}", url.PathEscape(parameterToString(r.inventoryReturnLocationId, "")), -1)

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

type ApiGETInventoryStockLocationIdInventoryModelRequest struct {
	ctx                      context.Context
	ApiService               *InventoryModelsApiService
	inventoryStockLocationId string
}

func (r ApiGETInventoryStockLocationIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryStockLocationIdInventoryModelExecute(r)
}

/*
GETInventoryStockLocationIdInventoryModel Retrieve the inventory model associated to the inventory stock location

Retrieve the inventory model associated to the inventory stock location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryStockLocationId The resource's id
 @return ApiGETInventoryStockLocationIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETInventoryStockLocationIdInventoryModel(ctx context.Context, inventoryStockLocationId string) ApiGETInventoryStockLocationIdInventoryModelRequest {
	return ApiGETInventoryStockLocationIdInventoryModelRequest{
		ApiService:               a,
		ctx:                      ctx,
		inventoryStockLocationId: inventoryStockLocationId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETInventoryStockLocationIdInventoryModelExecute(r ApiGETInventoryStockLocationIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryStockLocationIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations/{inventoryStockLocationId}/inventory_model"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryStockLocationId"+"}", url.PathEscape(parameterToString(r.inventoryStockLocationId, "")), -1)

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

type ApiGETMarketIdInventoryModelRequest struct {
	ctx        context.Context
	ApiService *InventoryModelsApiService
	marketId   string
}

func (r ApiGETMarketIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdInventoryModelExecute(r)
}

/*
GETMarketIdInventoryModel Retrieve the inventory model associated to the market

Retrieve the inventory model associated to the market

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketId The resource's id
 @return ApiGETMarketIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETMarketIdInventoryModel(ctx context.Context, marketId string) ApiGETMarketIdInventoryModelRequest {
	return ApiGETMarketIdInventoryModelRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETMarketIdInventoryModelExecute(r ApiGETMarketIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETMarketIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/inventory_model"
	localVarPath = strings.Replace(localVarPath, "{"+"marketId"+"}", url.PathEscape(parameterToString(r.marketId, "")), -1)

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

type ApiPATCHInventoryModelsInventoryModelIdRequest struct {
	ctx                  context.Context
	ApiService           *InventoryModelsApiService
	inventoryModelUpdate *InventoryModelUpdate
	inventoryModelId     string
}

func (r ApiPATCHInventoryModelsInventoryModelIdRequest) InventoryModelUpdate(inventoryModelUpdate InventoryModelUpdate) ApiPATCHInventoryModelsInventoryModelIdRequest {
	r.inventoryModelUpdate = &inventoryModelUpdate
	return r
}

func (r ApiPATCHInventoryModelsInventoryModelIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHInventoryModelsInventoryModelIdExecute(r)
}

/*
PATCHInventoryModelsInventoryModelId Update an inventory model

Update an inventory model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryModelId The resource's id
 @return ApiPATCHInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) PATCHInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId string) ApiPATCHInventoryModelsInventoryModelIdRequest {
	return ApiPATCHInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) PATCHInventoryModelsInventoryModelIdExecute(r ApiPATCHInventoryModelsInventoryModelIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.PATCHInventoryModelsInventoryModelId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterToString(r.inventoryModelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inventoryModelUpdate == nil {
		return nil, reportError("inventoryModelUpdate is required and must be specified")
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
	localVarPostBody = r.inventoryModelUpdate
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

type ApiPOSTInventoryModelsRequest struct {
	ctx                  context.Context
	ApiService           *InventoryModelsApiService
	inventoryModelCreate *InventoryModelCreate
}

func (r ApiPOSTInventoryModelsRequest) InventoryModelCreate(inventoryModelCreate InventoryModelCreate) ApiPOSTInventoryModelsRequest {
	r.inventoryModelCreate = &inventoryModelCreate
	return r
}

func (r ApiPOSTInventoryModelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTInventoryModelsExecute(r)
}

/*
POSTInventoryModels Create an inventory model

Create an inventory model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTInventoryModelsRequest
*/
func (a *InventoryModelsApiService) POSTInventoryModels(ctx context.Context) ApiPOSTInventoryModelsRequest {
	return ApiPOSTInventoryModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) POSTInventoryModelsExecute(r ApiPOSTInventoryModelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.POSTInventoryModels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inventoryModelCreate == nil {
		return nil, reportError("inventoryModelCreate is required and must be specified")
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
	localVarPostBody = r.inventoryModelCreate
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

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

// StockLineItemsApiService StockLineItemsApi service
type StockLineItemsApiService service

type ApiGETLineItemIdStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
	lineItemId string
}

func (r ApiGETLineItemIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemIdStockLineItemsExecute(r)
}

/*
GETLineItemIdStockLineItems Retrieve the stock line items associated to the line item

Retrieve the stock line items associated to the line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lineItemId The resource's id
 @return ApiGETLineItemIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETLineItemIdStockLineItems(ctx context.Context, lineItemId string) ApiGETLineItemIdStockLineItemsRequest {
	return ApiGETLineItemIdStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETLineItemIdStockLineItemsExecute(r ApiGETLineItemIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETLineItemIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemId"+"}", url.PathEscape(parameterToString(r.lineItemId, "")), -1)

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

type ApiGETParcelLineItemIdStockLineItemRequest struct {
	ctx              context.Context
	ApiService       *StockLineItemsApiService
	parcelLineItemId string
}

func (r ApiGETParcelLineItemIdStockLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelLineItemIdStockLineItemExecute(r)
}

/*
GETParcelLineItemIdStockLineItem Retrieve the stock line item associated to the parcel line item

Retrieve the stock line item associated to the parcel line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelLineItemId The resource's id
 @return ApiGETParcelLineItemIdStockLineItemRequest
*/
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItem(ctx context.Context, parcelLineItemId string) ApiGETParcelLineItemIdStockLineItemRequest {
	return ApiGETParcelLineItemIdStockLineItemRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItemExecute(r ApiGETParcelLineItemIdStockLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETParcelLineItemIdStockLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}/stock_line_item"
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

type ApiGETShipmentIdStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
	shipmentId string
}

func (r ApiGETShipmentIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdStockLineItemsExecute(r)
}

/*
GETShipmentIdStockLineItems Retrieve the stock line items associated to the shipment

Retrieve the stock line items associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ApiGETShipmentIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETShipmentIdStockLineItems(ctx context.Context, shipmentId string) ApiGETShipmentIdStockLineItemsRequest {
	return ApiGETShipmentIdStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETShipmentIdStockLineItemsExecute(r ApiGETShipmentIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETShipmentIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterToString(r.shipmentId, "")), -1)

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

type ApiGETStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
}

func (r ApiGETStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLineItemsExecute(r)
}

/*
GETStockLineItems List all stock line items

List all stock line items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETStockLineItems(ctx context.Context) ApiGETStockLineItemsRequest {
	return ApiGETStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETStockLineItemsExecute(r ApiGETStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items"

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

type ApiGETStockLineItemsStockLineItemIdRequest struct {
	ctx             context.Context
	ApiService      *StockLineItemsApiService
	stockLineItemId string
}

func (r ApiGETStockLineItemsStockLineItemIdRequest) Execute() (*StockLineItem, *http.Response, error) {
	return r.ApiService.GETStockLineItemsStockLineItemIdExecute(r)
}

/*
GETStockLineItemsStockLineItemId Retrieve a stock line item

Retrieve a stock line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockLineItemId The resource's id
 @return ApiGETStockLineItemsStockLineItemIdRequest
*/
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemId(ctx context.Context, stockLineItemId string) ApiGETStockLineItemsStockLineItemIdRequest {
	return ApiGETStockLineItemsStockLineItemIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
//  @return StockLineItem
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemIdExecute(r ApiGETStockLineItemsStockLineItemIdRequest) (*StockLineItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StockLineItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItemsStockLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterToString(r.stockLineItemId, "")), -1)

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

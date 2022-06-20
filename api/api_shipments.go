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

type ShipmentsApi interface {

	/*
		GETOrderIdShipments Retrieve the shipments associated to the order

		Retrieve the shipments associated to the order

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orderId The resource's id
		@return ApiGETOrderIdShipmentsRequest
	*/
	GETOrderIdShipments(ctx context.Context, orderId string) ApiGETOrderIdShipmentsRequest

	// GETOrderIdShipmentsExecute executes the request
	GETOrderIdShipmentsExecute(r ApiGETOrderIdShipmentsRequest) (*http.Response, error)

	/*
		GETParcelIdShipment Retrieve the shipment associated to the parcel

		Retrieve the shipment associated to the parcel

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parcelId The resource's id
		@return ApiGETParcelIdShipmentRequest
	*/
	GETParcelIdShipment(ctx context.Context, parcelId string) ApiGETParcelIdShipmentRequest

	// GETParcelIdShipmentExecute executes the request
	GETParcelIdShipmentExecute(r ApiGETParcelIdShipmentRequest) (*http.Response, error)

	/*
		GETShipments List all shipments

		List all shipments

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETShipmentsRequest
	*/
	GETShipments(ctx context.Context) ApiGETShipmentsRequest

	// GETShipmentsExecute executes the request
	GETShipmentsExecute(r ApiGETShipmentsRequest) (*http.Response, error)

	/*
		GETShipmentsShipmentId Retrieve a shipment

		Retrieve a shipment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param shipmentId The resource's id
		@return ApiGETShipmentsShipmentIdRequest
	*/
	GETShipmentsShipmentId(ctx context.Context, shipmentId string) ApiGETShipmentsShipmentIdRequest

	// GETShipmentsShipmentIdExecute executes the request
	//  @return Shipment
	GETShipmentsShipmentIdExecute(r ApiGETShipmentsShipmentIdRequest) (*Shipment, *http.Response, error)

	/*
		GETStockLineItemIdShipment Retrieve the shipment associated to the stock line item

		Retrieve the shipment associated to the stock line item

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stockLineItemId The resource's id
		@return ApiGETStockLineItemIdShipmentRequest
	*/
	GETStockLineItemIdShipment(ctx context.Context, stockLineItemId string) ApiGETStockLineItemIdShipmentRequest

	// GETStockLineItemIdShipmentExecute executes the request
	GETStockLineItemIdShipmentExecute(r ApiGETStockLineItemIdShipmentRequest) (*http.Response, error)

	/*
		GETStockTransferIdShipment Retrieve the shipment associated to the stock transfer

		Retrieve the shipment associated to the stock transfer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stockTransferId The resource's id
		@return ApiGETStockTransferIdShipmentRequest
	*/
	GETStockTransferIdShipment(ctx context.Context, stockTransferId string) ApiGETStockTransferIdShipmentRequest

	// GETStockTransferIdShipmentExecute executes the request
	GETStockTransferIdShipmentExecute(r ApiGETStockTransferIdShipmentRequest) (*http.Response, error)

	/*
		PATCHShipmentsShipmentId Update a shipment

		Update a shipment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param shipmentId The resource's id
		@return ApiPATCHShipmentsShipmentIdRequest
	*/
	PATCHShipmentsShipmentId(ctx context.Context, shipmentId string) ApiPATCHShipmentsShipmentIdRequest

	// PATCHShipmentsShipmentIdExecute executes the request
	PATCHShipmentsShipmentIdExecute(r ApiPATCHShipmentsShipmentIdRequest) (*http.Response, error)
}

// ShipmentsApiService ShipmentsApi service
type ShipmentsApiService service

type ApiGETOrderIdShipmentsRequest struct {
	ctx        context.Context
	ApiService ShipmentsApi
	orderId    string
}

func (r ApiGETOrderIdShipmentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdShipmentsExecute(r)
}

/*
GETOrderIdShipments Retrieve the shipments associated to the order

Retrieve the shipments associated to the order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderId The resource's id
 @return ApiGETOrderIdShipmentsRequest
*/
func (a *ShipmentsApiService) GETOrderIdShipments(ctx context.Context, orderId string) ApiGETOrderIdShipmentsRequest {
	return ApiGETOrderIdShipmentsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) GETOrderIdShipmentsExecute(r ApiGETOrderIdShipmentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETOrderIdShipments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/shipments"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterToString(r.orderId, "")), -1)

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

type ApiGETParcelIdShipmentRequest struct {
	ctx        context.Context
	ApiService ShipmentsApi
	parcelId   string
}

func (r ApiGETParcelIdShipmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelIdShipmentExecute(r)
}

/*
GETParcelIdShipment Retrieve the shipment associated to the parcel

Retrieve the shipment associated to the parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelId The resource's id
 @return ApiGETParcelIdShipmentRequest
*/
func (a *ShipmentsApiService) GETParcelIdShipment(ctx context.Context, parcelId string) ApiGETParcelIdShipmentRequest {
	return ApiGETParcelIdShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) GETParcelIdShipmentExecute(r ApiGETParcelIdShipmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETParcelIdShipment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}/shipment"
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

type ApiGETShipmentsRequest struct {
	ctx        context.Context
	ApiService ShipmentsApi
}

func (r ApiGETShipmentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentsExecute(r)
}

/*
GETShipments List all shipments

List all shipments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETShipmentsRequest
*/
func (a *ShipmentsApiService) GETShipments(ctx context.Context) ApiGETShipmentsRequest {
	return ApiGETShipmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) GETShipmentsExecute(r ApiGETShipmentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETShipments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments"

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

type ApiGETShipmentsShipmentIdRequest struct {
	ctx        context.Context
	ApiService ShipmentsApi
	shipmentId string
}

func (r ApiGETShipmentsShipmentIdRequest) Execute() (*Shipment, *http.Response, error) {
	return r.ApiService.GETShipmentsShipmentIdExecute(r)
}

/*
GETShipmentsShipmentId Retrieve a shipment

Retrieve a shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ApiGETShipmentsShipmentIdRequest
*/
func (a *ShipmentsApiService) GETShipmentsShipmentId(ctx context.Context, shipmentId string) ApiGETShipmentsShipmentIdRequest {
	return ApiGETShipmentsShipmentIdRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//  @return Shipment
func (a *ShipmentsApiService) GETShipmentsShipmentIdExecute(r ApiGETShipmentsShipmentIdRequest) (*Shipment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Shipment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETShipmentsShipmentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}"
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

type ApiGETStockLineItemIdShipmentRequest struct {
	ctx             context.Context
	ApiService      ShipmentsApi
	stockLineItemId string
}

func (r ApiGETStockLineItemIdShipmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLineItemIdShipmentExecute(r)
}

/*
GETStockLineItemIdShipment Retrieve the shipment associated to the stock line item

Retrieve the shipment associated to the stock line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockLineItemId The resource's id
 @return ApiGETStockLineItemIdShipmentRequest
*/
func (a *ShipmentsApiService) GETStockLineItemIdShipment(ctx context.Context, stockLineItemId string) ApiGETStockLineItemIdShipmentRequest {
	return ApiGETStockLineItemIdShipmentRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) GETStockLineItemIdShipmentExecute(r ApiGETStockLineItemIdShipmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETStockLineItemIdShipment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}/shipment"
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

type ApiGETStockTransferIdShipmentRequest struct {
	ctx             context.Context
	ApiService      ShipmentsApi
	stockTransferId string
}

func (r ApiGETStockTransferIdShipmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockTransferIdShipmentExecute(r)
}

/*
GETStockTransferIdShipment Retrieve the shipment associated to the stock transfer

Retrieve the shipment associated to the stock transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockTransferId The resource's id
 @return ApiGETStockTransferIdShipmentRequest
*/
func (a *ShipmentsApiService) GETStockTransferIdShipment(ctx context.Context, stockTransferId string) ApiGETStockTransferIdShipmentRequest {
	return ApiGETStockTransferIdShipmentRequest{
		ApiService:      a,
		ctx:             ctx,
		stockTransferId: stockTransferId,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) GETStockTransferIdShipmentExecute(r ApiGETStockTransferIdShipmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.GETStockTransferIdShipment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers/{stockTransferId}/shipment"
	localVarPath = strings.Replace(localVarPath, "{"+"stockTransferId"+"}", url.PathEscape(parameterToString(r.stockTransferId, "")), -1)

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

type ApiPATCHShipmentsShipmentIdRequest struct {
	ctx            context.Context
	ApiService     ShipmentsApi
	shipmentId     string
	shipmentUpdate *ShipmentUpdate
}

func (r ApiPATCHShipmentsShipmentIdRequest) ShipmentUpdate(shipmentUpdate ShipmentUpdate) ApiPATCHShipmentsShipmentIdRequest {
	r.shipmentUpdate = &shipmentUpdate
	return r
}

func (r ApiPATCHShipmentsShipmentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHShipmentsShipmentIdExecute(r)
}

/*
PATCHShipmentsShipmentId Update a shipment

Update a shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ApiPATCHShipmentsShipmentIdRequest
*/
func (a *ShipmentsApiService) PATCHShipmentsShipmentId(ctx context.Context, shipmentId string) ApiPATCHShipmentsShipmentIdRequest {
	return ApiPATCHShipmentsShipmentIdRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *ShipmentsApiService) PATCHShipmentsShipmentIdExecute(r ApiPATCHShipmentsShipmentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsApiService.PATCHShipmentsShipmentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterToString(r.shipmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shipmentUpdate == nil {
		return nil, reportError("shipmentUpdate is required and must be specified")
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
	localVarPostBody = r.shipmentUpdate
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

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

// ShippingMethodsApiService ShippingMethodsApi service
type ShippingMethodsApiService service

type ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest struct {
	ctx              context.Context
	ApiService       *ShippingMethodsApiService
	shippingMethodId string
}

func (r ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEShippingMethodsShippingMethodIdExecute(r)
}

/*
DELETEShippingMethodsShippingMethodId Delete a shipping method

Delete a shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest
*/
func (a *ShippingMethodsApiService) DELETEShippingMethodsShippingMethodId(ctx context.Context, shippingMethodId string) ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest {
	return ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) DELETEShippingMethodsShippingMethodIdExecute(r ShippingMethodsApiDELETEShippingMethodsShippingMethodIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.DELETEShippingMethodsShippingMethodId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

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

type ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest struct {
	ctx                context.Context
	ApiService         *ShippingMethodsApiService
	deliveryLeadTimeId string
}

func (r ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETDeliveryLeadTimeIdShippingMethodExecute(r)
}

/*
GETDeliveryLeadTimeIdShippingMethod Retrieve the shipping method associated to the delivery lead time

Retrieve the shipping method associated to the delivery lead time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deliveryLeadTimeId The resource's id
 @return ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest
*/
func (a *ShippingMethodsApiService) GETDeliveryLeadTimeIdShippingMethod(ctx context.Context, deliveryLeadTimeId string) ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest {
	return ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest{
		ApiService:         a,
		ctx:                ctx,
		deliveryLeadTimeId: deliveryLeadTimeId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) GETDeliveryLeadTimeIdShippingMethodExecute(r ShippingMethodsApiGETDeliveryLeadTimeIdShippingMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETDeliveryLeadTimeIdShippingMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times/{deliveryLeadTimeId}/shipping_method"
	localVarPath = strings.Replace(localVarPath, "{"+"deliveryLeadTimeId"+"}", url.PathEscape(parameterToString(r.deliveryLeadTimeId, "")), -1)

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

type ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest struct {
	ctx        context.Context
	ApiService *ShippingMethodsApiService
	shipmentId string
}

func (r ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdAvailableShippingMethodsExecute(r)
}

/*
GETShipmentIdAvailableShippingMethods Retrieve the available shipping methods associated to the shipment

Retrieve the available shipping methods associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest
*/
func (a *ShippingMethodsApiService) GETShipmentIdAvailableShippingMethods(ctx context.Context, shipmentId string) ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest {
	return ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) GETShipmentIdAvailableShippingMethodsExecute(r ShippingMethodsApiGETShipmentIdAvailableShippingMethodsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShipmentIdAvailableShippingMethods")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/available_shipping_methods"
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

type ShippingMethodsApiGETShipmentIdShippingMethodRequest struct {
	ctx        context.Context
	ApiService *ShippingMethodsApiService
	shipmentId string
}

func (r ShippingMethodsApiGETShipmentIdShippingMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdShippingMethodExecute(r)
}

/*
GETShipmentIdShippingMethod Retrieve the shipping method associated to the shipment

Retrieve the shipping method associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ShippingMethodsApiGETShipmentIdShippingMethodRequest
*/
func (a *ShippingMethodsApiService) GETShipmentIdShippingMethod(ctx context.Context, shipmentId string) ShippingMethodsApiGETShipmentIdShippingMethodRequest {
	return ShippingMethodsApiGETShipmentIdShippingMethodRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) GETShipmentIdShippingMethodExecute(r ShippingMethodsApiGETShipmentIdShippingMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShipmentIdShippingMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/shipping_method"
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

type ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest struct {
	ctx                  context.Context
	ApiService           *ShippingMethodsApiService
	shippingMethodTierId string
}

func (r ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodTierIdShippingMethodExecute(r)
}

/*
GETShippingMethodTierIdShippingMethod Retrieve the shipping method associated to the shipping method tier

Retrieve the shipping method associated to the shipping method tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodTierId The resource's id
 @return ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest
*/
func (a *ShippingMethodsApiService) GETShippingMethodTierIdShippingMethod(ctx context.Context, shippingMethodTierId string) ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest {
	return ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingMethodTierId: shippingMethodTierId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) GETShippingMethodTierIdShippingMethodExecute(r ShippingMethodsApiGETShippingMethodTierIdShippingMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShippingMethodTierIdShippingMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_method_tiers/{shippingMethodTierId}/shipping_method"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodTierId"+"}", url.PathEscape(parameterToString(r.shippingMethodTierId, "")), -1)

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

type ShippingMethodsApiGETShippingMethodsRequest struct {
	ctx        context.Context
	ApiService *ShippingMethodsApiService
}

func (r ShippingMethodsApiGETShippingMethodsRequest) Execute() (*GETShippingMethods200Response, *http.Response, error) {
	return r.ApiService.GETShippingMethodsExecute(r)
}

/*
GETShippingMethods List all shipping methods

List all shipping methods

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShippingMethodsApiGETShippingMethodsRequest
*/
func (a *ShippingMethodsApiService) GETShippingMethods(ctx context.Context) ShippingMethodsApiGETShippingMethodsRequest {
	return ShippingMethodsApiGETShippingMethodsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETShippingMethods200Response
func (a *ShippingMethodsApiService) GETShippingMethodsExecute(r ShippingMethodsApiGETShippingMethodsRequest) (*GETShippingMethods200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingMethods200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShippingMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods"

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

type ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest struct {
	ctx              context.Context
	ApiService       *ShippingMethodsApiService
	shippingMethodId string
}

func (r ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest) Execute() (*GETShippingMethodsShippingMethodId200Response, *http.Response, error) {
	return r.ApiService.GETShippingMethodsShippingMethodIdExecute(r)
}

/*
GETShippingMethodsShippingMethodId Retrieve a shipping method

Retrieve a shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest
*/
func (a *ShippingMethodsApiService) GETShippingMethodsShippingMethodId(ctx context.Context, shippingMethodId string) ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest {
	return ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
//  @return GETShippingMethodsShippingMethodId200Response
func (a *ShippingMethodsApiService) GETShippingMethodsShippingMethodIdExecute(r ShippingMethodsApiGETShippingMethodsShippingMethodIdRequest) (*GETShippingMethodsShippingMethodId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingMethodsShippingMethodId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShippingMethodsShippingMethodId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

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

type ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest struct {
	ctx                  context.Context
	ApiService           *ShippingMethodsApiService
	shippingWeightTierId string
}

func (r ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingWeightTierIdShippingMethodExecute(r)
}

/*
GETShippingWeightTierIdShippingMethod Retrieve the shipping method associated to the shipping weight tier

Retrieve the shipping method associated to the shipping weight tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingWeightTierId The resource's id
 @return ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest
*/
func (a *ShippingMethodsApiService) GETShippingWeightTierIdShippingMethod(ctx context.Context, shippingWeightTierId string) ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest {
	return ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
func (a *ShippingMethodsApiService) GETShippingWeightTierIdShippingMethodExecute(r ShippingMethodsApiGETShippingWeightTierIdShippingMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.GETShippingWeightTierIdShippingMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}/shipping_method"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterToString(r.shippingWeightTierId, "")), -1)

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

type ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingMethodsApiService
	shippingMethodUpdate *ShippingMethodUpdate
	shippingMethodId     string
}

func (r ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest) ShippingMethodUpdate(shippingMethodUpdate ShippingMethodUpdate) ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest {
	r.shippingMethodUpdate = &shippingMethodUpdate
	return r
}

func (r ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest) Execute() (*PATCHShippingMethodsShippingMethodId200Response, *http.Response, error) {
	return r.ApiService.PATCHShippingMethodsShippingMethodIdExecute(r)
}

/*
PATCHShippingMethodsShippingMethodId Update a shipping method

Update a shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest
*/
func (a *ShippingMethodsApiService) PATCHShippingMethodsShippingMethodId(ctx context.Context, shippingMethodId string) ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest {
	return ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
//  @return PATCHShippingMethodsShippingMethodId200Response
func (a *ShippingMethodsApiService) PATCHShippingMethodsShippingMethodIdExecute(r ShippingMethodsApiPATCHShippingMethodsShippingMethodIdRequest) (*PATCHShippingMethodsShippingMethodId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHShippingMethodsShippingMethodId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.PATCHShippingMethodsShippingMethodId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingMethodUpdate == nil {
		return localVarReturnValue, nil, reportError("shippingMethodUpdate is required and must be specified")
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
	localVarPostBody = r.shippingMethodUpdate
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

type ShippingMethodsApiPOSTShippingMethodsRequest struct {
	ctx                  context.Context
	ApiService           *ShippingMethodsApiService
	shippingMethodCreate *ShippingMethodCreate
}

func (r ShippingMethodsApiPOSTShippingMethodsRequest) ShippingMethodCreate(shippingMethodCreate ShippingMethodCreate) ShippingMethodsApiPOSTShippingMethodsRequest {
	r.shippingMethodCreate = &shippingMethodCreate
	return r
}

func (r ShippingMethodsApiPOSTShippingMethodsRequest) Execute() (*POSTShippingMethods201Response, *http.Response, error) {
	return r.ApiService.POSTShippingMethodsExecute(r)
}

/*
POSTShippingMethods Create a shipping method

Create a shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShippingMethodsApiPOSTShippingMethodsRequest
*/
func (a *ShippingMethodsApiService) POSTShippingMethods(ctx context.Context) ShippingMethodsApiPOSTShippingMethodsRequest {
	return ShippingMethodsApiPOSTShippingMethodsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTShippingMethods201Response
func (a *ShippingMethodsApiService) POSTShippingMethodsExecute(r ShippingMethodsApiPOSTShippingMethodsRequest) (*POSTShippingMethods201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTShippingMethods201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodsApiService.POSTShippingMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingMethodCreate == nil {
		return localVarReturnValue, nil, reportError("shippingMethodCreate is required and must be specified")
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
	localVarPostBody = r.shippingMethodCreate
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

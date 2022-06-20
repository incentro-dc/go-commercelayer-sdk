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

type CarrierAccountsApi interface {

	/*
		GETCarrierAccounts List all carrier accounts

		List all carrier accounts

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCarrierAccountsRequest
	*/
	GETCarrierAccounts(ctx context.Context) ApiGETCarrierAccountsRequest

	// GETCarrierAccountsExecute executes the request
	GETCarrierAccountsExecute(r ApiGETCarrierAccountsRequest) (*http.Response, error)

	/*
		GETCarrierAccountsCarrierAccountId Retrieve a carrier account

		Retrieve a carrier account

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param carrierAccountId The resource's id
		@return ApiGETCarrierAccountsCarrierAccountIdRequest
	*/
	GETCarrierAccountsCarrierAccountId(ctx context.Context, carrierAccountId string) ApiGETCarrierAccountsCarrierAccountIdRequest

	// GETCarrierAccountsCarrierAccountIdExecute executes the request
	//  @return CarrierAccount
	GETCarrierAccountsCarrierAccountIdExecute(r ApiGETCarrierAccountsCarrierAccountIdRequest) (*CarrierAccount, *http.Response, error)

	/*
		GETShipmentIdCarrierAccounts Retrieve the carrier accounts associated to the shipment

		Retrieve the carrier accounts associated to the shipment

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param shipmentId The resource's id
		@return ApiGETShipmentIdCarrierAccountsRequest
	*/
	GETShipmentIdCarrierAccounts(ctx context.Context, shipmentId string) ApiGETShipmentIdCarrierAccountsRequest

	// GETShipmentIdCarrierAccountsExecute executes the request
	GETShipmentIdCarrierAccountsExecute(r ApiGETShipmentIdCarrierAccountsRequest) (*http.Response, error)
}

// CarrierAccountsApiService CarrierAccountsApi service
type CarrierAccountsApiService service

type ApiGETCarrierAccountsRequest struct {
	ctx        context.Context
	ApiService CarrierAccountsApi
}

func (r ApiGETCarrierAccountsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCarrierAccountsExecute(r)
}

/*
GETCarrierAccounts List all carrier accounts

List all carrier accounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCarrierAccountsRequest
*/
func (a *CarrierAccountsApiService) GETCarrierAccounts(ctx context.Context) ApiGETCarrierAccountsRequest {
	return ApiGETCarrierAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CarrierAccountsApiService) GETCarrierAccountsExecute(r ApiGETCarrierAccountsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarrierAccountsApiService.GETCarrierAccounts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/carrier_accounts"

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

type ApiGETCarrierAccountsCarrierAccountIdRequest struct {
	ctx              context.Context
	ApiService       CarrierAccountsApi
	carrierAccountId string
}

func (r ApiGETCarrierAccountsCarrierAccountIdRequest) Execute() (*CarrierAccount, *http.Response, error) {
	return r.ApiService.GETCarrierAccountsCarrierAccountIdExecute(r)
}

/*
GETCarrierAccountsCarrierAccountId Retrieve a carrier account

Retrieve a carrier account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param carrierAccountId The resource's id
 @return ApiGETCarrierAccountsCarrierAccountIdRequest
*/
func (a *CarrierAccountsApiService) GETCarrierAccountsCarrierAccountId(ctx context.Context, carrierAccountId string) ApiGETCarrierAccountsCarrierAccountIdRequest {
	return ApiGETCarrierAccountsCarrierAccountIdRequest{
		ApiService:       a,
		ctx:              ctx,
		carrierAccountId: carrierAccountId,
	}
}

// Execute executes the request
//  @return CarrierAccount
func (a *CarrierAccountsApiService) GETCarrierAccountsCarrierAccountIdExecute(r ApiGETCarrierAccountsCarrierAccountIdRequest) (*CarrierAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CarrierAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarrierAccountsApiService.GETCarrierAccountsCarrierAccountId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/carrier_accounts/{carrierAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"carrierAccountId"+"}", url.PathEscape(parameterToString(r.carrierAccountId, "")), -1)

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

type ApiGETShipmentIdCarrierAccountsRequest struct {
	ctx        context.Context
	ApiService CarrierAccountsApi
	shipmentId string
}

func (r ApiGETShipmentIdCarrierAccountsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdCarrierAccountsExecute(r)
}

/*
GETShipmentIdCarrierAccounts Retrieve the carrier accounts associated to the shipment

Retrieve the carrier accounts associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ApiGETShipmentIdCarrierAccountsRequest
*/
func (a *CarrierAccountsApiService) GETShipmentIdCarrierAccounts(ctx context.Context, shipmentId string) ApiGETShipmentIdCarrierAccountsRequest {
	return ApiGETShipmentIdCarrierAccountsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *CarrierAccountsApiService) GETShipmentIdCarrierAccountsExecute(r ApiGETShipmentIdCarrierAccountsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarrierAccountsApiService.GETShipmentIdCarrierAccounts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/carrier_accounts"
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

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

// PriceVolumeTiersApiService PriceVolumeTiersApi service
type PriceVolumeTiersApiService service

type PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest struct {
	ctx               context.Context
	ApiService        *PriceVolumeTiersApiService
	priceVolumeTierId string
}

func (r PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPriceVolumeTiersPriceVolumeTierIdExecute(r)
}

/*
DELETEPriceVolumeTiersPriceVolumeTierId Delete a price volume tier

Delete a price volume tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceVolumeTierId The resource's id
 @return PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest
*/
func (a *PriceVolumeTiersApiService) DELETEPriceVolumeTiersPriceVolumeTierId(ctx context.Context, priceVolumeTierId string) PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest {
	return PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest{
		ApiService:        a,
		ctx:               ctx,
		priceVolumeTierId: priceVolumeTierId,
	}
}

// Execute executes the request
func (a *PriceVolumeTiersApiService) DELETEPriceVolumeTiersPriceVolumeTierIdExecute(r PriceVolumeTiersApiDELETEPriceVolumeTiersPriceVolumeTierIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.DELETEPriceVolumeTiersPriceVolumeTierId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_volume_tiers/{priceVolumeTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceVolumeTierId"+"}", url.PathEscape(parameterToString(r.priceVolumeTierId, "")), -1)

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

type PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest struct {
	ctx        context.Context
	ApiService *PriceVolumeTiersApiService
	priceId    string
}

func (r PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceIdPriceVolumeTiersExecute(r)
}

/*
GETPriceIdPriceVolumeTiers Retrieve the price volume tiers associated to the price

Retrieve the price volume tiers associated to the price

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceId The resource's id
 @return PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest
*/
func (a *PriceVolumeTiersApiService) GETPriceIdPriceVolumeTiers(ctx context.Context, priceId string) PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest {
	return PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PriceVolumeTiersApiService) GETPriceIdPriceVolumeTiersExecute(r PriceVolumeTiersApiGETPriceIdPriceVolumeTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.GETPriceIdPriceVolumeTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}/price_volume_tiers"
	localVarPath = strings.Replace(localVarPath, "{"+"priceId"+"}", url.PathEscape(parameterToString(r.priceId, "")), -1)

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

type PriceVolumeTiersApiGETPriceVolumeTiersRequest struct {
	ctx        context.Context
	ApiService *PriceVolumeTiersApiService
}

func (r PriceVolumeTiersApiGETPriceVolumeTiersRequest) Execute() (*GETPriceVolumeTiers200Response, *http.Response, error) {
	return r.ApiService.GETPriceVolumeTiersExecute(r)
}

/*
GETPriceVolumeTiers List all price volume tiers

List all price volume tiers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PriceVolumeTiersApiGETPriceVolumeTiersRequest
*/
func (a *PriceVolumeTiersApiService) GETPriceVolumeTiers(ctx context.Context) PriceVolumeTiersApiGETPriceVolumeTiersRequest {
	return PriceVolumeTiersApiGETPriceVolumeTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETPriceVolumeTiers200Response
func (a *PriceVolumeTiersApiService) GETPriceVolumeTiersExecute(r PriceVolumeTiersApiGETPriceVolumeTiersRequest) (*GETPriceVolumeTiers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceVolumeTiers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.GETPriceVolumeTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_volume_tiers"

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

type PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest struct {
	ctx               context.Context
	ApiService        *PriceVolumeTiersApiService
	priceVolumeTierId string
}

func (r PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest) Execute() (*GETPriceVolumeTiersPriceVolumeTierId200Response, *http.Response, error) {
	return r.ApiService.GETPriceVolumeTiersPriceVolumeTierIdExecute(r)
}

/*
GETPriceVolumeTiersPriceVolumeTierId Retrieve a price volume tier

Retrieve a price volume tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceVolumeTierId The resource's id
 @return PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest
*/
func (a *PriceVolumeTiersApiService) GETPriceVolumeTiersPriceVolumeTierId(ctx context.Context, priceVolumeTierId string) PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest {
	return PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest{
		ApiService:        a,
		ctx:               ctx,
		priceVolumeTierId: priceVolumeTierId,
	}
}

// Execute executes the request
//  @return GETPriceVolumeTiersPriceVolumeTierId200Response
func (a *PriceVolumeTiersApiService) GETPriceVolumeTiersPriceVolumeTierIdExecute(r PriceVolumeTiersApiGETPriceVolumeTiersPriceVolumeTierIdRequest) (*GETPriceVolumeTiersPriceVolumeTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceVolumeTiersPriceVolumeTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.GETPriceVolumeTiersPriceVolumeTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_volume_tiers/{priceVolumeTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceVolumeTierId"+"}", url.PathEscape(parameterToString(r.priceVolumeTierId, "")), -1)

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

type PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest struct {
	ctx                   context.Context
	ApiService            *PriceVolumeTiersApiService
	priceVolumeTierUpdate *PriceVolumeTierUpdate
	priceVolumeTierId     string
}

func (r PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest) PriceVolumeTierUpdate(priceVolumeTierUpdate PriceVolumeTierUpdate) PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest {
	r.priceVolumeTierUpdate = &priceVolumeTierUpdate
	return r
}

func (r PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest) Execute() (*PATCHPriceVolumeTiersPriceVolumeTierId200Response, *http.Response, error) {
	return r.ApiService.PATCHPriceVolumeTiersPriceVolumeTierIdExecute(r)
}

/*
PATCHPriceVolumeTiersPriceVolumeTierId Update a price volume tier

Update a price volume tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceVolumeTierId The resource's id
 @return PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest
*/
func (a *PriceVolumeTiersApiService) PATCHPriceVolumeTiersPriceVolumeTierId(ctx context.Context, priceVolumeTierId string) PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest {
	return PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest{
		ApiService:        a,
		ctx:               ctx,
		priceVolumeTierId: priceVolumeTierId,
	}
}

// Execute executes the request
//  @return PATCHPriceVolumeTiersPriceVolumeTierId200Response
func (a *PriceVolumeTiersApiService) PATCHPriceVolumeTiersPriceVolumeTierIdExecute(r PriceVolumeTiersApiPATCHPriceVolumeTiersPriceVolumeTierIdRequest) (*PATCHPriceVolumeTiersPriceVolumeTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPriceVolumeTiersPriceVolumeTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.PATCHPriceVolumeTiersPriceVolumeTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_volume_tiers/{priceVolumeTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceVolumeTierId"+"}", url.PathEscape(parameterToString(r.priceVolumeTierId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceVolumeTierUpdate == nil {
		return localVarReturnValue, nil, reportError("priceVolumeTierUpdate is required and must be specified")
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
	localVarPostBody = r.priceVolumeTierUpdate
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

type PriceVolumeTiersApiPOSTPriceVolumeTiersRequest struct {
	ctx                   context.Context
	ApiService            *PriceVolumeTiersApiService
	priceVolumeTierCreate *PriceVolumeTierCreate
}

func (r PriceVolumeTiersApiPOSTPriceVolumeTiersRequest) PriceVolumeTierCreate(priceVolumeTierCreate PriceVolumeTierCreate) PriceVolumeTiersApiPOSTPriceVolumeTiersRequest {
	r.priceVolumeTierCreate = &priceVolumeTierCreate
	return r
}

func (r PriceVolumeTiersApiPOSTPriceVolumeTiersRequest) Execute() (*POSTPriceVolumeTiers201Response, *http.Response, error) {
	return r.ApiService.POSTPriceVolumeTiersExecute(r)
}

/*
POSTPriceVolumeTiers Create a price volume tier

Create a price volume tier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PriceVolumeTiersApiPOSTPriceVolumeTiersRequest
*/
func (a *PriceVolumeTiersApiService) POSTPriceVolumeTiers(ctx context.Context) PriceVolumeTiersApiPOSTPriceVolumeTiersRequest {
	return PriceVolumeTiersApiPOSTPriceVolumeTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTPriceVolumeTiers201Response
func (a *PriceVolumeTiersApiService) POSTPriceVolumeTiersExecute(r PriceVolumeTiersApiPOSTPriceVolumeTiersRequest) (*POSTPriceVolumeTiers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPriceVolumeTiers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceVolumeTiersApiService.POSTPriceVolumeTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_volume_tiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceVolumeTierCreate == nil {
		return localVarReturnValue, nil, reportError("priceVolumeTierCreate is required and must be specified")
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
	localVarPostBody = r.priceVolumeTierCreate
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

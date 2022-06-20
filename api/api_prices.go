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

type PricesApi interface {

	/*
		DELETEPricesPriceId Delete a price

		Delete a price

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param priceId The resource's id
		@return ApiDELETEPricesPriceIdRequest
	*/
	DELETEPricesPriceId(ctx context.Context, priceId string) ApiDELETEPricesPriceIdRequest

	// DELETEPricesPriceIdExecute executes the request
	DELETEPricesPriceIdExecute(r ApiDELETEPricesPriceIdRequest) (*http.Response, error)

	/*
		GETPriceListIdPrices Retrieve the prices associated to the price list

		Retrieve the prices associated to the price list

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param priceListId The resource's id
		@return ApiGETPriceListIdPricesRequest
	*/
	GETPriceListIdPrices(ctx context.Context, priceListId string) ApiGETPriceListIdPricesRequest

	// GETPriceListIdPricesExecute executes the request
	GETPriceListIdPricesExecute(r ApiGETPriceListIdPricesRequest) (*http.Response, error)

	/*
		GETPrices List all prices

		List all prices

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETPricesRequest
	*/
	GETPrices(ctx context.Context) ApiGETPricesRequest

	// GETPricesExecute executes the request
	GETPricesExecute(r ApiGETPricesRequest) (*http.Response, error)

	/*
		GETPricesPriceId Retrieve a price

		Retrieve a price

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param priceId The resource's id
		@return ApiGETPricesPriceIdRequest
	*/
	GETPricesPriceId(ctx context.Context, priceId string) ApiGETPricesPriceIdRequest

	// GETPricesPriceIdExecute executes the request
	//  @return Price
	GETPricesPriceIdExecute(r ApiGETPricesPriceIdRequest) (*Price, *http.Response, error)

	/*
		GETSkuIdPrices Retrieve the prices associated to the SKU

		Retrieve the prices associated to the SKU

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param skuId The resource's id
		@return ApiGETSkuIdPricesRequest
	*/
	GETSkuIdPrices(ctx context.Context, skuId string) ApiGETSkuIdPricesRequest

	// GETSkuIdPricesExecute executes the request
	GETSkuIdPricesExecute(r ApiGETSkuIdPricesRequest) (*http.Response, error)

	/*
		PATCHPricesPriceId Update a price

		Update a price

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param priceId The resource's id
		@return ApiPATCHPricesPriceIdRequest
	*/
	PATCHPricesPriceId(ctx context.Context, priceId string) ApiPATCHPricesPriceIdRequest

	// PATCHPricesPriceIdExecute executes the request
	PATCHPricesPriceIdExecute(r ApiPATCHPricesPriceIdRequest) (*http.Response, error)

	/*
		POSTPrices Create a price

		Create a price

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTPricesRequest
	*/
	POSTPrices(ctx context.Context) ApiPOSTPricesRequest

	// POSTPricesExecute executes the request
	POSTPricesExecute(r ApiPOSTPricesRequest) (*http.Response, error)
}

// PricesApiService PricesApi service
type PricesApiService service

type ApiDELETEPricesPriceIdRequest struct {
	ctx        context.Context
	ApiService PricesApi
	priceId    string
}

func (r ApiDELETEPricesPriceIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPricesPriceIdExecute(r)
}

/*
DELETEPricesPriceId Delete a price

Delete a price

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceId The resource's id
 @return ApiDELETEPricesPriceIdRequest
*/
func (a *PricesApiService) DELETEPricesPriceId(ctx context.Context, priceId string) ApiDELETEPricesPriceIdRequest {
	return ApiDELETEPricesPriceIdRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PricesApiService) DELETEPricesPriceIdExecute(r ApiDELETEPricesPriceIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.DELETEPricesPriceId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}"
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

type ApiGETPriceListIdPricesRequest struct {
	ctx         context.Context
	ApiService  PricesApi
	priceListId string
}

func (r ApiGETPriceListIdPricesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceListIdPricesExecute(r)
}

/*
GETPriceListIdPrices Retrieve the prices associated to the price list

Retrieve the prices associated to the price list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceListId The resource's id
 @return ApiGETPriceListIdPricesRequest
*/
func (a *PricesApiService) GETPriceListIdPrices(ctx context.Context, priceListId string) ApiGETPriceListIdPricesRequest {
	return ApiGETPriceListIdPricesRequest{
		ApiService:  a,
		ctx:         ctx,
		priceListId: priceListId,
	}
}

// Execute executes the request
func (a *PricesApiService) GETPriceListIdPricesExecute(r ApiGETPriceListIdPricesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.GETPriceListIdPrices")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists/{priceListId}/prices"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListId"+"}", url.PathEscape(parameterToString(r.priceListId, "")), -1)

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

type ApiGETPricesRequest struct {
	ctx        context.Context
	ApiService PricesApi
}

func (r ApiGETPricesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPricesExecute(r)
}

/*
GETPrices List all prices

List all prices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETPricesRequest
*/
func (a *PricesApiService) GETPrices(ctx context.Context) ApiGETPricesRequest {
	return ApiGETPricesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PricesApiService) GETPricesExecute(r ApiGETPricesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.GETPrices")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices"

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

type ApiGETPricesPriceIdRequest struct {
	ctx        context.Context
	ApiService PricesApi
	priceId    string
}

func (r ApiGETPricesPriceIdRequest) Execute() (*Price, *http.Response, error) {
	return r.ApiService.GETPricesPriceIdExecute(r)
}

/*
GETPricesPriceId Retrieve a price

Retrieve a price

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceId The resource's id
 @return ApiGETPricesPriceIdRequest
*/
func (a *PricesApiService) GETPricesPriceId(ctx context.Context, priceId string) ApiGETPricesPriceIdRequest {
	return ApiGETPricesPriceIdRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
//  @return Price
func (a *PricesApiService) GETPricesPriceIdExecute(r ApiGETPricesPriceIdRequest) (*Price, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Price
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.GETPricesPriceId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}"
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

type ApiGETSkuIdPricesRequest struct {
	ctx        context.Context
	ApiService PricesApi
	skuId      string
}

func (r ApiGETSkuIdPricesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuIdPricesExecute(r)
}

/*
GETSkuIdPrices Retrieve the prices associated to the SKU

Retrieve the prices associated to the SKU

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skuId The resource's id
 @return ApiGETSkuIdPricesRequest
*/
func (a *PricesApiService) GETSkuIdPrices(ctx context.Context, skuId string) ApiGETSkuIdPricesRequest {
	return ApiGETSkuIdPricesRequest{
		ApiService: a,
		ctx:        ctx,
		skuId:      skuId,
	}
}

// Execute executes the request
func (a *PricesApiService) GETSkuIdPricesExecute(r ApiGETSkuIdPricesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.GETSkuIdPrices")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skus/{skuId}/prices"
	localVarPath = strings.Replace(localVarPath, "{"+"skuId"+"}", url.PathEscape(parameterToString(r.skuId, "")), -1)

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

type ApiPATCHPricesPriceIdRequest struct {
	ctx         context.Context
	ApiService  PricesApi
	priceId     string
	priceUpdate *PriceUpdate
}

func (r ApiPATCHPricesPriceIdRequest) PriceUpdate(priceUpdate PriceUpdate) ApiPATCHPricesPriceIdRequest {
	r.priceUpdate = &priceUpdate
	return r
}

func (r ApiPATCHPricesPriceIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHPricesPriceIdExecute(r)
}

/*
PATCHPricesPriceId Update a price

Update a price

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param priceId The resource's id
 @return ApiPATCHPricesPriceIdRequest
*/
func (a *PricesApiService) PATCHPricesPriceId(ctx context.Context, priceId string) ApiPATCHPricesPriceIdRequest {
	return ApiPATCHPricesPriceIdRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PricesApiService) PATCHPricesPriceIdExecute(r ApiPATCHPricesPriceIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.PATCHPricesPriceId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceId"+"}", url.PathEscape(parameterToString(r.priceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceUpdate == nil {
		return nil, reportError("priceUpdate is required and must be specified")
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
	localVarPostBody = r.priceUpdate
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

type ApiPOSTPricesRequest struct {
	ctx         context.Context
	ApiService  PricesApi
	priceCreate *PriceCreate
}

func (r ApiPOSTPricesRequest) PriceCreate(priceCreate PriceCreate) ApiPOSTPricesRequest {
	r.priceCreate = &priceCreate
	return r
}

func (r ApiPOSTPricesRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTPricesExecute(r)
}

/*
POSTPrices Create a price

Create a price

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTPricesRequest
*/
func (a *PricesApiService) POSTPrices(ctx context.Context) ApiPOSTPricesRequest {
	return ApiPOSTPricesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PricesApiService) POSTPricesExecute(r ApiPOSTPricesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesApiService.POSTPrices")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceCreate == nil {
		return nil, reportError("priceCreate is required and must be specified")
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
	localVarPostBody = r.priceCreate
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
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

type BundlesApi interface {

	/*
		DELETEBundlesBundleId Delete a bundle

		Delete a bundle

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bundleId The resource's id
		@return ApiDELETEBundlesBundleIdRequest
	*/
	DELETEBundlesBundleId(ctx context.Context, bundleId string) ApiDELETEBundlesBundleIdRequest

	// DELETEBundlesBundleIdExecute executes the request
	DELETEBundlesBundleIdExecute(r ApiDELETEBundlesBundleIdRequest) (*http.Response, error)

	/*
		GETBundles List all bundles

		List all bundles

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETBundlesRequest
	*/
	GETBundles(ctx context.Context) ApiGETBundlesRequest

	// GETBundlesExecute executes the request
	GETBundlesExecute(r ApiGETBundlesRequest) (*http.Response, error)

	/*
		GETBundlesBundleId Retrieve a bundle

		Retrieve a bundle

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bundleId The resource's id
		@return ApiGETBundlesBundleIdRequest
	*/
	GETBundlesBundleId(ctx context.Context, bundleId string) ApiGETBundlesBundleIdRequest

	// GETBundlesBundleIdExecute executes the request
	//  @return Bundle
	GETBundlesBundleIdExecute(r ApiGETBundlesBundleIdRequest) (*Bundle, *http.Response, error)

	/*
		GETOrderIdAvailableFreeBundles Retrieve the available free bundles associated to the order

		Retrieve the available free bundles associated to the order

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orderId The resource's id
		@return ApiGETOrderIdAvailableFreeBundlesRequest
	*/
	GETOrderIdAvailableFreeBundles(ctx context.Context, orderId string) ApiGETOrderIdAvailableFreeBundlesRequest

	// GETOrderIdAvailableFreeBundlesExecute executes the request
	GETOrderIdAvailableFreeBundlesExecute(r ApiGETOrderIdAvailableFreeBundlesRequest) (*http.Response, error)

	/*
		GETSkuListIdBundles Retrieve the bundles associated to the SKU list

		Retrieve the bundles associated to the SKU list

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param skuListId The resource's id
		@return ApiGETSkuListIdBundlesRequest
	*/
	GETSkuListIdBundles(ctx context.Context, skuListId string) ApiGETSkuListIdBundlesRequest

	// GETSkuListIdBundlesExecute executes the request
	GETSkuListIdBundlesExecute(r ApiGETSkuListIdBundlesRequest) (*http.Response, error)

	/*
		PATCHBundlesBundleId Update a bundle

		Update a bundle

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bundleId The resource's id
		@return ApiPATCHBundlesBundleIdRequest
	*/
	PATCHBundlesBundleId(ctx context.Context, bundleId string) ApiPATCHBundlesBundleIdRequest

	// PATCHBundlesBundleIdExecute executes the request
	PATCHBundlesBundleIdExecute(r ApiPATCHBundlesBundleIdRequest) (*http.Response, error)

	/*
		POSTBundles Create a bundle

		Create a bundle

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTBundlesRequest
	*/
	POSTBundles(ctx context.Context) ApiPOSTBundlesRequest

	// POSTBundlesExecute executes the request
	POSTBundlesExecute(r ApiPOSTBundlesRequest) (*http.Response, error)
}

// BundlesApiService BundlesApi service
type BundlesApiService service

type ApiDELETEBundlesBundleIdRequest struct {
	ctx        context.Context
	ApiService BundlesApi
	bundleId   string
}

func (r ApiDELETEBundlesBundleIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBundlesBundleIdExecute(r)
}

/*
DELETEBundlesBundleId Delete a bundle

Delete a bundle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bundleId The resource's id
 @return ApiDELETEBundlesBundleIdRequest
*/
func (a *BundlesApiService) DELETEBundlesBundleId(ctx context.Context, bundleId string) ApiDELETEBundlesBundleIdRequest {
	return ApiDELETEBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
func (a *BundlesApiService) DELETEBundlesBundleIdExecute(r ApiDELETEBundlesBundleIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.DELETEBundlesBundleId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterToString(r.bundleId, "")), -1)

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

type ApiGETBundlesRequest struct {
	ctx        context.Context
	ApiService BundlesApi
}

func (r ApiGETBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBundlesExecute(r)
}

/*
GETBundles List all bundles

List all bundles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETBundlesRequest
*/
func (a *BundlesApiService) GETBundles(ctx context.Context) ApiGETBundlesRequest {
	return ApiGETBundlesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BundlesApiService) GETBundlesExecute(r ApiGETBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles"

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

type ApiGETBundlesBundleIdRequest struct {
	ctx        context.Context
	ApiService BundlesApi
	bundleId   string
}

func (r ApiGETBundlesBundleIdRequest) Execute() (*Bundle, *http.Response, error) {
	return r.ApiService.GETBundlesBundleIdExecute(r)
}

/*
GETBundlesBundleId Retrieve a bundle

Retrieve a bundle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bundleId The resource's id
 @return ApiGETBundlesBundleIdRequest
*/
func (a *BundlesApiService) GETBundlesBundleId(ctx context.Context, bundleId string) ApiGETBundlesBundleIdRequest {
	return ApiGETBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
//  @return Bundle
func (a *BundlesApiService) GETBundlesBundleIdExecute(r ApiGETBundlesBundleIdRequest) (*Bundle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Bundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETBundlesBundleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterToString(r.bundleId, "")), -1)

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

type ApiGETOrderIdAvailableFreeBundlesRequest struct {
	ctx        context.Context
	ApiService BundlesApi
	orderId    string
}

func (r ApiGETOrderIdAvailableFreeBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdAvailableFreeBundlesExecute(r)
}

/*
GETOrderIdAvailableFreeBundles Retrieve the available free bundles associated to the order

Retrieve the available free bundles associated to the order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderId The resource's id
 @return ApiGETOrderIdAvailableFreeBundlesRequest
*/
func (a *BundlesApiService) GETOrderIdAvailableFreeBundles(ctx context.Context, orderId string) ApiGETOrderIdAvailableFreeBundlesRequest {
	return ApiGETOrderIdAvailableFreeBundlesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *BundlesApiService) GETOrderIdAvailableFreeBundlesExecute(r ApiGETOrderIdAvailableFreeBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETOrderIdAvailableFreeBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/available_free_bundles"
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

type ApiGETSkuListIdBundlesRequest struct {
	ctx        context.Context
	ApiService BundlesApi
	skuListId  string
}

func (r ApiGETSkuListIdBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuListIdBundlesExecute(r)
}

/*
GETSkuListIdBundles Retrieve the bundles associated to the SKU list

Retrieve the bundles associated to the SKU list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skuListId The resource's id
 @return ApiGETSkuListIdBundlesRequest
*/
func (a *BundlesApiService) GETSkuListIdBundles(ctx context.Context, skuListId string) ApiGETSkuListIdBundlesRequest {
	return ApiGETSkuListIdBundlesRequest{
		ApiService: a,
		ctx:        ctx,
		skuListId:  skuListId,
	}
}

// Execute executes the request
func (a *BundlesApiService) GETSkuListIdBundlesExecute(r ApiGETSkuListIdBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETSkuListIdBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_lists/{skuListId}/bundles"
	localVarPath = strings.Replace(localVarPath, "{"+"skuListId"+"}", url.PathEscape(parameterToString(r.skuListId, "")), -1)

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

type ApiPATCHBundlesBundleIdRequest struct {
	ctx          context.Context
	ApiService   BundlesApi
	bundleId     string
	bundleUpdate *BundleUpdate
}

func (r ApiPATCHBundlesBundleIdRequest) BundleUpdate(bundleUpdate BundleUpdate) ApiPATCHBundlesBundleIdRequest {
	r.bundleUpdate = &bundleUpdate
	return r
}

func (r ApiPATCHBundlesBundleIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHBundlesBundleIdExecute(r)
}

/*
PATCHBundlesBundleId Update a bundle

Update a bundle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bundleId The resource's id
 @return ApiPATCHBundlesBundleIdRequest
*/
func (a *BundlesApiService) PATCHBundlesBundleId(ctx context.Context, bundleId string) ApiPATCHBundlesBundleIdRequest {
	return ApiPATCHBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
func (a *BundlesApiService) PATCHBundlesBundleIdExecute(r ApiPATCHBundlesBundleIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.PATCHBundlesBundleId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterToString(r.bundleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleUpdate == nil {
		return nil, reportError("bundleUpdate is required and must be specified")
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
	localVarPostBody = r.bundleUpdate
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

type ApiPOSTBundlesRequest struct {
	ctx          context.Context
	ApiService   BundlesApi
	bundleCreate *BundleCreate
}

func (r ApiPOSTBundlesRequest) BundleCreate(bundleCreate BundleCreate) ApiPOSTBundlesRequest {
	r.bundleCreate = &bundleCreate
	return r
}

func (r ApiPOSTBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTBundlesExecute(r)
}

/*
POSTBundles Create a bundle

Create a bundle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTBundlesRequest
*/
func (a *BundlesApiService) POSTBundles(ctx context.Context) ApiPOSTBundlesRequest {
	return ApiPOSTBundlesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BundlesApiService) POSTBundlesExecute(r ApiPOSTBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.POSTBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleCreate == nil {
		return nil, reportError("bundleCreate is required and must be specified")
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
	localVarPostBody = r.bundleCreate
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

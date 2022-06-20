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

type CheckoutComGatewaysApi interface {

	/*
		DELETECheckoutComGatewaysCheckoutComGatewayId Delete a checkout.com gateway

		Delete a checkout.com gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComGatewayId The resource's id
		@return ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest
	*/
	DELETECheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest

	// DELETECheckoutComGatewaysCheckoutComGatewayIdExecute executes the request
	DELETECheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest) (*http.Response, error)

	/*
		GETCheckoutComGateways List all checkout.com gateways

		List all checkout.com gateways

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCheckoutComGatewaysRequest
	*/
	GETCheckoutComGateways(ctx context.Context) ApiGETCheckoutComGatewaysRequest

	// GETCheckoutComGatewaysExecute executes the request
	GETCheckoutComGatewaysExecute(r ApiGETCheckoutComGatewaysRequest) (*http.Response, error)

	/*
		GETCheckoutComGatewaysCheckoutComGatewayId Retrieve a checkout.com gateway

		Retrieve a checkout.com gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComGatewayId The resource's id
		@return ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest
	*/
	GETCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest

	// GETCheckoutComGatewaysCheckoutComGatewayIdExecute executes the request
	//  @return CheckoutComGateway
	GETCheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest) (*CheckoutComGateway, *http.Response, error)

	/*
		PATCHCheckoutComGatewaysCheckoutComGatewayId Update a checkout.com gateway

		Update a checkout.com gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param checkoutComGatewayId The resource's id
		@return ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest
	*/
	PATCHCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest

	// PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute executes the request
	PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) (*http.Response, error)

	/*
		POSTCheckoutComGateways Create a checkout.com gateway

		Create a checkout.com gateway

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTCheckoutComGatewaysRequest
	*/
	POSTCheckoutComGateways(ctx context.Context) ApiPOSTCheckoutComGatewaysRequest

	// POSTCheckoutComGatewaysExecute executes the request
	POSTCheckoutComGatewaysExecute(r ApiPOSTCheckoutComGatewaysRequest) (*http.Response, error)
}

// CheckoutComGatewaysApiService CheckoutComGatewaysApi service
type CheckoutComGatewaysApiService service

type ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                  context.Context
	ApiService           CheckoutComGatewaysApi
	checkoutComGatewayId string
}

func (r ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
DELETECheckoutComGatewaysCheckoutComGatewayId Delete a checkout.com gateway

Delete a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) DELETECheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest {
	return ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
func (a *CheckoutComGatewaysApiService) DELETECheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.DELETECheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

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

type ApiGETCheckoutComGatewaysRequest struct {
	ctx        context.Context
	ApiService CheckoutComGatewaysApi
}

func (r ApiGETCheckoutComGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCheckoutComGatewaysExecute(r)
}

/*
GETCheckoutComGateways List all checkout.com gateways

List all checkout.com gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCheckoutComGatewaysRequest
*/
func (a *CheckoutComGatewaysApiService) GETCheckoutComGateways(ctx context.Context) ApiGETCheckoutComGatewaysRequest {
	return ApiGETCheckoutComGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysExecute(r ApiGETCheckoutComGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.GETCheckoutComGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways"

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

type ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                  context.Context
	ApiService           CheckoutComGatewaysApi
	checkoutComGatewayId string
}

func (r ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*CheckoutComGateway, *http.Response, error) {
	return r.ApiService.GETCheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
GETCheckoutComGatewaysCheckoutComGatewayId Retrieve a checkout.com gateway

Retrieve a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest {
	return ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
//  @return CheckoutComGateway
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest) (*CheckoutComGateway, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CheckoutComGateway
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.GETCheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

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

type ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                      context.Context
	ApiService               CheckoutComGatewaysApi
	checkoutComGatewayId     string
	checkoutComGatewayUpdate *CheckoutComGatewayUpdate
}

func (r ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) CheckoutComGatewayUpdate(checkoutComGatewayUpdate CheckoutComGatewayUpdate) ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest {
	r.checkoutComGatewayUpdate = &checkoutComGatewayUpdate
	return r
}

func (r ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
PATCHCheckoutComGatewaysCheckoutComGatewayId Update a checkout.com gateway

Update a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) PATCHCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest {
	return ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
func (a *CheckoutComGatewaysApiService) PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute(r ApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.PATCHCheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComGatewayUpdate == nil {
		return nil, reportError("checkoutComGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.checkoutComGatewayUpdate
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

type ApiPOSTCheckoutComGatewaysRequest struct {
	ctx                      context.Context
	ApiService               CheckoutComGatewaysApi
	checkoutComGatewayCreate *CheckoutComGatewayCreate
}

func (r ApiPOSTCheckoutComGatewaysRequest) CheckoutComGatewayCreate(checkoutComGatewayCreate CheckoutComGatewayCreate) ApiPOSTCheckoutComGatewaysRequest {
	r.checkoutComGatewayCreate = &checkoutComGatewayCreate
	return r
}

func (r ApiPOSTCheckoutComGatewaysRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCheckoutComGatewaysExecute(r)
}

/*
POSTCheckoutComGateways Create a checkout.com gateway

Create a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCheckoutComGatewaysRequest
*/
func (a *CheckoutComGatewaysApiService) POSTCheckoutComGateways(ctx context.Context) ApiPOSTCheckoutComGatewaysRequest {
	return ApiPOSTCheckoutComGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CheckoutComGatewaysApiService) POSTCheckoutComGatewaysExecute(r ApiPOSTCheckoutComGatewaysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.POSTCheckoutComGateways")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComGatewayCreate == nil {
		return nil, reportError("checkoutComGatewayCreate is required and must be specified")
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
	localVarPostBody = r.checkoutComGatewayCreate
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
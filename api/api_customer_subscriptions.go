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

type CustomerSubscriptionsApi interface {

	/*
		DELETECustomerSubscriptionsCustomerSubscriptionId Delete a customer subscription

		Delete a customer subscription

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerSubscriptionId The resource's id
		@return ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest
	*/
	DELETECustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest

	// DELETECustomerSubscriptionsCustomerSubscriptionIdExecute executes the request
	DELETECustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest) (*http.Response, error)

	/*
		GETCustomerIdCustomerSubscriptions Retrieve the customer subscriptions associated to the customer

		Retrieve the customer subscriptions associated to the customer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerId The resource's id
		@return ApiGETCustomerIdCustomerSubscriptionsRequest
	*/
	GETCustomerIdCustomerSubscriptions(ctx context.Context, customerId string) ApiGETCustomerIdCustomerSubscriptionsRequest

	// GETCustomerIdCustomerSubscriptionsExecute executes the request
	GETCustomerIdCustomerSubscriptionsExecute(r ApiGETCustomerIdCustomerSubscriptionsRequest) (*http.Response, error)

	/*
		GETCustomerSubscriptions List all customer subscriptions

		List all customer subscriptions

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETCustomerSubscriptionsRequest
	*/
	GETCustomerSubscriptions(ctx context.Context) ApiGETCustomerSubscriptionsRequest

	// GETCustomerSubscriptionsExecute executes the request
	GETCustomerSubscriptionsExecute(r ApiGETCustomerSubscriptionsRequest) (*http.Response, error)

	/*
		GETCustomerSubscriptionsCustomerSubscriptionId Retrieve a customer subscription

		Retrieve a customer subscription

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerSubscriptionId The resource's id
		@return ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest
	*/
	GETCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest

	// GETCustomerSubscriptionsCustomerSubscriptionIdExecute executes the request
	//  @return CustomerSubscription
	GETCustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest) (*CustomerSubscription, *http.Response, error)

	/*
		PATCHCustomerSubscriptionsCustomerSubscriptionId Update a customer subscription

		Update a customer subscription

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerSubscriptionId The resource's id
		@return ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest
	*/
	PATCHCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest

	// PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute executes the request
	PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) (*http.Response, error)

	/*
		POSTCustomerSubscriptions Create a customer subscription

		Create a customer subscription

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTCustomerSubscriptionsRequest
	*/
	POSTCustomerSubscriptions(ctx context.Context) ApiPOSTCustomerSubscriptionsRequest

	// POSTCustomerSubscriptionsExecute executes the request
	POSTCustomerSubscriptionsExecute(r ApiPOSTCustomerSubscriptionsRequest) (*http.Response, error)
}

// CustomerSubscriptionsApiService CustomerSubscriptionsApi service
type CustomerSubscriptionsApiService service

type ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                    context.Context
	ApiService             CustomerSubscriptionsApi
	customerSubscriptionId string
}

func (r ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
DELETECustomerSubscriptionsCustomerSubscriptionId Delete a customer subscription

Delete a customer subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerSubscriptionId The resource's id
 @return ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) DELETECustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest {
	return ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) DELETECustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.DELETECustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

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

type ApiGETCustomerIdCustomerSubscriptionsRequest struct {
	ctx        context.Context
	ApiService CustomerSubscriptionsApi
	customerId string
}

func (r ApiGETCustomerIdCustomerSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerSubscriptionsExecute(r)
}

/*
GETCustomerIdCustomerSubscriptions Retrieve the customer subscriptions associated to the customer

Retrieve the customer subscriptions associated to the customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The resource's id
 @return ApiGETCustomerIdCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerIdCustomerSubscriptions(ctx context.Context, customerId string) ApiGETCustomerIdCustomerSubscriptionsRequest {
	return ApiGETCustomerIdCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) GETCustomerIdCustomerSubscriptionsExecute(r ApiGETCustomerIdCustomerSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerIdCustomerSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

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

type ApiGETCustomerSubscriptionsRequest struct {
	ctx        context.Context
	ApiService CustomerSubscriptionsApi
}

func (r ApiGETCustomerSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerSubscriptionsExecute(r)
}

/*
GETCustomerSubscriptions List all customer subscriptions

List all customer subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptions(ctx context.Context) ApiGETCustomerSubscriptionsRequest {
	return ApiGETCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsExecute(r ApiGETCustomerSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions"

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

type ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                    context.Context
	ApiService             CustomerSubscriptionsApi
	customerSubscriptionId string
}

func (r ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*CustomerSubscription, *http.Response, error) {
	return r.ApiService.GETCustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
GETCustomerSubscriptionsCustomerSubscriptionId Retrieve a customer subscription

Retrieve a customer subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerSubscriptionId The resource's id
 @return ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest {
	return ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
//  @return CustomerSubscription
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest) (*CustomerSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

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

type ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                        context.Context
	ApiService                 CustomerSubscriptionsApi
	customerSubscriptionId     string
	customerSubscriptionUpdate *CustomerSubscriptionUpdate
}

func (r ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) CustomerSubscriptionUpdate(customerSubscriptionUpdate CustomerSubscriptionUpdate) ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest {
	r.customerSubscriptionUpdate = &customerSubscriptionUpdate
	return r
}

func (r ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
PATCHCustomerSubscriptionsCustomerSubscriptionId Update a customer subscription

Update a customer subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerSubscriptionId The resource's id
 @return ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) PATCHCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest {
	return ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute(r ApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.PATCHCustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerSubscriptionUpdate == nil {
		return nil, reportError("customerSubscriptionUpdate is required and must be specified")
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
	localVarPostBody = r.customerSubscriptionUpdate
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

type ApiPOSTCustomerSubscriptionsRequest struct {
	ctx                        context.Context
	ApiService                 CustomerSubscriptionsApi
	customerSubscriptionCreate *CustomerSubscriptionCreate
}

func (r ApiPOSTCustomerSubscriptionsRequest) CustomerSubscriptionCreate(customerSubscriptionCreate CustomerSubscriptionCreate) ApiPOSTCustomerSubscriptionsRequest {
	r.customerSubscriptionCreate = &customerSubscriptionCreate
	return r
}

func (r ApiPOSTCustomerSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCustomerSubscriptionsExecute(r)
}

/*
POSTCustomerSubscriptions Create a customer subscription

Create a customer subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) POSTCustomerSubscriptions(ctx context.Context) ApiPOSTCustomerSubscriptionsRequest {
	return ApiPOSTCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) POSTCustomerSubscriptionsExecute(r ApiPOSTCustomerSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.POSTCustomerSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerSubscriptionCreate == nil {
		return nil, reportError("customerSubscriptionCreate is required and must be specified")
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
	localVarPostBody = r.customerSubscriptionCreate
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

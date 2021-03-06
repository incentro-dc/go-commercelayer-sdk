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

// OrderValidationRulesApiService OrderValidationRulesApi service
type OrderValidationRulesApiService service

type ApiGETOrderValidationRulesRequest struct {
	ctx        context.Context
	ApiService *OrderValidationRulesApiService
}

func (r ApiGETOrderValidationRulesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderValidationRulesExecute(r)
}

/*
GETOrderValidationRules List all order validation rules

List all order validation rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETOrderValidationRulesRequest
*/
func (a *OrderValidationRulesApiService) GETOrderValidationRules(ctx context.Context) ApiGETOrderValidationRulesRequest {
	return ApiGETOrderValidationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OrderValidationRulesApiService) GETOrderValidationRulesExecute(r ApiGETOrderValidationRulesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderValidationRulesApiService.GETOrderValidationRules")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_validation_rules"

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

type ApiGETOrderValidationRulesOrderValidationRuleIdRequest struct {
	ctx                   context.Context
	ApiService            *OrderValidationRulesApiService
	orderValidationRuleId string
}

func (r ApiGETOrderValidationRulesOrderValidationRuleIdRequest) Execute() (*OrderValidationRule, *http.Response, error) {
	return r.ApiService.GETOrderValidationRulesOrderValidationRuleIdExecute(r)
}

/*
GETOrderValidationRulesOrderValidationRuleId Retrieve an order validation rule

Retrieve an order validation rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orderValidationRuleId The resource's id
 @return ApiGETOrderValidationRulesOrderValidationRuleIdRequest
*/
func (a *OrderValidationRulesApiService) GETOrderValidationRulesOrderValidationRuleId(ctx context.Context, orderValidationRuleId string) ApiGETOrderValidationRulesOrderValidationRuleIdRequest {
	return ApiGETOrderValidationRulesOrderValidationRuleIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		orderValidationRuleId: orderValidationRuleId,
	}
}

// Execute executes the request
//  @return OrderValidationRule
func (a *OrderValidationRulesApiService) GETOrderValidationRulesOrderValidationRuleIdExecute(r ApiGETOrderValidationRulesOrderValidationRuleIdRequest) (*OrderValidationRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderValidationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderValidationRulesApiService.GETOrderValidationRulesOrderValidationRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_validation_rules/{orderValidationRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderValidationRuleId"+"}", url.PathEscape(parameterToString(r.orderValidationRuleId, "")), -1)

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

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

// BillingInfoValidationRulesApiService BillingInfoValidationRulesApi service
type BillingInfoValidationRulesApiService service

type BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct {
	ctx                         context.Context
	ApiService                  *BillingInfoValidationRulesApiService
	billingInfoValidationRuleId string
}

func (r BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r)
}

/*
DELETEBillingInfoValidationRulesBillingInfoValidationRuleId Delete a billing info validation rule

Delete a billing info validation rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param billingInfoValidationRuleId The resource's id
 @return BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest
*/
func (a *BillingInfoValidationRulesApiService) DELETEBillingInfoValidationRulesBillingInfoValidationRuleId(ctx context.Context, billingInfoValidationRuleId string) BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest {
	return BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest{
		ApiService:                  a,
		ctx:                         ctx,
		billingInfoValidationRuleId: billingInfoValidationRuleId,
	}
}

// Execute executes the request
func (a *BillingInfoValidationRulesApiService) DELETEBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r BillingInfoValidationRulesApiDELETEBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoValidationRulesApiService.DELETEBillingInfoValidationRulesBillingInfoValidationRuleId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing_info_validation_rules/{billingInfoValidationRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"billingInfoValidationRuleId"+"}", url.PathEscape(parameterToString(r.billingInfoValidationRuleId, "")), -1)

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

type BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest struct {
	ctx        context.Context
	ApiService *BillingInfoValidationRulesApiService
}

func (r BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest) Execute() (*GETBillingInfoValidationRules200Response, *http.Response, error) {
	return r.ApiService.GETBillingInfoValidationRulesExecute(r)
}

/*
GETBillingInfoValidationRules List all billing info validation rules

List all billing info validation rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest
*/
func (a *BillingInfoValidationRulesApiService) GETBillingInfoValidationRules(ctx context.Context) BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest {
	return BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETBillingInfoValidationRules200Response
func (a *BillingInfoValidationRulesApiService) GETBillingInfoValidationRulesExecute(r BillingInfoValidationRulesApiGETBillingInfoValidationRulesRequest) (*GETBillingInfoValidationRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBillingInfoValidationRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoValidationRulesApiService.GETBillingInfoValidationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing_info_validation_rules"

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

type BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct {
	ctx                         context.Context
	ApiService                  *BillingInfoValidationRulesApiService
	billingInfoValidationRuleId string
}

func (r BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) Execute() (*GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response, *http.Response, error) {
	return r.ApiService.GETBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r)
}

/*
GETBillingInfoValidationRulesBillingInfoValidationRuleId Retrieve a billing info validation rule

Retrieve a billing info validation rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param billingInfoValidationRuleId The resource's id
 @return BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest
*/
func (a *BillingInfoValidationRulesApiService) GETBillingInfoValidationRulesBillingInfoValidationRuleId(ctx context.Context, billingInfoValidationRuleId string) BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest {
	return BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest{
		ApiService:                  a,
		ctx:                         ctx,
		billingInfoValidationRuleId: billingInfoValidationRuleId,
	}
}

// Execute executes the request
//  @return GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response
func (a *BillingInfoValidationRulesApiService) GETBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r BillingInfoValidationRulesApiGETBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) (*GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoValidationRulesApiService.GETBillingInfoValidationRulesBillingInfoValidationRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing_info_validation_rules/{billingInfoValidationRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"billingInfoValidationRuleId"+"}", url.PathEscape(parameterToString(r.billingInfoValidationRuleId, "")), -1)

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

type BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest struct {
	ctx                             context.Context
	ApiService                      *BillingInfoValidationRulesApiService
	billingInfoValidationRuleUpdate *BillingInfoValidationRuleUpdate
	billingInfoValidationRuleId     string
}

func (r BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) BillingInfoValidationRuleUpdate(billingInfoValidationRuleUpdate BillingInfoValidationRuleUpdate) BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest {
	r.billingInfoValidationRuleUpdate = &billingInfoValidationRuleUpdate
	return r
}

func (r BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) Execute() (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response, *http.Response, error) {
	return r.ApiService.PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r)
}

/*
PATCHBillingInfoValidationRulesBillingInfoValidationRuleId Update a billing info validation rule

Update a billing info validation rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param billingInfoValidationRuleId The resource's id
 @return BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest
*/
func (a *BillingInfoValidationRulesApiService) PATCHBillingInfoValidationRulesBillingInfoValidationRuleId(ctx context.Context, billingInfoValidationRuleId string) BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest {
	return BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest{
		ApiService:                  a,
		ctx:                         ctx,
		billingInfoValidationRuleId: billingInfoValidationRuleId,
	}
}

// Execute executes the request
//  @return PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response
func (a *BillingInfoValidationRulesApiService) PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdExecute(r BillingInfoValidationRulesApiPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequest) (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoValidationRulesApiService.PATCHBillingInfoValidationRulesBillingInfoValidationRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing_info_validation_rules/{billingInfoValidationRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"billingInfoValidationRuleId"+"}", url.PathEscape(parameterToString(r.billingInfoValidationRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.billingInfoValidationRuleUpdate == nil {
		return localVarReturnValue, nil, reportError("billingInfoValidationRuleUpdate is required and must be specified")
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
	localVarPostBody = r.billingInfoValidationRuleUpdate
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

type BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest struct {
	ctx                             context.Context
	ApiService                      *BillingInfoValidationRulesApiService
	billingInfoValidationRuleCreate *BillingInfoValidationRuleCreate
}

func (r BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest) BillingInfoValidationRuleCreate(billingInfoValidationRuleCreate BillingInfoValidationRuleCreate) BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest {
	r.billingInfoValidationRuleCreate = &billingInfoValidationRuleCreate
	return r
}

func (r BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest) Execute() (*POSTBillingInfoValidationRules201Response, *http.Response, error) {
	return r.ApiService.POSTBillingInfoValidationRulesExecute(r)
}

/*
POSTBillingInfoValidationRules Create a billing info validation rule

Create a billing info validation rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest
*/
func (a *BillingInfoValidationRulesApiService) POSTBillingInfoValidationRules(ctx context.Context) BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest {
	return BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTBillingInfoValidationRules201Response
func (a *BillingInfoValidationRulesApiService) POSTBillingInfoValidationRulesExecute(r BillingInfoValidationRulesApiPOSTBillingInfoValidationRulesRequest) (*POSTBillingInfoValidationRules201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTBillingInfoValidationRules201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingInfoValidationRulesApiService.POSTBillingInfoValidationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billing_info_validation_rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.billingInfoValidationRuleCreate == nil {
		return localVarReturnValue, nil, reportError("billingInfoValidationRuleCreate is required and must be specified")
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
	localVarPostBody = r.billingInfoValidationRuleCreate
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

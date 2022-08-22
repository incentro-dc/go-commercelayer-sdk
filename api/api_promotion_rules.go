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

// PromotionRulesApiService PromotionRulesApi service
type PromotionRulesApiService service

type PromotionRulesApiGETPromotionRulesRequest struct {
	ctx        context.Context
	ApiService *PromotionRulesApiService
}

func (r PromotionRulesApiGETPromotionRulesRequest) Execute() (*GETPromotionRules200Response, *http.Response, error) {
	return r.ApiService.GETPromotionRulesExecute(r)
}

/*
GETPromotionRules List all promotion rules

List all promotion rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PromotionRulesApiGETPromotionRulesRequest
*/
func (a *PromotionRulesApiService) GETPromotionRules(ctx context.Context) PromotionRulesApiGETPromotionRulesRequest {
	return PromotionRulesApiGETPromotionRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETPromotionRules200Response
func (a *PromotionRulesApiService) GETPromotionRulesExecute(r PromotionRulesApiGETPromotionRulesRequest) (*GETPromotionRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPromotionRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PromotionRulesApiService.GETPromotionRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/promotion_rules"

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

type PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest struct {
	ctx             context.Context
	ApiService      *PromotionRulesApiService
	promotionRuleId string
}

func (r PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest) Execute() (*GETPromotionRulesPromotionRuleId200Response, *http.Response, error) {
	return r.ApiService.GETPromotionRulesPromotionRuleIdExecute(r)
}

/*
GETPromotionRulesPromotionRuleId Retrieve a promotion rule

Retrieve a promotion rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param promotionRuleId The resource's id
 @return PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest
*/
func (a *PromotionRulesApiService) GETPromotionRulesPromotionRuleId(ctx context.Context, promotionRuleId string) PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest {
	return PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest{
		ApiService:      a,
		ctx:             ctx,
		promotionRuleId: promotionRuleId,
	}
}

// Execute executes the request
//  @return GETPromotionRulesPromotionRuleId200Response
func (a *PromotionRulesApiService) GETPromotionRulesPromotionRuleIdExecute(r PromotionRulesApiGETPromotionRulesPromotionRuleIdRequest) (*GETPromotionRulesPromotionRuleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPromotionRulesPromotionRuleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PromotionRulesApiService.GETPromotionRulesPromotionRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/promotion_rules/{promotionRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"promotionRuleId"+"}", url.PathEscape(parameterToString(r.promotionRuleId, "")), -1)

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

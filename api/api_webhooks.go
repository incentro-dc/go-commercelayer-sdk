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

// WebhooksApiService WebhooksApi service
type WebhooksApiService service

type WebhooksApiDELETEWebhooksWebhookIdRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	webhookId  string
}

func (r WebhooksApiDELETEWebhooksWebhookIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEWebhooksWebhookIdExecute(r)
}

/*
DELETEWebhooksWebhookId Delete a webhook

Delete a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return WebhooksApiDELETEWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) DELETEWebhooksWebhookId(ctx context.Context, webhookId string) WebhooksApiDELETEWebhooksWebhookIdRequest {
	return WebhooksApiDELETEWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) DELETEWebhooksWebhookIdExecute(r WebhooksApiDELETEWebhooksWebhookIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.DELETEWebhooksWebhookId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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

type WebhooksApiGETEventCallbackIdWebhookRequest struct {
	ctx             context.Context
	ApiService      *WebhooksApiService
	eventCallbackId string
}

func (r WebhooksApiGETEventCallbackIdWebhookRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETEventCallbackIdWebhookExecute(r)
}

/*
GETEventCallbackIdWebhook Retrieve the webhook associated to the event callback

Retrieve the webhook associated to the event callback

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventCallbackId The resource's id
 @return WebhooksApiGETEventCallbackIdWebhookRequest
*/
func (a *WebhooksApiService) GETEventCallbackIdWebhook(ctx context.Context, eventCallbackId string) WebhooksApiGETEventCallbackIdWebhookRequest {
	return WebhooksApiGETEventCallbackIdWebhookRequest{
		ApiService:      a,
		ctx:             ctx,
		eventCallbackId: eventCallbackId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) GETEventCallbackIdWebhookExecute(r WebhooksApiGETEventCallbackIdWebhookRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETEventCallbackIdWebhook")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event_callbacks/{eventCallbackId}/webhook"
	localVarPath = strings.Replace(localVarPath, "{"+"eventCallbackId"+"}", url.PathEscape(parameterToString(r.eventCallbackId, "")), -1)

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

type WebhooksApiGETEventIdWebhooksRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	eventId    string
}

func (r WebhooksApiGETEventIdWebhooksRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETEventIdWebhooksExecute(r)
}

/*
GETEventIdWebhooks Retrieve the webhooks associated to the event

Retrieve the webhooks associated to the event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId The resource's id
 @return WebhooksApiGETEventIdWebhooksRequest
*/
func (a *WebhooksApiService) GETEventIdWebhooks(ctx context.Context, eventId string) WebhooksApiGETEventIdWebhooksRequest {
	return WebhooksApiGETEventIdWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
		eventId:    eventId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) GETEventIdWebhooksExecute(r WebhooksApiGETEventIdWebhooksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETEventIdWebhooks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{eventId}/webhooks"
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterToString(r.eventId, "")), -1)

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

type WebhooksApiGETWebhooksRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
}

func (r WebhooksApiGETWebhooksRequest) Execute() (*GETWebhooks200Response, *http.Response, error) {
	return r.ApiService.GETWebhooksExecute(r)
}

/*
GETWebhooks List all webhooks

List all webhooks

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WebhooksApiGETWebhooksRequest
*/
func (a *WebhooksApiService) GETWebhooks(ctx context.Context) WebhooksApiGETWebhooksRequest {
	return WebhooksApiGETWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETWebhooks200Response
func (a *WebhooksApiService) GETWebhooksExecute(r WebhooksApiGETWebhooksRequest) (*GETWebhooks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETWebhooks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

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

type WebhooksApiGETWebhooksWebhookIdRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	webhookId  string
}

func (r WebhooksApiGETWebhooksWebhookIdRequest) Execute() (*GETWebhooksWebhookId200Response, *http.Response, error) {
	return r.ApiService.GETWebhooksWebhookIdExecute(r)
}

/*
GETWebhooksWebhookId Retrieve a webhook

Retrieve a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return WebhooksApiGETWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) GETWebhooksWebhookId(ctx context.Context, webhookId string) WebhooksApiGETWebhooksWebhookIdRequest {
	return WebhooksApiGETWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//  @return GETWebhooksWebhookId200Response
func (a *WebhooksApiService) GETWebhooksWebhookIdExecute(r WebhooksApiGETWebhooksWebhookIdRequest) (*GETWebhooksWebhookId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETWebhooksWebhookId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETWebhooksWebhookId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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

type WebhooksApiPATCHWebhooksWebhookIdRequest struct {
	ctx           context.Context
	ApiService    *WebhooksApiService
	webhookUpdate *WebhookUpdate
	webhookId     string
}

func (r WebhooksApiPATCHWebhooksWebhookIdRequest) WebhookUpdate(webhookUpdate WebhookUpdate) WebhooksApiPATCHWebhooksWebhookIdRequest {
	r.webhookUpdate = &webhookUpdate
	return r
}

func (r WebhooksApiPATCHWebhooksWebhookIdRequest) Execute() (*PATCHWebhooksWebhookId200Response, *http.Response, error) {
	return r.ApiService.PATCHWebhooksWebhookIdExecute(r)
}

/*
PATCHWebhooksWebhookId Update a webhook

Update a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return WebhooksApiPATCHWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) PATCHWebhooksWebhookId(ctx context.Context, webhookId string) WebhooksApiPATCHWebhooksWebhookIdRequest {
	return WebhooksApiPATCHWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//  @return PATCHWebhooksWebhookId200Response
func (a *WebhooksApiService) PATCHWebhooksWebhookIdExecute(r WebhooksApiPATCHWebhooksWebhookIdRequest) (*PATCHWebhooksWebhookId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHWebhooksWebhookId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.PATCHWebhooksWebhookId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookUpdate == nil {
		return localVarReturnValue, nil, reportError("webhookUpdate is required and must be specified")
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
	localVarPostBody = r.webhookUpdate
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

type WebhooksApiPOSTWebhooksRequest struct {
	ctx           context.Context
	ApiService    *WebhooksApiService
	webhookCreate *WebhookCreate
}

func (r WebhooksApiPOSTWebhooksRequest) WebhookCreate(webhookCreate WebhookCreate) WebhooksApiPOSTWebhooksRequest {
	r.webhookCreate = &webhookCreate
	return r
}

func (r WebhooksApiPOSTWebhooksRequest) Execute() (*POSTWebhooks201Response, *http.Response, error) {
	return r.ApiService.POSTWebhooksExecute(r)
}

/*
POSTWebhooks Create a webhook

Create a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WebhooksApiPOSTWebhooksRequest
*/
func (a *WebhooksApiService) POSTWebhooks(ctx context.Context) WebhooksApiPOSTWebhooksRequest {
	return WebhooksApiPOSTWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTWebhooks201Response
func (a *WebhooksApiService) POSTWebhooksExecute(r WebhooksApiPOSTWebhooksRequest) (*POSTWebhooks201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTWebhooks201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.POSTWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookCreate == nil {
		return localVarReturnValue, nil, reportError("webhookCreate is required and must be specified")
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
	localVarPostBody = r.webhookCreate
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

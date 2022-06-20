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

type TransactionsApi interface {

	/*
		GETTransactions List all transactions

		List all transactions

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETTransactionsRequest
	*/
	GETTransactions(ctx context.Context) ApiGETTransactionsRequest

	// GETTransactionsExecute executes the request
	GETTransactionsExecute(r ApiGETTransactionsRequest) (*http.Response, error)

	/*
		GETTransactionsTransactionId Retrieve a transaction

		Retrieve a transaction

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param transactionId The resource's id
		@return ApiGETTransactionsTransactionIdRequest
	*/
	GETTransactionsTransactionId(ctx context.Context, transactionId string) ApiGETTransactionsTransactionIdRequest

	// GETTransactionsTransactionIdExecute executes the request
	//  @return Transaction
	GETTransactionsTransactionIdExecute(r ApiGETTransactionsTransactionIdRequest) (*Transaction, *http.Response, error)
}

// TransactionsApiService TransactionsApi service
type TransactionsApiService service

type ApiGETTransactionsRequest struct {
	ctx        context.Context
	ApiService TransactionsApi
}

func (r ApiGETTransactionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETTransactionsExecute(r)
}

/*
GETTransactions List all transactions

List all transactions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETTransactionsRequest
*/
func (a *TransactionsApiService) GETTransactions(ctx context.Context) ApiGETTransactionsRequest {
	return ApiGETTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TransactionsApiService) GETTransactionsExecute(r ApiGETTransactionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GETTransactions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions"

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

type ApiGETTransactionsTransactionIdRequest struct {
	ctx           context.Context
	ApiService    TransactionsApi
	transactionId string
}

func (r ApiGETTransactionsTransactionIdRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.GETTransactionsTransactionIdExecute(r)
}

/*
GETTransactionsTransactionId Retrieve a transaction

Retrieve a transaction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param transactionId The resource's id
 @return ApiGETTransactionsTransactionIdRequest
*/
func (a *TransactionsApiService) GETTransactionsTransactionId(ctx context.Context, transactionId string) ApiGETTransactionsTransactionIdRequest {
	return ApiGETTransactionsTransactionIdRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//  @return Transaction
func (a *TransactionsApiService) GETTransactionsTransactionIdExecute(r ApiGETTransactionsTransactionIdRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GETTransactionsTransactionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterToString(r.transactionId, "")), -1)

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
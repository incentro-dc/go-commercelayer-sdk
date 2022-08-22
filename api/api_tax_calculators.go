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

// TaxCalculatorsApiService TaxCalculatorsApi service
type TaxCalculatorsApiService service

type TaxCalculatorsApiGETMarketIdTaxCalculatorRequest struct {
	ctx        context.Context
	ApiService *TaxCalculatorsApiService
	marketId   string
}

func (r TaxCalculatorsApiGETMarketIdTaxCalculatorRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdTaxCalculatorExecute(r)
}

/*
GETMarketIdTaxCalculator Retrieve the tax calculator associated to the market

Retrieve the tax calculator associated to the market

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketId The resource's id
 @return TaxCalculatorsApiGETMarketIdTaxCalculatorRequest
*/
func (a *TaxCalculatorsApiService) GETMarketIdTaxCalculator(ctx context.Context, marketId string) TaxCalculatorsApiGETMarketIdTaxCalculatorRequest {
	return TaxCalculatorsApiGETMarketIdTaxCalculatorRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *TaxCalculatorsApiService) GETMarketIdTaxCalculatorExecute(r TaxCalculatorsApiGETMarketIdTaxCalculatorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCalculatorsApiService.GETMarketIdTaxCalculator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/tax_calculator"
	localVarPath = strings.Replace(localVarPath, "{"+"marketId"+"}", url.PathEscape(parameterToString(r.marketId, "")), -1)

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

type TaxCalculatorsApiGETTaxCalculatorsRequest struct {
	ctx        context.Context
	ApiService *TaxCalculatorsApiService
}

func (r TaxCalculatorsApiGETTaxCalculatorsRequest) Execute() (*GETTaxCalculators200Response, *http.Response, error) {
	return r.ApiService.GETTaxCalculatorsExecute(r)
}

/*
GETTaxCalculators List all tax calculators

List all tax calculators

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TaxCalculatorsApiGETTaxCalculatorsRequest
*/
func (a *TaxCalculatorsApiService) GETTaxCalculators(ctx context.Context) TaxCalculatorsApiGETTaxCalculatorsRequest {
	return TaxCalculatorsApiGETTaxCalculatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETTaxCalculators200Response
func (a *TaxCalculatorsApiService) GETTaxCalculatorsExecute(r TaxCalculatorsApiGETTaxCalculatorsRequest) (*GETTaxCalculators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxCalculators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCalculatorsApiService.GETTaxCalculators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_calculators"

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

type TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest struct {
	ctx             context.Context
	ApiService      *TaxCalculatorsApiService
	taxCalculatorId string
}

func (r TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest) Execute() (*GETTaxCalculatorsTaxCalculatorId200Response, *http.Response, error) {
	return r.ApiService.GETTaxCalculatorsTaxCalculatorIdExecute(r)
}

/*
GETTaxCalculatorsTaxCalculatorId Retrieve a tax calculator

Retrieve a tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxCalculatorId The resource's id
 @return TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest
*/
func (a *TaxCalculatorsApiService) GETTaxCalculatorsTaxCalculatorId(ctx context.Context, taxCalculatorId string) TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest {
	return TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest{
		ApiService:      a,
		ctx:             ctx,
		taxCalculatorId: taxCalculatorId,
	}
}

// Execute executes the request
//  @return GETTaxCalculatorsTaxCalculatorId200Response
func (a *TaxCalculatorsApiService) GETTaxCalculatorsTaxCalculatorIdExecute(r TaxCalculatorsApiGETTaxCalculatorsTaxCalculatorIdRequest) (*GETTaxCalculatorsTaxCalculatorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxCalculatorsTaxCalculatorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCalculatorsApiService.GETTaxCalculatorsTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_calculators/{taxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxCalculatorId"+"}", url.PathEscape(parameterToString(r.taxCalculatorId, "")), -1)

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

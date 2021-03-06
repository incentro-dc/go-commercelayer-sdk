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

// CouponsApiService CouponsApi service
type CouponsApiService service

type ApiDELETECouponsCouponIdRequest struct {
	ctx        context.Context
	ApiService *CouponsApiService
	couponId   string
}

func (r ApiDELETECouponsCouponIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECouponsCouponIdExecute(r)
}

/*
DELETECouponsCouponId Delete a coupon

Delete a coupon

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponId The resource's id
 @return ApiDELETECouponsCouponIdRequest
*/
func (a *CouponsApiService) DELETECouponsCouponId(ctx context.Context, couponId string) ApiDELETECouponsCouponIdRequest {
	return ApiDELETECouponsCouponIdRequest{
		ApiService: a,
		ctx:        ctx,
		couponId:   couponId,
	}
}

// Execute executes the request
func (a *CouponsApiService) DELETECouponsCouponIdExecute(r ApiDELETECouponsCouponIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.DELETECouponsCouponId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupons/{couponId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponId"+"}", url.PathEscape(parameterToString(r.couponId, "")), -1)

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

type ApiGETCouponCodesPromotionRuleIdCouponsRequest struct {
	ctx                        context.Context
	ApiService                 *CouponsApiService
	couponCodesPromotionRuleId string
}

func (r ApiGETCouponCodesPromotionRuleIdCouponsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCouponCodesPromotionRuleIdCouponsExecute(r)
}

/*
GETCouponCodesPromotionRuleIdCoupons Retrieve the coupons associated to the coupon codes promotion rule

Retrieve the coupons associated to the coupon codes promotion rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponCodesPromotionRuleId The resource's id
 @return ApiGETCouponCodesPromotionRuleIdCouponsRequest
*/
func (a *CouponsApiService) GETCouponCodesPromotionRuleIdCoupons(ctx context.Context, couponCodesPromotionRuleId string) ApiGETCouponCodesPromotionRuleIdCouponsRequest {
	return ApiGETCouponCodesPromotionRuleIdCouponsRequest{
		ApiService:                 a,
		ctx:                        ctx,
		couponCodesPromotionRuleId: couponCodesPromotionRuleId,
	}
}

// Execute executes the request
func (a *CouponsApiService) GETCouponCodesPromotionRuleIdCouponsExecute(r ApiGETCouponCodesPromotionRuleIdCouponsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.GETCouponCodesPromotionRuleIdCoupons")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_codes_promotion_rules/{couponCodesPromotionRuleId}/coupons"
	localVarPath = strings.Replace(localVarPath, "{"+"couponCodesPromotionRuleId"+"}", url.PathEscape(parameterToString(r.couponCodesPromotionRuleId, "")), -1)

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

type ApiGETCouponsRequest struct {
	ctx        context.Context
	ApiService *CouponsApiService
}

func (r ApiGETCouponsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCouponsExecute(r)
}

/*
GETCoupons List all coupons

List all coupons

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETCouponsRequest
*/
func (a *CouponsApiService) GETCoupons(ctx context.Context) ApiGETCouponsRequest {
	return ApiGETCouponsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CouponsApiService) GETCouponsExecute(r ApiGETCouponsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.GETCoupons")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupons"

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

type ApiGETCouponsCouponIdRequest struct {
	ctx        context.Context
	ApiService *CouponsApiService
	couponId   string
}

func (r ApiGETCouponsCouponIdRequest) Execute() (*Coupon, *http.Response, error) {
	return r.ApiService.GETCouponsCouponIdExecute(r)
}

/*
GETCouponsCouponId Retrieve a coupon

Retrieve a coupon

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponId The resource's id
 @return ApiGETCouponsCouponIdRequest
*/
func (a *CouponsApiService) GETCouponsCouponId(ctx context.Context, couponId string) ApiGETCouponsCouponIdRequest {
	return ApiGETCouponsCouponIdRequest{
		ApiService: a,
		ctx:        ctx,
		couponId:   couponId,
	}
}

// Execute executes the request
//  @return Coupon
func (a *CouponsApiService) GETCouponsCouponIdExecute(r ApiGETCouponsCouponIdRequest) (*Coupon, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Coupon
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.GETCouponsCouponId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupons/{couponId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponId"+"}", url.PathEscape(parameterToString(r.couponId, "")), -1)

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

type ApiPATCHCouponsCouponIdRequest struct {
	ctx          context.Context
	ApiService   *CouponsApiService
	couponUpdate *CouponUpdate
	couponId     string
}

func (r ApiPATCHCouponsCouponIdRequest) CouponUpdate(couponUpdate CouponUpdate) ApiPATCHCouponsCouponIdRequest {
	r.couponUpdate = &couponUpdate
	return r
}

func (r ApiPATCHCouponsCouponIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHCouponsCouponIdExecute(r)
}

/*
PATCHCouponsCouponId Update a coupon

Update a coupon

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponId The resource's id
 @return ApiPATCHCouponsCouponIdRequest
*/
func (a *CouponsApiService) PATCHCouponsCouponId(ctx context.Context, couponId string) ApiPATCHCouponsCouponIdRequest {
	return ApiPATCHCouponsCouponIdRequest{
		ApiService: a,
		ctx:        ctx,
		couponId:   couponId,
	}
}

// Execute executes the request
func (a *CouponsApiService) PATCHCouponsCouponIdExecute(r ApiPATCHCouponsCouponIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.PATCHCouponsCouponId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupons/{couponId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponId"+"}", url.PathEscape(parameterToString(r.couponId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.couponUpdate == nil {
		return nil, reportError("couponUpdate is required and must be specified")
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
	localVarPostBody = r.couponUpdate
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

type ApiPOSTCouponsRequest struct {
	ctx          context.Context
	ApiService   *CouponsApiService
	couponCreate *CouponCreate
}

func (r ApiPOSTCouponsRequest) CouponCreate(couponCreate CouponCreate) ApiPOSTCouponsRequest {
	r.couponCreate = &couponCreate
	return r
}

func (r ApiPOSTCouponsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTCouponsExecute(r)
}

/*
POSTCoupons Create a coupon

Create a coupon

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTCouponsRequest
*/
func (a *CouponsApiService) POSTCoupons(ctx context.Context) ApiPOSTCouponsRequest {
	return ApiPOSTCouponsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CouponsApiService) POSTCouponsExecute(r ApiPOSTCouponsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponsApiService.POSTCoupons")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupons"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.couponCreate == nil {
		return nil, reportError("couponCreate is required and must be specified")
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
	localVarPostBody = r.couponCreate
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

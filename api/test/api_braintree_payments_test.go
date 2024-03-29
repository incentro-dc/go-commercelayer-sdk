/*
Commerce Layer API

Testing BraintreePaymentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_BraintreePaymentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BraintreePaymentsApiService DELETEBraintreePaymentsBraintreePaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreePaymentId interface{}

		httpRes, err := apiClient.BraintreePaymentsApi.DELETEBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BraintreePaymentsApiService GETBraintreeGatewayIdBraintreePayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreeGatewayId interface{}

		httpRes, err := apiClient.BraintreePaymentsApi.GETBraintreeGatewayIdBraintreePayments(context.Background(), braintreeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BraintreePaymentsApiService GETBraintreePayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BraintreePaymentsApi.GETBraintreePayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BraintreePaymentsApiService GETBraintreePaymentsBraintreePaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreePaymentId interface{}

		resp, httpRes, err := apiClient.BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BraintreePaymentsApiService PATCHBraintreePaymentsBraintreePaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreePaymentId interface{}

		resp, httpRes, err := apiClient.BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BraintreePaymentsApiService POSTBraintreePayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BraintreePaymentsApi.POSTBraintreePayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

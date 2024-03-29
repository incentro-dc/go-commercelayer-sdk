/*
Commerce Layer API

Testing PaypalPaymentsApiService

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

func Test_api_PaypalPaymentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaypalPaymentsApiService DELETEPaypalPaymentsPaypalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalPaymentId interface{}

		httpRes, err := apiClient.PaypalPaymentsApi.DELETEPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaypalPaymentsApiService GETPaypalGatewayIdPaypalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalGatewayId interface{}

		httpRes, err := apiClient.PaypalPaymentsApi.GETPaypalGatewayIdPaypalPayments(context.Background(), paypalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaypalPaymentsApiService GETPaypalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaypalPaymentsApi.GETPaypalPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaypalPaymentsApiService GETPaypalPaymentsPaypalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalPaymentId interface{}

		resp, httpRes, err := apiClient.PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaypalPaymentsApiService PATCHPaypalPaymentsPaypalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalPaymentId interface{}

		resp, httpRes, err := apiClient.PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaypalPaymentsApiService POSTPaypalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaypalPaymentsApi.POSTPaypalPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

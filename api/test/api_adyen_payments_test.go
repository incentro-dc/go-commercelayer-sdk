/*
Commerce Layer API

Testing AdyenPaymentsApiService

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

func Test_api_AdyenPaymentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdyenPaymentsApiService DELETEAdyenPaymentsAdyenPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenPaymentId interface{}

		httpRes, err := apiClient.AdyenPaymentsApi.DELETEAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdyenPaymentsApiService GETAdyenGatewayIdAdyenPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenGatewayId interface{}

		httpRes, err := apiClient.AdyenPaymentsApi.GETAdyenGatewayIdAdyenPayments(context.Background(), adyenGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdyenPaymentsApiService GETAdyenPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AdyenPaymentsApi.GETAdyenPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdyenPaymentsApiService GETAdyenPaymentsAdyenPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenPaymentId interface{}

		resp, httpRes, err := apiClient.AdyenPaymentsApi.GETAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdyenPaymentsApiService PATCHAdyenPaymentsAdyenPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenPaymentId interface{}

		resp, httpRes, err := apiClient.AdyenPaymentsApi.PATCHAdyenPaymentsAdyenPaymentId(context.Background(), adyenPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdyenPaymentsApiService POSTAdyenPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AdyenPaymentsApi.POSTAdyenPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

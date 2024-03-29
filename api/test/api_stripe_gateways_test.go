/*
Commerce Layer API

Testing StripeGatewaysApiService

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

func Test_api_StripeGatewaysApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StripeGatewaysApiService DELETEStripeGatewaysStripeGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripeGatewayId interface{}

		httpRes, err := apiClient.StripeGatewaysApi.DELETEStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StripeGatewaysApiService GETStripeGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StripeGatewaysApi.GETStripeGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StripeGatewaysApiService GETStripeGatewaysStripeGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripeGatewayId interface{}

		resp, httpRes, err := apiClient.StripeGatewaysApi.GETStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StripeGatewaysApiService PATCHStripeGatewaysStripeGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripeGatewayId interface{}

		resp, httpRes, err := apiClient.StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StripeGatewaysApiService POSTStripeGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StripeGatewaysApi.POSTStripeGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

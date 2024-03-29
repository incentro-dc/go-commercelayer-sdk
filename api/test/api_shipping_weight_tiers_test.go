/*
Commerce Layer API

Testing ShippingWeightTiersApiService

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

func Test_api_ShippingWeightTiersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShippingWeightTiersApiService DELETEShippingWeightTiersShippingWeightTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		httpRes, err := apiClient.ShippingWeightTiersApi.DELETEShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingWeightTiersApiService GETShippingMethodIdShippingWeightTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.ShippingWeightTiersApi.GETShippingMethodIdShippingWeightTiers(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingWeightTiersApiService GETShippingWeightTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingWeightTiersApiService GETShippingWeightTiersShippingWeightTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		resp, httpRes, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingWeightTiersApiService PATCHShippingWeightTiersShippingWeightTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingWeightTierId interface{}

		resp, httpRes, err := apiClient.ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingWeightTiersApiService POSTShippingWeightTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShippingWeightTiersApi.POSTShippingWeightTiers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

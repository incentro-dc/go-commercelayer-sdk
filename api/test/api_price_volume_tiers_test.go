/*
Commerce Layer API

Testing PriceVolumeTiersApiService

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

func Test_api_PriceVolumeTiersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PriceVolumeTiersApiService DELETEPriceVolumeTiersPriceVolumeTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		httpRes, err := apiClient.PriceVolumeTiersApi.DELETEPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceVolumeTiersApiService GETPriceIdPriceVolumeTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceId interface{}

		httpRes, err := apiClient.PriceVolumeTiersApi.GETPriceIdPriceVolumeTiers(context.Background(), priceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceVolumeTiersApiService GETPriceVolumeTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceVolumeTiersApiService GETPriceVolumeTiersPriceVolumeTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		resp, httpRes, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceVolumeTiersApiService PATCHPriceVolumeTiersPriceVolumeTierId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		resp, httpRes, err := apiClient.PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceVolumeTiersApiService POSTPriceVolumeTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PriceVolumeTiersApi.POSTPriceVolumeTiers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

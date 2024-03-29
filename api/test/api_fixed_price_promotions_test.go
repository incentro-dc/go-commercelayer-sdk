/*
Commerce Layer API

Testing FixedPricePromotionsApiService

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

func Test_api_FixedPricePromotionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FixedPricePromotionsApiService DELETEFixedPricePromotionsFixedPricePromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.FixedPricePromotionsApi.DELETEFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FixedPricePromotionsApiService GETFixedPricePromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FixedPricePromotionsApiService GETFixedPricePromotionsFixedPricePromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		resp, httpRes, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FixedPricePromotionsApiService PATCHFixedPricePromotionsFixedPricePromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		resp, httpRes, err := apiClient.FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FixedPricePromotionsApiService POSTFixedPricePromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FixedPricePromotionsApi.POSTFixedPricePromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

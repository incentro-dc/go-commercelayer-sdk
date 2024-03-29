/*
Commerce Layer API

Testing PercentageDiscountPromotionsApiService

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

func Test_api_PercentageDiscountPromotionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PercentageDiscountPromotionsApiService DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.PercentageDiscountPromotionsApi.DELETEPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PercentageDiscountPromotionsApiService GETPercentageDiscountPromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PercentageDiscountPromotionsApi.GETPercentageDiscountPromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PercentageDiscountPromotionsApiService GETPercentageDiscountPromotionsPercentageDiscountPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		resp, httpRes, err := apiClient.PercentageDiscountPromotionsApi.GETPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PercentageDiscountPromotionsApiService PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		resp, httpRes, err := apiClient.PercentageDiscountPromotionsApi.PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PercentageDiscountPromotionsApiService POSTPercentageDiscountPromotions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PercentageDiscountPromotionsApi.POSTPercentageDiscountPromotions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
Commerce Layer API

Testing SubscriptionModelsApiService

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

func Test_api_SubscriptionModelsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionModelsApiService DELETESubscriptionModelsSubscriptionModelId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionModelId interface{}

		httpRes, err := apiClient.SubscriptionModelsApi.DELETESubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService GETMarketIdSubscriptionModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		httpRes, err := apiClient.SubscriptionModelsApi.GETMarketIdSubscriptionModel(context.Background(), marketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService GETOrderSubscriptionIdSubscriptionModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.SubscriptionModelsApi.GETOrderSubscriptionIdSubscriptionModel(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService GETSubscriptionModels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SubscriptionModelsApi.GETSubscriptionModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService GETSubscriptionModelsSubscriptionModelId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionModelId interface{}

		resp, httpRes, err := apiClient.SubscriptionModelsApi.GETSubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService PATCHSubscriptionModelsSubscriptionModelId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionModelId interface{}

		resp, httpRes, err := apiClient.SubscriptionModelsApi.PATCHSubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionModelsApiService POSTSubscriptionModels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SubscriptionModelsApi.POSTSubscriptionModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

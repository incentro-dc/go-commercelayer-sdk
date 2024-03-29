/*
Commerce Layer API

Testing BingGeocodersApiService

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

func Test_api_BingGeocodersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BingGeocodersApiService DELETEBingGeocodersBingGeocoderId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bingGeocoderId interface{}

		httpRes, err := apiClient.BingGeocodersApi.DELETEBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BingGeocodersApiService GETBingGeocoders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BingGeocodersApi.GETBingGeocoders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BingGeocodersApiService GETBingGeocodersBingGeocoderId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bingGeocoderId interface{}

		resp, httpRes, err := apiClient.BingGeocodersApi.GETBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BingGeocodersApiService PATCHBingGeocodersBingGeocoderId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bingGeocoderId interface{}

		resp, httpRes, err := apiClient.BingGeocodersApi.PATCHBingGeocodersBingGeocoderId(context.Background(), bingGeocoderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BingGeocodersApiService POSTBingGeocoders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BingGeocodersApi.POSTBingGeocoders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
Commerce Layer API

Testing SatispayGatewaysApiService

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

func Test_api_SatispayGatewaysApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SatispayGatewaysApiService DELETESatispayGatewaysSatispayGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayGatewayId interface{}

		httpRes, err := apiClient.SatispayGatewaysApi.DELETESatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SatispayGatewaysApiService GETSatispayGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SatispayGatewaysApi.GETSatispayGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SatispayGatewaysApiService GETSatispayGatewaysSatispayGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayGatewayId interface{}

		resp, httpRes, err := apiClient.SatispayGatewaysApi.GETSatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SatispayGatewaysApiService PATCHSatispayGatewaysSatispayGatewayId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayGatewayId interface{}

		resp, httpRes, err := apiClient.SatispayGatewaysApi.PATCHSatispayGatewaysSatispayGatewayId(context.Background(), satispayGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SatispayGatewaysApiService POSTSatispayGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SatispayGatewaysApi.POSTSatispayGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

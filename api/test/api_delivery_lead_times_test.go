/*
Commerce Layer API

Testing DeliveryLeadTimesApiService

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

func Test_api_DeliveryLeadTimesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeliveryLeadTimesApiService DELETEDeliveryLeadTimesDeliveryLeadTimeId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		httpRes, err := apiClient.DeliveryLeadTimesApi.DELETEDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService GETDeliveryLeadTimes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService GETDeliveryLeadTimesDeliveryLeadTimeId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		resp, httpRes, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService GETShipmentIdDeliveryLeadTime", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.DeliveryLeadTimesApi.GETShipmentIdDeliveryLeadTime(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService GETShippingMethodIdDeliveryLeadTimeForShipment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.DeliveryLeadTimesApi.GETShippingMethodIdDeliveryLeadTimeForShipment(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService GETSkuIdDeliveryLeadTimes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuId interface{}

		httpRes, err := apiClient.DeliveryLeadTimesApi.GETSkuIdDeliveryLeadTimes(context.Background(), skuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService PATCHDeliveryLeadTimesDeliveryLeadTimeId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deliveryLeadTimeId interface{}

		resp, httpRes, err := apiClient.DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryLeadTimesApiService POSTDeliveryLeadTimes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeliveryLeadTimesApi.POSTDeliveryLeadTimes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

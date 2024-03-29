/*
Commerce Layer API

Testing GiftCardRecipientsApiService

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

func Test_api_GiftCardRecipientsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GiftCardRecipientsApiService DELETEGiftCardRecipientsGiftCardRecipientId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardRecipientId interface{}

		httpRes, err := apiClient.GiftCardRecipientsApi.DELETEGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardRecipientsApiService GETGiftCardIdGiftCardRecipient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.GiftCardRecipientsApi.GETGiftCardIdGiftCardRecipient(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardRecipientsApiService GETGiftCardRecipients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GiftCardRecipientsApi.GETGiftCardRecipients(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardRecipientsApiService GETGiftCardRecipientsGiftCardRecipientId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardRecipientId interface{}

		resp, httpRes, err := apiClient.GiftCardRecipientsApi.GETGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardRecipientsApiService PATCHGiftCardRecipientsGiftCardRecipientId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardRecipientId interface{}

		resp, httpRes, err := apiClient.GiftCardRecipientsApi.PATCHGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardRecipientsApiService POSTGiftCardRecipients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GiftCardRecipientsApi.POSTGiftCardRecipients(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

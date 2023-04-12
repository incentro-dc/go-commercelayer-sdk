/*
Commerce Layer API

Testing TaxRulesApiService

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

func Test_api_TaxRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaxRulesApiService DELETETaxRulesTaxRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxRuleId interface{}

		httpRes, err := apiClient.TaxRulesApi.DELETETaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxRulesApiService GETManualTaxCalculatorIdTaxRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualTaxCalculatorId interface{}

		httpRes, err := apiClient.TaxRulesApi.GETManualTaxCalculatorIdTaxRules(context.Background(), manualTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxRulesApiService GETTaxRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TaxRulesApi.GETTaxRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxRulesApiService GETTaxRulesTaxRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxRuleId interface{}

		resp, httpRes, err := apiClient.TaxRulesApi.GETTaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxRulesApiService PATCHTaxRulesTaxRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxRuleId interface{}

		resp, httpRes, err := apiClient.TaxRulesApi.PATCHTaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxRulesApiService POSTTaxRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TaxRulesApi.POSTTaxRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

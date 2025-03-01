/*
API Avito shop

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService ApiAuthPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ApiAuthPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ApiBuyItemGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var item string

		httpRes, err := apiClient.DefaultAPI.ApiBuyItemGet(context.Background(), item).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ApiInfoGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ApiInfoGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ApiSendCoinPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.DefaultAPI.ApiSendCoinPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
Clover API

Testing PrivateSubnetApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mohi-kalantari/ubicloud-go-client"
)

func Test_openapi_PrivateSubnetApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PrivateSubnetApiService CreatePrivateSubnet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string
		var privateSubnetName string

		resp, httpRes, err := apiClient.PrivateSubnetApi.CreatePrivateSubnet(context.Background(), projectId, location, privateSubnetName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService DeletePSWithId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string
		var privateSubnetId string

		httpRes, err := apiClient.PrivateSubnetApi.DeletePSWithId(context.Background(), projectId, location, privateSubnetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService DeletePrivateSubnet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string
		var privateSubnetName string

		httpRes, err := apiClient.PrivateSubnetApi.DeletePrivateSubnet(context.Background(), projectId, location, privateSubnetName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService GetPSDetailsWithId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string
		var privateSubnetId string

		resp, httpRes, err := apiClient.PrivateSubnetApi.GetPSDetailsWithId(context.Background(), projectId, location, privateSubnetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService GetPrivateSubnetDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string
		var privateSubnetName string

		resp, httpRes, err := apiClient.PrivateSubnetApi.GetPrivateSubnetDetails(context.Background(), projectId, location, privateSubnetName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService ListLocationPrivateSubnets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var location string

		resp, httpRes, err := apiClient.PrivateSubnetApi.ListLocationPrivateSubnets(context.Background(), projectId, location).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrivateSubnetApiService ListPSs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.PrivateSubnetApi.ListPSs(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

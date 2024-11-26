/*
Clover API

Testing LoadBalancerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/mohi-kalantari/ubicloud-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_LoadBalancerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LoadBalancerApiService AttachVmLocationLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.AttachVmLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService CreateLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.CreateLoadBalancer(context.Background(), projectId, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService CreateLocationLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.CreateLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService DeleteLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		httpRes, err := apiClient.LoadBalancerApi.DeleteLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService DeleteLoadBalancerWithID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerId string

		httpRes, err := apiClient.LoadBalancerApi.DeleteLoadBalancerWithID(context.Background(), projectId, location, loadBalancerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService DetachVmLocationLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.DetachVmLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService GetLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.GetLoadBalancer(context.Background(), projectId, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService GetLoadBalancerDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.GetLoadBalancerDetails(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService GetLoadBalancerDetailsWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerId string

		resp, httpRes, err := apiClient.LoadBalancerApi.GetLoadBalancerDetailsWithId(context.Background(), projectId, location, loadBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService ListLoadBalancers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.LoadBalancerApi.ListLoadBalancers(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService ListLocationLoadBalancers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string

		resp, httpRes, err := apiClient.LoadBalancerApi.ListLocationLoadBalancers(context.Background(), projectId, location).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LoadBalancerApiService PatchLocationLoadBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var loadBalancerName string

		resp, httpRes, err := apiClient.LoadBalancerApi.PatchLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

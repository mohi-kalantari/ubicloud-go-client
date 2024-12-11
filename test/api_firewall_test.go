/*
Clover API

Testing FirewallApiService

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

func Test_openapi_FirewallApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirewallApiService ActionLocationFirewallAttachSubnet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallId string

		resp, httpRes, err := apiClient.FirewallApi.ActionLocationFirewallAttachSubnet(context.Background(), location, projectId, firewallId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService ActionLocationFirewallDetachSubnet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallId string

		resp, httpRes, err := apiClient.FirewallApi.ActionLocationFirewallDetachSubnet(context.Background(), location, projectId, firewallId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService CreateFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.FirewallApi.CreateFirewall(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService CreateLocationFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallName string

		resp, httpRes, err := apiClient.FirewallApi.CreateLocationFirewall(context.Background(), location, projectId, firewallName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService DeleteFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var firewallName string

		httpRes, err := apiClient.FirewallApi.DeleteFirewall(context.Background(), projectId, firewallName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService DeleteLocationFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallName string

		httpRes, err := apiClient.FirewallApi.DeleteLocationFirewall(context.Background(), location, projectId, firewallName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService DeleteLocationFirewallWithId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallId string

		httpRes, err := apiClient.FirewallApi.DeleteLocationFirewallWithId(context.Background(), location, projectId, firewallId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.FirewallApi.GetFirewall(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetFirewallDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var firewallName string

		resp, httpRes, err := apiClient.FirewallApi.GetFirewallDetails(context.Background(), projectId, firewallName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetLocationFirewall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string

		resp, httpRes, err := apiClient.FirewallApi.GetLocationFirewall(context.Background(), location, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetLocationFirewallDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallName string

		resp, httpRes, err := apiClient.FirewallApi.GetLocationFirewallDetails(context.Background(), location, projectId, firewallName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallApiService GetLocationFirewallDetailsWithId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var location string
		var projectId string
		var firewallId string

		resp, httpRes, err := apiClient.FirewallApi.GetLocationFirewallDetailsWithId(context.Background(), location, projectId, firewallId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

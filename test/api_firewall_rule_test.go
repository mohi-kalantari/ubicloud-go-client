/*
Clover API

Testing FirewallRuleApiService

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

func Test_openapi_FirewallRuleApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirewallRuleApiService CreateFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var firewallName string

		resp, httpRes, err := apiClient.FirewallRuleApi.CreateFirewallRule(context.Background(), projectId, firewallName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService CreateLocationFirewallFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		resp, httpRes, err := apiClient.FirewallRuleApi.CreateLocationFirewallFirewallRule(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService CreateLocationFirewallFirewallRuleWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		resp, httpRes, err := apiClient.FirewallRuleApi.CreateLocationFirewallFirewallRuleWithId(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService CreateLocationFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var firewallName string
		var location string
		var projectId string

		resp, httpRes, err := apiClient.FirewallRuleApi.CreateLocationFirewallRule(context.Background(), firewallName, location, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService CreateLocationFirewallRuleWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var firewallName string
		var location string
		var projectId string

		resp, httpRes, err := apiClient.FirewallRuleApi.CreateLocationFirewallRuleWithId(context.Background(), firewallName, location, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService DeleteFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var firewallName string
		var firewallRuleId string

		httpRes, err := apiClient.FirewallRuleApi.DeleteFirewallRule(context.Background(), projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService DeleteLocationFirewallFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		httpRes, err := apiClient.FirewallRuleApi.DeleteLocationFirewallFirewallRule(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService DeleteLocationFirewallFirewallRuleWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		httpRes, err := apiClient.FirewallRuleApi.DeleteLocationFirewallFirewallRuleWithId(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService DeleteLocationPostgresFirewallRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var postgresDatabaseName string
		var firewallRuleId string

		httpRes, err := apiClient.FirewallRuleApi.DeleteLocationPostgresFirewallRule(context.Background(), projectId, location, postgresDatabaseName, firewallRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService DeleteLocationPostgresFirewallRuleWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var postgresDatabaseId string
		var projectId string
		var firewallRuleId string

		httpRes, err := apiClient.FirewallRuleApi.DeleteLocationPostgresFirewallRuleWithId(context.Background(), location, postgresDatabaseId, projectId, firewallRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService GetFirewallRuleDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var firewallName string
		var firewallRuleId string

		resp, httpRes, err := apiClient.FirewallRuleApi.GetFirewallRuleDetails(context.Background(), projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService GetLocationFirewallFirewallRuleDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		resp, httpRes, err := apiClient.FirewallRuleApi.GetLocationFirewallFirewallRuleDetails(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallRuleApiService GetLocationFirewallFirewallRuleDetailsWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string
		var firewallName string
		var firewallRuleId string

		resp, httpRes, err := apiClient.FirewallRuleApi.GetLocationFirewallFirewallRuleDetailsWithId(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

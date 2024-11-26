/*
Clover API

Testing VirtualMachineApiService

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

func Test_openapi_VirtualMachineApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VirtualMachineApiService CreateVM", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var vmName string

		resp, httpRes, err := apiClient.VirtualMachineApi.CreateVM(context.Background(), projectId, location, vmName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService DeleteVM", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var vmName string

		httpRes, err := apiClient.VirtualMachineApi.DeleteVM(context.Background(), projectId, location, vmName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService DeleteVMWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var vmId string

		httpRes, err := apiClient.VirtualMachineApi.DeleteVMWithId(context.Background(), projectId, location, vmId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService GetVMDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var vmName string

		resp, httpRes, err := apiClient.VirtualMachineApi.GetVMDetails(context.Background(), projectId, location, vmName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService GetVMDetailsWithId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string
		var location string
		var vmId string

		resp, httpRes, err := apiClient.VirtualMachineApi.GetVMDetailsWithId(context.Background(), projectId, location, vmId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService ListLocationVMs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var location string
		var projectId string

		resp, httpRes, err := apiClient.VirtualMachineApi.ListLocationVMs(context.Background(), location, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineApiService ListProjectVMs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.VirtualMachineApi.ListProjectVMs(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
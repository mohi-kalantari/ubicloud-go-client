# PatchLocationLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm of the Load Balancer | [optional] 
**DstPort** | Pointer to **int32** | Destination port for the Load Balancer | [optional] 
**HealthCheckEndpoint** | Pointer to **string** | Health check endpoint URL | [optional] 
**SrcPort** | Pointer to **int32** | Source port for the Load Balancer | [optional] 
**Vms** | Pointer to **[]string** | List of VM apids for the Load Balancer | [optional] 

## Methods

### NewPatchLocationLoadBalancerRequest

`func NewPatchLocationLoadBalancerRequest() *PatchLocationLoadBalancerRequest`

NewPatchLocationLoadBalancerRequest instantiates a new PatchLocationLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLocationLoadBalancerRequestWithDefaults

`func NewPatchLocationLoadBalancerRequestWithDefaults() *PatchLocationLoadBalancerRequest`

NewPatchLocationLoadBalancerRequestWithDefaults instantiates a new PatchLocationLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *PatchLocationLoadBalancerRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PatchLocationLoadBalancerRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PatchLocationLoadBalancerRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PatchLocationLoadBalancerRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDstPort

`func (o *PatchLocationLoadBalancerRequest) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *PatchLocationLoadBalancerRequest) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *PatchLocationLoadBalancerRequest) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *PatchLocationLoadBalancerRequest) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetHealthCheckEndpoint

`func (o *PatchLocationLoadBalancerRequest) GetHealthCheckEndpoint() string`

GetHealthCheckEndpoint returns the HealthCheckEndpoint field if non-nil, zero value otherwise.

### GetHealthCheckEndpointOk

`func (o *PatchLocationLoadBalancerRequest) GetHealthCheckEndpointOk() (*string, bool)`

GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEndpoint

`func (o *PatchLocationLoadBalancerRequest) SetHealthCheckEndpoint(v string)`

SetHealthCheckEndpoint sets HealthCheckEndpoint field to given value.

### HasHealthCheckEndpoint

`func (o *PatchLocationLoadBalancerRequest) HasHealthCheckEndpoint() bool`

HasHealthCheckEndpoint returns a boolean if a field has been set.

### GetSrcPort

`func (o *PatchLocationLoadBalancerRequest) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *PatchLocationLoadBalancerRequest) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *PatchLocationLoadBalancerRequest) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *PatchLocationLoadBalancerRequest) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetVms

`func (o *PatchLocationLoadBalancerRequest) GetVms() []string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *PatchLocationLoadBalancerRequest) GetVmsOk() (*[]string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *PatchLocationLoadBalancerRequest) SetVms(v []string)`

SetVms sets Vms field to given value.

### HasVms

`func (o *PatchLocationLoadBalancerRequest) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



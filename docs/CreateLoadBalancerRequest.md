# CreateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Algorithm of the Load Balancer | 
**DstPort** | **int32** | Destination port for the Load Balancer | 
**HealthCheckEndpoint** | Pointer to **string** | Health check endpoint URL | [optional] 
**HealthCheckProtocol** | **string** | Health check endpoint protocol | 
**PrivateSubnetId** | **string** | ID of Private Subnet | 
**SrcPort** | **int32** | Source port for the Load Balancer | 

## Methods

### NewCreateLoadBalancerRequest

`func NewCreateLoadBalancerRequest(algorithm string, dstPort int32, healthCheckProtocol string, privateSubnetId string, srcPort int32, ) *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequest instantiates a new CreateLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerRequestWithDefaults

`func NewCreateLoadBalancerRequestWithDefaults() *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequestWithDefaults instantiates a new CreateLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *CreateLoadBalancerRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CreateLoadBalancerRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CreateLoadBalancerRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetDstPort

`func (o *CreateLoadBalancerRequest) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *CreateLoadBalancerRequest) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *CreateLoadBalancerRequest) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.


### GetHealthCheckEndpoint

`func (o *CreateLoadBalancerRequest) GetHealthCheckEndpoint() string`

GetHealthCheckEndpoint returns the HealthCheckEndpoint field if non-nil, zero value otherwise.

### GetHealthCheckEndpointOk

`func (o *CreateLoadBalancerRequest) GetHealthCheckEndpointOk() (*string, bool)`

GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEndpoint

`func (o *CreateLoadBalancerRequest) SetHealthCheckEndpoint(v string)`

SetHealthCheckEndpoint sets HealthCheckEndpoint field to given value.

### HasHealthCheckEndpoint

`func (o *CreateLoadBalancerRequest) HasHealthCheckEndpoint() bool`

HasHealthCheckEndpoint returns a boolean if a field has been set.

### GetHealthCheckProtocol

`func (o *CreateLoadBalancerRequest) GetHealthCheckProtocol() string`

GetHealthCheckProtocol returns the HealthCheckProtocol field if non-nil, zero value otherwise.

### GetHealthCheckProtocolOk

`func (o *CreateLoadBalancerRequest) GetHealthCheckProtocolOk() (*string, bool)`

GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProtocol

`func (o *CreateLoadBalancerRequest) SetHealthCheckProtocol(v string)`

SetHealthCheckProtocol sets HealthCheckProtocol field to given value.


### GetPrivateSubnetId

`func (o *CreateLoadBalancerRequest) GetPrivateSubnetId() string`

GetPrivateSubnetId returns the PrivateSubnetId field if non-nil, zero value otherwise.

### GetPrivateSubnetIdOk

`func (o *CreateLoadBalancerRequest) GetPrivateSubnetIdOk() (*string, bool)`

GetPrivateSubnetIdOk returns a tuple with the PrivateSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnetId

`func (o *CreateLoadBalancerRequest) SetPrivateSubnetId(v string)`

SetPrivateSubnetId sets PrivateSubnetId field to given value.


### GetSrcPort

`func (o *CreateLoadBalancerRequest) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *CreateLoadBalancerRequest) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *CreateLoadBalancerRequest) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



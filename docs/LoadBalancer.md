# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Algorithm of the Load Balancer | 
**DstPort** | **int32** | Destination port for the Load Balancer | 
**HealthCheckEndpoint** | **string** | Health check endpoint URL | 
**HealthCheckProtocol** | **string** | Health check endpoint protocol | 
**Hostname** | **string** | Hostname of the Load Balancer | 
**Id** | **string** | ID of the Load Balancer | 
**Name** | **string** | Name of the Load Balancer | 
**SrcPort** | **int32** | Source port for the Load Balancer | 

## Methods

### NewLoadBalancer

`func NewLoadBalancer(algorithm string, dstPort int32, healthCheckEndpoint string, healthCheckProtocol string, hostname string, id string, name string, srcPort int32, ) *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *LoadBalancer) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *LoadBalancer) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *LoadBalancer) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetDstPort

`func (o *LoadBalancer) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *LoadBalancer) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *LoadBalancer) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.


### GetHealthCheckEndpoint

`func (o *LoadBalancer) GetHealthCheckEndpoint() string`

GetHealthCheckEndpoint returns the HealthCheckEndpoint field if non-nil, zero value otherwise.

### GetHealthCheckEndpointOk

`func (o *LoadBalancer) GetHealthCheckEndpointOk() (*string, bool)`

GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEndpoint

`func (o *LoadBalancer) SetHealthCheckEndpoint(v string)`

SetHealthCheckEndpoint sets HealthCheckEndpoint field to given value.


### GetHealthCheckProtocol

`func (o *LoadBalancer) GetHealthCheckProtocol() string`

GetHealthCheckProtocol returns the HealthCheckProtocol field if non-nil, zero value otherwise.

### GetHealthCheckProtocolOk

`func (o *LoadBalancer) GetHealthCheckProtocolOk() (*string, bool)`

GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProtocol

`func (o *LoadBalancer) SetHealthCheckProtocol(v string)`

SetHealthCheckProtocol sets HealthCheckProtocol field to given value.


### GetHostname

`func (o *LoadBalancer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoadBalancer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoadBalancer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetId

`func (o *LoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancer) SetName(v string)`

SetName sets Name field to given value.


### GetSrcPort

`func (o *LoadBalancer) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *LoadBalancer) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *LoadBalancer) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



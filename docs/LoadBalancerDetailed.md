# LoadBalancerDetailed

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
**Location** | **string** | Location of the Load Balancer | 
**Subnet** | **string** | Subnet of the Load Balancer | 
**Vms** | [**[]Vm**](Vm.md) |  | 

## Methods

### NewLoadBalancerDetailed

`func NewLoadBalancerDetailed(algorithm string, dstPort int32, healthCheckEndpoint string, healthCheckProtocol string, hostname string, id string, name string, srcPort int32, location string, subnet string, vms []Vm, ) *LoadBalancerDetailed`

NewLoadBalancerDetailed instantiates a new LoadBalancerDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerDetailedWithDefaults

`func NewLoadBalancerDetailedWithDefaults() *LoadBalancerDetailed`

NewLoadBalancerDetailedWithDefaults instantiates a new LoadBalancerDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *LoadBalancerDetailed) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *LoadBalancerDetailed) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *LoadBalancerDetailed) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetDstPort

`func (o *LoadBalancerDetailed) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *LoadBalancerDetailed) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *LoadBalancerDetailed) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.


### GetHealthCheckEndpoint

`func (o *LoadBalancerDetailed) GetHealthCheckEndpoint() string`

GetHealthCheckEndpoint returns the HealthCheckEndpoint field if non-nil, zero value otherwise.

### GetHealthCheckEndpointOk

`func (o *LoadBalancerDetailed) GetHealthCheckEndpointOk() (*string, bool)`

GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEndpoint

`func (o *LoadBalancerDetailed) SetHealthCheckEndpoint(v string)`

SetHealthCheckEndpoint sets HealthCheckEndpoint field to given value.


### GetHealthCheckProtocol

`func (o *LoadBalancerDetailed) GetHealthCheckProtocol() string`

GetHealthCheckProtocol returns the HealthCheckProtocol field if non-nil, zero value otherwise.

### GetHealthCheckProtocolOk

`func (o *LoadBalancerDetailed) GetHealthCheckProtocolOk() (*string, bool)`

GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProtocol

`func (o *LoadBalancerDetailed) SetHealthCheckProtocol(v string)`

SetHealthCheckProtocol sets HealthCheckProtocol field to given value.


### GetHostname

`func (o *LoadBalancerDetailed) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoadBalancerDetailed) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoadBalancerDetailed) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetId

`func (o *LoadBalancerDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerDetailed) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LoadBalancerDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetSrcPort

`func (o *LoadBalancerDetailed) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *LoadBalancerDetailed) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *LoadBalancerDetailed) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.


### GetLocation

`func (o *LoadBalancerDetailed) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoadBalancerDetailed) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoadBalancerDetailed) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetSubnet

`func (o *LoadBalancerDetailed) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *LoadBalancerDetailed) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *LoadBalancerDetailed) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetVms

`func (o *LoadBalancerDetailed) GetVms() []Vm`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *LoadBalancerDetailed) GetVmsOk() (*[]Vm, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *LoadBalancerDetailed) SetVms(v []Vm)`

SetVms sets Vms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



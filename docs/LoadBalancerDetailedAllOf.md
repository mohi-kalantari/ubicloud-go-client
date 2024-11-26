# LoadBalancerDetailedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of the Load Balancer | 
**Subnet** | **string** | Subnet of the Load Balancer | 
**Vms** | [**[]Vm**](Vm.md) |  | 

## Methods

### NewLoadBalancerDetailedAllOf

`func NewLoadBalancerDetailedAllOf(location string, subnet string, vms []Vm, ) *LoadBalancerDetailedAllOf`

NewLoadBalancerDetailedAllOf instantiates a new LoadBalancerDetailedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerDetailedAllOfWithDefaults

`func NewLoadBalancerDetailedAllOfWithDefaults() *LoadBalancerDetailedAllOf`

NewLoadBalancerDetailedAllOfWithDefaults instantiates a new LoadBalancerDetailedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *LoadBalancerDetailedAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoadBalancerDetailedAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoadBalancerDetailedAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetSubnet

`func (o *LoadBalancerDetailedAllOf) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *LoadBalancerDetailedAllOf) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *LoadBalancerDetailedAllOf) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetVms

`func (o *LoadBalancerDetailedAllOf) GetVms() []Vm`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *LoadBalancerDetailedAllOf) GetVmsOk() (*[]Vm, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *LoadBalancerDetailedAllOf) SetVms(v []Vm)`

SetVms sets Vms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



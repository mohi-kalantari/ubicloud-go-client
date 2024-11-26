# DetachVmLocationLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmId** | **string** | Apid of VM to detach from Load Balancer | 

## Methods

### NewDetachVmLocationLoadBalancerRequest

`func NewDetachVmLocationLoadBalancerRequest(vmId string, ) *DetachVmLocationLoadBalancerRequest`

NewDetachVmLocationLoadBalancerRequest instantiates a new DetachVmLocationLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachVmLocationLoadBalancerRequestWithDefaults

`func NewDetachVmLocationLoadBalancerRequestWithDefaults() *DetachVmLocationLoadBalancerRequest`

NewDetachVmLocationLoadBalancerRequestWithDefaults instantiates a new DetachVmLocationLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmId

`func (o *DetachVmLocationLoadBalancerRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *DetachVmLocationLoadBalancerRequest) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *DetachVmLocationLoadBalancerRequest) SetVmId(v string)`

SetVmId sets VmId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



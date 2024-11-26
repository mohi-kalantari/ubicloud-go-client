/*
Clover API

API for managing resources on Ubicloud

API version: 0.1.0
Contact: support@ubicloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateLoadBalancerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerRequest{}

// CreateLoadBalancerRequest struct for CreateLoadBalancerRequest
type CreateLoadBalancerRequest struct {
	// Algorithm of the Load Balancer
	Algorithm string `json:"algorithm"`
	// Destination port for the Load Balancer
	DstPort int32 `json:"dst_port"`
	// Health check endpoint URL
	HealthCheckEndpoint *string `json:"health_check_endpoint,omitempty"`
	// Health check endpoint protocol
	HealthCheckProtocol string `json:"health_check_protocol"`
	// ID of Private Subnet
	PrivateSubnetId string `json:"private_subnet_id"`
	// Source port for the Load Balancer
	SrcPort int32 `json:"src_port"`
}

// NewCreateLoadBalancerRequest instantiates a new CreateLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerRequest(algorithm string, dstPort int32, healthCheckProtocol string, privateSubnetId string, srcPort int32) *CreateLoadBalancerRequest {
	this := CreateLoadBalancerRequest{}
	this.Algorithm = algorithm
	this.DstPort = dstPort
	this.HealthCheckProtocol = healthCheckProtocol
	this.PrivateSubnetId = privateSubnetId
	this.SrcPort = srcPort
	return &this
}

// NewCreateLoadBalancerRequestWithDefaults instantiates a new CreateLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerRequestWithDefaults() *CreateLoadBalancerRequest {
	this := CreateLoadBalancerRequest{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *CreateLoadBalancerRequest) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *CreateLoadBalancerRequest) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetDstPort returns the DstPort field value
func (o *CreateLoadBalancerRequest) GetDstPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetDstPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DstPort, true
}

// SetDstPort sets field value
func (o *CreateLoadBalancerRequest) SetDstPort(v int32) {
	o.DstPort = v
}

// GetHealthCheckEndpoint returns the HealthCheckEndpoint field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetHealthCheckEndpoint() string {
	if o == nil || IsNil(o.HealthCheckEndpoint) {
		var ret string
		return ret
	}
	return *o.HealthCheckEndpoint
}

// GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetHealthCheckEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.HealthCheckEndpoint) {
		return nil, false
	}
	return o.HealthCheckEndpoint, true
}

// HasHealthCheckEndpoint returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasHealthCheckEndpoint() bool {
	if o != nil && !IsNil(o.HealthCheckEndpoint) {
		return true
	}

	return false
}

// SetHealthCheckEndpoint gets a reference to the given string and assigns it to the HealthCheckEndpoint field.
func (o *CreateLoadBalancerRequest) SetHealthCheckEndpoint(v string) {
	o.HealthCheckEndpoint = &v
}

// GetHealthCheckProtocol returns the HealthCheckProtocol field value
func (o *CreateLoadBalancerRequest) GetHealthCheckProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthCheckProtocol
}

// GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetHealthCheckProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthCheckProtocol, true
}

// SetHealthCheckProtocol sets field value
func (o *CreateLoadBalancerRequest) SetHealthCheckProtocol(v string) {
	o.HealthCheckProtocol = v
}

// GetPrivateSubnetId returns the PrivateSubnetId field value
func (o *CreateLoadBalancerRequest) GetPrivateSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateSubnetId
}

// GetPrivateSubnetIdOk returns a tuple with the PrivateSubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetPrivateSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateSubnetId, true
}

// SetPrivateSubnetId sets field value
func (o *CreateLoadBalancerRequest) SetPrivateSubnetId(v string) {
	o.PrivateSubnetId = v
}

// GetSrcPort returns the SrcPort field value
func (o *CreateLoadBalancerRequest) GetSrcPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetSrcPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcPort, true
}

// SetSrcPort sets field value
func (o *CreateLoadBalancerRequest) SetSrcPort(v int32) {
	o.SrcPort = v
}

func (o CreateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLoadBalancerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["dst_port"] = o.DstPort
	if !IsNil(o.HealthCheckEndpoint) {
		toSerialize["health_check_endpoint"] = o.HealthCheckEndpoint
	}
	toSerialize["health_check_protocol"] = o.HealthCheckProtocol
	toSerialize["private_subnet_id"] = o.PrivateSubnetId
	toSerialize["src_port"] = o.SrcPort
	return toSerialize, nil
}

type NullableCreateLoadBalancerRequest struct {
	value *CreateLoadBalancerRequest
	isSet bool
}

func (v NullableCreateLoadBalancerRequest) Get() *CreateLoadBalancerRequest {
	return v.value
}

func (v *NullableCreateLoadBalancerRequest) Set(val *CreateLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerRequest(val *CreateLoadBalancerRequest) *NullableCreateLoadBalancerRequest {
	return &NullableCreateLoadBalancerRequest{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


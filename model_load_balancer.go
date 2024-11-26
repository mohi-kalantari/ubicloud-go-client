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

// checks if the LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancer{}

// LoadBalancer struct for LoadBalancer
type LoadBalancer struct {
	// Algorithm of the Load Balancer
	Algorithm string `json:"algorithm"`
	// Destination port for the Load Balancer
	DstPort int32 `json:"dst_port"`
	// Health check endpoint URL
	HealthCheckEndpoint string `json:"health_check_endpoint"`
	// Health check endpoint protocol
	HealthCheckProtocol string `json:"health_check_protocol"`
	// Hostname of the Load Balancer
	Hostname string `json:"hostname"`
	// ID of the Load Balancer
	Id string `json:"id"`
	// Name of the Load Balancer
	Name string `json:"name"`
	// Source port for the Load Balancer
	SrcPort int32 `json:"src_port"`
}

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer(algorithm string, dstPort int32, healthCheckEndpoint string, healthCheckProtocol string, hostname string, id string, name string, srcPort int32) *LoadBalancer {
	this := LoadBalancer{}
	this.Algorithm = algorithm
	this.DstPort = dstPort
	this.HealthCheckEndpoint = healthCheckEndpoint
	this.HealthCheckProtocol = healthCheckProtocol
	this.Hostname = hostname
	this.Id = id
	this.Name = name
	this.SrcPort = srcPort
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *LoadBalancer) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *LoadBalancer) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetDstPort returns the DstPort field value
func (o *LoadBalancer) GetDstPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetDstPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DstPort, true
}

// SetDstPort sets field value
func (o *LoadBalancer) SetDstPort(v int32) {
	o.DstPort = v
}

// GetHealthCheckEndpoint returns the HealthCheckEndpoint field value
func (o *LoadBalancer) GetHealthCheckEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthCheckEndpoint
}

// GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetHealthCheckEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthCheckEndpoint, true
}

// SetHealthCheckEndpoint sets field value
func (o *LoadBalancer) SetHealthCheckEndpoint(v string) {
	o.HealthCheckEndpoint = v
}

// GetHealthCheckProtocol returns the HealthCheckProtocol field value
func (o *LoadBalancer) GetHealthCheckProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthCheckProtocol
}

// GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetHealthCheckProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthCheckProtocol, true
}

// SetHealthCheckProtocol sets field value
func (o *LoadBalancer) SetHealthCheckProtocol(v string) {
	o.HealthCheckProtocol = v
}

// GetHostname returns the Hostname field value
func (o *LoadBalancer) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *LoadBalancer) SetHostname(v string) {
	o.Hostname = v
}

// GetId returns the Id field value
func (o *LoadBalancer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LoadBalancer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancer) SetName(v string) {
	o.Name = v
}

// GetSrcPort returns the SrcPort field value
func (o *LoadBalancer) GetSrcPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSrcPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcPort, true
}

// SetSrcPort sets field value
func (o *LoadBalancer) SetSrcPort(v int32) {
	o.SrcPort = v
}

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["dst_port"] = o.DstPort
	toSerialize["health_check_endpoint"] = o.HealthCheckEndpoint
	toSerialize["health_check_protocol"] = o.HealthCheckProtocol
	toSerialize["hostname"] = o.Hostname
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["src_port"] = o.SrcPort
	return toSerialize, nil
}

type NullableLoadBalancer struct {
	value *LoadBalancer
	isSet bool
}

func (v NullableLoadBalancer) Get() *LoadBalancer {
	return v.value
}

func (v *NullableLoadBalancer) Set(val *LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancer(val *LoadBalancer) *NullableLoadBalancer {
	return &NullableLoadBalancer{value: val, isSet: true}
}

func (v NullableLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


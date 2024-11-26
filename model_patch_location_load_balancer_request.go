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

// checks if the PatchLocationLoadBalancerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchLocationLoadBalancerRequest{}

// PatchLocationLoadBalancerRequest struct for PatchLocationLoadBalancerRequest
type PatchLocationLoadBalancerRequest struct {
	// Algorithm of the Load Balancer
	Algorithm *string `json:"algorithm,omitempty"`
	// Destination port for the Load Balancer
	DstPort *int32 `json:"dst_port,omitempty"`
	// Health check endpoint URL
	HealthCheckEndpoint *string `json:"health_check_endpoint,omitempty"`
	// Source port for the Load Balancer
	SrcPort *int32 `json:"src_port,omitempty"`
	// List of VM apids for the Load Balancer
	Vms []string `json:"vms,omitempty"`
}

// NewPatchLocationLoadBalancerRequest instantiates a new PatchLocationLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchLocationLoadBalancerRequest() *PatchLocationLoadBalancerRequest {
	this := PatchLocationLoadBalancerRequest{}
	return &this
}

// NewPatchLocationLoadBalancerRequestWithDefaults instantiates a new PatchLocationLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchLocationLoadBalancerRequestWithDefaults() *PatchLocationLoadBalancerRequest {
	this := PatchLocationLoadBalancerRequest{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PatchLocationLoadBalancerRequest) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLocationLoadBalancerRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PatchLocationLoadBalancerRequest) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *PatchLocationLoadBalancerRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *PatchLocationLoadBalancerRequest) GetDstPort() int32 {
	if o == nil || IsNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLocationLoadBalancerRequest) GetDstPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DstPort) {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *PatchLocationLoadBalancerRequest) HasDstPort() bool {
	if o != nil && !IsNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *PatchLocationLoadBalancerRequest) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetHealthCheckEndpoint returns the HealthCheckEndpoint field value if set, zero value otherwise.
func (o *PatchLocationLoadBalancerRequest) GetHealthCheckEndpoint() string {
	if o == nil || IsNil(o.HealthCheckEndpoint) {
		var ret string
		return ret
	}
	return *o.HealthCheckEndpoint
}

// GetHealthCheckEndpointOk returns a tuple with the HealthCheckEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLocationLoadBalancerRequest) GetHealthCheckEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.HealthCheckEndpoint) {
		return nil, false
	}
	return o.HealthCheckEndpoint, true
}

// HasHealthCheckEndpoint returns a boolean if a field has been set.
func (o *PatchLocationLoadBalancerRequest) HasHealthCheckEndpoint() bool {
	if o != nil && !IsNil(o.HealthCheckEndpoint) {
		return true
	}

	return false
}

// SetHealthCheckEndpoint gets a reference to the given string and assigns it to the HealthCheckEndpoint field.
func (o *PatchLocationLoadBalancerRequest) SetHealthCheckEndpoint(v string) {
	o.HealthCheckEndpoint = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *PatchLocationLoadBalancerRequest) GetSrcPort() int32 {
	if o == nil || IsNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLocationLoadBalancerRequest) GetSrcPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *PatchLocationLoadBalancerRequest) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *PatchLocationLoadBalancerRequest) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *PatchLocationLoadBalancerRequest) GetVms() []string {
	if o == nil || IsNil(o.Vms) {
		var ret []string
		return ret
	}
	return o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLocationLoadBalancerRequest) GetVmsOk() ([]string, bool) {
	if o == nil || IsNil(o.Vms) {
		return nil, false
	}
	return o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *PatchLocationLoadBalancerRequest) HasVms() bool {
	if o != nil && !IsNil(o.Vms) {
		return true
	}

	return false
}

// SetVms gets a reference to the given []string and assigns it to the Vms field.
func (o *PatchLocationLoadBalancerRequest) SetVms(v []string) {
	o.Vms = v
}

func (o PatchLocationLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLocationLoadBalancerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.DstPort) {
		toSerialize["dst_port"] = o.DstPort
	}
	if !IsNil(o.HealthCheckEndpoint) {
		toSerialize["health_check_endpoint"] = o.HealthCheckEndpoint
	}
	if !IsNil(o.SrcPort) {
		toSerialize["src_port"] = o.SrcPort
	}
	if !IsNil(o.Vms) {
		toSerialize["vms"] = o.Vms
	}
	return toSerialize, nil
}

type NullablePatchLocationLoadBalancerRequest struct {
	value *PatchLocationLoadBalancerRequest
	isSet bool
}

func (v NullablePatchLocationLoadBalancerRequest) Get() *PatchLocationLoadBalancerRequest {
	return v.value
}

func (v *NullablePatchLocationLoadBalancerRequest) Set(val *PatchLocationLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLocationLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLocationLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLocationLoadBalancerRequest(val *PatchLocationLoadBalancerRequest) *NullablePatchLocationLoadBalancerRequest {
	return &NullablePatchLocationLoadBalancerRequest{value: val, isSet: true}
}

func (v NullablePatchLocationLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLocationLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


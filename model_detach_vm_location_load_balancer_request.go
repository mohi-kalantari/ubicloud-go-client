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

// checks if the DetachVmLocationLoadBalancerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetachVmLocationLoadBalancerRequest{}

// DetachVmLocationLoadBalancerRequest struct for DetachVmLocationLoadBalancerRequest
type DetachVmLocationLoadBalancerRequest struct {
	// Apid of VM to detach from Load Balancer
	VmId string `json:"vm_id"`
}

// NewDetachVmLocationLoadBalancerRequest instantiates a new DetachVmLocationLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetachVmLocationLoadBalancerRequest(vmId string) *DetachVmLocationLoadBalancerRequest {
	this := DetachVmLocationLoadBalancerRequest{}
	this.VmId = vmId
	return &this
}

// NewDetachVmLocationLoadBalancerRequestWithDefaults instantiates a new DetachVmLocationLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetachVmLocationLoadBalancerRequestWithDefaults() *DetachVmLocationLoadBalancerRequest {
	this := DetachVmLocationLoadBalancerRequest{}
	return &this
}

// GetVmId returns the VmId field value
func (o *DetachVmLocationLoadBalancerRequest) GetVmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value
// and a boolean to check if the value has been set.
func (o *DetachVmLocationLoadBalancerRequest) GetVmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmId, true
}

// SetVmId sets field value
func (o *DetachVmLocationLoadBalancerRequest) SetVmId(v string) {
	o.VmId = v
}

func (o DetachVmLocationLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetachVmLocationLoadBalancerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vm_id"] = o.VmId
	return toSerialize, nil
}

type NullableDetachVmLocationLoadBalancerRequest struct {
	value *DetachVmLocationLoadBalancerRequest
	isSet bool
}

func (v NullableDetachVmLocationLoadBalancerRequest) Get() *DetachVmLocationLoadBalancerRequest {
	return v.value
}

func (v *NullableDetachVmLocationLoadBalancerRequest) Set(val *DetachVmLocationLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetachVmLocationLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetachVmLocationLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetachVmLocationLoadBalancerRequest(val *DetachVmLocationLoadBalancerRequest) *NullableDetachVmLocationLoadBalancerRequest {
	return &NullableDetachVmLocationLoadBalancerRequest{value: val, isSet: true}
}

func (v NullableDetachVmLocationLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetachVmLocationLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


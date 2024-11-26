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

// checks if the LoadBalancerDetailedAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerDetailedAllOf{}

// LoadBalancerDetailedAllOf struct for LoadBalancerDetailedAllOf
type LoadBalancerDetailedAllOf struct {
	// Location of the Load Balancer
	Location string `json:"location"`
	// Subnet of the Load Balancer
	Subnet string `json:"subnet"`
	Vms []Vm `json:"vms"`
}

// NewLoadBalancerDetailedAllOf instantiates a new LoadBalancerDetailedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerDetailedAllOf(location string, subnet string, vms []Vm) *LoadBalancerDetailedAllOf {
	this := LoadBalancerDetailedAllOf{}
	this.Location = location
	this.Subnet = subnet
	this.Vms = vms
	return &this
}

// NewLoadBalancerDetailedAllOfWithDefaults instantiates a new LoadBalancerDetailedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerDetailedAllOfWithDefaults() *LoadBalancerDetailedAllOf {
	this := LoadBalancerDetailedAllOf{}
	return &this
}

// GetLocation returns the Location field value
func (o *LoadBalancerDetailedAllOf) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetailedAllOf) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LoadBalancerDetailedAllOf) SetLocation(v string) {
	o.Location = v
}

// GetSubnet returns the Subnet field value
func (o *LoadBalancerDetailedAllOf) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetailedAllOf) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *LoadBalancerDetailedAllOf) SetSubnet(v string) {
	o.Subnet = v
}

// GetVms returns the Vms field value
func (o *LoadBalancerDetailedAllOf) GetVms() []Vm {
	if o == nil {
		var ret []Vm
		return ret
	}

	return o.Vms
}

// GetVmsOk returns a tuple with the Vms field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetailedAllOf) GetVmsOk() ([]Vm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vms, true
}

// SetVms sets field value
func (o *LoadBalancerDetailedAllOf) SetVms(v []Vm) {
	o.Vms = v
}

func (o LoadBalancerDetailedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerDetailedAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["subnet"] = o.Subnet
	toSerialize["vms"] = o.Vms
	return toSerialize, nil
}

type NullableLoadBalancerDetailedAllOf struct {
	value *LoadBalancerDetailedAllOf
	isSet bool
}

func (v NullableLoadBalancerDetailedAllOf) Get() *LoadBalancerDetailedAllOf {
	return v.value
}

func (v *NullableLoadBalancerDetailedAllOf) Set(val *LoadBalancerDetailedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerDetailedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerDetailedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerDetailedAllOf(val *LoadBalancerDetailedAllOf) *NullableLoadBalancerDetailedAllOf {
	return &NullableLoadBalancerDetailedAllOf{value: val, isSet: true}
}

func (v NullableLoadBalancerDetailedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerDetailedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


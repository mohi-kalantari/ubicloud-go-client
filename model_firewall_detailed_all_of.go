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

// checks if the FirewallDetailedAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallDetailedAllOf{}

// FirewallDetailedAllOf struct for FirewallDetailedAllOf
type FirewallDetailedAllOf struct {
	// List of private subnets
	PrivateSubnets []string `json:"private_subnets"`
}

// NewFirewallDetailedAllOf instantiates a new FirewallDetailedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallDetailedAllOf(privateSubnets []string) *FirewallDetailedAllOf {
	this := FirewallDetailedAllOf{}
	this.PrivateSubnets = privateSubnets
	return &this
}

// NewFirewallDetailedAllOfWithDefaults instantiates a new FirewallDetailedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallDetailedAllOfWithDefaults() *FirewallDetailedAllOf {
	this := FirewallDetailedAllOf{}
	return &this
}

// GetPrivateSubnets returns the PrivateSubnets field value
func (o *FirewallDetailedAllOf) GetPrivateSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailedAllOf) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// SetPrivateSubnets sets field value
func (o *FirewallDetailedAllOf) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

func (o FirewallDetailedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallDetailedAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["private_subnets"] = o.PrivateSubnets
	return toSerialize, nil
}

type NullableFirewallDetailedAllOf struct {
	value *FirewallDetailedAllOf
	isSet bool
}

func (v NullableFirewallDetailedAllOf) Get() *FirewallDetailedAllOf {
	return v.value
}

func (v *NullableFirewallDetailedAllOf) Set(val *FirewallDetailedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallDetailedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallDetailedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallDetailedAllOf(val *FirewallDetailedAllOf) *NullableFirewallDetailedAllOf {
	return &NullableFirewallDetailedAllOf{value: val, isSet: true}
}

func (v NullableFirewallDetailedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallDetailedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the VmDetailedAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmDetailedAllOf{}

// VmDetailedAllOf struct for VmDetailedAllOf
type VmDetailedAllOf struct {
	// List of firewalls
	Firewalls []Firewall `json:"firewalls"`
	// Private IPv4 address
	PrivateIpv4 string `json:"private_ipv4"`
	// Private IPv6 address
	PrivateIpv6 string `json:"private_ipv6"`
	// Subnet of the VM
	Subnet string `json:"subnet"`
	// Unix user of the VM
	UnixUser string `json:"unix_user"`
}

// NewVmDetailedAllOf instantiates a new VmDetailedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmDetailedAllOf(firewalls []Firewall, privateIpv4 string, privateIpv6 string, subnet string, unixUser string) *VmDetailedAllOf {
	this := VmDetailedAllOf{}
	this.Firewalls = firewalls
	this.PrivateIpv4 = privateIpv4
	this.PrivateIpv6 = privateIpv6
	this.Subnet = subnet
	this.UnixUser = unixUser
	return &this
}

// NewVmDetailedAllOfWithDefaults instantiates a new VmDetailedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmDetailedAllOfWithDefaults() *VmDetailedAllOf {
	this := VmDetailedAllOf{}
	return &this
}

// GetFirewalls returns the Firewalls field value
func (o *VmDetailedAllOf) GetFirewalls() []Firewall {
	if o == nil {
		var ret []Firewall
		return ret
	}

	return o.Firewalls
}

// GetFirewallsOk returns a tuple with the Firewalls field value
// and a boolean to check if the value has been set.
func (o *VmDetailedAllOf) GetFirewallsOk() ([]Firewall, bool) {
	if o == nil {
		return nil, false
	}
	return o.Firewalls, true
}

// SetFirewalls sets field value
func (o *VmDetailedAllOf) SetFirewalls(v []Firewall) {
	o.Firewalls = v
}

// GetPrivateIpv4 returns the PrivateIpv4 field value
func (o *VmDetailedAllOf) GetPrivateIpv4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateIpv4
}

// GetPrivateIpv4Ok returns a tuple with the PrivateIpv4 field value
// and a boolean to check if the value has been set.
func (o *VmDetailedAllOf) GetPrivateIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateIpv4, true
}

// SetPrivateIpv4 sets field value
func (o *VmDetailedAllOf) SetPrivateIpv4(v string) {
	o.PrivateIpv4 = v
}

// GetPrivateIpv6 returns the PrivateIpv6 field value
func (o *VmDetailedAllOf) GetPrivateIpv6() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateIpv6
}

// GetPrivateIpv6Ok returns a tuple with the PrivateIpv6 field value
// and a boolean to check if the value has been set.
func (o *VmDetailedAllOf) GetPrivateIpv6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateIpv6, true
}

// SetPrivateIpv6 sets field value
func (o *VmDetailedAllOf) SetPrivateIpv6(v string) {
	o.PrivateIpv6 = v
}

// GetSubnet returns the Subnet field value
func (o *VmDetailedAllOf) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *VmDetailedAllOf) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *VmDetailedAllOf) SetSubnet(v string) {
	o.Subnet = v
}

// GetUnixUser returns the UnixUser field value
func (o *VmDetailedAllOf) GetUnixUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnixUser
}

// GetUnixUserOk returns a tuple with the UnixUser field value
// and a boolean to check if the value has been set.
func (o *VmDetailedAllOf) GetUnixUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnixUser, true
}

// SetUnixUser sets field value
func (o *VmDetailedAllOf) SetUnixUser(v string) {
	o.UnixUser = v
}

func (o VmDetailedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmDetailedAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firewalls"] = o.Firewalls
	toSerialize["private_ipv4"] = o.PrivateIpv4
	toSerialize["private_ipv6"] = o.PrivateIpv6
	toSerialize["subnet"] = o.Subnet
	toSerialize["unix_user"] = o.UnixUser
	return toSerialize, nil
}

type NullableVmDetailedAllOf struct {
	value *VmDetailedAllOf
	isSet bool
}

func (v NullableVmDetailedAllOf) Get() *VmDetailedAllOf {
	return v.value
}

func (v *NullableVmDetailedAllOf) Set(val *VmDetailedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVmDetailedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVmDetailedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmDetailedAllOf(val *VmDetailedAllOf) *NullableVmDetailedAllOf {
	return &NullableVmDetailedAllOf{value: val, isSet: true}
}

func (v NullableVmDetailedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmDetailedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the PrivateSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateSubnet{}

// PrivateSubnet struct for PrivateSubnet
type PrivateSubnet struct {
	Firewalls []Firewall `json:"firewalls"`
	// ID of the subnet
	Id string `json:"id"`
	// Location of the subnet
	Location string `json:"location"`
	// Name of the subnet
	Name string `json:"name"`
	// IPv4 CIDR of the subnet
	Net4 string `json:"net4"`
	// IPv6 CIDR of the subnet
	Net6 string `json:"net6"`
	// List of NICs
	Nics []Nic `json:"nics"`
	// State of the subnet
	State string `json:"state"`
}

// NewPrivateSubnet instantiates a new PrivateSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateSubnet(firewalls []Firewall, id string, location string, name string, net4 string, net6 string, nics []Nic, state string) *PrivateSubnet {
	this := PrivateSubnet{}
	this.Firewalls = firewalls
	this.Id = id
	this.Location = location
	this.Name = name
	this.Net4 = net4
	this.Net6 = net6
	this.Nics = nics
	this.State = state
	return &this
}

// NewPrivateSubnetWithDefaults instantiates a new PrivateSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateSubnetWithDefaults() *PrivateSubnet {
	this := PrivateSubnet{}
	return &this
}

// GetFirewalls returns the Firewalls field value
func (o *PrivateSubnet) GetFirewalls() []Firewall {
	if o == nil {
		var ret []Firewall
		return ret
	}

	return o.Firewalls
}

// GetFirewallsOk returns a tuple with the Firewalls field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetFirewallsOk() ([]Firewall, bool) {
	if o == nil {
		return nil, false
	}
	return o.Firewalls, true
}

// SetFirewalls sets field value
func (o *PrivateSubnet) SetFirewalls(v []Firewall) {
	o.Firewalls = v
}

// GetId returns the Id field value
func (o *PrivateSubnet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateSubnet) SetId(v string) {
	o.Id = v
}

// GetLocation returns the Location field value
func (o *PrivateSubnet) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PrivateSubnet) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value
func (o *PrivateSubnet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateSubnet) SetName(v string) {
	o.Name = v
}

// GetNet4 returns the Net4 field value
func (o *PrivateSubnet) GetNet4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Net4
}

// GetNet4Ok returns a tuple with the Net4 field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetNet4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Net4, true
}

// SetNet4 sets field value
func (o *PrivateSubnet) SetNet4(v string) {
	o.Net4 = v
}

// GetNet6 returns the Net6 field value
func (o *PrivateSubnet) GetNet6() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Net6
}

// GetNet6Ok returns a tuple with the Net6 field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetNet6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Net6, true
}

// SetNet6 sets field value
func (o *PrivateSubnet) SetNet6(v string) {
	o.Net6 = v
}

// GetNics returns the Nics field value
func (o *PrivateSubnet) GetNics() []Nic {
	if o == nil {
		var ret []Nic
		return ret
	}

	return o.Nics
}

// GetNicsOk returns a tuple with the Nics field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetNicsOk() ([]Nic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nics, true
}

// SetNics sets field value
func (o *PrivateSubnet) SetNics(v []Nic) {
	o.Nics = v
}

// GetState returns the State field value
func (o *PrivateSubnet) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PrivateSubnet) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PrivateSubnet) SetState(v string) {
	o.State = v
}

func (o PrivateSubnet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firewalls"] = o.Firewalls
	toSerialize["id"] = o.Id
	toSerialize["location"] = o.Location
	toSerialize["name"] = o.Name
	toSerialize["net4"] = o.Net4
	toSerialize["net6"] = o.Net6
	toSerialize["nics"] = o.Nics
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullablePrivateSubnet struct {
	value *PrivateSubnet
	isSet bool
}

func (v NullablePrivateSubnet) Get() *PrivateSubnet {
	return v.value
}

func (v *NullablePrivateSubnet) Set(val *PrivateSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateSubnet(val *PrivateSubnet) *NullablePrivateSubnet {
	return &NullablePrivateSubnet{value: val, isSet: true}
}

func (v NullablePrivateSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



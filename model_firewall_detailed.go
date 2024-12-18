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

// checks if the FirewallDetailed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallDetailed{}

// FirewallDetailed struct for FirewallDetailed
type FirewallDetailed struct {
	// Description of the firewall
	Description string `json:"description"`
	// List of firewall rules
	FirewallRules []FirewallRule `json:"firewall_rules"`
	// ID of the firewall
	Id string `json:"id"`
	// Location of the the firewall
	Location string `json:"location"`
	// Name of the firewall
	Name string `json:"name"`
	// List of private subnets
	PrivateSubnets []string `json:"private_subnets"`
}

// NewFirewallDetailed instantiates a new FirewallDetailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallDetailed(description string, firewallRules []FirewallRule, id string, location string, name string, privateSubnets []string) *FirewallDetailed {
	this := FirewallDetailed{}
	this.Description = description
	this.FirewallRules = firewallRules
	this.Id = id
	this.Location = location
	this.Name = name
	this.PrivateSubnets = privateSubnets
	return &this
}

// NewFirewallDetailedWithDefaults instantiates a new FirewallDetailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallDetailedWithDefaults() *FirewallDetailed {
	this := FirewallDetailed{}
	return &this
}

// GetDescription returns the Description field value
func (o *FirewallDetailed) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FirewallDetailed) SetDescription(v string) {
	o.Description = v
}

// GetFirewallRules returns the FirewallRules field value
func (o *FirewallDetailed) GetFirewallRules() []FirewallRule {
	if o == nil {
		var ret []FirewallRule
		return ret
	}

	return o.FirewallRules
}

// GetFirewallRulesOk returns a tuple with the FirewallRules field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetFirewallRulesOk() ([]FirewallRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirewallRules, true
}

// SetFirewallRules sets field value
func (o *FirewallDetailed) SetFirewallRules(v []FirewallRule) {
	o.FirewallRules = v
}

// GetId returns the Id field value
func (o *FirewallDetailed) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FirewallDetailed) SetId(v string) {
	o.Id = v
}

// GetLocation returns the Location field value
func (o *FirewallDetailed) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *FirewallDetailed) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value
func (o *FirewallDetailed) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FirewallDetailed) SetName(v string) {
	o.Name = v
}

// GetPrivateSubnets returns the PrivateSubnets field value
func (o *FirewallDetailed) GetPrivateSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value
// and a boolean to check if the value has been set.
func (o *FirewallDetailed) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// SetPrivateSubnets sets field value
func (o *FirewallDetailed) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

func (o FirewallDetailed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallDetailed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["firewall_rules"] = o.FirewallRules
	toSerialize["id"] = o.Id
	toSerialize["location"] = o.Location
	toSerialize["name"] = o.Name
	toSerialize["private_subnets"] = o.PrivateSubnets
	return toSerialize, nil
}

type NullableFirewallDetailed struct {
	value *FirewallDetailed
	isSet bool
}

func (v NullableFirewallDetailed) Get() *FirewallDetailed {
	return v.value
}

func (v *NullableFirewallDetailed) Set(val *FirewallDetailed) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallDetailed) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallDetailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallDetailed(val *FirewallDetailed) *NullableFirewallDetailed {
	return &NullableFirewallDetailed{value: val, isSet: true}
}

func (v NullableFirewallDetailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallDetailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



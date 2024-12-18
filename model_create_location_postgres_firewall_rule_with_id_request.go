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

// checks if the CreateLocationPostgresFirewallRuleWithIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLocationPostgresFirewallRuleWithIdRequest{}

// CreateLocationPostgresFirewallRuleWithIdRequest struct for CreateLocationPostgresFirewallRuleWithIdRequest
type CreateLocationPostgresFirewallRuleWithIdRequest struct {
	// CIDR of the firewall rule
	Cidr string `json:"cidr"`
}

// NewCreateLocationPostgresFirewallRuleWithIdRequest instantiates a new CreateLocationPostgresFirewallRuleWithIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLocationPostgresFirewallRuleWithIdRequest(cidr string) *CreateLocationPostgresFirewallRuleWithIdRequest {
	this := CreateLocationPostgresFirewallRuleWithIdRequest{}
	this.Cidr = cidr
	return &this
}

// NewCreateLocationPostgresFirewallRuleWithIdRequestWithDefaults instantiates a new CreateLocationPostgresFirewallRuleWithIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLocationPostgresFirewallRuleWithIdRequestWithDefaults() *CreateLocationPostgresFirewallRuleWithIdRequest {
	this := CreateLocationPostgresFirewallRuleWithIdRequest{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *CreateLocationPostgresFirewallRuleWithIdRequest) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *CreateLocationPostgresFirewallRuleWithIdRequest) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *CreateLocationPostgresFirewallRuleWithIdRequest) SetCidr(v string) {
	o.Cidr = v
}

func (o CreateLocationPostgresFirewallRuleWithIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLocationPostgresFirewallRuleWithIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	return toSerialize, nil
}

type NullableCreateLocationPostgresFirewallRuleWithIdRequest struct {
	value *CreateLocationPostgresFirewallRuleWithIdRequest
	isSet bool
}

func (v NullableCreateLocationPostgresFirewallRuleWithIdRequest) Get() *CreateLocationPostgresFirewallRuleWithIdRequest {
	return v.value
}

func (v *NullableCreateLocationPostgresFirewallRuleWithIdRequest) Set(val *CreateLocationPostgresFirewallRuleWithIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLocationPostgresFirewallRuleWithIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLocationPostgresFirewallRuleWithIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLocationPostgresFirewallRuleWithIdRequest(val *CreateLocationPostgresFirewallRuleWithIdRequest) *NullableCreateLocationPostgresFirewallRuleWithIdRequest {
	return &NullableCreateLocationPostgresFirewallRuleWithIdRequest{value: val, isSet: true}
}

func (v NullableCreateLocationPostgresFirewallRuleWithIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLocationPostgresFirewallRuleWithIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



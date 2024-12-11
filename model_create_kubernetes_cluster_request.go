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

// checks if the CreateKubernetesClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKubernetesClusterRequest{}

// CreateKubernetesClusterRequest struct for CreateKubernetesClusterRequest
type CreateKubernetesClusterRequest struct {
	// Name of subnet
	Subnet string `json:"subnet"`
	// Location of the kubernetes cluster
	Location string `json:"location"`
	// Version of the kubernetes cluster
	KubernetesVersion string `json:"kubernetes_version"`
}

// NewCreateKubernetesClusterRequest instantiates a new CreateKubernetesClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKubernetesClusterRequest(subnet string, location string, kubernetesVersion string) *CreateKubernetesClusterRequest {
	this := CreateKubernetesClusterRequest{}
	this.Subnet = subnet
	this.Location = location
	this.KubernetesVersion = kubernetesVersion
	return &this
}

// NewCreateKubernetesClusterRequestWithDefaults instantiates a new CreateKubernetesClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKubernetesClusterRequestWithDefaults() *CreateKubernetesClusterRequest {
	this := CreateKubernetesClusterRequest{}
	return &this
}

// GetSubnet returns the Subnet field value
func (o *CreateKubernetesClusterRequest) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *CreateKubernetesClusterRequest) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *CreateKubernetesClusterRequest) SetSubnet(v string) {
	o.Subnet = v
}

// GetLocation returns the Location field value
func (o *CreateKubernetesClusterRequest) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *CreateKubernetesClusterRequest) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *CreateKubernetesClusterRequest) SetLocation(v string) {
	o.Location = v
}

// GetKubernetesVersion returns the KubernetesVersion field value
func (o *CreateKubernetesClusterRequest) GetKubernetesVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value
// and a boolean to check if the value has been set.
func (o *CreateKubernetesClusterRequest) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesVersion, true
}

// SetKubernetesVersion sets field value
func (o *CreateKubernetesClusterRequest) SetKubernetesVersion(v string) {
	o.KubernetesVersion = v
}

func (o CreateKubernetesClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKubernetesClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subnet"] = o.Subnet
	toSerialize["location"] = o.Location
	toSerialize["kubernetes_version"] = o.KubernetesVersion
	return toSerialize, nil
}

type NullableCreateKubernetesClusterRequest struct {
	value *CreateKubernetesClusterRequest
	isSet bool
}

func (v NullableCreateKubernetesClusterRequest) Get() *CreateKubernetesClusterRequest {
	return v.value
}

func (v *NullableCreateKubernetesClusterRequest) Set(val *CreateKubernetesClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKubernetesClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKubernetesClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKubernetesClusterRequest(val *CreateKubernetesClusterRequest) *NullableCreateKubernetesClusterRequest {
	return &NullableCreateKubernetesClusterRequest{value: val, isSet: true}
}

func (v NullableCreateKubernetesClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKubernetesClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


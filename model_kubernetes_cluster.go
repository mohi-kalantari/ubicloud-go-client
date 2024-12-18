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

// checks if the KubernetesCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCluster{}

// KubernetesCluster struct for KubernetesCluster
type KubernetesCluster struct {
	// ID of the kubernetes cluster
	Id string `json:"id"`
	// Name of the kubernetes cluster
	Name string `json:"name"`
	// Number of control plane replicas
	Replica int32 `json:"replica"`
	// Version of the kubernetes cluster
	KubernetesVersion string `json:"kubernetes_version"`
	// Name of the kubernetes cluster's subnet
	Subnet string `json:"subnet"`
	// Location of the kubernetes cluster
	Location string `json:"location"`
}

// NewKubernetesCluster instantiates a new KubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCluster(id string, name string, replica int32, kubernetesVersion string, subnet string, location string) *KubernetesCluster {
	this := KubernetesCluster{}
	this.Id = id
	this.Name = name
	this.Replica = replica
	this.KubernetesVersion = kubernetesVersion
	this.Subnet = subnet
	this.Location = location
	return &this
}

// NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterWithDefaults() *KubernetesCluster {
	this := KubernetesCluster{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesCluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesCluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesCluster) SetName(v string) {
	o.Name = v
}

// GetReplica returns the Replica field value
func (o *KubernetesCluster) GetReplica() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Replica
}

// GetReplicaOk returns a tuple with the Replica field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetReplicaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replica, true
}

// SetReplica sets field value
func (o *KubernetesCluster) SetReplica(v int32) {
	o.Replica = v
}

// GetKubernetesVersion returns the KubernetesVersion field value
func (o *KubernetesCluster) GetKubernetesVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesVersion, true
}

// SetKubernetesVersion sets field value
func (o *KubernetesCluster) SetKubernetesVersion(v string) {
	o.KubernetesVersion = v
}

// GetSubnet returns the Subnet field value
func (o *KubernetesCluster) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *KubernetesCluster) SetSubnet(v string) {
	o.Subnet = v
}

// GetLocation returns the Location field value
func (o *KubernetesCluster) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *KubernetesCluster) SetLocation(v string) {
	o.Location = v
}

func (o KubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["replica"] = o.Replica
	toSerialize["kubernetes_version"] = o.KubernetesVersion
	toSerialize["subnet"] = o.Subnet
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableKubernetesCluster struct {
	value *KubernetesCluster
	isSet bool
}

func (v NullableKubernetesCluster) Get() *KubernetesCluster {
	return v.value
}

func (v *NullableKubernetesCluster) Set(val *KubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCluster(val *KubernetesCluster) *NullableKubernetesCluster {
	return &NullableKubernetesCluster{value: val, isSet: true}
}

func (v NullableKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



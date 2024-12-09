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

// checks if the Vm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vm{}

// Vm struct for Vm
type Vm struct {
	// ID of the VM
	Id string `json:"id"`
	// IPv4 address
	Ip4 NullableString `json:"ip4"`
	// Whether IPv4 is enabled
	Ip4Enabled bool `json:"ip4_enabled"`
	// IPv6 address
	Ip6 NullableString `json:"ip6"`
	// Location of the VM
	Location string `json:"location"`
	// Name of the VM
	Name string `json:"name"`
	// Size of the underlying VM
	Size string `json:"size"`
	// State of the VM
	State string `json:"state"`
	// Storage size in GiB
	StorageSizeGib int32 `json:"storage_size_gib"`
	// Unix user of the VM
	UnixUser string `json:"unix_user"`
}

// NewVm instantiates a new Vm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVm(id string, ip4 NullableString, ip4Enabled bool, ip6 NullableString, location string, name string, size string, state string, storageSizeGib int32, unixUser string) *Vm {
	this := Vm{}
	this.Id = id
	this.Ip4 = ip4
	this.Ip4Enabled = ip4Enabled
	this.Ip6 = ip6
	this.Location = location
	this.Name = name
	this.Size = size
	this.State = state
	this.StorageSizeGib = storageSizeGib
	this.UnixUser = unixUser
	return &this
}

// NewVmWithDefaults instantiates a new Vm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmWithDefaults() *Vm {
	this := Vm{}
	return &this
}

// GetId returns the Id field value
func (o *Vm) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Vm) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Vm) SetId(v string) {
	o.Id = v
}

// GetIp4 returns the Ip4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vm) GetIp4() string {
	if o == nil || o.Ip4.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip4.Get()
}

// GetIp4Ok returns a tuple with the Ip4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vm) GetIp4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip4.Get(), o.Ip4.IsSet()
}

// SetIp4 sets field value
func (o *Vm) SetIp4(v string) {
	o.Ip4.Set(&v)
}

// GetIp4Enabled returns the Ip4Enabled field value
func (o *Vm) GetIp4Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ip4Enabled
}

// GetIp4EnabledOk returns a tuple with the Ip4Enabled field value
// and a boolean to check if the value has been set.
func (o *Vm) GetIp4EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip4Enabled, true
}

// SetIp4Enabled sets field value
func (o *Vm) SetIp4Enabled(v bool) {
	o.Ip4Enabled = v
}

// GetIp6 returns the Ip6 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vm) GetIp6() string {
	if o == nil || o.Ip6.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip6.Get()
}

// GetIp6Ok returns a tuple with the Ip6 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vm) GetIp6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip6.Get(), o.Ip6.IsSet()
}

// SetIp6 sets field value
func (o *Vm) SetIp6(v string) {
	o.Ip6.Set(&v)
}

// GetLocation returns the Location field value
func (o *Vm) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Vm) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Vm) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value
func (o *Vm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Vm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Vm) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value
func (o *Vm) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Vm) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Vm) SetSize(v string) {
	o.Size = v
}

// GetState returns the State field value
func (o *Vm) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Vm) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Vm) SetState(v string) {
	o.State = v
}

// GetStorageSizeGib returns the StorageSizeGib field value
func (o *Vm) GetStorageSizeGib() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageSizeGib
}

// GetStorageSizeGibOk returns a tuple with the StorageSizeGib field value
// and a boolean to check if the value has been set.
func (o *Vm) GetStorageSizeGibOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageSizeGib, true
}

// SetStorageSizeGib sets field value
func (o *Vm) SetStorageSizeGib(v int32) {
	o.StorageSizeGib = v
}

// GetUnixUser returns the UnixUser field value
func (o *Vm) GetUnixUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnixUser
}

// GetUnixUserOk returns a tuple with the UnixUser field value
// and a boolean to check if the value has been set.
func (o *Vm) GetUnixUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnixUser, true
}

// SetUnixUser sets field value
func (o *Vm) SetUnixUser(v string) {
	o.UnixUser = v
}

func (o Vm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ip4"] = o.Ip4.Get()
	toSerialize["ip4_enabled"] = o.Ip4Enabled
	toSerialize["ip6"] = o.Ip6.Get()
	toSerialize["location"] = o.Location
	toSerialize["name"] = o.Name
	toSerialize["size"] = o.Size
	toSerialize["state"] = o.State
	toSerialize["storage_size_gib"] = o.StorageSizeGib
	toSerialize["unix_user"] = o.UnixUser
	return toSerialize, nil
}

type NullableVm struct {
	value *Vm
	isSet bool
}

func (v NullableVm) Get() *Vm {
	return v.value
}

func (v *NullableVm) Set(val *Vm) {
	v.value = val
	v.isSet = true
}

func (v NullableVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVm(val *Vm) *NullableVm {
	return &NullableVm{value: val, isSet: true}
}

func (v NullableVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



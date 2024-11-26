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

// checks if the RestorePostgresDatabaseWithIDRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestorePostgresDatabaseWithIDRequest{}

// RestorePostgresDatabaseWithIDRequest struct for RestorePostgresDatabaseWithIDRequest
type RestorePostgresDatabaseWithIDRequest struct {
	Name string `json:"name"`
	RestoreTarget string `json:"restore_target"`
}

// NewRestorePostgresDatabaseWithIDRequest instantiates a new RestorePostgresDatabaseWithIDRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestorePostgresDatabaseWithIDRequest(name string, restoreTarget string) *RestorePostgresDatabaseWithIDRequest {
	this := RestorePostgresDatabaseWithIDRequest{}
	this.Name = name
	this.RestoreTarget = restoreTarget
	return &this
}

// NewRestorePostgresDatabaseWithIDRequestWithDefaults instantiates a new RestorePostgresDatabaseWithIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestorePostgresDatabaseWithIDRequestWithDefaults() *RestorePostgresDatabaseWithIDRequest {
	this := RestorePostgresDatabaseWithIDRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RestorePostgresDatabaseWithIDRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RestorePostgresDatabaseWithIDRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RestorePostgresDatabaseWithIDRequest) SetName(v string) {
	o.Name = v
}

// GetRestoreTarget returns the RestoreTarget field value
func (o *RestorePostgresDatabaseWithIDRequest) GetRestoreTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestoreTarget
}

// GetRestoreTargetOk returns a tuple with the RestoreTarget field value
// and a boolean to check if the value has been set.
func (o *RestorePostgresDatabaseWithIDRequest) GetRestoreTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestoreTarget, true
}

// SetRestoreTarget sets field value
func (o *RestorePostgresDatabaseWithIDRequest) SetRestoreTarget(v string) {
	o.RestoreTarget = v
}

func (o RestorePostgresDatabaseWithIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestorePostgresDatabaseWithIDRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["restore_target"] = o.RestoreTarget
	return toSerialize, nil
}

type NullableRestorePostgresDatabaseWithIDRequest struct {
	value *RestorePostgresDatabaseWithIDRequest
	isSet bool
}

func (v NullableRestorePostgresDatabaseWithIDRequest) Get() *RestorePostgresDatabaseWithIDRequest {
	return v.value
}

func (v *NullableRestorePostgresDatabaseWithIDRequest) Set(val *RestorePostgresDatabaseWithIDRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestorePostgresDatabaseWithIDRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestorePostgresDatabaseWithIDRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestorePostgresDatabaseWithIDRequest(val *RestorePostgresDatabaseWithIDRequest) *NullableRestorePostgresDatabaseWithIDRequest {
	return &NullableRestorePostgresDatabaseWithIDRequest{value: val, isSet: true}
}

func (v NullableRestorePostgresDatabaseWithIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestorePostgresDatabaseWithIDRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



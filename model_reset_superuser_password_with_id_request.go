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

// checks if the ResetSuperuserPasswordWithIDRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetSuperuserPasswordWithIDRequest{}

// ResetSuperuserPasswordWithIDRequest struct for ResetSuperuserPasswordWithIDRequest
type ResetSuperuserPasswordWithIDRequest struct {
	Password string `json:"password"`
}

// NewResetSuperuserPasswordWithIDRequest instantiates a new ResetSuperuserPasswordWithIDRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetSuperuserPasswordWithIDRequest(password string) *ResetSuperuserPasswordWithIDRequest {
	this := ResetSuperuserPasswordWithIDRequest{}
	this.Password = password
	return &this
}

// NewResetSuperuserPasswordWithIDRequestWithDefaults instantiates a new ResetSuperuserPasswordWithIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetSuperuserPasswordWithIDRequestWithDefaults() *ResetSuperuserPasswordWithIDRequest {
	this := ResetSuperuserPasswordWithIDRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *ResetSuperuserPasswordWithIDRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ResetSuperuserPasswordWithIDRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ResetSuperuserPasswordWithIDRequest) SetPassword(v string) {
	o.Password = v
}

func (o ResetSuperuserPasswordWithIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetSuperuserPasswordWithIDRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableResetSuperuserPasswordWithIDRequest struct {
	value *ResetSuperuserPasswordWithIDRequest
	isSet bool
}

func (v NullableResetSuperuserPasswordWithIDRequest) Get() *ResetSuperuserPasswordWithIDRequest {
	return v.value
}

func (v *NullableResetSuperuserPasswordWithIDRequest) Set(val *ResetSuperuserPasswordWithIDRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResetSuperuserPasswordWithIDRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResetSuperuserPasswordWithIDRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetSuperuserPasswordWithIDRequest(val *ResetSuperuserPasswordWithIDRequest) *NullableResetSuperuserPasswordWithIDRequest {
	return &NullableResetSuperuserPasswordWithIDRequest{value: val, isSet: true}
}

func (v NullableResetSuperuserPasswordWithIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetSuperuserPasswordWithIDRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



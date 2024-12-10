# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the VM | 
**Ip4** | **NullableString** | IPv4 address | 
**Ip4Enabled** | **bool** | Whether IPv4 is enabled | 
**Ip6** | **NullableString** | IPv6 address | 
**Location** | **string** | Location of the VM | 
**Name** | **string** | Name of the VM | 
**Size** | **string** | Size of the underlying VM | 
**State** | **string** | State of the VM | 
**StorageSizeGib** | **int32** | Storage size in GiB | 
**UnixUser** | **string** | Unix user of the VM | 

## Methods

### NewVm

`func NewVm(id string, ip4 NullableString, ip4Enabled bool, ip6 NullableString, location string, name string, size string, state string, storageSizeGib int32, unixUser string, ) *Vm`

NewVm instantiates a new Vm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmWithDefaults

`func NewVmWithDefaults() *Vm`

NewVmWithDefaults instantiates a new Vm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vm) SetId(v string)`

SetId sets Id field to given value.


### GetIp4

`func (o *Vm) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *Vm) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *Vm) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.


### SetIp4Nil

`func (o *Vm) SetIp4Nil(b bool)`

 SetIp4Nil sets the value for Ip4 to be an explicit nil

### UnsetIp4
`func (o *Vm) UnsetIp4()`

UnsetIp4 ensures that no value is present for Ip4, not even an explicit nil
### GetIp4Enabled

`func (o *Vm) GetIp4Enabled() bool`

GetIp4Enabled returns the Ip4Enabled field if non-nil, zero value otherwise.

### GetIp4EnabledOk

`func (o *Vm) GetIp4EnabledOk() (*bool, bool)`

GetIp4EnabledOk returns a tuple with the Ip4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Enabled

`func (o *Vm) SetIp4Enabled(v bool)`

SetIp4Enabled sets Ip4Enabled field to given value.


### GetIp6

`func (o *Vm) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *Vm) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *Vm) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.


### SetIp6Nil

`func (o *Vm) SetIp6Nil(b bool)`

 SetIp6Nil sets the value for Ip6 to be an explicit nil

### UnsetIp6
`func (o *Vm) UnsetIp6()`

UnsetIp6 ensures that no value is present for Ip6, not even an explicit nil
### GetLocation

`func (o *Vm) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Vm) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Vm) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *Vm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vm) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *Vm) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Vm) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Vm) SetSize(v string)`

SetSize sets Size field to given value.


### GetState

`func (o *Vm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Vm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Vm) SetState(v string)`

SetState sets State field to given value.


### GetStorageSizeGib

`func (o *Vm) GetStorageSizeGib() int32`

GetStorageSizeGib returns the StorageSizeGib field if non-nil, zero value otherwise.

### GetStorageSizeGibOk

`func (o *Vm) GetStorageSizeGibOk() (*int32, bool)`

GetStorageSizeGibOk returns a tuple with the StorageSizeGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeGib

`func (o *Vm) SetStorageSizeGib(v int32)`

SetStorageSizeGib sets StorageSizeGib field to given value.


### GetUnixUser

`func (o *Vm) GetUnixUser() string`

GetUnixUser returns the UnixUser field if non-nil, zero value otherwise.

### GetUnixUserOk

`func (o *Vm) GetUnixUserOk() (*string, bool)`

GetUnixUserOk returns a tuple with the UnixUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUser

`func (o *Vm) SetUnixUser(v string)`

SetUnixUser sets UnixUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



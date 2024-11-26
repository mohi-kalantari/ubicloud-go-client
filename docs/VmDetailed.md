# VmDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the VM | 
**Ip4** | **NullableString** | IPv4 address | 
**Ip6** | **NullableString** | IPv6 address | 
**Location** | **string** | Location of the VM | 
**Name** | **string** | Name of the VM | 
**Size** | **string** | Size of the underlying VM | 
**State** | **string** | State of the VM | 
**StorageSizeGib** | **int32** | Storage size in GiB | 
**UnixUser** | **string** | Unix user of the VM | 
**Firewalls** | [**[]Firewall**](Firewall.md) | List of firewalls | 
**PrivateIpv4** | **string** | Private IPv4 address | 
**PrivateIpv6** | **string** | Private IPv6 address | 
**Subnet** | **string** | Subnet of the VM | 

## Methods

### NewVmDetailed

`func NewVmDetailed(id string, ip4 NullableString, ip6 NullableString, location string, name string, size string, state string, storageSizeGib int32, unixUser string, firewalls []Firewall, privateIpv4 string, privateIpv6 string, subnet string, ) *VmDetailed`

NewVmDetailed instantiates a new VmDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDetailedWithDefaults

`func NewVmDetailedWithDefaults() *VmDetailed`

NewVmDetailedWithDefaults instantiates a new VmDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmDetailed) SetId(v string)`

SetId sets Id field to given value.


### GetIp4

`func (o *VmDetailed) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *VmDetailed) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *VmDetailed) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.


### SetIp4Nil

`func (o *VmDetailed) SetIp4Nil(b bool)`

 SetIp4Nil sets the value for Ip4 to be an explicit nil

### UnsetIp4
`func (o *VmDetailed) UnsetIp4()`

UnsetIp4 ensures that no value is present for Ip4, not even an explicit nil
### GetIp6

`func (o *VmDetailed) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *VmDetailed) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *VmDetailed) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.


### SetIp6Nil

`func (o *VmDetailed) SetIp6Nil(b bool)`

 SetIp6Nil sets the value for Ip6 to be an explicit nil

### UnsetIp6
`func (o *VmDetailed) UnsetIp6()`

UnsetIp6 ensures that no value is present for Ip6, not even an explicit nil
### GetLocation

`func (o *VmDetailed) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VmDetailed) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VmDetailed) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *VmDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *VmDetailed) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VmDetailed) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VmDetailed) SetSize(v string)`

SetSize sets Size field to given value.


### GetState

`func (o *VmDetailed) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VmDetailed) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VmDetailed) SetState(v string)`

SetState sets State field to given value.


### GetStorageSizeGib

`func (o *VmDetailed) GetStorageSizeGib() int32`

GetStorageSizeGib returns the StorageSizeGib field if non-nil, zero value otherwise.

### GetStorageSizeGibOk

`func (o *VmDetailed) GetStorageSizeGibOk() (*int32, bool)`

GetStorageSizeGibOk returns a tuple with the StorageSizeGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeGib

`func (o *VmDetailed) SetStorageSizeGib(v int32)`

SetStorageSizeGib sets StorageSizeGib field to given value.


### GetUnixUser

`func (o *VmDetailed) GetUnixUser() string`

GetUnixUser returns the UnixUser field if non-nil, zero value otherwise.

### GetUnixUserOk

`func (o *VmDetailed) GetUnixUserOk() (*string, bool)`

GetUnixUserOk returns a tuple with the UnixUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUser

`func (o *VmDetailed) SetUnixUser(v string)`

SetUnixUser sets UnixUser field to given value.


### GetFirewalls

`func (o *VmDetailed) GetFirewalls() []Firewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *VmDetailed) GetFirewallsOk() (*[]Firewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *VmDetailed) SetFirewalls(v []Firewall)`

SetFirewalls sets Firewalls field to given value.


### GetPrivateIpv4

`func (o *VmDetailed) GetPrivateIpv4() string`

GetPrivateIpv4 returns the PrivateIpv4 field if non-nil, zero value otherwise.

### GetPrivateIpv4Ok

`func (o *VmDetailed) GetPrivateIpv4Ok() (*string, bool)`

GetPrivateIpv4Ok returns a tuple with the PrivateIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4

`func (o *VmDetailed) SetPrivateIpv4(v string)`

SetPrivateIpv4 sets PrivateIpv4 field to given value.


### GetPrivateIpv6

`func (o *VmDetailed) GetPrivateIpv6() string`

GetPrivateIpv6 returns the PrivateIpv6 field if non-nil, zero value otherwise.

### GetPrivateIpv6Ok

`func (o *VmDetailed) GetPrivateIpv6Ok() (*string, bool)`

GetPrivateIpv6Ok returns a tuple with the PrivateIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv6

`func (o *VmDetailed) SetPrivateIpv6(v string)`

SetPrivateIpv6 sets PrivateIpv6 field to given value.


### GetSubnet

`func (o *VmDetailed) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VmDetailed) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VmDetailed) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VmDetailedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firewalls** | [**[]Firewall**](Firewall.md) | List of firewalls | 
**PrivateIpv4** | **string** | Private IPv4 address | 
**PrivateIpv6** | **string** | Private IPv6 address | 
**Subnet** | **string** | Subnet of the VM | 
**UnixUser** | **string** | Unix user of the VM | 

## Methods

### NewVmDetailedAllOf

`func NewVmDetailedAllOf(firewalls []Firewall, privateIpv4 string, privateIpv6 string, subnet string, unixUser string, ) *VmDetailedAllOf`

NewVmDetailedAllOf instantiates a new VmDetailedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDetailedAllOfWithDefaults

`func NewVmDetailedAllOfWithDefaults() *VmDetailedAllOf`

NewVmDetailedAllOfWithDefaults instantiates a new VmDetailedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewalls

`func (o *VmDetailedAllOf) GetFirewalls() []Firewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *VmDetailedAllOf) GetFirewallsOk() (*[]Firewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *VmDetailedAllOf) SetFirewalls(v []Firewall)`

SetFirewalls sets Firewalls field to given value.


### GetPrivateIpv4

`func (o *VmDetailedAllOf) GetPrivateIpv4() string`

GetPrivateIpv4 returns the PrivateIpv4 field if non-nil, zero value otherwise.

### GetPrivateIpv4Ok

`func (o *VmDetailedAllOf) GetPrivateIpv4Ok() (*string, bool)`

GetPrivateIpv4Ok returns a tuple with the PrivateIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4

`func (o *VmDetailedAllOf) SetPrivateIpv4(v string)`

SetPrivateIpv4 sets PrivateIpv4 field to given value.


### GetPrivateIpv6

`func (o *VmDetailedAllOf) GetPrivateIpv6() string`

GetPrivateIpv6 returns the PrivateIpv6 field if non-nil, zero value otherwise.

### GetPrivateIpv6Ok

`func (o *VmDetailedAllOf) GetPrivateIpv6Ok() (*string, bool)`

GetPrivateIpv6Ok returns a tuple with the PrivateIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv6

`func (o *VmDetailedAllOf) SetPrivateIpv6(v string)`

SetPrivateIpv6 sets PrivateIpv6 field to given value.


### GetSubnet

`func (o *VmDetailedAllOf) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VmDetailedAllOf) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VmDetailedAllOf) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetUnixUser

`func (o *VmDetailedAllOf) GetUnixUser() string`

GetUnixUser returns the UnixUser field if non-nil, zero value otherwise.

### GetUnixUserOk

`func (o *VmDetailedAllOf) GetUnixUserOk() (*string, bool)`

GetUnixUserOk returns a tuple with the UnixUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUser

`func (o *VmDetailedAllOf) SetUnixUser(v string)`

SetUnixUser sets UnixUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



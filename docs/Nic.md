# Nic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the NIC | 
**Name** | **string** | Name of the NIC | 
**PrivateIpv4** | **string** | Private IPv4 address | 
**PrivateIpv6** | **string** | Private IPv6 address | 
**VmName** | **NullableString** | Name of the VM | 

## Methods

### NewNic

`func NewNic(id string, name string, privateIpv4 string, privateIpv6 string, vmName NullableString, ) *Nic`

NewNic instantiates a new Nic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicWithDefaults

`func NewNicWithDefaults() *Nic`

NewNicWithDefaults instantiates a new Nic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Nic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Nic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Nic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Nic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Nic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Nic) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateIpv4

`func (o *Nic) GetPrivateIpv4() string`

GetPrivateIpv4 returns the PrivateIpv4 field if non-nil, zero value otherwise.

### GetPrivateIpv4Ok

`func (o *Nic) GetPrivateIpv4Ok() (*string, bool)`

GetPrivateIpv4Ok returns a tuple with the PrivateIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4

`func (o *Nic) SetPrivateIpv4(v string)`

SetPrivateIpv4 sets PrivateIpv4 field to given value.


### GetPrivateIpv6

`func (o *Nic) GetPrivateIpv6() string`

GetPrivateIpv6 returns the PrivateIpv6 field if non-nil, zero value otherwise.

### GetPrivateIpv6Ok

`func (o *Nic) GetPrivateIpv6Ok() (*string, bool)`

GetPrivateIpv6Ok returns a tuple with the PrivateIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv6

`func (o *Nic) SetPrivateIpv6(v string)`

SetPrivateIpv6 sets PrivateIpv6 field to given value.


### GetVmName

`func (o *Nic) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *Nic) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *Nic) SetVmName(v string)`

SetVmName sets VmName field to given value.


### SetVmNameNil

`func (o *Nic) SetVmNameNil(b bool)`

 SetVmNameNil sets the value for VmName to be an explicit nil

### UnsetVmName
`func (o *Nic) UnsetVmName()`

UnsetVmName ensures that no value is present for VmName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootImage** | Pointer to **string** | Boot image of the VM | [optional] 
**EnableIp4** | Pointer to **bool** | Enable IPv4 | [optional] 
**PrivateSubnetId** | Pointer to **string** | ID of the private subnet | [optional] 
**PublicKey** | **string** | Public SSH key for the VM | 
**Size** | Pointer to **string** | Size of the VM | [optional] 
**StorageSize** | Pointer to **int32** | Requested storage size in GiB | [optional] 
**UnixUser** | Pointer to **string** | Unix user of the VM | [optional] 

## Methods

### NewCreateVMRequest

`func NewCreateVMRequest(publicKey string, ) *CreateVMRequest`

NewCreateVMRequest instantiates a new CreateVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestWithDefaults

`func NewCreateVMRequestWithDefaults() *CreateVMRequest`

NewCreateVMRequestWithDefaults instantiates a new CreateVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootImage

`func (o *CreateVMRequest) GetBootImage() string`

GetBootImage returns the BootImage field if non-nil, zero value otherwise.

### GetBootImageOk

`func (o *CreateVMRequest) GetBootImageOk() (*string, bool)`

GetBootImageOk returns a tuple with the BootImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootImage

`func (o *CreateVMRequest) SetBootImage(v string)`

SetBootImage sets BootImage field to given value.

### HasBootImage

`func (o *CreateVMRequest) HasBootImage() bool`

HasBootImage returns a boolean if a field has been set.

### GetEnableIp4

`func (o *CreateVMRequest) GetEnableIp4() bool`

GetEnableIp4 returns the EnableIp4 field if non-nil, zero value otherwise.

### GetEnableIp4Ok

`func (o *CreateVMRequest) GetEnableIp4Ok() (*bool, bool)`

GetEnableIp4Ok returns a tuple with the EnableIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIp4

`func (o *CreateVMRequest) SetEnableIp4(v bool)`

SetEnableIp4 sets EnableIp4 field to given value.

### HasEnableIp4

`func (o *CreateVMRequest) HasEnableIp4() bool`

HasEnableIp4 returns a boolean if a field has been set.

### GetPrivateSubnetId

`func (o *CreateVMRequest) GetPrivateSubnetId() string`

GetPrivateSubnetId returns the PrivateSubnetId field if non-nil, zero value otherwise.

### GetPrivateSubnetIdOk

`func (o *CreateVMRequest) GetPrivateSubnetIdOk() (*string, bool)`

GetPrivateSubnetIdOk returns a tuple with the PrivateSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnetId

`func (o *CreateVMRequest) SetPrivateSubnetId(v string)`

SetPrivateSubnetId sets PrivateSubnetId field to given value.

### HasPrivateSubnetId

`func (o *CreateVMRequest) HasPrivateSubnetId() bool`

HasPrivateSubnetId returns a boolean if a field has been set.

### GetPublicKey

`func (o *CreateVMRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateVMRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreateVMRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSize

`func (o *CreateVMRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageSize

`func (o *CreateVMRequest) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CreateVMRequest) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CreateVMRequest) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *CreateVMRequest) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetUnixUser

`func (o *CreateVMRequest) GetUnixUser() string`

GetUnixUser returns the UnixUser field if non-nil, zero value otherwise.

### GetUnixUserOk

`func (o *CreateVMRequest) GetUnixUserOk() (*string, bool)`

GetUnixUserOk returns a tuple with the UnixUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUser

`func (o *CreateVMRequest) SetUnixUser(v string)`

SetUnixUser sets UnixUser field to given value.

### HasUnixUser

`func (o *CreateVMRequest) HasUnixUser() bool`

HasUnixUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



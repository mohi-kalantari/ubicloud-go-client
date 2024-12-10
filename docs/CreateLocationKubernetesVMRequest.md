# CreateLocationKubernetesVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootImage** | **string** | Boot image of the VM | 
**EnableIp4** | **bool** | Enable IPv4 | 
**PrivateSubnetId** | **string** | ID of the private subnet | 
**Size** | **string** | Size of the VM | 
**StorageSize** | **int32** | Requested storage size in GiB | 
**UnixUser** | **string** | Unix user of the VM | 
**Commands** | **[]string** | List of commands for the Kubernetes VM | 

## Methods

### NewCreateLocationKubernetesVMRequest

`func NewCreateLocationKubernetesVMRequest(bootImage string, enableIp4 bool, privateSubnetId string, size string, storageSize int32, unixUser string, commands []string, ) *CreateLocationKubernetesVMRequest`

NewCreateLocationKubernetesVMRequest instantiates a new CreateLocationKubernetesVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLocationKubernetesVMRequestWithDefaults

`func NewCreateLocationKubernetesVMRequestWithDefaults() *CreateLocationKubernetesVMRequest`

NewCreateLocationKubernetesVMRequestWithDefaults instantiates a new CreateLocationKubernetesVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootImage

`func (o *CreateLocationKubernetesVMRequest) GetBootImage() string`

GetBootImage returns the BootImage field if non-nil, zero value otherwise.

### GetBootImageOk

`func (o *CreateLocationKubernetesVMRequest) GetBootImageOk() (*string, bool)`

GetBootImageOk returns a tuple with the BootImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootImage

`func (o *CreateLocationKubernetesVMRequest) SetBootImage(v string)`

SetBootImage sets BootImage field to given value.


### GetEnableIp4

`func (o *CreateLocationKubernetesVMRequest) GetEnableIp4() bool`

GetEnableIp4 returns the EnableIp4 field if non-nil, zero value otherwise.

### GetEnableIp4Ok

`func (o *CreateLocationKubernetesVMRequest) GetEnableIp4Ok() (*bool, bool)`

GetEnableIp4Ok returns a tuple with the EnableIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIp4

`func (o *CreateLocationKubernetesVMRequest) SetEnableIp4(v bool)`

SetEnableIp4 sets EnableIp4 field to given value.


### GetPrivateSubnetId

`func (o *CreateLocationKubernetesVMRequest) GetPrivateSubnetId() string`

GetPrivateSubnetId returns the PrivateSubnetId field if non-nil, zero value otherwise.

### GetPrivateSubnetIdOk

`func (o *CreateLocationKubernetesVMRequest) GetPrivateSubnetIdOk() (*string, bool)`

GetPrivateSubnetIdOk returns a tuple with the PrivateSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnetId

`func (o *CreateLocationKubernetesVMRequest) SetPrivateSubnetId(v string)`

SetPrivateSubnetId sets PrivateSubnetId field to given value.


### GetSize

`func (o *CreateLocationKubernetesVMRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateLocationKubernetesVMRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateLocationKubernetesVMRequest) SetSize(v string)`

SetSize sets Size field to given value.


### GetStorageSize

`func (o *CreateLocationKubernetesVMRequest) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CreateLocationKubernetesVMRequest) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CreateLocationKubernetesVMRequest) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.


### GetUnixUser

`func (o *CreateLocationKubernetesVMRequest) GetUnixUser() string`

GetUnixUser returns the UnixUser field if non-nil, zero value otherwise.

### GetUnixUserOk

`func (o *CreateLocationKubernetesVMRequest) GetUnixUserOk() (*string, bool)`

GetUnixUserOk returns a tuple with the UnixUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUser

`func (o *CreateLocationKubernetesVMRequest) SetUnixUser(v string)`

SetUnixUser sets UnixUser field to given value.


### GetCommands

`func (o *CreateLocationKubernetesVMRequest) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *CreateLocationKubernetesVMRequest) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *CreateLocationKubernetesVMRequest) SetCommands(v []string)`

SetCommands sets Commands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreatePostgresDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | Pointer to **string** | Kind of database | [optional] 
**HaType** | Pointer to **string** | High availability type | [optional] 
**Size** | **string** | Requested size for the underlying VM | 
**StorageSize** | Pointer to **int32** | Requested storage size in GiB | [optional] 

## Methods

### NewCreatePostgresDatabaseRequest

`func NewCreatePostgresDatabaseRequest(size string, ) *CreatePostgresDatabaseRequest`

NewCreatePostgresDatabaseRequest instantiates a new CreatePostgresDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePostgresDatabaseRequestWithDefaults

`func NewCreatePostgresDatabaseRequestWithDefaults() *CreatePostgresDatabaseRequest`

NewCreatePostgresDatabaseRequestWithDefaults instantiates a new CreatePostgresDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *CreatePostgresDatabaseRequest) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *CreatePostgresDatabaseRequest) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *CreatePostgresDatabaseRequest) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *CreatePostgresDatabaseRequest) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetHaType

`func (o *CreatePostgresDatabaseRequest) GetHaType() string`

GetHaType returns the HaType field if non-nil, zero value otherwise.

### GetHaTypeOk

`func (o *CreatePostgresDatabaseRequest) GetHaTypeOk() (*string, bool)`

GetHaTypeOk returns a tuple with the HaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaType

`func (o *CreatePostgresDatabaseRequest) SetHaType(v string)`

SetHaType sets HaType field to given value.

### HasHaType

`func (o *CreatePostgresDatabaseRequest) HasHaType() bool`

HasHaType returns a boolean if a field has been set.

### GetSize

`func (o *CreatePostgresDatabaseRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreatePostgresDatabaseRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreatePostgresDatabaseRequest) SetSize(v string)`

SetSize sets Size field to given value.


### GetStorageSize

`func (o *CreatePostgresDatabaseRequest) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CreatePostgresDatabaseRequest) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CreatePostgresDatabaseRequest) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *CreatePostgresDatabaseRequest) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



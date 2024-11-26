# Postgres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | **string** | Kind of Postgres database | 
**HaType** | **string** | High availability type | 
**Id** | **string** | ID of the Postgres database | 
**Location** | **string** | Location of the Postgres database | 
**Name** | **string** | Name of the Postgres database | 
**State** | **string** | State of the Postgres database | 
**StorageSizeGib** | **int32** | Storage size in GiB | 
**Version** | **string** | Postgres version | 
**VmSize** | **string** | Size of the underlying VM | 

## Methods

### NewPostgres

`func NewPostgres(flavor string, haType string, id string, location string, name string, state string, storageSizeGib int32, version string, vmSize string, ) *Postgres`

NewPostgres instantiates a new Postgres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresWithDefaults

`func NewPostgresWithDefaults() *Postgres`

NewPostgresWithDefaults instantiates a new Postgres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *Postgres) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Postgres) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Postgres) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetHaType

`func (o *Postgres) GetHaType() string`

GetHaType returns the HaType field if non-nil, zero value otherwise.

### GetHaTypeOk

`func (o *Postgres) GetHaTypeOk() (*string, bool)`

GetHaTypeOk returns a tuple with the HaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaType

`func (o *Postgres) SetHaType(v string)`

SetHaType sets HaType field to given value.


### GetId

`func (o *Postgres) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Postgres) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Postgres) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *Postgres) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Postgres) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Postgres) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *Postgres) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Postgres) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Postgres) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *Postgres) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Postgres) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Postgres) SetState(v string)`

SetState sets State field to given value.


### GetStorageSizeGib

`func (o *Postgres) GetStorageSizeGib() int32`

GetStorageSizeGib returns the StorageSizeGib field if non-nil, zero value otherwise.

### GetStorageSizeGibOk

`func (o *Postgres) GetStorageSizeGibOk() (*int32, bool)`

GetStorageSizeGibOk returns a tuple with the StorageSizeGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeGib

`func (o *Postgres) SetStorageSizeGib(v int32)`

SetStorageSizeGib sets StorageSizeGib field to given value.


### GetVersion

`func (o *Postgres) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Postgres) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Postgres) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVmSize

`func (o *Postgres) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *Postgres) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *Postgres) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



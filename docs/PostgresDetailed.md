# PostgresDetailed

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
**ConnectionString** | **NullableString** | Connection string to the Postgres database | 
**EarliestRestoreTime** | **NullableString** | Earliest restore time (if primary) | 
**FirewallRules** | [**[]PostgresFirewallRule**](PostgresFirewallRule.md) | List of Postgres firewall rules | 
**LatestRestoreTime** | **string** | Latest restore time (if primary)\&quot; | 
**Primary** | **bool** | Is the database primary | 

## Methods

### NewPostgresDetailed

`func NewPostgresDetailed(flavor string, haType string, id string, location string, name string, state string, storageSizeGib int32, version string, vmSize string, connectionString NullableString, earliestRestoreTime NullableString, firewallRules []PostgresFirewallRule, latestRestoreTime string, primary bool, ) *PostgresDetailed`

NewPostgresDetailed instantiates a new PostgresDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresDetailedWithDefaults

`func NewPostgresDetailedWithDefaults() *PostgresDetailed`

NewPostgresDetailedWithDefaults instantiates a new PostgresDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *PostgresDetailed) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *PostgresDetailed) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *PostgresDetailed) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetHaType

`func (o *PostgresDetailed) GetHaType() string`

GetHaType returns the HaType field if non-nil, zero value otherwise.

### GetHaTypeOk

`func (o *PostgresDetailed) GetHaTypeOk() (*string, bool)`

GetHaTypeOk returns a tuple with the HaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaType

`func (o *PostgresDetailed) SetHaType(v string)`

SetHaType sets HaType field to given value.


### GetId

`func (o *PostgresDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresDetailed) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *PostgresDetailed) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostgresDetailed) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostgresDetailed) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *PostgresDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostgresDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostgresDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *PostgresDetailed) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostgresDetailed) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostgresDetailed) SetState(v string)`

SetState sets State field to given value.


### GetStorageSizeGib

`func (o *PostgresDetailed) GetStorageSizeGib() int32`

GetStorageSizeGib returns the StorageSizeGib field if non-nil, zero value otherwise.

### GetStorageSizeGibOk

`func (o *PostgresDetailed) GetStorageSizeGibOk() (*int32, bool)`

GetStorageSizeGibOk returns a tuple with the StorageSizeGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeGib

`func (o *PostgresDetailed) SetStorageSizeGib(v int32)`

SetStorageSizeGib sets StorageSizeGib field to given value.


### GetVersion

`func (o *PostgresDetailed) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostgresDetailed) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostgresDetailed) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVmSize

`func (o *PostgresDetailed) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *PostgresDetailed) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *PostgresDetailed) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.


### GetConnectionString

`func (o *PostgresDetailed) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *PostgresDetailed) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *PostgresDetailed) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### SetConnectionStringNil

`func (o *PostgresDetailed) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *PostgresDetailed) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil
### GetEarliestRestoreTime

`func (o *PostgresDetailed) GetEarliestRestoreTime() string`

GetEarliestRestoreTime returns the EarliestRestoreTime field if non-nil, zero value otherwise.

### GetEarliestRestoreTimeOk

`func (o *PostgresDetailed) GetEarliestRestoreTimeOk() (*string, bool)`

GetEarliestRestoreTimeOk returns a tuple with the EarliestRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRestoreTime

`func (o *PostgresDetailed) SetEarliestRestoreTime(v string)`

SetEarliestRestoreTime sets EarliestRestoreTime field to given value.


### SetEarliestRestoreTimeNil

`func (o *PostgresDetailed) SetEarliestRestoreTimeNil(b bool)`

 SetEarliestRestoreTimeNil sets the value for EarliestRestoreTime to be an explicit nil

### UnsetEarliestRestoreTime
`func (o *PostgresDetailed) UnsetEarliestRestoreTime()`

UnsetEarliestRestoreTime ensures that no value is present for EarliestRestoreTime, not even an explicit nil
### GetFirewallRules

`func (o *PostgresDetailed) GetFirewallRules() []PostgresFirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PostgresDetailed) GetFirewallRulesOk() (*[]PostgresFirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PostgresDetailed) SetFirewallRules(v []PostgresFirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetLatestRestoreTime

`func (o *PostgresDetailed) GetLatestRestoreTime() string`

GetLatestRestoreTime returns the LatestRestoreTime field if non-nil, zero value otherwise.

### GetLatestRestoreTimeOk

`func (o *PostgresDetailed) GetLatestRestoreTimeOk() (*string, bool)`

GetLatestRestoreTimeOk returns a tuple with the LatestRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRestoreTime

`func (o *PostgresDetailed) SetLatestRestoreTime(v string)`

SetLatestRestoreTime sets LatestRestoreTime field to given value.


### GetPrimary

`func (o *PostgresDetailed) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PostgresDetailed) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PostgresDetailed) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostgresDetailedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | **NullableString** | Connection string to the Postgres database | 
**EarliestRestoreTime** | **NullableString** | Earliest restore time (if primary) | 
**FirewallRules** | [**[]PostgresFirewallRule**](PostgresFirewallRule.md) | List of Postgres firewall rules | 
**LatestRestoreTime** | **string** | Latest restore time (if primary)\&quot; | 
**Primary** | **bool** | Is the database primary | 

## Methods

### NewPostgresDetailedAllOf

`func NewPostgresDetailedAllOf(connectionString NullableString, earliestRestoreTime NullableString, firewallRules []PostgresFirewallRule, latestRestoreTime string, primary bool, ) *PostgresDetailedAllOf`

NewPostgresDetailedAllOf instantiates a new PostgresDetailedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresDetailedAllOfWithDefaults

`func NewPostgresDetailedAllOfWithDefaults() *PostgresDetailedAllOf`

NewPostgresDetailedAllOfWithDefaults instantiates a new PostgresDetailedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *PostgresDetailedAllOf) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *PostgresDetailedAllOf) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *PostgresDetailedAllOf) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### SetConnectionStringNil

`func (o *PostgresDetailedAllOf) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *PostgresDetailedAllOf) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil
### GetEarliestRestoreTime

`func (o *PostgresDetailedAllOf) GetEarliestRestoreTime() string`

GetEarliestRestoreTime returns the EarliestRestoreTime field if non-nil, zero value otherwise.

### GetEarliestRestoreTimeOk

`func (o *PostgresDetailedAllOf) GetEarliestRestoreTimeOk() (*string, bool)`

GetEarliestRestoreTimeOk returns a tuple with the EarliestRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRestoreTime

`func (o *PostgresDetailedAllOf) SetEarliestRestoreTime(v string)`

SetEarliestRestoreTime sets EarliestRestoreTime field to given value.


### SetEarliestRestoreTimeNil

`func (o *PostgresDetailedAllOf) SetEarliestRestoreTimeNil(b bool)`

 SetEarliestRestoreTimeNil sets the value for EarliestRestoreTime to be an explicit nil

### UnsetEarliestRestoreTime
`func (o *PostgresDetailedAllOf) UnsetEarliestRestoreTime()`

UnsetEarliestRestoreTime ensures that no value is present for EarliestRestoreTime, not even an explicit nil
### GetFirewallRules

`func (o *PostgresDetailedAllOf) GetFirewallRules() []PostgresFirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PostgresDetailedAllOf) GetFirewallRulesOk() (*[]PostgresFirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PostgresDetailedAllOf) SetFirewallRules(v []PostgresFirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetLatestRestoreTime

`func (o *PostgresDetailedAllOf) GetLatestRestoreTime() string`

GetLatestRestoreTime returns the LatestRestoreTime field if non-nil, zero value otherwise.

### GetLatestRestoreTimeOk

`func (o *PostgresDetailedAllOf) GetLatestRestoreTimeOk() (*string, bool)`

GetLatestRestoreTimeOk returns a tuple with the LatestRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRestoreTime

`func (o *PostgresDetailedAllOf) SetLatestRestoreTime(v string)`

SetLatestRestoreTime sets LatestRestoreTime field to given value.


### GetPrimary

`func (o *PostgresDetailedAllOf) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PostgresDetailedAllOf) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PostgresDetailedAllOf) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostgresFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR of the Postgres firewall rule | 
**Id** | **string** | ID of the Postgres firewall rule | 

## Methods

### NewPostgresFirewallRule

`func NewPostgresFirewallRule(cidr string, id string, ) *PostgresFirewallRule`

NewPostgresFirewallRule instantiates a new PostgresFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresFirewallRuleWithDefaults

`func NewPostgresFirewallRuleWithDefaults() *PostgresFirewallRule`

NewPostgresFirewallRuleWithDefaults instantiates a new PostgresFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *PostgresFirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PostgresFirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PostgresFirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetId

`func (o *PostgresFirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresFirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresFirewallRule) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListLocationPostgresFirewallRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Items** | [**[]PostgresFirewallRule**](PostgresFirewallRule.md) |  | 

## Methods

### NewListLocationPostgresFirewallRules200Response

`func NewListLocationPostgresFirewallRules200Response(count int32, items []PostgresFirewallRule, ) *ListLocationPostgresFirewallRules200Response`

NewListLocationPostgresFirewallRules200Response instantiates a new ListLocationPostgresFirewallRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLocationPostgresFirewallRules200ResponseWithDefaults

`func NewListLocationPostgresFirewallRules200ResponseWithDefaults() *ListLocationPostgresFirewallRules200Response`

NewListLocationPostgresFirewallRules200ResponseWithDefaults instantiates a new ListLocationPostgresFirewallRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListLocationPostgresFirewallRules200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListLocationPostgresFirewallRules200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListLocationPostgresFirewallRules200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *ListLocationPostgresFirewallRules200Response) GetItems() []PostgresFirewallRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListLocationPostgresFirewallRules200Response) GetItemsOk() (*[]PostgresFirewallRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListLocationPostgresFirewallRules200Response) SetItems(v []PostgresFirewallRule)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



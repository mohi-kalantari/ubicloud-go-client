# FirewallDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the firewall | 
**FirewallRules** | [**[]FirewallRule**](FirewallRule.md) | List of firewall rules | 
**Id** | **string** | ID of the firewall | 
**Location** | **string** | Location of the the firewall | 
**Name** | **string** | Name of the firewall | 
**PrivateSubnets** | **[]string** | List of private subnets | 

## Methods

### NewFirewallDetailed

`func NewFirewallDetailed(description string, firewallRules []FirewallRule, id string, location string, name string, privateSubnets []string, ) *FirewallDetailed`

NewFirewallDetailed instantiates a new FirewallDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallDetailedWithDefaults

`func NewFirewallDetailedWithDefaults() *FirewallDetailed`

NewFirewallDetailedWithDefaults instantiates a new FirewallDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FirewallDetailed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallDetailed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallDetailed) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFirewallRules

`func (o *FirewallDetailed) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *FirewallDetailed) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *FirewallDetailed) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetId

`func (o *FirewallDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallDetailed) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *FirewallDetailed) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FirewallDetailed) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FirewallDetailed) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *FirewallDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateSubnets

`func (o *FirewallDetailed) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *FirewallDetailed) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *FirewallDetailed) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



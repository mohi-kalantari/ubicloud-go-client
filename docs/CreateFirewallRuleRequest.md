# CreateFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR of the firewall rule | 
**FirewallId** | Pointer to **string** |  | [optional] 
**PortRange** | Pointer to **string** | Port range of the firewall rule | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFirewallRuleRequest

`func NewCreateFirewallRuleRequest(cidr string, ) *CreateFirewallRuleRequest`

NewCreateFirewallRuleRequest instantiates a new CreateFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRuleRequestWithDefaults

`func NewCreateFirewallRuleRequestWithDefaults() *CreateFirewallRuleRequest`

NewCreateFirewallRuleRequestWithDefaults instantiates a new CreateFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CreateFirewallRuleRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateFirewallRuleRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateFirewallRuleRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetFirewallId

`func (o *CreateFirewallRuleRequest) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *CreateFirewallRuleRequest) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *CreateFirewallRuleRequest) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *CreateFirewallRuleRequest) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### GetPortRange

`func (o *CreateFirewallRuleRequest) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *CreateFirewallRuleRequest) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *CreateFirewallRuleRequest) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *CreateFirewallRuleRequest) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateFirewallRuleRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateFirewallRuleRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateFirewallRuleRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateFirewallRuleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



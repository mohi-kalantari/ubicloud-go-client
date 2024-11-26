# PrivateSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firewalls** | [**[]Firewall**](Firewall.md) |  | 
**Id** | **string** | ID of the subnet | 
**Location** | **string** | Location of the subnet | 
**Name** | **string** | Name of the subnet | 
**Net4** | **string** | IPv4 CIDR of the subnet | 
**Net6** | **string** | IPv6 CIDR of the subnet | 
**Nics** | [**[]Nic**](Nic.md) | List of NICs | 
**State** | **string** | State of the subnet | 

## Methods

### NewPrivateSubnet

`func NewPrivateSubnet(firewalls []Firewall, id string, location string, name string, net4 string, net6 string, nics []Nic, state string, ) *PrivateSubnet`

NewPrivateSubnet instantiates a new PrivateSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSubnetWithDefaults

`func NewPrivateSubnetWithDefaults() *PrivateSubnet`

NewPrivateSubnetWithDefaults instantiates a new PrivateSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewalls

`func (o *PrivateSubnet) GetFirewalls() []Firewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *PrivateSubnet) GetFirewallsOk() (*[]Firewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *PrivateSubnet) SetFirewalls(v []Firewall)`

SetFirewalls sets Firewalls field to given value.


### GetId

`func (o *PrivateSubnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSubnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSubnet) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *PrivateSubnet) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrivateSubnet) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrivateSubnet) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *PrivateSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSubnet) SetName(v string)`

SetName sets Name field to given value.


### GetNet4

`func (o *PrivateSubnet) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *PrivateSubnet) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *PrivateSubnet) SetNet4(v string)`

SetNet4 sets Net4 field to given value.


### GetNet6

`func (o *PrivateSubnet) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *PrivateSubnet) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *PrivateSubnet) SetNet6(v string)`

SetNet6 sets Net6 field to given value.


### GetNics

`func (o *PrivateSubnet) GetNics() []Nic`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *PrivateSubnet) GetNicsOk() (*[]Nic, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *PrivateSubnet) SetNics(v []Nic)`

SetNics sets Nics field to given value.


### GetState

`func (o *PrivateSubnet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrivateSubnet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrivateSubnet) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



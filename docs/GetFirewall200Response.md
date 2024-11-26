# GetFirewall200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Items** | [**[]Firewall**](Firewall.md) |  | 

## Methods

### NewGetFirewall200Response

`func NewGetFirewall200Response(count int32, items []Firewall, ) *GetFirewall200Response`

NewGetFirewall200Response instantiates a new GetFirewall200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFirewall200ResponseWithDefaults

`func NewGetFirewall200ResponseWithDefaults() *GetFirewall200Response`

NewGetFirewall200ResponseWithDefaults instantiates a new GetFirewall200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetFirewall200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetFirewall200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetFirewall200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *GetFirewall200Response) GetItems() []Firewall`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetFirewall200Response) GetItemsOk() (*[]Firewall, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetFirewall200Response) SetItems(v []Firewall)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



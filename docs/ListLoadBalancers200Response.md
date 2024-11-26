# ListLoadBalancers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Items** | [**[]LoadBalancer**](LoadBalancer.md) |  | 

## Methods

### NewListLoadBalancers200Response

`func NewListLoadBalancers200Response(count int32, items []LoadBalancer, ) *ListLoadBalancers200Response`

NewListLoadBalancers200Response instantiates a new ListLoadBalancers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancers200ResponseWithDefaults

`func NewListLoadBalancers200ResponseWithDefaults() *ListLoadBalancers200Response`

NewListLoadBalancers200ResponseWithDefaults instantiates a new ListLoadBalancers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListLoadBalancers200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListLoadBalancers200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListLoadBalancers200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *ListLoadBalancers200Response) GetItems() []LoadBalancer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListLoadBalancers200Response) GetItemsOk() (*[]LoadBalancer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListLoadBalancers200Response) SetItems(v []LoadBalancer)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



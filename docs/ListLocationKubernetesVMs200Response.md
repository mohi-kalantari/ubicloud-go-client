# ListLocationKubernetesVMs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Items** | [**[]VmDetailed**](VmDetailed.md) |  | 

## Methods

### NewListLocationKubernetesVMs200Response

`func NewListLocationKubernetesVMs200Response(count int32, items []VmDetailed, ) *ListLocationKubernetesVMs200Response`

NewListLocationKubernetesVMs200Response instantiates a new ListLocationKubernetesVMs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLocationKubernetesVMs200ResponseWithDefaults

`func NewListLocationKubernetesVMs200ResponseWithDefaults() *ListLocationKubernetesVMs200Response`

NewListLocationKubernetesVMs200ResponseWithDefaults instantiates a new ListLocationKubernetesVMs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListLocationKubernetesVMs200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListLocationKubernetesVMs200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListLocationKubernetesVMs200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *ListLocationKubernetesVMs200Response) GetItems() []VmDetailed`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListLocationKubernetesVMs200Response) GetItemsOk() (*[]VmDetailed, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListLocationKubernetesVMs200Response) SetItems(v []VmDetailed)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



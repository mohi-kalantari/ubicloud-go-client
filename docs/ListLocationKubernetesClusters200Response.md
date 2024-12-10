# ListLocationKubernetesClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Items** | [**[]KubernetesCluster**](KubernetesCluster.md) |  | 

## Methods

### NewListLocationKubernetesClusters200Response

`func NewListLocationKubernetesClusters200Response(count int32, items []KubernetesCluster, ) *ListLocationKubernetesClusters200Response`

NewListLocationKubernetesClusters200Response instantiates a new ListLocationKubernetesClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLocationKubernetesClusters200ResponseWithDefaults

`func NewListLocationKubernetesClusters200ResponseWithDefaults() *ListLocationKubernetesClusters200Response`

NewListLocationKubernetesClusters200ResponseWithDefaults instantiates a new ListLocationKubernetesClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListLocationKubernetesClusters200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListLocationKubernetesClusters200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListLocationKubernetesClusters200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *ListLocationKubernetesClusters200Response) GetItems() []KubernetesCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListLocationKubernetesClusters200Response) GetItemsOk() (*[]KubernetesCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListLocationKubernetesClusters200Response) SetItems(v []KubernetesCluster)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



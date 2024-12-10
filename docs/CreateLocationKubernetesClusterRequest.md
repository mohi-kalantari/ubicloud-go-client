# CreateLocationKubernetesClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | **string** | Name of subnet | 
**KubernetesVersion** | **string** | Version of the kubernetes cluster | 

## Methods

### NewCreateLocationKubernetesClusterRequest

`func NewCreateLocationKubernetesClusterRequest(subnet string, kubernetesVersion string, ) *CreateLocationKubernetesClusterRequest`

NewCreateLocationKubernetesClusterRequest instantiates a new CreateLocationKubernetesClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLocationKubernetesClusterRequestWithDefaults

`func NewCreateLocationKubernetesClusterRequestWithDefaults() *CreateLocationKubernetesClusterRequest`

NewCreateLocationKubernetesClusterRequestWithDefaults instantiates a new CreateLocationKubernetesClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *CreateLocationKubernetesClusterRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateLocationKubernetesClusterRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateLocationKubernetesClusterRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetKubernetesVersion

`func (o *CreateLocationKubernetesClusterRequest) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *CreateLocationKubernetesClusterRequest) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *CreateLocationKubernetesClusterRequest) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



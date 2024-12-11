# CreateKubernetesClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | **string** | Name of subnet | 
**Location** | **string** | Location of the kubernetes cluster | 
**KubernetesVersion** | **string** | Version of the kubernetes cluster | 

## Methods

### NewCreateKubernetesClusterRequest

`func NewCreateKubernetesClusterRequest(subnet string, location string, kubernetesVersion string, ) *CreateKubernetesClusterRequest`

NewCreateKubernetesClusterRequest instantiates a new CreateKubernetesClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKubernetesClusterRequestWithDefaults

`func NewCreateKubernetesClusterRequestWithDefaults() *CreateKubernetesClusterRequest`

NewCreateKubernetesClusterRequestWithDefaults instantiates a new CreateKubernetesClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *CreateKubernetesClusterRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateKubernetesClusterRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateKubernetesClusterRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetLocation

`func (o *CreateKubernetesClusterRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateKubernetesClusterRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateKubernetesClusterRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetKubernetesVersion

`func (o *CreateKubernetesClusterRequest) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *CreateKubernetesClusterRequest) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *CreateKubernetesClusterRequest) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



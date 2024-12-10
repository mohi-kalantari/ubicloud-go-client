# KubernetesCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the kubernetes cluster | 
**Name** | **string** | Name of the kubernetes cluster | 
**Replica** | **int32** | Number of control plane replicas | 
**KubernetesVersion** | **string** | Version of the kubernetes cluster | 
**Subnet** | **string** | Name of the kubernetes cluster&#39;s subnet | 
**Location** | **string** | Location of the kubernetes cluster | 

## Methods

### NewKubernetesCluster

`func NewKubernetesCluster(id string, name string, replica int32, kubernetesVersion string, subnet string, location string, ) *KubernetesCluster`

NewKubernetesCluster instantiates a new KubernetesCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterWithDefaults

`func NewKubernetesClusterWithDefaults() *KubernetesCluster`

NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesCluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCluster) SetName(v string)`

SetName sets Name field to given value.


### GetReplica

`func (o *KubernetesCluster) GetReplica() int32`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *KubernetesCluster) GetReplicaOk() (*int32, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *KubernetesCluster) SetReplica(v int32)`

SetReplica sets Replica field to given value.


### GetKubernetesVersion

`func (o *KubernetesCluster) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesCluster) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesCluster) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetSubnet

`func (o *KubernetesCluster) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *KubernetesCluster) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *KubernetesCluster) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetLocation

`func (o *KubernetesCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesCluster) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



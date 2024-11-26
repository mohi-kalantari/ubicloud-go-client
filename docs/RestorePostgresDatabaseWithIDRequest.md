# RestorePostgresDatabaseWithIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RestoreTarget** | **string** |  | 

## Methods

### NewRestorePostgresDatabaseWithIDRequest

`func NewRestorePostgresDatabaseWithIDRequest(name string, restoreTarget string, ) *RestorePostgresDatabaseWithIDRequest`

NewRestorePostgresDatabaseWithIDRequest instantiates a new RestorePostgresDatabaseWithIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePostgresDatabaseWithIDRequestWithDefaults

`func NewRestorePostgresDatabaseWithIDRequestWithDefaults() *RestorePostgresDatabaseWithIDRequest`

NewRestorePostgresDatabaseWithIDRequestWithDefaults instantiates a new RestorePostgresDatabaseWithIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestorePostgresDatabaseWithIDRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestorePostgresDatabaseWithIDRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestorePostgresDatabaseWithIDRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRestoreTarget

`func (o *RestorePostgresDatabaseWithIDRequest) GetRestoreTarget() string`

GetRestoreTarget returns the RestoreTarget field if non-nil, zero value otherwise.

### GetRestoreTargetOk

`func (o *RestorePostgresDatabaseWithIDRequest) GetRestoreTargetOk() (*string, bool)`

GetRestoreTargetOk returns a tuple with the RestoreTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTarget

`func (o *RestorePostgresDatabaseWithIDRequest) SetRestoreTarget(v string)`

SetRestoreTarget sets RestoreTarget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



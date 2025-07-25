# V0042ClusterRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**V0040ClusterRecController**](V0040ClusterRecController.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | ClusterName | [optional] 
**Nodes** | Pointer to **string** | Node names | [optional] 
**SelectPlugin** | Pointer to **string** |  | [optional] 
**Associations** | Pointer to [**V0042ClusterRecAssociations**](V0042ClusterRecAssociations.md) |  | [optional] 
**RpcVersion** | Pointer to **int32** | RPC version used in the cluster | [optional] 
**Tres** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 

## Methods

### NewV0042ClusterRec

`func NewV0042ClusterRec() *V0042ClusterRec`

NewV0042ClusterRec instantiates a new V0042ClusterRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042ClusterRecWithDefaults

`func NewV0042ClusterRecWithDefaults() *V0042ClusterRec`

NewV0042ClusterRecWithDefaults instantiates a new V0042ClusterRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *V0042ClusterRec) GetController() V0040ClusterRecController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *V0042ClusterRec) GetControllerOk() (*V0040ClusterRecController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *V0042ClusterRec) SetController(v V0040ClusterRecController)`

SetController sets Controller field to given value.

### HasController

`func (o *V0042ClusterRec) HasController() bool`

HasController returns a boolean if a field has been set.

### GetFlags

`func (o *V0042ClusterRec) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0042ClusterRec) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0042ClusterRec) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0042ClusterRec) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0042ClusterRec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0042ClusterRec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0042ClusterRec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0042ClusterRec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *V0042ClusterRec) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042ClusterRec) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042ClusterRec) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042ClusterRec) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSelectPlugin

`func (o *V0042ClusterRec) GetSelectPlugin() string`

GetSelectPlugin returns the SelectPlugin field if non-nil, zero value otherwise.

### GetSelectPluginOk

`func (o *V0042ClusterRec) GetSelectPluginOk() (*string, bool)`

GetSelectPluginOk returns a tuple with the SelectPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPlugin

`func (o *V0042ClusterRec) SetSelectPlugin(v string)`

SetSelectPlugin sets SelectPlugin field to given value.

### HasSelectPlugin

`func (o *V0042ClusterRec) HasSelectPlugin() bool`

HasSelectPlugin returns a boolean if a field has been set.

### GetAssociations

`func (o *V0042ClusterRec) GetAssociations() V0042ClusterRecAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0042ClusterRec) GetAssociationsOk() (*V0042ClusterRecAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0042ClusterRec) SetAssociations(v V0042ClusterRecAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0042ClusterRec) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetRpcVersion

`func (o *V0042ClusterRec) GetRpcVersion() int32`

GetRpcVersion returns the RpcVersion field if non-nil, zero value otherwise.

### GetRpcVersionOk

`func (o *V0042ClusterRec) GetRpcVersionOk() (*int32, bool)`

GetRpcVersionOk returns a tuple with the RpcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcVersion

`func (o *V0042ClusterRec) SetRpcVersion(v int32)`

SetRpcVersion sets RpcVersion field to given value.

### HasRpcVersion

`func (o *V0042ClusterRec) HasRpcVersion() bool`

HasRpcVersion returns a boolean if a field has been set.

### GetTres

`func (o *V0042ClusterRec) GetTres() []V0042Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0042ClusterRec) GetTresOk() (*[]V0042Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0042ClusterRec) SetTres(v []V0042Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0042ClusterRec) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



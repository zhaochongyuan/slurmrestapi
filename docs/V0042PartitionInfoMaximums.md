# V0042PartitionInfoMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** | MaxMemPerCPU or MaxMemPerNode | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**Nodes** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Shares** | Pointer to **int32** | OverSubscribe | [optional] 
**Oversubscribe** | Pointer to [**V0042PartitionInfoMaximumsOversubscribe**](V0042PartitionInfoMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0042Uint16NoValStruct**](V0042Uint16NoValStruct.md) |  | [optional] 

## Methods

### NewV0042PartitionInfoMaximums

`func NewV0042PartitionInfoMaximums() *V0042PartitionInfoMaximums`

NewV0042PartitionInfoMaximums instantiates a new V0042PartitionInfoMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042PartitionInfoMaximumsWithDefaults

`func NewV0042PartitionInfoMaximumsWithDefaults() *V0042PartitionInfoMaximums`

NewV0042PartitionInfoMaximumsWithDefaults instantiates a new V0042PartitionInfoMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0042PartitionInfoMaximums) GetCpusPerNode() V0042Uint32NoValStruct`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0042PartitionInfoMaximums) GetCpusPerNodeOk() (*V0042Uint32NoValStruct, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0042PartitionInfoMaximums) SetCpusPerNode(v V0042Uint32NoValStruct)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0042PartitionInfoMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0042PartitionInfoMaximums) GetCpusPerSocket() V0042Uint32NoValStruct`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0042PartitionInfoMaximums) GetCpusPerSocketOk() (*V0042Uint32NoValStruct, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0042PartitionInfoMaximums) SetCpusPerSocket(v V0042Uint32NoValStruct)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0042PartitionInfoMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0042PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) GetPartitionMemoryPerCpu() V0042Uint64NoValStruct`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0042PartitionInfoMaximums) GetPartitionMemoryPerCpuOk() (*V0042Uint64NoValStruct, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) SetPartitionMemoryPerCpu(v V0042Uint64NoValStruct)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0042PartitionInfoMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0042PartitionInfoMaximums) GetPartitionMemoryPerNode() V0042Uint64NoValStruct`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0042PartitionInfoMaximums) GetPartitionMemoryPerNodeOk() (*V0042Uint64NoValStruct, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0042PartitionInfoMaximums) SetPartitionMemoryPerNode(v V0042Uint64NoValStruct)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0042PartitionInfoMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0042PartitionInfoMaximums) GetNodes() V0042Uint32NoValStruct`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042PartitionInfoMaximums) GetNodesOk() (*V0042Uint32NoValStruct, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042PartitionInfoMaximums) SetNodes(v V0042Uint32NoValStruct)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042PartitionInfoMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0042PartitionInfoMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0042PartitionInfoMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0042PartitionInfoMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0042PartitionInfoMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0042PartitionInfoMaximums) GetOversubscribe() V0042PartitionInfoMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0042PartitionInfoMaximums) GetOversubscribeOk() (*V0042PartitionInfoMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0042PartitionInfoMaximums) SetOversubscribe(v V0042PartitionInfoMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0042PartitionInfoMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0042PartitionInfoMaximums) GetTime() V0042Uint32NoValStruct`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042PartitionInfoMaximums) GetTimeOk() (*V0042Uint32NoValStruct, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042PartitionInfoMaximums) SetTime(v V0042Uint32NoValStruct)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042PartitionInfoMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0042PartitionInfoMaximums) GetOverTimeLimit() V0042Uint16NoValStruct`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0042PartitionInfoMaximums) GetOverTimeLimitOk() (*V0042Uint16NoValStruct, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0042PartitionInfoMaximums) SetOverTimeLimit(v V0042Uint16NoValStruct)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0042PartitionInfoMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



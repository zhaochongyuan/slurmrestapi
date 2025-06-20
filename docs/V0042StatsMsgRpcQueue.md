# V0042StatsMsgRpcQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Message type as integer | 
**MessageType** | **string** | Message type as string | 
**Count** | **int32** | Number of pending RPCs queued | 

## Methods

### NewV0042StatsMsgRpcQueue

`func NewV0042StatsMsgRpcQueue(typeId int32, messageType string, count int32, ) *V0042StatsMsgRpcQueue`

NewV0042StatsMsgRpcQueue instantiates a new V0042StatsMsgRpcQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsMsgRpcQueueWithDefaults

`func NewV0042StatsMsgRpcQueueWithDefaults() *V0042StatsMsgRpcQueue`

NewV0042StatsMsgRpcQueueWithDefaults instantiates a new V0042StatsMsgRpcQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *V0042StatsMsgRpcQueue) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *V0042StatsMsgRpcQueue) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *V0042StatsMsgRpcQueue) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *V0042StatsMsgRpcQueue) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V0042StatsMsgRpcQueue) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V0042StatsMsgRpcQueue) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *V0042StatsMsgRpcQueue) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0042StatsMsgRpcQueue) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0042StatsMsgRpcQueue) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# V0041KillJobsMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Filter jobs to a specific account | [optional] 
**Flags** | Pointer to **[]string** | Filter jobs according to flags | [optional] 
**JobName** | Pointer to **string** | Filter jobs to a specific name | [optional] 
**Jobs** | Pointer to **[]string** | List of jobs to signal | [optional] 
**Partition** | Pointer to **string** | Filter jobs to a specific partition | [optional] 
**Qos** | Pointer to **string** | Filter jobs to a specific QOS | [optional] 
**Reservation** | Pointer to **string** | Filter jobs to a specific reservation | [optional] 
**Signal** | Pointer to **string** | Signal to send to jobs | [optional] 
**JobState** | Pointer to **[]string** | Filter jobs to a specific state | [optional] 
**UserId** | Pointer to **string** | Filter jobs to a specific numeric user id | [optional] 
**UserName** | Pointer to **string** | Filter jobs to a specific user name | [optional] 
**Wckey** | Pointer to **string** | Filter jobs to a specific wckey | [optional] 
**Nodes** | Pointer to **[]string** | Filter jobs to a set of nodes | [optional] 

## Methods

### NewV0041KillJobsMsg

`func NewV0041KillJobsMsg() *V0041KillJobsMsg`

NewV0041KillJobsMsg instantiates a new V0041KillJobsMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041KillJobsMsgWithDefaults

`func NewV0041KillJobsMsgWithDefaults() *V0041KillJobsMsg`

NewV0041KillJobsMsgWithDefaults instantiates a new V0041KillJobsMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0041KillJobsMsg) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041KillJobsMsg) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041KillJobsMsg) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041KillJobsMsg) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFlags

`func (o *V0041KillJobsMsg) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041KillJobsMsg) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041KillJobsMsg) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041KillJobsMsg) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetJobName

`func (o *V0041KillJobsMsg) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *V0041KillJobsMsg) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *V0041KillJobsMsg) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *V0041KillJobsMsg) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobs

`func (o *V0041KillJobsMsg) GetJobs() []string`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0041KillJobsMsg) GetJobsOk() (*[]string, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0041KillJobsMsg) SetJobs(v []string)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0041KillJobsMsg) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPartition

`func (o *V0041KillJobsMsg) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041KillJobsMsg) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041KillJobsMsg) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041KillJobsMsg) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetQos

`func (o *V0041KillJobsMsg) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041KillJobsMsg) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041KillJobsMsg) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041KillJobsMsg) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReservation

`func (o *V0041KillJobsMsg) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0041KillJobsMsg) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0041KillJobsMsg) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0041KillJobsMsg) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetSignal

`func (o *V0041KillJobsMsg) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0041KillJobsMsg) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0041KillJobsMsg) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0041KillJobsMsg) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetJobState

`func (o *V0041KillJobsMsg) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0041KillJobsMsg) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0041KillJobsMsg) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0041KillJobsMsg) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetUserId

`func (o *V0041KillJobsMsg) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0041KillJobsMsg) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0041KillJobsMsg) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0041KillJobsMsg) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0041KillJobsMsg) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0041KillJobsMsg) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0041KillJobsMsg) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0041KillJobsMsg) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWckey

`func (o *V0041KillJobsMsg) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0041KillJobsMsg) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0041KillJobsMsg) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0041KillJobsMsg) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetNodes

`func (o *V0041KillJobsMsg) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041KillJobsMsg) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041KillJobsMsg) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041KillJobsMsg) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# V0042StatsMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartsPacked** | Pointer to **int32** | Zero if only RPC statistic included | [optional] 
**ReqTime** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**ReqTimeStart** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**ServerThreadCount** | Pointer to **int32** | Number of current active slurmctld threads | [optional] 
**AgentQueueSize** | Pointer to **int32** | Number of enqueued outgoing RPC requests in an internal retry list | [optional] 
**AgentCount** | Pointer to **int32** | Number of agent threads | [optional] 
**AgentThreadCount** | Pointer to **int32** | Total number of active threads created by all agent threads | [optional] 
**DbdAgentQueueSize** | Pointer to **int32** | Number of messages for SlurmDBD that are queued | [optional] 
**GettimeofdayLatency** | Pointer to **int32** | Latency of 1000 calls to the gettimeofday() syscall in microseconds, as measured at controller startup | [optional] 
**ScheduleCycleMax** | Pointer to **int32** | Max time of any scheduling cycle in microseconds since last reset | [optional] 
**ScheduleCycleLast** | Pointer to **int32** | Time in microseconds for last scheduling cycle | [optional] 
**ScheduleCycleSum** | Pointer to **int32** | Total run time in microseconds for all scheduling cycles since last reset | [optional] 
**ScheduleCycleTotal** | Pointer to **int32** | Number of scheduling cycles since last reset | [optional] 
**ScheduleCycleMean** | Pointer to **int64** | Mean time in microseconds for all scheduling cycles since last reset | [optional] 
**ScheduleCycleMeanDepth** | Pointer to **int64** | Mean of the number of jobs processed in a scheduling cycle | [optional] 
**ScheduleCyclePerMinute** | Pointer to **int64** | Number of scheduling executions per minute | [optional] 
**ScheduleCycleDepth** | Pointer to **int32** | Total number of jobs processed in scheduling cycles | [optional] 
**ScheduleExit** | Pointer to [**V0042ScheduleExitFields**](V0042ScheduleExitFields.md) |  | [optional] 
**ScheduleQueueLength** | Pointer to **int32** | Number of jobs pending in queue | [optional] 
**JobsSubmitted** | Pointer to **int32** | Number of jobs submitted since last reset | [optional] 
**JobsStarted** | Pointer to **int32** | Number of jobs started since last reset | [optional] 
**JobsCompleted** | Pointer to **int32** | Number of jobs completed since last reset | [optional] 
**JobsCanceled** | Pointer to **int32** | Number of jobs canceled since the last reset | [optional] 
**JobsFailed** | Pointer to **int32** | Number of jobs failed due to slurmd or other internal issues since last reset | [optional] 
**JobsPending** | Pointer to **int32** | Number of jobs pending at the time of listed in job_state_ts | [optional] 
**JobsRunning** | Pointer to **int32** | Number of jobs running at the time of listed in job_state_ts | [optional] 
**JobStatesTs** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**BfBackfilledJobs** | Pointer to **int32** | Number of jobs started through backfilling since last slurm start | [optional] 
**BfLastBackfilledJobs** | Pointer to **int32** | Number of jobs started through backfilling since last reset | [optional] 
**BfBackfilledHetJobs** | Pointer to **int32** | Number of heterogeneous job components started through backfilling since last Slurm start | [optional] 
**BfCycleCounter** | Pointer to **int32** | Number of backfill scheduling cycles since last reset | [optional] 
**BfCycleMean** | Pointer to **int64** | Mean time in microseconds of backfilling scheduling cycles since last reset | [optional] 
**BfDepthMean** | Pointer to **int64** | Mean number of eligible to run jobs processed during all backfilling scheduling cycles since last reset | [optional] 
**BfDepthMeanTry** | Pointer to **int64** | The subset of Depth Mean that the backfill scheduler attempted to schedule | [optional] 
**BfCycleSum** | Pointer to **int64** | Total time in microseconds of backfilling scheduling cycles since last reset | [optional] 
**BfCycleLast** | Pointer to **int32** | Execution time in microseconds of last backfill scheduling cycle | [optional] 
**BfCycleMax** | Pointer to **int32** | Execution time in microseconds of longest backfill scheduling cycle | [optional] 
**BfExit** | Pointer to [**V0042BfExitFields**](V0042BfExitFields.md) |  | [optional] 
**BfLastDepth** | Pointer to **int32** | Number of processed jobs during last backfilling scheduling cycle | [optional] 
**BfLastDepthTry** | Pointer to **int32** | Number of processed jobs during last backfilling scheduling cycle that had a chance to start using available resources | [optional] 
**BfDepthSum** | Pointer to **int32** | Total number of jobs processed during all backfilling scheduling cycles since last reset | [optional] 
**BfDepthTrySum** | Pointer to **int32** | Subset of bf_depth_sum that the backfill scheduler attempted to schedule | [optional] 
**BfQueueLen** | Pointer to **int32** | Number of jobs pending to be processed by backfilling algorithm | [optional] 
**BfQueueLenMean** | Pointer to **int64** | Mean number of jobs pending to be processed by backfilling algorithm | [optional] 
**BfQueueLenSum** | Pointer to **int32** | Total number of jobs pending to be processed by backfilling algorithm since last reset | [optional] 
**BfTableSize** | Pointer to **int32** | Number of different time slots tested by the backfill scheduler in its last iteration | [optional] 
**BfTableSizeSum** | Pointer to **int32** | Total number of different time slots tested by the backfill scheduler | [optional] 
**BfTableSizeMean** | Pointer to **int64** | Mean number of different time slots tested by the backfill scheduler | [optional] 
**BfWhenLastCycle** | Pointer to [**V0042Uint64NoValStruct**](V0042Uint64NoValStruct.md) |  | [optional] 
**BfActive** | Pointer to **bool** | Backfill scheduler currently running | [optional] 
**RpcsByMessageType** | Pointer to [**[]V0042StatsMsgRpcType**](V0042StatsMsgRpcType.md) | RPCs by type | [optional] 
**RpcsByUser** | Pointer to [**[]V0042StatsMsgRpcUser**](V0042StatsMsgRpcUser.md) | RPCs by user | [optional] 
**PendingRpcs** | Pointer to [**[]V0042StatsMsgRpcQueue**](V0042StatsMsgRpcQueue.md) | Pending RPCs | [optional] 
**PendingRpcsByHostlist** | Pointer to [**[]V0042StatsMsgRpcDump**](V0042StatsMsgRpcDump.md) | Pending RPCs by hostlist | [optional] 

## Methods

### NewV0042StatsMsg

`func NewV0042StatsMsg() *V0042StatsMsg`

NewV0042StatsMsg instantiates a new V0042StatsMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StatsMsgWithDefaults

`func NewV0042StatsMsgWithDefaults() *V0042StatsMsg`

NewV0042StatsMsgWithDefaults instantiates a new V0042StatsMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartsPacked

`func (o *V0042StatsMsg) GetPartsPacked() int32`

GetPartsPacked returns the PartsPacked field if non-nil, zero value otherwise.

### GetPartsPackedOk

`func (o *V0042StatsMsg) GetPartsPackedOk() (*int32, bool)`

GetPartsPackedOk returns a tuple with the PartsPacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPacked

`func (o *V0042StatsMsg) SetPartsPacked(v int32)`

SetPartsPacked sets PartsPacked field to given value.

### HasPartsPacked

`func (o *V0042StatsMsg) HasPartsPacked() bool`

HasPartsPacked returns a boolean if a field has been set.

### GetReqTime

`func (o *V0042StatsMsg) GetReqTime() V0042Uint64NoValStruct`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *V0042StatsMsg) GetReqTimeOk() (*V0042Uint64NoValStruct, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *V0042StatsMsg) SetReqTime(v V0042Uint64NoValStruct)`

SetReqTime sets ReqTime field to given value.

### HasReqTime

`func (o *V0042StatsMsg) HasReqTime() bool`

HasReqTime returns a boolean if a field has been set.

### GetReqTimeStart

`func (o *V0042StatsMsg) GetReqTimeStart() V0042Uint64NoValStruct`

GetReqTimeStart returns the ReqTimeStart field if non-nil, zero value otherwise.

### GetReqTimeStartOk

`func (o *V0042StatsMsg) GetReqTimeStartOk() (*V0042Uint64NoValStruct, bool)`

GetReqTimeStartOk returns a tuple with the ReqTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeStart

`func (o *V0042StatsMsg) SetReqTimeStart(v V0042Uint64NoValStruct)`

SetReqTimeStart sets ReqTimeStart field to given value.

### HasReqTimeStart

`func (o *V0042StatsMsg) HasReqTimeStart() bool`

HasReqTimeStart returns a boolean if a field has been set.

### GetServerThreadCount

`func (o *V0042StatsMsg) GetServerThreadCount() int32`

GetServerThreadCount returns the ServerThreadCount field if non-nil, zero value otherwise.

### GetServerThreadCountOk

`func (o *V0042StatsMsg) GetServerThreadCountOk() (*int32, bool)`

GetServerThreadCountOk returns a tuple with the ServerThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThreadCount

`func (o *V0042StatsMsg) SetServerThreadCount(v int32)`

SetServerThreadCount sets ServerThreadCount field to given value.

### HasServerThreadCount

`func (o *V0042StatsMsg) HasServerThreadCount() bool`

HasServerThreadCount returns a boolean if a field has been set.

### GetAgentQueueSize

`func (o *V0042StatsMsg) GetAgentQueueSize() int32`

GetAgentQueueSize returns the AgentQueueSize field if non-nil, zero value otherwise.

### GetAgentQueueSizeOk

`func (o *V0042StatsMsg) GetAgentQueueSizeOk() (*int32, bool)`

GetAgentQueueSizeOk returns a tuple with the AgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentQueueSize

`func (o *V0042StatsMsg) SetAgentQueueSize(v int32)`

SetAgentQueueSize sets AgentQueueSize field to given value.

### HasAgentQueueSize

`func (o *V0042StatsMsg) HasAgentQueueSize() bool`

HasAgentQueueSize returns a boolean if a field has been set.

### GetAgentCount

`func (o *V0042StatsMsg) GetAgentCount() int32`

GetAgentCount returns the AgentCount field if non-nil, zero value otherwise.

### GetAgentCountOk

`func (o *V0042StatsMsg) GetAgentCountOk() (*int32, bool)`

GetAgentCountOk returns a tuple with the AgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCount

`func (o *V0042StatsMsg) SetAgentCount(v int32)`

SetAgentCount sets AgentCount field to given value.

### HasAgentCount

`func (o *V0042StatsMsg) HasAgentCount() bool`

HasAgentCount returns a boolean if a field has been set.

### GetAgentThreadCount

`func (o *V0042StatsMsg) GetAgentThreadCount() int32`

GetAgentThreadCount returns the AgentThreadCount field if non-nil, zero value otherwise.

### GetAgentThreadCountOk

`func (o *V0042StatsMsg) GetAgentThreadCountOk() (*int32, bool)`

GetAgentThreadCountOk returns a tuple with the AgentThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentThreadCount

`func (o *V0042StatsMsg) SetAgentThreadCount(v int32)`

SetAgentThreadCount sets AgentThreadCount field to given value.

### HasAgentThreadCount

`func (o *V0042StatsMsg) HasAgentThreadCount() bool`

HasAgentThreadCount returns a boolean if a field has been set.

### GetDbdAgentQueueSize

`func (o *V0042StatsMsg) GetDbdAgentQueueSize() int32`

GetDbdAgentQueueSize returns the DbdAgentQueueSize field if non-nil, zero value otherwise.

### GetDbdAgentQueueSizeOk

`func (o *V0042StatsMsg) GetDbdAgentQueueSizeOk() (*int32, bool)`

GetDbdAgentQueueSizeOk returns a tuple with the DbdAgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbdAgentQueueSize

`func (o *V0042StatsMsg) SetDbdAgentQueueSize(v int32)`

SetDbdAgentQueueSize sets DbdAgentQueueSize field to given value.

### HasDbdAgentQueueSize

`func (o *V0042StatsMsg) HasDbdAgentQueueSize() bool`

HasDbdAgentQueueSize returns a boolean if a field has been set.

### GetGettimeofdayLatency

`func (o *V0042StatsMsg) GetGettimeofdayLatency() int32`

GetGettimeofdayLatency returns the GettimeofdayLatency field if non-nil, zero value otherwise.

### GetGettimeofdayLatencyOk

`func (o *V0042StatsMsg) GetGettimeofdayLatencyOk() (*int32, bool)`

GetGettimeofdayLatencyOk returns a tuple with the GettimeofdayLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettimeofdayLatency

`func (o *V0042StatsMsg) SetGettimeofdayLatency(v int32)`

SetGettimeofdayLatency sets GettimeofdayLatency field to given value.

### HasGettimeofdayLatency

`func (o *V0042StatsMsg) HasGettimeofdayLatency() bool`

HasGettimeofdayLatency returns a boolean if a field has been set.

### GetScheduleCycleMax

`func (o *V0042StatsMsg) GetScheduleCycleMax() int32`

GetScheduleCycleMax returns the ScheduleCycleMax field if non-nil, zero value otherwise.

### GetScheduleCycleMaxOk

`func (o *V0042StatsMsg) GetScheduleCycleMaxOk() (*int32, bool)`

GetScheduleCycleMaxOk returns a tuple with the ScheduleCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMax

`func (o *V0042StatsMsg) SetScheduleCycleMax(v int32)`

SetScheduleCycleMax sets ScheduleCycleMax field to given value.

### HasScheduleCycleMax

`func (o *V0042StatsMsg) HasScheduleCycleMax() bool`

HasScheduleCycleMax returns a boolean if a field has been set.

### GetScheduleCycleLast

`func (o *V0042StatsMsg) GetScheduleCycleLast() int32`

GetScheduleCycleLast returns the ScheduleCycleLast field if non-nil, zero value otherwise.

### GetScheduleCycleLastOk

`func (o *V0042StatsMsg) GetScheduleCycleLastOk() (*int32, bool)`

GetScheduleCycleLastOk returns a tuple with the ScheduleCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleLast

`func (o *V0042StatsMsg) SetScheduleCycleLast(v int32)`

SetScheduleCycleLast sets ScheduleCycleLast field to given value.

### HasScheduleCycleLast

`func (o *V0042StatsMsg) HasScheduleCycleLast() bool`

HasScheduleCycleLast returns a boolean if a field has been set.

### GetScheduleCycleSum

`func (o *V0042StatsMsg) GetScheduleCycleSum() int32`

GetScheduleCycleSum returns the ScheduleCycleSum field if non-nil, zero value otherwise.

### GetScheduleCycleSumOk

`func (o *V0042StatsMsg) GetScheduleCycleSumOk() (*int32, bool)`

GetScheduleCycleSumOk returns a tuple with the ScheduleCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleSum

`func (o *V0042StatsMsg) SetScheduleCycleSum(v int32)`

SetScheduleCycleSum sets ScheduleCycleSum field to given value.

### HasScheduleCycleSum

`func (o *V0042StatsMsg) HasScheduleCycleSum() bool`

HasScheduleCycleSum returns a boolean if a field has been set.

### GetScheduleCycleTotal

`func (o *V0042StatsMsg) GetScheduleCycleTotal() int32`

GetScheduleCycleTotal returns the ScheduleCycleTotal field if non-nil, zero value otherwise.

### GetScheduleCycleTotalOk

`func (o *V0042StatsMsg) GetScheduleCycleTotalOk() (*int32, bool)`

GetScheduleCycleTotalOk returns a tuple with the ScheduleCycleTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleTotal

`func (o *V0042StatsMsg) SetScheduleCycleTotal(v int32)`

SetScheduleCycleTotal sets ScheduleCycleTotal field to given value.

### HasScheduleCycleTotal

`func (o *V0042StatsMsg) HasScheduleCycleTotal() bool`

HasScheduleCycleTotal returns a boolean if a field has been set.

### GetScheduleCycleMean

`func (o *V0042StatsMsg) GetScheduleCycleMean() int64`

GetScheduleCycleMean returns the ScheduleCycleMean field if non-nil, zero value otherwise.

### GetScheduleCycleMeanOk

`func (o *V0042StatsMsg) GetScheduleCycleMeanOk() (*int64, bool)`

GetScheduleCycleMeanOk returns a tuple with the ScheduleCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMean

`func (o *V0042StatsMsg) SetScheduleCycleMean(v int64)`

SetScheduleCycleMean sets ScheduleCycleMean field to given value.

### HasScheduleCycleMean

`func (o *V0042StatsMsg) HasScheduleCycleMean() bool`

HasScheduleCycleMean returns a boolean if a field has been set.

### GetScheduleCycleMeanDepth

`func (o *V0042StatsMsg) GetScheduleCycleMeanDepth() int64`

GetScheduleCycleMeanDepth returns the ScheduleCycleMeanDepth field if non-nil, zero value otherwise.

### GetScheduleCycleMeanDepthOk

`func (o *V0042StatsMsg) GetScheduleCycleMeanDepthOk() (*int64, bool)`

GetScheduleCycleMeanDepthOk returns a tuple with the ScheduleCycleMeanDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMeanDepth

`func (o *V0042StatsMsg) SetScheduleCycleMeanDepth(v int64)`

SetScheduleCycleMeanDepth sets ScheduleCycleMeanDepth field to given value.

### HasScheduleCycleMeanDepth

`func (o *V0042StatsMsg) HasScheduleCycleMeanDepth() bool`

HasScheduleCycleMeanDepth returns a boolean if a field has been set.

### GetScheduleCyclePerMinute

`func (o *V0042StatsMsg) GetScheduleCyclePerMinute() int64`

GetScheduleCyclePerMinute returns the ScheduleCyclePerMinute field if non-nil, zero value otherwise.

### GetScheduleCyclePerMinuteOk

`func (o *V0042StatsMsg) GetScheduleCyclePerMinuteOk() (*int64, bool)`

GetScheduleCyclePerMinuteOk returns a tuple with the ScheduleCyclePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCyclePerMinute

`func (o *V0042StatsMsg) SetScheduleCyclePerMinute(v int64)`

SetScheduleCyclePerMinute sets ScheduleCyclePerMinute field to given value.

### HasScheduleCyclePerMinute

`func (o *V0042StatsMsg) HasScheduleCyclePerMinute() bool`

HasScheduleCyclePerMinute returns a boolean if a field has been set.

### GetScheduleCycleDepth

`func (o *V0042StatsMsg) GetScheduleCycleDepth() int32`

GetScheduleCycleDepth returns the ScheduleCycleDepth field if non-nil, zero value otherwise.

### GetScheduleCycleDepthOk

`func (o *V0042StatsMsg) GetScheduleCycleDepthOk() (*int32, bool)`

GetScheduleCycleDepthOk returns a tuple with the ScheduleCycleDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleDepth

`func (o *V0042StatsMsg) SetScheduleCycleDepth(v int32)`

SetScheduleCycleDepth sets ScheduleCycleDepth field to given value.

### HasScheduleCycleDepth

`func (o *V0042StatsMsg) HasScheduleCycleDepth() bool`

HasScheduleCycleDepth returns a boolean if a field has been set.

### GetScheduleExit

`func (o *V0042StatsMsg) GetScheduleExit() V0042ScheduleExitFields`

GetScheduleExit returns the ScheduleExit field if non-nil, zero value otherwise.

### GetScheduleExitOk

`func (o *V0042StatsMsg) GetScheduleExitOk() (*V0042ScheduleExitFields, bool)`

GetScheduleExitOk returns a tuple with the ScheduleExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExit

`func (o *V0042StatsMsg) SetScheduleExit(v V0042ScheduleExitFields)`

SetScheduleExit sets ScheduleExit field to given value.

### HasScheduleExit

`func (o *V0042StatsMsg) HasScheduleExit() bool`

HasScheduleExit returns a boolean if a field has been set.

### GetScheduleQueueLength

`func (o *V0042StatsMsg) GetScheduleQueueLength() int32`

GetScheduleQueueLength returns the ScheduleQueueLength field if non-nil, zero value otherwise.

### GetScheduleQueueLengthOk

`func (o *V0042StatsMsg) GetScheduleQueueLengthOk() (*int32, bool)`

GetScheduleQueueLengthOk returns a tuple with the ScheduleQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleQueueLength

`func (o *V0042StatsMsg) SetScheduleQueueLength(v int32)`

SetScheduleQueueLength sets ScheduleQueueLength field to given value.

### HasScheduleQueueLength

`func (o *V0042StatsMsg) HasScheduleQueueLength() bool`

HasScheduleQueueLength returns a boolean if a field has been set.

### GetJobsSubmitted

`func (o *V0042StatsMsg) GetJobsSubmitted() int32`

GetJobsSubmitted returns the JobsSubmitted field if non-nil, zero value otherwise.

### GetJobsSubmittedOk

`func (o *V0042StatsMsg) GetJobsSubmittedOk() (*int32, bool)`

GetJobsSubmittedOk returns a tuple with the JobsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSubmitted

`func (o *V0042StatsMsg) SetJobsSubmitted(v int32)`

SetJobsSubmitted sets JobsSubmitted field to given value.

### HasJobsSubmitted

`func (o *V0042StatsMsg) HasJobsSubmitted() bool`

HasJobsSubmitted returns a boolean if a field has been set.

### GetJobsStarted

`func (o *V0042StatsMsg) GetJobsStarted() int32`

GetJobsStarted returns the JobsStarted field if non-nil, zero value otherwise.

### GetJobsStartedOk

`func (o *V0042StatsMsg) GetJobsStartedOk() (*int32, bool)`

GetJobsStartedOk returns a tuple with the JobsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsStarted

`func (o *V0042StatsMsg) SetJobsStarted(v int32)`

SetJobsStarted sets JobsStarted field to given value.

### HasJobsStarted

`func (o *V0042StatsMsg) HasJobsStarted() bool`

HasJobsStarted returns a boolean if a field has been set.

### GetJobsCompleted

`func (o *V0042StatsMsg) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *V0042StatsMsg) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *V0042StatsMsg) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.

### HasJobsCompleted

`func (o *V0042StatsMsg) HasJobsCompleted() bool`

HasJobsCompleted returns a boolean if a field has been set.

### GetJobsCanceled

`func (o *V0042StatsMsg) GetJobsCanceled() int32`

GetJobsCanceled returns the JobsCanceled field if non-nil, zero value otherwise.

### GetJobsCanceledOk

`func (o *V0042StatsMsg) GetJobsCanceledOk() (*int32, bool)`

GetJobsCanceledOk returns a tuple with the JobsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCanceled

`func (o *V0042StatsMsg) SetJobsCanceled(v int32)`

SetJobsCanceled sets JobsCanceled field to given value.

### HasJobsCanceled

`func (o *V0042StatsMsg) HasJobsCanceled() bool`

HasJobsCanceled returns a boolean if a field has been set.

### GetJobsFailed

`func (o *V0042StatsMsg) GetJobsFailed() int32`

GetJobsFailed returns the JobsFailed field if non-nil, zero value otherwise.

### GetJobsFailedOk

`func (o *V0042StatsMsg) GetJobsFailedOk() (*int32, bool)`

GetJobsFailedOk returns a tuple with the JobsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFailed

`func (o *V0042StatsMsg) SetJobsFailed(v int32)`

SetJobsFailed sets JobsFailed field to given value.

### HasJobsFailed

`func (o *V0042StatsMsg) HasJobsFailed() bool`

HasJobsFailed returns a boolean if a field has been set.

### GetJobsPending

`func (o *V0042StatsMsg) GetJobsPending() int32`

GetJobsPending returns the JobsPending field if non-nil, zero value otherwise.

### GetJobsPendingOk

`func (o *V0042StatsMsg) GetJobsPendingOk() (*int32, bool)`

GetJobsPendingOk returns a tuple with the JobsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsPending

`func (o *V0042StatsMsg) SetJobsPending(v int32)`

SetJobsPending sets JobsPending field to given value.

### HasJobsPending

`func (o *V0042StatsMsg) HasJobsPending() bool`

HasJobsPending returns a boolean if a field has been set.

### GetJobsRunning

`func (o *V0042StatsMsg) GetJobsRunning() int32`

GetJobsRunning returns the JobsRunning field if non-nil, zero value otherwise.

### GetJobsRunningOk

`func (o *V0042StatsMsg) GetJobsRunningOk() (*int32, bool)`

GetJobsRunningOk returns a tuple with the JobsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsRunning

`func (o *V0042StatsMsg) SetJobsRunning(v int32)`

SetJobsRunning sets JobsRunning field to given value.

### HasJobsRunning

`func (o *V0042StatsMsg) HasJobsRunning() bool`

HasJobsRunning returns a boolean if a field has been set.

### GetJobStatesTs

`func (o *V0042StatsMsg) GetJobStatesTs() V0042Uint64NoValStruct`

GetJobStatesTs returns the JobStatesTs field if non-nil, zero value otherwise.

### GetJobStatesTsOk

`func (o *V0042StatsMsg) GetJobStatesTsOk() (*V0042Uint64NoValStruct, bool)`

GetJobStatesTsOk returns a tuple with the JobStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatesTs

`func (o *V0042StatsMsg) SetJobStatesTs(v V0042Uint64NoValStruct)`

SetJobStatesTs sets JobStatesTs field to given value.

### HasJobStatesTs

`func (o *V0042StatsMsg) HasJobStatesTs() bool`

HasJobStatesTs returns a boolean if a field has been set.

### GetBfBackfilledJobs

`func (o *V0042StatsMsg) GetBfBackfilledJobs() int32`

GetBfBackfilledJobs returns the BfBackfilledJobs field if non-nil, zero value otherwise.

### GetBfBackfilledJobsOk

`func (o *V0042StatsMsg) GetBfBackfilledJobsOk() (*int32, bool)`

GetBfBackfilledJobsOk returns a tuple with the BfBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledJobs

`func (o *V0042StatsMsg) SetBfBackfilledJobs(v int32)`

SetBfBackfilledJobs sets BfBackfilledJobs field to given value.

### HasBfBackfilledJobs

`func (o *V0042StatsMsg) HasBfBackfilledJobs() bool`

HasBfBackfilledJobs returns a boolean if a field has been set.

### GetBfLastBackfilledJobs

`func (o *V0042StatsMsg) GetBfLastBackfilledJobs() int32`

GetBfLastBackfilledJobs returns the BfLastBackfilledJobs field if non-nil, zero value otherwise.

### GetBfLastBackfilledJobsOk

`func (o *V0042StatsMsg) GetBfLastBackfilledJobsOk() (*int32, bool)`

GetBfLastBackfilledJobsOk returns a tuple with the BfLastBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastBackfilledJobs

`func (o *V0042StatsMsg) SetBfLastBackfilledJobs(v int32)`

SetBfLastBackfilledJobs sets BfLastBackfilledJobs field to given value.

### HasBfLastBackfilledJobs

`func (o *V0042StatsMsg) HasBfLastBackfilledJobs() bool`

HasBfLastBackfilledJobs returns a boolean if a field has been set.

### GetBfBackfilledHetJobs

`func (o *V0042StatsMsg) GetBfBackfilledHetJobs() int32`

GetBfBackfilledHetJobs returns the BfBackfilledHetJobs field if non-nil, zero value otherwise.

### GetBfBackfilledHetJobsOk

`func (o *V0042StatsMsg) GetBfBackfilledHetJobsOk() (*int32, bool)`

GetBfBackfilledHetJobsOk returns a tuple with the BfBackfilledHetJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledHetJobs

`func (o *V0042StatsMsg) SetBfBackfilledHetJobs(v int32)`

SetBfBackfilledHetJobs sets BfBackfilledHetJobs field to given value.

### HasBfBackfilledHetJobs

`func (o *V0042StatsMsg) HasBfBackfilledHetJobs() bool`

HasBfBackfilledHetJobs returns a boolean if a field has been set.

### GetBfCycleCounter

`func (o *V0042StatsMsg) GetBfCycleCounter() int32`

GetBfCycleCounter returns the BfCycleCounter field if non-nil, zero value otherwise.

### GetBfCycleCounterOk

`func (o *V0042StatsMsg) GetBfCycleCounterOk() (*int32, bool)`

GetBfCycleCounterOk returns a tuple with the BfCycleCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleCounter

`func (o *V0042StatsMsg) SetBfCycleCounter(v int32)`

SetBfCycleCounter sets BfCycleCounter field to given value.

### HasBfCycleCounter

`func (o *V0042StatsMsg) HasBfCycleCounter() bool`

HasBfCycleCounter returns a boolean if a field has been set.

### GetBfCycleMean

`func (o *V0042StatsMsg) GetBfCycleMean() int64`

GetBfCycleMean returns the BfCycleMean field if non-nil, zero value otherwise.

### GetBfCycleMeanOk

`func (o *V0042StatsMsg) GetBfCycleMeanOk() (*int64, bool)`

GetBfCycleMeanOk returns a tuple with the BfCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMean

`func (o *V0042StatsMsg) SetBfCycleMean(v int64)`

SetBfCycleMean sets BfCycleMean field to given value.

### HasBfCycleMean

`func (o *V0042StatsMsg) HasBfCycleMean() bool`

HasBfCycleMean returns a boolean if a field has been set.

### GetBfDepthMean

`func (o *V0042StatsMsg) GetBfDepthMean() int64`

GetBfDepthMean returns the BfDepthMean field if non-nil, zero value otherwise.

### GetBfDepthMeanOk

`func (o *V0042StatsMsg) GetBfDepthMeanOk() (*int64, bool)`

GetBfDepthMeanOk returns a tuple with the BfDepthMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMean

`func (o *V0042StatsMsg) SetBfDepthMean(v int64)`

SetBfDepthMean sets BfDepthMean field to given value.

### HasBfDepthMean

`func (o *V0042StatsMsg) HasBfDepthMean() bool`

HasBfDepthMean returns a boolean if a field has been set.

### GetBfDepthMeanTry

`func (o *V0042StatsMsg) GetBfDepthMeanTry() int64`

GetBfDepthMeanTry returns the BfDepthMeanTry field if non-nil, zero value otherwise.

### GetBfDepthMeanTryOk

`func (o *V0042StatsMsg) GetBfDepthMeanTryOk() (*int64, bool)`

GetBfDepthMeanTryOk returns a tuple with the BfDepthMeanTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMeanTry

`func (o *V0042StatsMsg) SetBfDepthMeanTry(v int64)`

SetBfDepthMeanTry sets BfDepthMeanTry field to given value.

### HasBfDepthMeanTry

`func (o *V0042StatsMsg) HasBfDepthMeanTry() bool`

HasBfDepthMeanTry returns a boolean if a field has been set.

### GetBfCycleSum

`func (o *V0042StatsMsg) GetBfCycleSum() int64`

GetBfCycleSum returns the BfCycleSum field if non-nil, zero value otherwise.

### GetBfCycleSumOk

`func (o *V0042StatsMsg) GetBfCycleSumOk() (*int64, bool)`

GetBfCycleSumOk returns a tuple with the BfCycleSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleSum

`func (o *V0042StatsMsg) SetBfCycleSum(v int64)`

SetBfCycleSum sets BfCycleSum field to given value.

### HasBfCycleSum

`func (o *V0042StatsMsg) HasBfCycleSum() bool`

HasBfCycleSum returns a boolean if a field has been set.

### GetBfCycleLast

`func (o *V0042StatsMsg) GetBfCycleLast() int32`

GetBfCycleLast returns the BfCycleLast field if non-nil, zero value otherwise.

### GetBfCycleLastOk

`func (o *V0042StatsMsg) GetBfCycleLastOk() (*int32, bool)`

GetBfCycleLastOk returns a tuple with the BfCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleLast

`func (o *V0042StatsMsg) SetBfCycleLast(v int32)`

SetBfCycleLast sets BfCycleLast field to given value.

### HasBfCycleLast

`func (o *V0042StatsMsg) HasBfCycleLast() bool`

HasBfCycleLast returns a boolean if a field has been set.

### GetBfCycleMax

`func (o *V0042StatsMsg) GetBfCycleMax() int32`

GetBfCycleMax returns the BfCycleMax field if non-nil, zero value otherwise.

### GetBfCycleMaxOk

`func (o *V0042StatsMsg) GetBfCycleMaxOk() (*int32, bool)`

GetBfCycleMaxOk returns a tuple with the BfCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMax

`func (o *V0042StatsMsg) SetBfCycleMax(v int32)`

SetBfCycleMax sets BfCycleMax field to given value.

### HasBfCycleMax

`func (o *V0042StatsMsg) HasBfCycleMax() bool`

HasBfCycleMax returns a boolean if a field has been set.

### GetBfExit

`func (o *V0042StatsMsg) GetBfExit() V0042BfExitFields`

GetBfExit returns the BfExit field if non-nil, zero value otherwise.

### GetBfExitOk

`func (o *V0042StatsMsg) GetBfExitOk() (*V0042BfExitFields, bool)`

GetBfExitOk returns a tuple with the BfExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfExit

`func (o *V0042StatsMsg) SetBfExit(v V0042BfExitFields)`

SetBfExit sets BfExit field to given value.

### HasBfExit

`func (o *V0042StatsMsg) HasBfExit() bool`

HasBfExit returns a boolean if a field has been set.

### GetBfLastDepth

`func (o *V0042StatsMsg) GetBfLastDepth() int32`

GetBfLastDepth returns the BfLastDepth field if non-nil, zero value otherwise.

### GetBfLastDepthOk

`func (o *V0042StatsMsg) GetBfLastDepthOk() (*int32, bool)`

GetBfLastDepthOk returns a tuple with the BfLastDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepth

`func (o *V0042StatsMsg) SetBfLastDepth(v int32)`

SetBfLastDepth sets BfLastDepth field to given value.

### HasBfLastDepth

`func (o *V0042StatsMsg) HasBfLastDepth() bool`

HasBfLastDepth returns a boolean if a field has been set.

### GetBfLastDepthTry

`func (o *V0042StatsMsg) GetBfLastDepthTry() int32`

GetBfLastDepthTry returns the BfLastDepthTry field if non-nil, zero value otherwise.

### GetBfLastDepthTryOk

`func (o *V0042StatsMsg) GetBfLastDepthTryOk() (*int32, bool)`

GetBfLastDepthTryOk returns a tuple with the BfLastDepthTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepthTry

`func (o *V0042StatsMsg) SetBfLastDepthTry(v int32)`

SetBfLastDepthTry sets BfLastDepthTry field to given value.

### HasBfLastDepthTry

`func (o *V0042StatsMsg) HasBfLastDepthTry() bool`

HasBfLastDepthTry returns a boolean if a field has been set.

### GetBfDepthSum

`func (o *V0042StatsMsg) GetBfDepthSum() int32`

GetBfDepthSum returns the BfDepthSum field if non-nil, zero value otherwise.

### GetBfDepthSumOk

`func (o *V0042StatsMsg) GetBfDepthSumOk() (*int32, bool)`

GetBfDepthSumOk returns a tuple with the BfDepthSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthSum

`func (o *V0042StatsMsg) SetBfDepthSum(v int32)`

SetBfDepthSum sets BfDepthSum field to given value.

### HasBfDepthSum

`func (o *V0042StatsMsg) HasBfDepthSum() bool`

HasBfDepthSum returns a boolean if a field has been set.

### GetBfDepthTrySum

`func (o *V0042StatsMsg) GetBfDepthTrySum() int32`

GetBfDepthTrySum returns the BfDepthTrySum field if non-nil, zero value otherwise.

### GetBfDepthTrySumOk

`func (o *V0042StatsMsg) GetBfDepthTrySumOk() (*int32, bool)`

GetBfDepthTrySumOk returns a tuple with the BfDepthTrySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthTrySum

`func (o *V0042StatsMsg) SetBfDepthTrySum(v int32)`

SetBfDepthTrySum sets BfDepthTrySum field to given value.

### HasBfDepthTrySum

`func (o *V0042StatsMsg) HasBfDepthTrySum() bool`

HasBfDepthTrySum returns a boolean if a field has been set.

### GetBfQueueLen

`func (o *V0042StatsMsg) GetBfQueueLen() int32`

GetBfQueueLen returns the BfQueueLen field if non-nil, zero value otherwise.

### GetBfQueueLenOk

`func (o *V0042StatsMsg) GetBfQueueLenOk() (*int32, bool)`

GetBfQueueLenOk returns a tuple with the BfQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLen

`func (o *V0042StatsMsg) SetBfQueueLen(v int32)`

SetBfQueueLen sets BfQueueLen field to given value.

### HasBfQueueLen

`func (o *V0042StatsMsg) HasBfQueueLen() bool`

HasBfQueueLen returns a boolean if a field has been set.

### GetBfQueueLenMean

`func (o *V0042StatsMsg) GetBfQueueLenMean() int64`

GetBfQueueLenMean returns the BfQueueLenMean field if non-nil, zero value otherwise.

### GetBfQueueLenMeanOk

`func (o *V0042StatsMsg) GetBfQueueLenMeanOk() (*int64, bool)`

GetBfQueueLenMeanOk returns a tuple with the BfQueueLenMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenMean

`func (o *V0042StatsMsg) SetBfQueueLenMean(v int64)`

SetBfQueueLenMean sets BfQueueLenMean field to given value.

### HasBfQueueLenMean

`func (o *V0042StatsMsg) HasBfQueueLenMean() bool`

HasBfQueueLenMean returns a boolean if a field has been set.

### GetBfQueueLenSum

`func (o *V0042StatsMsg) GetBfQueueLenSum() int32`

GetBfQueueLenSum returns the BfQueueLenSum field if non-nil, zero value otherwise.

### GetBfQueueLenSumOk

`func (o *V0042StatsMsg) GetBfQueueLenSumOk() (*int32, bool)`

GetBfQueueLenSumOk returns a tuple with the BfQueueLenSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenSum

`func (o *V0042StatsMsg) SetBfQueueLenSum(v int32)`

SetBfQueueLenSum sets BfQueueLenSum field to given value.

### HasBfQueueLenSum

`func (o *V0042StatsMsg) HasBfQueueLenSum() bool`

HasBfQueueLenSum returns a boolean if a field has been set.

### GetBfTableSize

`func (o *V0042StatsMsg) GetBfTableSize() int32`

GetBfTableSize returns the BfTableSize field if non-nil, zero value otherwise.

### GetBfTableSizeOk

`func (o *V0042StatsMsg) GetBfTableSizeOk() (*int32, bool)`

GetBfTableSizeOk returns a tuple with the BfTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSize

`func (o *V0042StatsMsg) SetBfTableSize(v int32)`

SetBfTableSize sets BfTableSize field to given value.

### HasBfTableSize

`func (o *V0042StatsMsg) HasBfTableSize() bool`

HasBfTableSize returns a boolean if a field has been set.

### GetBfTableSizeSum

`func (o *V0042StatsMsg) GetBfTableSizeSum() int32`

GetBfTableSizeSum returns the BfTableSizeSum field if non-nil, zero value otherwise.

### GetBfTableSizeSumOk

`func (o *V0042StatsMsg) GetBfTableSizeSumOk() (*int32, bool)`

GetBfTableSizeSumOk returns a tuple with the BfTableSizeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeSum

`func (o *V0042StatsMsg) SetBfTableSizeSum(v int32)`

SetBfTableSizeSum sets BfTableSizeSum field to given value.

### HasBfTableSizeSum

`func (o *V0042StatsMsg) HasBfTableSizeSum() bool`

HasBfTableSizeSum returns a boolean if a field has been set.

### GetBfTableSizeMean

`func (o *V0042StatsMsg) GetBfTableSizeMean() int64`

GetBfTableSizeMean returns the BfTableSizeMean field if non-nil, zero value otherwise.

### GetBfTableSizeMeanOk

`func (o *V0042StatsMsg) GetBfTableSizeMeanOk() (*int64, bool)`

GetBfTableSizeMeanOk returns a tuple with the BfTableSizeMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeMean

`func (o *V0042StatsMsg) SetBfTableSizeMean(v int64)`

SetBfTableSizeMean sets BfTableSizeMean field to given value.

### HasBfTableSizeMean

`func (o *V0042StatsMsg) HasBfTableSizeMean() bool`

HasBfTableSizeMean returns a boolean if a field has been set.

### GetBfWhenLastCycle

`func (o *V0042StatsMsg) GetBfWhenLastCycle() V0042Uint64NoValStruct`

GetBfWhenLastCycle returns the BfWhenLastCycle field if non-nil, zero value otherwise.

### GetBfWhenLastCycleOk

`func (o *V0042StatsMsg) GetBfWhenLastCycleOk() (*V0042Uint64NoValStruct, bool)`

GetBfWhenLastCycleOk returns a tuple with the BfWhenLastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfWhenLastCycle

`func (o *V0042StatsMsg) SetBfWhenLastCycle(v V0042Uint64NoValStruct)`

SetBfWhenLastCycle sets BfWhenLastCycle field to given value.

### HasBfWhenLastCycle

`func (o *V0042StatsMsg) HasBfWhenLastCycle() bool`

HasBfWhenLastCycle returns a boolean if a field has been set.

### GetBfActive

`func (o *V0042StatsMsg) GetBfActive() bool`

GetBfActive returns the BfActive field if non-nil, zero value otherwise.

### GetBfActiveOk

`func (o *V0042StatsMsg) GetBfActiveOk() (*bool, bool)`

GetBfActiveOk returns a tuple with the BfActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfActive

`func (o *V0042StatsMsg) SetBfActive(v bool)`

SetBfActive sets BfActive field to given value.

### HasBfActive

`func (o *V0042StatsMsg) HasBfActive() bool`

HasBfActive returns a boolean if a field has been set.

### GetRpcsByMessageType

`func (o *V0042StatsMsg) GetRpcsByMessageType() []V0042StatsMsgRpcType`

GetRpcsByMessageType returns the RpcsByMessageType field if non-nil, zero value otherwise.

### GetRpcsByMessageTypeOk

`func (o *V0042StatsMsg) GetRpcsByMessageTypeOk() (*[]V0042StatsMsgRpcType, bool)`

GetRpcsByMessageTypeOk returns a tuple with the RpcsByMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByMessageType

`func (o *V0042StatsMsg) SetRpcsByMessageType(v []V0042StatsMsgRpcType)`

SetRpcsByMessageType sets RpcsByMessageType field to given value.

### HasRpcsByMessageType

`func (o *V0042StatsMsg) HasRpcsByMessageType() bool`

HasRpcsByMessageType returns a boolean if a field has been set.

### GetRpcsByUser

`func (o *V0042StatsMsg) GetRpcsByUser() []V0042StatsMsgRpcUser`

GetRpcsByUser returns the RpcsByUser field if non-nil, zero value otherwise.

### GetRpcsByUserOk

`func (o *V0042StatsMsg) GetRpcsByUserOk() (*[]V0042StatsMsgRpcUser, bool)`

GetRpcsByUserOk returns a tuple with the RpcsByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByUser

`func (o *V0042StatsMsg) SetRpcsByUser(v []V0042StatsMsgRpcUser)`

SetRpcsByUser sets RpcsByUser field to given value.

### HasRpcsByUser

`func (o *V0042StatsMsg) HasRpcsByUser() bool`

HasRpcsByUser returns a boolean if a field has been set.

### GetPendingRpcs

`func (o *V0042StatsMsg) GetPendingRpcs() []V0042StatsMsgRpcQueue`

GetPendingRpcs returns the PendingRpcs field if non-nil, zero value otherwise.

### GetPendingRpcsOk

`func (o *V0042StatsMsg) GetPendingRpcsOk() (*[]V0042StatsMsgRpcQueue, bool)`

GetPendingRpcsOk returns a tuple with the PendingRpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcs

`func (o *V0042StatsMsg) SetPendingRpcs(v []V0042StatsMsgRpcQueue)`

SetPendingRpcs sets PendingRpcs field to given value.

### HasPendingRpcs

`func (o *V0042StatsMsg) HasPendingRpcs() bool`

HasPendingRpcs returns a boolean if a field has been set.

### GetPendingRpcsByHostlist

`func (o *V0042StatsMsg) GetPendingRpcsByHostlist() []V0042StatsMsgRpcDump`

GetPendingRpcsByHostlist returns the PendingRpcsByHostlist field if non-nil, zero value otherwise.

### GetPendingRpcsByHostlistOk

`func (o *V0042StatsMsg) GetPendingRpcsByHostlistOk() (*[]V0042StatsMsgRpcDump, bool)`

GetPendingRpcsByHostlistOk returns a tuple with the PendingRpcsByHostlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRpcsByHostlist

`func (o *V0042StatsMsg) SetPendingRpcsByHostlist(v []V0042StatsMsgRpcDump)`

SetPendingRpcsByHostlist sets PendingRpcsByHostlist field to given value.

### HasPendingRpcsByHostlist

`func (o *V0042StatsMsg) HasPendingRpcsByHostlist() bool`

HasPendingRpcsByHostlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# V0042Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**V0042StepTime**](V0042StepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0042ProcessExitCodeVerbose**](V0042ProcessExitCodeVerbose.md) |  | [optional] 
**Nodes** | Pointer to [**V0042StepNodes**](V0042StepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**V0040StepTasks**](V0040StepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | Deprecated; Process ID | [optional] 
**CPU** | Pointer to [**V0042StepCPU**](V0042StepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User ID that requested termination of the step | [optional] 
**State** | Pointer to **[]string** |  | [optional] 
**Statistics** | Pointer to [**V0042StepStatistics**](V0042StepStatistics.md) |  | [optional] 
**Step** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep.md) |  | [optional] 
**Task** | Pointer to [**V0040StepTask**](V0040StepTask.md) |  | [optional] 
**Tres** | Pointer to [**V0042StepTres**](V0042StepTres.md) |  | [optional] 

## Methods

### NewV0042Step

`func NewV0042Step() *V0042Step`

NewV0042Step instantiates a new V0042Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042StepWithDefaults

`func NewV0042StepWithDefaults() *V0042Step`

NewV0042StepWithDefaults instantiates a new V0042Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V0042Step) GetTime() V0042StepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0042Step) GetTimeOk() (*V0042StepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0042Step) SetTime(v V0042StepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0042Step) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0042Step) GetExitCode() V0042ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0042Step) GetExitCodeOk() (*V0042ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0042Step) SetExitCode(v V0042ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0042Step) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *V0042Step) GetNodes() V0042StepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0042Step) GetNodesOk() (*V0042StepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0042Step) SetNodes(v V0042StepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0042Step) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *V0042Step) GetTasks() V0040StepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0042Step) GetTasksOk() (*V0040StepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0042Step) SetTasks(v V0040StepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0042Step) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *V0042Step) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *V0042Step) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *V0042Step) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *V0042Step) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *V0042Step) GetCPU() V0042StepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0042Step) GetCPUOk() (*V0042StepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0042Step) SetCPU(v V0042StepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0042Step) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0042Step) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0042Step) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0042Step) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0042Step) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *V0042Step) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0042Step) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0042Step) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0042Step) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *V0042Step) GetStatistics() V0042StepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0042Step) GetStatisticsOk() (*V0042StepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0042Step) SetStatistics(v V0042StepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0042Step) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *V0042Step) GetStep() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V0042Step) GetStepOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V0042Step) SetStep(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *V0042Step) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *V0042Step) GetTask() V0040StepTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0042Step) GetTaskOk() (*V0040StepTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0042Step) SetTask(v V0040StepTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0042Step) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *V0042Step) GetTres() V0042StepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0042Step) GetTresOk() (*V0042StepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0042Step) SetTres(v V0042StepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0042Step) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



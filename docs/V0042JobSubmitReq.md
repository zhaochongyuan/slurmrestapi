# V0042JobSubmitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | Deprecated; Populate script field in jobs[0] or job | [optional] 
**Jobs** | Pointer to [**[]V0042JobDescMsg**](V0042JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0042JobDescMsg**](V0042JobDescMsg.md) |  | [optional] 

## Methods

### NewV0042JobSubmitReq

`func NewV0042JobSubmitReq() *V0042JobSubmitReq`

NewV0042JobSubmitReq instantiates a new V0042JobSubmitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobSubmitReqWithDefaults

`func NewV0042JobSubmitReqWithDefaults() *V0042JobSubmitReq`

NewV0042JobSubmitReqWithDefaults instantiates a new V0042JobSubmitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0042JobSubmitReq) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0042JobSubmitReq) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0042JobSubmitReq) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0042JobSubmitReq) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJobs

`func (o *V0042JobSubmitReq) GetJobs() []V0042JobDescMsg`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0042JobSubmitReq) GetJobsOk() (*[]V0042JobDescMsg, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0042JobSubmitReq) SetJobs(v []V0042JobDescMsg)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0042JobSubmitReq) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetJob

`func (o *V0042JobSubmitReq) GetJob() V0042JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0042JobSubmitReq) GetJobOk() (*V0042JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0042JobSubmitReq) SetJob(v V0042JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0042JobSubmitReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# V0041OpenapiInstancesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]V0041OpenapiSlurmdbdConfigRespInstancesInner**](V0041OpenapiSlurmdbdConfigRespInstancesInner.md) | instances | 
**Meta** | Pointer to [**V0041OpenapiSharesRespMeta**](V0041OpenapiSharesRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSharesRespErrorsInner**](V0041OpenapiSharesRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSharesRespWarningsInner**](V0041OpenapiSharesRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiInstancesResp

`func NewV0041OpenapiInstancesResp(instances []V0041OpenapiSlurmdbdConfigRespInstancesInner, ) *V0041OpenapiInstancesResp`

NewV0041OpenapiInstancesResp instantiates a new V0041OpenapiInstancesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiInstancesRespWithDefaults

`func NewV0041OpenapiInstancesRespWithDefaults() *V0041OpenapiInstancesResp`

NewV0041OpenapiInstancesRespWithDefaults instantiates a new V0041OpenapiInstancesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *V0041OpenapiInstancesResp) GetInstances() []V0041OpenapiSlurmdbdConfigRespInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0041OpenapiInstancesResp) GetInstancesOk() (*[]V0041OpenapiSlurmdbdConfigRespInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0041OpenapiInstancesResp) SetInstances(v []V0041OpenapiSlurmdbdConfigRespInstancesInner)`

SetInstances sets Instances field to given value.


### GetMeta

`func (o *V0041OpenapiInstancesResp) GetMeta() V0041OpenapiSharesRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiInstancesResp) GetMetaOk() (*V0041OpenapiSharesRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiInstancesResp) SetMeta(v V0041OpenapiSharesRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiInstancesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiInstancesResp) GetErrors() []V0041OpenapiSharesRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiInstancesResp) GetErrorsOk() (*[]V0041OpenapiSharesRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiInstancesResp) SetErrors(v []V0041OpenapiSharesRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiInstancesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiInstancesResp) GetWarnings() []V0041OpenapiSharesRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiInstancesResp) GetWarningsOk() (*[]V0041OpenapiSharesRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiInstancesResp) SetWarnings(v []V0041OpenapiSharesRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiInstancesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



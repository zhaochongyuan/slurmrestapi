# V0042OpenapiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Long form error description | [optional] 
**ErrorNumber** | Pointer to **int32** | Slurm numeric error identifier | [optional] 
**Error** | Pointer to **string** | Short form error description | [optional] 
**Source** | Pointer to **string** | Source of error or where error was first detected | [optional] 

## Methods

### NewV0042OpenapiError

`func NewV0042OpenapiError() *V0042OpenapiError`

NewV0042OpenapiError instantiates a new V0042OpenapiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042OpenapiErrorWithDefaults

`func NewV0042OpenapiErrorWithDefaults() *V0042OpenapiError`

NewV0042OpenapiErrorWithDefaults instantiates a new V0042OpenapiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V0042OpenapiError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0042OpenapiError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0042OpenapiError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V0042OpenapiError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorNumber

`func (o *V0042OpenapiError) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *V0042OpenapiError) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *V0042OpenapiError) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *V0042OpenapiError) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.

### GetError

`func (o *V0042OpenapiError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0042OpenapiError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0042OpenapiError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0042OpenapiError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSource

`func (o *V0042OpenapiError) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V0042OpenapiError) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V0042OpenapiError) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V0042OpenapiError) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# V0042License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseName** | Pointer to **string** | Name of the license | [optional] 
**Total** | Pointer to **int32** | Total number of licenses present | [optional] 
**Used** | Pointer to **int32** | Number of licenses in use | [optional] 
**Free** | Pointer to **int32** | Number of licenses currently available | [optional] 
**Remote** | Pointer to **bool** | Indicates whether licenses are served by the database | [optional] 
**Reserved** | Pointer to **int32** | Number of licenses reserved | [optional] 
**LastConsumed** | Pointer to **int32** | Last known number of licenses that were consumed in the license manager (Remote Only) | [optional] 
**LastDeficit** | Pointer to **int32** | Number of \&quot;missing licenses\&quot; from the cluster&#39;s perspective | [optional] 
**LastUpdate** | Pointer to **int64** | When the license information was last updated (UNIX Timestamp) | [optional] 

## Methods

### NewV0042License

`func NewV0042License() *V0042License`

NewV0042License instantiates a new V0042License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042LicenseWithDefaults

`func NewV0042LicenseWithDefaults() *V0042License`

NewV0042LicenseWithDefaults instantiates a new V0042License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseName

`func (o *V0042License) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *V0042License) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *V0042License) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *V0042License) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetTotal

`func (o *V0042License) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0042License) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0042License) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0042License) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *V0042License) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *V0042License) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *V0042License) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *V0042License) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetFree

`func (o *V0042License) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *V0042License) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *V0042License) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *V0042License) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetRemote

`func (o *V0042License) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *V0042License) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *V0042License) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *V0042License) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetReserved

`func (o *V0042License) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *V0042License) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *V0042License) SetReserved(v int32)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *V0042License) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetLastConsumed

`func (o *V0042License) GetLastConsumed() int32`

GetLastConsumed returns the LastConsumed field if non-nil, zero value otherwise.

### GetLastConsumedOk

`func (o *V0042License) GetLastConsumedOk() (*int32, bool)`

GetLastConsumedOk returns a tuple with the LastConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumed

`func (o *V0042License) SetLastConsumed(v int32)`

SetLastConsumed sets LastConsumed field to given value.

### HasLastConsumed

`func (o *V0042License) HasLastConsumed() bool`

HasLastConsumed returns a boolean if a field has been set.

### GetLastDeficit

`func (o *V0042License) GetLastDeficit() int32`

GetLastDeficit returns the LastDeficit field if non-nil, zero value otherwise.

### GetLastDeficitOk

`func (o *V0042License) GetLastDeficitOk() (*int32, bool)`

GetLastDeficitOk returns a tuple with the LastDeficit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeficit

`func (o *V0042License) SetLastDeficit(v int32)`

SetLastDeficit sets LastDeficit field to given value.

### HasLastDeficit

`func (o *V0042License) HasLastDeficit() bool`

HasLastDeficit returns a boolean if a field has been set.

### GetLastUpdate

`func (o *V0042License) GetLastUpdate() int64`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0042License) GetLastUpdateOk() (*int64, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0042License) SetLastUpdate(v int64)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *V0042License) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



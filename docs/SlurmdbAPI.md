# \SlurmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbV0040DeleteAccount**](SlurmdbAPI.md#SlurmdbV0040DeleteAccount) | **Delete** /slurmdb/v0.0.40/account/{account_name} | Delete account
[**SlurmdbV0040DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0040DeleteAssociation) | **Delete** /slurmdb/v0.0.40/association/ | Delete association
[**SlurmdbV0040DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0040DeleteAssociations) | **Delete** /slurmdb/v0.0.40/associations/ | Delete associations
[**SlurmdbV0040DeleteCluster**](SlurmdbAPI.md#SlurmdbV0040DeleteCluster) | **Delete** /slurmdb/v0.0.40/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0040DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0040DeleteSingleQos) | **Delete** /slurmdb/v0.0.40/qos/{qos} | Delete QOS
[**SlurmdbV0040DeleteUser**](SlurmdbAPI.md#SlurmdbV0040DeleteUser) | **Delete** /slurmdb/v0.0.40/user/{name} | Delete user
[**SlurmdbV0040DeleteWckey**](SlurmdbAPI.md#SlurmdbV0040DeleteWckey) | **Delete** /slurmdb/v0.0.40/wckey/{id} | Delete wckey
[**SlurmdbV0040GetAccount**](SlurmdbAPI.md#SlurmdbV0040GetAccount) | **Get** /slurmdb/v0.0.40/account/{account_name} | Get account info
[**SlurmdbV0040GetAccounts**](SlurmdbAPI.md#SlurmdbV0040GetAccounts) | **Get** /slurmdb/v0.0.40/accounts/ | Get account list
[**SlurmdbV0040GetAssociation**](SlurmdbAPI.md#SlurmdbV0040GetAssociation) | **Get** /slurmdb/v0.0.40/association/ | Get association info
[**SlurmdbV0040GetAssociations**](SlurmdbAPI.md#SlurmdbV0040GetAssociations) | **Get** /slurmdb/v0.0.40/associations/ | Get association list
[**SlurmdbV0040GetCluster**](SlurmdbAPI.md#SlurmdbV0040GetCluster) | **Get** /slurmdb/v0.0.40/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0040GetClusters**](SlurmdbAPI.md#SlurmdbV0040GetClusters) | **Get** /slurmdb/v0.0.40/clusters/ | Get cluster list
[**SlurmdbV0040GetConfig**](SlurmdbAPI.md#SlurmdbV0040GetConfig) | **Get** /slurmdb/v0.0.40/config | Dump all configuration information
[**SlurmdbV0040GetDiag**](SlurmdbAPI.md#SlurmdbV0040GetDiag) | **Get** /slurmdb/v0.0.40/diag/ | Get slurmdb diagnostics
[**SlurmdbV0040GetInstance**](SlurmdbAPI.md#SlurmdbV0040GetInstance) | **Get** /slurmdb/v0.0.40/instance/ | Get instance info
[**SlurmdbV0040GetInstances**](SlurmdbAPI.md#SlurmdbV0040GetInstances) | **Get** /slurmdb/v0.0.40/instances/ | Get instance list
[**SlurmdbV0040GetJob**](SlurmdbAPI.md#SlurmdbV0040GetJob) | **Get** /slurmdb/v0.0.40/job/{job_id} | Get job info
[**SlurmdbV0040GetJobs**](SlurmdbAPI.md#SlurmdbV0040GetJobs) | **Get** /slurmdb/v0.0.40/jobs/ | Get job list
[**SlurmdbV0040GetQos**](SlurmdbAPI.md#SlurmdbV0040GetQos) | **Get** /slurmdb/v0.0.40/qos/ | Get QOS list
[**SlurmdbV0040GetSingleQos**](SlurmdbAPI.md#SlurmdbV0040GetSingleQos) | **Get** /slurmdb/v0.0.40/qos/{qos} | Get QOS info
[**SlurmdbV0040GetTres**](SlurmdbAPI.md#SlurmdbV0040GetTres) | **Get** /slurmdb/v0.0.40/tres/ | Get TRES info
[**SlurmdbV0040GetUser**](SlurmdbAPI.md#SlurmdbV0040GetUser) | **Get** /slurmdb/v0.0.40/user/{name} | Get user info
[**SlurmdbV0040GetUsers**](SlurmdbAPI.md#SlurmdbV0040GetUsers) | **Get** /slurmdb/v0.0.40/users/ | Get user list
[**SlurmdbV0040GetWckey**](SlurmdbAPI.md#SlurmdbV0040GetWckey) | **Get** /slurmdb/v0.0.40/wckey/{id} | Get wckey info
[**SlurmdbV0040GetWckeys**](SlurmdbAPI.md#SlurmdbV0040GetWckeys) | **Get** /slurmdb/v0.0.40/wckeys/ | Get wckey list
[**SlurmdbV0040PostAccounts**](SlurmdbAPI.md#SlurmdbV0040PostAccounts) | **Post** /slurmdb/v0.0.40/accounts/ | Add/update list of accounts
[**SlurmdbV0040PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0040PostAccountsAssociation) | **Post** /slurmdb/v0.0.40/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0040PostAssociations**](SlurmdbAPI.md#SlurmdbV0040PostAssociations) | **Post** /slurmdb/v0.0.40/associations/ | Set associations info
[**SlurmdbV0040PostClusters**](SlurmdbAPI.md#SlurmdbV0040PostClusters) | **Post** /slurmdb/v0.0.40/clusters/ | Get cluster list
[**SlurmdbV0040PostConfig**](SlurmdbAPI.md#SlurmdbV0040PostConfig) | **Post** /slurmdb/v0.0.40/config | Load all configuration information
[**SlurmdbV0040PostQos**](SlurmdbAPI.md#SlurmdbV0040PostQos) | **Post** /slurmdb/v0.0.40/qos/ | Add or update QOSs
[**SlurmdbV0040PostTres**](SlurmdbAPI.md#SlurmdbV0040PostTres) | **Post** /slurmdb/v0.0.40/tres/ | Add TRES
[**SlurmdbV0040PostUsers**](SlurmdbAPI.md#SlurmdbV0040PostUsers) | **Post** /slurmdb/v0.0.40/users/ | Update users
[**SlurmdbV0040PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0040PostUsersAssociation) | **Post** /slurmdb/v0.0.40/users_association/ | Add users with conditional association
[**SlurmdbV0040PostWckeys**](SlurmdbAPI.md#SlurmdbV0040PostWckeys) | **Post** /slurmdb/v0.0.40/wckeys/ | Add or update wckeys
[**SlurmdbV0041DeleteAccount**](SlurmdbAPI.md#SlurmdbV0041DeleteAccount) | **Delete** /slurmdb/v0.0.41/account/{account_name} | Delete account
[**SlurmdbV0041DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociation) | **Delete** /slurmdb/v0.0.41/association/ | Delete association
[**SlurmdbV0041DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociations) | **Delete** /slurmdb/v0.0.41/associations/ | Delete associations
[**SlurmdbV0041DeleteCluster**](SlurmdbAPI.md#SlurmdbV0041DeleteCluster) | **Delete** /slurmdb/v0.0.41/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0041DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0041DeleteSingleQos) | **Delete** /slurmdb/v0.0.41/qos/{qos} | Delete QOS
[**SlurmdbV0041DeleteUser**](SlurmdbAPI.md#SlurmdbV0041DeleteUser) | **Delete** /slurmdb/v0.0.41/user/{name} | Delete user
[**SlurmdbV0041DeleteWckey**](SlurmdbAPI.md#SlurmdbV0041DeleteWckey) | **Delete** /slurmdb/v0.0.41/wckey/{id} | Delete wckey
[**SlurmdbV0041GetAccount**](SlurmdbAPI.md#SlurmdbV0041GetAccount) | **Get** /slurmdb/v0.0.41/account/{account_name} | Get account info
[**SlurmdbV0041GetAccounts**](SlurmdbAPI.md#SlurmdbV0041GetAccounts) | **Get** /slurmdb/v0.0.41/accounts/ | Get account list
[**SlurmdbV0041GetAssociation**](SlurmdbAPI.md#SlurmdbV0041GetAssociation) | **Get** /slurmdb/v0.0.41/association/ | Get association info
[**SlurmdbV0041GetAssociations**](SlurmdbAPI.md#SlurmdbV0041GetAssociations) | **Get** /slurmdb/v0.0.41/associations/ | Get association list
[**SlurmdbV0041GetCluster**](SlurmdbAPI.md#SlurmdbV0041GetCluster) | **Get** /slurmdb/v0.0.41/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0041GetClusters**](SlurmdbAPI.md#SlurmdbV0041GetClusters) | **Get** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041GetConfig**](SlurmdbAPI.md#SlurmdbV0041GetConfig) | **Get** /slurmdb/v0.0.41/config | Dump all configuration information
[**SlurmdbV0041GetDiag**](SlurmdbAPI.md#SlurmdbV0041GetDiag) | **Get** /slurmdb/v0.0.41/diag/ | Get slurmdb diagnostics
[**SlurmdbV0041GetInstance**](SlurmdbAPI.md#SlurmdbV0041GetInstance) | **Get** /slurmdb/v0.0.41/instance/ | Get instance info
[**SlurmdbV0041GetInstances**](SlurmdbAPI.md#SlurmdbV0041GetInstances) | **Get** /slurmdb/v0.0.41/instances/ | Get instance list
[**SlurmdbV0041GetJob**](SlurmdbAPI.md#SlurmdbV0041GetJob) | **Get** /slurmdb/v0.0.41/job/{job_id} | Get job info
[**SlurmdbV0041GetJobs**](SlurmdbAPI.md#SlurmdbV0041GetJobs) | **Get** /slurmdb/v0.0.41/jobs/ | Get job list
[**SlurmdbV0041GetQos**](SlurmdbAPI.md#SlurmdbV0041GetQos) | **Get** /slurmdb/v0.0.41/qos/ | Get QOS list
[**SlurmdbV0041GetSingleQos**](SlurmdbAPI.md#SlurmdbV0041GetSingleQos) | **Get** /slurmdb/v0.0.41/qos/{qos} | Get QOS info
[**SlurmdbV0041GetTres**](SlurmdbAPI.md#SlurmdbV0041GetTres) | **Get** /slurmdb/v0.0.41/tres/ | Get TRES info
[**SlurmdbV0041GetUser**](SlurmdbAPI.md#SlurmdbV0041GetUser) | **Get** /slurmdb/v0.0.41/user/{name} | Get user info
[**SlurmdbV0041GetUsers**](SlurmdbAPI.md#SlurmdbV0041GetUsers) | **Get** /slurmdb/v0.0.41/users/ | Get user list
[**SlurmdbV0041GetWckey**](SlurmdbAPI.md#SlurmdbV0041GetWckey) | **Get** /slurmdb/v0.0.41/wckey/{id} | Get wckey info
[**SlurmdbV0041GetWckeys**](SlurmdbAPI.md#SlurmdbV0041GetWckeys) | **Get** /slurmdb/v0.0.41/wckeys/ | Get wckey list
[**SlurmdbV0041PostAccounts**](SlurmdbAPI.md#SlurmdbV0041PostAccounts) | **Post** /slurmdb/v0.0.41/accounts/ | Add/update list of accounts
[**SlurmdbV0041PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0041PostAccountsAssociation) | **Post** /slurmdb/v0.0.41/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0041PostAssociations**](SlurmdbAPI.md#SlurmdbV0041PostAssociations) | **Post** /slurmdb/v0.0.41/associations/ | Set associations info
[**SlurmdbV0041PostClusters**](SlurmdbAPI.md#SlurmdbV0041PostClusters) | **Post** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041PostConfig**](SlurmdbAPI.md#SlurmdbV0041PostConfig) | **Post** /slurmdb/v0.0.41/config | Load all configuration information
[**SlurmdbV0041PostQos**](SlurmdbAPI.md#SlurmdbV0041PostQos) | **Post** /slurmdb/v0.0.41/qos/ | Add or update QOSs
[**SlurmdbV0041PostTres**](SlurmdbAPI.md#SlurmdbV0041PostTres) | **Post** /slurmdb/v0.0.41/tres/ | Add TRES
[**SlurmdbV0041PostUsers**](SlurmdbAPI.md#SlurmdbV0041PostUsers) | **Post** /slurmdb/v0.0.41/users/ | Update users
[**SlurmdbV0041PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0041PostUsersAssociation) | **Post** /slurmdb/v0.0.41/users_association/ | Add users with conditional association
[**SlurmdbV0041PostWckeys**](SlurmdbAPI.md#SlurmdbV0041PostWckeys) | **Post** /slurmdb/v0.0.41/wckeys/ | Add or update wckeys
[**SlurmdbV0042DeleteAccount**](SlurmdbAPI.md#SlurmdbV0042DeleteAccount) | **Delete** /slurmdb/v0.0.42/account/{account_name} | Delete account
[**SlurmdbV0042DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0042DeleteAssociation) | **Delete** /slurmdb/v0.0.42/association/ | Delete association
[**SlurmdbV0042DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0042DeleteAssociations) | **Delete** /slurmdb/v0.0.42/associations/ | Delete associations
[**SlurmdbV0042DeleteCluster**](SlurmdbAPI.md#SlurmdbV0042DeleteCluster) | **Delete** /slurmdb/v0.0.42/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0042DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0042DeleteSingleQos) | **Delete** /slurmdb/v0.0.42/qos/{qos} | Delete QOS
[**SlurmdbV0042DeleteUser**](SlurmdbAPI.md#SlurmdbV0042DeleteUser) | **Delete** /slurmdb/v0.0.42/user/{name} | Delete user
[**SlurmdbV0042DeleteWckey**](SlurmdbAPI.md#SlurmdbV0042DeleteWckey) | **Delete** /slurmdb/v0.0.42/wckey/{id} | Delete wckey
[**SlurmdbV0042GetAccount**](SlurmdbAPI.md#SlurmdbV0042GetAccount) | **Get** /slurmdb/v0.0.42/account/{account_name} | Get account info
[**SlurmdbV0042GetAccounts**](SlurmdbAPI.md#SlurmdbV0042GetAccounts) | **Get** /slurmdb/v0.0.42/accounts/ | Get account list
[**SlurmdbV0042GetAssociation**](SlurmdbAPI.md#SlurmdbV0042GetAssociation) | **Get** /slurmdb/v0.0.42/association/ | Get association info
[**SlurmdbV0042GetAssociations**](SlurmdbAPI.md#SlurmdbV0042GetAssociations) | **Get** /slurmdb/v0.0.42/associations/ | Get association list
[**SlurmdbV0042GetCluster**](SlurmdbAPI.md#SlurmdbV0042GetCluster) | **Get** /slurmdb/v0.0.42/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0042GetClusters**](SlurmdbAPI.md#SlurmdbV0042GetClusters) | **Get** /slurmdb/v0.0.42/clusters/ | Get cluster list
[**SlurmdbV0042GetConfig**](SlurmdbAPI.md#SlurmdbV0042GetConfig) | **Get** /slurmdb/v0.0.42/config | Dump all configuration information
[**SlurmdbV0042GetDiag**](SlurmdbAPI.md#SlurmdbV0042GetDiag) | **Get** /slurmdb/v0.0.42/diag/ | Get slurmdb diagnostics
[**SlurmdbV0042GetInstance**](SlurmdbAPI.md#SlurmdbV0042GetInstance) | **Get** /slurmdb/v0.0.42/instance/ | Get instance info
[**SlurmdbV0042GetInstances**](SlurmdbAPI.md#SlurmdbV0042GetInstances) | **Get** /slurmdb/v0.0.42/instances/ | Get instance list
[**SlurmdbV0042GetJob**](SlurmdbAPI.md#SlurmdbV0042GetJob) | **Get** /slurmdb/v0.0.42/job/{job_id} | Get job info
[**SlurmdbV0042GetJobs**](SlurmdbAPI.md#SlurmdbV0042GetJobs) | **Get** /slurmdb/v0.0.42/jobs/ | Get job list
[**SlurmdbV0042GetPing**](SlurmdbAPI.md#SlurmdbV0042GetPing) | **Get** /slurmdb/v0.0.42/ping/ | ping test
[**SlurmdbV0042GetQos**](SlurmdbAPI.md#SlurmdbV0042GetQos) | **Get** /slurmdb/v0.0.42/qos/ | Get QOS list
[**SlurmdbV0042GetSingleQos**](SlurmdbAPI.md#SlurmdbV0042GetSingleQos) | **Get** /slurmdb/v0.0.42/qos/{qos} | Get QOS info
[**SlurmdbV0042GetTres**](SlurmdbAPI.md#SlurmdbV0042GetTres) | **Get** /slurmdb/v0.0.42/tres/ | Get TRES info
[**SlurmdbV0042GetUser**](SlurmdbAPI.md#SlurmdbV0042GetUser) | **Get** /slurmdb/v0.0.42/user/{name} | Get user info
[**SlurmdbV0042GetUsers**](SlurmdbAPI.md#SlurmdbV0042GetUsers) | **Get** /slurmdb/v0.0.42/users/ | Get user list
[**SlurmdbV0042GetWckey**](SlurmdbAPI.md#SlurmdbV0042GetWckey) | **Get** /slurmdb/v0.0.42/wckey/{id} | Get wckey info
[**SlurmdbV0042GetWckeys**](SlurmdbAPI.md#SlurmdbV0042GetWckeys) | **Get** /slurmdb/v0.0.42/wckeys/ | Get wckey list
[**SlurmdbV0042PostAccounts**](SlurmdbAPI.md#SlurmdbV0042PostAccounts) | **Post** /slurmdb/v0.0.42/accounts/ | Add/update list of accounts
[**SlurmdbV0042PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0042PostAccountsAssociation) | **Post** /slurmdb/v0.0.42/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0042PostAssociations**](SlurmdbAPI.md#SlurmdbV0042PostAssociations) | **Post** /slurmdb/v0.0.42/associations/ | Set associations info
[**SlurmdbV0042PostClusters**](SlurmdbAPI.md#SlurmdbV0042PostClusters) | **Post** /slurmdb/v0.0.42/clusters/ | Get cluster list
[**SlurmdbV0042PostConfig**](SlurmdbAPI.md#SlurmdbV0042PostConfig) | **Post** /slurmdb/v0.0.42/config | Load all configuration information
[**SlurmdbV0042PostQos**](SlurmdbAPI.md#SlurmdbV0042PostQos) | **Post** /slurmdb/v0.0.42/qos/ | Add or update QOSs
[**SlurmdbV0042PostTres**](SlurmdbAPI.md#SlurmdbV0042PostTres) | **Post** /slurmdb/v0.0.42/tres/ | Add TRES
[**SlurmdbV0042PostUsers**](SlurmdbAPI.md#SlurmdbV0042PostUsers) | **Post** /slurmdb/v0.0.42/users/ | Update users
[**SlurmdbV0042PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0042PostUsersAssociation) | **Post** /slurmdb/v0.0.42/users_association/ | Add users with conditional association
[**SlurmdbV0042PostWckeys**](SlurmdbAPI.md#SlurmdbV0042PostWckeys) | **Post** /slurmdb/v0.0.42/wckeys/ | Add or update wckeys



## SlurmdbV0040DeleteAccount

> V0040OpenapiAccountsRemovedResp SlurmdbV0040DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAccount`: V0040OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiAccountsRemovedResp**](V0040OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteAssociation

> V0040OpenapiAssocsRemovedResp SlurmdbV0040DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAssociation`: V0040OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information also | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0040OpenapiAssocsRemovedResp**](V0040OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteAssociations

> V0040OpenapiAssocsRemovedResp SlurmdbV0040DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAssociations`: V0040OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information also | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0040OpenapiAssocsRemovedResp**](V0040OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteCluster

> V0040OpenapiClustersRemovedResp SlurmdbV0040DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteCluster`: V0040OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0040OpenapiClustersRemovedResp**](V0040OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteSingleQos

> V0040OpenapiSlurmdbdQosRemovedResp SlurmdbV0040DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteSingleQos`: V0040OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiSlurmdbdQosRemovedResp**](V0040OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteUser

> V0040OpenapiResp SlurmdbV0040DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteUser`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteWckey

> V0040OpenapiWckeyRemovedResp SlurmdbV0040DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteWckey`: V0040OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiWckeyRemovedResp**](V0040OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAccount

> V0040OpenapiAccountsResp SlurmdbV0040GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAccount`: V0040OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAccounts

> V0040OpenapiAccountsResp SlurmdbV0040GetAccounts(ctx).Description(description).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted accounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAccounts(context.Background()).Description(description).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAccounts`: V0040OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted accounts | 

### Return type

[**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAssociation

> V0040OpenapiAssocsResp SlurmdbV0040GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAssociation`: V0040OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information also | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAssociations

> V0040OpenapiAssocsResp SlurmdbV0040GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAssociations`: V0040OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information also | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetCluster

> V0040OpenapiClustersResp SlurmdbV0040GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetCluster`: V0040OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetClusters

> V0040OpenapiClustersResp SlurmdbV0040GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetClusters`: V0040OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetConfig

> V0040OpenapiSlurmdbdConfigResp SlurmdbV0040GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetConfig`: V0040OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetConfigRequest struct via the builder pattern


### Return type

[**V0040OpenapiSlurmdbdConfigResp**](V0040OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetDiag

> V0040OpenapiSlurmdbdStatsResp SlurmdbV0040GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetDiag`: V0040OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetDiagRequest struct via the builder pattern


### Return type

[**V0040OpenapiSlurmdbdStatsResp**](V0040OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetInstance

> V0040OpenapiInstancesResp SlurmdbV0040GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetInstance`: V0040OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0040OpenapiInstancesResp**](V0040OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetInstances

> V0040OpenapiInstancesResp SlurmdbV0040GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetInstances`: V0040OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0040OpenapiInstancesResp**](V0040OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetJob

> V0040OpenapiSlurmdbdJobsResp SlurmdbV0040GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	jobId := "jobId_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetJob`: V0040OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiSlurmdbdJobsResp**](V0040OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetJobs

> V0040OpenapiSlurmdbdJobsResp SlurmdbV0040GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).TimelimitMax(timelimitMax).TimelimitMin(timelimitMin).EndTime(endTime).StartTime(startTime).SubmitTime(submitTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	cpusMax := "cpusMax_example" // string | Maximum number of cpus (optional)
	cpusMin := "cpusMin_example" // string | Minimum number of cpus (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	nodesMax := "nodesMax_example" // string | Maximum number of nodes (optional)
	nodesMin := "nodesMin_example" // string | Minimum number of nodes (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	timelimitMax := "timelimitMax_example" // string | Maximum timelimit (seconds) (optional)
	timelimitMin := "timelimitMin_example" // string | Minimum timelimit (seconds) (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	submitTime := "submitTime_example" // string | Submit time (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV wckey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).TimelimitMax(timelimitMax).TimelimitMin(timelimitMin).EndTime(endTime).StartTime(startTime).SubmitTime(submitTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetJobs`: V0040OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **cpusMax** | **string** | Maximum number of cpus | 
 **cpusMin** | **string** | Minimum number of cpus | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **nodesMax** | **string** | Maximum number of nodes | 
 **nodesMin** | **string** | Minimum number of nodes | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **timelimitMax** | **string** | Maximum timelimit (seconds) | 
 **timelimitMin** | **string** | Minimum timelimit (seconds) | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **submitTime** | **string** | Submit time (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV wckey list | 

### Return type

[**V0040OpenapiSlurmdbdJobsResp**](V0040OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetQos

> V0040OpenapiSlurmdbdQosResp SlurmdbV0040GetQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetQos`: V0040OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 

### Return type

[**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetSingleQos

> V0040OpenapiSlurmdbdQosResp SlurmdbV0040GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetSingleQos`: V0040OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetTres

> V0040OpenapiTresResp SlurmdbV0040GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetTres`: V0040OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetTresRequest struct via the builder pattern


### Return type

[**V0040OpenapiTresResp**](V0040OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetUser

> V0040OpenapiUsersResp SlurmdbV0040GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetUser`: V0040OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include wckeys | 

### Return type

[**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetUsers

> V0040OpenapiUsersResp SlurmdbV0040GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default wckey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With wckeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetUsers`: V0040OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default wckey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With wckeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetWckey

> V0040OpenapiWckeyResp SlurmdbV0040GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetWckey`: V0040OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetWckeys

> V0040OpenapiWckeyResp SlurmdbV0040GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetWckeys`: V0040OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 

### Return type

[**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAccounts

> V0040OpenapiResp SlurmdbV0040PostAccounts(ctx).V0040OpenapiAccountsResp(v0040OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiAccountsResp := *openapiclient.NewV0040OpenapiAccountsResp([]openapiclient.V0040Account{*openapiclient.NewV0040Account("Description_example", "Name_example", "Organization_example")}) // V0040OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAccounts(context.Background()).V0040OpenapiAccountsResp(v0040OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAccounts`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAccountsResp** | [**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAccountsAssociation

> V0040OpenapiAccountsAddCondRespStr SlurmdbV0040PostAccountsAssociation(ctx).V0040OpenapiAccountsAddCondResp(v0040OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiAccountsAddCondResp := *openapiclient.NewV0040OpenapiAccountsAddCondResp() // V0040OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAccountsAssociation(context.Background()).V0040OpenapiAccountsAddCondResp(v0040OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAccountsAssociation`: V0040OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAccountsAddCondResp** | [**V0040OpenapiAccountsAddCondResp**](V0040OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0040OpenapiAccountsAddCondRespStr**](V0040OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAssociations

> V0040OpenapiResp SlurmdbV0040PostAssociations(ctx).V0040OpenapiAssocsResp(v0040OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiAssocsResp := *openapiclient.NewV0040OpenapiAssocsResp([]openapiclient.V0040Assoc{*openapiclient.NewV0040Assoc("User_example")}) // V0040OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAssociations(context.Background()).V0040OpenapiAssocsResp(v0040OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAssociations`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAssocsResp** | [**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostClusters

> V0040OpenapiResp SlurmdbV0040PostClusters(ctx).UpdateTime(updateTime).V0040OpenapiClustersResp(v0040OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)
	v0040OpenapiClustersResp := *openapiclient.NewV0040OpenapiClustersResp([]openapiclient.V0040ClusterRec{*openapiclient.NewV0040ClusterRec()}) // V0040OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostClusters(context.Background()).UpdateTime(updateTime).V0040OpenapiClustersResp(v0040OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostClusters`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 
 **v0040OpenapiClustersResp** | [**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostConfig

> V0040OpenapiResp SlurmdbV0040PostConfig(ctx).V0040OpenapiSlurmdbdConfigResp(v0040OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiSlurmdbdConfigResp := *openapiclient.NewV0040OpenapiSlurmdbdConfigResp() // V0040OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostConfig(context.Background()).V0040OpenapiSlurmdbdConfigResp(v0040OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostConfig`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiSlurmdbdConfigResp** | [**V0040OpenapiSlurmdbdConfigResp**](V0040OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostQos

> V0040OpenapiResp SlurmdbV0040PostQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0040OpenapiSlurmdbdQosResp(v0040OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)
	v0040OpenapiSlurmdbdQosResp := *openapiclient.NewV0040OpenapiSlurmdbdQosResp([]openapiclient.V0040Qos{*openapiclient.NewV0040Qos()}) // V0040OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0040OpenapiSlurmdbdQosResp(v0040OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostQos`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 
 **v0040OpenapiSlurmdbdQosResp** | [**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostTres

> V0040OpenapiResp SlurmdbV0040PostTres(ctx).V0040OpenapiTresResp(v0040OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiTresResp := *openapiclient.NewV0040OpenapiTresResp([]openapiclient.V0040Tres{*openapiclient.NewV0040Tres("Type_example")}) // V0040OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostTres(context.Background()).V0040OpenapiTresResp(v0040OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostTres`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiTresResp** | [**V0040OpenapiTresResp**](V0040OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostUsers

> V0040OpenapiResp SlurmdbV0040PostUsers(ctx).V0040OpenapiUsersResp(v0040OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0040OpenapiUsersResp := *openapiclient.NewV0040OpenapiUsersResp([]openapiclient.V0040User{*openapiclient.NewV0040User("Name_example")}) // V0040OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostUsers(context.Background()).V0040OpenapiUsersResp(v0040OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostUsers`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiUsersResp** | [**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostUsersAssociation

> V0040OpenapiUsersAddCondRespStr SlurmdbV0040PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0040OpenapiUsersAddCondResp(v0040OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0040OpenapiUsersAddCondResp := *openapiclient.NewV0040OpenapiUsersAddCondResp(*openapiclient.NewV0040UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0040UserShort()) // V0040OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0040OpenapiUsersAddCondResp(v0040OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostUsersAssociation`: V0040OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 
 **v0040OpenapiUsersAddCondResp** | [**V0040OpenapiUsersAddCondResp**](V0040OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0040OpenapiUsersAddCondRespStr**](V0040OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostWckeys

> V0040OpenapiResp SlurmdbV0040PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0040OpenapiWckeyResp(v0040OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)
	v0040OpenapiWckeyResp := *openapiclient.NewV0040OpenapiWckeyResp([]openapiclient.V0040Wckey{*openapiclient.NewV0040Wckey("Cluster_example", "Name_example", "User_example")}) // V0040OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0040OpenapiWckeyResp(v0040OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostWckeys`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 
 **v0040OpenapiWckeyResp** | [**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAccount

> V0041OpenapiAccountsRemovedResp SlurmdbV0041DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAccount`: V0041OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiAccountsRemovedResp**](V0041OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociation

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociation`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociations

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociations`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteCluster

> V0041OpenapiClustersRemovedResp SlurmdbV0041DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteCluster`: V0041OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0041OpenapiClustersRemovedResp**](V0041OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteSingleQos

> V0041OpenapiSlurmdbdQosRemovedResp SlurmdbV0041DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteSingleQos`: V0041OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiSlurmdbdQosRemovedResp**](V0041OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteUser

> V0041OpenapiResp SlurmdbV0041DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteUser`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteWckey

> V0041OpenapiWckeyRemovedResp SlurmdbV0041DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteWckey`: V0041OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiWckeyRemovedResp**](V0041OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccount

> V0041OpenapiAccountsResp SlurmdbV0041GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccount`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccounts

> V0041OpenapiAccountsResp SlurmdbV0041GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccounts`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociation

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociation`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociations

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociations`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetCluster

> V0041OpenapiClustersResp SlurmdbV0041GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetCluster`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetClusters

> V0041OpenapiClustersResp SlurmdbV0041GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetClusters`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetConfig

> V0041OpenapiSlurmdbdConfigResp SlurmdbV0041GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetConfig`: V0041OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetConfigRequest struct via the builder pattern


### Return type

[**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetDiag

> V0041OpenapiSlurmdbdStatsResp SlurmdbV0041GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetDiag`: V0041OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetDiagRequest struct via the builder pattern


### Return type

[**V0041OpenapiSlurmdbdStatsResp**](V0041OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstance

> V0041OpenapiInstancesResp SlurmdbV0041GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstance`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstances

> V0041OpenapiInstancesResp SlurmdbV0041GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstances`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJob

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	jobId := "jobId_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJob`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJobs

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV wckey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJobs`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV wckey list | 

### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetSingleQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetSingleQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetTres

> V0041OpenapiTresResp SlurmdbV0041GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetTres`: V0041OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetTresRequest struct via the builder pattern


### Return type

[**V0041OpenapiTresResp**](V0041OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUser

> V0041OpenapiUsersResp SlurmdbV0041GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUser`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include wckeys | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUsers

> V0041OpenapiUsersResp SlurmdbV0041GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default wckey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With wckeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUsers`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default wckey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With wckeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckey

> V0041OpenapiWckeyResp SlurmdbV0041GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckey`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckeys

> V0041OpenapiWckeyResp SlurmdbV0041GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckeys`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 

### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccounts

> V0041OpenapiResp SlurmdbV0041PostAccounts(ctx).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiAccountsResp := *openapiclient.NewV0041OpenapiAccountsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAccountsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAccountsInner("Description_example", "Name_example", "Organization_example")}) // V0041OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccounts(context.Background()).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccounts`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAccountsResp** | [**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccountsAssociation

> V0041OpenapiAccountsAddCondRespStr SlurmdbV0041PostAccountsAssociation(ctx).V0041OpenapiAccountsAddCondResp(v0041OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiAccountsAddCondResp := *openapiclient.NewV0041OpenapiAccountsAddCondResp() // V0041OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccountsAssociation(context.Background()).V0041OpenapiAccountsAddCondResp(v0041OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccountsAssociation`: V0041OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAccountsAddCondResp** | [**V0041OpenapiAccountsAddCondResp**](V0041OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0041OpenapiAccountsAddCondRespStr**](V0041OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAssociations

> V0041OpenapiResp SlurmdbV0041PostAssociations(ctx).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiAssocsResp := *openapiclient.NewV0041OpenapiAssocsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAssociationsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAssociationsInner("User_example")}) // V0041OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAssociations(context.Background()).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAssociations`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAssocsResp** | [**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostClusters

> V0041OpenapiResp SlurmdbV0041PostClusters(ctx).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)
	v0041OpenapiClustersResp := *openapiclient.NewV0041OpenapiClustersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespClustersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespClustersInner()}) // V0041OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostClusters(context.Background()).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostClusters`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 
 **v0041OpenapiClustersResp** | [**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostConfig

> V0041OpenapiResp SlurmdbV0041PostConfig(ctx).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiSlurmdbdConfigResp := *openapiclient.NewV0041OpenapiSlurmdbdConfigResp() // V0041OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostConfig(context.Background()).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostConfig`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiSlurmdbdConfigResp** | [**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostQos

> V0041OpenapiResp SlurmdbV0041PostQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)
	v0041OpenapiSlurmdbdQosResp := *openapiclient.NewV0041OpenapiSlurmdbdQosResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespQosInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespQosInner()}) // V0041OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostQos`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 
 **v0041OpenapiSlurmdbdQosResp** | [**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostTres

> V0041OpenapiResp SlurmdbV0041PostTres(ctx).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiTresResp := *openapiclient.NewV0041OpenapiTresResp([]openapiclient.V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner{*openapiclient.NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner("Type_example")}) // V0041OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostTres(context.Background()).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostTres`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiTresResp** | [**V0041OpenapiTresResp**](V0041OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsers

> V0041OpenapiResp SlurmdbV0041PostUsers(ctx).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0041OpenapiUsersResp := *openapiclient.NewV0041OpenapiUsersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInner("Name_example")}) // V0041OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsers(context.Background()).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsers`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiUsersResp** | [**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsersAssociation

> V0041OpenapiUsersAddCondRespStr SlurmdbV0041PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0041OpenapiUsersAddCondResp(v0041OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0041OpenapiUsersAddCondResp := *openapiclient.NewV0041OpenapiUsersAddCondResp(*openapiclient.NewV0041OpenapiUsersAddCondRespAssociationCondition([]string{"Users_example"}), *openapiclient.NewV0041OpenapiUsersAddCondRespUser()) // V0041OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0041OpenapiUsersAddCondResp(v0041OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsersAssociation`: V0041OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 
 **v0041OpenapiUsersAddCondResp** | [**V0041OpenapiUsersAddCondResp**](V0041OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0041OpenapiUsersAddCondRespStr**](V0041OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostWckeys

> V0041OpenapiResp SlurmdbV0041PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)
	v0041OpenapiWckeyResp := *openapiclient.NewV0041OpenapiWckeyResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner("Cluster_example", "Name_example", "User_example")}) // V0041OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostWckeys`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 
 **v0041OpenapiWckeyResp** | [**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAccount

> V0042OpenapiAccountsRemovedResp SlurmdbV0042DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAccount`: V0042OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiAccountsRemovedResp**](V0042OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAssociation

> V0042OpenapiAssocsRemovedResp SlurmdbV0042DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAssociation`: V0042OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsRemovedResp**](V0042OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAssociations

> V0042OpenapiAssocsRemovedResp SlurmdbV0042DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAssociations`: V0042OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsRemovedResp**](V0042OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteCluster

> V0042OpenapiClustersRemovedResp SlurmdbV0042DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteCluster`: V0042OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0042OpenapiClustersRemovedResp**](V0042OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteSingleQos

> V0042OpenapiSlurmdbdQosRemovedResp SlurmdbV0042DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteSingleQos`: V0042OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiSlurmdbdQosRemovedResp**](V0042OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteUser

> V0042OpenapiResp SlurmdbV0042DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteUser`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteWckey

> V0042OpenapiWckeyRemovedResp SlurmdbV0042DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteWckey`: V0042OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiWckeyRemovedResp**](V0042OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAccount

> V0042OpenapiAccountsResp SlurmdbV0042GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAccount`: V0042OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAccounts

> V0042OpenapiAccountsResp SlurmdbV0042GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAccounts`: V0042OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAssociation

> V0042OpenapiAssocsResp SlurmdbV0042GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAssociation`: V0042OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAssociations

> V0042OpenapiAssocsResp SlurmdbV0042GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAssociations`: V0042OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetCluster

> V0042OpenapiClustersResp SlurmdbV0042GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetCluster`: V0042OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetClusters

> V0042OpenapiClustersResp SlurmdbV0042GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetClusters`: V0042OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetConfig

> V0042OpenapiSlurmdbdConfigResp SlurmdbV0042GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetConfig`: V0042OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetConfigRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdConfigResp**](V0042OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetDiag

> V0042OpenapiSlurmdbdStatsResp SlurmdbV0042GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetDiag`: V0042OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetDiagRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdStatsResp**](V0042OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetInstance

> V0042OpenapiInstancesResp SlurmdbV0042GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetInstance`: V0042OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0042OpenapiInstancesResp**](V0042OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetInstances

> V0042OpenapiInstancesResp SlurmdbV0042GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetInstances`: V0042OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0042OpenapiInstancesResp**](V0042OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetJob

> V0042OpenapiSlurmdbdJobsResp SlurmdbV0042GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	jobId := "jobId_example" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetJob`: V0042OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiSlurmdbdJobsResp**](V0042OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetJobs

> V0042OpenapiSlurmdbdJobsResp SlurmdbV0042GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV WCKey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetJobs`: V0042OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV WCKey list | 

### Return type

[**V0042OpenapiSlurmdbdJobsResp**](V0042OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetPing

> V0042OpenapiSlurmdbdPingResp SlurmdbV0042GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetPing`: V0042OpenapiSlurmdbdPingResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetPingRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdPingResp**](V0042OpenapiSlurmdbdPingResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetQos

> V0042OpenapiSlurmdbdQosResp SlurmdbV0042GetQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetQos`: V0042OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 

### Return type

[**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetSingleQos

> V0042OpenapiSlurmdbdQosResp SlurmdbV0042GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetSingleQos`: V0042OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetTres

> V0042OpenapiTresResp SlurmdbV0042GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetTres`: V0042OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetTresRequest struct via the builder pattern


### Return type

[**V0042OpenapiTresResp**](V0042OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetUser

> V0042OpenapiUsersResp SlurmdbV0042GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetUser`: V0042OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include WCKeys | 

### Return type

[**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetUsers

> V0042OpenapiUsersResp SlurmdbV0042GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default WCKey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With WCKeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetUsers`: V0042OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default WCKey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With WCKeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetWckey

> V0042OpenapiWckeyResp SlurmdbV0042GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetWckey`: V0042OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetWckeys

> V0042OpenapiWckeyResp SlurmdbV0042GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetWckeys`: V0042OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 

### Return type

[**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAccounts

> V0042OpenapiResp SlurmdbV0042PostAccounts(ctx).V0042OpenapiAccountsResp(v0042OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiAccountsResp := *openapiclient.NewV0042OpenapiAccountsResp([]openapiclient.V0042Account{*openapiclient.NewV0042Account("Description_example", "Name_example", "Organization_example")}) // V0042OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAccounts(context.Background()).V0042OpenapiAccountsResp(v0042OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAccounts`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAccountsResp** | [**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAccountsAssociation

> V0042OpenapiAccountsAddCondRespStr SlurmdbV0042PostAccountsAssociation(ctx).V0042OpenapiAccountsAddCondResp(v0042OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiAccountsAddCondResp := *openapiclient.NewV0042OpenapiAccountsAddCondResp() // V0042OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAccountsAssociation(context.Background()).V0042OpenapiAccountsAddCondResp(v0042OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAccountsAssociation`: V0042OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAccountsAddCondResp** | [**V0042OpenapiAccountsAddCondResp**](V0042OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0042OpenapiAccountsAddCondRespStr**](V0042OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAssociations

> V0042OpenapiResp SlurmdbV0042PostAssociations(ctx).V0042OpenapiAssocsResp(v0042OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiAssocsResp := *openapiclient.NewV0042OpenapiAssocsResp([]openapiclient.V0042Assoc{*openapiclient.NewV0042Assoc("User_example")}) // V0042OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAssociations(context.Background()).V0042OpenapiAssocsResp(v0042OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAssociations`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAssocsResp** | [**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostClusters

> V0042OpenapiResp SlurmdbV0042PostClusters(ctx).UpdateTime(updateTime).V0042OpenapiClustersResp(v0042OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)
	v0042OpenapiClustersResp := *openapiclient.NewV0042OpenapiClustersResp([]openapiclient.V0042ClusterRec{*openapiclient.NewV0042ClusterRec()}) // V0042OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostClusters(context.Background()).UpdateTime(updateTime).V0042OpenapiClustersResp(v0042OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostClusters`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 
 **v0042OpenapiClustersResp** | [**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostConfig

> V0042OpenapiResp SlurmdbV0042PostConfig(ctx).V0042OpenapiSlurmdbdConfigResp(v0042OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiSlurmdbdConfigResp := *openapiclient.NewV0042OpenapiSlurmdbdConfigResp() // V0042OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostConfig(context.Background()).V0042OpenapiSlurmdbdConfigResp(v0042OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostConfig`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiSlurmdbdConfigResp** | [**V0042OpenapiSlurmdbdConfigResp**](V0042OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostQos

> V0042OpenapiResp SlurmdbV0042PostQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0042OpenapiSlurmdbdQosResp(v0042OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	v0042OpenapiSlurmdbdQosResp := *openapiclient.NewV0042OpenapiSlurmdbdQosResp([]openapiclient.V0042Qos{*openapiclient.NewV0042Qos()}) // V0042OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0042OpenapiSlurmdbdQosResp(v0042OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostQos`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **v0042OpenapiSlurmdbdQosResp** | [**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostTres

> V0042OpenapiResp SlurmdbV0042PostTres(ctx).V0042OpenapiTresResp(v0042OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiTresResp := *openapiclient.NewV0042OpenapiTresResp([]openapiclient.V0042Tres{*openapiclient.NewV0042Tres("Type_example")}) // V0042OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostTres(context.Background()).V0042OpenapiTresResp(v0042OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostTres`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiTresResp** | [**V0042OpenapiTresResp**](V0042OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostUsers

> V0042OpenapiResp SlurmdbV0042PostUsers(ctx).V0042OpenapiUsersResp(v0042OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	v0042OpenapiUsersResp := *openapiclient.NewV0042OpenapiUsersResp([]openapiclient.V0042User{*openapiclient.NewV0042User("Name_example")}) // V0042OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostUsers(context.Background()).V0042OpenapiUsersResp(v0042OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostUsers`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiUsersResp** | [**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostUsersAssociation

> V0042OpenapiUsersAddCondRespStr SlurmdbV0042PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0042OpenapiUsersAddCondResp(v0042OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0042OpenapiUsersAddCondResp := *openapiclient.NewV0042OpenapiUsersAddCondResp(*openapiclient.NewV0042UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0042UserShort()) // V0042OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0042OpenapiUsersAddCondResp(v0042OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostUsersAssociation`: V0042OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 
 **v0042OpenapiUsersAddCondResp** | [**V0042OpenapiUsersAddCondResp**](V0042OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0042OpenapiUsersAddCondRespStr**](V0042OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostWckeys

> V0042OpenapiResp SlurmdbV0042PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0042OpenapiWckeyResp(v0042OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zhuqw12/slurmrestapi/slurmrestapi"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)
	v0042OpenapiWckeyResp := *openapiclient.NewV0042OpenapiWckeyResp([]openapiclient.V0042Wckey{*openapiclient.NewV0042Wckey("Cluster_example", "Name_example", "User_example")}) // V0042OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0042OpenapiWckeyResp(v0042OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostWckeys`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 
 **v0042OpenapiWckeyResp** | [**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


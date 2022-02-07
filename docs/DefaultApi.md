# \DefaultApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchJob**](DefaultApi.md#CreateSearchJob) | **Post** /v1/search/jobs | Create a Search Job.
[**DeleteSearchJob**](DefaultApi.md#DeleteSearchJob) | **Delete** /v1/search/jobs/{searchJobId} | Delete a search job.
[**GetSearchJobMessages**](DefaultApi.md#GetSearchJobMessages) | **Get** /v1/search/jobs/{searchJobId}/messages | Fetch messages found by a search job.
[**GetSearchJobRecords**](DefaultApi.md#GetSearchJobRecords) | **Get** /v1/search/jobs/{searchJobId}/records | Fetch records produced by a search job.
[**GetSearchJobStatus**](DefaultApi.md#GetSearchJobStatus) | **Get** /v1/search/jobs/{searchJobId} | Get status for a search job.



## CreateSearchJob

> CreateSearchJob(ctx).SearchJobDefinition(searchJobDefinition).Execute()

Create a Search Job.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchJobDefinition := *openapiclient.NewSearchJobDefinition() // SearchJobDefinition | Create Search Job request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSearchJob(context.Background()).SearchJobDefinition(searchJobDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSearchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchJobDefinition** | [**SearchJobDefinition**](SearchJobDefinition.md) | Create Search Job request body. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSearchJob

> DeleteSearchJob(ctx, searchJobId).Execute()

Delete a search job.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchJobId := "2207828B17FE1D99" // string | The identifier of the search job to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSearchJob(context.Background(), searchJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSearchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchJobId** | **string** | The identifier of the search job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchJobMessages

> SearchJobMessages GetSearchJobMessages(ctx, searchJobId).Offset(offset).Limit(limit).Execute()

Fetch messages found by a search job.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchJobId := "2207828B17FE1D99" // string | The identifier of the search job to retrieve.
    offset := true // bool | The position from where to start the search operation.
    limit := true // bool | Limit the number of messages returned in the response. The number of messages returned may be less than the `limit`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSearchJobMessages(context.Background(), searchJobId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSearchJobMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchJobMessages`: SearchJobMessages
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSearchJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchJobId** | **string** | The identifier of the search job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchJobMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **bool** | The position from where to start the search operation. | 
 **limit** | **bool** | Limit the number of messages returned in the response. The number of messages returned may be less than the &#x60;limit&#x60;. | 

### Return type

[**SearchJobMessages**](SearchJobMessages.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchJobRecords

> SearchJobRecords GetSearchJobRecords(ctx, searchJobId).Offset(offset).Limit(limit).Execute()

Fetch records produced by a search job.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchJobId := "2207828B17FE1D99" // string | The identifier of the search job to retrieve.
    offset := true // bool | The position from where to start the search operation.
    limit := true // bool | Limit the number of records returned in the response. The number of records returned may be less than the `limit`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSearchJobRecords(context.Background(), searchJobId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSearchJobRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchJobRecords`: SearchJobRecords
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSearchJobRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchJobId** | **string** | The identifier of the search job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchJobRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **bool** | The position from where to start the search operation. | 
 **limit** | **bool** | Limit the number of records returned in the response. The number of records returned may be less than the &#x60;limit&#x60;. | 

### Return type

[**SearchJobRecords**](SearchJobRecords.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchJobStatus

> SearchJobState GetSearchJobStatus(ctx, searchJobId).Execute()

Get status for a search job.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchJobId := "2207828B17FE1D99" // string | The identifier of the search job to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSearchJobStatus(context.Background(), searchJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSearchJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchJobStatus`: SearchJobState
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSearchJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchJobId** | **string** | The identifier of the search job to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchJobState**](SearchJobState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


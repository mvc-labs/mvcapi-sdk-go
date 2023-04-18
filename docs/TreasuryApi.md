# \TreasuryApi

All URIs are relative to *https://testnet.mvcapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TreasuryGet**](TreasuryApi.md#TreasuryGet) | **Get** /treasury | Get current treasury info.
[**TreasuryHistoryGet**](TreasuryApi.md#TreasuryHistoryGet) | **Get** /treasury/history | Get all treasury history.



## TreasuryGet

> TreasuryInfo TreasuryGet(ctx).Execute()

Get current treasury info.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TreasuryApi.TreasuryGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TreasuryApi.TreasuryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TreasuryGet`: TreasuryInfo
    fmt.Fprintf(os.Stdout, "Response from `TreasuryApi.TreasuryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTreasuryGetRequest struct via the builder pattern


### Return type

[**TreasuryInfo**](TreasuryInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TreasuryHistoryGet

> []TreasuryHistory TreasuryHistoryGet(ctx).Execute()

Get all treasury history.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TreasuryApi.TreasuryHistoryGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TreasuryApi.TreasuryHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TreasuryHistoryGet`: []TreasuryHistory
    fmt.Fprintf(os.Stdout, "Response from `TreasuryApi.TreasuryHistoryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTreasuryHistoryGetRequest struct via the builder pattern


### Return type

[**[]TreasuryHistory**](TreasuryHistory.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


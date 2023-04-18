# \XpubApi

All URIs are relative to *https://testnet.mvcapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XpubLiteXpubAddressAddressGet**](XpubApi.md#XpubLiteXpubAddressAddressGet) | **Get** /xpubLite/{xpub}/address/{address} | Get xpub address type and index. Only index under /0/70 /1/30 is valid. Otherwise not found.
[**XpubLiteXpubBalanceGet**](XpubApi.md#XpubLiteXpubBalanceGet) | **Get** /xpubLite/{xpub}/balance | Get xpub balances including confirmed and unconfirmed.
[**XpubLiteXpubUtxoGet**](XpubApi.md#XpubLiteXpubUtxoGet) | **Get** /xpubLite/{xpub}/utxo | Get xpub utxos by specific xpub(300 per page).



## XpubLiteXpubAddressAddressGet

> XpubAddress XpubLiteXpubAddressAddressGet(ctx, xpub, address).Execute()

Get xpub address type and index. Only index under /0/70 /1/30 is valid. Otherwise not found.

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
    xpub := "xpub_example" // string | the requested xpub
    address := "address_example" // string | the requested address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubAddressAddressGet(context.Background(), xpub, address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubAddressAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubAddressAddressGet`: XpubAddress
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubAddressAddressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubAddressAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**XpubAddress**](XpubAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubLiteXpubBalanceGet

> XpubLiteBalance XpubLiteXpubBalanceGet(ctx, xpub).Execute()

Get xpub balances including confirmed and unconfirmed.



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
    xpub := "xpub_example" // string | the xpub to search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubBalanceGet(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubBalanceGet`: XpubLiteBalance
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the xpub to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XpubLiteBalance**](XpubLiteBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubLiteXpubUtxoGet

> []XpubUtxo XpubLiteXpubUtxoGet(ctx, xpub).Limit(limit).Execute()

Get xpub utxos by specific xpub(300 per page).

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
    xpub := "xpub_example" // string | the requested xpub
    limit := int32(56) // int32 | The max items returned in this query(default 300), not bigger than 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubUtxoGet(context.Background(), xpub).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubUtxoGet`: []XpubUtxo
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max items returned in this query(default 300), not bigger than 5000. | 

### Return type

[**[]XpubUtxo**](XpubUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


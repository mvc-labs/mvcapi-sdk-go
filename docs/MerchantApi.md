# \MerchantApi

All URIs are relative to *https://testnet.mvcapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantUtxoPost**](MerchantApi.md#MerchantUtxoPost) | **Post** /merchant/utxo | Pick utxos by addresses and amount.



## MerchantUtxoPost

> []AddressUtxo MerchantUtxoPost(ctx).UtxoPickRequest(utxoPickRequest).Execute()

Pick utxos by addresses and amount.



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
    utxoPickRequest := *openapiclient.NewUtxoPickRequest() // UtxoPickRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MerchantApi.MerchantUtxoPost(context.Background()).UtxoPickRequest(utxoPickRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantApi.MerchantUtxoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MerchantUtxoPost`: []AddressUtxo
    fmt.Fprintf(os.Stdout, "Response from `MerchantApi.MerchantUtxoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUtxoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **utxoPickRequest** | [**UtxoPickRequest**](UtxoPickRequest.md) |  | 

### Return type

[**[]AddressUtxo**](AddressUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


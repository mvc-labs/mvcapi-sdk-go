# \ContractApi

All URIs are relative to *https://api-mvc-testnet.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractFtAddressAddressBalanceConfirmedGet**](ContractApi.md#ContractFtAddressAddressBalanceConfirmedGet) | **Get** /contract/ft/address/{address}/balance/confirmed | Get all contract token balances for specific address ignoring all unconfirmed txs.
[**ContractFtAddressAddressBalanceGet**](ContractApi.md#ContractFtAddressAddressBalanceGet) | **Get** /contract/ft/address/{address}/balance | Get all contract token balances for specific address.
[**ContractFtAddressAddressUtxoGet**](ContractApi.md#ContractFtAddressAddressUtxoGet) | **Get** /contract/ft/address/{address}/utxo | Get all contract token utxos for specific address.
[**ContractNftAddressAddressCountConfirmedGet**](ContractApi.md#ContractNftAddressAddressCountConfirmedGet) | **Get** /contract/nft/address/{address}/count/confirmed | Get confirmed utxo count for specific nft(ignore all unconfirmed txs).
[**ContractNftAddressAddressSummaryGet**](ContractApi.md#ContractNftAddressAddressSummaryGet) | **Get** /contract/nft/address/{address}/summary | Get nft summary(NFT count group by genesis) for address.
[**ContractNftAddressAddressUtxoGet**](ContractApi.md#ContractNftAddressAddressUtxoGet) | **Get** /contract/nft/address/{address}/utxo | Get all contract nft token utxos for specific address.
[**ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet**](ContractApi.md#ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet) | **Get** /contract/nft/auction/codeHash/{codeHash}/nftId/{nftId}/utxo | Get all contract nft token utxos by codeHash and genesisId.
[**ContractNftGenesisCodeHashGenesisSummaryGet**](ContractApi.md#ContractNftGenesisCodeHashGenesisSummaryGet) | **Get** /contract/nft/genesis/{codeHash}/{genesis}/summary | Get nft summary(count group by address) for specific codeHash and genesisId(result cached for 60s).
[**ContractNftGenesisCodeHashGenesisUtxoGet**](ContractApi.md#ContractNftGenesisCodeHashGenesisUtxoGet) | **Get** /contract/nft/genesis/{codeHash}/{genesis}/utxo | Get all contract nft token utxos by codeHash and genesisId.
[**ContractNftSellAddressAddressUtxoGet**](ContractApi.md#ContractNftSellAddressAddressUtxoGet) | **Get** /contract/nft/sell/address/{address}/utxo | Get all contract sell sell utxos for specific address.
[**ContractNftSellGenesisCodeHashGenesisUtxoGet**](ContractApi.md#ContractNftSellGenesisCodeHashGenesisUtxoGet) | **Get** /contract/nft/sell/genesis/{codeHash}/{genesis}/utxo | Get all contract nft token utxos by codeHash and genesisId.
[**ContractNftSellV2AddressAddressUtxoGet**](ContractApi.md#ContractNftSellV2AddressAddressUtxoGet) | **Get** /contract/nft/sellV2/address/{address}/utxo | Get all contract sell sell utxos for specific address.
[**ContractNftSellV2GenesisCodeHashGenesisUtxoGet**](ContractApi.md#ContractNftSellV2GenesisCodeHashGenesisUtxoGet) | **Get** /contract/nft/sellV2/genesis/{codeHash}/{genesis}/utxo | Get all contract nft token utxos by codeHash and genesisId.
[**ContractUniqueGenesisCodeHashGenesisUtxoGet**](ContractApi.md#ContractUniqueGenesisCodeHashGenesisUtxoGet) | **Get** /contract/unique/genesis/{codeHash}/{genesis}/utxo | Get contract unique utxos by codeHash and genesisId.



## ContractFtAddressAddressBalanceConfirmedGet

> int64 ContractFtAddressAddressBalanceConfirmedGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get all contract token balances for specific address ignoring all unconfirmed txs.

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
    address := "address_example" // string | the requested address
    codeHash := "codeHash_example" // string | Filter by contract code hash
    genesis := "genesis_example" // string | Filter by contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractFtAddressAddressBalanceConfirmedGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractFtAddressAddressBalanceConfirmedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractFtAddressAddressBalanceConfirmedGet`: int64
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractFtAddressAddressBalanceConfirmedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractFtAddressAddressBalanceConfirmedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

### Return type

**int64**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractFtAddressAddressBalanceGet

> []ContractFtBalance ContractFtAddressAddressBalanceGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get all contract token balances for specific address.

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
    address := "address_example" // string | the requested address
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractFtAddressAddressBalanceGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractFtAddressAddressBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractFtAddressAddressBalanceGet`: []ContractFtBalance
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractFtAddressAddressBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractFtAddressAddressBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

### Return type

[**[]ContractFtBalance**](ContractFtBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractFtAddressAddressUtxoGet

> []ContractFtUtxo ContractFtAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all contract token utxos for specific address.

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
    address := "address_example" // string | the requested address
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractFtAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractFtAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractFtAddressAddressUtxoGet`: []ContractFtUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractFtAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractFtAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]ContractFtUtxo**](ContractFtUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftAddressAddressCountConfirmedGet

> int32 ContractNftAddressAddressCountConfirmedGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get confirmed utxo count for specific nft(ignore all unconfirmed txs).

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
    address := "address_example" // string | the requested address
    codeHash := "codeHash_example" // string | Filter by contract code hash
    genesis := "genesis_example" // string | Filter by contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftAddressAddressCountConfirmedGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftAddressAddressCountConfirmedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftAddressAddressCountConfirmedGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftAddressAddressCountConfirmedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftAddressAddressCountConfirmedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

### Return type

**int32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftAddressAddressSummaryGet

> []ContractNftAddressSummary ContractNftAddressAddressSummaryGet(ctx, address).Execute()

Get nft summary(NFT count group by genesis) for address.

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
    address := "address_example" // string | the requested address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftAddressAddressSummaryGet(context.Background(), address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftAddressAddressSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftAddressAddressSummaryGet`: []ContractNftAddressSummary
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftAddressAddressSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftAddressAddressSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContractNftAddressSummary**](ContractNftAddressSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftAddressAddressUtxoGet

> []ContractNftUtxo ContractNftAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Limit(limit).Flag(flag).Execute()

Get all contract nft token utxos for specific address.

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
    address := "address_example" // string | the requested address
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    limit := int32(56) // int32 | Limit the return count(no more than 300) (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Limit(limit).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftAddressAddressUtxoGet`: []ContractNftUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **limit** | **int32** | Limit the return count(no more than 300) | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]ContractNftUtxo**](ContractNftUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet

> []ContractNftAuctionUtxo ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet(ctx, codeHash, nftId).Execute()

Get all contract nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    nftId := "nftId_example" // string | Nft id of this auction.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet(context.Background(), codeHash, nftId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet`: []ContractNftAuctionUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**nftId** | **string** | Nft id of this auction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftAuctionCodeHashCodeHashNftIdNftIdUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ContractNftAuctionUtxo**](ContractNftAuctionUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftGenesisCodeHashGenesisSummaryGet

> []ContractNftGenesisSummary ContractNftGenesisCodeHashGenesisSummaryGet(ctx, codeHash, genesis).Execute()

Get nft summary(count group by address) for specific codeHash and genesisId(result cached for 60s).

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftGenesisCodeHashGenesisSummaryGet(context.Background(), codeHash, genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftGenesisCodeHashGenesisSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftGenesisCodeHashGenesisSummaryGet`: []ContractNftGenesisSummary
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftGenesisCodeHashGenesisSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftGenesisCodeHashGenesisSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ContractNftGenesisSummary**](ContractNftGenesisSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftGenesisCodeHashGenesisUtxoGet

> []ContractNftUtxo ContractNftGenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all contract nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftGenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftGenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftGenesisCodeHashGenesisUtxoGet`: []ContractNftUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftGenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftGenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]ContractNftUtxo**](ContractNftUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftSellAddressAddressUtxoGet

> []ContractNftSellUtxo ContractNftSellAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all contract sell sell utxos for specific address.

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
    address := "address_example" // string | Owner address.
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftSellAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftSellAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftSellAddressAddressUtxoGet`: []ContractNftSellUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftSellAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Owner address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftSellAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]ContractNftSellUtxo**](ContractNftSellUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftSellGenesisCodeHashGenesisUtxoGet

> []ContractNftSellUtxo ContractNftSellGenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all contract nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftSellGenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftSellGenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftSellGenesisCodeHashGenesisUtxoGet`: []ContractNftSellUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftSellGenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftSellGenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]ContractNftSellUtxo**](ContractNftSellUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftSellV2AddressAddressUtxoGet

> []ContractNftSellV2Utxo ContractNftSellV2AddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all contract sell sell utxos for specific address.

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
    address := "address_example" // string | Owner address.
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftSellV2AddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftSellV2AddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftSellV2AddressAddressUtxoGet`: []ContractNftSellV2Utxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftSellV2AddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Owner address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftSellV2AddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]ContractNftSellV2Utxo**](ContractNftSellV2Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractNftSellV2GenesisCodeHashGenesisUtxoGet

> []ContractNftSellV2Utxo ContractNftSellV2GenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all contract nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractNftSellV2GenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractNftSellV2GenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractNftSellV2GenesisCodeHashGenesisUtxoGet`: []ContractNftSellV2Utxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractNftSellV2GenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractNftSellV2GenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]ContractNftSellV2Utxo**](ContractNftSellV2Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractUniqueGenesisCodeHashGenesisUtxoGet

> []ContractUniqueUtxo ContractUniqueGenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).Execute()

Get contract unique utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.ContractUniqueGenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractUniqueGenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractUniqueGenesisCodeHashGenesisUtxoGet`: []ContractUniqueUtxo
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.ContractUniqueGenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractUniqueGenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ContractUniqueUtxo**](ContractUniqueUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


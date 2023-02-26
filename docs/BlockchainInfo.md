# BlockchainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | mainnet or testnet | [optional] 
**Blocks** | Pointer to **int32** | current block count | [optional] 
**Headers** | Pointer to **int32** | current block header count | [optional] 
**BestBlockHash** | Pointer to **string** | current highest blockhash | [optional] 
**Difficulty** | Pointer to **string** | decimal string for current difficulty | [optional] 
**MedianTime** | Pointer to **int64** | current median time | [optional] 
**Chainwork** | Pointer to **string** | current pow | [optional] 
**NetworkHashPerSecond** | Pointer to **string** | estimated current network hash rate. | [optional] 
**MempoolTxCount** | Pointer to **int32** | current mempool transaction count. | [optional] 
**MempoolUsage** | Pointer to **int64** | current mempool usage. | [optional] 
**EstimatedBlockSize** | Pointer to **int64** | next estimated block size. | [optional] 

## Methods

### NewBlockchainInfo

`func NewBlockchainInfo() *BlockchainInfo`

NewBlockchainInfo instantiates a new BlockchainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainInfoWithDefaults

`func NewBlockchainInfoWithDefaults() *BlockchainInfo`

NewBlockchainInfoWithDefaults instantiates a new BlockchainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *BlockchainInfo) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *BlockchainInfo) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *BlockchainInfo) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *BlockchainInfo) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetBlocks

`func (o *BlockchainInfo) GetBlocks() int32`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlockchainInfo) GetBlocksOk() (*int32, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlockchainInfo) SetBlocks(v int32)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BlockchainInfo) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetHeaders

`func (o *BlockchainInfo) GetHeaders() int32`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BlockchainInfo) GetHeadersOk() (*int32, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BlockchainInfo) SetHeaders(v int32)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BlockchainInfo) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBestBlockHash

`func (o *BlockchainInfo) GetBestBlockHash() string`

GetBestBlockHash returns the BestBlockHash field if non-nil, zero value otherwise.

### GetBestBlockHashOk

`func (o *BlockchainInfo) GetBestBlockHashOk() (*string, bool)`

GetBestBlockHashOk returns a tuple with the BestBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestBlockHash

`func (o *BlockchainInfo) SetBestBlockHash(v string)`

SetBestBlockHash sets BestBlockHash field to given value.

### HasBestBlockHash

`func (o *BlockchainInfo) HasBestBlockHash() bool`

HasBestBlockHash returns a boolean if a field has been set.

### GetDifficulty

`func (o *BlockchainInfo) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *BlockchainInfo) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *BlockchainInfo) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *BlockchainInfo) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetMedianTime

`func (o *BlockchainInfo) GetMedianTime() int64`

GetMedianTime returns the MedianTime field if non-nil, zero value otherwise.

### GetMedianTimeOk

`func (o *BlockchainInfo) GetMedianTimeOk() (*int64, bool)`

GetMedianTimeOk returns a tuple with the MedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTime

`func (o *BlockchainInfo) SetMedianTime(v int64)`

SetMedianTime sets MedianTime field to given value.

### HasMedianTime

`func (o *BlockchainInfo) HasMedianTime() bool`

HasMedianTime returns a boolean if a field has been set.

### GetChainwork

`func (o *BlockchainInfo) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *BlockchainInfo) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *BlockchainInfo) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.

### HasChainwork

`func (o *BlockchainInfo) HasChainwork() bool`

HasChainwork returns a boolean if a field has been set.

### GetNetworkHashPerSecond

`func (o *BlockchainInfo) GetNetworkHashPerSecond() string`

GetNetworkHashPerSecond returns the NetworkHashPerSecond field if non-nil, zero value otherwise.

### GetNetworkHashPerSecondOk

`func (o *BlockchainInfo) GetNetworkHashPerSecondOk() (*string, bool)`

GetNetworkHashPerSecondOk returns a tuple with the NetworkHashPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHashPerSecond

`func (o *BlockchainInfo) SetNetworkHashPerSecond(v string)`

SetNetworkHashPerSecond sets NetworkHashPerSecond field to given value.

### HasNetworkHashPerSecond

`func (o *BlockchainInfo) HasNetworkHashPerSecond() bool`

HasNetworkHashPerSecond returns a boolean if a field has been set.

### GetMempoolTxCount

`func (o *BlockchainInfo) GetMempoolTxCount() int32`

GetMempoolTxCount returns the MempoolTxCount field if non-nil, zero value otherwise.

### GetMempoolTxCountOk

`func (o *BlockchainInfo) GetMempoolTxCountOk() (*int32, bool)`

GetMempoolTxCountOk returns a tuple with the MempoolTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMempoolTxCount

`func (o *BlockchainInfo) SetMempoolTxCount(v int32)`

SetMempoolTxCount sets MempoolTxCount field to given value.

### HasMempoolTxCount

`func (o *BlockchainInfo) HasMempoolTxCount() bool`

HasMempoolTxCount returns a boolean if a field has been set.

### GetMempoolUsage

`func (o *BlockchainInfo) GetMempoolUsage() int64`

GetMempoolUsage returns the MempoolUsage field if non-nil, zero value otherwise.

### GetMempoolUsageOk

`func (o *BlockchainInfo) GetMempoolUsageOk() (*int64, bool)`

GetMempoolUsageOk returns a tuple with the MempoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMempoolUsage

`func (o *BlockchainInfo) SetMempoolUsage(v int64)`

SetMempoolUsage sets MempoolUsage field to given value.

### HasMempoolUsage

`func (o *BlockchainInfo) HasMempoolUsage() bool`

HasMempoolUsage returns a boolean if a field has been set.

### GetEstimatedBlockSize

`func (o *BlockchainInfo) GetEstimatedBlockSize() int64`

GetEstimatedBlockSize returns the EstimatedBlockSize field if non-nil, zero value otherwise.

### GetEstimatedBlockSizeOk

`func (o *BlockchainInfo) GetEstimatedBlockSizeOk() (*int64, bool)`

GetEstimatedBlockSizeOk returns a tuple with the EstimatedBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedBlockSize

`func (o *BlockchainInfo) SetEstimatedBlockSize(v int64)`

SetEstimatedBlockSize sets EstimatedBlockSize field to given value.

### HasEstimatedBlockSize

`func (o *BlockchainInfo) HasEstimatedBlockSize() bool`

HasEstimatedBlockSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



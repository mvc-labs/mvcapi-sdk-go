# BlockHeaderIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | Pointer to **string** | Block hash. | [optional] 
**Height** | Pointer to **int64** | Block height. | [optional] 
**Version** | Pointer to **int64** | Block version. | [optional] 
**PrevBlockHash** | Pointer to **string** | The block hash of the previous block. | [optional] 
**MerkleRoot** | Pointer to **string** | Hex formatted block merkle root. | [optional] 
**Timestamp** | Pointer to **int64** | Block timestamp. | [optional] 
**MedianTime** | Pointer to **int64** | Block median timestamp. | [optional] 
**Reward** | Pointer to **int64** | Miner total rewards, including miner fee. | [optional] 
**Miner** | Pointer to **string** | Guessed miner name. | [optional] 
**MinerAddress** | Pointer to **string** | Miner address that received rewards. | [optional] 
**TxCount** | Pointer to **int32** | Total txs count included in the block. | [optional] 
**InputCount** | Pointer to **int32** | Total input count in the block. | [optional] 
**OutputCount** | Pointer to **int32** | Total output count in the block | [optional] 
**Size** | Pointer to **int64** | Size of the block | [optional] 
**Bits** | Pointer to **int64** | Target bits. | [optional] 
**Nonce** | Pointer to **int64** | Block nonce | [optional] 
**Coinbase** | Pointer to **string** | Hex formated coinbase info. | [optional] 

## Methods

### NewBlockHeaderIndex

`func NewBlockHeaderIndex() *BlockHeaderIndex`

NewBlockHeaderIndex instantiates a new BlockHeaderIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeaderIndexWithDefaults

`func NewBlockHeaderIndexWithDefaults() *BlockHeaderIndex`

NewBlockHeaderIndexWithDefaults instantiates a new BlockHeaderIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *BlockHeaderIndex) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *BlockHeaderIndex) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *BlockHeaderIndex) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *BlockHeaderIndex) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetHeight

`func (o *BlockHeaderIndex) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockHeaderIndex) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockHeaderIndex) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BlockHeaderIndex) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetVersion

`func (o *BlockHeaderIndex) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockHeaderIndex) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockHeaderIndex) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlockHeaderIndex) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPrevBlockHash

`func (o *BlockHeaderIndex) GetPrevBlockHash() string`

GetPrevBlockHash returns the PrevBlockHash field if non-nil, zero value otherwise.

### GetPrevBlockHashOk

`func (o *BlockHeaderIndex) GetPrevBlockHashOk() (*string, bool)`

GetPrevBlockHashOk returns a tuple with the PrevBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevBlockHash

`func (o *BlockHeaderIndex) SetPrevBlockHash(v string)`

SetPrevBlockHash sets PrevBlockHash field to given value.

### HasPrevBlockHash

`func (o *BlockHeaderIndex) HasPrevBlockHash() bool`

HasPrevBlockHash returns a boolean if a field has been set.

### GetMerkleRoot

`func (o *BlockHeaderIndex) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *BlockHeaderIndex) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *BlockHeaderIndex) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.

### HasMerkleRoot

`func (o *BlockHeaderIndex) HasMerkleRoot() bool`

HasMerkleRoot returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlockHeaderIndex) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockHeaderIndex) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockHeaderIndex) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockHeaderIndex) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMedianTime

`func (o *BlockHeaderIndex) GetMedianTime() int64`

GetMedianTime returns the MedianTime field if non-nil, zero value otherwise.

### GetMedianTimeOk

`func (o *BlockHeaderIndex) GetMedianTimeOk() (*int64, bool)`

GetMedianTimeOk returns a tuple with the MedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTime

`func (o *BlockHeaderIndex) SetMedianTime(v int64)`

SetMedianTime sets MedianTime field to given value.

### HasMedianTime

`func (o *BlockHeaderIndex) HasMedianTime() bool`

HasMedianTime returns a boolean if a field has been set.

### GetReward

`func (o *BlockHeaderIndex) GetReward() int64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *BlockHeaderIndex) GetRewardOk() (*int64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *BlockHeaderIndex) SetReward(v int64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *BlockHeaderIndex) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetMiner

`func (o *BlockHeaderIndex) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *BlockHeaderIndex) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *BlockHeaderIndex) SetMiner(v string)`

SetMiner sets Miner field to given value.

### HasMiner

`func (o *BlockHeaderIndex) HasMiner() bool`

HasMiner returns a boolean if a field has been set.

### GetMinerAddress

`func (o *BlockHeaderIndex) GetMinerAddress() string`

GetMinerAddress returns the MinerAddress field if non-nil, zero value otherwise.

### GetMinerAddressOk

`func (o *BlockHeaderIndex) GetMinerAddressOk() (*string, bool)`

GetMinerAddressOk returns a tuple with the MinerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinerAddress

`func (o *BlockHeaderIndex) SetMinerAddress(v string)`

SetMinerAddress sets MinerAddress field to given value.

### HasMinerAddress

`func (o *BlockHeaderIndex) HasMinerAddress() bool`

HasMinerAddress returns a boolean if a field has been set.

### GetTxCount

`func (o *BlockHeaderIndex) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *BlockHeaderIndex) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *BlockHeaderIndex) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.

### HasTxCount

`func (o *BlockHeaderIndex) HasTxCount() bool`

HasTxCount returns a boolean if a field has been set.

### GetInputCount

`func (o *BlockHeaderIndex) GetInputCount() int32`

GetInputCount returns the InputCount field if non-nil, zero value otherwise.

### GetInputCountOk

`func (o *BlockHeaderIndex) GetInputCountOk() (*int32, bool)`

GetInputCountOk returns a tuple with the InputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCount

`func (o *BlockHeaderIndex) SetInputCount(v int32)`

SetInputCount sets InputCount field to given value.

### HasInputCount

`func (o *BlockHeaderIndex) HasInputCount() bool`

HasInputCount returns a boolean if a field has been set.

### GetOutputCount

`func (o *BlockHeaderIndex) GetOutputCount() int32`

GetOutputCount returns the OutputCount field if non-nil, zero value otherwise.

### GetOutputCountOk

`func (o *BlockHeaderIndex) GetOutputCountOk() (*int32, bool)`

GetOutputCountOk returns a tuple with the OutputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCount

`func (o *BlockHeaderIndex) SetOutputCount(v int32)`

SetOutputCount sets OutputCount field to given value.

### HasOutputCount

`func (o *BlockHeaderIndex) HasOutputCount() bool`

HasOutputCount returns a boolean if a field has been set.

### GetSize

`func (o *BlockHeaderIndex) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockHeaderIndex) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockHeaderIndex) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockHeaderIndex) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetBits

`func (o *BlockHeaderIndex) GetBits() int64`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *BlockHeaderIndex) GetBitsOk() (*int64, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *BlockHeaderIndex) SetBits(v int64)`

SetBits sets Bits field to given value.

### HasBits

`func (o *BlockHeaderIndex) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetNonce

`func (o *BlockHeaderIndex) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockHeaderIndex) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockHeaderIndex) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BlockHeaderIndex) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetCoinbase

`func (o *BlockHeaderIndex) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *BlockHeaderIndex) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *BlockHeaderIndex) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *BlockHeaderIndex) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



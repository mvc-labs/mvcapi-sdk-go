# BlockHeaderPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** | Block height. | [optional] 
**BlockHash** | Pointer to **string** | Block hash. | [optional] 
**Timestamp** | Pointer to **int64** | Block timestamp. | [optional] 
**MedianTime** | Pointer to **int64** | current median time | [optional] 
**Reward** | Pointer to **int64** | Miner total rewards, including miner fee. | [optional] 
**Miner** | Pointer to **string** | Guessed miner name. | [optional] 
**TxCount** | Pointer to **int32** | Total txs count included in the block. | [optional] 
**Size** | Pointer to **int64** | Size of the block | [optional] 

## Methods

### NewBlockHeaderPage

`func NewBlockHeaderPage() *BlockHeaderPage`

NewBlockHeaderPage instantiates a new BlockHeaderPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeaderPageWithDefaults

`func NewBlockHeaderPageWithDefaults() *BlockHeaderPage`

NewBlockHeaderPageWithDefaults instantiates a new BlockHeaderPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *BlockHeaderPage) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockHeaderPage) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockHeaderPage) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BlockHeaderPage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBlockHash

`func (o *BlockHeaderPage) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *BlockHeaderPage) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *BlockHeaderPage) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *BlockHeaderPage) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlockHeaderPage) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockHeaderPage) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockHeaderPage) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockHeaderPage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMedianTime

`func (o *BlockHeaderPage) GetMedianTime() int64`

GetMedianTime returns the MedianTime field if non-nil, zero value otherwise.

### GetMedianTimeOk

`func (o *BlockHeaderPage) GetMedianTimeOk() (*int64, bool)`

GetMedianTimeOk returns a tuple with the MedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTime

`func (o *BlockHeaderPage) SetMedianTime(v int64)`

SetMedianTime sets MedianTime field to given value.

### HasMedianTime

`func (o *BlockHeaderPage) HasMedianTime() bool`

HasMedianTime returns a boolean if a field has been set.

### GetReward

`func (o *BlockHeaderPage) GetReward() int64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *BlockHeaderPage) GetRewardOk() (*int64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *BlockHeaderPage) SetReward(v int64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *BlockHeaderPage) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetMiner

`func (o *BlockHeaderPage) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *BlockHeaderPage) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *BlockHeaderPage) SetMiner(v string)`

SetMiner sets Miner field to given value.

### HasMiner

`func (o *BlockHeaderPage) HasMiner() bool`

HasMiner returns a boolean if a field has been set.

### GetTxCount

`func (o *BlockHeaderPage) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *BlockHeaderPage) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *BlockHeaderPage) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.

### HasTxCount

`func (o *BlockHeaderPage) HasTxCount() bool`

HasTxCount returns a boolean if a field has been set.

### GetSize

`func (o *BlockHeaderPage) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockHeaderPage) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockHeaderPage) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockHeaderPage) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



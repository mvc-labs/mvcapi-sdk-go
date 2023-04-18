/*
 * MicrovisionChain API Document
 *
 * API definition for MicrovisionChain provided apis
 *
 * API version: 3.0.8
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// BlockchainInfo struct for BlockchainInfo
type BlockchainInfo struct {
	// mainnet or testnet
	Chain *string `json:"chain,omitempty"`
	// current block count
	Blocks *int32 `json:"blocks,omitempty"`
	// current block header count
	Headers *int32 `json:"headers,omitempty"`
	// current highest blockhash
	BestBlockHash *string `json:"bestBlockHash,omitempty"`
	// decimal string for current difficulty
	Difficulty *string `json:"difficulty,omitempty"`
	// current median time
	MedianTime *int64 `json:"medianTime,omitempty"`
	// current pow
	Chainwork *string `json:"chainwork,omitempty"`
	// estimated current network hash rate.
	NetworkHashPerSecond *string `json:"networkHashPerSecond,omitempty"`
	// current mempool transaction count.
	MempoolTxCount *int32 `json:"mempoolTxCount,omitempty"`
	// current mempool usage.
	MempoolUsage *int64 `json:"mempoolUsage,omitempty"`
	// next estimated block size.
	EstimatedBlockSize *int64 `json:"estimatedBlockSize,omitempty"`
}

// NewBlockchainInfo instantiates a new BlockchainInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainInfo() *BlockchainInfo {
	this := BlockchainInfo{}
	return &this
}

// NewBlockchainInfoWithDefaults instantiates a new BlockchainInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainInfoWithDefaults() *BlockchainInfo {
	this := BlockchainInfo{}
	return &this
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *BlockchainInfo) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *BlockchainInfo) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *BlockchainInfo) SetChain(v string) {
	o.Chain = &v
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *BlockchainInfo) GetBlocks() int32 {
	if o == nil || o.Blocks == nil {
		var ret int32
		return ret
	}
	return *o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetBlocksOk() (*int32, bool) {
	if o == nil || o.Blocks == nil {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *BlockchainInfo) HasBlocks() bool {
	if o != nil && o.Blocks != nil {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given int32 and assigns it to the Blocks field.
func (o *BlockchainInfo) SetBlocks(v int32) {
	o.Blocks = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *BlockchainInfo) GetHeaders() int32 {
	if o == nil || o.Headers == nil {
		var ret int32
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetHeadersOk() (*int32, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BlockchainInfo) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given int32 and assigns it to the Headers field.
func (o *BlockchainInfo) SetHeaders(v int32) {
	o.Headers = &v
}

// GetBestBlockHash returns the BestBlockHash field value if set, zero value otherwise.
func (o *BlockchainInfo) GetBestBlockHash() string {
	if o == nil || o.BestBlockHash == nil {
		var ret string
		return ret
	}
	return *o.BestBlockHash
}

// GetBestBlockHashOk returns a tuple with the BestBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetBestBlockHashOk() (*string, bool) {
	if o == nil || o.BestBlockHash == nil {
		return nil, false
	}
	return o.BestBlockHash, true
}

// HasBestBlockHash returns a boolean if a field has been set.
func (o *BlockchainInfo) HasBestBlockHash() bool {
	if o != nil && o.BestBlockHash != nil {
		return true
	}

	return false
}

// SetBestBlockHash gets a reference to the given string and assigns it to the BestBlockHash field.
func (o *BlockchainInfo) SetBestBlockHash(v string) {
	o.BestBlockHash = &v
}

// GetDifficulty returns the Difficulty field value if set, zero value otherwise.
func (o *BlockchainInfo) GetDifficulty() string {
	if o == nil || o.Difficulty == nil {
		var ret string
		return ret
	}
	return *o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetDifficultyOk() (*string, bool) {
	if o == nil || o.Difficulty == nil {
		return nil, false
	}
	return o.Difficulty, true
}

// HasDifficulty returns a boolean if a field has been set.
func (o *BlockchainInfo) HasDifficulty() bool {
	if o != nil && o.Difficulty != nil {
		return true
	}

	return false
}

// SetDifficulty gets a reference to the given string and assigns it to the Difficulty field.
func (o *BlockchainInfo) SetDifficulty(v string) {
	o.Difficulty = &v
}

// GetMedianTime returns the MedianTime field value if set, zero value otherwise.
func (o *BlockchainInfo) GetMedianTime() int64 {
	if o == nil || o.MedianTime == nil {
		var ret int64
		return ret
	}
	return *o.MedianTime
}

// GetMedianTimeOk returns a tuple with the MedianTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetMedianTimeOk() (*int64, bool) {
	if o == nil || o.MedianTime == nil {
		return nil, false
	}
	return o.MedianTime, true
}

// HasMedianTime returns a boolean if a field has been set.
func (o *BlockchainInfo) HasMedianTime() bool {
	if o != nil && o.MedianTime != nil {
		return true
	}

	return false
}

// SetMedianTime gets a reference to the given int64 and assigns it to the MedianTime field.
func (o *BlockchainInfo) SetMedianTime(v int64) {
	o.MedianTime = &v
}

// GetChainwork returns the Chainwork field value if set, zero value otherwise.
func (o *BlockchainInfo) GetChainwork() string {
	if o == nil || o.Chainwork == nil {
		var ret string
		return ret
	}
	return *o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetChainworkOk() (*string, bool) {
	if o == nil || o.Chainwork == nil {
		return nil, false
	}
	return o.Chainwork, true
}

// HasChainwork returns a boolean if a field has been set.
func (o *BlockchainInfo) HasChainwork() bool {
	if o != nil && o.Chainwork != nil {
		return true
	}

	return false
}

// SetChainwork gets a reference to the given string and assigns it to the Chainwork field.
func (o *BlockchainInfo) SetChainwork(v string) {
	o.Chainwork = &v
}

// GetNetworkHashPerSecond returns the NetworkHashPerSecond field value if set, zero value otherwise.
func (o *BlockchainInfo) GetNetworkHashPerSecond() string {
	if o == nil || o.NetworkHashPerSecond == nil {
		var ret string
		return ret
	}
	return *o.NetworkHashPerSecond
}

// GetNetworkHashPerSecondOk returns a tuple with the NetworkHashPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetNetworkHashPerSecondOk() (*string, bool) {
	if o == nil || o.NetworkHashPerSecond == nil {
		return nil, false
	}
	return o.NetworkHashPerSecond, true
}

// HasNetworkHashPerSecond returns a boolean if a field has been set.
func (o *BlockchainInfo) HasNetworkHashPerSecond() bool {
	if o != nil && o.NetworkHashPerSecond != nil {
		return true
	}

	return false
}

// SetNetworkHashPerSecond gets a reference to the given string and assigns it to the NetworkHashPerSecond field.
func (o *BlockchainInfo) SetNetworkHashPerSecond(v string) {
	o.NetworkHashPerSecond = &v
}

// GetMempoolTxCount returns the MempoolTxCount field value if set, zero value otherwise.
func (o *BlockchainInfo) GetMempoolTxCount() int32 {
	if o == nil || o.MempoolTxCount == nil {
		var ret int32
		return ret
	}
	return *o.MempoolTxCount
}

// GetMempoolTxCountOk returns a tuple with the MempoolTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetMempoolTxCountOk() (*int32, bool) {
	if o == nil || o.MempoolTxCount == nil {
		return nil, false
	}
	return o.MempoolTxCount, true
}

// HasMempoolTxCount returns a boolean if a field has been set.
func (o *BlockchainInfo) HasMempoolTxCount() bool {
	if o != nil && o.MempoolTxCount != nil {
		return true
	}

	return false
}

// SetMempoolTxCount gets a reference to the given int32 and assigns it to the MempoolTxCount field.
func (o *BlockchainInfo) SetMempoolTxCount(v int32) {
	o.MempoolTxCount = &v
}

// GetMempoolUsage returns the MempoolUsage field value if set, zero value otherwise.
func (o *BlockchainInfo) GetMempoolUsage() int64 {
	if o == nil || o.MempoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.MempoolUsage
}

// GetMempoolUsageOk returns a tuple with the MempoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetMempoolUsageOk() (*int64, bool) {
	if o == nil || o.MempoolUsage == nil {
		return nil, false
	}
	return o.MempoolUsage, true
}

// HasMempoolUsage returns a boolean if a field has been set.
func (o *BlockchainInfo) HasMempoolUsage() bool {
	if o != nil && o.MempoolUsage != nil {
		return true
	}

	return false
}

// SetMempoolUsage gets a reference to the given int64 and assigns it to the MempoolUsage field.
func (o *BlockchainInfo) SetMempoolUsage(v int64) {
	o.MempoolUsage = &v
}

// GetEstimatedBlockSize returns the EstimatedBlockSize field value if set, zero value otherwise.
func (o *BlockchainInfo) GetEstimatedBlockSize() int64 {
	if o == nil || o.EstimatedBlockSize == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedBlockSize
}

// GetEstimatedBlockSizeOk returns a tuple with the EstimatedBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainInfo) GetEstimatedBlockSizeOk() (*int64, bool) {
	if o == nil || o.EstimatedBlockSize == nil {
		return nil, false
	}
	return o.EstimatedBlockSize, true
}

// HasEstimatedBlockSize returns a boolean if a field has been set.
func (o *BlockchainInfo) HasEstimatedBlockSize() bool {
	if o != nil && o.EstimatedBlockSize != nil {
		return true
	}

	return false
}

// SetEstimatedBlockSize gets a reference to the given int64 and assigns it to the EstimatedBlockSize field.
func (o *BlockchainInfo) SetEstimatedBlockSize(v int64) {
	o.EstimatedBlockSize = &v
}

func (o BlockchainInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	if o.Blocks != nil {
		toSerialize["blocks"] = o.Blocks
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.BestBlockHash != nil {
		toSerialize["bestBlockHash"] = o.BestBlockHash
	}
	if o.Difficulty != nil {
		toSerialize["difficulty"] = o.Difficulty
	}
	if o.MedianTime != nil {
		toSerialize["medianTime"] = o.MedianTime
	}
	if o.Chainwork != nil {
		toSerialize["chainwork"] = o.Chainwork
	}
	if o.NetworkHashPerSecond != nil {
		toSerialize["networkHashPerSecond"] = o.NetworkHashPerSecond
	}
	if o.MempoolTxCount != nil {
		toSerialize["mempoolTxCount"] = o.MempoolTxCount
	}
	if o.MempoolUsage != nil {
		toSerialize["mempoolUsage"] = o.MempoolUsage
	}
	if o.EstimatedBlockSize != nil {
		toSerialize["estimatedBlockSize"] = o.EstimatedBlockSize
	}
	return json.Marshal(toSerialize)
}

type NullableBlockchainInfo struct {
	value *BlockchainInfo
	isSet bool
}

func (v NullableBlockchainInfo) Get() *BlockchainInfo {
	return v.value
}

func (v *NullableBlockchainInfo) Set(val *BlockchainInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainInfo(val *BlockchainInfo) *NullableBlockchainInfo {
	return &NullableBlockchainInfo{value: val, isSet: true}
}

func (v NullableBlockchainInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

# TreasuryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | current treasury utxo txid | [optional] 
**Index** | Pointer to **int32** | current treasury utxo index | [optional] 
**Amount** | Pointer to **int64** | current treasury amount | [optional] 
**Height** | Pointer to **int64** | current treasury utxo height | [optional] 
**BlockHash** | Pointer to **string** | current treasury utxo block hash | [optional] 
**Timestamp** | Pointer to **int64** | current treasury utxo timestamp | [optional] 

## Methods

### NewTreasuryInfo

`func NewTreasuryInfo() *TreasuryInfo`

NewTreasuryInfo instantiates a new TreasuryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreasuryInfoWithDefaults

`func NewTreasuryInfoWithDefaults() *TreasuryInfo`

NewTreasuryInfoWithDefaults instantiates a new TreasuryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *TreasuryInfo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *TreasuryInfo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *TreasuryInfo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *TreasuryInfo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetIndex

`func (o *TreasuryInfo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TreasuryInfo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TreasuryInfo) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TreasuryInfo) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAmount

`func (o *TreasuryInfo) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TreasuryInfo) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TreasuryInfo) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TreasuryInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetHeight

`func (o *TreasuryInfo) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TreasuryInfo) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TreasuryInfo) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TreasuryInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBlockHash

`func (o *TreasuryInfo) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TreasuryInfo) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TreasuryInfo) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TreasuryInfo) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetTimestamp

`func (o *TreasuryInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TreasuryInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TreasuryInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TreasuryInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



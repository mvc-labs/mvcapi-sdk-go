# ContractNftAuctionUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**BidMvcPrice** | Pointer to **int64** | bidMvcPrice | [optional] 
**BidTimestamp** | Pointer to **int64** | bidTimestamp | [optional] 
**BidderAddressPkh** | Pointer to **string** | bidderAddressPkh | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**EndTimestamp** | Pointer to **int64** | endTimestamp | [optional] 
**FeeAddressPkh** | Pointer to **string** | feeAddressPkh | [optional] 
**FeeAmount** | Pointer to **int64** | feeAmount | [optional] 
**FeeRate** | Pointer to **int32** | feeRate | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**NftCodeHash** | Pointer to **string** | nftCodeHash | [optional] 
**NftId** | Pointer to **string** | nftId of the auctioning fnt | [optional] 
**SenderAddressPkh** | Pointer to **string** | senderAddressPkh | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the utxo | [optional] 
**StartMvcPrice** | Pointer to **int64** | startMvcPrice | [optional] 
**Satoshi** | Pointer to **int64** | Mvc value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Mvc value of the utxo(In string format) | [optional] 
**ContractAddress** | Pointer to **string** | The hash160 of the script(p2ch address) | [optional] 
**IsReady** | Pointer to **bool** | this the nft send to this contract address | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewContractNftAuctionUtxo

`func NewContractNftAuctionUtxo() *ContractNftAuctionUtxo`

NewContractNftAuctionUtxo instantiates a new ContractNftAuctionUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractNftAuctionUtxoWithDefaults

`func NewContractNftAuctionUtxoWithDefaults() *ContractNftAuctionUtxo`

NewContractNftAuctionUtxoWithDefaults instantiates a new ContractNftAuctionUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractNftAuctionUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractNftAuctionUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractNftAuctionUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractNftAuctionUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxid

`func (o *ContractNftAuctionUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ContractNftAuctionUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ContractNftAuctionUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ContractNftAuctionUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *ContractNftAuctionUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *ContractNftAuctionUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *ContractNftAuctionUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *ContractNftAuctionUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetBidMvcPrice

`func (o *ContractNftAuctionUtxo) GetBidMvcPrice() int64`

GetBidMvcPrice returns the BidMvcPrice field if non-nil, zero value otherwise.

### GetBidMvcPriceOk

`func (o *ContractNftAuctionUtxo) GetBidMvcPriceOk() (*int64, bool)`

GetBidMvcPriceOk returns a tuple with the BidMvcPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMvcPrice

`func (o *ContractNftAuctionUtxo) SetBidMvcPrice(v int64)`

SetBidMvcPrice sets BidMvcPrice field to given value.

### HasBidMvcPrice

`func (o *ContractNftAuctionUtxo) HasBidMvcPrice() bool`

HasBidMvcPrice returns a boolean if a field has been set.

### GetBidTimestamp

`func (o *ContractNftAuctionUtxo) GetBidTimestamp() int64`

GetBidTimestamp returns the BidTimestamp field if non-nil, zero value otherwise.

### GetBidTimestampOk

`func (o *ContractNftAuctionUtxo) GetBidTimestampOk() (*int64, bool)`

GetBidTimestampOk returns a tuple with the BidTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidTimestamp

`func (o *ContractNftAuctionUtxo) SetBidTimestamp(v int64)`

SetBidTimestamp sets BidTimestamp field to given value.

### HasBidTimestamp

`func (o *ContractNftAuctionUtxo) HasBidTimestamp() bool`

HasBidTimestamp returns a boolean if a field has been set.

### GetBidderAddressPkh

`func (o *ContractNftAuctionUtxo) GetBidderAddressPkh() string`

GetBidderAddressPkh returns the BidderAddressPkh field if non-nil, zero value otherwise.

### GetBidderAddressPkhOk

`func (o *ContractNftAuctionUtxo) GetBidderAddressPkhOk() (*string, bool)`

GetBidderAddressPkhOk returns a tuple with the BidderAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidderAddressPkh

`func (o *ContractNftAuctionUtxo) SetBidderAddressPkh(v string)`

SetBidderAddressPkh sets BidderAddressPkh field to given value.

### HasBidderAddressPkh

`func (o *ContractNftAuctionUtxo) HasBidderAddressPkh() bool`

HasBidderAddressPkh returns a boolean if a field has been set.

### GetCodeHash

`func (o *ContractNftAuctionUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractNftAuctionUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractNftAuctionUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractNftAuctionUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractNftAuctionUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractNftAuctionUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractNftAuctionUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractNftAuctionUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetEndTimestamp

`func (o *ContractNftAuctionUtxo) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *ContractNftAuctionUtxo) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *ContractNftAuctionUtxo) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *ContractNftAuctionUtxo) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetFeeAddressPkh

`func (o *ContractNftAuctionUtxo) GetFeeAddressPkh() string`

GetFeeAddressPkh returns the FeeAddressPkh field if non-nil, zero value otherwise.

### GetFeeAddressPkhOk

`func (o *ContractNftAuctionUtxo) GetFeeAddressPkhOk() (*string, bool)`

GetFeeAddressPkhOk returns a tuple with the FeeAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddressPkh

`func (o *ContractNftAuctionUtxo) SetFeeAddressPkh(v string)`

SetFeeAddressPkh sets FeeAddressPkh field to given value.

### HasFeeAddressPkh

`func (o *ContractNftAuctionUtxo) HasFeeAddressPkh() bool`

HasFeeAddressPkh returns a boolean if a field has been set.

### GetFeeAmount

`func (o *ContractNftAuctionUtxo) GetFeeAmount() int64`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *ContractNftAuctionUtxo) GetFeeAmountOk() (*int64, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *ContractNftAuctionUtxo) SetFeeAmount(v int64)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *ContractNftAuctionUtxo) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeRate

`func (o *ContractNftAuctionUtxo) GetFeeRate() int32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ContractNftAuctionUtxo) GetFeeRateOk() (*int32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ContractNftAuctionUtxo) SetFeeRate(v int32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *ContractNftAuctionUtxo) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetHeight

`func (o *ContractNftAuctionUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContractNftAuctionUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContractNftAuctionUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContractNftAuctionUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetNftCodeHash

`func (o *ContractNftAuctionUtxo) GetNftCodeHash() string`

GetNftCodeHash returns the NftCodeHash field if non-nil, zero value otherwise.

### GetNftCodeHashOk

`func (o *ContractNftAuctionUtxo) GetNftCodeHashOk() (*string, bool)`

GetNftCodeHashOk returns a tuple with the NftCodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftCodeHash

`func (o *ContractNftAuctionUtxo) SetNftCodeHash(v string)`

SetNftCodeHash sets NftCodeHash field to given value.

### HasNftCodeHash

`func (o *ContractNftAuctionUtxo) HasNftCodeHash() bool`

HasNftCodeHash returns a boolean if a field has been set.

### GetNftId

`func (o *ContractNftAuctionUtxo) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *ContractNftAuctionUtxo) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *ContractNftAuctionUtxo) SetNftId(v string)`

SetNftId sets NftId field to given value.

### HasNftId

`func (o *ContractNftAuctionUtxo) HasNftId() bool`

HasNftId returns a boolean if a field has been set.

### GetSenderAddressPkh

`func (o *ContractNftAuctionUtxo) GetSenderAddressPkh() string`

GetSenderAddressPkh returns the SenderAddressPkh field if non-nil, zero value otherwise.

### GetSenderAddressPkhOk

`func (o *ContractNftAuctionUtxo) GetSenderAddressPkhOk() (*string, bool)`

GetSenderAddressPkhOk returns a tuple with the SenderAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddressPkh

`func (o *ContractNftAuctionUtxo) SetSenderAddressPkh(v string)`

SetSenderAddressPkh sets SenderAddressPkh field to given value.

### HasSenderAddressPkh

`func (o *ContractNftAuctionUtxo) HasSenderAddressPkh() bool`

HasSenderAddressPkh returns a boolean if a field has been set.

### GetSensibleId

`func (o *ContractNftAuctionUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *ContractNftAuctionUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *ContractNftAuctionUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *ContractNftAuctionUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetStartMvcPrice

`func (o *ContractNftAuctionUtxo) GetStartMvcPrice() int64`

GetStartMvcPrice returns the StartMvcPrice field if non-nil, zero value otherwise.

### GetStartMvcPriceOk

`func (o *ContractNftAuctionUtxo) GetStartMvcPriceOk() (*int64, bool)`

GetStartMvcPriceOk returns a tuple with the StartMvcPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMvcPrice

`func (o *ContractNftAuctionUtxo) SetStartMvcPrice(v int64)`

SetStartMvcPrice sets StartMvcPrice field to given value.

### HasStartMvcPrice

`func (o *ContractNftAuctionUtxo) HasStartMvcPrice() bool`

HasStartMvcPrice returns a boolean if a field has been set.

### GetSatoshi

`func (o *ContractNftAuctionUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *ContractNftAuctionUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *ContractNftAuctionUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *ContractNftAuctionUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *ContractNftAuctionUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *ContractNftAuctionUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *ContractNftAuctionUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *ContractNftAuctionUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetContractAddress

`func (o *ContractNftAuctionUtxo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ContractNftAuctionUtxo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ContractNftAuctionUtxo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *ContractNftAuctionUtxo) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetIsReady

`func (o *ContractNftAuctionUtxo) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *ContractNftAuctionUtxo) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *ContractNftAuctionUtxo) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *ContractNftAuctionUtxo) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetFlag

`func (o *ContractNftAuctionUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ContractNftAuctionUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ContractNftAuctionUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ContractNftAuctionUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



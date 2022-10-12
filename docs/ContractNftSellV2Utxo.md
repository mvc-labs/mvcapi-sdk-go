# ContractNftSellV2Utxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**ContractAddress** | Pointer to **string** | Address calculated from contract hash(p2ch). | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**TokenIndex** | Pointer to **int64** | The index of this NFT. | [optional] 
**Price** | Pointer to **int64** | the price of nft. | [optional] 
**FeeAddressPkh** | Pointer to **string** | The address to receive fees. | [optional] 
**FeeRate** | Pointer to **int32** | feeRate | [optional] 
**NftId** | Pointer to **string** | nftId | [optional] 
**SellerAddressPkh** | Pointer to **string** | The address pkh of seller | [optional] 
**Satoshi** | Pointer to **int64** | Mvc value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Mvc value of the utxo(In string format) | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**IsReady** | Pointer to **bool** | Is current nft transfered into sell contract, If not ready, the following fields will be null | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**MetaTxid** | Pointer to **string** | The metanet tx describing the nft. | [optional] 
**MetaOutputIndex** | Pointer to **int32** | Symbol of the token. | [optional] 
**TokenSupply** | Pointer to **int64** | The total supply of this NFT. | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewContractNftSellV2Utxo

`func NewContractNftSellV2Utxo() *ContractNftSellV2Utxo`

NewContractNftSellV2Utxo instantiates a new ContractNftSellV2Utxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractNftSellV2UtxoWithDefaults

`func NewContractNftSellV2UtxoWithDefaults() *ContractNftSellV2Utxo`

NewContractNftSellV2UtxoWithDefaults instantiates a new ContractNftSellV2Utxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractNftSellV2Utxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractNftSellV2Utxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractNftSellV2Utxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractNftSellV2Utxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContractAddress

`func (o *ContractNftSellV2Utxo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ContractNftSellV2Utxo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ContractNftSellV2Utxo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *ContractNftSellV2Utxo) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTxid

`func (o *ContractNftSellV2Utxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ContractNftSellV2Utxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ContractNftSellV2Utxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ContractNftSellV2Utxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *ContractNftSellV2Utxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *ContractNftSellV2Utxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *ContractNftSellV2Utxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *ContractNftSellV2Utxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetCodeHash

`func (o *ContractNftSellV2Utxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractNftSellV2Utxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractNftSellV2Utxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractNftSellV2Utxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractNftSellV2Utxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractNftSellV2Utxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractNftSellV2Utxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractNftSellV2Utxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetTokenIndex

`func (o *ContractNftSellV2Utxo) GetTokenIndex() int64`

GetTokenIndex returns the TokenIndex field if non-nil, zero value otherwise.

### GetTokenIndexOk

`func (o *ContractNftSellV2Utxo) GetTokenIndexOk() (*int64, bool)`

GetTokenIndexOk returns a tuple with the TokenIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIndex

`func (o *ContractNftSellV2Utxo) SetTokenIndex(v int64)`

SetTokenIndex sets TokenIndex field to given value.

### HasTokenIndex

`func (o *ContractNftSellV2Utxo) HasTokenIndex() bool`

HasTokenIndex returns a boolean if a field has been set.

### GetPrice

`func (o *ContractNftSellV2Utxo) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ContractNftSellV2Utxo) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ContractNftSellV2Utxo) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ContractNftSellV2Utxo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeeAddressPkh

`func (o *ContractNftSellV2Utxo) GetFeeAddressPkh() string`

GetFeeAddressPkh returns the FeeAddressPkh field if non-nil, zero value otherwise.

### GetFeeAddressPkhOk

`func (o *ContractNftSellV2Utxo) GetFeeAddressPkhOk() (*string, bool)`

GetFeeAddressPkhOk returns a tuple with the FeeAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddressPkh

`func (o *ContractNftSellV2Utxo) SetFeeAddressPkh(v string)`

SetFeeAddressPkh sets FeeAddressPkh field to given value.

### HasFeeAddressPkh

`func (o *ContractNftSellV2Utxo) HasFeeAddressPkh() bool`

HasFeeAddressPkh returns a boolean if a field has been set.

### GetFeeRate

`func (o *ContractNftSellV2Utxo) GetFeeRate() int32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ContractNftSellV2Utxo) GetFeeRateOk() (*int32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ContractNftSellV2Utxo) SetFeeRate(v int32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *ContractNftSellV2Utxo) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetNftId

`func (o *ContractNftSellV2Utxo) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *ContractNftSellV2Utxo) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *ContractNftSellV2Utxo) SetNftId(v string)`

SetNftId sets NftId field to given value.

### HasNftId

`func (o *ContractNftSellV2Utxo) HasNftId() bool`

HasNftId returns a boolean if a field has been set.

### GetSellerAddressPkh

`func (o *ContractNftSellV2Utxo) GetSellerAddressPkh() string`

GetSellerAddressPkh returns the SellerAddressPkh field if non-nil, zero value otherwise.

### GetSellerAddressPkhOk

`func (o *ContractNftSellV2Utxo) GetSellerAddressPkhOk() (*string, bool)`

GetSellerAddressPkhOk returns a tuple with the SellerAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAddressPkh

`func (o *ContractNftSellV2Utxo) SetSellerAddressPkh(v string)`

SetSellerAddressPkh sets SellerAddressPkh field to given value.

### HasSellerAddressPkh

`func (o *ContractNftSellV2Utxo) HasSellerAddressPkh() bool`

HasSellerAddressPkh returns a boolean if a field has been set.

### GetSatoshi

`func (o *ContractNftSellV2Utxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *ContractNftSellV2Utxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *ContractNftSellV2Utxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *ContractNftSellV2Utxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *ContractNftSellV2Utxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *ContractNftSellV2Utxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *ContractNftSellV2Utxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *ContractNftSellV2Utxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetHeight

`func (o *ContractNftSellV2Utxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContractNftSellV2Utxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContractNftSellV2Utxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContractNftSellV2Utxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIsReady

`func (o *ContractNftSellV2Utxo) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *ContractNftSellV2Utxo) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *ContractNftSellV2Utxo) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *ContractNftSellV2Utxo) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetSensibleId

`func (o *ContractNftSellV2Utxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *ContractNftSellV2Utxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *ContractNftSellV2Utxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *ContractNftSellV2Utxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetMetaTxid

`func (o *ContractNftSellV2Utxo) GetMetaTxid() string`

GetMetaTxid returns the MetaTxid field if non-nil, zero value otherwise.

### GetMetaTxidOk

`func (o *ContractNftSellV2Utxo) GetMetaTxidOk() (*string, bool)`

GetMetaTxidOk returns a tuple with the MetaTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTxid

`func (o *ContractNftSellV2Utxo) SetMetaTxid(v string)`

SetMetaTxid sets MetaTxid field to given value.

### HasMetaTxid

`func (o *ContractNftSellV2Utxo) HasMetaTxid() bool`

HasMetaTxid returns a boolean if a field has been set.

### GetMetaOutputIndex

`func (o *ContractNftSellV2Utxo) GetMetaOutputIndex() int32`

GetMetaOutputIndex returns the MetaOutputIndex field if non-nil, zero value otherwise.

### GetMetaOutputIndexOk

`func (o *ContractNftSellV2Utxo) GetMetaOutputIndexOk() (*int32, bool)`

GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOutputIndex

`func (o *ContractNftSellV2Utxo) SetMetaOutputIndex(v int32)`

SetMetaOutputIndex sets MetaOutputIndex field to given value.

### HasMetaOutputIndex

`func (o *ContractNftSellV2Utxo) HasMetaOutputIndex() bool`

HasMetaOutputIndex returns a boolean if a field has been set.

### GetTokenSupply

`func (o *ContractNftSellV2Utxo) GetTokenSupply() int64`

GetTokenSupply returns the TokenSupply field if non-nil, zero value otherwise.

### GetTokenSupplyOk

`func (o *ContractNftSellV2Utxo) GetTokenSupplyOk() (*int64, bool)`

GetTokenSupplyOk returns a tuple with the TokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSupply

`func (o *ContractNftSellV2Utxo) SetTokenSupply(v int64)`

SetTokenSupply sets TokenSupply field to given value.

### HasTokenSupply

`func (o *ContractNftSellV2Utxo) HasTokenSupply() bool`

HasTokenSupply returns a boolean if a field has been set.

### GetFlag

`func (o *ContractNftSellV2Utxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ContractNftSellV2Utxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ContractNftSellV2Utxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ContractNftSellV2Utxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



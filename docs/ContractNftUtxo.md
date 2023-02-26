# ContractNftUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**Height** | Pointer to **int64** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**MetaTxid** | Pointer to **string** | The metanet tx describing the nft. | [optional] 
**MetaOutputIndex** | Pointer to **int32** | Symbol of the token. | [optional] 
**TokenSupply** | Pointer to **int64** | The total supply of this NFT. | [optional] 
**TokenIndex** | Pointer to **int64** | The index of this NFT. | [optional] 
**Satoshi** | Pointer to **int64** | Mvc value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Mvc value of the utxo(In string format) | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewContractNftUtxo

`func NewContractNftUtxo() *ContractNftUtxo`

NewContractNftUtxo instantiates a new ContractNftUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractNftUtxoWithDefaults

`func NewContractNftUtxoWithDefaults() *ContractNftUtxo`

NewContractNftUtxoWithDefaults instantiates a new ContractNftUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractNftUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractNftUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractNftUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractNftUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxid

`func (o *ContractNftUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ContractNftUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ContractNftUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ContractNftUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *ContractNftUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *ContractNftUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *ContractNftUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *ContractNftUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetCodeHash

`func (o *ContractNftUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractNftUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractNftUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractNftUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractNftUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractNftUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractNftUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractNftUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetSensibleId

`func (o *ContractNftUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *ContractNftUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *ContractNftUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *ContractNftUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetHeight

`func (o *ContractNftUtxo) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContractNftUtxo) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContractNftUtxo) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContractNftUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMetaTxid

`func (o *ContractNftUtxo) GetMetaTxid() string`

GetMetaTxid returns the MetaTxid field if non-nil, zero value otherwise.

### GetMetaTxidOk

`func (o *ContractNftUtxo) GetMetaTxidOk() (*string, bool)`

GetMetaTxidOk returns a tuple with the MetaTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTxid

`func (o *ContractNftUtxo) SetMetaTxid(v string)`

SetMetaTxid sets MetaTxid field to given value.

### HasMetaTxid

`func (o *ContractNftUtxo) HasMetaTxid() bool`

HasMetaTxid returns a boolean if a field has been set.

### GetMetaOutputIndex

`func (o *ContractNftUtxo) GetMetaOutputIndex() int32`

GetMetaOutputIndex returns the MetaOutputIndex field if non-nil, zero value otherwise.

### GetMetaOutputIndexOk

`func (o *ContractNftUtxo) GetMetaOutputIndexOk() (*int32, bool)`

GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOutputIndex

`func (o *ContractNftUtxo) SetMetaOutputIndex(v int32)`

SetMetaOutputIndex sets MetaOutputIndex field to given value.

### HasMetaOutputIndex

`func (o *ContractNftUtxo) HasMetaOutputIndex() bool`

HasMetaOutputIndex returns a boolean if a field has been set.

### GetTokenSupply

`func (o *ContractNftUtxo) GetTokenSupply() int64`

GetTokenSupply returns the TokenSupply field if non-nil, zero value otherwise.

### GetTokenSupplyOk

`func (o *ContractNftUtxo) GetTokenSupplyOk() (*int64, bool)`

GetTokenSupplyOk returns a tuple with the TokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSupply

`func (o *ContractNftUtxo) SetTokenSupply(v int64)`

SetTokenSupply sets TokenSupply field to given value.

### HasTokenSupply

`func (o *ContractNftUtxo) HasTokenSupply() bool`

HasTokenSupply returns a boolean if a field has been set.

### GetTokenIndex

`func (o *ContractNftUtxo) GetTokenIndex() int64`

GetTokenIndex returns the TokenIndex field if non-nil, zero value otherwise.

### GetTokenIndexOk

`func (o *ContractNftUtxo) GetTokenIndexOk() (*int64, bool)`

GetTokenIndexOk returns a tuple with the TokenIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIndex

`func (o *ContractNftUtxo) SetTokenIndex(v int64)`

SetTokenIndex sets TokenIndex field to given value.

### HasTokenIndex

`func (o *ContractNftUtxo) HasTokenIndex() bool`

HasTokenIndex returns a boolean if a field has been set.

### GetSatoshi

`func (o *ContractNftUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *ContractNftUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *ContractNftUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *ContractNftUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *ContractNftUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *ContractNftUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *ContractNftUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *ContractNftUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetFlag

`func (o *ContractNftUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ContractNftUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ContractNftUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ContractNftUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



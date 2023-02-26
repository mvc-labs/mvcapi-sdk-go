# ContractFtUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**Name** | Pointer to **string** | Name of the token. | [optional] 
**Symbol** | Pointer to **string** | Symbol of the token. | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**Decimal** | Pointer to **int32** | The decimal position. | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**Value** | Pointer to **int64** | Token value of the utxo(Irrelavant to satoshi value). | [optional] 
**ValueString** | Pointer to **string** | Token value of the utxo(In string format) | [optional] 
**Satoshi** | Pointer to **int64** | Mvc value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Mvc value of the utxo(In string format) | [optional] 
**Height** | Pointer to **int64** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewContractFtUtxo

`func NewContractFtUtxo() *ContractFtUtxo`

NewContractFtUtxo instantiates a new ContractFtUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractFtUtxoWithDefaults

`func NewContractFtUtxoWithDefaults() *ContractFtUtxo`

NewContractFtUtxoWithDefaults instantiates a new ContractFtUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractFtUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractFtUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractFtUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractFtUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCodeHash

`func (o *ContractFtUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractFtUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractFtUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractFtUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractFtUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractFtUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractFtUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractFtUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetName

`func (o *ContractFtUtxo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractFtUtxo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractFtUtxo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractFtUtxo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *ContractFtUtxo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ContractFtUtxo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ContractFtUtxo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ContractFtUtxo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSensibleId

`func (o *ContractFtUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *ContractFtUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *ContractFtUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *ContractFtUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetDecimal

`func (o *ContractFtUtxo) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *ContractFtUtxo) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *ContractFtUtxo) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *ContractFtUtxo) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetTxid

`func (o *ContractFtUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ContractFtUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ContractFtUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ContractFtUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *ContractFtUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *ContractFtUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *ContractFtUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *ContractFtUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetValue

`func (o *ContractFtUtxo) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractFtUtxo) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractFtUtxo) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContractFtUtxo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueString

`func (o *ContractFtUtxo) GetValueString() string`

GetValueString returns the ValueString field if non-nil, zero value otherwise.

### GetValueStringOk

`func (o *ContractFtUtxo) GetValueStringOk() (*string, bool)`

GetValueStringOk returns a tuple with the ValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueString

`func (o *ContractFtUtxo) SetValueString(v string)`

SetValueString sets ValueString field to given value.

### HasValueString

`func (o *ContractFtUtxo) HasValueString() bool`

HasValueString returns a boolean if a field has been set.

### GetSatoshi

`func (o *ContractFtUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *ContractFtUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *ContractFtUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *ContractFtUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *ContractFtUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *ContractFtUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *ContractFtUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *ContractFtUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetHeight

`func (o *ContractFtUtxo) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContractFtUtxo) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContractFtUtxo) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContractFtUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetFlag

`func (o *ContractFtUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ContractFtUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ContractFtUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ContractFtUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



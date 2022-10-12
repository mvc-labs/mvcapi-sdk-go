# ContractFtBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeHash** | Pointer to **string** | Codehash of the token. | [optional] 
**Genesis** | Pointer to **string** | Genesis of the token. | [optional] 
**Name** | Pointer to **string** | Name of the token. | [optional] 
**Symbol** | Pointer to **string** | Symbol of the token. | [optional] 
**Decimal** | Pointer to **int32** | The decimal position. | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**UtxoCount** | Pointer to **int32** | Number of utxos for this token. | [optional] 
**Confirmed** | Pointer to **int64** | Confirmed balance of the token. | [optional] 
**ConfirmedString** | Pointer to **string** | Confirmed balance of the token(In string format). | [optional] 
**Unconfirmed** | Pointer to **int64** | Unconfirmed balance of the token. | [optional] 
**UnconfirmedString** | Pointer to **string** | Unconfirmed balance of the token(In string format). | [optional] 

## Methods

### NewContractFtBalance

`func NewContractFtBalance() *ContractFtBalance`

NewContractFtBalance instantiates a new ContractFtBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractFtBalanceWithDefaults

`func NewContractFtBalanceWithDefaults() *ContractFtBalance`

NewContractFtBalanceWithDefaults instantiates a new ContractFtBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeHash

`func (o *ContractFtBalance) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractFtBalance) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractFtBalance) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractFtBalance) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractFtBalance) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractFtBalance) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractFtBalance) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractFtBalance) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetName

`func (o *ContractFtBalance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractFtBalance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractFtBalance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractFtBalance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *ContractFtBalance) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ContractFtBalance) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ContractFtBalance) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ContractFtBalance) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimal

`func (o *ContractFtBalance) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *ContractFtBalance) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *ContractFtBalance) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *ContractFtBalance) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetSensibleId

`func (o *ContractFtBalance) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *ContractFtBalance) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *ContractFtBalance) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *ContractFtBalance) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetUtxoCount

`func (o *ContractFtBalance) GetUtxoCount() int32`

GetUtxoCount returns the UtxoCount field if non-nil, zero value otherwise.

### GetUtxoCountOk

`func (o *ContractFtBalance) GetUtxoCountOk() (*int32, bool)`

GetUtxoCountOk returns a tuple with the UtxoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoCount

`func (o *ContractFtBalance) SetUtxoCount(v int32)`

SetUtxoCount sets UtxoCount field to given value.

### HasUtxoCount

`func (o *ContractFtBalance) HasUtxoCount() bool`

HasUtxoCount returns a boolean if a field has been set.

### GetConfirmed

`func (o *ContractFtBalance) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *ContractFtBalance) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *ContractFtBalance) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *ContractFtBalance) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetConfirmedString

`func (o *ContractFtBalance) GetConfirmedString() string`

GetConfirmedString returns the ConfirmedString field if non-nil, zero value otherwise.

### GetConfirmedStringOk

`func (o *ContractFtBalance) GetConfirmedStringOk() (*string, bool)`

GetConfirmedStringOk returns a tuple with the ConfirmedString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedString

`func (o *ContractFtBalance) SetConfirmedString(v string)`

SetConfirmedString sets ConfirmedString field to given value.

### HasConfirmedString

`func (o *ContractFtBalance) HasConfirmedString() bool`

HasConfirmedString returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *ContractFtBalance) GetUnconfirmed() int64`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *ContractFtBalance) GetUnconfirmedOk() (*int64, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *ContractFtBalance) SetUnconfirmed(v int64)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *ContractFtBalance) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.

### GetUnconfirmedString

`func (o *ContractFtBalance) GetUnconfirmedString() string`

GetUnconfirmedString returns the UnconfirmedString field if non-nil, zero value otherwise.

### GetUnconfirmedStringOk

`func (o *ContractFtBalance) GetUnconfirmedStringOk() (*string, bool)`

GetUnconfirmedStringOk returns a tuple with the UnconfirmedString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmedString

`func (o *ContractFtBalance) SetUnconfirmedString(v string)`

SetUnconfirmedString sets UnconfirmedString field to given value.

### HasUnconfirmedString

`func (o *ContractFtBalance) HasUnconfirmedString() bool`

HasUnconfirmedString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



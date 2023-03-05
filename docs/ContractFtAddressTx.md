# ContractFtAddressTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | Pointer to **string** | Paging flag | [optional] 
**Address** | Pointer to **string** | The address of the record | [optional] 
**CodeHash** | Pointer to **string** | Specific codeHash | [optional] 
**Genesis** | Pointer to **string** | specific genesis | [optional] 
**Time** | Pointer to **int64** | timestamp of the tx | [optional] 
**Height** | Pointer to **int64** | Block Height of the tx, -1 if not confirmed | [optional] 
**Income** | Pointer to **int64** | total income of the address from this tx | [optional] 
**Outcome** | Pointer to **int64** | total outcome of the address from this tx | [optional] 
**Txid** | Pointer to **string** | txid of this record | [optional] 

## Methods

### NewContractFtAddressTx

`func NewContractFtAddressTx() *ContractFtAddressTx`

NewContractFtAddressTx instantiates a new ContractFtAddressTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractFtAddressTxWithDefaults

`func NewContractFtAddressTxWithDefaults() *ContractFtAddressTx`

NewContractFtAddressTxWithDefaults instantiates a new ContractFtAddressTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *ContractFtAddressTx) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ContractFtAddressTx) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ContractFtAddressTx) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ContractFtAddressTx) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetAddress

`func (o *ContractFtAddressTx) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractFtAddressTx) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractFtAddressTx) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractFtAddressTx) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCodeHash

`func (o *ContractFtAddressTx) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractFtAddressTx) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractFtAddressTx) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *ContractFtAddressTx) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *ContractFtAddressTx) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *ContractFtAddressTx) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *ContractFtAddressTx) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *ContractFtAddressTx) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetTime

`func (o *ContractFtAddressTx) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ContractFtAddressTx) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ContractFtAddressTx) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ContractFtAddressTx) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetHeight

`func (o *ContractFtAddressTx) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ContractFtAddressTx) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ContractFtAddressTx) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ContractFtAddressTx) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIncome

`func (o *ContractFtAddressTx) GetIncome() int64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *ContractFtAddressTx) GetIncomeOk() (*int64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *ContractFtAddressTx) SetIncome(v int64)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *ContractFtAddressTx) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetOutcome

`func (o *ContractFtAddressTx) GetOutcome() int64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ContractFtAddressTx) GetOutcomeOk() (*int64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ContractFtAddressTx) SetOutcome(v int64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ContractFtAddressTx) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetTxid

`func (o *ContractFtAddressTx) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ContractFtAddressTx) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ContractFtAddressTx) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ContractFtAddressTx) HasTxid() bool`

HasTxid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



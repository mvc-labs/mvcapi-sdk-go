# ContractFtGenesisCirculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **int64** | Confirmed circulation of the token(Sum of confirmed amount). | [optional] 
**Unconfirmed** | Pointer to **int64** | Unconfirmed circulation of the token(Unconfirmed new token - Unconfirmed spent). | [optional] 

## Methods

### NewContractFtGenesisCirculation

`func NewContractFtGenesisCirculation() *ContractFtGenesisCirculation`

NewContractFtGenesisCirculation instantiates a new ContractFtGenesisCirculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractFtGenesisCirculationWithDefaults

`func NewContractFtGenesisCirculationWithDefaults() *ContractFtGenesisCirculation`

NewContractFtGenesisCirculationWithDefaults instantiates a new ContractFtGenesisCirculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *ContractFtGenesisCirculation) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *ContractFtGenesisCirculation) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *ContractFtGenesisCirculation) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *ContractFtGenesisCirculation) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *ContractFtGenesisCirculation) GetUnconfirmed() int64`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *ContractFtGenesisCirculation) GetUnconfirmedOk() (*int64, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *ContractFtGenesisCirculation) SetUnconfirmed(v int64)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *ContractFtGenesisCirculation) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



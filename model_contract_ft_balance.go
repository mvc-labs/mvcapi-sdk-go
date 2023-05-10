/*
 * MicrovisionChain API Document
 *
 * API definition for MicrovisionChain provided apis
 *
 * API version: 3.0.9
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mvcapi

import (
	"encoding/json"
)

// ContractFtBalance Contract fungible token balance
type ContractFtBalance struct {
	// Codehash of the token.
	CodeHash *string `json:"codeHash,omitempty"`
	// Genesis of the token.
	Genesis *string `json:"genesis,omitempty"`
	// Name of the token.
	Name *string `json:"name,omitempty"`
	// Symbol of the token.
	Symbol *string `json:"symbol,omitempty"`
	// The decimal position.
	Decimal *int32 `json:"decimal,omitempty"`
	// SensibleId of the token
	SensibleId *string `json:"sensibleId,omitempty"`
	// Number of utxos for this token.
	UtxoCount *int32 `json:"utxoCount,omitempty"`
	// Confirmed balance of the token.
	Confirmed *int64 `json:"confirmed,omitempty"`
	// Confirmed balance of the token(In string format).
	ConfirmedString *string `json:"confirmedString,omitempty"`
	// Unconfirmed balance of the token.
	Unconfirmed *int64 `json:"unconfirmed,omitempty"`
	// Unconfirmed balance of the token(In string format).
	UnconfirmedString *string `json:"unconfirmedString,omitempty"`
}

// NewContractFtBalance instantiates a new ContractFtBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractFtBalance() *ContractFtBalance {
	this := ContractFtBalance{}
	return &this
}

// NewContractFtBalanceWithDefaults instantiates a new ContractFtBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractFtBalanceWithDefaults() *ContractFtBalance {
	this := ContractFtBalance{}
	return &this
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *ContractFtBalance) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *ContractFtBalance) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *ContractFtBalance) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *ContractFtBalance) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *ContractFtBalance) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *ContractFtBalance) SetGenesis(v string) {
	o.Genesis = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContractFtBalance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContractFtBalance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContractFtBalance) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ContractFtBalance) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ContractFtBalance) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ContractFtBalance) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDecimal returns the Decimal field value if set, zero value otherwise.
func (o *ContractFtBalance) GetDecimal() int32 {
	if o == nil || o.Decimal == nil {
		var ret int32
		return ret
	}
	return *o.Decimal
}

// GetDecimalOk returns a tuple with the Decimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetDecimalOk() (*int32, bool) {
	if o == nil || o.Decimal == nil {
		return nil, false
	}
	return o.Decimal, true
}

// HasDecimal returns a boolean if a field has been set.
func (o *ContractFtBalance) HasDecimal() bool {
	if o != nil && o.Decimal != nil {
		return true
	}

	return false
}

// SetDecimal gets a reference to the given int32 and assigns it to the Decimal field.
func (o *ContractFtBalance) SetDecimal(v int32) {
	o.Decimal = &v
}

// GetSensibleId returns the SensibleId field value if set, zero value otherwise.
func (o *ContractFtBalance) GetSensibleId() string {
	if o == nil || o.SensibleId == nil {
		var ret string
		return ret
	}
	return *o.SensibleId
}

// GetSensibleIdOk returns a tuple with the SensibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetSensibleIdOk() (*string, bool) {
	if o == nil || o.SensibleId == nil {
		return nil, false
	}
	return o.SensibleId, true
}

// HasSensibleId returns a boolean if a field has been set.
func (o *ContractFtBalance) HasSensibleId() bool {
	if o != nil && o.SensibleId != nil {
		return true
	}

	return false
}

// SetSensibleId gets a reference to the given string and assigns it to the SensibleId field.
func (o *ContractFtBalance) SetSensibleId(v string) {
	o.SensibleId = &v
}

// GetUtxoCount returns the UtxoCount field value if set, zero value otherwise.
func (o *ContractFtBalance) GetUtxoCount() int32 {
	if o == nil || o.UtxoCount == nil {
		var ret int32
		return ret
	}
	return *o.UtxoCount
}

// GetUtxoCountOk returns a tuple with the UtxoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetUtxoCountOk() (*int32, bool) {
	if o == nil || o.UtxoCount == nil {
		return nil, false
	}
	return o.UtxoCount, true
}

// HasUtxoCount returns a boolean if a field has been set.
func (o *ContractFtBalance) HasUtxoCount() bool {
	if o != nil && o.UtxoCount != nil {
		return true
	}

	return false
}

// SetUtxoCount gets a reference to the given int32 and assigns it to the UtxoCount field.
func (o *ContractFtBalance) SetUtxoCount(v int32) {
	o.UtxoCount = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *ContractFtBalance) GetConfirmed() int64 {
	if o == nil || o.Confirmed == nil {
		var ret int64
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetConfirmedOk() (*int64, bool) {
	if o == nil || o.Confirmed == nil {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *ContractFtBalance) HasConfirmed() bool {
	if o != nil && o.Confirmed != nil {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given int64 and assigns it to the Confirmed field.
func (o *ContractFtBalance) SetConfirmed(v int64) {
	o.Confirmed = &v
}

// GetConfirmedString returns the ConfirmedString field value if set, zero value otherwise.
func (o *ContractFtBalance) GetConfirmedString() string {
	if o == nil || o.ConfirmedString == nil {
		var ret string
		return ret
	}
	return *o.ConfirmedString
}

// GetConfirmedStringOk returns a tuple with the ConfirmedString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetConfirmedStringOk() (*string, bool) {
	if o == nil || o.ConfirmedString == nil {
		return nil, false
	}
	return o.ConfirmedString, true
}

// HasConfirmedString returns a boolean if a field has been set.
func (o *ContractFtBalance) HasConfirmedString() bool {
	if o != nil && o.ConfirmedString != nil {
		return true
	}

	return false
}

// SetConfirmedString gets a reference to the given string and assigns it to the ConfirmedString field.
func (o *ContractFtBalance) SetConfirmedString(v string) {
	o.ConfirmedString = &v
}

// GetUnconfirmed returns the Unconfirmed field value if set, zero value otherwise.
func (o *ContractFtBalance) GetUnconfirmed() int64 {
	if o == nil || o.Unconfirmed == nil {
		var ret int64
		return ret
	}
	return *o.Unconfirmed
}

// GetUnconfirmedOk returns a tuple with the Unconfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetUnconfirmedOk() (*int64, bool) {
	if o == nil || o.Unconfirmed == nil {
		return nil, false
	}
	return o.Unconfirmed, true
}

// HasUnconfirmed returns a boolean if a field has been set.
func (o *ContractFtBalance) HasUnconfirmed() bool {
	if o != nil && o.Unconfirmed != nil {
		return true
	}

	return false
}

// SetUnconfirmed gets a reference to the given int64 and assigns it to the Unconfirmed field.
func (o *ContractFtBalance) SetUnconfirmed(v int64) {
	o.Unconfirmed = &v
}

// GetUnconfirmedString returns the UnconfirmedString field value if set, zero value otherwise.
func (o *ContractFtBalance) GetUnconfirmedString() string {
	if o == nil || o.UnconfirmedString == nil {
		var ret string
		return ret
	}
	return *o.UnconfirmedString
}

// GetUnconfirmedStringOk returns a tuple with the UnconfirmedString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtBalance) GetUnconfirmedStringOk() (*string, bool) {
	if o == nil || o.UnconfirmedString == nil {
		return nil, false
	}
	return o.UnconfirmedString, true
}

// HasUnconfirmedString returns a boolean if a field has been set.
func (o *ContractFtBalance) HasUnconfirmedString() bool {
	if o != nil && o.UnconfirmedString != nil {
		return true
	}

	return false
}

// SetUnconfirmedString gets a reference to the given string and assigns it to the UnconfirmedString field.
func (o *ContractFtBalance) SetUnconfirmedString(v string) {
	o.UnconfirmedString = &v
}

func (o ContractFtBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CodeHash != nil {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.Genesis != nil {
		toSerialize["genesis"] = o.Genesis
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Decimal != nil {
		toSerialize["decimal"] = o.Decimal
	}
	if o.SensibleId != nil {
		toSerialize["sensibleId"] = o.SensibleId
	}
	if o.UtxoCount != nil {
		toSerialize["utxoCount"] = o.UtxoCount
	}
	if o.Confirmed != nil {
		toSerialize["confirmed"] = o.Confirmed
	}
	if o.ConfirmedString != nil {
		toSerialize["confirmedString"] = o.ConfirmedString
	}
	if o.Unconfirmed != nil {
		toSerialize["unconfirmed"] = o.Unconfirmed
	}
	if o.UnconfirmedString != nil {
		toSerialize["unconfirmedString"] = o.UnconfirmedString
	}
	return json.Marshal(toSerialize)
}

type NullableContractFtBalance struct {
	value *ContractFtBalance
	isSet bool
}

func (v NullableContractFtBalance) Get() *ContractFtBalance {
	return v.value
}

func (v *NullableContractFtBalance) Set(val *ContractFtBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableContractFtBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableContractFtBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractFtBalance(val *ContractFtBalance) *NullableContractFtBalance {
	return &NullableContractFtBalance{value: val, isSet: true}
}

func (v NullableContractFtBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractFtBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

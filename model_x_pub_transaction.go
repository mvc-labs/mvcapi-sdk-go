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

// XPubTransaction Xpub transaction history
type XPubTransaction struct {
	// query xpub
	Xpub *string `json:"xpub,omitempty"`
	// Txid for this transaction.
	Txid *string `json:"txid,omitempty"`
	// Max lookahead receive index when processing this transaction.
	MaxReceiveIndex *int32 `json:"maxReceiveIndex,omitempty"`
	// Max lookahead change index when processing this transaction.
	MaxChangeIndex *int32 `json:"maxChangeIndex,omitempty"`
	// Total received satoshis(Including all address)
	Income *int64 `json:"income,omitempty"`
	// Total spent satoshis(Including all address)
	Outcome *int64 `json:"outcome,omitempty"`
	// Height for this transaction. -1 for unconfirmed
	Height *int64 `json:"height,omitempty"`
	// Block index for this transaction, -1 for unconfirmed
	BlockIndex *int32 `json:"blockIndex,omitempty"`
	// Block timestamp for this transaction, if unconfirmed, the time is first seen time.
	BlockTime *int64 `json:"blockTime,omitempty"`
	// Paging flag, format blockTimestamp_blockIndex
	Flag *string `json:"flag,omitempty"`
}

// NewXPubTransaction instantiates a new XPubTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXPubTransaction() *XPubTransaction {
	this := XPubTransaction{}
	return &this
}

// NewXPubTransactionWithDefaults instantiates a new XPubTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXPubTransactionWithDefaults() *XPubTransaction {
	this := XPubTransaction{}
	return &this
}

// GetXpub returns the Xpub field value if set, zero value otherwise.
func (o *XPubTransaction) GetXpub() string {
	if o == nil || o.Xpub == nil {
		var ret string
		return ret
	}
	return *o.Xpub
}

// GetXpubOk returns a tuple with the Xpub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetXpubOk() (*string, bool) {
	if o == nil || o.Xpub == nil {
		return nil, false
	}
	return o.Xpub, true
}

// HasXpub returns a boolean if a field has been set.
func (o *XPubTransaction) HasXpub() bool {
	if o != nil && o.Xpub != nil {
		return true
	}

	return false
}

// SetXpub gets a reference to the given string and assigns it to the Xpub field.
func (o *XPubTransaction) SetXpub(v string) {
	o.Xpub = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *XPubTransaction) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *XPubTransaction) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *XPubTransaction) SetTxid(v string) {
	o.Txid = &v
}

// GetMaxReceiveIndex returns the MaxReceiveIndex field value if set, zero value otherwise.
func (o *XPubTransaction) GetMaxReceiveIndex() int32 {
	if o == nil || o.MaxReceiveIndex == nil {
		var ret int32
		return ret
	}
	return *o.MaxReceiveIndex
}

// GetMaxReceiveIndexOk returns a tuple with the MaxReceiveIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetMaxReceiveIndexOk() (*int32, bool) {
	if o == nil || o.MaxReceiveIndex == nil {
		return nil, false
	}
	return o.MaxReceiveIndex, true
}

// HasMaxReceiveIndex returns a boolean if a field has been set.
func (o *XPubTransaction) HasMaxReceiveIndex() bool {
	if o != nil && o.MaxReceiveIndex != nil {
		return true
	}

	return false
}

// SetMaxReceiveIndex gets a reference to the given int32 and assigns it to the MaxReceiveIndex field.
func (o *XPubTransaction) SetMaxReceiveIndex(v int32) {
	o.MaxReceiveIndex = &v
}

// GetMaxChangeIndex returns the MaxChangeIndex field value if set, zero value otherwise.
func (o *XPubTransaction) GetMaxChangeIndex() int32 {
	if o == nil || o.MaxChangeIndex == nil {
		var ret int32
		return ret
	}
	return *o.MaxChangeIndex
}

// GetMaxChangeIndexOk returns a tuple with the MaxChangeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetMaxChangeIndexOk() (*int32, bool) {
	if o == nil || o.MaxChangeIndex == nil {
		return nil, false
	}
	return o.MaxChangeIndex, true
}

// HasMaxChangeIndex returns a boolean if a field has been set.
func (o *XPubTransaction) HasMaxChangeIndex() bool {
	if o != nil && o.MaxChangeIndex != nil {
		return true
	}

	return false
}

// SetMaxChangeIndex gets a reference to the given int32 and assigns it to the MaxChangeIndex field.
func (o *XPubTransaction) SetMaxChangeIndex(v int32) {
	o.MaxChangeIndex = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *XPubTransaction) GetIncome() int64 {
	if o == nil || o.Income == nil {
		var ret int64
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetIncomeOk() (*int64, bool) {
	if o == nil || o.Income == nil {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *XPubTransaction) HasIncome() bool {
	if o != nil && o.Income != nil {
		return true
	}

	return false
}

// SetIncome gets a reference to the given int64 and assigns it to the Income field.
func (o *XPubTransaction) SetIncome(v int64) {
	o.Income = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *XPubTransaction) GetOutcome() int64 {
	if o == nil || o.Outcome == nil {
		var ret int64
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetOutcomeOk() (*int64, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *XPubTransaction) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given int64 and assigns it to the Outcome field.
func (o *XPubTransaction) SetOutcome(v int64) {
	o.Outcome = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *XPubTransaction) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *XPubTransaction) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *XPubTransaction) SetHeight(v int64) {
	o.Height = &v
}

// GetBlockIndex returns the BlockIndex field value if set, zero value otherwise.
func (o *XPubTransaction) GetBlockIndex() int32 {
	if o == nil || o.BlockIndex == nil {
		var ret int32
		return ret
	}
	return *o.BlockIndex
}

// GetBlockIndexOk returns a tuple with the BlockIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetBlockIndexOk() (*int32, bool) {
	if o == nil || o.BlockIndex == nil {
		return nil, false
	}
	return o.BlockIndex, true
}

// HasBlockIndex returns a boolean if a field has been set.
func (o *XPubTransaction) HasBlockIndex() bool {
	if o != nil && o.BlockIndex != nil {
		return true
	}

	return false
}

// SetBlockIndex gets a reference to the given int32 and assigns it to the BlockIndex field.
func (o *XPubTransaction) SetBlockIndex(v int32) {
	o.BlockIndex = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *XPubTransaction) GetBlockTime() int64 {
	if o == nil || o.BlockTime == nil {
		var ret int64
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetBlockTimeOk() (*int64, bool) {
	if o == nil || o.BlockTime == nil {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *XPubTransaction) HasBlockTime() bool {
	if o != nil && o.BlockTime != nil {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given int64 and assigns it to the BlockTime field.
func (o *XPubTransaction) SetBlockTime(v int64) {
	o.BlockTime = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *XPubTransaction) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XPubTransaction) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *XPubTransaction) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *XPubTransaction) SetFlag(v string) {
	o.Flag = &v
}

func (o XPubTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Xpub != nil {
		toSerialize["xpub"] = o.Xpub
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.MaxReceiveIndex != nil {
		toSerialize["maxReceiveIndex"] = o.MaxReceiveIndex
	}
	if o.MaxChangeIndex != nil {
		toSerialize["maxChangeIndex"] = o.MaxChangeIndex
	}
	if o.Income != nil {
		toSerialize["income"] = o.Income
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.BlockIndex != nil {
		toSerialize["blockIndex"] = o.BlockIndex
	}
	if o.BlockTime != nil {
		toSerialize["blockTime"] = o.BlockTime
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableXPubTransaction struct {
	value *XPubTransaction
	isSet bool
}

func (v NullableXPubTransaction) Get() *XPubTransaction {
	return v.value
}

func (v *NullableXPubTransaction) Set(val *XPubTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableXPubTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableXPubTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXPubTransaction(val *XPubTransaction) *NullableXPubTransaction {
	return &NullableXPubTransaction{value: val, isSet: true}
}

func (v NullableXPubTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXPubTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

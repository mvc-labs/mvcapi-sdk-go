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

// InvalidBroadcastCollide This field is only present in case of double spent transaction and contains transactions we collided with
type InvalidBroadcastCollide struct {
	// The transaction id
	Txid *string `json:"txid,omitempty"`
	// Transaction size in bytes
	Size *int32 `json:"size,omitempty"`
	// Whole transaction in hex
	Hex *string `json:"hex,omitempty"`
}

// NewInvalidBroadcastCollide instantiates a new InvalidBroadcastCollide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidBroadcastCollide() *InvalidBroadcastCollide {
	this := InvalidBroadcastCollide{}
	return &this
}

// NewInvalidBroadcastCollideWithDefaults instantiates a new InvalidBroadcastCollide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidBroadcastCollideWithDefaults() *InvalidBroadcastCollide {
	this := InvalidBroadcastCollide{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *InvalidBroadcastCollide) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadcastCollide) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *InvalidBroadcastCollide) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *InvalidBroadcastCollide) SetTxid(v string) {
	o.Txid = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InvalidBroadcastCollide) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadcastCollide) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InvalidBroadcastCollide) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *InvalidBroadcastCollide) SetSize(v int32) {
	o.Size = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *InvalidBroadcastCollide) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadcastCollide) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *InvalidBroadcastCollide) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *InvalidBroadcastCollide) SetHex(v string) {
	o.Hex = &v
}

func (o InvalidBroadcastCollide) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Hex != nil {
		toSerialize["hex"] = o.Hex
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidBroadcastCollide struct {
	value *InvalidBroadcastCollide
	isSet bool
}

func (v NullableInvalidBroadcastCollide) Get() *InvalidBroadcastCollide {
	return v.value
}

func (v *NullableInvalidBroadcastCollide) Set(val *InvalidBroadcastCollide) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidBroadcastCollide) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidBroadcastCollide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidBroadcastCollide(val *InvalidBroadcastCollide) *NullableInvalidBroadcastCollide {
	return &NullableInvalidBroadcastCollide{value: val, isSet: true}
}

func (v NullableInvalidBroadcastCollide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidBroadcastCollide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

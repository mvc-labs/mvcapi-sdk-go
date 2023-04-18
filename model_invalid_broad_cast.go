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

// InvalidBroadCast Invalid transactions detected during validation (if there are any)
type InvalidBroadCast struct {
	// return txid if broadcast success
	Txid *string `json:"txid,omitempty"`
	// return messages if broadcast failed
	RejectCode *int32 `json:"reject_code,omitempty"`
	// return messages if broadcast failed
	RejectReason *string `json:"reject_reason,omitempty"`
	// return messages if broadcast failed
	CollideWith *[]InvalidBroadcastCollide `json:"collideWith,omitempty"`
}

// NewInvalidBroadCast instantiates a new InvalidBroadCast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidBroadCast() *InvalidBroadCast {
	this := InvalidBroadCast{}
	return &this
}

// NewInvalidBroadCastWithDefaults instantiates a new InvalidBroadCast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidBroadCastWithDefaults() *InvalidBroadCast {
	this := InvalidBroadCast{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *InvalidBroadCast) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadCast) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *InvalidBroadCast) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *InvalidBroadCast) SetTxid(v string) {
	o.Txid = &v
}

// GetRejectCode returns the RejectCode field value if set, zero value otherwise.
func (o *InvalidBroadCast) GetRejectCode() int32 {
	if o == nil || o.RejectCode == nil {
		var ret int32
		return ret
	}
	return *o.RejectCode
}

// GetRejectCodeOk returns a tuple with the RejectCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadCast) GetRejectCodeOk() (*int32, bool) {
	if o == nil || o.RejectCode == nil {
		return nil, false
	}
	return o.RejectCode, true
}

// HasRejectCode returns a boolean if a field has been set.
func (o *InvalidBroadCast) HasRejectCode() bool {
	if o != nil && o.RejectCode != nil {
		return true
	}

	return false
}

// SetRejectCode gets a reference to the given int32 and assigns it to the RejectCode field.
func (o *InvalidBroadCast) SetRejectCode(v int32) {
	o.RejectCode = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *InvalidBroadCast) GetRejectReason() string {
	if o == nil || o.RejectReason == nil {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadCast) GetRejectReasonOk() (*string, bool) {
	if o == nil || o.RejectReason == nil {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *InvalidBroadCast) HasRejectReason() bool {
	if o != nil && o.RejectReason != nil {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *InvalidBroadCast) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetCollideWith returns the CollideWith field value if set, zero value otherwise.
func (o *InvalidBroadCast) GetCollideWith() []InvalidBroadcastCollide {
	if o == nil || o.CollideWith == nil {
		var ret []InvalidBroadcastCollide
		return ret
	}
	return *o.CollideWith
}

// GetCollideWithOk returns a tuple with the CollideWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidBroadCast) GetCollideWithOk() (*[]InvalidBroadcastCollide, bool) {
	if o == nil || o.CollideWith == nil {
		return nil, false
	}
	return o.CollideWith, true
}

// HasCollideWith returns a boolean if a field has been set.
func (o *InvalidBroadCast) HasCollideWith() bool {
	if o != nil && o.CollideWith != nil {
		return true
	}

	return false
}

// SetCollideWith gets a reference to the given []InvalidBroadcastCollide and assigns it to the CollideWith field.
func (o *InvalidBroadCast) SetCollideWith(v []InvalidBroadcastCollide) {
	o.CollideWith = &v
}

func (o InvalidBroadCast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.RejectCode != nil {
		toSerialize["reject_code"] = o.RejectCode
	}
	if o.RejectReason != nil {
		toSerialize["reject_reason"] = o.RejectReason
	}
	if o.CollideWith != nil {
		toSerialize["collideWith"] = o.CollideWith
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidBroadCast struct {
	value *InvalidBroadCast
	isSet bool
}

func (v NullableInvalidBroadCast) Get() *InvalidBroadCast {
	return v.value
}

func (v *NullableInvalidBroadCast) Set(val *InvalidBroadCast) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidBroadCast) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidBroadCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidBroadCast(val *InvalidBroadCast) *NullableInvalidBroadCast {
	return &NullableInvalidBroadCast{value: val, isSet: true}
}

func (v NullableInvalidBroadCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidBroadCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

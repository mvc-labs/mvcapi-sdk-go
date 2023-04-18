# InvalidBroadCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | return txid if broadcast success | [optional] 
**RejectCode** | Pointer to **int32** | return messages if broadcast failed | [optional] 
**RejectReason** | Pointer to **string** | return messages if broadcast failed | [optional] 
**CollideWith** | Pointer to [**[]InvalidBroadcastCollide**](InvalidBroadcastCollide.md) | return messages if broadcast failed | [optional] 

## Methods

### NewInvalidBroadCast

`func NewInvalidBroadCast() *InvalidBroadCast`

NewInvalidBroadCast instantiates a new InvalidBroadCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidBroadCastWithDefaults

`func NewInvalidBroadCastWithDefaults() *InvalidBroadCast`

NewInvalidBroadCastWithDefaults instantiates a new InvalidBroadCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *InvalidBroadCast) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *InvalidBroadCast) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *InvalidBroadCast) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *InvalidBroadCast) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetRejectCode

`func (o *InvalidBroadCast) GetRejectCode() int32`

GetRejectCode returns the RejectCode field if non-nil, zero value otherwise.

### GetRejectCodeOk

`func (o *InvalidBroadCast) GetRejectCodeOk() (*int32, bool)`

GetRejectCodeOk returns a tuple with the RejectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectCode

`func (o *InvalidBroadCast) SetRejectCode(v int32)`

SetRejectCode sets RejectCode field to given value.

### HasRejectCode

`func (o *InvalidBroadCast) HasRejectCode() bool`

HasRejectCode returns a boolean if a field has been set.

### GetRejectReason

`func (o *InvalidBroadCast) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *InvalidBroadCast) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *InvalidBroadCast) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *InvalidBroadCast) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetCollideWith

`func (o *InvalidBroadCast) GetCollideWith() []InvalidBroadcastCollide`

GetCollideWith returns the CollideWith field if non-nil, zero value otherwise.

### GetCollideWithOk

`func (o *InvalidBroadCast) GetCollideWithOk() (*[]InvalidBroadcastCollide, bool)`

GetCollideWithOk returns a tuple with the CollideWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollideWith

`func (o *InvalidBroadCast) SetCollideWith(v []InvalidBroadcastCollide)`

SetCollideWith sets CollideWith field to given value.

### HasCollideWith

`func (o *InvalidBroadCast) HasCollideWith() bool`

HasCollideWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



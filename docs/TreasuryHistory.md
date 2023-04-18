# TreasuryHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | treasury transaction txid. | [optional] 
**Outcome** | Pointer to **int64** | treasury transaction spent outcome in satoshi. | [optional] 
**Timestamp** | Pointer to **int64** | treasury transaction timestamp. | [optional] 
**Reason** | Pointer to **string** | Reason for treasury transaction. | [optional] 
**AnnouncementUrl** | Pointer to **string** | Announcement link for treasury transaction. | [optional] 

## Methods

### NewTreasuryHistory

`func NewTreasuryHistory() *TreasuryHistory`

NewTreasuryHistory instantiates a new TreasuryHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreasuryHistoryWithDefaults

`func NewTreasuryHistoryWithDefaults() *TreasuryHistory`

NewTreasuryHistoryWithDefaults instantiates a new TreasuryHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *TreasuryHistory) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *TreasuryHistory) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *TreasuryHistory) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *TreasuryHistory) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetOutcome

`func (o *TreasuryHistory) GetOutcome() int64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TreasuryHistory) GetOutcomeOk() (*int64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TreasuryHistory) SetOutcome(v int64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TreasuryHistory) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetTimestamp

`func (o *TreasuryHistory) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TreasuryHistory) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TreasuryHistory) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TreasuryHistory) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReason

`func (o *TreasuryHistory) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TreasuryHistory) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TreasuryHistory) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TreasuryHistory) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAnnouncementUrl

`func (o *TreasuryHistory) GetAnnouncementUrl() string`

GetAnnouncementUrl returns the AnnouncementUrl field if non-nil, zero value otherwise.

### GetAnnouncementUrlOk

`func (o *TreasuryHistory) GetAnnouncementUrlOk() (*string, bool)`

GetAnnouncementUrlOk returns a tuple with the AnnouncementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementUrl

`func (o *TreasuryHistory) SetAnnouncementUrl(v string)`

SetAnnouncementUrl sets AnnouncementUrl field to given value.

### HasAnnouncementUrl

`func (o *TreasuryHistory) HasAnnouncementUrl() bool`

HasAnnouncementUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



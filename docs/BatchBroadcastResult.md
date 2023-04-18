# BatchBroadcastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Known** | Pointer to **[]string** | Already known transactions detected during processing (if there are any) | [optional] 
**Evicted** | Pointer to **[]string** | Transactions accepted by the mempool and then evicted due to insufficient fee (if there are any) | [optional] 
**Invalid** | Pointer to [**[]InvalidBroadCast**](InvalidBroadCast.md) | Transactions that failed to be accepted by the mempool (if there are any) | [optional] 

## Methods

### NewBatchBroadcastResult

`func NewBatchBroadcastResult() *BatchBroadcastResult`

NewBatchBroadcastResult instantiates a new BatchBroadcastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchBroadcastResultWithDefaults

`func NewBatchBroadcastResultWithDefaults() *BatchBroadcastResult`

NewBatchBroadcastResultWithDefaults instantiates a new BatchBroadcastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnown

`func (o *BatchBroadcastResult) GetKnown() []string`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *BatchBroadcastResult) GetKnownOk() (*[]string, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *BatchBroadcastResult) SetKnown(v []string)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *BatchBroadcastResult) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetEvicted

`func (o *BatchBroadcastResult) GetEvicted() []string`

GetEvicted returns the Evicted field if non-nil, zero value otherwise.

### GetEvictedOk

`func (o *BatchBroadcastResult) GetEvictedOk() (*[]string, bool)`

GetEvictedOk returns a tuple with the Evicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvicted

`func (o *BatchBroadcastResult) SetEvicted(v []string)`

SetEvicted sets Evicted field to given value.

### HasEvicted

`func (o *BatchBroadcastResult) HasEvicted() bool`

HasEvicted returns a boolean if a field has been set.

### GetInvalid

`func (o *BatchBroadcastResult) GetInvalid() []InvalidBroadCast`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *BatchBroadcastResult) GetInvalidOk() (*[]InvalidBroadCast, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *BatchBroadcastResult) SetInvalid(v []InvalidBroadCast)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *BatchBroadcastResult) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



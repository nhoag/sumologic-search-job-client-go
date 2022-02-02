# SearchJobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Search job state. | [optional] 
**MessageCount** | Pointer to **int32** | The number of raw messages found so far. | [optional] 
**HistogramBuckets** | Pointer to [**[]HistogramBucket**](HistogramBucket.md) | A list of histogram buckets. | [optional] 
**PendingErrors** | Pointer to **[]interface{}** |  | [optional] 
**PendingWarnings** | Pointer to **[]interface{}** |  | [optional] 
**RecordCount** | Pointer to **int32** | The number of aggregated records produced so far. | [optional] 

## Methods

### NewSearchJobState

`func NewSearchJobState() *SearchJobState`

NewSearchJobState instantiates a new SearchJobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchJobStateWithDefaults

`func NewSearchJobStateWithDefaults() *SearchJobState`

NewSearchJobStateWithDefaults instantiates a new SearchJobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SearchJobState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SearchJobState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SearchJobState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SearchJobState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageCount

`func (o *SearchJobState) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *SearchJobState) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *SearchJobState) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *SearchJobState) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetHistogramBuckets

`func (o *SearchJobState) GetHistogramBuckets() []HistogramBucket`

GetHistogramBuckets returns the HistogramBuckets field if non-nil, zero value otherwise.

### GetHistogramBucketsOk

`func (o *SearchJobState) GetHistogramBucketsOk() (*[]HistogramBucket, bool)`

GetHistogramBucketsOk returns a tuple with the HistogramBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBuckets

`func (o *SearchJobState) SetHistogramBuckets(v []HistogramBucket)`

SetHistogramBuckets sets HistogramBuckets field to given value.

### HasHistogramBuckets

`func (o *SearchJobState) HasHistogramBuckets() bool`

HasHistogramBuckets returns a boolean if a field has been set.

### GetPendingErrors

`func (o *SearchJobState) GetPendingErrors() []interface{}`

GetPendingErrors returns the PendingErrors field if non-nil, zero value otherwise.

### GetPendingErrorsOk

`func (o *SearchJobState) GetPendingErrorsOk() (*[]interface{}, bool)`

GetPendingErrorsOk returns a tuple with the PendingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingErrors

`func (o *SearchJobState) SetPendingErrors(v []interface{})`

SetPendingErrors sets PendingErrors field to given value.

### HasPendingErrors

`func (o *SearchJobState) HasPendingErrors() bool`

HasPendingErrors returns a boolean if a field has been set.

### GetPendingWarnings

`func (o *SearchJobState) GetPendingWarnings() []interface{}`

GetPendingWarnings returns the PendingWarnings field if non-nil, zero value otherwise.

### GetPendingWarningsOk

`func (o *SearchJobState) GetPendingWarningsOk() (*[]interface{}, bool)`

GetPendingWarningsOk returns a tuple with the PendingWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWarnings

`func (o *SearchJobState) SetPendingWarnings(v []interface{})`

SetPendingWarnings sets PendingWarnings field to given value.

### HasPendingWarnings

`func (o *SearchJobState) HasPendingWarnings() bool`

HasPendingWarnings returns a boolean if a field has been set.

### GetRecordCount

`func (o *SearchJobState) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *SearchJobState) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *SearchJobState) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *SearchJobState) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



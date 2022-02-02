# HistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int32** | Width of the bucket in milliseconds. | [optional] 
**Count** | Pointer to **int32** | Number of messages in the bucket. | [optional] 
**StartTimestamp** | Pointer to **int32** | The start timestamp of the bucket in milliseconds. | [optional] 

## Methods

### NewHistogramBucket

`func NewHistogramBucket() *HistogramBucket`

NewHistogramBucket instantiates a new HistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistogramBucketWithDefaults

`func NewHistogramBucketWithDefaults() *HistogramBucket`

NewHistogramBucketWithDefaults instantiates a new HistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *HistogramBucket) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *HistogramBucket) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *HistogramBucket) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *HistogramBucket) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetCount

`func (o *HistogramBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HistogramBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HistogramBucket) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HistogramBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *HistogramBucket) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *HistogramBucket) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *HistogramBucket) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *HistogramBucket) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
Sumo Logic API

The Search Job API provides third-party scripts and applications access to your log data. The API follows Representational State Transfer (REST) patterns and is optimized for ease of use and consistency. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// HistogramBucket Vertical slice for time series data.
type HistogramBucket struct {
	// Width of the bucket in milliseconds.
	Length *int32 `json:"length,omitempty"`
	// Number of messages in the bucket.
	Count *int32 `json:"count,omitempty"`
	// The start timestamp of the bucket in milliseconds.
	StartTimestamp *float32 `json:"startTimestamp,omitempty"`
}

// NewHistogramBucket instantiates a new HistogramBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistogramBucket() *HistogramBucket {
	this := HistogramBucket{}
	return &this
}

// NewHistogramBucketWithDefaults instantiates a new HistogramBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistogramBucketWithDefaults() *HistogramBucket {
	this := HistogramBucket{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *HistogramBucket) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramBucket) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *HistogramBucket) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *HistogramBucket) SetLength(v int32) {
	o.Length = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *HistogramBucket) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramBucket) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *HistogramBucket) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *HistogramBucket) SetCount(v int32) {
	o.Count = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *HistogramBucket) GetStartTimestamp() float32 {
	if o == nil || o.StartTimestamp == nil {
		var ret float32
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistogramBucket) GetStartTimestampOk() (*float32, bool) {
	if o == nil || o.StartTimestamp == nil {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *HistogramBucket) HasStartTimestamp() bool {
	if o != nil && o.StartTimestamp != nil {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given float32 and assigns it to the StartTimestamp field.
func (o *HistogramBucket) SetStartTimestamp(v float32) {
	o.StartTimestamp = &v
}

func (o HistogramBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.StartTimestamp != nil {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableHistogramBucket struct {
	value *HistogramBucket
	isSet bool
}

func (v NullableHistogramBucket) Get() *HistogramBucket {
	return v.value
}

func (v *NullableHistogramBucket) Set(val *HistogramBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableHistogramBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableHistogramBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistogramBucket(val *HistogramBucket) *NullableHistogramBucket {
	return &NullableHistogramBucket{value: val, isSet: true}
}

func (v NullableHistogramBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistogramBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



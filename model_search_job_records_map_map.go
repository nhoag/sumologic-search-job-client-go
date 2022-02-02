/*
Sumo Logic API

YOLO! 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SearchJobRecordsMapMap struct for SearchJobRecordsMapMap
type SearchJobRecordsMapMap struct {
	Count *string `json:"_count,omitempty"`
	SourceCategory *string `json:"_sourceCategory,omitempty"`
}

// NewSearchJobRecordsMapMap instantiates a new SearchJobRecordsMapMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchJobRecordsMapMap() *SearchJobRecordsMapMap {
	this := SearchJobRecordsMapMap{}
	return &this
}

// NewSearchJobRecordsMapMapWithDefaults instantiates a new SearchJobRecordsMapMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchJobRecordsMapMapWithDefaults() *SearchJobRecordsMapMap {
	this := SearchJobRecordsMapMap{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchJobRecordsMapMap) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobRecordsMapMap) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchJobRecordsMapMap) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *SearchJobRecordsMapMap) SetCount(v string) {
	o.Count = &v
}

// GetSourceCategory returns the SourceCategory field value if set, zero value otherwise.
func (o *SearchJobRecordsMapMap) GetSourceCategory() string {
	if o == nil || o.SourceCategory == nil {
		var ret string
		return ret
	}
	return *o.SourceCategory
}

// GetSourceCategoryOk returns a tuple with the SourceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobRecordsMapMap) GetSourceCategoryOk() (*string, bool) {
	if o == nil || o.SourceCategory == nil {
		return nil, false
	}
	return o.SourceCategory, true
}

// HasSourceCategory returns a boolean if a field has been set.
func (o *SearchJobRecordsMapMap) HasSourceCategory() bool {
	if o != nil && o.SourceCategory != nil {
		return true
	}

	return false
}

// SetSourceCategory gets a reference to the given string and assigns it to the SourceCategory field.
func (o *SearchJobRecordsMapMap) SetSourceCategory(v string) {
	o.SourceCategory = &v
}

func (o SearchJobRecordsMapMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["_count"] = o.Count
	}
	if o.SourceCategory != nil {
		toSerialize["_sourceCategory"] = o.SourceCategory
	}
	return json.Marshal(toSerialize)
}

type NullableSearchJobRecordsMapMap struct {
	value *SearchJobRecordsMapMap
	isSet bool
}

func (v NullableSearchJobRecordsMapMap) Get() *SearchJobRecordsMapMap {
	return v.value
}

func (v *NullableSearchJobRecordsMapMap) Set(val *SearchJobRecordsMapMap) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchJobRecordsMapMap) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchJobRecordsMapMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchJobRecordsMapMap(val *SearchJobRecordsMapMap) *NullableSearchJobRecordsMapMap {
	return &NullableSearchJobRecordsMapMap{value: val, isSet: true}
}

func (v NullableSearchJobRecordsMapMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchJobRecordsMapMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// SearchJobRecords blah
type SearchJobRecords struct {
	Fields []SearchJobField `json:"fields,omitempty"`
	Records []SearchJobRecordsMap `json:"records,omitempty"`
}

// NewSearchJobRecords instantiates a new SearchJobRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchJobRecords() *SearchJobRecords {
	this := SearchJobRecords{}
	return &this
}

// NewSearchJobRecordsWithDefaults instantiates a new SearchJobRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchJobRecordsWithDefaults() *SearchJobRecords {
	this := SearchJobRecords{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SearchJobRecords) GetFields() []SearchJobField {
	if o == nil || o.Fields == nil {
		var ret []SearchJobField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobRecords) GetFieldsOk() ([]SearchJobField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SearchJobRecords) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []SearchJobField and assigns it to the Fields field.
func (o *SearchJobRecords) SetFields(v []SearchJobField) {
	o.Fields = v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SearchJobRecords) GetRecords() []SearchJobRecordsMap {
	if o == nil || o.Records == nil {
		var ret []SearchJobRecordsMap
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobRecords) GetRecordsOk() ([]SearchJobRecordsMap, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SearchJobRecords) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []SearchJobRecordsMap and assigns it to the Records field.
func (o *SearchJobRecords) SetRecords(v []SearchJobRecordsMap) {
	o.Records = v
}

func (o SearchJobRecords) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableSearchJobRecords struct {
	value *SearchJobRecords
	isSet bool
}

func (v NullableSearchJobRecords) Get() *SearchJobRecords {
	return v.value
}

func (v *NullableSearchJobRecords) Set(val *SearchJobRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchJobRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchJobRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchJobRecords(val *SearchJobRecords) *NullableSearchJobRecords {
	return &NullableSearchJobRecords{value: val, isSet: true}
}

func (v NullableSearchJobRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchJobRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// SearchJobField Search Job field.
type SearchJobField struct {
	Name *string `json:"name,omitempty"`
	FieldType *string `json:"fieldType,omitempty"`
	KeyField *bool `json:"keyField,omitempty"`
}

// NewSearchJobField instantiates a new SearchJobField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchJobField() *SearchJobField {
	this := SearchJobField{}
	return &this
}

// NewSearchJobFieldWithDefaults instantiates a new SearchJobField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchJobFieldWithDefaults() *SearchJobField {
	this := SearchJobField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchJobField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchJobField) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchJobField) SetName(v string) {
	o.Name = &v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *SearchJobField) GetFieldType() string {
	if o == nil || o.FieldType == nil {
		var ret string
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobField) GetFieldTypeOk() (*string, bool) {
	if o == nil || o.FieldType == nil {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *SearchJobField) HasFieldType() bool {
	if o != nil && o.FieldType != nil {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given string and assigns it to the FieldType field.
func (o *SearchJobField) SetFieldType(v string) {
	o.FieldType = &v
}

// GetKeyField returns the KeyField field value if set, zero value otherwise.
func (o *SearchJobField) GetKeyField() bool {
	if o == nil || o.KeyField == nil {
		var ret bool
		return ret
	}
	return *o.KeyField
}

// GetKeyFieldOk returns a tuple with the KeyField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobField) GetKeyFieldOk() (*bool, bool) {
	if o == nil || o.KeyField == nil {
		return nil, false
	}
	return o.KeyField, true
}

// HasKeyField returns a boolean if a field has been set.
func (o *SearchJobField) HasKeyField() bool {
	if o != nil && o.KeyField != nil {
		return true
	}

	return false
}

// SetKeyField gets a reference to the given bool and assigns it to the KeyField field.
func (o *SearchJobField) SetKeyField(v bool) {
	o.KeyField = &v
}

func (o SearchJobField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FieldType != nil {
		toSerialize["fieldType"] = o.FieldType
	}
	if o.KeyField != nil {
		toSerialize["keyField"] = o.KeyField
	}
	return json.Marshal(toSerialize)
}

type NullableSearchJobField struct {
	value *SearchJobField
	isSet bool
}

func (v NullableSearchJobField) Get() *SearchJobField {
	return v.value
}

func (v *NullableSearchJobField) Set(val *SearchJobField) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchJobField) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchJobField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchJobField(val *SearchJobField) *NullableSearchJobField {
	return &NullableSearchJobField{value: val, isSet: true}
}

func (v NullableSearchJobField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchJobField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


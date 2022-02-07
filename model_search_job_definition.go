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

// SearchJobDefinition Search job definition.
type SearchJobDefinition struct {
	// Search query.
	Query *string `json:"query,omitempty"`
	// Search start time.
	From *string `json:"from,omitempty"`
	// Search end time.
	To *string `json:"to,omitempty"`
	// Timezone context for search.
	TimeZone *string `json:"timeZone,omitempty"`
	// Use the receipt time rather than timestamps log messages.
	ByReceiptTime *bool `json:"byReceiptTime,omitempty"`
}

// NewSearchJobDefinition instantiates a new SearchJobDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchJobDefinition() *SearchJobDefinition {
	this := SearchJobDefinition{}
	var byReceiptTime bool = false
	this.ByReceiptTime = &byReceiptTime
	return &this
}

// NewSearchJobDefinitionWithDefaults instantiates a new SearchJobDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchJobDefinitionWithDefaults() *SearchJobDefinition {
	this := SearchJobDefinition{}
	var byReceiptTime bool = false
	this.ByReceiptTime = &byReceiptTime
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchJobDefinition) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobDefinition) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchJobDefinition) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SearchJobDefinition) SetQuery(v string) {
	o.Query = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SearchJobDefinition) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobDefinition) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SearchJobDefinition) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SearchJobDefinition) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SearchJobDefinition) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobDefinition) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SearchJobDefinition) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SearchJobDefinition) SetTo(v string) {
	o.To = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *SearchJobDefinition) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobDefinition) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *SearchJobDefinition) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *SearchJobDefinition) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetByReceiptTime returns the ByReceiptTime field value if set, zero value otherwise.
func (o *SearchJobDefinition) GetByReceiptTime() bool {
	if o == nil || o.ByReceiptTime == nil {
		var ret bool
		return ret
	}
	return *o.ByReceiptTime
}

// GetByReceiptTimeOk returns a tuple with the ByReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchJobDefinition) GetByReceiptTimeOk() (*bool, bool) {
	if o == nil || o.ByReceiptTime == nil {
		return nil, false
	}
	return o.ByReceiptTime, true
}

// HasByReceiptTime returns a boolean if a field has been set.
func (o *SearchJobDefinition) HasByReceiptTime() bool {
	if o != nil && o.ByReceiptTime != nil {
		return true
	}

	return false
}

// SetByReceiptTime gets a reference to the given bool and assigns it to the ByReceiptTime field.
func (o *SearchJobDefinition) SetByReceiptTime(v bool) {
	o.ByReceiptTime = &v
}

func (o SearchJobDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.ByReceiptTime != nil {
		toSerialize["byReceiptTime"] = o.ByReceiptTime
	}
	return json.Marshal(toSerialize)
}

type NullableSearchJobDefinition struct {
	value *SearchJobDefinition
	isSet bool
}

func (v NullableSearchJobDefinition) Get() *SearchJobDefinition {
	return v.value
}

func (v *NullableSearchJobDefinition) Set(val *SearchJobDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchJobDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchJobDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchJobDefinition(val *SearchJobDefinition) *NullableSearchJobDefinition {
	return &NullableSearchJobDefinition{value: val, isSet: true}
}

func (v NullableSearchJobDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchJobDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



# SearchJobRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]SearchJobField**](SearchJobField.md) |  | [optional] 
**Records** | Pointer to [**[]SearchJobRecordsMap**](SearchJobRecordsMap.md) |  | [optional] 

## Methods

### NewSearchJobRecords

`func NewSearchJobRecords() *SearchJobRecords`

NewSearchJobRecords instantiates a new SearchJobRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchJobRecordsWithDefaults

`func NewSearchJobRecordsWithDefaults() *SearchJobRecords`

NewSearchJobRecordsWithDefaults instantiates a new SearchJobRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SearchJobRecords) GetFields() []SearchJobField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchJobRecords) GetFieldsOk() (*[]SearchJobField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchJobRecords) SetFields(v []SearchJobField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchJobRecords) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetRecords

`func (o *SearchJobRecords) GetRecords() []SearchJobRecordsMap`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SearchJobRecords) GetRecordsOk() (*[]SearchJobRecordsMap, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SearchJobRecords) SetRecords(v []SearchJobRecordsMap)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SearchJobRecords) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


